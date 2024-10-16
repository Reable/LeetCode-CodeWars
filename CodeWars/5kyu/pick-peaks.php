<?php 

//https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/php

function pickPeaks(array $arr) {
    if (count($arr) < 2) {
        return ["pos" => [], "peaks" => []]; // Возвращаем пустые массивы
    }


    $pos = [];
    $peaks = [];
    $count = count($arr);

    for($i = 1; $i < $count - 1; $i++){
        if($arr[$i] > $arr[$i - 1] && $arr[$i] > $arr[$i + 1]){
            $pos[] = $i;
            $peaks[] = $arr[$i];
        } else if($arr[$i] > $arr[$i - 1] && $arr[$i] == $arr[$i + 1]){
            $s = $i;

            while($i < $count - 1 && $arr[$i] == $arr[$i + 1]){
                $i++;
            }

            if($i < $count - 1 && $arr[$i] > $arr[$i+1]){
                $pos[] = $s;
                $peaks[] = $arr[$s];
            }
        }
    }
    
    return [
        "pos" => $pos,
        "peaks" => $peaks
    ];
}