<?php

function minimumDifference($nums, $k) {
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

echo minimumDifference([87063,61094,44530,21297,95857,93551,9918], 6);