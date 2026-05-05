<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
      $this->val = $val;
      $this->next = $next;
    }
}
function rotateRight($head, $k) {
    if($head == null || $head->next === null || $k == 0) return $head;

    $arrayLinks = array();
    while($head != null) {
        $temp = $head;
        $head = $head->next;
        $temp->next = null;
        $arrayLinks[] = $temp;
    }

    $len = count($arrayLinks);

    $index = $len - $k;
    if($k > $len) $index = $len - ($k % $len);
    if($k == $len) $index = 0;

    $tempList = new ListNode();
    $result = $tempList;

    for($i = $index; $i < $len; $i++) {
        $tempList->next = $arrayLinks[$i];
        $tempList = $tempList->next;
    }

    for($i = 0; $i < $index; $i++) {
        $tempList->next = $arrayLinks[$i];
        $tempList = $tempList->next;
    }

    return $result->next;
}


$a1 = new ListNode(1);
$a2 = new ListNode(2);
$a3 = new ListNode(3);
$a4 = new ListNode(4);
$a5 = new ListNode(5);

$a1->next = $a2;
$a2->next = $a3;
$a3->next = $a4;
$a4->next = $a5;

print_r(rotateRight($a1, 5));

$a1 = new ListNode(1);
$a2 = new ListNode(2);

$a1->next = $a2;

print_r(rotateRight($a1, 3));