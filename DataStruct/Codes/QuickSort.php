<?php

function swap(array &$arr, $a, $b){
    $temp = $arr[$a];
    $arr[$a] = $arr[$b];
    $arr[$b] = $temp;
}

function Partition(array &$arr, $low, $high){
    // $high 尾部下标 $low 头部下标
    $pivot = $arr[$low];   //选取子数组第一个元素作为枢轴
    while($low < $high){  //从数组的两端交替向中间扫描
        while($low < $high && $arr[$high] >= $pivot) {
            $high --;
        }
        echo PHP_EOL."before hight swap low-val: ". $arr[$low]. " hight-val: ". $arr[$high].PHP_EOL;
        echo PHP_EOL."before hight swap ". $low. " hight: ". $high.PHP_EOL;
        swap($arr, $low, $high);    //终于遇到一个比$pivot小的数，将其放到数组低端
        echo "hight swap ". implode("-", $arr).PHP_EOL;
        while($low < $high && $arr[$low] <= $pivot) {
            $low ++;
        }
        echo PHP_EOL."before low swap low-val: ". $arr[$low]. " hight-val: ". $arr[$high].PHP_EOL;
        echo PHP_EOL."before low swap ". $low. " hight: ". $high.PHP_EOL;
        swap($arr, $low, $high);    //终于遇到一个比$pivot大的数，将其放到数组高端
        echo "low swap ". implode("-", $arr).PHP_EOL;
    }

    return $low;   //返回high也行，毕竟最后low和high都是停留在pivot下标处
}

function QSort(array &$arr, $low, $high){
    if ($low < $high) {
        $mid = Partition($arr, $low, $high);  //将$arr[$low...$high]一分为二，算出枢轴值
        echo "mid ". $mid.PHP_EOL;
        QSort($arr, $low, $mid - 1);   //对低子表进行递归排序(左边界)
        QSort($arr, $mid + 1, $high);  //对高子表进行递归排序(有边界)
    }
}

function QuickSort(array &$arr){
    $low = 0;
    $high = count($arr) - 1;
    QSort($arr, $low, $high);
    
    return $arr;
}

// $arr = array(9, 1, 5, 8, 3);
$arr = [6, 5, 3, 1, 7, 2, 4];
$data = QuickSort($arr);
echo implode("-", $data).PHP_EOL;
// var_dump($data);
