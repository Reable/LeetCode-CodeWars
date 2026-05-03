<?php

function mergeAlternately1768($word1, $word2) {
    $result = [];
    $len = min(strlen($word1), strlen($word2));
    $p1 = 0;
    $p2 = 0;

    while($p1 < $len || $p2 < $len) {
        if($p1 <= $p2){
            $result[] = $word1[$p1];
            $p1++;
        } else {
            $result[] = $word2[$p2];
            $p2++;
        }
    }

    while($p1 < strlen($word1)) {
        $result[] = $word1[$p1];
        $p1++;
    }
    while($p2 < strlen($word2)) {
        $result[] = $word2[$p2];
        $p2++;
    }

    return implode('', $result);
}
