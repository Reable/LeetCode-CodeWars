<?php

function strStr28($haystack, $needle): int {
    if (strlen($needle) < 1) {
        return -1;
    }

    for ($i = 0; $i <= strlen($haystack) - strlen($needle); $i++) {
        $match = true;
        for($j = 0; $j < strlen($needle); $j++) {
            if ($haystack[$i+$j] !== $needle[$j]) {
                $match = false;
                break;
            }
        }

        if($match){
            return $i;
        }

    }

    return -1;
}

print_r(strStr28("sadbutsad", "sad") . "\n");
print_r(strStr28("leetcode", "leeto") . "\n");
print_r(strStr28("a", "a") . "\n");


