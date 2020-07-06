<?php

class tree
{
    public $data;
    public $left =  null ;
    public $right = null;
    public function __construct($data) 
    {
        $this->data = $data;
    }

    // DLR
    public function preOrder()
    {
        echo $this->data." ";
        if($this->left)
            $this->left->preOrder();
        if($this->right)
            $this->right->preOrder();
    }
    // LDR
    public function inOrder()
    {
        if($this->left)
            $this->left->inOrder();
        echo $this->data." ";
        if($this->right)
            $this->right->inOrder();
    }
    // LRD
    public function postOrder()
    {
        if($this->left)
            $this->left->postOrder();
        if($this->right)
            $this->right->postOrder();
        echo $this->data." ";
    }

    public function printPath($tree, $target)
    {
        $res = [];
        if ($this->hasPath($tree, $res, $target)) {
            $data = implode('->', $res);
            echo $data . "\n";
        } else {
            echo "no find target". "\n";
        }
    }

    protected function hasPath($tree, &$res, $target)
    {
        // 没有下一级节点 结束递归
        if ($tree == null) {
            return false;
        }
        // 每经过一个节点, 都记录踪迹
        array_push($res, $tree->data);
        if ($tree->data == $target) {
            return true;
        }
        // 判断当前节点左右节点 并进行递归处理
        if ($this->hasPath($tree->left, $res, $target) ||
            $this->hasPath($tree->right, $res, $target)) {
            return true;
        }
        // 路径无法匹配目标, 移除当前入栈的路径
        array_pop($res);
        // 继续遍历同级其他节点 或者退出当前递归，返回上一层递归路线
        // 待定
        return false;
    }
}

class Solution
{
    /**
     * 
     * 计算二叉树路径节点的和满足目标数值(sum)
     * 
     * https://leetcode.com/problems/path-sum-iii
     */
    public function pathSum($root, int $sum)
    {
        if ($root === null) {
            return 0;
        }
        echo "process ". $root->data.PHP_EOL;
        $pathImLeading = $this->countSum($root, $sum); // 自己为开头的路径数
        echo $pathImLeading. " what ". $root->data.PHP_EOL;
        $leftPathSum = $this->pathSum($root->left, $sum); // 左边路径总数（相信他能算出来）
        $rightPathSum = $this->pathSum($root->right, $sum); // 右边路径总数（相信他能算出来）
        echo $root->data. " d". $pathImLeading.PHP_EOL;
        // print_r($pathImLeading);
        // print_r($leftPathSum);
        // print_r($rightPathSum);
    }

    public function countSum($node, $sum)
    {
        // 终止递归条件
        if ($node === null) return 0;
        echo "count target ". $node->data. " sum is ". $sum. PHP_EOL;
        // 进行目标数值匹配
        $isMe = ($node->data == $sum) ? 1 : 0;
        // 左叉子树递归查询，对剩余数值进行计算($sum - $node->data)
        $leftBrother = $this->countSum($node->left, $sum - $node->data);
        // 右叉子树递归查询 对剩余数值进行计算($sum - $node->data)
        $rightBrother = $this->countSum($node->right, $sum - $node->data);
        
        $res = $isMe + $leftBrother + $rightBrother;
        // 结束递归
        echo "target : ".$node->data." res ".$res.PHP_EOL;
        return $res; 
    }
}


// $trees = new tree(8);
// $trees->left =  new tree(3);
// $trees->left->left =  new tree(1);
// $trees->left->right = new tree(6);
// $trees->left->right->left = new tree(4);
// $trees->left->right->right = new tree(7);

$trees->right =  new tree(10);
$trees->right->right = new tree(14);
$trees->right->right->left =  new tree(13);

$trees = new tree(10);
$trees->left =  new tree(5);
$trees->right =  new tree(-3);
$trees->left->left =  new tree(3);
$trees->left->right = new tree(2);
$trees->right->right =  new tree(11);
$trees->left->left->left =  new tree(3);
$trees->left->left->right = new tree(-2);
$trees->left->right->right = new tree(1);

// $slo = new Solution();
// $slo->pathSum($trees, 8);

$trees->preOrder();
echo "\n";
$trees->inOrder();
echo "\n";
$trees->postOrder();
echo "\n";

// 打印寻数路径
$trees->printPath($trees, $argv[1]);
