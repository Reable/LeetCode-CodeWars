<?php

function moveZeroes283(&$nums) {
    $n = count($nums);

    $slow = 0;
    $fast = 0;

    while($fast < $n){
        if($nums[$fast] != 0){
            $t = $nums[$slow];
            $nums[$slow] = $nums[$fast];
            $nums[$fast] = $t;

            $slow++;
            $fast++;
        } else {
            $fast++;
        }
    }

}

$nums = [1,3,4,0,5,9,7];
moveZeroes283($nums);
print_r($nums);