namespace Solutions.Tests.Quests.Array1;

public class Array_3_Tests
{
    private readonly Array_3 _handler = new();

    [Fact]
    public void SingleOne_ReturnsOne()
    {
        Assert.Equal(1, _handler.FindMaxConsecutiveOnes([1]));
    }

    [Fact]
    public void SingleZero_ReturnsZero()
    {
        Assert.Equal(0, _handler.FindMaxConsecutiveOnes([0]));
    }

    [Fact]
    public void AllOnes_ReturnsArrayLength()
    {
        Assert.Equal(5, _handler.FindMaxConsecutiveOnes([1, 1, 1, 1, 1]));
    }

    [Fact]
    public void AllZeros_ReturnsZero()
    {
        Assert.Equal(0, _handler.FindMaxConsecutiveOnes([0, 0, 0, 0]));
    }

    [Fact]
    public void OnesAtBeginning()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([1, 1, 1, 0, 0]));
    }

    [Fact]
    public void OnesAtEnd()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([0, 0, 1, 1, 1]));
    }

    [Fact]
    public void OnesInMiddle()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([0, 1, 1, 1, 0]));
    }

    [Fact]
    public void AlternatingOnesAndZeros_ReturnsOne()
    {
        Assert.Equal(1, _handler.FindMaxConsecutiveOnes([1, 0, 1, 0, 1, 0]));
    }

    [Fact]
    public void TwoSeparateSequences_ReturnsMax()
    {
        Assert.Equal(4, _handler.FindMaxConsecutiveOnes([1, 1, 0, 1, 1, 1, 1]));
    }

    [Fact]
    public void MultipleSequencesSameLength()
    {
        Assert.Equal(2, _handler.FindMaxConsecutiveOnes([1, 1, 0, 1, 1, 0]));
    }

    [Fact]
    public void LongSequenceAfterShortOne()
    {
        Assert.Equal(5, _handler.FindMaxConsecutiveOnes([1, 0, 1, 1, 1, 1, 1]));
    }

    [Fact]
    public void LongSequenceBeforeShortOne()
    {
        Assert.Equal(5, _handler.FindMaxConsecutiveOnes([1, 1, 1, 1, 1, 0, 1]));
    }

    [Fact]
    public void ZeroSplitsLongestSequence()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([1, 1, 1, 0, 1, 1]));
    }

    [Fact]
    public void MultipleZerosBetweenOnes()
    {
        Assert.Equal(2, _handler.FindMaxConsecutiveOnes([1, 1, 0, 0, 0, 1, 1]));
    }

    [Fact]
    public void LeadingZerosIgnored()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([0, 0, 1, 1, 1]));
    }

    [Fact]
    public void TrailingZerosIgnored()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([1, 1, 1, 0, 0]));
    }

    [Fact]
    public void ZerosBetweenEveryOne()
    {
        Assert.Equal(1, _handler.FindMaxConsecutiveOnes([1, 0, 1, 0, 1, 0, 1]));
    }

    [Fact]
    public void OneLongSequenceAmongManyShort()
    {
        Assert.Equal(6, _handler.FindMaxConsecutiveOnes([1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1]));
    }

    [Fact]
    public void SingleOneAmongZeros()
    {
        Assert.Equal(1, _handler.FindMaxConsecutiveOnes([0, 0, 1, 0, 0]));
    }

    [Fact]
    public void SequenceEndsWithOne_NoTrailingZero()
    {
        Assert.Equal(4, _handler.FindMaxConsecutiveOnes([0, 1, 1, 1, 1]));
    }

    [Fact]
    public void SequenceStartsWithOne_NoLeadingZero()
    {
        Assert.Equal(4, _handler.FindMaxConsecutiveOnes([1, 1, 1, 1, 0]));
    }

    [Fact]
    public void LargeInput_AllOnes()
    {
        var nums = Enumerable.Repeat(1, 100000).ToArray();
        Assert.Equal(100000, _handler.FindMaxConsecutiveOnes(nums));
    }

    [Fact]
    public void LargeInput_AllZeros()
    {
        var nums = Enumerable.Repeat(0, 100000).ToArray();
        Assert.Equal(0, _handler.FindMaxConsecutiveOnes(nums));
    }

    [Fact]
    public void LargeInput_SingleZeroInMiddle()
    {
        var nums = Enumerable.Repeat(1, 50000)
            .Concat([0])
            .Concat(Enumerable.Repeat(1, 49999))
            .ToArray();

        Assert.Equal(50000, _handler.FindMaxConsecutiveOnes(nums));
    }

    [Fact]
    public void ClassicLeetCodeExample()
    {
        Assert.Equal(3, _handler.FindMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]));
    }
}
