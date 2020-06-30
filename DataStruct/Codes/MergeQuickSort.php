<?php

class SortHelper
{
    /**
     * 已排序数组
     */
    public $sortedArray;
    
    /**
     * 快速排序
     */
    function MergeQuickSort($arr)
    {
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

    /**
     * 冒泡排序
     */
    public function bindSort(array &$arr)
    {
        for ($i = 0; $i < count($arr) - 1; $i++) {
            for ($j = 0; $j < count($arr) - 1 - $i; $j++) {
                if ($arr[$j] > $arr[$j + 1]) {
                    list($arr[$j], $arr[$j+1]) = [$arr[$j + 1], $arr[$j]];
                }
            }
        }

        return $arr;
    }

    /**
     * 二分法查找
     */
    public function find(int $target)
    {
        return $this->repFind($target, 0, count($this->sortedArray) - 1);
    }

    /**
     * 递归处理二分法
     */
    protected function repFind($target, $left, $right)
    {
        if ($left > $right) {
            return false;
        }
        $mid = $left + (int) (($right - $left) / 2);
        // $mid = (int) (($right + $left) / 2);
        if ($this->sortedArray[$mid] === $target) {
            return $mid;
        }
        if ($this->sortedArray[$mid] > $target) {
            // 向左缩小范围
            return $this->repFind($target, $left, $mid - 1);
        } else {
            // 向右缩小范围
            return $this->repFind($target, $mid + 1, $right);
        }
    }
}

$arr = array(2, 1, 5, 8, 3, 7);
$obj = new SortHelper();
// $data = $obj->MergeQuickSort($arr);

$data = $obj->bindSort($arr);
$obj->sortedArray = $data;
$res = $obj->find(7);
print_r($data);
print_r($res);
