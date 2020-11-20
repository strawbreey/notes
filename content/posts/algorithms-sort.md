---
title: "Algorithms Sort"
date: 2020-11-18T19:23:53+08:00
draft: false
---

There are three practical reasons for you to study sorting algorithms, even though 
you might just use a system sort:

-	Analyzing sorting algorithms is a thorough introduction to the approach that we 
use to compare algorithm performance throughout the book.
-	Similar techniques are effective in addressing other problems.
-	We often use sorting algorithms as a starting point to solve other problems. 
More important than these practical reasons is that the algorithms are elegant, classic, 
and effective.

即使你只是使用标准库中的排序函数，学习排序算法仍然有三大实际意义： 

 对排序算法的分析将有助于你全面理解本书中比较算法性能的方法； 

 类似的技术也能有效解决其他类型的问题； 

 排序算法常常是我们解决其他问题的第一步。


```java
public class Example 
{ 
  public static void sort(Comparable[] a) 
  { 
    /* See Algorithms 2.1, 2.2, 2.3, 2.4, 2.5, or 2.7. */ 
  }

  private static boolean less(Comparable v, Comparable w) {
    return v.compareTo(w) < 0; 
  }
  
  private static void exch(Comparable[] a, int i, int j) { 
    Comparable t = a[i]; a[i] = a[j]; a[j] = t; 
  }
  
  private static void show(Comparable[] a) { 
    // Print the array, on a single line.
    for (int i = 0; i < a.length; i++) 
    StdOut.print(a[i] + " "); 
    StdOut.println(); 
  }
  
  public static boolean isSorted(Comparable[] a) { 
    // Test whether the array entries are in order.
    for (int i = 1; i < a.length; i++) 
    if (less(a[i], a[i-1])) return false; 
    return true; 
  }
  
  public static void main(String[] args) { 
    // Read strings from standard input, sort them, and print.
    String[] a = In.readStrings(); 
    sort(a); 
    assert isSorted(a); 
    show(a); 
  } 
}
```


### words

```
 variation 变异

 Among the reasons for 其中的原因

 relatively  相对

 as we will see 正如我们将看到的

 illustrates  说明

 conventions 约定

 present 当下
```