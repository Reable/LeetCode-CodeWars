namespace BinarySearch;

public class FindSmallestLetterGreaterThanTarget744
{
    
    public char NextGreatestLetter(char[] letters, char target) {
        int l = 0;
        int r = letters.Length;
        while (l < r)
        {
            int mid = ( l + r ) / 2;
            
            if(letters[mid] > target) 
                r = mid;
            else 
                l = mid + 1;
        }

        if (l < 0 || l >= letters.Length)
            return letters[0];
        
        return letters[l];
    }

}