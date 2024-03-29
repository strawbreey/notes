---
title: "Html Entities"
date: 2021-09-01T11:16:37+08:00
draft: false
---

## HTML 中的预留字符必须被替换为字符实体。

在 HTML 中，某些字符是预留的。

在 HTML 中不能使用小于号 `<` 和大于号`>`，这是因为浏览器会误认为它们是标签。

如果希望正确地显示预留字符，我们必须在 HTML 源代码中使用字符实体（character entities）。

字符实体类似这样：

```
&entity_name;

或者

&#entity_number;
```

如需显示小于号，我们必须这样写：`&lt;` 或 `&#60;`

提示：使用实体名而不是数字的好处是，名称易于记忆。不过坏处是，浏览器也许并不支持所有实体名称（对实体数字的支持却很好）。

### 不间断空格（non-breaking space）

HTML 中的常用字符实体是不间断空格 `&nbsp;`。

浏览器总是会截短 HTML 页面中的空格。如果您在文本中写 10 个空格，在显示该页面之前，浏览器会删除它们中的 9 个。如需在页面中增加空格的数量，您需要使用 `&nbsp;` 字符实体。


### HTML 中有用的字符实体

注释：实体名称对大小写敏感！

```
显示结果	描述	    实体名称	 实体编号
 	       空格	        &nbsp;	    &#160;
<	       小于号	    &lt;	    &#60;
>   	   大于号	    &gt;	    &#62;
&	       和号	        &amp;	    &#38;
"	       引号	        &quot;	    &#34;
'	       撇号 	    &apos;	    &#39;
￠	      分(cent)	   &cent;	   &#162;
£	       镑(pound)	&pound;	    &#163;
¥	       元(yen)	    &yen;	    &#165;
€	       欧元(euro)	&euro;	    &#8364;
§	       小节	        &sect;	    &#167;
©	       版权(copyright)&copy;	&#169;
®	       注册商标	     &reg;	    &#174;
™	       商标	        &trade;	    &#8482;
×	       乘号	        &times; 	&#215;
÷	       除号	        &divide;	&#247;
```
如需完整的实体符号参考，请访问我们的 HTML 实体符号参考手册。


- [HTML 实体符号参考手册](https://www.w3school.com.cn/charsets/ref_html_8859.asp)