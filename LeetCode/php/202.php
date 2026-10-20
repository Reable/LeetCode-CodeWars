<?php

function isHappy202($n) {
    $sum = $n;
    $hash = [$sum => true];

    do{
        $temp = 0;
        while($sum != 0){
            $temp += ($sum % 10) ** 2;
            $sum = intdiv($sum, 10);
        }

        $hash[$temp] = isset($hash[$temp]) ? false : true;

        if($temp == 1) return "true";
        else $sum = $temp;

    } while($hash[$sum]);

    return "false";
}

echo isHappy202(19);
echo isHappy202(2);
