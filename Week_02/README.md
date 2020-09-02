# 学习笔记 Week2

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
