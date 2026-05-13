<?php

function intersection349_v2($nums1, $nums2){
    $map = [];
    $result = [];
    foreach($nums1 as $num){
        $map[$num] = true;
    }
    foreach($nums2 as $num){
        if(isset($map[$num])) $result[$num] = $num;
    }
    return array_values($result);
}

//two pointers
function intersection349_v1($nums1, $nums2) {
    $result = [];

    sort($nums1);
    sort($nums2);

    $map = [];

    $p1 = 0;
    $p2 = 0;

    while($p1 < count($nums1) && $p2 < count($nums2)) {
        if($nums1[$p1] == $nums2[$p2]) {
            if(!isset($map[$nums1[$p1]])) {
                $map[$nums1[$p1]] = 0;
                $result[] = $nums1[$p1];
            }
            $p1++;
            $p2++;
        } else if($nums1[$p1] < $nums2[$p2]) {
            $p1++;
        } else {
            $p2++;
        }
    }

    return $result;
}

print_r(intersection349_v2([1,2,2,1], [2,2]));
//print_r(intersection349_v2([1,2,2,1], [1,1]));
//print_r(intersection349_v2([4,9,5], [9,4,9,8,4]));