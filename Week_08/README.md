# 学习笔记 Week08

## 位运算

### 运算符

|含义|运算符|示例|
|---|---|---|
|左移|<<|0011->0110|
|右移|>>|0110->0011|
|位或|\||0011\|1011->1011|
|位与|&|0011&1011->0011|
|取反|~|~0011->1100|
|异或|^|0011^1011->1000|

### 异或的特性

相同为0，不同为1
```
x^0=x
x^1s=~x //1s=~0
x^(~x)=1s
x^x=0
c=a^b -> a^c=b -> b^c=a
a^b^c=a^(b^c)=(a^b)^c
```

### 指定位置的位运算

右边n位清零

```
x&(~0<<n)
```

获取x的第n位置

```
(x>>n)&1
```

获取x的第n位幂值

```
x&(1<<n)
```

将第n位置为1

```
x|(1<<n)
```

仅将第n位置为0 

```
x&(~(1<<n))
```

将x最高位至第n位（含）清零

```
x&((1<<n)-1)
```

### 实战位运算要点

判断奇偶

```
x%2==1 -> (x&1) == 1
x%2==0 -> (x&1) == 0
```

x>>1 -> x/2

```
x=x/2 -> x=x>>1
mid=(left+right)/2 -> mid=(left+right)>>1
```

清零最低位的1

```
x=x&(x-1)
```

得到最低位的1

```
x&(-x)
```

x&(~x) => 0

## 布隆过滤器（bloom filter）

布隆过滤器是由一个很长的二进制向量和一系列随机映射函数组成，可以用于检索一个元素是否在一个集合中

优点:`空间效率和查询时间都远远超过一般的算法`

缺点:`有一定的误识别率和删除困难`

### 案列

* 比特币网络
* 分布式系统
* redis缓存
* 垃圾邮件、评论过滤

### 参考资料

https://www.cnblogs.com/cpselvis/p/6265825.html

https://blog.csdn.net/tianyaleixiaowu/article/details/74721877

https://shimo.im/docs/UITYMj1eK88JCJTH/read

https://www.geeksforgeeks.org/bloom-filters-introduction-and-python-implementation/

https://github.com/jhgg/pybloof

https://github.com/lovasoa/bloomfilter/blob/master/src/main/java/BloomFilter.java

https://github.com/Baqend/Orestes-Bloomfilter

## LRU Cache

要素

* 大小
* 替换策略

实现方式:`HashTable+Double LinkedList`

效率

* O(1）查询
* O(1）修改，更新

### 参考资料

https://shimo.im/docs/CoyPAyXooGcDuLQo/read

https://www.sqlpassion.at/archive/2018/01/06/understanding-the-meltdown-exploit-in-my-own-simple-words/

https://en.wikipedia.org/wiki/Cache_replacement_policies

## 排序

* 比较类排序
  * 通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn),因此也称为非线性时间比较类排序
  * 交换排序
    * 冒泡 `嵌套循环，每次查看相邻的元素，如果逆序则交换`
    * 快排 `分治的思想，数组取表哥pivot，将小元素放到pivot左边，大元素放到右边，然后依次对右边和右边的子数组继续快排；以达到整个数组有序`
  * 插入排序
    * 简单插入排序 `从前到后逐步构建有序序列；对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入`
    * 希尔排序
  * 选择排序
    * 简单选择排序 `每次找最小值，然后放到待排序数组的起始位置`
    * 堆排序 `1.数组元素依次建立小顶堆，2.依次取堆顶元素，并删除`
  * 归并排序 `分治思想，1.把长度为n的输入序列分成两个长度为n/2的子序列；2.对这两个子序列分别采用归并排序；3.将两个排序号的子序列合并成一个最终的排序序列`
    * 二路归并排序
    * 多路归并排序
* 非比较类排序
  * 不通过比较来决定元素间的相对次序，它可以突破基于比较排序的时间下界，以线性时间运行，因此也称为线性时间非比较类排序
  * 计数排序 `要求输入的数据必须是有确定范围的整数。将输入的数据值转化为键存储在额外开辟的数组空间中；然后依次把计数大于1的填充回原数组`
  * 桶排序 `假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或以递归方式继续使用桶排序进行排序）`
  * 基数排序 `按照地位先排序，然后收集；再按高位排序，然后再收集；依次类推，直到最高位。在有些属性有优先级顺序的情况下，先按低优先级排序，再按高优先级排序`

|排序方法|平均时间复杂度|最坏时间复杂度|最好时间复杂度|空间复杂度|稳定性|
|------|-----------|-------------|----------|-----|----|
|插入|O(n^2)|O(n^2)|O(n)|O(1)|稳定|
|希尔排序|O(n^1.3)|O(n^2)|O(n)|O(1)|不稳定|
|选择排序|O(n^2)|O(n^2)|O(n^2)|O(1)|不稳定|
|堆排序|O(nlogn)|O(nlogn)|O(nlogn)|O(1)|不稳定|
|冒泡排序|O(n^2)|O(n^2)|O(n)|O(1)|稳定|
|快速排序|O(nlogn)|O(n^2)|O(nlogn)|O(nlogn)|不稳定|
|归并排序|O(nlogn)|O(nlogn)|O(nlogn)|O(n)|稳定|
|计数排序|O(n+k)|O(n+k)|O(n+k)|O(n+k)|稳定|
|桶排序|O(n+k)|O(n^2)|O(n)|O(n+k)|稳定|
|基数排序|O(n*k)|O(n*k)|O(n*k)|O(n+k)|稳定|

### 参考链接

https://www.cnblogs.com/onepixel/p/7674659.html

https://shimo.im/docs/TX9bDbSC7C0CR5XO/read

https://shimo.im/docs/sDXxjjiKf3gLVVAU/read

https://shimo.im/docs/M2xfacKvwzAykhz6/read

https://www.bilibili.com/video/av25136272

https://www.bilibili.com/video/av63851336