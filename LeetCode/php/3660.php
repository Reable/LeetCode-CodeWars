<?php

function maxValue3660($nums) {
    $n = count($nums);
    $ans = array_fill(0, $n, 0);
    $stack = []; // [val, left, right]

    for($i = 0; $i < $n; $i++) {
        $val = $nums[$i];
        $left = $i;
        $right = $i;

        while(count($stack) > 0 and $stack[count($stack) - 1][0] > $nums[$i]){
            $elem = array_pop($stack);
            $val = max($val, $elem[0]);
            $left = $elem[1];
        }
        $stack[] = [$val, $left, $right];
    }

    for ($i = 0; $i < count($stack); $i++) {
        for ($j = $stack[$i][1]; $j <= $stack[$i][2]; $j++) {
            $ans[$j] = $stack[$i][0];
        }
    }

    return $ans;
}

print_r(maxValue3660([2,1,3]));