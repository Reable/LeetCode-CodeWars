namespace Solutions;

public class Array_3
{
    public int FindMaxConsecutiveOnes(int[] nums) {
        int max = 0;
        int current = 0;

        for(int i = 0; i < nums.Length; i++){
            if(nums[i] == 1){
                current += 1;
            } else {
                if (current > max)
                {
                    max = current;
                }
                current = 0;
                
            }
        }

        if (current > max)
            return current;
        return max;
    }
}