<?php

function canPlaceFlowers605($flowerbed, $n) {
    $p = 0;
    $count = 0;
    $len = count($flowerbed);

    while($p < $len){
        if($flowerbed[$p] == 0){
            $left = $p == 0 || $flowerbed[$p-1] == 0;
            $right = $p == $len - 1 || $flowerbed[$p+1] == 0;

            if($left && $right){
                $flowerbed[$p] = 1;
                $p += 2;
            } else {
                $p++;
            }
        } else {
            $p += 2;
        }
    }

    return $count >= $n;
}
