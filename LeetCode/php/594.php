<?php

function findLHS594($nums) {
    sort($nums);

    $l = 0;
    $len = 0;

    for($r = 0; $r < count($nums); $r++) {
        while($nums[$r] - $nums[$l] > 1) {
            $l++;
        }

        if($nums[$r] - $nums[$l] === 1) {
            $len = max($len, $r - $l + 1);
        }
    }

    return $len;
}

echo findLHS594([1,3,2,2,5,2,3,7]);