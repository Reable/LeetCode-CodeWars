<?php

function reverseString344(&$s) {
    $len = count($s);

    $l = 0;
    $r = $len - 1;

    while ($l < $r) {
        $temp = $s[$l];
        $s[$l] = $s[$r];
        $s[$r] = $temp;

        $l++;
        $r--;
    }

}

$s = ["h","e","l","l","o"];
reverseString344($s);
echo "s = " . json_encode($s) . PHP_EOL;