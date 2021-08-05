---
title: "Json Yaml Toml"
date: 2021-08-04T15:06:34+08:00
draft: false
---


互联网的早期，一些非常聪明的家伙决定整合一种让每个系统都能理解的标准语言，并创造性地将其命名为标准通用标记语言 `Standard Generalized Markup Language`（简称 SGML）。SGML 非常灵活，发布者也很好地定义了它。它成为了 XML、SVG 和 HTML 等语言之父。

随后，人们开始看到简洁、易读的数据的好处，这些数据可以在系统之间以编程的方式共享，而开销很小。JSON 诞生了并且能够满足所有的需求。如 CSON，TOML 和 YAML。

### XML

XML 语言非常灵活且易于编写，但它的缺点是冗长，人类难以阅读、计算机非常难以读取，并且有很多语法对于传达信息并不是完全必要的。

```xml
<book id="bk101">
    <author>Gambardella, Matthew</author>
    <title>XML Developer's Guide</title>
    <genre>Computer</genre>
    <price>44.95</price>
    <publish_date>2000-10-01</publish_date>
    <description>An in-depth look at creating applications with XML.</description>
</book>
```

太棒了。易于阅读、理解、写入，也容易编码一个可以读写它的系统。但请考虑这个例子：

```html
<!DOCTYPE r [ <!ENTITY y "a]>b"> ]>
<r>
<a b="&y;>" />
<![CDATA[[a>b <a>b <a]]>
<?x <a> <!-- <b> ?> c --> d
</r>
```

这上面是 100% 有效的 XML。几乎不可能阅读、理解或推理。编写可以使用和理解这个的代码将花费至少 36 根头发和 248 磅咖啡渣。我们没有那么多时间或咖啡，而且我们大多数老程序员们现在都是秃头。所以，让它活在我们的记忆里，就像 css hacks、IE 6 浏览器 和真空管一样好了。

### JSON
JavaScript 对象表示法JavaScript Object Notation，简称 JSON。JSON（读起来像 Jason 这个名字） 是 Brendan Eich 发明的，并且得到了伟大而强力的 JavaScript 意见领袖 Douglas Crockford 的推广。它现在几乎用在任何地方。这种格式很容易由人和机器编写，按规范中的严格规则解析也相当容易，并且灵活 —— 允许深层嵌套数据，支持所有的原始数据类型，及将集合解释为数组或对象。JSON 成为了将数据从一个系统传输到另一个系统的事实标准。几乎所有语言都有内置读写它的功能。

JSON语法很简单。方括号表示数组，花括号表示记录，由冒号分隔的两个值分别表示属性或“键”（在左边）、值（在右边）。所有键必须用双引号括起来：

```json
  {
    "books": [
      {
        "id": "bk102",
        "author": "Crockford, Douglas",
        "title": "JavaScript: The Good Parts",
        "genre": "Computer",
        "price": 29.99,
        "publish_date": "2008-05-01",
        "description": "Unearthing the Excellence in JavaScript"
      }
    ]
  }
```

### TOML: 缩短到彻底的利他主义

TOML（Tom 的显而易见的最小化语言Tom’s Obvious, Minimal Language）允许以相当快捷、简洁的方式定义深层嵌套的数据结构。名字中的 Tom 是指发明者 Tom Preston Werner，他是一位活跃于我们行业的创造者和软件开发人员。与 JSON 相比，语法有点尴尬，更类似 ini 文件。这不是一个糟糕的语法，但是需要一些时间适应。

```toml
[[books]]
id = 'bk101'
author = 'Crockford, Douglas'
title = 'JavaScript: The Good Parts'
genre = 'Computer'
price = 29.99
publish_date = 2008-05-01T00:00:00+00:00
description = 'Unearthing the Excellence in JavaScript'
```
TOML 中集成了一些很棒的功能，例如多行字符串、保留字符的自动转义、日期、时间、整数、浮点数、科学记数法和“表扩展”等数据类型。

### CSON: 特定系统所包含的简单样本

首先，有两个 CSON 规范。 一个代表 `CoffeeScript Object Notation`，另一个代表 `Cursive Script Object Notation`。后者不经常使用，所以我们不会关注它。我们只关注 CoffeeScript。

CSON 需要一点介绍。首先，我们来谈谈 CoffeeScript。CoffeeScript 是一种通过运行编译器生成 JavaScript 的语言。它允许你以更加简洁的语法编写 JavaScript 并转译成实际的 JavaScript，然后你可以在你的 web 应用程序中使用它。CoffeeScript 通过删除 JavaScript 中必需的许多额外语法，使编写 JavaScript 变得更容易。CoffeeScript 摆脱的一个大问题是花括号 —— 不需要它们。同样，CSON 是没有大括号的 JSON。它依赖于缩进来确定数据的层次结构。CSON 非常易于读写，并且通常比 JSON 需要更少的代码行，因为没有括号。

CSON 还提供一些 JSON 不提供的额外细节。多行字符串非常容易编写，你可以通过使用 # 符号开始一行来输入注释，并且不需要用逗号分隔键值对。

```cson
books: [
  id: 'bk102'
  author: 'Crockford, Douglas'
  title: 'JavaScript: The Good Parts'
  genre: 'Computer'
  price: 29.99
  publish_date: '2008-05-01'
  description: 'Unearthing the Excellence in JavaScript'
]
```

### YAML：年轻人的呼喊

开发人员感到高兴，因为 YAML 来自一个 Python 的贡献者。YAML 具有与 CSON 相同的功能集和类似的语法，有一系列新功能，以及几乎所有 web 编程语言都可用的解析器。它还有一些额外的功能，如循环引用、软包装、多行键、类型转换标签、二进制数据、对象合并和集合映射。它具有非常好的可读性和可写性，并且是 JSON 的超集，因此你可以在 YAML 中使用完全合格的 JSON 语法并且一切正常工作。你几乎不需要引号，它可以解释大多数基本数据类型（字符串、整数、浮点数、布尔值等）。

```yaml
books:
  - id: bk102
  author: Crockford, Douglas
  title: 'JavaScript: The Good Parts'
  genre: Computer
  price: 29.99
  publish_date: !!str 2008-05-01
  description: Unearthing the Excellence in JavaScript
```

YAML 有两个问题，对我而言，第一个是大问题。在撰写本文时，YAML 解析器尚未内置于多种语言，因此你需要使用第三方库或扩展来为你选择的语言解析 .yaml 文件。这不是什么大问题，可似乎大多数为 YAML 创建解析器的开发人员都选择随机将“附加功能”放入解析器中。有些允许标记化，有些允许链引用，有些甚至允许内联计算。这一切都很好（某种意义上），只是这些功能都不是规范的一部分，因此很难在其他语言的其他解析器中找到。这导致系统限定，你最终遇到了与 CSON 相同的问题。如果你使用仅在一个解析器中找到的功能，则其他解析器将无法解释输入。大多数这些功能都是无意义的，不属于数据集，而是属于你的应用程序逻辑，因此最好简单地忽略它们和编写符合规范的 YAML。
