<?php 

// https://www.codewars.com/kata/5848565e273af816fb000449/train/php

function encryptThis($text): string
{
    $string = [];

    foreach(explode(" ", $text) as $word){
        $first = ord($word[0]);
        $word = substr($word, 1);

        if(strlen($word) > 0){
            $last = $word[strlen($word) - 1];
            $word[strlen($word) - 1] = $word[0];
            $word[0] = $last;
        }
        $word = $first . $word;
        $string[] = $word;
    }


    return implode(" ", $string);
}


echo encryptThis("А") . PHP_EOL;
echo encryptThis("A wise old owl lived in an oak") . PHP_EOL;
echo encryptThis("The more he saw the less he spoke") . PHP_EOL;
echo encryptThis("The less he spoke the more he heard") . PHP_EOL;
