/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target) {
    int i, j;
    int* res = (int*)malloc(2*sizeof(int));
    if (res == NULL)
        return NULL;
    for (i=0; i<numsSize; i++) {
        int delta = target - nums[i];
        for (j=i+1; j<numsSize; j++) {
            if (delta == nums[j]) {
                res[0] = i + 1;
                res[1] = j + 1;
                return res;
            }
        }
    }
    return NULL;
}
