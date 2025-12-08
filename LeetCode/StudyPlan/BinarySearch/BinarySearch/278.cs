namespace Solutions;

public class TaskFirstBadVersion {
    
    public int FirstBadVersion(int version) {
        int l = 0;
        int r = version;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            bool result = IsBadVersion(mid);
            
            if(result)
                r = mid - 1;
            else 
                l = mid + 1;
        }

        return l;
    }

    protected virtual bool IsBadVersion(int version)
    {
        throw new NotImplementedException(
            "Метод IsBadVersion должен быть переопределён в тестовой среде");
    }
}