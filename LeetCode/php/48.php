<?php

function rotate48(&$matrix) {
    $len = count($matrix);

    for($i = 0; $i < $len; $i++) {
        for($j = $i+1; $j < $len; $j++) {
            $a = $matrix[$i][$j];
            $matrix[$i][$j] = $matrix[$j][$i];
            $matrix[$j][$i] = $a;
        }
    }
    for($i = 0; $i < $len; $i++) {
        $matrix[$i] = array_reverse($matrix[$i]);
    }
}

$m = [[1,2,3]
    , [4,5,6]
    , [7,8,9]
];
rotate48($m);
print_r($m);