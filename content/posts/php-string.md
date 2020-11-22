---
title: "Php String"
date: 2020-09-24T15:09:55+08:00
draft: false
---

一个字符串 string 就是由一系列的字符组成，其中每个字符等同于一个字节。这意味着 PHP 只能支持 256 的字符集，因此不支持 Unicode 。详见字符串类型详解。


### 语法

一个字符串可以用 4 种方式表达：

- 单引号
- 双引号
- heredoc 语法结构
- nowdoc 语法结构（自 PHP 5.3.0 起


### 单引号

定义一个字符串的最简单的方法是用单引号把它包围起来（字符 '）。

要表达一个单引号自身，需在它的前面加个反斜线（\）来转义。要表达一个反斜线自身，则用两个反斜线（\\）。其它任何方式的反斜线都会被当成反斜线本身：也就是说如果想使用其它转义序列例如 \r 或者 \n，并不代表任何特殊含义，就单纯是这两个字符本身。

```php
echo 'this is a simple string';

echo 'You can also have embedded newlines in 
strings this way as it is
okay to do'; // 可以录入多行


echo 'Arnold once said: "I\'ll be back"'; // 输出： Arnold once said: "I'll be back"  单引号转义

echo 'You deleted C:\\*.*?'; // 输出：  You deleted C:\*.*? 单斜杠转义

echo 'You deleted C:\*.*?'; // 输出： You deleted C:\*.*?

echo 'This will not expand: \n a newline'; // 输出： This will not expand: \n a newline

echo 'Variables do not $expand $either'; // 输出： Variables do not $expand $either
```

### 双引号

如果字符串是包围在双引号（"）中， PHP 将对一些特殊的字符进行解析：

\n	换行（ASCII 字符集中的 LF 或 0x0A (10)）
\r	回车（ASCII 字符集中的 CR 或 0x0D (13)）
\t	水平制表符（ASCII 字符集中的 HT 或 0x09 (9)）
\v	垂直制表符（ASCII 字符集中的 VT 或 0x0B (11)）（自 PHP 5.2.5 起）
\e	Escape（ASCII 字符集中的 ESC 或 0x1B (27)）（自 PHP 5.4.0 起）
\f	换页（ASCII 字符集中的 FF 或 0x0C (12)）（自 PHP 5.2.5 起）
\\	反斜线
\$	美元标记
\"	双引号
\[0-7]{1,3}	符合该正则表达式序列的是一个以八进制方式来表达的字符
\x[0-9A-Fa-f]{1,2}	符合该正则表达式序列的是一个以十六进制方式来表达的字符

和单引号字符串一样，转义任何其它字符都会导致反斜线被显示出来。PHP 5.1.1 以前，\{$var} 中的反斜线还不会被显示出来。


### Heredoc 结构 

第三种表达字符串的方法是用 heredoc 句法结构：<<<。在该运算符之后要提供一个标识符，然后换行。接下来是字符串 string 本身，最后要用前面定义的标识符作为结束标志。


### Nowdoc 结构

就象 heredoc 结构类似于双引号字符串，Nowdoc 结构是类似于单引号字符串的。Nowdoc 结构很象 heredoc 结构，但是 nowdoc 中不进行解析操作。这种结构很适合用于嵌入 PHP 代码或其它大段文本而无需对其中的特殊字符进行转义。与 SGML 的 <![CDATA[ ]]> 结构是用来声明大段的不用解析的文本类似，nowdoc 结构也有相同的特征。