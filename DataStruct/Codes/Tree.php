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

$trees = new tree(8);
$trees->left =  new tree(3);
$trees->left->left =  new tree(1);
$trees->left->right = new tree(6);
$trees->left->right->left = new tree(4);
$trees->left->right->right = new tree(7);

$trees->right =  new tree(10);
$trees->right->right = new tree(14);
$trees->right->right->left =  new tree(13);

$trees->preOrder();
echo "\n";
$trees->inOrder();
echo "\n";
$trees->postOrder();
echo "\n";

// 打印寻数路径
$trees->printPath($trees, $argv[1]);