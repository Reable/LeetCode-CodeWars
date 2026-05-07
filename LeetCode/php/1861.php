<?php

function rotateTheBox1861($boxGrid) {
    $len = count($boxGrid);
    $lenColumns = count($boxGrid[0]);

    $result = array_fill(0, $lenColumns, []);

    for($i = 0; $i < $len; $i++) {
        $step = 0;
        for($j = $lenColumns - 1; $j >= 0; $j--) {
            if($boxGrid[$i][$j] == '.') {
                $step += 1;
                continue;
            }

            if($boxGrid[$i][$j] == '*') {
                $step = 0;
                continue;
            }

            if($boxGrid[$i][$j] == '#') {
                $index = $j + $step;
                if($index != $j){
                    $boxGrid[$i][$j + $step] = '#';
                    $boxGrid[$i][$j] = '.';
                }
            }
        }
    }
    echo PHP_EOL . json_encode($boxGrid) . PHP_EOL;

    for($i = $len - 1; $i >= 0; $i--) {
        for($j = 0; $j < count($boxGrid[$i]); $j++) {
            $result[$j][] = $boxGrid[$i][$j];
        }
    }


    print_r(json_encode($result));
}

$grid = [["#",".","*","."], ["#","#","*","."]];

//print_r(json_encode($grid) . PHP_EOL);
//rotateTheBox1861($grid);

echo PHP_EOL . "---------------" . PHP_EOL;
$grid2 = [["#",".","*","."],["#","#","*","."]];
print_r(json_encode($grid2) . PHP_EOL);
rotateTheBox1861($grid2);