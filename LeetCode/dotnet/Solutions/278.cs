namespace Solutions;

public class TaskFirstBadVersion {
    
    public int FirstBadVersion(int version, int isBadVersion) {
        int l = 0;
        int r = version;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            bool result = IsBadVersion(mid, isBadVersion);
            
            if(result)
                r = mid - 1;
            else 
                l = mid + 1;
        }

        return l;
    }

    private static bool IsBadVersion(int version, int isBadVersion)
    {
        if (version == isBadVersion)
            return true;
        else
            return false;
    }
}