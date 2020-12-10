---
title: "Bootstrap Helper"
date: 2020-12-09T14:46:32+08:00
draft: false
---


### Clearfix

Quickly and easily clear floated content within a container by adding a clearfix utility.

Easily clear floats by adding .clearfix to the parent element. Can also be used as a mixin.

```html
<div class="clearfix">...</div>
```

The mixin source code:
```css
@mixin clearfix() {
  &::after {
    display: block;
    clear: both;
    content: "";
  }
}
```

Use the mixin in SCSS:
```scss
.element {
  @include clearfix;
}
```


### Colored links

You can use the .link-* classes to colorize links. Unlike the .text-* classes, these classes have a :hover and :focus state.

```
<a href="#" class="link-primary">Primary link</a>
<a href="#" class="link-secondary">Secondary link</a>
<a href="#" class="link-success">Success link</a>
<a href="#" class="link-danger">Danger link</a>
<a href="#" class="link-warning">Warning link</a>
<a href="#" class="link-info">Info link</a>
<a href="#" class="link-light">Light link</a>
<a href="#" class="link-dark">Dark link</a>
```

### Ratios

Use generated pseudo elements to make an element maintain the aspect ratio of your choosing. Perfect for responsively handling video or slideshow embeds based on the width of the parent.

Use the ratio helper to manage the aspect ratios of external content like <iframe>s, <embed>s, <video>s, and <object>s. These helpers also can be used on any standard HTML child element (e.g., a <div> or <img>). Styles are applied from the parent .ratio class directly to the child.

Aspect ratios are declared in a Sass map and included in each class via CSS variable, which also allows custom aspect ratios.

Example

```html

<div class="ratio ratio-16x9">
  <iframe src="https://www.youtube.com/embed/zpOULjyy-n8?rel=0" title="YouTube video" allowfullscreen></iframe>
</div>

```

```css
.ratio-16x9 {
    --aspect-ratio: calc(9 / 16 * 100%);
}
```

```html
<div class="ratio ratio-1x1">
  <div>1x1</div>
</div>
<div class="ratio ratio-4x3">
  <div>4x3</div>
</div>
<div class="ratio ratio-16x9">
  <div>16x9</div>
</div>
<div class="ratio ratio-21x9">
  <div>21x9</div>
</div>

<div class="ratio" style="--aspect-ratio: 50%;">
  <div>2x1</div>
</div>
```

Sass map
Within _variables.scss, you can change the aspect ratios you want to use. Here’s our default $ratio-aspect-ratios map. Modify the map as you like and recompile your Sass to put them to use.

```scss
$aspect-ratios: (
  "1x1": 100%,
  "4x3": calc(3 / 4 * 100%),
  "16x9": calc(9 / 16 * 100%),
  "21x9": calc(9 / 21 * 100%)
);
```

### Position

Use these helpers for quickly configuring the position of an element.


- Fixed top

Position an element at the top of the viewport, from edge to edge. Be sure you understand the ramifications of fixed position in your project; you may need to add additional CSS.

```html
<div class="fixed-top">...</div>
```
- Fixed bottom

Position an element at the bottom of the viewport, from edge to edge. Be sure you understand the ramifications of fixed position in your project; you may need to add additional CSS.

```html
<div class="fixed-bottom">...</div>
```

- Sticky top

Position an element at the top of the viewport, from edge to edge, but only after you scroll past it. The .sticky-top utility uses CSS’s position: sticky, which isn’t fully supported in all browsers.

```html
<div class="sticky-top">...</div>
```

- Responsive sticky top

Responsive variations also exist for .sticky-top utility.

```html
<div class="sticky-sm-top">Stick to the top on viewports sized SM (small) or wider</div>
<div class="sticky-md-top">Stick to the top on viewports sized MD (medium) or wider</div>
<div class="sticky-lg-top">Stick to the top on viewports sized LG (large) or wider</div>
<div class="sticky-xl-top">Stick to the top on viewports sized XL (extra-large) or wider</div>
```

### Visually hidden

Use these helpers to visually hide elements but keep them accessible to assistive technologies.

Visually hide an element while still allowing it to be exposed to assistive technologies (such as screen readers) with .visually-hidden. Use .visually-hidden-focusable to visually hide an element by default, but to display it when it’s focused (e.g. by a keyboard-only user). Can also be used as mixins.

```html
<h2 class="visually-hidden">Title for screen readers</h2>
<a class="visually-hidden-focusable" href="#content">Skip to main content</a>
```

```scss
.visually-hidden-title {
  @include visually-hidden;
}

.skip-navigation {
  @include visually-hidden-focusable;
}
```

### Stretched link
Make any HTML element or Bootstrap component clickable by “stretching” a nested link via CSS.

Add .stretched-link to a link to make its containing block clickable via a ::after pseudo element. In most cases, this means that an element with position: relative; that contains a link with the .stretched-link class is clickable.

Cards have position: relative by default in Bootstrap, so in this case you can safely add the .stretched-link class to a link in the card without any other HTML changes.

Multiple links and tap targets are not recommended with stretched links. However, some position and z-index styles can help should this be required.

```html
<div class="card" style="width: 18rem;">
  <img src="..." class="card-img-top" alt="...">
  <div class="card-body">
    <h5 class="card-title">Card with stretched link</h5>
    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
    <a href="#" class="btn btn-primary stretched-link">Go somewhere</a>
  </div>
</div>
```

Most custom components do not have position: relative by default, so we need to add the .position-relative here to prevent the link from stretching outside the parent element.

```html
<div class="d-flex position-relative">
  <img src="..." class="flex-shrink-0 me-3" alt="...">
  <div>
    <h5 class="mt-0">Custom component with stretched link</h5>
    <p>Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.</p>
    <a href="#" class="stretched-link">Go somewhere</a>
  </div>
</div>
```

Identifying the containing block

If the stretched link doesn’t seem to work, the containing block will probably be the cause. The following CSS properties will make an element the containing block:

A position value other than static
A transform or perspective value other than none
A will-change value of transform or perspective
A filter value other than none or a will-change value of filter (only works on Firefox)

```html
<div class="card" style="width: 18rem;">
  <img src="..." class="card-img-top" alt="...">
  <div class="card-body">
    <h5 class="card-title">Card with stretched links</h5>
    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
    <p class="card-text">
      <a href="#" class="stretched-link text-danger" style="position: relative;">Stretched link will not work here, because <code>position: relative</code> is added to the link</a>
    </p>
    <p class="card-text bg-light" style="transform: rotate(0);">
      This <a href="#" class="text-warning stretched-link">stretched link</a> will only be spread over the <code>p</code>-tag, because a transform is applied to it.
    </p>
  </div>
</div>
```

### Text truncation

Truncate long strings of text with an ellipsis.

For longer content, you can add a .text-truncate class to truncate the text with an ellipsis. Requires display: inline-block or display: block.

```html
<!-- Block level -->
<div class="row">
  <div class="col-2 text-truncate">
    Praeterea iter est quasdam res quas ex communi.
  </div>
</div>

<!-- Inline level -->
<span class="d-inline-block text-truncate" style="max-width: 150px;">
  Praeterea iter est quasdam res quas ex communi.
</span>
```