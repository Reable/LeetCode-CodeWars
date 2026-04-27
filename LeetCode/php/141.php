<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

function hasCycle141_1($head) {
    if($head == null) return false;

    $slow = $head;
    $fast = $head->next;

    while($slow !== $fast || $slow != null || $fast != null){
        if($slow != $fast){
            $slow = $slow->next;
            $fast = $fast->next->next ?? null;
        } else {
            return true;
        }
    }

    return false;
}

function hasCycle141_2($head) {
    $slow = $head;
    $fast = $head->next;

    while ($fast !== null) {
        $slow = $slow->next;
        $fast = $fast->next->next ?? null;

        if($slow === $fast) return true;
    }

    return false;
}

$head = new ListNode(1);
$head->next = $head;


echo hasCycle($head);

