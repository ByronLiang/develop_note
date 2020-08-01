<?php

class Node
{
    public $data, $next;

    public function __construct($data) {
        $this->data = $data;
        $this->next = null;
    }
}

class SingleLinkedList
{
    public $header, $last, $size;

    public function __construct() {
        $this->header = null;
        $this->last = null;
        $this->size = 0;
    }

    public function add(Node $node) {
        if ($this->header == null && $this->last == null) {
            $this->header = $node;
            $this->last = $node;
        } else {
            $this->last->next = $node;
            $this->last = $node;
        }
        $this->size +=1;
    }

    public function getSize() :int
    {
        return $this->size;
    }

    public function findNode($data)
    {
        $node = $this->header;
        if ($node->data == $data) {
            return $data;
        }
        while($node->next != null) {
            if ($node->next->data == $data) {
                return $data;
            }
            $node = $node->next;
        }

        echo 'unfound'. "\n";
    }

    /**
     * 迭代翻转链表
     * 翻转后, last属性是作为head属性来使用
     */
    public function reverse()
    {
        // 保存上一节点的指向对象
        $prev = null;
        // 迭代对象
        $current = $this->header;
        // 暂存当前节点的下一节点对象
        $next = null;
        while($current != null) {
            // 暂存当前节点的下一个节点
            $next = $current->next;
            // 将当前节点的下一节点指向上一个节点
            $current->next = $prev;
            // 保存当前节点 以作为下一节点提供指向节点的next对象
            $prev = $current;
            // 下一节点继续完成迭代业务
            $current = $next;
        }
        $this->header = $prev;
    }

    /**
     * 递归翻转链表
     */
    public function reverseMap(Node $node)
    {
        // 中止产生递归条件: 已经遍历到链表尾部
        if ($node == null || $node->next == null) {
            echo "emd map loop".$node->data.PHP_EOL;
            return $node;
        }
        $this->reverseMap($node->next);
        echo "current: " . $node->data." next: ". $node->next->data .PHP_EOL;
        // 将一层层递归完成;
        $node->next->next = $node;
        $node->next = null;
    }

    public function getAll($node)
    {
        while($node->next != null) {
            echo $node->data . "\n";
            $node = $node->next;
        }
        echo "last node: ". $node->data ."\n";
    }

    /**
     * 进行闭环处理
     *
     * @param $node 设置尾部节点
     * @param $circlePointData 对指定节点的值建立闭环
     */
    public function addCirclePoint(Node $node, $circlePointData)
    {
        if ($this->header == null && $this->last == null) {
            return false;
        }
        $this->last->next = $node;
        $this->last = $node;
        $this->size +=1;
        // 开始对指定节点进行闭环处理
        $node = $this->header;
        if ($node->data == $circlePointData) {
            $this->last->next = $node;
            return true;
        }
        while($node->next != null) {
            $target = $node->next;
            if ($target->data == $circlePointData) {
                $this->last->next = $target;
                return true;
            }
            $node = $node->next;
        }
    }

    /**
     * 判断是否闭环
     * 
     * @return array 返回判断结果及返回相遇节点
     */
    public function judgeCycle()
    {
        $fast_node = $slow_node = $this->header;
        while ($fast_node != null && $fast_node->next != null) {
            $fast_node = $fast_node->next->next;
            $slow_node = $slow_node->next;
            if ($fast_node->data == $slow_node->data) return [1, $fast_node];
        }

        return [0, null];
    }

    /**
     * 查询闭环节点的节点值
     */
    public function findCycleNode($fastNode)
    {
        $slowNode = $this->header;
        while ($slowNode != $fastNode) {
            $fastNode = $fastNode->next;
            $slowNode = $slowNode->next;
        }

        return $slowNode;
    }

    public function removeOneRepeatMap(Node $node)
    {
        if ($node == null || $node->next == null) {
            return $node;
        }
        echo "ff ". $node->data .PHP_EOL;
        if ($node->data == $node->next->data) {
            $node = $this->removeOneRepeatMap($node->next);
            echo "xx ". $node->data .PHP_EOL;
        } else {
            $node->next = $this->removeOneRepeatMap($node->next);
            echo "yy ". $node->data .PHP_EOL;
        }
        // echo $node->data .PHP_EOL;
        return $node;
    }

    public function removeOneRepeat()
    {
        $node = $this->header;
        while($node != null) {
            while ($node->next != null) {
                if ($node->next->data == $node->data) {
                    $node->next = $node->next->next;
                } else {
                    break;
                }
            }
            $node = $node->next;
        }
    }

    public function removeAllRepeated()
    {
        $temp = new Node(0);
        $temp->next = $this->header;
        $fast = $temp->next;
        $slow = $temp;
        while($fast != null) {
            if ($fast->next != null && $fast->next->data == $fast->data) {
                while ($fast->next != null && $fast->next->data == $fast->data) {
                    $fast = $fast->next;
                }
                $slow->next = $fast->next;
                $fast = $fast->next;
            } else {
                $slow = $slow->next;
                $fast = $fast->next;
            }
        }
        return $temp->next;
    }
}

$singleLinkedList = new SingleLinkedList();

$nodes = [2, 5, 5, 8, 8, 10, 15];
foreach ($nodes as $value) {
    $node = new Node($value);
    $singleLinkedList->add($node);
}

// 递归翻转测试
// $headNode = $singleLinkedList->header;
// $singleLinkedList->reverseMap($headNode);
// // 当尾部节点作为头节点属性 进行遍历显示节点数据
// $node = $singleLinkedList->last;
// $singleLinkedList->getAll($node);

// 移除重复数值
$singleLinkedList->removeOneRepeatMap($singleLinkedList->header);
// $singleLinkedList->removeOneRepeat();
// $node = $singleLinkedList->removeAllRepeated();
// $singleLinkedList->getAll($node);

// 闭环操作
// $singleLinkedList->addCirclePoint(new Node(10), 5);
// list($res, $fastNode) = $singleLinkedList->judgeCycle();
// echo $res."\n";
// if ($res > 0) {
//     $node = $singleLinkedList->findCycleNode($fastNode);
//     echo "cycle point data: ".$node->data.PHP_EOL;
// }

// $headNode = $singleLinkedList->header;
// $singleLinkedList->getAll($headNode);
// $singleLinkedList->reverse();
// $singleLinkedList->getAll();