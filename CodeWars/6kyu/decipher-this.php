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


