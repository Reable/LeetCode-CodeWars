<?php

function kidsWithCandies($candies, $extraCandies) {
    $maxCandie = max($candies);
    $result = array_fill(0, $maxCandie, false);
    for($i = 0; $i < count($candies); $i++) {
        $sum = $candies[$i] + $extraCandies;
        $result[$i] = $sum >= $maxCandie;
    }
    return $result;
}