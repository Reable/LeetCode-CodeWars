<?php

//Binary search O(log n)
function search33_v2($nums, $target) { // Beats 100%
    $l = 0;
    $r = count($nums) - 1;

    while ($l <= $r) {
        $mid = (int)(($r + $l) / 2);

        if($nums[$mid] == $target) return $mid;

        if($nums[$l] <= $nums[$mid]) {
            if($nums[$l] <= $target && $target < $nums[$mid]) {
                $r = $mid - 1;
            } else {
                $l = $mid + 1;
            }
        } else {
            if($nums[$mid] < $target && $target <= $nums[$r]) {
                $l = $mid + 1;
            } else {
                $r = $mid - 1;
            }
        }
    }

    return -1;
}


//Linear search O(n)
function search33_v1($nums, $target) { // Beats 100%
    $res = -1;
    for($i = 0; $i < count($nums); $i++){
        if($nums[$i] == $target) $res = $i;
    }
    return $res;
}
