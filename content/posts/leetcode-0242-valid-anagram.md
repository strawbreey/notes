---
title: "Leetcode 0242 Valid Anagram"
date: 2020-11-22T22:38:27+08:00
draft: false
tags: ['leetcode']
---

Given two strings s and t , write a function to determine if t is an anagram of s.

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

Example:

```
Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
```

Note:

You may assume the string contains only lowercase alphabets.
你可以假设字符串只包含小写字母。

Follow up:

What if the inputs contain unicode characters? How would you adapt your solution to such case?
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

Solution

最简单的方式

```js
var isAnagram = function(s, t) {
    return s.split("").sort().join("") == t.split("").sort().join("")
};

// 更精简的写法
var isAnagram = function(s, t) {
    return s.length == t.length && [...s].sort().join('') === [...t].sort().join('')
};
```

复杂度分析

    时间复杂度：O(n \log n)O(nlogn)，其中 nn 为 ss 的长度。排序的时间复杂度为 O(n\log n)O(nlogn)，比较两个字符串是否相等时间复杂度为 O(n)O(n)，因此总体时间复杂度为 O(n \log n+n)=O(n\log n)O(nlogn+n)=O(nlogn)。

    空间复杂度：O(\log n)O(logn)。排序需要 O(\log n)O(logn) 的空间复杂度。注意，在某些语言（比如 Java & JavaScript）中字符串是不可变的，因此我们需要额外的 O(n)O(n) 的空间来拷贝字符串。但是我们忽略这一复杂度分析，因为：


```js
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    if (s.length != t.length) {
        return false
    }

    let sObj = {}, tObj = {}

    s.split("").forEach((key) => {
        if (sObj[key]) {
            sObj[key]++
        } else {
            sObj[key] = 1
        }
    })

    t.split("").forEach((key) => {
        if (tObj[key]) {
            tObj[key]++
        } else {
            tObj[key] = 1
        }
    })

    if (Object.getOwnPropertyNames(sObj).length !=  Object.getOwnPropertyNames(tObj).length) {
        return false
    }

    for(let key in sObj) {
        if (sObj[key] != tObj[key]) {
            return false
        }
    }


    return true

};

// =
var isAnagram = function(s, t) {
    if (s.length !== t.length) {
        return false;
    }
    const table = new Array(26).fill(0);
    for (let i = 0; i < s.length; ++i) {
        // unicode
        table[s.codePointAt(i) - 'a'.codePointAt(0)]++;
    }
    for (let i = 0; i < t.length; ++i) {
        // unicode
        table[t.codePointAt(i) - 'a'.codePointAt(0)]--;
        if (table[t.codePointAt(i) - 'a'.codePointAt(0)] < 0) {
            return false;
        }
    }
    return true;
};
```

### 参考资料
- [Valid Anagram](https://leetcode.com/problems/valid-anagram/)
- [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)
