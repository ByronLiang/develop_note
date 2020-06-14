<?php

class Tree
{
    public $data;
    public $left =  null ;
    public $right = null;
    public function __construct($data) 
    {
        $this->data = $data;
    }
}

class Solution
{
    protected $minPre = 0;

    /**
     * 检验是否二叉树(BST)
     */
    public function verify($tree = null)
    {
        if ($tree == null) {
            return true;
        }
        // 访问左子树
        if (!$this->verify($tree->left)) {
            return false;
        }
        echo $this->minPre .PHP_EOL;
        // 访问当前节点：如果当前节点小于等于中序遍历的前一个节点，说明不满足BST，返回 false；否则继续遍历。
        if ($tree->data <= $this->minPre) {
            echo $tree->data." error ".$this->minPre.PHP_EOL;
            return false;
        }
        $this->minPre = $tree->data;
        echo $tree->data .PHP_EOL;
        // 访问右子树
        return $this->verify($tree->right);
    }
}

$trees = new Tree(5);
$trees->left =  new Tree(1);
$trees->right = new Tree(4);
$trees->right->left = new Tree(3);
$trees->right->right = new Tree(6);

$sol = new Solution();
$res = $sol->verify($trees);
var_dump($res);



