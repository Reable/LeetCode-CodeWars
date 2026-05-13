<?php

function reverseWords151($s){
    $s = trim($s);
    $temp = "";
    $result = [];

    for($i = strlen($s) - 1; $i >= 0; $i--){
        if($s[$i] == " "){
            if(strlen($temp) > 0){
                $result[] = strrev($temp);
                $temp = "";
            }
            continue;
        }

        $temp .= $s[$i];
    }
    if(strlen($temp) > 0){
        $result[] = strrev($temp);
    }

    return implode(" ", $result);
}

echo "result = " . reverseWords151("a good   example");