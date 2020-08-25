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

    /**
     * 原始归并排序
     * 不断进行拆分
     * 完成排序逐一进行合并
     */
    public function OriginMergeSort($arr)
    {
        $len = count($arr);
        if ($len < 2) {
            // 终止再生成递归
            return $arr;
        }
        $middle = (int) ($len / 2);
        $left = array_slice($arr, 0, $middle);
        $right = array_slice($arr, $middle);
        $leftDiv = $this->OriginMergeSort($left);
        $rightDiv = $this->OriginMergeSort($right);
        $res = $this->merge($leftDiv, $rightDiv);
        return $res;
    }

    private function merge($left, $right)
    {
        echo "left".PHP_EOL;
        print_r($left);
        echo "right".PHP_EOL;
        print_r($right);
        while (count($left) > 0 && count($right) > 0) {
            if ($left[0] <= $right[0]) {
                $result[] = array_shift($left);
            } else {
                $result[] = array_shift($right);
            }
        }
        while (count($left))
            $result[] = array_shift($left);
        while (count($right))
            $result[] = array_shift($right);

        return $result;
    }

    public function segment(&$arr, $left, $right)
    {
        if ($left >= $right) return;

        $mid = $left + (int) (($right - $left) / 2);

        echo "left index: ". $left. " mid: ". $mid. " rig: ". $right.PHP_EOL;

        $this->segment($arr, $left, $mid);

        echo "finished left; process right; left: ". $left. " mid: ". $mid. " rig: ". $right.PHP_EOL;

        $this->segment($arr, $mid + 1, $right);

        echo "after left index: ". $left. " mid: ". $mid. " rig: ". $right.PHP_EOL;

        $this->MergeWithSegment($arr, $left, $mid, $right);
    }

    private function MergeWithSegment(&$arr, $left, $mid, $right)
    {
        $tempArr = [];
        $rightIndex = $mid + 1;
        $tempArrIndex = $left;
        $beginIndex = $left;

        // 左右切片分支进行对比 组成具备排序的临时数组
        while ($left <= $mid && $rightIndex <= $right) {
            // 比较大小 进入临时数组 则对应下标递进1
            if ($arr[$left] <= $arr[$rightIndex]) {
                $tempArr[$tempArrIndex] = $arr[$left];
                $left++;
            } else {
                $tempArr[$tempArrIndex] = $arr[$rightIndex];
                $rightIndex++;
            }
            $tempArrIndex++;
        }
        // 若仍剩余, 追加到临时数组末尾
        while ($left <= $mid) {
            $tempArr[$tempArrIndex] = $arr[$left];
            $left ++;
            $tempArrIndex ++;
        }
        // 右分支
        while ($rightIndex <= $right) {
            $tempArr[$tempArrIndex] = $arr[$rightIndex];
            $rightIndex ++;
            $tempArrIndex ++;
        }
        // 将临时数组的排序 对原数组进行调整处理
        while ($beginIndex <= $right) {
            $arr[$beginIndex] = $tempArr[$beginIndex];
            $beginIndex ++;
        }
    }
}

// $arr = array(2, 1, 5, 8, 3, 7);
$arr = [6, 5, 3, 1, 7, 2, 4];
$obj = new SortHelper();
// $data = $obj->MergeQuickSort($arr);
// $data = $obj->OriginMergeSort($arr);
$obj->segment($arr, 0, count($arr) - 1);
$data = $arr;
echo implode("-", $data).PHP_EOL;
// print_r($data);
return;
$data = $obj->bindSort($arr);
$obj->sortedArray = $data;
$res = $obj->find(7);
print_r($data);
print_r($res);
