<?php

function longestOnes1004_v2($nums, $k) {
    $l = 0;
    $r = 0;
    $zeros = 0;
    $result = 0;

    while($r < count($nums)) {
        if($nums[$r] == 0) $zeros += 1;

        while($zeros > $k) {
            if($nums[$l++] == 0){
                $zeros -= 1;
            }
        }

        $current = $r - $l + 1;
        if($current > $result) {
            $result = $current;
        }

        $r++;
    }

    return $result;
}


function longestOnes1004_v1($nums, $k) {
    $l = 0;
    $r = 0;
    $zCount = 0;
    $result = 0;

    while($l < count($nums)) {
        while($r < count($nums) && ($nums[$r] == 1 || $zCount < $k)) {
            if($nums[$r] == 0){
                $zCount++;
            }
            $r++;
        }

        $result = max($result, $r - $l);

        if($nums[$l] == 0){
            $zCount--;
        }
        $l++;
    }

    return $result;
}

print_r(longestOnes1004_v2([1,1,1,0,0,0,1,1,1,1,0], 2));