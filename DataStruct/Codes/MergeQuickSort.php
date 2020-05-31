<?php

class SortHelper
{
    function MergeQuickSort($arr) {

        if (empty($arr)) {
            return [];
        }
        $provit = $arr[0];
        $count = count($arr);
        if ($count == 1) {
            return $arr;
        }
        
        $left = [];
        $right = [];
        for ($i = 1; $i < $count; $i++) {
            if ($arr[$i] < $provit) {
                $left[] = $arr[$i];
            } else {
                $right[] = $arr[$i];
            }
        }
        echo(PHP_EOL."before".PHP_EOL);
        print_r($left);
        print_r($right);
        $left = $this->MergeQuickSort($left);
        $right = $this->MergeQuickSort($right);
        echo(PHP_EOL."after".PHP_EOL);
        print_r($left);
        print_r($right);
        print_r($provit);
        echo (PHP_EOL."end".PHP_EOL);
        return array_merge($left, [$provit], $right);
    }
}

$arr = array(2, 1, 5, 8, 3, 7);
$obj = new SortHelper();
$data = $obj->MergeQuickSort($arr);
var_dump($data);
