<?php 

//https://www.codewars.com/kata/58235a167a8cb37e1a0000db/train/php

function numberOfPairs($gloves)
{
    $arr = [];
    
    $arrGloves = [];
    foreach($gloves as $v){
      if(isset($arr[$v])){
        $arr[$v]++;
        if($arr[$v] % 2 == 0){
          $arrGloves[] = $v;
        }
      }else{
        $arr[$v] = 1;
      }
    }
  
    return count($arrGloves);
}

echo numberOfPairs(['gray','black','purple','purple','gray','black']);