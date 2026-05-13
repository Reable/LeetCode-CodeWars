<?php

function intersect350($nums1, $nums2) {
    $p1 = 0;
    $p2 = 0;

    sort($nums1);
    sort($nums2);

    $result = [];
    while($p2 < count($nums2) && $p1 < count($nums1)){
        if($nums1[$p1] == $nums2[$p2]){
            $result[] = $nums1[$p1];
            $p1++;
            $p2++;
        } else if($nums1[$p1] > $nums2[$p2]){
            $p2++;
        } else {
            $p1++;
        }
    }

    return $result;
}

print_r(intersect350([4,9,5], [9,4,9,8,4])); // [4,5,9]   [4,4,8,9,9]
print_r(intersect350([1,2,2,1], [2,2]));     // [1,1,2,2] [2,2]
