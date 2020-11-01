# 学习笔记 Week09

## 高级动态规划

动态规划和递归或者分治没有根本上的区别（关键看有无最优子结构）

和递归分治的共性：找到重复子问题

和递归分治的差异：最优子结构，中途可以淘汰次优解

### DP 递推模板

```
function DB():
    dp=[][] #二维情况

    for i=0...m {
        for i=0...n{
            dp[i][j] = _Function(dp[i'][j'])
        }
    }

    return dp[m][n]
```

### 高级dp复杂度来源

1. 状态拥有更多的维度（二维，三维或更多，甚至需要压缩）
2. 状态方程更加复杂

## 字符串算法


[字符串是否可变](https://lemire.me/blog/2017/07/07/are-your-strings-immutable/)

[atoi代码](https://shimo.im/docs/5kykuLmt7a4DdjSP/read)

### 遍历字符串

c++
```
string x("abc")
for (int i = 0; i<x.lenght();i++){
    count << x[i]
}
```

golang
```
x:="abc"
for i:=0; i<len(x);i++ {
    fmt.Println(x[i])
}
```

### 字符串匹配算法

1. [暴力法](https://shimo.im/docs/8G0aJqNL86wWrPUE/read)
2. [Rabin-Karp算法](https://shimo.im/docs/1wnsM7eaZ6Ab9j9M/read)
3. [KMP算法](http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html)
   
   3.1. [视频](https://www.bilibili.com/video/av11866460?from=search&seid=17425875345653862171)
4. [Boyer-Moore算法](https://www.ruanyifeng.com/blog/2013/05/boyer-moore_string_search_algorithm.html)
5. [Sunday算法](https://blog.csdn.net/u012505432/article/details/52210975)

#### Rabin-Karp算法思想

* 假设子串长度为M（pat）目标字符串长度为N（txt)
* 计算子串的hash值hash_pat
* 计算目标字符串txt中每个长度为M的子串的hash值（共需要计算N-M+1次）
* 比较hash值，如果hash值不同则必然不匹配；如果相同则需要使用朴素算法进行进一步判断

