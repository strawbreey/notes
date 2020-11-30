# notes
http://note.strawbreey.com


Markdown 是一种方便记忆、书写的纯文本标记语言，用户可以使用这些标记符号以最小的输入代价生成极富表现力的文档：譬如您正在阅读的这份文档。它使用简单的符号标记不同的标题，分割不同的段落，**粗体** 或者 *斜体* 某些文字，更棒的是，它还可以

## 目录

<!-- TOC -->

- [notes](#notes)
  - [title: "Markdown"](#title-markdown)
  - [目录](#%E7%9B%AE%E5%BD%95)
    - [斜体和粗体](#%E6%96%9C%E4%BD%93%E5%92%8C%E7%B2%97%E4%BD%93)
    - [分级标题](#%E5%88%86%E7%BA%A7%E6%A0%87%E9%A2%98)
    - [外链接](#%E5%A4%96%E9%93%BE%E6%8E%A5)
    - [无序列表](#%E6%97%A0%E5%BA%8F%E5%88%97%E8%A1%A8)
    - [有序列表](#%E6%9C%89%E5%BA%8F%E5%88%97%E8%A1%A8)
    - [文字引用](#%E6%96%87%E5%AD%97%E5%BC%95%E7%94%A8)
    - [行内代码块](#%E8%A1%8C%E5%86%85%E4%BB%A3%E7%A0%81%E5%9D%97)
    - [代码块](#%E4%BB%A3%E7%A0%81%E5%9D%97)
    - [插入图像](#%E6%8F%92%E5%85%A5%E5%9B%BE%E5%83%8F)
    - [内容目录](#%E5%86%85%E5%AE%B9%E7%9B%AE%E5%BD%95)
    - [标签分类](#%E6%A0%87%E7%AD%BE%E5%88%86%E7%B1%BB)
    - [删除线](#%E5%88%A0%E9%99%A4%E7%BA%BF)
    - [注脚](#%E6%B3%A8%E8%84%9A)
    - [LaTeX 公式](#latex-%E5%85%AC%E5%BC%8F)
    - [加强的代码块](#%E5%8A%A0%E5%BC%BA%E7%9A%84%E4%BB%A3%E7%A0%81%E5%9D%97)
    - [流程图](#%E6%B5%81%E7%A8%8B%E5%9B%BE)
      - [示例](#%E7%A4%BA%E4%BE%8B)
      - [更多语法参考：流程图语法参考](#%E6%9B%B4%E5%A4%9A%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83%E6%B5%81%E7%A8%8B%E5%9B%BE%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83)
    - [序列图](#%E5%BA%8F%E5%88%97%E5%9B%BE)
      - [示例 1](#%E7%A4%BA%E4%BE%8B-1)
      - [示例 2](#%E7%A4%BA%E4%BE%8B-2)
      - [更多语法参考：序列图语法参考](#%E6%9B%B4%E5%A4%9A%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83%E5%BA%8F%E5%88%97%E5%9B%BE%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83)
    - [甘特图](#%E7%94%98%E7%89%B9%E5%9B%BE)
      - [更多语法参考：甘特图语法参考](#%E6%9B%B4%E5%A4%9A%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83%E7%94%98%E7%89%B9%E5%9B%BE%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83)
    - [Mermaid 流程图](#mermaid-%E6%B5%81%E7%A8%8B%E5%9B%BE)
      - [更多语法参考：Mermaid 流程图语法参考](#%E6%9B%B4%E5%A4%9A%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83mermaid-%E6%B5%81%E7%A8%8B%E5%9B%BE%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83)
    - [Mermaid 序列图](#mermaid-%E5%BA%8F%E5%88%97%E5%9B%BE)
      - [更多语法参考：Mermaid 序列图语法参考](#%E6%9B%B4%E5%A4%9A%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83mermaid-%E5%BA%8F%E5%88%97%E5%9B%BE%E8%AF%AD%E6%B3%95%E5%8F%82%E8%80%83)
    - [表格支持](#%E8%A1%A8%E6%A0%BC%E6%94%AF%E6%8C%81)
    - [定义型列表](#%E5%AE%9A%E4%B9%89%E5%9E%8B%E5%88%97%E8%A1%A8)
    - [Html 标签](#html-%E6%A0%87%E7%AD%BE)
    - [内嵌图标](#%E5%86%85%E5%B5%8C%E5%9B%BE%E6%A0%87)
    - [待办事宜 Todo 列表](#%E5%BE%85%E5%8A%9E%E4%BA%8B%E5%AE%9C-todo-%E5%88%97%E8%A1%A8)
    - [注脚](#%E6%B3%A8%E8%84%9A)
    - [参考链接](#%E5%8F%82%E8%80%83%E9%93%BE%E6%8E%A5)
    - [相关文章](#%E7%9B%B8%E5%85%B3%E6%96%87%E7%AB%A0)



### 1. 斜体和粗体

使用 * 和 ** 表示斜体和粗体。

示例：

这是 *斜体*，这是 **粗体**。

### 2. 分级标题

使用 === 表示一级标题，使用 --- 表示二级标题。

示例：

```
这是一个一级标题
============================

这是一个二级标题
--------------------------------------------------

### 这是一个三级标题
```

你也可以选择在行首加井号表示不同级别的标题 (H1-H6)，例如：# H1, ## H2, ### H3，#### H4。

### 3. 外链接

使用 \[描述](链接地址) 为文字增加外链接。

示例： `[本人博客](http://strawbreey.github.com)`

这是去往  [本人博客](http://strawbreey.github.com) 的链接。

### 4. 无序列表

使用 *，+，- 表示无序列表。

示例：

- 无序列表项 一
- 无序列表项 二
- 无序列表项 三

### 5. 有序列表

使用数字和点表示有序列表。

示例：

1. 有序列表项 一
2. 有序列表项 二
3. 有序列表项 三

### 6. 文字引用

使用 > 表示文字引用。

示例：

> 野火烧不尽，春风吹又生。

### 7. 行内代码块

使用 \`代码` 表示行内代码块。

示例：

让我们聊聊 `html`。

### 8.  代码块

使用 四个缩进空格 表示代码块。

示例：

    这是一个代码块，此行左侧有四个不可见的空格。

### 9.  插入图像

使用 \!\[描述](图片链接地址) 插入图像。

示例：

![我的头像](https://www.zybuluo.com/static/img/my_head.jpg)


### 10. 内容目录

在段落中填写 `[TOC]` 以显示全文内容的目录结构。

[TOC]

### 11. 标签分类

在编辑区任意行的列首位置输入以下代码给文稿标签：

标签： 数学 英语 Markdown

或者

Tags： 数学 英语 Markdown

### 12. 删除线

使用 ~~ 表示删除线。

~~这是一段错误的文本。~~

### 13. 注脚

使用 [^keyword] 表示注脚。

这是一个注脚[^footnote]的样例。

这是第二个注脚[^footnote2]的样例。

### 14. LaTeX 公式

$ 表示行内公式： 

质能守恒方程可以用一个很简洁的方程式 $E=mc^2$ 来表达。

$$ 表示整行公式：

$$\sum_{i=1}^n a_i=0$$

$$f(x_1,x_x,\ldots,x_n) = x_1^2 + x_2^2 + \cdots + x_n^2 $$

$$\sum^{j-1}_{k=0}{\widehat{\gamma}_{kj} z_k}$$

访问 [MathJax](http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference) 参考更多使用方法。


### 15. 加强的代码块

支持四十一种编程语言的语法高亮的显示，行号显示。

非代码示例：

```
$ sudo apt-get install vim-gnome
```

Python 示例：

```python
@requires_authorization
def somefunc(param1='', param2=0):
    '''A docstring'''
    if param1 > param2: # interesting
        print 'Greater'
    return (param2 - param1 + 1) or None

class SomeClass:
    pass

>>> message = '''interpreter
... prompt'''
```

JavaScript 示例：

``` javascript
/**
* nth element in the fibonacci series.
* @param n >= 0
* @return the nth element, >= 0.
*/
function fib(n) {
  var a = 1, b = 1;
  var tmp;
  while (--n >= 0) {
    tmp = a;
    a += b;
    b = tmp;
  }
  return a;
}

document.write(fib(10));
```

### 16. 流程图

#### 示例

```flow
st=>start: Start:>https://www.zybuluo.com
io=>inputoutput: verification
op=>operation: Your Operation
cond=>condition: Yes or No?
sub=>subroutine: Your Subroutine
e=>end

st->io->op->cond
cond(yes)->e
cond(no)->sub->io
```

#### 更多语法参考：[流程图语法参考](http://adrai.github.io/flowchart.js/)

### 17. 序列图

#### 示例 1

```seq
Alice->Bob: Hello Bob, how are you?
Note right of Bob: Bob thinks
Bob-->Alice: I am good thanks!
```

#### 示例 2

```seq
Title: Here is a title
A->B: Normal line
B-->C: Dashed line
C->>D: Open arrow
D-->>A: Dashed open arrow
```

#### 更多语法参考：[序列图语法参考](http://bramp.github.io/js-sequence-diagrams/)

### 18. 甘特图

甘特图内在思想简单。基本是一条线条图，横轴表示时间，纵轴表示活动（项目），线条表示在整个期间上计划和实际的活动完成情况。它直观地表明任务计划在什么时候进行，及实际进展与计划要求的对比。

```gantt
    title 项目开发流程
    section 项目确定
        需求分析       :a1, 2016-06-22, 3d
        可行性报告     :after a1, 5d
        概念验证       : 5d
    section 项目实施
        概要设计      :2016-07-05  , 5d
        详细设计      :2016-07-08, 10d
        编码          :2016-07-15, 10d
        测试          :2016-07-22, 5d
    section 发布验收
        发布: 2d
        验收: 3d
```

#### 更多语法参考：[甘特图语法参考](https://knsv.github.io/mermaid/#gant-diagrams)

### 19. Mermaid 流程图

```graphLR
    A[Hard edge] -->|Link text| B(Round edge)
    B --> C{Decision}
    C -->|One| D[Result one]
    C -->|Two| E[Result two]
```

#### 更多语法参考：[Mermaid 流程图语法参考](https://knsv.github.io/mermaid/#flowcharts-basic-syntax)

### 20. Mermaid 序列图

```sequence
    Alice->John: Hello John, how are you?
    loop every minute
        John-->Alice: Great!
    end
```

#### 更多语法参考：[Mermaid 序列图语法参考](https://knsv.github.io/mermaid/#sequence-diagrams)

### 21. 表格支持

| 项目        | 价格   |  数量  |
| --------   | -----:  | :----:  |
| 计算机     | \$1600 |   5     |
| 手机        |   \$12   |   12   |
| 管线        |    \$1    |  234  |


### 22. 定义型列表

名词 1
:   定义 1（左侧有一个可见的冒号和四个不可见的空格）

代码块 2
:   这是代码块的定义（左侧有一个可见的冒号和四个不可见的空格）

        代码块（左侧有八个不可见的空格）

### 23. Html 标签

本站支持在 Markdown 语法中嵌套 Html 标签，譬如，你可以用 Html 写一个纵跨两行的表格：

    <table>
        <tr>
            <th rowspan="2">值班人员</th>
            <th>星期一</th>
            <th>星期二</th>
            <th>星期三</th>
        </tr>
        <tr>
            <td>李强</td>
            <td>张明</td>
            <td>王平</td>
        </tr>
    </table>


<table>
    <tr>
        <th rowspan="2">值班人员</th>
        <th>星期一</th>
        <th>星期二</th>
        <th>星期三</th>
    </tr>
    <tr>
        <td>李强</td>
        <td>张明</td>
        <td>王平</td>
    </tr>
</table>

### 24. 内嵌图标

本站的图标系统对外开放，在文档中输入

    <i class="icon-weibo"></i>

即显示微博的图标： <i class="icon-weibo icon-2x"></i>

替换 上述 `i 标签` 内的 `icon-weibo` 以显示不同的图标，例如：

    <i class="icon-renren"></i>

即显示人人的图标： <i class="icon-renren icon-2x"></i>

更多的图标和玩法可以参看 [font-awesome](http://fortawesome.github.io/Font-Awesome/3.2.1/icons/) 官方网站。

### 25. 待办事宜 Todo 列表

使用带有 [ ] 或 [x] （未完成或已完成）项的列表语法撰写一个待办事宜列表，并且支持子列表嵌套以及混用Markdown语法，例如：

    - [ ] **Cmd Markdown 开发**
        - [ ] 改进 Cmd 渲染算法，使用局部渲染技术提高渲染效率
        - [ ] 支持以 PDF 格式导出文稿
        - [x] 新增Todo列表功能 [语法参考](https://github.com/blog/1375-task-lists-in-gfm-issues-pulls-comments)
        - [x] 改进 LaTex 功能
            - [x] 修复 LaTex 公式渲染问题
            - [x] 新增 LaTex 公式编号功能 [语法参考](http://docs.mathjax.org/en/latest/tex.html#tex-eq-numbers)
    - [ ] **七月旅行准备**
        - [ ] 准备邮轮上需要携带的物品
        - [ ] 浏览日本免税店的物品
        - [x] 购买蓝宝石公主号七月一日的船票
        
对应显示如下待办事宜 Todo 列表：
        
- [ ] **Cmd Markdown 开发**
    - [ ] 改进 Cmd 渲染算法，使用局部渲染技术提高渲染效率
    - [ ] 支持以 PDF 格式导出文稿
    - [x] 新增Todo列表功能 [语法参考](https://github.com/blog/1375-task-lists-in-gfm-issues-pulls-comments)
    - [x] 改进 LaTex 功能
        - [x] 修复 LaTex 公式渲染问题
        - [x] 新增 LaTex 公式编号功能 [语法参考](http://docs.mathjax.org/en/latest/tex.html#tex-eq-numbers)
- [ ] **七月旅行准备**
    - [ ] 准备邮轮上需要携带的物品
    - [ ] 浏览日本免税店的物品
    - [x] 购买蓝宝石公主号七月一日的船票
        
### 注脚

[^footnote]: 这是一个 *注脚* 的 **文本**。

[^footnote2]: 这是另一个 *注脚* 的 **文本**。


### 参考链接


### 相关文章




