<?php

function countGoodSubstrings1876_2($s) {

    if(strlen($s) < 3) return 0;

    $result = 0;
    $temp = [];

    for($i = 0; $i < 3; $i++) {
        $temp[$s[$i]] = ($temp[$s[$i]] ?? 0) + 1;
    }

    if(count($temp) === 3) $result += 1;

    for($i = 3; $i < strlen($s); $i++) {
        $temp[$s[$i-3]] -= 1;
        if($temp[$s[$i-3]] === 0) unset($temp[$s[$i-3]]);

        $right = $s[$i];
        $temp[$right] = ($temp[$right] ?? 0) + 1;

        if(count($temp) === 3) $result += 1;
    }

    return $result;
}

function countGoodSubstrings1876_1($s) {
    if(strlen($s) < 3) return 0;

    $result = 0;

    $temp = [];
    $window = substr($s, 0 , 3);

    for($i = 0; $i < strlen($window); $i++) {
        if(isset($temp[$window[$i]])) {
            $temp[$window[$i]] += 1;
        } else {
            $temp[$window[$i]] = 1;
        }
    }
    if(self::checkUnique($temp)) $result += 1;

    for($i = 1; $i < strlen($s)-2; $i++) {
        $temp[$s[$i - 1]] -= 1;
        if(isset($temp[$s[$i+2]])) {
            $temp[$s[$i+2]] += 1;
        } else {
            $temp[$s[$i+2]] = 1;
        }

        if(self::checkUnique($temp)) $result += 1;
    }

    return $result;
}

function checkUnique($temp){
    foreach($temp as $k=>$v){
        if($v >= 2){
            return false;
        }
    }
    return true;
}

echo countGoodSubstrings("aababcabc");