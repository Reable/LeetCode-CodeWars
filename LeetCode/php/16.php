<?php

function threeSumClosest16($nums, $target) {
    if(count($nums) < 3) return -1;

    sort($nums);

    $tempSum = $nums[0] + $nums[1] + $nums[2];
    $result = abs($tempSum - $target);

    print_r($nums);

    $n = count($nums);

    for($i = 0; $i < $n - 2; $i++) {
        if ($i > 0 && $nums[$i] == $nums[$i - 1]) {
            continue;
        }

        $l = $i + 1;
        $r = $n - 1;


        while($l < $r) {
            $currentSum = $nums[$i] + $nums[$l] + $nums[$r];
            $currentMin = abs($currentSum - $target);

            if($currentMin < $result) {
                $result = $currentMin;
                $tempSum = $currentSum;
            }

            if($currentSum < $target) {
                $l++;
            } else if($currentSum > $target) {
                $r--;
            } else {
                return $currentSum;
            }
        }
    }


    return $tempSum;
}

print_r(threeSumClosest16([0,3,97,102,200], 300));