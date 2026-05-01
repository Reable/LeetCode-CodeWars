<?php

function longestNiceSubstring1763($s) {
    if(strlen($s) < 2) return "";

    return checkNiceString($s, 0, strlen($s));
}

function checkNiceString($s, $l, $r) {
    if($r - $l < 2) return null;

    $temp = [];
    for($i = $l; $i < $r; $i++) {
        $temp[$s[$i]] = true;
    }

    for($i = $l; $i < $r; $i++) {
        $chr = $s[$i];
        $caseS = ctype_upper($chr) ? strtolower($chr) : strtoupper($chr);

        if(!isset($temp[$caseS])){
            $leftS = checkNiceString($s, $l, $i);
            $rightS = checkNiceString($s, $i + 1, $r);

            if(is_null($leftS)) return $rightS;
            if(is_null($rightS)) return $leftS;

            return strlen($leftS) >= strlen($rightS)
                ? $leftS
                : $rightS;
        }
    }

    return substr($s, $l, $r - $l);
}

$s = "jcJ";
echo json_encode(
    checkNiceString("jcJ", 0, strlen($s), 1)
). PHP_EOL;
