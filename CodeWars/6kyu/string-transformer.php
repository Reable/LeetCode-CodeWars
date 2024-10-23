<?php

function string_transformer(string $s): string {
    $output = preg_replace_callback('/[a-zA-Z]/', 'toggleCase', $s);

    $arr = explode(" ", $output);

    if(count($arr) > 1){
        $arr = array_reverse($arr);
    }

    return implode(" ", $arr);
}

function toggleCase($matches) {
    return ctype_upper($matches[0]) ? strtolower($matches[0]) : strtoupper($matches[0]);
}

echo string_transformer("Hello World");