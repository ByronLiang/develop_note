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

    public function reverse()
    {
        $prev = null;
        $current = $this->header;
        $next = null;
        while($current != null) {
            $next = $current->next;
            $current->next = $prev;
            $prev = $current;
            $current = $next;
        }
        $this->header = $prev;
    }

    public function getAll()
    {
        $node = $this->header;
        while($node->next != null) {
            echo $node->data . "\n";
            $node = $node->next;
        }
        echo "last node: ". $node->data ."\n";
    }

    /**
     * 进行闭环处理
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
     */
    public function judgeCycle()
    {
        $fast_node = $slow_node = $this->header;
        while ($fast_node != null && $fast_node->next != null) {
            $fast_node = $fast_node->next->next;
            $slow_node = $slow_node->next;
            if ($fast_node->data == $slow_node->data) return 1;
        }

        return 0;
    }
}

$singleLinkedList = new SingleLinkedList();

$nodes = [2, 5, 8, 10, 15];
foreach ($nodes as $value) {
    $node = new Node($value);
    $singleLinkedList->add($node);
}
// 闭环操作
// $singleLinkedList->addCirclePoint(new Node(10), 5);
// $res = $singleLinkedList->judgeCycle();
// echo $res."\n";

// $singleLinkedList->getAll();
// $singleLinkedList->reverse();
// $singleLinkedList->getAll();