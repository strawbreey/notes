---
title: "Leetcode 1662 Check if Two String Arrays Are Equivalent"
date: 2020-11-26T10:45:53+08:00
draft: false
---

给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。

数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。

```
输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
输出：true
解释：
word1 表示的字符串为 "ab" + "c" -> "abc"
word2 表示的字符串为 "a" + "bc" -> "abc"
两个字符串相同，返回 true

输入：word1 = ["a", "cb"], word2 = ["ab", "c"]
输出：false


输入：word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
输出：true

```

提示：

  - 1 <= word1.length, word2.length <= 103
  - 1 <= word1[i].length, word2[i].length <= 103
  - 1 <= sum(word1[i].length), sum(word2[i].length) <= 103
  - word1[i] 和 word2[i] 由小写字母组成


Solution

```js
/**
 * @param {string[]} word1
 * @param {string[]} word2
 * @return {boolean}
 */
var arrayStringsAreEqual = function(word1, word2) {
    return word1.join("") == word2.join("")
};
```