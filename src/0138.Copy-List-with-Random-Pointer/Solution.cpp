/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node() {}

    Node(int _val, Node* _next, Node* _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};
*/
class Solution {
public:
    Node* copyRandomList(Node* head) {

    }
    RandomListNode *copyRandomList(RandomListNode *head) {
        RandomListNode *newHead, *l1, *l2;
        if (head == NULL) return NULL;
        for (l1 = head; l1 != NULL; l1 = l1->next->next) {
            l2 = new RandomListNode(l1->label);
            l2->next = l1->next;
            l1->next = l2;
        }

        newHead = head->next;
        for (l1 = head; l1 != NULL; l1 = l1->next->next) {
            if (l1->random != NULL) l1->next->random = l1->random->next;
        }

        for (l1 = head; l1 != NULL; l1 = l1->next) {
            l2 = l1->next;
            l1->next = l2->next;
            if (l2->next != NULL) l2->next = l2->next->next;
        }

        return newHead;
    }
     RandomListNode *copyRandomList_2(RandomListNode *head) {
    	 RandomListNode *newHead, *l1, *l2;
    	 if (head == NULL) return NULL;
    	 /*origin
    	 l1			1 - 2 - 3   NULL
    	 rand p  		  \ |
    	 l1			1 - 2 - 3   NULL
    	 */

    	 /*connect
    	 l2			1   2   3   NULL
    	 next p  		  \ |
    	 l1			1 - 2 - 3   NULL
    	 rand p  		|   |   |
    	 l2			1   2   3   NULL
    	 */
    	 for (l1 = head; l1 != NULL; l1 = l1->next) {
    		 l2 = new RandomListNode(l1->label);
    		 l2->next = l1->random;
    		 l1->random = l2;
    	 }

    	 /* build rand p
    	 l2			1   2   3   NULL
    	 next p  		  \ |
    	 l1			1 - 2 - 3   NULL
    	 rand p  		|   |   |
    	 l2			1   2   3   NULL
    	 rand p  		  \ |
    	 l2			1   2   3   NULL
    	 */
    	 newHead = head->random;
    	 for (l1 = head; l1 != NULL; l1 = l1->next) {
    		 l2 = l1->random;
    		 l2->random = l2->next ? l2->next->random : NULL;
    	 }
    	 /* restore	 */
    	 for (l1 = head; l1 != NULL; l1 = l1->next) {
    		 l2 = l1->random;
    		 l1->random = l2->next;
    		 l2->next = l1->next ? l1->next->random : NULL;
    	 }

    	 return newHead;
     }
};