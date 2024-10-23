<?php 

// https://www.codewars.com/kata/decipher-this

function decipherThis($str): string
{
    $string = [];

    foreach(explode(" ", $str) as $word){
        $chrCode = (int) $word;
        $firstSymbol = chr($chrCode);
        $word = str_split(str_replace($chrCode, $firstSymbol, $word));
        if(count($word) > 1){
            $end = $word[count($word) - 1];
            $word[count($word) - 1] = $word[1];
            $word[1] = $end;
        }
        $string[] = implode("", $word);
    }

    return implode(" ", $string);
}

echo decipherThis("72olle 103doo 100ya") . PHP_EOL;                             // Hello good day
echo decipherThis("82yade 115te 103o") . PHP_EOL;                               // Ready set go
echo decipherThis("65 119esi 111dl 111lw 108dvei 105n 97n 111ka") . PHP_EOL;    //A wise old owl lived in an oak


// c#
// using System;

// namespace Net
// {
//     internal class Program
//     {
//         static string decode(string str)
//         {
//             string[] ints = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"];
//             string[] arr = str.Split(" ");

//             for(int i = 0; i < arr.Length; i++)
//             {
//                 char[] charsArr = arr[i].ToCharArray();
//                 int idStart = 0;
//                 string first = "";

//                 for(int c = 0; c < 3; c++)
//                 {
//                     string asciiValue = charsArr[c].ToString();
//                     if (ints.Contains(asciiValue)) {
//                         idStart++;
//                         first += asciiValue;
//                     }
//                 }

//                 if(int.TryParse(first, out int number)){
//                     string stringNotFirst = arr[i].Substring(idStart);
//                     string r = "";
//                     string l = "";

//                     Console.WriteLine(stringNotFirst + " " + idStart);

//                     if(stringNotFirst.Length >= 2){
//                         r = stringNotFirst[stringNotFirst.Length - 1].ToString();
//                         l = stringNotFirst[0].ToString();
//                     }

//                     arr[i] = (char) number + r + stringNotFirst.Substring(1, stringNotFirst.Length - 2) + l;
                
//                 }
//             }

//             return String.Join(" ", arr);
//         }
//     }
// }