namespace BinarySearch.Tests;

public class GuessNumberHigherOrLowerTests : GuessNumberHigherOrLower {
    
    private int _actualNumber;
    
    protected override int guess(int number)
    {
        if (number > _actualNumber)
            return -1;
        else if (number < _actualNumber)
            return 1;
        else 
            return 0;
    }

    [Fact]
    public void LeetCodeSimpleTest_Case1()
    {
        _actualNumber = 6;
        Assert.Equal(6, GuessNumber(10));
    }
    
    [Fact]
    public void LeetCodeSimpleTest_Case2()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(1));
    }
    
    [Fact]
    public void LeetCodeSimpleTest_Case3()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(2));
    }
    
    [Fact]
    public void GuessNumber_TwoElements_PickIs2()
    {
        _actualNumber = 2;
        Assert.Equal(2, GuessNumber(2));
    }

    [Fact]
    public void GuessNumber_ThreeElements_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(3));
    }

    [Fact]
    public void GuessNumber_ThreeElements_PickIs2()
    {
        _actualNumber = 2;
        Assert.Equal(2, GuessNumber(3));
    }

    [Fact]
    public void GuessNumber_ThreeElements_PickIs3()
    {
        _actualNumber = 3;
        Assert.Equal(3, GuessNumber(3));
    }

    [Fact]
    public void GuessNumber_TenElements_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(10));
    }

    [Fact]
    public void GuessNumber_TenElements_PickIs10()
    {
        _actualNumber = 10;
        Assert.Equal(10, GuessNumber(10));
    }

    [Fact]
    public void GuessNumber_TenElements_PickIs5()
    {
        _actualNumber = 5;
        Assert.Equal(5, GuessNumber(10));
    }

    [Fact]
    public void GuessNumber_TenElements_PickIs7()
    {
        _actualNumber = 7;
        Assert.Equal(7, GuessNumber(10));
    }

    [Fact]
    public void GuessNumber_HundredElements_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(100));
    }

    [Fact]
    public void GuessNumber_HundredElements_PickIs100()
    {
        _actualNumber = 100;
        Assert.Equal(100, GuessNumber(100));
    }

    [Fact]
    public void GuessNumber_HundredElements_PickIs50()
    {
        _actualNumber = 50;
        Assert.Equal(50, GuessNumber(100));
    }

    [Fact]
    public void GuessNumber_HundredElements_PickIs99()
    {
        _actualNumber = 99;
        Assert.Equal(99, GuessNumber(100));
    }

    [Fact]
    public void GuessNumber_ThousandElements_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(1000));
    }

    [Fact]
    public void GuessNumber_ThousandElements_PickIs1000()
    {
        _actualNumber = 1000;
        Assert.Equal(1000, GuessNumber(1000));
    }

    [Fact]
    public void GuessNumber_ThousandElements_PickIs500()
    {
        _actualNumber = 500;
        Assert.Equal(500, GuessNumber(1000));
    }

    [Fact]
    public void GuessNumber_ThousandElements_PickIs999()
    {
        _actualNumber = 999;
        Assert.Equal(999, GuessNumber(1000));
    }

    [Fact]
    public void GuessNumber_LargeRange_PickIsMiddle()
    {
        int n = 9999;
        _actualNumber = n / 2;
        Assert.Equal(n / 2, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_LargeRange_PickIsNearEnd()
    {
        int n = 10000;
        _actualNumber = n - 1;
        Assert.Equal(n - 1, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_MaxInt_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(int.MaxValue));
    }

    [Fact]
    public void GuessNumber_MaxInt_PickIsMax()
    {
        _actualNumber = int.MaxValue;
        Assert.Equal(int.MaxValue, GuessNumber(int.MaxValue));
    }

    [Fact]
    public void GuessNumber_OddRange_PickIsCenter()
    {
        int n = 99;
        _actualNumber = 50;
        Assert.Equal(50, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_EvenRange_PickIsCenterLeft()
    {
        int n = 100;
        _actualNumber = 50;
        Assert.Equal(50, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_EvenRange_PickIsCenterRight()
    {
        int n = 100;
        _actualNumber = 51;
        Assert.Equal(51, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_SmallRange_PickIsLowerBound()
    {
        _actualNumber = 3;
        Assert.Equal(3, GuessNumber(5));
    }

    [Fact]
    public void GuessNumber_SmallRange_PickIsUpperBound()
    {
        _actualNumber = 5;
        Assert.Equal(5, GuessNumber(5));
    }

    [Fact]
    public void GuessNumber_RepeatedCalls_DifferentPicks()
    {
        _actualNumber = 42;
        Assert.Equal(42, GuessNumber(100));

        _actualNumber = 7;
        Assert.Equal(7, GuessNumber(10));
    }

    [Fact]
    public void GuessNumber_Boundary_PickEqualsN()
    {
        int n = 42;
        _actualNumber = n;
        Assert.Equal(n, GuessNumber(n));
    }

    [Fact]
    public void GuessNumber_MinimalRange_PickIs1()
    {
        _actualNumber = 1;
        Assert.Equal(1, GuessNumber(1));
    }

    [Fact]
    public void GuessNumber_AsymmetricSearchPath_PickIs3In7()
    {
        _actualNumber = 3;
        Assert.Equal(3, GuessNumber(7));
    }
    
    [Fact]
    public void GuessNumber_RepeatedGuessesSameNumber_ConsistentResult()
    {
        _actualNumber = 42;
        Assert.Equal(42, GuessNumber(100));
    
        _actualNumber = 15;
        Assert.Equal(15, GuessNumber(50));
    
        _actualNumber = 99;
        Assert.Equal(99, GuessNumber(99));
    }
}