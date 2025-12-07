using System.Collections;

namespace BinarySearch;

public class FindFirstAndLastPositionElementInSortedArray34
{
    public int[] SearchRange(int[] nums, int target)
    {
        int left = FindLeftIndex(nums, target);

        if (left < 0)
            return [-1, -1];
        
        int right = FindRightIndex(nums, target);
        
        return [left, right];
    }

    private static int FindLeftIndex(int[] nums, int target)
    {
        int l = 0;
        int r = nums.Length - 1;
        int result = -1;
        
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            
            if (nums[mid] == target)
            {
                result = mid;
                r -= 1;
            }

            if (nums[mid] < target)
                l = mid + 1;
            else if (nums[mid] > target)
                r = mid - 1;
        }
        
        return result;
    }
    
    private static int FindRightIndex(int[] nums, int target)
    {
        
        int l = 0;
        int r = nums.Length - 1;
        int result = -1;
        
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            
            if (nums[mid] == target)
            {
                result = mid;
                l += 1;
            }

            if (nums[mid] < target)
                l = mid + 1;
            else if (nums[mid] > target)
                r = mid - 1;
        }
        
        return result;
    }
}