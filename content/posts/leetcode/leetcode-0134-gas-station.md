---
title: "Leetcode 0134 Gas Station"
date: 2020-11-18T10:05:57+08:00
draft: false
tags: ['leetcode']
---

There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。


Note:

  - If there exists a solution, it is guaranteed to be unique.
  - Both input arrays are non-empty and have the same length.
  - Each element in the input arrays is a non-negative integer.

说明: 

  - 如果题目有解，该答案即为唯一答案。
  - 输入数组均为非空数组，且长度相同。
  - 输入数组中的元素均为非负数。

Example:

```
  Input: 
  gas  = [1,2,3,4,5]
  cost = [3,4,5,1,2]

  Output: 3

  Explanation:
  Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
  Travel to station 4. Your tank = 4 - 1 + 5 = 8
  Travel to station 0. Your tank = 8 - 2 + 1 = 7
  Travel to station 1. Your tank = 7 - 3 + 2 = 6
  Travel to station 2. Your tank = 6 - 4 + 3 = 5
  Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
  Therefore, return 3 as the starting index.

  Input: 
  gas  = [2,3,4]
  cost = [3,4,3]

  Output: -1

  Explanation:
  You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
  Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
  Travel to station 0. Your tank = 4 - 3 + 2 = 3
  Travel to station 1. Your tank = 3 - 3 + 3 = 3
  You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
  Therefore, you can't travel around the circuit once no matter where you start.
```



Solution: 

```js
// javascript
var canCompleteCircuit = function(gas, cost) {
    let sum = 0
    let index = 0

    if (cost.reduce((total,num) => total + num) > gas.reduce((total,num) => total + num)) {
        return -1
    }

    for (let i = 0; i < gas.length; i++) {
        sum += gas[i] - cost[i];
        if (sum < 0) {
            sum = 0
            index = i + 1 > gas.length ? -1 : i + 1
        }
    }

    return index
};
```

最容易想到的解法是：从头到尾遍历每个加油站，并检查以该加油站为起点，最终能否行驶一周。我们可以通过减小被检查的加油站数目，来降低总的时间复杂度。

```golang
// golang
func canCompleteCircuit(gas []int, cost []int) int {
    for i, n := 0, len(gas); i < n; {
        sumOfGas, sumOfCost, cnt := 0, 0, 0
        for cnt < n {
            j := (i + cnt) % n
            sumOfGas += gas[j]
            sumOfCost += cost[j]
            if sumOfCost > sumOfGas {
                break
            }
            cnt++
        }
        if cnt == n {
            return i
        } else {
            i += cnt + 1
        }
    }
    return -1
}
```