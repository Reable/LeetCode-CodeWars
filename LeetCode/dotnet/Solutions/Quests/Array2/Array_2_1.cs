namespace Solutions.Quests.Array2;

public class Array_2_1
{
    public int[] FindErrorNums(int[] nums)
    {
        int current = 0;
        // [1,2,2,4]
        for (int i = 0; i < nums.Length; i++) 
        {
            if (nums[i] > current)
            {
                current = nums[i];
                continue;
            }

            if (nums[i] < current)
                continue;

            return [current, ++current];
        }

        return [];
    }

    public void Main()
    {
        int[] nums = { 1, 2, 2, 4 };

        var data = FindErrorNums(nums);

        foreach (int i in data) {
            Console.WriteLine(i);
        }
    }
}
