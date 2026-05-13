<?php

function findAnagrams438($s, $p) {

    $result = [];
    $bitMap = array_fill(0, 26, 0);
    for ($i = 0; $i < strlen($p); $i++) {
        $bit = ord($p[$i]) - ord("a");
        $bitMap[$bit]++;
    }

    $windowBitMap = array_fill(0, 26, 0);
    for ($i = 0; $i < strlen($p); $i++) {
        $bit = ord($s[$i]) - ord("a");
        $windowBitMap[$bit]++;
    }

    if($windowBitMap == $bitMap) {
        $result[] = 0;
    }

    for ($i = 1; $i <= strlen($s) - strlen($p); $i++) {
        $lBit = ord($s[$i - 1]) - ord("a");
        $windowBitMap[$lBit] -= 1;

        $rBit = ord($s[$i + strlen($p) - 1]) - ord("a");
        $windowBitMap[$rBit] += 1;

        if($windowBitMap == $bitMap) {
            $result[] = $i;
        }

    }

    return $result;
}

findAnagrams438("cbaebabacd", "abc");