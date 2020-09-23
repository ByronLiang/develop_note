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
    public $math = [];

    public $res = [];
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

    public function rangeSumBST($node, $left, $right)
    {
        if ($node == null) {
            return 0;
        }
        if ($node->data > $right) {
            return $this->rangeSumBST($node->left, $left, $right);
        }
        if ($node->data < $left) {
            return $this->rangeSumBST($node->right, $left, $right);
        }
        $this->math[] = $node->data;
        $leftSum = $this->rangeSumBST($node->left, $left, $right);
        $rightSum = $this->rangeSumBST($node->right, $left, $right);

        return $node->data + $leftSum + $rightSum;
    }

    /**
     * 行遍历
     */
    public function loadBST($node)
    {
        $queue = array();
        $val = array();
        $queue[] = $node;
        $height = 0;
        $heightDetail = array();
        for (; count($queue) > 0 ;) {
            $height++;
            $heightNodeLength = count($queue);
            $heightDetail[$height] = $heightNodeLength;
            for ($i=0; $i < $heightNodeLength; $i++) { 
                $val[] = $queue[$i]->data;
                $leftNode = $queue[$i]->left;
                $rightNode = $queue[$i]->right;
                if ($leftNode != null) {
                    $queue[] = $leftNode;
                }
                if ($rightNode != null) {
                    $queue[] = $rightNode;
                }
            }
            $queue = array_slice($queue, $heightNodeLength);
        }

        return [$val, $heightDetail];
    }

    /**
     * 交叉行遍历
     */
    public function crossBST($node)
    {
        $queue = array();
        $val = array();
        $queue[] = $node;
        $height = 0;
        $heightDetail = array();
        for (; count($queue) > 0 ;) {
            $height++;
            $heightNodeLength = count($queue);
            $heightDetail[$height]['num'] = $heightNodeLength;
            for ($i=0; $i < $heightNodeLength; $i++) {
                if ($height % 2 == 0) {
                    $j = $heightNodeLength - 1 - $i;
                    $heightDetail[$height]['data'][] = $queue[$j]->data;
                } else {
                    $heightDetail[$height]['data'][] = $queue[$i]->data;
                }
                $leftNode = $queue[$i]->left;
                $rightNode = $queue[$i]->right;
                if ($leftNode != null) {
                    $queue[] = $leftNode;
                }
                if ($rightNode != null) {
                    $queue[] = $rightNode;
                }
            }
            $queue = array_slice($queue, $heightNodeLength);
        }

        return $heightDetail;
    }

    /**
     * 中序迭代遍历
     */
    public function midBST($node) {
        $queue = array();
        $val = array();
        while ($node != null | count($queue) > 0) {
            while ($node != null) {
                // 继续入栈
                $queue[] = $node;
                // 只向左分支入栈处理
                $node = $node->left;
            }
            // 出栈
            $currentNode = $queue[count($queue) - 1];
            $queue = array_slice($queue, 0, count($queue) - 1);
            $val[] = $currentNode->data;
            // 右分支树节点
            $node = $currentNode->right;
        }

        return $val;
    }

    /**
     * 最大深度
     */
    public function findMaxDepth($node) {
        if ($node == null) {
            return 0;
        }
        $leftNodeDepth = $this->findMaxDepth($node->left);
        echo "left: ". $leftNodeDepth. " node: " . $node->data . PHP_EOL;
        $rightNodeDepth = $this->findMaxDepth($node->right);
        echo "right: ". $rightNodeDepth. " node: " . $node->data . PHP_EOL;
        
        $res = max($leftNodeDepth, $rightNodeDepth) + 1;
        echo "res: ". $res. " node: " . $node->data . PHP_EOL;
        return $res;
    }

    /**
     * 最小深度
     */
    public function findMinDepth($node) {
        if ($node == null) {
            return 0;
        }
        $leftNodeDepth = $this->findMinDepth($node->left);
        echo "left: ". $leftNodeDepth. " node: " . $node->data . PHP_EOL;
        $rightNodeDepth = $this->findMinDepth($node->right);
        echo "right: ". $rightNodeDepth. " node: " . $node->data . PHP_EOL;
        if ($node->left == null || $node->right == null) {
            $res = $leftNodeDepth + $rightNodeDepth + 1;
        } else {
            $res = min($leftNodeDepth, $rightNodeDepth) + 1;
        }
        echo "res: ". $res. " node: " . $node->data . PHP_EOL;
        return $res;
    }

    /**
     * 有序数组建立二叉树(BST建立)
     */
    public function sortedArrayToBST($target)
    {
        return count($target) == 0 ? null : $this->buildTree($target, 0, count($target) - 1);
    }

    private function buildTree($target, $left, $right)
    {
        if ($left > $right) {
            return null;
        }
        $mid = $left + (int) (($right - $left) / 2);
        echo "left: ". $left . " right: ". $right ." index: ". $mid. " val ". $target[$mid].PHP_EOL;
        $node = new tree($target[$mid]);
        $node->left = $this->buildTree($target, $left, $mid - 1);
        echo "processing right: ". $mid. " val ". $right.PHP_EOL;
        $node->right = $this->buildTree($target, $mid + 1, $right);

        return $node;
    }

    public function printAllPath($root, &$path)
    {
        if ($root->left == null & $root->right == null) {
            $path[] = $root->data;
            $this->res[] = implode("->", $path);
            array_pop($path);
            return;
        }
        $path[] = $root->data;
        echo "current: ". implode("->", $path).PHP_EOL;
        if ($root->left != null) {
            $this->printAllPath($root->left, $path);
        }
        if ($root->right != null) {
            $this->printAllPath($root->right, $path);   
        }
        array_pop($path);
        echo "finished processed: ". implode("->", $path).PHP_EOL;
    }

    /**
     * 镜像二叉树
     */
    public function mirrorTree($node)
    {
        if ($node == null) {
            return;
        }
        if ($node->left == null && $node->right == null) {
            return;
        }
        $temp = $node->left;
        $node->left = $node->right;
        $node->right = $temp;
        if ($node->left != null) {
            $this->mirrorTree($node->left);
        }
        if ($node->right != null) {
            $this->mirrorTree($node->right);
        }
    }
}

$slo = new Solution();

// $trees = new tree(8);
// $trees->left =  new tree(3);
// $trees->left->left =  new tree(1);
// $trees->left->right = new tree(6);
// $trees->left->right->left = new tree(4);
// $trees->left->right->right = new tree(7);
// $trees->right =  new tree(10);
// $trees->right->right = new tree(14);
// $trees->right->right->left =  new tree(13);

// $trees = new tree(10);
// $trees->left =  new tree(5);
// $trees->right =  new tree(-3);
// $trees->left->left =  new tree(3);
// $trees->left->right = new tree(2);
// $trees->right->right =  new tree(11);
// $trees->left->left->left =  new tree(3);
// $trees->left->left->right = new tree(-2);
// $trees->left->right->right = new tree(1);


// $slo->pathSum($trees, 8);

// $trees->preOrder();
// echo "\n";
// $trees->inOrder();
// echo "\n";
// $trees->postOrder();
// echo "\n";

// // 打印寻数路径
// $trees->printPath($trees, $argv[1]);


/**
 *          10
 *         /  \
 *        5    15
 *       / \     \
 *      3   7    18
 */

$trees = new tree(10);
$trees->left =  new tree(5);
$trees->right =  new tree(15);
$trees->left->left =  new tree(3);
$trees->left->right = new tree(7);
$trees->right->right =  new tree(18);

// $res = $slo->rangeSumBST($trees, 1, 14);
// print_r($slo->math);
// print_r($res);

// list($vals, $levels) = $slo->loadBST($trees);
// $vals = $slo->midBST($trees);
$vals = $slo->findMaxDepth($trees);
echo "max-depth: ".$vals .PHP_EOL;
$vals = $slo->findMinDepth($trees);
echo "min-depth: ".$vals .PHP_EOL;
// print_r($levels);
// print_r($slo->crossBST($trees));
$path = [];
// 镜像处理
$slo->mirrorTree($trees);
$slo->printAllPath($trees, $path);
print_r($slo->res);
return;

$target = [3, 5, 7, 10, 15, 18];
$trees = $slo->sortedArrayToBST($target);
list($vals, $levels) = $slo->loadBST($trees);
print_r($vals);
print_r($levels);
