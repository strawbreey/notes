---
title: "Go String Range"
date: 2020-12-03T14:41:24+08:00
draft: false
---



索引值是返回的第二个值的当前“字符”的的一个byte的索引。它不是当前“字符”的索引，这与其他语言不同。注意真实的字符可能会由多个rune表示。如果你需要处理字符，确保你使用了“norm”包。

string变量的for range语句将会尝试把数据翻译为UTF8文本。对于它无法理解的任何byte序列，它将返回0xfffd runes，而不是真实的数据。如果你任意（非UTF8文本）的数据保存在string变量中，确保把它们转换为byte slice，以得到所有保存的数据。

```go
for pos, char := range "日本語" {
    fmt.Printf("character %c starts at byte position %d\n", char, pos)
}

// character 日 starts at byte position 0
// character 本 starts at byte position 3
// character 語 starts at byte position 6
```


### 参考资料

- [The Go Blog Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
- [How can I iterate over a string by runes in Go?](https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go)

- [全面认识golang string](https://www.cnblogs.com/fanbi/p/11956208.html)