---
title: "Leetcode 0020 Valid Parentheses"
date: 2020-11-24T23:28:48+08:00
draft: false
tags: ['leetcode']
---

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

```
输入: "()"
输出: true

输入: "()[]{}"
输出: true

输入: "(]"
输出: false

输入: "([)]"
输出: false
```

solution:

```js
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let arr = []
    let count = 0
    let obj = {
        '(': ')',
        '{': '}',
        '[': ']'
    }
    while (s) {
        if (s[0] == '(' || s[0] == '{' || s[0] == '[') {
           arr.push(s[0])
           count++
        } else {
            if (obj[arr[count - 1]] == s[0]) {
                arr.pop()
                count--
            } else {
                count++
            }
        }
        s= s.substring(1)
    }
    return count == 0 ? true : false
};
```

### 参考资料

- [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)