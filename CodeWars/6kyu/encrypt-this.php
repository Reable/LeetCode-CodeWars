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


// c#
// using System;
// namespace Net
// {
//     internal class Program
//     {
//         static string encode(string str)
//         {
//             string[] arr = str.Split(" ");

//             for (int i = 0; i < arr.Length; i++)
//             {
//                 int a = Convert.ToInt32(arr[i][0]);
                
//                 if(arr[i].Length > 1){
//                     char start = arr[i][1];
//                     string last = arr[i][arr[i].Length - 1].ToString();

//                     string stringOut = arr[i].Substring(1);

//                     if(stringOut.Length > 1){
//                         stringOut = stringOut.Substring(1, stringOut.Length - 1);

//                         arr[i] = a + last + stringOut + start;
//                     }else if(stringOut.Length == 1){
//                         arr[i] = a + stringOut;
//                     } else {
//                         arr[i] = a + last + stringOut + start;
//                     }
//                 } else {
//                     arr[i] = a.ToString();
//                 }
//             }

//             return string.Join(" ", arr);
//         }
//     }
// }