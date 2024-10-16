<?php

//https://www.codewars.com/kata/556deca17c58da83c00002db

function tribonacci($signature, $n) {
    $arr = array_merge([], $signature);

    for($i = 0; $i < $n - 3; $i++){
        $arr[] = $arr[$i] + $arr[$i+1] + $arr[$i+2];
    }

    return array_splice($arr, 0, $n);
}