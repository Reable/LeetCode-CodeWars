<?php

function rotateString796($s, $goal) {
    if(strlen($s) !== strlen($goal)) return false;
    $str = str_repeat($s,2);
    return str_contains($str, $goal);
}

