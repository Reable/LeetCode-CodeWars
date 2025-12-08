namespace Solutions;

public class ConcatenationOfArray {
    public int[] GetConcatenation(int[] nums) {
        int len = nums.Length;
        int[] result = new int[2 * len];

        for (int i = 0; i < len; i++)
        {
            result[i] = nums[i];
            result[i + len] = nums[i];
        }

        return result;
    }
}