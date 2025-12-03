namespace BinarySearch;

public class BinarySearch704 {
    public int Search(int[] nums, int target)
    {
        int result = -1;
        
        int l = 0;
        int r = nums.Length - 1;
        while (l <= r) {
            int mid =  l + (r - l) / 2;
            if (nums[mid] == target) {
                result = mid;
                r = mid - 1;
            } 
            if (nums[mid] < target) l = mid + 1;
            else r = mid - 1;
        }
        return result;
    }

}