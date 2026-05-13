<?php

function findMaxConsecutiveOnes485_v2($nums){
    $counter = 0;
    $max = 0;

    foreach($nums as $num){
        if($num == 1){
            $counter++;
        } else {
            if($counter > $max){
                $max = $counter;
            }
            $counter = 0;
        }
    }
    if($counter > $max){
        $max = $counter;
    }
    return $max;

}

function findMaxConsecutiveOnes485_v1($nums) {

    $l = 0;
    $r = 0;
    $result = 0;

    while($r < count($nums)) {
        $max = -1;

        while($r < count($nums) && $nums[$r] == 1) {
            $max = $r - $l + 1;
            $r++;
        }

        $result = max($result, $max);
        $r++;
        $l = $r;
    }

    return $result;
}

