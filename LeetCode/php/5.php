<?php

// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=two-pointers
// Given a string s, return the longest palindromic substring in s.

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"


function longestPalindrome5($s) {
    $maxLen = 1;
    $start = 0;

    for($i = 0; $i < strlen($s); $i++)
    {
        $oddLen  = findCentralPolindrom($s, $i, $i);
        $evenLen = findCentralPolindrom($s, $i, $i+1);

        if ($oddLen > $maxLen) {
            $maxLen = $oddLen;
            $start = $i - intdiv($oddLen, 2);
        }

        if ($evenLen > $maxLen) {
            $maxLen = $evenLen;
            $start = $i - intdiv($evenLen, 2) + 1;
        }
    }

    return substr($s, $start, $maxLen);

}

function findCentralPolindrom($s, $l, $r)
{
    while($l >=0 && $r < strlen($s) && $s[$l] == $s[$r])
    {
        $l--;
        $r++;
    }

    return $r - $l - 1;
}

echo longestPalindrome5("abbd");
