---
title: "Leetcode 1030 Matrix Cells in Distance Order"
date: 2020-11-17T09:56:46+08:00
draft: false
tags: ['leetcode']
---

We are given a matrix with R rows and C columns has cells with integer coordinates (r, c), where 0 <= r < R and 0 <= c < C.

Additionally, we are given a cell in that matrix with coordinates (r0, c0).

Return the coordinates of all cells in the matrix, sorted by their distance from (r0, c0) from smallest distance to largest distance.  Here, the distance between two cells (r1, c1) and (r2, c2) is the Manhattan distance, |r1 - r2| + |c1 - c2|.  (You may return the answer in any order that satisfies this condition.

给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）

Example:

```
Input: R = 1, C = 2, r0 = 0, c0 = 0
Output: [[0,0],[0,1]]
Explanation: The distances from (r0, c0) to other cells are: [0,1]

Input: R = 2, C = 2, r0 = 0, c0 = 1
Output: [[0,1],[0,0],[1,1],[1,0]]
Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2]
The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.
Example 3:

Input: R = 2, C = 3, r0 = 1, c0 = 2
Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2,2,3]
There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].
```

Note:

```
1 <= R <= 100
1 <= C <= 100
0 <= r0 < R
0 <= c0 < C
```

Solution:

方法一：直接排序

最容易想到的方法是首先存储矩阵内所有的点，然后将其按照哈曼顿距离直接排序。

复杂度分析

- 时间复杂度：O(RC \log(RC))O(RClog(RC))，存储所有点时间复杂度 O(RC)O(RC)，排序时间复杂度 O(RC \log(RC))O(RClog(RC))。

- 空间复杂度：O(\log(RC))O(log(RC))，即为排序需要使用的栈空间，不考虑返回值的空间占用。

```js
var allCellsDistOrder = function(R, C, r0, c0) {
    let arr = []
    for (let r = 0; r < R; r++) {
      for (c = 0; c < C; c++) {
        arr.push([r,c])
      }
    }
    return arr.sort((a, b) => Math.abs(a[0] - r0) +  Math.abs(a[1] - c0)  - Math.abs(b[0] - r0) -  Math.abs(b[1] - c0))
};
```

方法二：桶排序

注意到方法一中排序的时间复杂度太高。实际在枚举所有点时，我们可以直接按照哈曼顿距离分桶。这样我们就可以实现线性的桶排序。

时间复杂度：O(RC)O(RC)，存储所有点时间复杂度 O(RC)O(RC)，桶排序时间复杂度 O(RC)O(RC)。

空间复杂度：O(RC)O(RC)，需要存储矩阵内所有点。


```go
```

方法三：几何法

我们也可以直接变换枚举矩阵的顺序，直接按照曼哈顿距离遍历该矩形即可。

注意到曼哈顿距离相同的位置恰好构成一个斜着的正方形边框，因此我们可以从小到大枚举曼哈顿距离，并使用循环来直接枚举该距离对应的边框。我们每次从该正方形边框的上顶点出发，依次经过右顶点、下顶点和左顶点，最后回到上顶点。这样即可完成当前层的遍历。

```php
```

![matrix-cells-in-distance-order](/images/matrix-cells-in-distance-order.png)


### 参考链接 

- [距离顺序排列矩阵单元格](https://leetcode-cn.com/problems/matrix-cells-in-distance-order/)
- [Matrix Cells in Distance Order](https://leetcode.com/problems/matrix-cells-in-distance-order//)

