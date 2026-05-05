<?php

function maxRotateFunction396($nums) {
    $len = count($nums);

    $totalSum = array_sum($nums);
    $sum = 0;
    for ($i = 0; $i < $len; $i++) {
        $sum += $i * $nums[$i];
    }
    $result = $sum;

    for ($i = $len - 1; $i > 0; $i--) {
        $sum = $sum + $totalSum - $len * $nums[$i];
        $result = max($result, $sum);
    }

    return $result;
}

echo maxRotateFunction396([4,3,2,6]);