<?php

function minimumDifference1984($nums, $k) {
    if(count($nums) <= 1) return 0;

    sort($nums);

    $sum = PHP_INT_MAX;
    $len = count($nums);

    for($i = 0; $i <= $len - $k; $i++) {
        $currentDiff = $nums[$i + $k - 1] - $nums[$i];
        $sum = min($sum, $currentDiff);
    }

    return $sum;
}
