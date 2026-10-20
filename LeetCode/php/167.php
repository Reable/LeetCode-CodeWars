<?php

function twoSum167($numbers, $target) {
    $l = 0;
    $r = count($numbers) - 1;

    while($l < $r){
        $tempSum = $numbers[$l] + $numbers[$r];

        if($tempSum == $target){
            return [$l + 1, $r + 1];
        } else if ($tempSum > $target){
            $r--;
        } else {
            $l++;
        }
    }

    return [-1, -1];
}

echo json_encode(twoSum167([2,7,11,15], 9));
