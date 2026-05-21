<?php

function isGood2784($nums) {
    sort($nums);

    $n = max($nums);

    if (count($nums) != $n + 1) {
        return false;
    }

    $n = count($nums);
    if($n < 2){
        return false;
    }

    for($i = 1; $i < $n - 1; $i++){
        if($nums[$i] - $nums[$i - 1] != 1){
            return false;
        }
    }

    if($nums[$n - 2] == $nums[$n - 1]){
        return true;
    }

    return false;
}

print_r(isGood2784([6,7,8,9,9]));