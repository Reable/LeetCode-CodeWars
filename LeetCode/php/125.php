<?php

// Two pointer
function isPalindrome125(string $s) {
    $str = strtolower($s);
    $str = preg_replace('/[^a-zA-Z0-9]/', '', $str);

    if (strlen($str) == 0) return 1;

    $l = 0;
    $r = strlen($str) - 1;

    while ($l < $r) {
        if ($str[$l] != $str[$r]) {
            return 0;
        }
        $l++;
        $r--;
    }

    return 1;
}

function isPalindrome125_2(string $s) {
    $str = strtolower($s);
    $str = preg_replace('/[^a-zA-Z0-9]/', '', $str);
    echo $str;
    return strrev($str) == $str;
}

echo isPalindrome125_2("Marge, let's \"[went].\" I await {news} telegram.");
