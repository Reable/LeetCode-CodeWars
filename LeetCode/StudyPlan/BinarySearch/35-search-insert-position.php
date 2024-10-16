<?php 

//https://leetcode.com/problems/search-insert-position/submissions/1402274551/?envType=study-plan-v2&envId=binary-search

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function searchInsert($nums, $target): int
    {
        $l = 0;
        $r = count($nums);

        while($l < $r){
            $m = floor(($l + $r) / 2);

            if($nums[$m] == $target) return $m;

            if($nums[$m] < $target){
                if($nums[$m + 1] > $target) return $m + 1;
                $l = $m + 1;
            } else {
                if($nums[$m] < $target) return $m;
                $r = $m;
            }
        }

        return $l;
    }
}