<?php
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]

function moveZeroes(&$nums) {
    $slow = 0;
    $fast = 0;

    while($fast < count($nums))
    {
        if($nums[$fast] != 0)
        {
            $temp = $nums[$slow];
            $nums[$slow] = $nums[$fast];
            $nums[$fast] = $temp;
            $slow++;
        }

        $fast++;
    }
}

$nums = [1,7,0,0,0,3,12,2,2];
var_dump(moveZeroes($nums));
var_dump(json_encode($nums));
