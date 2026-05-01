<?php

function decrypt1652(array $code, int $k) {
    $len = count($code);
    $result = array_fill(0, $len, 0);

    if($k === 0) return $result;

    $steps = abs($k);
    $direction = $k > 0 ? 1 : -1;

    for($i = 0; $i < $len; $i++) {
        $sum = 0;

        for($j = 1; $j <= $steps; $j++) {
            $index = ($i + $j * $direction) % $len;

            if($index < 0) {
                $index += $len;
            }

            $sum += $code[$index];
        }

        $result[$i] = $sum;
    }

    return $result;
}
