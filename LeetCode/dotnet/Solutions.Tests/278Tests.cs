namespace Solutions.Tests;

public class TaskFirstBadVersionTests
{
    private TaskFirstBadVersion _handler = new ();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        Assert.Equal(4, _handler.FirstBadVersion(version:5, isBadVersion:4));
        Assert.Equal(1, _handler.FirstBadVersion(version:1, isBadVersion:1));
    }
}
