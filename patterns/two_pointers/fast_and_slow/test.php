<?php

// Getting array [' ',' ','a',' ',' ',' ','b',' '] you need delete enter


function delEnter(&$chars)
{
    $slow = 0;
    $fast = 1;

    while($fast < count($chars))
    {
        if($chars[$slow] == " " && $chars[$fast] == " ")
        {
            array_splice($chars, $slow, 1);
            continue;
        }

        if($chars[$slow] != " " || $chars[$fast] != " "){
            $slow++;
        }

        $fast++;
    }
}

$chars = [" ", " ", " ", "h"];

var_dump(delEnter($chars));
var_dump(json_encode($chars));
