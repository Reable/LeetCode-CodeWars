<?php

function reverseVowels345_v2($s){
    $arrSymbols = [
        'a'=>true, 'e'=>true, 'i'=>true,'o'=>true,'u'=>true,
        'A' => true,'E' => true,'I' => true,'O' => true,'U' => true
    ];

    $l = 0;
    $r = strlen($s) - 1;

    while($l < $r){
        $left = $s[$l];
        $right = $s[$r];

        if(!isset($arrSymbols[$left])){
            $l++;
            continue;
        }
        if(!isset($arrSymbols[$right])){
            $r--;
            continue;
        }

        $s[$l] = $right;
        $s[$r] = $left;

        $l++;
        $r--;
    }
    return $s;
}

function reverseVowels345_v1($s) {

    $symbolsArr = [
        'a' => 0,
        'e'=>0,
        'i'=>0,
        'o'=>0,
        'u'=>0
    ];

    $result = [];
    for($i = 0; $i < strlen($s); $i++) {
        $result[$i] = $s[$i];
    }

    $l = 0;
    $r = strlen($s) - 1;
    while ($l < $r) {
        $leftChr = strtolower($result[$l]);
        $rightChr = strtolower($result[$r]);

        if(isset($symbolsArr[$leftChr]) && (isset($symbolsArr[$rightChr]))
        ) {
            $temp = $result[$r];
            $result[$r] = $result[$l];
            $result[$l] = $temp;
            $l++;
            $r--;

        } else {
            if(isset($symbolsArr[$leftChr]) && !isset($symbolsArr[$rightChr])) {
                $r--;
            } else if(!isset($symbolsArr[$leftChr]) && isset($symbolsArr[$rightChr])) {
                $l++;
            } else {
                $r--;
                $l++;
            }
        }
    }
    return implode("",$result);
}
