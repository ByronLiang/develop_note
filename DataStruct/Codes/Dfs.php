<?php

class dfs {

    public $data = [];

    public function __construct() {
        $this->data = json_decode(file_get_contents('./data/sample.json'), true);
    }

    public function dfsMap()
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

class Solution {

    protected $result = [];

    /**
     * https://leetcode-cn.com/problems/permutations
     * 
     * 全排列: 回溯-递归
     */
    function permute($nums) {
        $count = count($nums);
        if ($count == 0) return $this->result; 
        $this->backtrackPermute($nums, 0, []);
        return $this->result;
    }

    private function backtrackPermute($nums, $depth, $path)
    {
        if ($depth == count($nums)) {
            $this->result[] = $path;
            return;
        }

        for ($i = 0; $i < count($nums); $i++) {
            if (in_array($nums[$i], $path)) continue;
            $path[] = $nums[$i];
            echo "before-i: ". $i. " ". $depth. " path: ". implode("-", $path). PHP_EOL;
            $this->backtrackPermute($nums, $depth + 1, $path);
            echo "after-i: ". $i. " ". $depth. " path: ". implode("-", $path). PHP_EOL;
            // 回溯，恢复状态
            array_pop($path);
        }
        echo "end-i: ". $i. " ". $depth. " path: ". implode("-", $path). PHP_EOL;
    }
}

// $obj = new dfs();
// $res = $obj->dfsMap();
// print_r($res);

$obj = new Solution();
$res = $obj->permute([1, 2, 4]);
print_r($res);
