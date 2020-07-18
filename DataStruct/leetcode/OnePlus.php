<?php

class Solution {
    /**
     * @param Integer[] $digits
     * @return Integer[]
     */
    function plusOne($digits) {
        for ($i = count($digits) - 1; $i >= 0; $i--) {
            $digits[$i]++;
            $digits[$i] = (int) ($digits[$i] % 10);
            if ($digits[$i] != 0) {
                break;
            }
        }
        if ($digits[0] == 0) {
            $digits = array_merge([1], $digits);
        }
        print_r($digits);
    }
}

$s = new Solution();

$s->plusOne([2, 9, 9, 1]);