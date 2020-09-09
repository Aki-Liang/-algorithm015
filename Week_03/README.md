# 学习笔记 Week 03

## 递归

通过函数体实现的循环

### 思维要点

* 不要进行人肉递归
* 寻找最近子问题
* 数学归纳法思维

### 递归公式

```
func recurision() {
    if 满足结束条件 {
        return resule
    }

    处理本层逻辑

    recusion()

    清理当前层
}
```

1. recursion terminator,先写退出条件
2. process logic incurrent level,处理本层逻辑
3. drill down，下探到下一层
4. reverse the current level status if needed，退出时清理本层

## 分治， 回溯

本质就是递归

### 分治

把复杂问题拆分成子问题

#### 代码模板

```
func divide_conquer(问题) {
    if 满足结束条件 {
        return result
    }

    准备数据

    res1 := divide_conquer(子问题1)
    res2 := divide_conquer(子问题2)
    res3 := divide_conquer(子问题3)

    组装结果
    退出清理
}
```

1. recursion terminator，先写退出条件
2. prepare data，准备数据
3. conquer subproblems，处理子问题
4. process and generate the final result,组装最终结果
5. reverse the current level states if needed，清理当前层

### 回溯

通过试错找到正确的解

#### 代码模板


```
result = []
func backtrack(路径，选择列表) {
    if 满足结束条件 {
        result.add(路径)
        return
    }

    for 选择 in 选择列表 {
        做选择
        backtrack(路径，选择列表)
        撤销选择
    }
}
```