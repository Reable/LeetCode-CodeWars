<?php

// O(n * k)
function groupAnagrams49(array $strs) {
    $map = [];
    $result = [];

    $i = 0;
    while($i < count($strs) ) {
        $bitMap = getBitMap($strs[$i]);

        if(isset($map[$bitMap])) {
            $map[$bitMap][] = $strs[$i];
        } else {
            $map[$bitMap] = [$strs[$i]];
        }

        $i++;
    }

    // or use return array_values($map)
    foreach($map as $bit => $values) {
        $result[] = $values;
    }

    return $result;
}

function getBitMap($string){
    $bits = array_fill(0,26,0);
    $length = strlen($string);

    $charA = ord('a');
    for ($i = 0; $i < $length; $i++) {
        $bits[ord($string[$i]) - $charA]++;
    }
    return implode("#", $bits);
}

//groupAnagrams49(["eat","tea","tan","ate","nat","bat"]);
