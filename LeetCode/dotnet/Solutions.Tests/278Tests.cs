namespace Solutions.Tests;

public class TaskFirstBadVersionTests : TaskFirstBadVersion
{
    private int _actualBadVersion;
    
    protected override bool IsBadVersion(int version)
    {
        return version >= _actualBadVersion;
    }
    
    [Fact]
    public void LeetCodeSimpleTest_Case1()
    {
        _actualBadVersion = 4;
        Assert.Equal(4, FirstBadVersion(5));
    }
    
    [Fact]
    public void LeetCodeSimpleTest_Case2()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(1));
    }

    [Fact]
    public void FirstVersionIsBad_SingleVersion()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(1));
    }

    [Fact]
    public void LastVersionIsBad_SingleVersion()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(1));
    }

    [Fact]
    public void FirstVersionIsBad_MultipleVersions()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(5));
    }

    [Fact]
    public void LastVersionIsBad_MultipleVersions()
    {
        _actualBadVersion = 5;
        Assert.Equal(5, FirstBadVersion(5));
    }

    [Fact]
    public void MiddleVersionIsBad()
    {
        _actualBadVersion = 3;
        Assert.Equal(3, FirstBadVersion(5));
    }

    [Fact]
    public void BadVersionIsSecondOfTwo()
    {
        _actualBadVersion = 2;
        Assert.Equal(2, FirstBadVersion(2));
    }

    [Fact]
    public void BadVersionIsFirstOfTwo()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(2));
    }

    [Fact]
    public void AllVersionsAreBad()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(10));
    }

    [Fact]
    public void OnlyLastVersionIsBad_LargeSet()
    {
        _actualBadVersion = 100;
        Assert.Equal(100, FirstBadVersion(100));
    }

    [Fact]
    public void BadVersionNearStart_LargeSet()
    {
        _actualBadVersion = 5;
        Assert.Equal(5, FirstBadVersion(1000));
    }

    [Fact]
    public void BadVersionNearEnd_LargeSet()
    {
        _actualBadVersion = 995;
        Assert.Equal(995, FirstBadVersion(1000));
    }

    [Fact]
    public void BadVersionInExactMiddle_LargeSet()
    {
        _actualBadVersion = 500;
        Assert.Equal(500, FirstBadVersion(1000));
    }

    [Fact]
    public void TwoVersions_BothBad()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(2));
    }

    [Fact]
    public void ThreeVersions_MiddleIsBad()
    {
        _actualBadVersion = 2;
        Assert.Equal(2, FirstBadVersion(3));
    }

    [Fact]
    public void FourVersions_ThirdIsBad()
    {
        _actualBadVersion = 3;
        Assert.Equal(3, FirstBadVersion(4));
    }

    [Fact]
    public void Version1000_BadIs1()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(1000));
    }

    [Fact]
    public void Version1000_BadIs1000()
    {
        _actualBadVersion = 1000;
        Assert.Equal(1000, FirstBadVersion(1000));
    }

    [Fact]
    public void Version10000_BadIs5000()
    {
        _actualBadVersion = 5000;
        Assert.Equal(5000, FirstBadVersion(10000));
    }

    [Fact]
    public void Version100000_BadIs1()
    {
        _actualBadVersion = 1;
        Assert.Equal(1, FirstBadVersion(100000));
    }

    [Fact]
    public void Version100000_BadIs99999()
    {
        _actualBadVersion = 99999;
        Assert.Equal(99999, FirstBadVersion(100000));
    }
}
