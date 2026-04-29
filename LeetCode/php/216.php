<?php

function containsNearbyDuplicate216($nums, $k) {
    $window = [];

    for($i = 0; $i < count($nums); $i++) {
        if(isset($window[$nums[$i]])) {
            return true;
        }

        $window[$nums[$i]] = true;

        if(count($window) > $k) {
            unset($window[$nums[$i - $k]]);
        }
    }

    return false;
}

echo containsNearbyDuplicate216([1,2,3,1], 3) . " --- true \n";
echo containsNearbyDuplicate216([1,0,1,1], 1) . " --- true \n";
echo containsNearbyDuplicate216([1,2,3,1,2,3], 2) . " --- false \n";
