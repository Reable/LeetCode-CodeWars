<?php

function isAnagram242($s, $t)
{
    if(strlen($s) !== strlen($t)) return false;

    $map = [];

    for ($i = 0; $i < strlen($s); $i++) {
        if(isset($map[$s[$i]])){
            $map[$s[$i]] += 1;
        } else {
            $map[$s[$i]] = 1;
        }
    }

    for($i = 0; $i < strlen($t); $i++) {
        if(isset($map[$t[$i]])){
            $map[$t[$i]] -= 1;
        } else {
            return false;
        }
    }

    return max($map) == 0;
}

print_r(isAnagram242("anagram", "nagaram"));
