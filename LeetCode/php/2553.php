<?php

function separateDigits2553_v2($nums){
    $result = [];

    for($i = count($nums) - 1; $i >= 0; $i--){
        $n = $nums[$i];
        while($n != 0){
            $result[] = $n % 10;
            $n = (int)($n / 10);
        }
    }

    return array_reverse($result);
}

function separateDigits2553_v1($nums) {
    $result = [];

    for ($i = 0; $i < count($nums); $i++) {
        $num = $nums[$i];
        $temp = [];

        while($num != 0){
            array_unshift($temp, $num % 10);
            $num = (int)($num / 10);
        }

        for($j = 0; $j < count($temp); $j++){
            $result[] = $temp[$j];
        }
    }

    return $result;
}

print_r(separateDigits2553_v2([13,25,83,77]));