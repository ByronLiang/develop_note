# Heap(堆) 数据结构笔记

## 基本

- 堆是一种完全二叉树;除了树的最后一层节点不需要是满的，其它的每一层从左到右都是满的。

- 以数组实现的二叉树; 每个树节点对应数组里的下标(index)

### 相关节点计算

1. 获取数组长度:len；共有`len/3`父树节点
2. 父节点的下标分布: 
    - 最后一个父节点下标(至少有一个子节点): `(len/2) - 1`；
    - 节点的父节点：`(index - 1) / 2`;
3. 子节点下标分布: `pNode`为父节点下标;
节点的左子节点: `(2 * pNode) + 1`; 节点的右子节点: `(2 * pNode) + 2`

4. 基于节点下标`index`获取其父节点的下标`pNode = (index - 1) / 2`

## 堆的分类

- 大堆/大根堆: 当父节点大于等于其左右子节点;

- 小堆/小根堆: 当父节点小于等于其左右子节点;

## 核心算法

若以小堆为例，从最底层父节点(下标最大值)向上遍历数据；

- 比较左右子节点, 选取较小值与其父节点进行比较，若子节点的值小于父节点，进行交换处理(上浮)
- 父节点的移除，需要重新进行堆的调整

- 新增节点, 需要进行上浮处理, 比较新增节点与其父节点的大小, 进行交互，并进行堆调整

## 应用

实现队列优先级

### 原理

- 按照小堆分类实现队列优先级: 优先级最高，处于堆的顶部; 

- 每次从堆的顶部消费消息(数组里首尾位置进行交换), 对于置换后，若影响堆结构，需要完成堆的调整(下沉处理)