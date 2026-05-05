<?php

function lengthOfLastWord58($s) {
    $words = explode(" ", $s);
    $len = count($words);

    for ($i = $len - 1; $i >= 0; $i--) {
        echo $words[$i] . " len = " . strlen($words[$i]) . "\n";
        if(strlen($words[$i]) == 0) continue;
        else {
            return strlen($words[$i]);
        }
    }
}

echo lengthOfLastWord58("   fly me   to   the moon  ");