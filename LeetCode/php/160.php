<?php

function getIntersectionNode160($headA, $headB) {
    $pA = $headA;
    $pB = $headB;

    while ($pA !== $pB) {
        if($pA === null) {
            $pA = $headB;
        } else {
            $pA = $pA->next;
        }

        if($pB === null) {
            $pB = $headA;
        } else {
            $pB = $pB->next;
        }
    }

    return $pA;
}

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val)
    {
        $this->val = $val;
    }
}
//
//$a1 = new ListNode("4");
//$a2 = new ListNode("1");
//$b1 = new ListNode("5");
//$b2 = new ListNode("6");
//$b3 = new ListNode("1");
//$c1 = new ListNode("8");
//$c2 = new ListNode("4");
//$c3 = new ListNode("5");
//$a1->next = &$a2;
//$a2->next = &$c1;
//$b1->next = &$b2;
//$b2->next = &$b3;
//$b3->next = &$c1;
//$c1->next = &$c2;
//$c2->next = &$c3;

$a1 = new ListNode(1);
$a2 = new ListNode(2);

$b1 = new ListNode(4);
$b2 = new ListNode(6);
$b3 = new ListNode(5);

$c1 = new ListNode(2);
$c2 = new ListNode(3);
$c3 = new ListNode(4);

$a1->next = $a2;
$a2->next = $c1;

$b1->next = $b2;
$b2->next = $b3;
$b3->next = $c1;

$c1->next = $c2;
$c2->next = $c3;
print_r(getIntersectionNode160($a1, $b1));