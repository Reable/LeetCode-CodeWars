<?php 

//https://www.codewars.com/kata/556e0fccc392c527f20000c5

function Xbonacci_v1($s, $n) {
    if ($n <= count($s)) {
        return array_slice($s, 0, $n);
    }

    $arr = array_merge([], $s);
    
    for($i = 0; $i < $n - count($s); $i++){
        $v = 0;
        foreach(array_slice($arr, $i, $i + count($s)) as $num){
            $v += $num;
        }
        array_push($arr, $v);
    }

    return $arr;
}

function Xbonacci_v2($s, $n) {
    if ($n <= count($s)) {
        return array_slice($s, 0, $n);
    }

    $arr = array_merge([], $s);
    
    for($i = 0; $i < $n - count($s); $i++){
        $arr[] = array_sum(array_slice($arr, $i, count($arr)));
    }

    return $arr;
}