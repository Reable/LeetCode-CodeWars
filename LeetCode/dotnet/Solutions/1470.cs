namespace Solutions;

public class ShuffleTheArray {
    public int[] Shuffle(int[] nums, int n) {
        int id = 0;
        List<int> result = new List<int>();
        for (int i = 0; i < n; i++)
        {
            result.Add(nums[id]);
            result.Add(nums[n + id]);
            id += 1;
        }
        return result.ToArray();
    }
}