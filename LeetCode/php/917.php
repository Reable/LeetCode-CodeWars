<?php
// Given a string s, reverse the string according to the following rules:
// All the characters that are not English letters remain in the same position.
// All the English letters (lowercase or uppercase) should be reversed.
// Return s after reversing it.

// Example 1:
// Input: s = "ab-cd"
// Output: "dc-ba"
// Example 2:

// Input: s = "a-bC-dEf-ghIj"
// Output: "j-Ih-gfE-dCba"
// Example 3:

// Input: s = "Test1ng-Leet=code-Q!"
// Output: "Qedo1ct-eeLg=ntse-T!"

function reverseOnlyLetters917($s) {
    $l = 0;
    $r = strlen($s) - 1;

    while($l < $r)
    {
        if(!ctype_alpha($s[$l])) {
            $l++;
        } else if(!ctype_alpha($s[$r])){
            $r--;
        } else {
            $temp = $s[$l];
            $s[$l] = $s[$r];
            $s[$r] = $temp;
            $l++;
            $r--;
        }
    }

    return $s;
}

echo reverseOnlyLetters917("a-bC-dEf-ghIj");
