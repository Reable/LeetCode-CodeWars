<?php 

//https://leetcode.com/problems/binary-search/description/?envType=study-plan-v2&envId=binary-search

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function search($nums, $target) {
        $l = 0;
        $r = count($nums);

        while($l < $r){
            $m = floor(($l + $r) / 2);
            
            if($nums[$m] === $target) return $m;
            
            if($nums[$m] < $target){
                $l = $m + 1;
            } else {
                $r = $m;
            }
        }

        return -1;
    }
}