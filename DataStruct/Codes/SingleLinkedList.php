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
}

$singleLinkedList = new SingleLinkedList();

$nodes = [2, 5, 8];
foreach ($nodes as $value) {
    $node = new Node($value);
    $singleLinkedList->add($node);
}

$singleLinkedList->getAll();
$singleLinkedList->reverse();
$singleLinkedList->getAll();