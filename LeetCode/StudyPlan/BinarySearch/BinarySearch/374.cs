namespace BinarySearch;

public class GuessNumberHigherOrLower {
    public int GuessNumber(int n)
    {
        int l = 1;
        int r = n;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            int res = guess(mid);

            if (res == 0) {
                return mid;
            }

            if (res == -1)
                r = mid - 1;
            else
               l = mid + 1;
        }

        return -1;
    }
    
    protected virtual int guess(int number)
    {
        throw new NotImplementedException(
            "Метод guess должен быть переопределён в тестовой среде");
    }

}