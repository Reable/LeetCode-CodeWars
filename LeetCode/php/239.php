<?php

function maxSlidingWindow239_v2($nums, $k)
{
    $q = [];
    $result = [];

    for($i = 0; $i < count($nums); $i++) {
        while(count($q) > 0 && $nums[$q[count($q) - 1]] <= $nums[$i]) {
            array_pop($q);
        }

        $q[] = $i;

        if($q[0] == $i - $k) {
            array_shift($q);
        }

        if($i >= $k - 1){
            $result[] = $nums[$q[0]];
        }
    }

    return $result;


}

// error Time limit exceeded for very big array $nums
function maxSlidingWindow231_v1($nums, $k) {
    if (empty($nums)) {
        return [];
    }

    if ($k >= count($nums)) {
        return [max($nums)];
    }

    $l = 0;
    $r = $l + $k;
    $window = array_slice($nums, $l, $r);

    $max = max($window);
    $result[] = $max;

    while($r < count($nums)) {
        unset($window[$l]);
        $window[] = $nums[$r];
        $result[] = max($window);
        $l++;
        $r++;

    }

    return $result;
}