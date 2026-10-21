<?php
// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.

// Example 1:
// Input: nums = [3,1,2,4]
// Output: [2,4,3,1]
// Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
// Example 2:

// Input: nums = [0]
// Output: [0]

function sortArrayByParity905($nums) {

    $s = 0;
    $f = 0;

    while($f < count($nums))
    {
        if($nums[$f] % 2 == 0){
            $temp = $nums[$s];
            $nums[$s] = $nums[$f];
            $nums[$f] = $temp;
            $s++;
        }

        $f++;
    }

    return $nums;
}


echo json_encode(sortArrayByParity905([3,1,2,4]));
echo json_encode(sortArrayByParity905([0]));
