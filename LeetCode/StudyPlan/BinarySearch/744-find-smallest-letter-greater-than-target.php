<?php 

//https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/?envType=study-plan-v2&envId=binary-search

class Solution {

    /**
     * @param String[] $letters
     * @param String $target
     * @return String
     */
    function nextGreatestLetter($letters, $target) {
        foreach($letters as $letter){
            if(ord($letter) > ord($target)) return $letter;
        }
        return $letters[0];
    }
}