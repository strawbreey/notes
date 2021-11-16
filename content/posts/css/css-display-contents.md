---
title: "Css Display Contents"
date: 2021-11-15T11:32:00+08:00
draft: false
---

## CSS 的 display:contents

display 这一规则属性有了`display:contents` 这个新值，就能把一个元素从边界框树型结构(box tree)中移去，但内容保留。详细要求里有准确定义：

元素本身不产生任何边界框，而元素的子元素与伪元素仍然生成边界框，元素文字照常显示。为了同时照顾边界框与布局，处理这个元素时，要想象这个元素不在元素树型结构里，而只有内容留下。这包括元素在元文档中的子元素与伪元素，比如`::before`和`::after`这两个伪元素，如平常一样，前者仍然在元素子元素之前生成，后者在之后生成。

给一个简单的例子帮助正确理解：

```html
<div style="display: contents;
background: magenta; border: solid thick black; ng: 20px;
color: cyan; font: 30px/1 Monospace;">
    <span style="background: black;">foobar</span>
</div>
```
display: contents样式规则使div元素不产生任何边界框，因此元素的背景、边框和填充部分都不会渲染。然而，继承的属性如颜色(color)和字体(font)却能照常影响到span这个子元素。

就这个例子而言，最终结果应该看上去是这个样子的：

<span style="background: black; color: cyan; font: 30px/1 Monospace;">foobar</span>


上一个例子在浏览器不支持时的效果，对比实际效果与支持时的效果。

想了解更多细节，Rachel Andrew写了一篇关于这个话题的博客贴子，很不错。

### CSS网格布局与display: contents

display: contents样式能取代子网格(subgrids)功能，目前还没有任何浏览器支持次网格。不过，子网格(subgrids)有些情况还是需要用到的。

典型的例子是网格布局(Grid Layout)的自动就位(auto-placement)效果，下面是一个简单的表格元素，看起来是这样的：

```html
<style>
    form { 
        display:grid;   
    }
    label { 
        grid-column: 1;      
    }
    input { 
        grid-column: 2;      
    }
    button { 
        grid-column: span 2; 
    }
</style>
<form>
    <label>Name</label><input />
    <label>Mail</label><input />
    <button>Send</button>
</form>
```

然而这不是一个典型的HTML网页表格，因为通常我们会在表格内部使用列表，这样使用读屏软件的用户就能预先知道有多少空要填。所以HTML网页看起来更可能会是这样的：

<form>
    <ul>
        <li><label>Name</label><input /></li>
        <li><label>Mail</label><input /></li>
        <li><button>Send</button></li>
    </ul>
</form>

有了display: contents样式，就可以做出和第一个例子相同的布局，用的CSS也差不多：

```css
ul{ 
    display: grid;
}
li{ 
    display: contents;
}
label{ 
    grid-column: 1;
}
input{ 
    grid-column: 2;
}
button { 
    grid-column: span 2; 
}
```

现在这样，网站转用CSS网格布局(Grid Layout)时，HTML代码不用大改，也不需要舍去一些确实有用的HTML元素，如上面例子里的列表元素，真的很不错。

### 在Chromium浏览器上实这个功能

之前介绍部分说过，火狐三年前就已经支持display: contents功能了，但Chromium浏览器对此却没有任何实现。CSS网格布局(Grid Layout)是由Igalia公司实现的，所以公司非常希望这个功能得到支持，因为在好几个网格布局(Grid Layout)用例中，这个功能都可以方便解决问题。

Igalia公司代码之旅(Coding Experience)提出的计划是在Blink引擎上实现display: contents样式规则，并将其作为主要任务。Emilio做得很棒，大部分工作都成功完成了，发现的问题应需报告给了CSS工作组和其它浏览器团队，并为网站平台测试(web-platform-tests)仓库编写了测试，以保证不同实现方法之间的互用性。

还有一些工作要在代码之旅(Coding Experience)结束后做好，然后才能默认开启display: contents规则。前Opera公司、现Google公司的@Rune Lillesveen在整个开发过程中都给予了帮助，并将剩下的工作做完，一星期前发布。

### 在WebKit引擎上实现这个功能

WebKit引擎已经对display: contents规则有了初步支持 ，但只是在实现Shadow DOM技术时内部应用，终端用户无法接触，代码其它部分也不支持。

我们把那部分也重新激活了，虽然没时间做完，但之后苹果公司的Antti Koivisto完成了。到2017年11月，这个功能已在主干开发版本上设为默认启动。

### 总结

Igalia公司作为外部力量，致力于开放式网站平台项目，是这一领域顶尖的公司之一，所以我们有机会能在不同的开源项目中实现新的功能。这要感谢我们成员的集体参与，以及在此领域内通过几年经验积累起来的外部知识。关于display: contents的实现，如果没有Igalia公司的支持，特别是代码之旅(Coding Experience)的经历，这个功能是不可能在今天的Chromium浏览器和WebKit引擎上实现的。

代码之旅(Coding Experience)取得了好的结果，我们非常高兴，我们也期望在将来能再次大获成功。

当然，所有的功劳都应该归Emilio，他是一个了不起的工程师，在代码之旅(Coding Experience)期间做得非常出色。在这一过程中，他拥有了在Chromium项目和WebKit项目里提交的特权。赞！

最后，感谢Antti和Rune把剩下的任务完成了，才让WebKit和Chromium用户能用到display: contents样式规则。
