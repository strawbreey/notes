---
title: "Typescript 4"
date: 2020-12-07T22:59:26+08:00
draft: false
---

### Template Literal Types

String literal types in TypeScript allow us to model functions and APIs that expect a set of specific strings.

```ts
function setVerticalAlignment(color: "top" | "middle" | "bottom") {
  // ...
}

setVerticalAlignment("middel");
// Argument of type '"middel"' is not assignable to parameter of type '"top" | "middle" | "bottom"'.

```

This is pretty nice because string literal types can basically spell-check our string values.

We also like that string literals can be used as property names in mapped types. In this sense, they’re also usable as building blocks:

```ts
type Options = {
  [K in "noImplicitAny" | "strictNullChecks" | "strictFunctionTypes"]?: boolean;
};
// same as
//   type Options = {
//       noImplicitAny?: boolean,
//       strictNullChecks?: boolean,
//       strictFunctionTypes?: boolean
//   };
```

But there’s another place that that string literal types could be used as building blocks: building other string literal types.

That’s why TypeScript 4.1 brings the template literal string type. It has the same syntax as template literal strings in JavaScript, but is used in type positions. When you use it with concrete literal types, it produces a new string literal type by concatenating the contents.

```ts
type World = "world";

type Greeting = `hello ${World}`;
//   ^ = type Greeting = "hello world"
```

What happens when you have unions in substitution positions? It produces the set of every possible string literal that could be represented by each union member.

```ts
type Color = "red" | "blue";
type Quantity = "one" | "two";

type SeussFish = `${Quantity | Color} fish`;
//   ^ = type SeussFish = "one fish" | "two fish" | "red fish" | "blue fish"
```

This can be used beyond cute examples in release notes. For example, several libraries for UI components have a way to specify both vertical and horizontal alignment in their APIs, often with both at once using a single string like "bottom-right". Between vertically aligning with "top", "middle", and "bottom", and horizontally aligning with "left", "center", and "right", there are 9 possible strings where each of the former strings is connected with each of the latter strings using a dash.

```ts
type VerticalAlignment = "top" | "middle" | "bottom";
type HorizontalAlignment = "left" | "center" | "right";

// Takes
//   | "top-left"    | "top-center"    | "top-right"
//   | "middle-left" | "middle-center" | "middle-right"
//   | "bottom-left" | "bottom-center" | "bottom-right"

declare function setAlignment(value: `${VerticalAlignment}-${HorizontalAlignment}`): void;

setAlignment("top-left");   // works!
setAlignment("top-middel"); // error!
setAlignment("top-pot");  // error! 
```

While there are lots of examples of this sort of API in the wild, this is still a bit of a toy example since we could write these out manually. In fact, for 9 strings, this is likely fine; but when you need a ton of strings, you should consider automatically generating them ahead of time to save work on every type-check (or just use string, which will be much simpler to comprehend).

Some of the real value comes from dynamically creating new string literals. For example, imagine a makeWatchedObject API that takes an object and produces a mostly identical object, but with a new on method to detect for changes to the properties.

```ts
let person = makeWatchedObject({
  firstName: "Homer",
  age: 42, // give-or-take
  location: "Springfield",
});

person.on("firstNameChanged", () => {
  console.log(`firstName was changed!`);
});
```

Notice that on listens on the event "firstNameChanged", not just "firstName". How would we type this?

```ts
type PropEventSource<T> = {
    on(eventName: `${string & keyof T}Changed`, callback: () => void): void;
};

/// Create a "watched object" with an 'on' method
/// so that you can watch for changes to properties.
declare function makeWatchedObject<T>(obj: T): T & PropEventSource<T>;
```
With this, we can build something that errors when we give the wrong property!

```ts
// error!
person.on("firstName", () => {});
// Argument of type '"firstName"' is not assignable to parameter of type '"firstNameChanged" | "ageChanged" | "locationChanged"'.

// error!
person.on("frstNameChanged", () => {});
// Argument of type '"frstNameChanged"' is not assignable to parameter of type '"firstNameChanged" | "ageChanged" | "locationChanged"'.
```

We can also do something special in template literal types: we can infer from substitution positions. We can make our last example generic to infer from parts of the eventName string to figure out the associated property.

```ts
type PropEventSource<T> = {
    on<K extends string & keyof T>
        (eventName: `${K}Changed`, callback: (newValue: T[K]) => void ): void;
};

declare function makeWatchedObject<T>(obj: T): T & PropEventSource<T>;

let person = makeWatchedObject({
    firstName: "Homer",
    age: 42,
    location: "Springfield",
});

// works! 'newName' is typed as 'string'
person.on("firstNameChanged", newName => {
    // 'newName' has the type of 'firstName'
    console.log(`new name is ${newName.toUpperCase()}`);
});

// works! 'newAge' is typed as 'number'
person.on("ageChanged", newAge => {
    if (newAge < 0) {
        console.log("warning! negative age");
    }
})
```

Here we made on into a generic method. When a user calls with the string "firstNameChanged', TypeScript will try to infer the right type for K. To do that, it will match K against the content prior to "Changed" and infer the string "firstName". Once TypeScript figures that out, the on method can fetch the type of firstName on the original object, which is string in this case. Similarly, when we call with "ageChanged", it finds the type for the property age which is number).

Inference can be combined in different ways, often to deconstruct strings, and reconstruct them in different ways. In fact, to help with modifying these string literal types, we’ve added a few new utility type aliases for modifying casing in letters (i.e. converting to lowercase and uppercase characters).

```ts
type EnthusiasticGreeting<T extends string> = `${Uppercase<T>}`

type HELLO = EnthusiasticGreeting<"hello">;
//   ^ = type HELLO = "HELLO"s
```

The new type aliases are Uppercase, Lowercase, Capitalize and Uncapitalize. The first two transform every character in a string, and the latter two transform only the first character in a string.

For more details, [see the original pull request](https://github.com/microsoft/TypeScript/pull/40336) and [the in-progress pull request to switch to type alias helpers](https://github.com/microsoft/TypeScript/pull/40580).


### Key Remapping in Mapped Types

Just as a refresher, a mapped type can create new object types based on arbitrary keys

```ts
type Options = {
  [K in "noImplicitAny" | "strictNullChecks" | "strictFunctionTypes"]?: boolean;
};
// same as
//   type Options = {
//       noImplicitAny?: boolean,
//       strictNullChecks?: boolean,
//       strictFunctionTypes?: boolean
//   };
```

or new object types based on other object types.

```ts
/// 'Partial<T>' is the same as 'T', but with each property marked optional.
type Partial<T> = {
  [K in keyof T]?: T[K];
};
```

Until now, mapped types could only produce new object types with keys that you provided them; however, lots of the time you want to be able to create new keys, or filter out keys, based on the inputs.

That’s why TypeScript 4.1 allows you to re-map keys in mapped types with a new as clause.

```ts
type MappedTypeWithNewKeys<T> = {
    [K in keyof T as NewKeyType]: T[K]
    //            ^^^^^^^^^^^^^
    //            This is the new syntax!
}
```
With this new as clause, you can leverage features like template literal types to easily create property names based off of old ones.

```ts
type Getters<T> = {
    [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K]
};

interface Person {
    name: string;
    age: number;
    location: string;
}

type LazyPerson = Getters<Person>;
//   ^ = type LazyPerson = {
//       getName: () => string;
//       getAge: () => number;
//       getLocation: () => string;
//   }
```

and you can even filter out keys by producing never. That means you don’t have to use an extra Omit helper type in some cases.

```ts
// Remove the 'kind' property
type RemoveKindField<T> = {
    [K in keyof T as Exclude<K, "kind">]: T[K]
};

interface Circle {
    kind: "circle";
    radius: number;
}

type KindlessCircle = RemoveKindField<Circle>;
//   ^ = type KindlessCircle = {
//       radius: number;
//   }
```

For more information, take a look at [the original pull request over on GitHub](https://github.com/microsoft/TypeScript/pull/40336).

### Recursive (递归) Conditional Types

In JavaScript it’s fairly common to see functions that can flatten and build up container types at arbitrary levels. For example, consider the .then() method on instances of Promise. .then(...) unwraps each promise until it finds a value that’s not “promise-like”, and passes that value to a callback. There’s also a relatively new flat method on Arrays that can take a depth of how deep to flatten.

Expressing this in TypeScript’s type system was, for all practical intents and purposes, not possible. While there were hacks to achieve this, the types ended up looking very unreasonable.

That’s why TypeScript 4.1 eases some restrictions on conditional types - so that they can model these patterns. In TypeScript 4.1, conditional types can now immediately reference themselves within their branches, making it easier to write recursive type aliases.

For example, if we wanted to write a type to get the element types of nested arrays, we could write the following deepFlatten type.

```ts
type ElementType<T> = T extends ReadonlyArray<infer U> ? ElementType<U> : T;

function deepFlatten<T extends readonly unknown[]>(x: T): ElementType<T>[] {
  throw "not implemented";
}

// All of these return the type 'number[]':
deepFlatten([1, 2, 3]);
deepFlatten([[1], [2, 3]]);
deepFlatten([[1], [[2]], [[[3]]]]);
```

[todo]
https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-1.html#recursive-conditional-types