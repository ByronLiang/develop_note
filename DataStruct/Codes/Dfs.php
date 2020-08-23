<?php

class dfs {

    public $data = [];

    public function __construct() {
        $this->data = json_decode(file_get_contents('./data/sample.json'), true);
    }

    public function DfsMap()
    {
        $stack = [];
        $list = [];
        $stack[] = $this->data;
        while (count($stack) > 0) {
            $node = $stack[count($stack) - 1];
            $list[] = $node['name'];
            $stack = array_slice($stack, 0, count($stack) - 1);
            if (isset($node['children']) && count($node['children'])) {
                $childTotal = count($node['children']);
                for ($i = ($childTotal - 1); $i >= 0; $i--) {
                    $stack[] = $node['children'][$i]; 
                }
            }
        }
        return $list;
    }
}

$obj = new dfs();
$res = $obj->DfsMap();
print_r($res);
