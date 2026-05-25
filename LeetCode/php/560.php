<?php

function subarraySum($nums, $k) {
    $result = 0;
    $map = [0 => 1];
    $prefixSum = 0;

    foreach ($nums as $num) {
        $prefixSum += $num;
        $need = $prefixSum - $k;

        if (isset($map[$need])) {
            $result += $map[$need];
        }

        if (isset($map[$prefixSum])) {
            $map[$prefixSum]++;
        } else {
            $map[$prefixSum] = 1;
        }
    }

    return $result;
}

echo "2 -> " . json_encode(subarraySum([1,2,3], 3)) . PHP_EOL;
echo "3 -> " . json_encode(subarraySum([1,1,1], 1)) . PHP_EOL;
echo "0 -> " . json_encode(subarraySum([1], 0)) . PHP_EOL;
echo "1 -> " . json_encode(subarraySum([-1,-1,0], 0)) . PHP_EOL;
echo "3 -> " . json_encode(subarraySum([-1,1,0], 0)) . PHP_EOL;
