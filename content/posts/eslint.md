---
title: "Eslint"
date: 2020-12-24T10:26:32+08:00
draft: false
---

### 最佳实践
accessor-pairs	      最佳实践	一般	有 setter 的地方必须有 getter，有 getter 的地方可以没有 setter
array-callback-return	最佳实践	一般	数组的一些方法（map, reduce 等）的回调函数中，必须有返回值
block-scoped-var	    最佳实践	一般	将 var 定义的变量视为块作用域，禁止在块外使用
class-methods-use-this	最佳实践	一般	在类的非静态方法中，必须存在对 this 的引用
consistent-return	最佳实践	一般	禁止函数在不同分支返回不同类型的值
complexity	最佳实践	一般	禁止函数的循环复杂度超过 20
curly	最佳实践	一般	if 后面必须要有 {，除非是单行 if
default-case	最佳实践	一般	switch 语句必须有 default
dot-location	最佳实践	一般	链式调用的时候，点号必须放在第二行开头处，禁止放在第一行结尾处
dot-notation	最佳实践	一般	禁止出现 foo['bar']，必须写成 foo.bar
eqeqeq	最佳实践	一般	必须使用 === 或 !==，禁止使用 == 或 !=，与 null 比较时除外
guard-for-in	最佳实践	一般	for in 内部必须有 hasOwnProperty
no-alert	最佳实践	一般	禁止使用 alert
no-caller	最佳实践	一般	禁止使用 caller 或 callee
no-case-declarations	最佳实践	一般	switch 的 case 内有变量定义的时候，必须使用大括号将 case 内变成一个代码块
no-div-regex	最佳实践	一般	禁止在正则表达式中出现没必要的转义符
no-else-return	最佳实践	一般	禁止在 else 内使用 return，必须改为提前结束
no-empty-function	最佳实践	一般	不允许有空函数，除非是将一个空函数设置为某个项的默认值
no-empty-pattern	最佳实践	一般	禁止解构中出现空 {} 或 []
no-eq-null	最佳实践	一般	禁止使用 foo == null 或 foo != null，必须使用 foo === null 或 foo !== null
no-eval	最佳实践	一般	禁止使用 eval
no-extend-native	最佳实践	一般	禁止修改原生对象
no-extra-bind	最佳实践	一般	禁止出现没必要的 bind
no-extra-label	最佳实践	一般	禁止出现没必要的 label
no-fallthrough	最佳实践	一般	switch 的 case 内必须有 break, return 或 throw
no-floating-decimal	最佳实践	一般	表示小数时，禁止省略 0，比如 .5
no-global-assign	最佳实践	一般	禁止对全局变量赋值
no-implicit-coercion	最佳实践	一般	禁止使用 !! ~ 等难以理解的运算符
no-implicit-globals	最佳实践	一般	禁止在全局作用域下定义变量或申明函数
no-implied-eval	最佳实践	一般	禁止在 setTimeout 或 setInterval 中传入字符串，如 setTimeout('alert("Hi!")', 100);
no-invalid-this	最佳实践	一般	禁止在类之外的地方使用 this
no-iterator	最佳实践	一般	禁止使用 iterator
no-labels	最佳实践	一般	禁止使用 label
no-lone-blocks	最佳实践	一般	禁止使用没必要的 {} 作为代码块
no-loop-func	最佳实践	一般	禁止在循环内的函数中出现循环体条件语句中定义的变量，
no-magic-numbers	最佳实践	一般	禁止使用 magic numbers
no-multi-spaces	最佳实践	一般	禁止出现连续的多个空格，除非是注释前，或对齐对象的属性、变量定义、import 等
no-multi-str	最佳实践	一般	禁止使用 \ 来换行字符串
no-new	最佳实践	一般	禁止直接 new 一个类而不赋值
no-new-func	最佳实践	一般	禁止使用 new Function，比如 let x = new Function("a", "b", "return a + b");
no-new-wrappers	最佳实践	一般	禁止使用 new 来生成 String, Number 或 Boolean
no-octal	最佳实践	一般	禁止使用 0 开头的数字表示八进制数
no-octal-escape	最佳实践	一般	禁止使用八进制的转义符
no-param-reassign	最佳实践	一般	禁止对函数的参数重新赋值
no-proto	最佳实践	一般	禁止使用__proto

no-redeclare	最佳实践	一般	禁止重复定义变量

no-restricted-properties	最佳实践	一般	禁止使用指定的对象属性

no-return-assign	      最佳实践	一般	禁止在 return 语句里赋值
no-return-await	        最佳实践	一般	禁止在 return 语句里使用 await
no-script-url	          最佳实践	一般	禁止出现 location.href = 'javascript:void(0)';
no-self-assign	        最佳实践	一般	禁止将自己赋值给自己
no-self-compare	        最佳实践	一般	禁止将自己与自己比较
no-sequences	          最佳实践	一般	禁止使用逗号操作符
no-throw-literal	      最佳实践	一般	禁止 throw 字面量，必须 throw 一个 Error 对象
no-unmodified-loop-condition	最佳实践	一般	循环内必须对循环条件的变量有修改
no-unused-expressions	  最佳实践	一般	禁止无用的表达式
no-unused-labels	      最佳实践	一般	禁止出现没用的 label
no-useless-call	        最佳实践	一般	禁止出现没必要的 call 或 apply
no-useless-concat	      最佳实践	一般	禁止出现没必要的字符串连接
no-useless-escape	      最佳实践	一般	禁止出现没必要的转义
no-useless-return	      最佳实践	一般	禁止没必要的 return
no-void	                最佳实践	一般	禁止使用 void
no-warning-comments	    最佳实践	一般	禁止注释中出现 TODO 和 FIXME
no-with	                最佳实践	一般	禁止使用 with
prefer-promise-reject-errors	最佳实践	一般	Promise 的 reject 中必须传入 Error 对象，而不是字面量
radix	                  最佳实践	一般	parseInt 必须传入第二个参数


### 风格相关
array-bracket-newline	风格相关	提示	配置数组的中括号内前后的换行格式
array-bracket-spacing	风格相关	提示	数组的括号内的前后禁止有空格
array-element-newline	风格相关	提示	配置数组的元素之间的换行格式
block-spacing	风格相关	提示	代码块如果在一行内，那么大括号内的首尾必须有空格，比如 function () { alert('Hello') }
brace-style	风格相关	提示	if 与 else 的大括号风格必须一致

camelcase	风格相关	提示	变量名必须是 camelcase 风格的
capitalized-comments	风格相关	提示	注释的首字母必须大写

comma-dangle	风格相关	提示	对象的最后一个属性末尾必须有逗号
comma-spacing	风格相关	提示	逗号前禁止有空格，逗号后必须要有空格
comma-style	风格相关	提示	禁止在行首写逗号

computed-property-spacing	风格相关	提示	用作对象的计算属性时，中括号内的首尾禁止有空格

consistent-this	风格相关	提示	限制 this 的别名


eol-last	风格相关	提示	文件最后一行必须有一个空行


func-call-spacing	风格相关	提示	函数名和执行它的括号之间禁止有空格
func-name-matching	风格相关	提示	函数赋值给变量的时候，函数名必须与变量名一致
func-names	风格相关	提示	函数必须有名字
func-style	风格相关	提示	必须只使用函数申明或只使用函数表达式
id-blacklist	风格相关	提示	禁止使用指定的标识符
id-length	风格相关	提示	限制变量名长度
id-match	风格相关	提示	限制变量名必须匹配指定的正则表达式
indent	风格相关	提示	一个缩进必须用四个空格替代

jsx-quotes	风格相关	提示	jsx 中的属性必须用双引号
key-spacing	风格相关	提示	对象字面量中冒号前面禁止有空格，后面必须有空格
keyword-spacing	风格相关	提示	关键字前后必须有空格
linebreak-style	风格相关	提示	限制换行符为 LF 或 CRLF
line-comment-position	风格相关	提示	单行注释必须写在上一行
lines-around-comment	风格相关	提示	注释前后必须有空行
max-depth	风格相关	提示	代码块嵌套的深度禁止超过 5 层
max-len	风格相关	提示	现在编辑器已经很智能了，不需要限制一行的长度
max-lines	风格相关	提示	限制一个文件最多的行数
max-nested-callbacks	风格相关	提示	回调函数嵌套禁止超过 3 层，多了请用 async await 替代
max-params	风格相关	提示	函数的参数禁止超过 7 个
max-statements	风格相关	提示	限制函数块中的语句数量
max-statements-per-line	风格相关	提示	限制一行中的语句数量
multiline-ternary	风格相关	提示	三元表达式必须得换行
new-cap	风格相关	提示	new 后面的类名必须首字母大写
newline-per-chained-call	风格相关	提示	链式调用必须换行
new-parens	风格相关	提示	new 后面的类必须有小括号

no-array-constructor	风格相关	提示	禁止使用 Array 构造函数
no-await-in-loop	逻辑相关	一般	禁止将 await 写在循环里，因为这样就无法同时发送多个异步请求了
no-bitwise	风格相关	提示	禁止使用位运算

no-continue	风格相关	提示	禁止使用 continue
no-inline-comments	风格相关	提示	禁止在代码后添加内联注释
no-lonely-if	风格相关	提示	禁止 else 中只有一个单独的 if
no-mixed-operators	风格相关	提示	禁止混用不同的操作符，比如 let foo = a && b < 0 || c > 0 || d + 1 === 0


no-mixed-spaces-and-tabs	风格相关	提示	禁止混用空格和缩进
no-multi-assign	        风格相关	提示	禁止连续赋值，比如 a = b = c = 5
no-multiple-empty-lines	风格相关	提示	禁止出现超过三行的连续空行
nonblock-statement-body-position	风格相关	提示	禁止 if 后面不加大括号而写两行代码
no-negated-condition	  风格相关	提示	否定的表达式可以把逻辑表达的更清楚
no-nested-ternary	      风格相关	提示	禁止使用嵌套的三元表达式，比如 a ? b : c ? d : e
no-new-object	          风格相关	提示	禁止直接 new Object

no-plusplus	            风格相关	提示	禁止使用 ++ 或 --
no-restricted-syntax	  风格相关	提示	禁止使用特定的语法
no-tabs	                风格相关	提示	禁止使用 tabs
no-ternary	            风格相关	提示	禁止使用三元表达式

no-trailing-spaces	    风格相关	提示	禁止行尾有空格

no-underscore-dangle	  风格相关	提示	禁止变量名出现下划线

no-unneeded-ternary	    风格相关	提示	必须使用 !a 替代 a ? false : true

no-whitespace-before-property	风格相关	提示	禁止属性前有空格，比如 foo. bar()


object-curly-newline	  风格相关	提示	大括号内的首尾必须有换行
object-curly-spacing	  风格相关	提示	对象字面量只有一行时，大括号内的首尾必须有空格
object-property-newline	风格相关	提示	对象字面量内的属性每行必须只有一个

one-var	                风格相关	提示	禁止变量申明时用逗号一次申明多个
one-var-declaration-per-line	风格相关	提示	变量申明必须每行一个
operator-assignment	    风格相关	提示	必须使用 x = x + y 而不是 x += y
operator-linebreak	    风格相关	提示	需要换行的时候，操作符必须放在行末
padded-blocks	          风格相关	提示	代码块首尾必须要空行
padding-line-between-statements	风格相关	提示	限制语句之间的空行规则，比如变量定义完之后必须要空行

quote-props	            风格相关	提示	对象字面量的键名禁止用引号括起来
quotes	                风格相关	提示	必须使用单引号，禁止使用双引号

### ES6相关
arrow-body-style	      ES6相关	一般	箭头函数能够省略 return 的时候，必须省略，比如必须写成 () => 0，禁止写成 () => { return 0 }
arrow-parens	          ES6相关	一般	箭头函数只有一个参数的时候，必须加括号
arrow-spacing	          ES6相关	一般	箭头函数的箭头前后必须有空格
constructor-super	      ES6相关	一般	constructor 中必须有 super
generator-star-spacing	ES6相关	一般	generator 的 * 前面禁止有空格，后面必须有空格
no-class-assign	        ES6相关	一般	禁止对定义过的 class 重新赋值
no-confusing-arrow	    ES6相关	一般	禁止出现难以理解的箭头函数，比如 let x = a => 1 ? 2 : 3
no-const-assign	        ES6相关	一般	禁止对使用 const 定义的常量重新赋值

no-dupe-class-members	  ES6相关	一般	禁止重复定义类
no-duplicate-imports	  ES6相关	一般	禁止重复 import 模块
no-new-symbol	          ES6相关	一般	禁止使用 new 来生成 Symbol
no-restricted-imports	  ES6相关	一般	禁止 import 指定的模块
no-this-before-super	  ES6相关	一般	禁止在 super 被调用之前使用 this 或 super
no-useless-computed-key	ES6相关	一般	禁止出现没必要的计算键名，比如 let a = { ['0']: 0 };
no-useless-constructor	ES6相关	一般	禁止出现没必要的 constructor，比如 constructor(value) { super(value) }
no-useless-rename	      ES6相关	一般	禁止解构时出现同样名字的的重命名，比如 let { foo: foo } = bar;
no-var	                ES6相关	一般	禁止出现 var
object-shorthand	      ES6相关	一般	必须使用 a = {b} 而不是 a = {b: b}
prefer-arrow-callback	  ES6相关	一般	必须使用箭头函数作为回调
prefer-const	          ES6相关	一般	申明后不再被修改的变量必须使用 const 来申明
prefer-destructuring	  ES6相关	一般	必须使用解构
prefer-numeric-literals	ES6相关	一般	必须使用 0b11111011 而不是 parseInt('111110111', 2)

prefer-rest-params	    ES6相关	一般	必须使用 ...args 而不是 arguments
prefer-spread	          ES6相关	一般	必须使用 ... 而不是 apply，比如 foo(...args)
prefer-template	        ES6相关	一般	必须使用模版字面量而不是字符串连接

callback-return	        Node.js或CommonJS相关	一般	callback 之后必须立即 return
global-require	        Node.js或CommonJS相关	一般	require 必须在全局作用域下
handle-callback-err	    Node.js或CommonJS相关	一般	callback 中的 error 必须被处理
no-buffer-constructor	  Node.js或CommonJS相关	一般	禁止直接使用 Buffer
no-mixed-requires	      Node.js或CommonJS相关	一般	相同类型的 require 必须放在一起
no-new-require	        Node.js或CommonJS相关	一般	禁止直接 new require('foo')
no-path-concat	        Node.js或CommonJS相关	一般	禁止对__dirname 或__filename 使用字符串连接
no-process-env	        Node.js或CommonJS相关	一般	禁止使用 process.env.NODE_ENV
no-process-exit	        Node.js或CommonJS相关	一般	禁止使用 process.exit(0)
no-restricted-modules	  Node.js或CommonJS相关	一般	禁止使用指定的模块
no-sync	                Node.js或CommonJS相关	一般	禁止使用 node 中的同步的方法，比如 fs.readFileSync



### 逻辑相关

for-direction	          逻辑相关	严重	禁止 for 循环出现方向错误的循环，比如 for (i = 0; i < 10; i--)
getter-return	          逻辑相关	严重	getter 必须有返回值，并且禁止返回空，比如 return;
no-compare-neg-zero	    逻辑相关	严重	禁止与负零进行比较
no-cond-assign  	      逻辑相关	严重	禁止在 if, for, while 里使用赋值语句，除非这个赋值语句被括号包起来了
no-console	            逻辑相关	严重	禁止使用 console
no-constant-condition	  逻辑相关	严重	禁止将常量作为 if 或三元表达式的测试条件，比如 if (true), let foo = 0 ? 'foo' : 'bar'
no-control-regex	      逻辑相关	严重	禁止在正则表达式中出现 Ctrl 键的 ASCII 表示，即禁止使用 /\x1f/
no-debugger	            逻辑相关	严重	禁止使用 debugger
no-dupe-args        	  逻辑相关	严重	禁止在函数参数中出现重复名称的参数
no-dupe-keys	          逻辑相关	严重	禁止在对象字面量中出现重复名称的键名
no-duplicate-case	      逻辑相关	严重	禁止在 switch 语句中出现重复测试表达式的 case
no-empty	              逻辑相关	严重	禁止出现空代码块
no-empty-character-class	逻辑相关	严重	禁止在正则表达式中使用空的字符集
no-ex-assign	          逻辑相关	严重	禁止将 catch 的第一个参数 error 重新赋值
no-extra-boolean-cast	  逻辑相关	严重	禁止不必要的布尔转换
no-extra-parens	        逻辑相关	严重	禁止出现多余的括号，比如 (a * b) + c
no-extra-semi	          逻辑相关	严重	禁止出现多余的分号
no-func-assign	        逻辑相关	严重	禁止将一个函数申明重新赋值
no-inner-declarations	  逻辑相关	严重	禁止在 if 内出现函数申明或使用 var 定义变量
no-invalid-regexp	      逻辑相关	严重	禁止出现非法的正则表达式
no-irregular-whitespace	逻辑相关	严重	禁止使用特殊空白符（比如全角空格），除非是出现在字符串、正则表达式或模版字符串中
no-obj-calls	          逻辑相关	严重	禁止将 Math, JSON 或 Reflect 直接作为函数调用，必须作为类使用
no-prototype-builtins	  逻辑相关	严重	禁止使用 hasOwnProperty, isPrototypeOf 或 propertyIsEnumerable
no-regex-spaces	        逻辑相关	严重	禁止在正则表达式中出现连续的空格，必须使用 /foo {3}bar/ 代替
no-sparse-arrays	      逻辑相关	严重	禁止在数组中出现连续的逗号，如 let foo = [,,]
no-template-curly-in-string	逻辑相关	严重	禁止在普通字符串中出现模版字符串的变量形式，如 'Hello ${name}!'
no-unexpected-multiline	逻辑相关	严重	禁止出现难以理解的多行表达式
no-unreachable	        逻辑相关	严重	禁止在 return, throw, break 或 continue 之后还有代码
no-unsafe-finally	      逻辑相关	严重	禁止在 finally 中出现 return, throw, break 或 continue
no-unsafe-negation	    逻辑相关	严重	禁止在 in 或 instanceof 操作符的左侧使用感叹号，如 if (!key in object)

### 变量相关
init-declarations	      变量相关	一般	变量必须在定义的时候赋值
no-catch-shadow	        变量相关	一般	禁止 catch 的参数名与定义过的变量重复
no-delete-var         	变量相关	一般	禁止使用 delete
no-label-var	          变量相关	一般	禁止 label 名称与定义过的变量重复
no-restricted-globals	  变量相关	一般	禁止使用指定的全局变量
no-shadow	              变量相关	一般	禁止变量名与上层作用域内的定义过的变量重复
no-shadow-restricted-names	变量相关	一般	禁止使用保留字作为变量名
no-undef	              变量相关	一般	禁止使用未定义的变量
no-undefined	          变量相关	一般	禁止对 undefined 重新赋值
no-undef-init	          变量相关	一般	禁止将 undefined 赋值给变量
no-unused-vars	        变量相关	一般	定义过的变量必须使用
no-use-before-define	  变量相关	一般	变量必须先定义后使用

