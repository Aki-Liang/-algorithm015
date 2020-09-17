# 学习笔记 Week 04

## 搜索

遍历搜索

* 每个节点都要访问一次
* 每个节点仅访问一次
  
### 深度优先搜索 (DFS， Depth First Search)

代码模板:
```
递归写法
visited = set()

def dfs(node, visited)
    if node in visited:
        return 

    visited.add(node)

    # process current node
    ...

    for next_node in node.children():
        if not next_node in visited:
            dfs(next_node, visited)
```
```
循环写法 手动维护一个栈
def dfs(self, tree):
    if tree.root is None:
        return []
    
    visited, stack = [], [tree.root]

    while stack:
        node=stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_relatex_nodes(node)
        stack.push(nodes)

```

### 广度优先搜索 (BFS, Breadth First Search)

一层一层的往下遍历

代码模板:
```
使用队列

def bfs(graph, start, end):
    queue = []
    queue.appedn({start})
    visited.add(start)

    while queue:
        node = queue.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)

```

### 优先级优先搜索

    启发搜索

## 贪心算法

在每一步选择中都采取在当前状态下最好活最优的选择，从而希望到时结果是全局最好或最优的算法

### 适用场景

一旦一个问题可以使用贪心算法，那么贪心算法一般是解决这个问题的最好办法

* 解决一些最优化的问题
  * 最小生成树
  * 哈夫曼编码
* 用作辅助算法
* 解决一些结果要求不特别精确的问题
* 能分解成子问题的问题，子问题的最优解能递推到最终问题的最优解(子问题最优解被称为最优子结构)

### 与动态规划的区别

* 贪心
  * 对每个子问题的解决方案都做出选择
  * 不能回退
* 动态规划
  * 保存以前的运算结果，并根据以前的结果对当前进行选择
  * 有回退功能

## 二分查找

代码模板：

```
left, right = 0, len(arry)-1
while left <= right
    mid = (left+right)/2
    if array[mid] == target:
        break or return result
    elif array[mid] < target:
        left=mid+1
    else:
        right=mid-1
```

### 前提条件

* 目标函数单调性(单调递增或递减)
* 存在上下界
* 能够通过索引访问

