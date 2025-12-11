namespace Solutions;

public class CountCharactersInString
{
    public static Dictionary<char, int> Count(string str)
    {
        var counts = new Dictionary<char, int>();

        foreach (var c in str)
        {
            if (!counts.TryAdd(c, 1))
                counts[c]++;
        }
    
        return counts;
    }
}