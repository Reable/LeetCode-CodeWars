<?php

function detectCycle($head) {
    $slow = $head;
    $fast= $head;

    while($fast !== null && $fast->next !== null){
        $slow = $slow->next;
        $fast = $fast->next->next;

        if ($slow === $fast){
            $slow = $head;

            while ($slow !== $fast){
                $slow = $slow->next;
                $fast = $fast->next;
            }

            return $slow;
        }
    }

    return null;
}

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

$head = new ListNode(1);
echo detectCycle($head);
