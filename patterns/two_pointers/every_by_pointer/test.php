<?php
// Geting 2 sorted array, you need return all elements who inspect in arrayes

function findCommonElements(array $arr1, array $arr2): array
{
    $i = 0;
    $j = 0;

    $countArr1 = count($arr1);
    $countArr2 = count($arr2);

    $commonElements = [];

    while($i < $countArr1 && $j < $countArr2)
    {
        if($arr1[$i] === $arr2[$j])
        {
            $commonElements[] = $arr1[$i];
            $i++;
            $j++;
            continue;
        }

        if($arr1[$i] > $arr2[$j])
        {
            $j++;
        }
        else
        {
            $i++;
        }
    }

    return $commonElements;
}

var_dump(findCommonElements([0,2,4,8,8], [1,2,2,7,8,8,8]));
