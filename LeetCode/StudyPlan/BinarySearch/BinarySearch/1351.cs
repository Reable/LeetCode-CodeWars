namespace BinarySearch;

public class CountNegativeNumbersInSortedMatrix
{
    public int CountNegatives(int[][] grid)
    {
        int m = grid.Length;
        int n = grid[0].Length;
        int i = 0;
        int j = n - 1;

        int count = 0;

        while (i < m && j >= 0)
        {
            if (grid[i][j] < 0)
            {
                count += m - i;
                j--;
            }
            else
                i++;
        }
        
        return count;
    }
}