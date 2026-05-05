<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
         $this->val = $val;
         $this->next = $next;
    }
}

function mergeTwoLists21($list1, $list2) {
    $temp = new ListNode(0);

    $result = $temp;

    $p1 = $list1;
    $p2 = $list2;

    while($p1 != null && $p2 != null) {
        if($p1->val < $p2->val) {
            $temp->next = new ListNode($p1->val);
            echo "p1 = " . $p1->val . "\n";
            $p1 = $p1->next;
        } else {
            $temp->next = new ListNode($p2->val);
            echo "p2 = " . $p2->val . "\n";
            $p2 = $p2->next;
        }
        $temp = $temp->next;
    }

    while($p1 != null){
        $temp->next = new ListNode($p1->val);
        $temp = $temp->next;
        $p1 = $p1->next;
    }

    while($p2 != null){
        $temp->next = new ListNode($p2->val);
        $temp = $temp->next;
        $p2 = $p2->next;
    }

    return $result->next;
}

$a1 = new ListNode(1);
$a2 = new ListNode(2);
$a3 = new ListNode(4);

$a1->next = $a2;
$a2->next = $a3;

$b1 = new ListNode(1);
$b2 = new ListNode(3);
$b3 = new ListNode(4);

$b1->next = $b2;
$b2->next = $b3;


print_r(mergeTwoLists21($a1, $b1));

