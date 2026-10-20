<?php
// You are given a 1-indexed array of integers numbers that is already sorted in non-decreasing order.
// Find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers index1 and index2 as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

// Example 1:
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
// Example 2:

// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
// Example 3:

// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].


function twoSum($numbers, $target) {
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

echo json_encode(twoSum([2,7,11,15], 9));
