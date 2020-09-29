# 学习笔记 Week06

## 动态规划 (Dynamic Programming 动态递推)

使用一种递归的方式将一个复杂的问题，分解成简单的子问题

采用分治+最优子结构（Optimal substructure）的思想

### 和递归分治的区别和共同点

动态规划和递归或者分治没有本质上的区别

* 共性
  * 寻找重复子问题
* 差别
  * 动态规划有最优子结构
  * 动态规划中途可以淘汰次优解（如果不淘汰，复杂度一般是指数级；如果进行淘汰会优化成O(n^2)或者O(n)的复杂度）
  
### 关键点

* 最优子结构-1
  * opt[n]=best_of(opt[n-1],opt[n-2],...)
* 存储中间状态
  * opt[i]
* 递推公式(状态转移方程或者DP方程)
  * Fib: opt[i] = opt[n-1] + opt[n-2]

### 麻省理工五步法

* 进行分治把复杂问题转化成简单重复子问题（define subproblems）
* 猜递推方程（guess part of solution）
* 合并子问题的解（relate subproblem solutions）
* 递归&记忆化或自底向上建立DP状态表（recurse&memories or build DP table bottom-uo）
* 解决原问题（solve original problem）

## 思维小结

* 打破自己的思维惯性，形成机器思维（找重复性）
* 理解复杂逻辑的关键
* 也是职业进阶的要点要领（不要人肉递归，不要亲力亲为）