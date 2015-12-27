/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode* head, **cursor;       
    int carry = 0, sum = 0, x = 0, y = 0;

    cursor = &head;
    while (l1 != NULL || l2 != NULL) {
        if (l1 == NULL) {
            x = 0;
        } else {
            x = l1->val;
            l1 = l1->next;
        }

        if (l2 == NULL) {
            y = 0;
        } else {
            y = l2->val;
            l2 = l2->next;
        }

        sum = x + y + carry;
        struct ListNode* node = (struct ListNode*)malloc(sizeof(struct ListNode));
        if (node == NULL) {
            goto err;
        }
        node->val = sum%10;
        node->next = NULL;
        *cursor = node;
        cursor = &(node->next);
        carry = (int)(sum/10);
    }
    
    if (carry > 0) {
        struct ListNode* node = (struct ListNode*)malloc(sizeof(struct ListNode));
        if (node == NULL) {
            goto err;
        }
        node->val = carry;
        node->next = NULL;
        *cursor = node;
    }
    return head;
err:
    while (head != NULL) {
        struct ListNode* tmp = head;
        head = head->next;
        free(tmp);
    }
    return NULL;
}
