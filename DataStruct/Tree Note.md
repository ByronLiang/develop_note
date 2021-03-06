# 树数据结构记录

## 基本定义

1. 树是n个结点的有限集. n=0, 称为空树；在任意一颗非空树中，有且仅有一个特定的称为根节点(root)

2. 其余结点可分为m个互不相交的有限集, 每一个集合本身又是一颗树, 称为根的子树(SubTree)

3. 结点度：结点拥有的子树数；若结点度为0, 则为终端结点/叶结点；结点度大于0, 非终端结点/分支结点

4. 树的深度`Depth`：对于任意节点n,n的深度为从根到n的唯一路径长，根的深度为0;

5. 树的高度`height`: 对于任意节点n,n的高度为从n到一片树叶的最长路径长，所有终端叶结点的高度为0；

## 二叉树基本理论及其遍历

### 遍历

- 先序遍历: 先访问根结点, 然后前序遍历左子树, 再前序遍历右子树
- 中序遍历: 遍历根结点的左子树，访问根结点, 遍历右子树 `遍历出来的数据具备升序特性`
- 后序遍历: 遍历访问左右子树，最后访问根结点

## 线段树 segment tree

主要处理数组区间分段与区间总和问题

### 数据结构

1. 完全二叉树来存储对应于其每一个区间（segment）的数据

2. 每一个结点中保存着相对应于这一个区间的信息。使用一个数组保存的，与二叉堆（Heap）的实现方式相同


