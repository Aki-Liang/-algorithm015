# 学习笔记 Week2

## 哈希表

* 也叫散列表
* 根据关键码值(key value)直接进行访问的数据结构
* 通过散列函数把关键码值映射到表中的位置来访问记录，以加快访问速度
* 会出现哈希碰撞的问题（拉链式解决冲突法）
* 时间复杂度O(1)，出现大量重复数据后最坏情况O(n)

## 树，二叉树，二叉搜索树

* 有根节点，有叶子节点
* 每个最多只有两个叶子节点的就是二叉树

### 树的遍历

* 前序遍历：根左右
* 中序遍历：左根右
* 后序遍历：左右根

### 二叉搜索树

* 空树
* 左子树上所有节点的值都小于它的根节点的值
* 右子树上所有节点的值都大于它的根节点的值
* 查询、插入 O(log n)

## 堆

* 堆是可以迅速找到一堆数中最大或者最小值的数据结构
* 根节点最大的堆叫大顶堆或大根堆
* 根节点最小的堆叫小顶堆或小根堆
* 常见的堆：二叉堆、斐波那契堆

###  二叉堆

* 通过完全二叉树实现
* 树中任意节点的值>=其子节点的值（大顶堆）

#### 实现

* 一般通过数组实现
  
```
假设第一个元素在数组索引为0
索引为i的节点的左孩子索引是2*i+1
索引为i的节点的右孩子索引是2*i+1
索引为i的节点的父节点索引是floor((i-1)/2)
```

* 插入操作

```
HeapifyUp

新元素一律先插入到堆的尾部
依次向上调整整个堆的结构（一直到根）
```

* 删除堆顶操作

```
HeapifyDown

将堆尾元素替换到顶部（堆顶被删除）
依次从根部乡下调整整个堆的结构
```

## 图

* 有点
  * 度：入度和出度
  * 点与点之间：连通与否
* 有边
  * 有向、无向
  * 权重（边长）
  

### 属性和分类

表示方法

  * 邻接矩阵
  * 邻接表

### 相关算法

* DFS 深度优先搜索
* BFS 广度优先搜搜

高级算法参考链接：

* 连通图个数： https://leetcode-cn.com/problems/number-of-islands/
* 拓扑排序（Topological Sorting）： https://zhuanlan.zhihu.com/p/34871092
* 最短路径（Shortest Path）：Dijkstra https://www.bilibili.com/video/av25829980?from=search&seid=13391343514095937158
* 最小生成树（Minimum Spanning Tree）： https://www.bilibili.com/video/av84820276?from=search&seid=17476598104352152051

## go语言HeapSort源码学习笔记

go语言的`container/heap`包提供了堆容器，其中包含堆排序的相关算法，采用二叉堆方案实现的小根堆

### 数据存储结构

go中的堆使用数组或其他任意支持下标操作的数据结构作为存储数据的载体

设根节点下标为`i`,则该节点的两个叶子节点的下标分别为`2i+1`、`2i+2`

### Push方法

源码如下

```
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

Step1. 调用存储结构的push方法，将新元素追加到尾部

Step2. 调用up方法，将尾部的新元素上浮到正确节点
```

up源码如下：

```
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}
```

找到当前要上浮的节点的父节点，与其比较大小，当前节点小于父节点则与父节点交换，循环反复，直到到达堆顶（下标为0）或该节点大于父节点为止。


### Pop方法

源码如下

```
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

Step1.将堆顶的元素和位于存储数据结构末尾的元素交换

Step2.将除了位于存储数据结构末尾的元素之外的所有元素进行排序

Step3.弹出尾部元素
```

down源码如下

```
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

```
先找到当前节点的子节点，与其比较大小，如果大于子节点则交换，循环反复，直到所有子节点大于当前节点。


### Interface的Pop和Push方法实现示例

```
func (a *Heap) Push(x interface{}) {
	*a = append(*a, x.(*HeapNode))
}

func (a *Heap) Pop() interface{} {
	x := (*a)[len(*a)-1]
	(*a) = (*a)[:len(*a)-1]
	return x
}
```
