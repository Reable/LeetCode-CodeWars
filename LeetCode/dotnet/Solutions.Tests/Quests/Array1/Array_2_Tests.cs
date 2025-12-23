namespace Solutions.Tests.Quests.Array1;

public class Array_2_Tests
{
    private readonly Array_2 _handler = new();

    [Fact]
    public void Example1_FromLeetCode()
    {
        Assert.Equal(
            [2, 3, 5, 4, 1, 7],
            _handler.Shuffle([2, 5, 1, 3, 4, 7], 3)
        );
    }

    [Fact]
    public void Example2_FromLeetCode()
    {
        Assert.Equal(
            [1, 4, 2, 3, 3, 2, 4, 1],
            _handler.Shuffle([1, 2, 3, 4, 4, 3, 2, 1], 4)
        );
    }

    [Fact]
    public void Example3_FromLeetCode()
    {
        Assert.Equal(
            [1, 2, 1, 2],
            _handler.Shuffle([1, 1, 2, 2], 2)
        );
    }

    [Fact]
    public void MinimalValidInput_nEqualsOne()
    {
        Assert.Equal(
            [5, 9],
            _handler.Shuffle([5, 9], 1)
        );
    }

    [Fact]
    public void TwoPairs_ShuffledCorrectly()
    {
        Assert.Equal(
            [1, 3, 2, 4],
            _handler.Shuffle([1, 2, 3, 4], 2)
        );
    }

    [Fact]
    public void AllElementsSame()
    {
        Assert.Equal(
            [1, 1, 1, 1],
            _handler.Shuffle([1, 1, 1, 1], 2)
        );
    }

    [Fact]
    public void IncreasingXAndYSequences()
    {
        Assert.Equal(
            [1, 4, 2, 5, 3, 6],
            _handler.Shuffle([1, 2, 3, 4, 5, 6], 3)
        );
    }

    [Fact]
    public void DecreasingXAndYSequences()
    {
        Assert.Equal(
            [3, 6, 2, 5, 1, 4],
            _handler.Shuffle([3, 2, 1, 6, 5, 4], 3)
        );
    }

    [Fact]
    public void AlternatingValuesInResult()
    {
        Assert.Equal(
            [10, 40, 20, 50, 30, 60],
            _handler.Shuffle([10, 20, 30, 40, 50, 60], 3)
        );
    }

    [Fact]
    public void RepeatingPattern()
    {
        Assert.Equal(
            [2, 3, 2, 3, 2, 3],
            _handler.Shuffle([2, 2, 2, 3, 3, 3], 3)
        );
    }

    [Fact]
    public void WorksWithMaxConstraintValues()
    {
        Assert.Equal(
            [1000, 1000, 1000, 1000],
            _handler.Shuffle([1000, 1000, 1000, 1000], 2)
        );
    }

    [Fact]
    public void XPartContainsSameValue_YPartDifferent()
    {
        Assert.Equal(
            [1, 2, 1, 3, 1, 4],
            _handler.Shuffle([1, 1, 1, 2, 3, 4], 3)
        );
    }

    [Fact]
    public void YPartContainsSameValue_XPartDifferent()
    {
        Assert.Equal(
            [1, 5, 2, 5, 3, 5],
            _handler.Shuffle([1, 2, 3, 5, 5, 5], 3)
        );
    }

    [Fact]
    public void ResultLengthEqualsInputLength()
    {
        int[] nums = [1, 2, 3, 4, 5, 6];
        var result = _handler.Shuffle(nums, 3);

        Assert.Equal(nums.Length, result.Length);
    }

    [Fact]
    public void ResultIsNewArray_NotSameReference()
    {
        var nums = new[] { 1, 2, 3, 4 };
        var result = _handler.Shuffle(nums, 2);

        Assert.NotSame(nums, result);
    }

    [Fact]
    public void OriginalArrayIsNotModified()
    {
        var nums = new[] { 1, 2, 3, 4 };
        _handler.Shuffle(nums, 2);

        Assert.Equal([1, 2, 3, 4], nums);
    }

    [Fact]
    public void LargeInput_nEqualsFiveHundred()
    {
        var x = Enumerable.Range(1, 500);
        var y = Enumerable.Range(501, 500);
        var nums = x.Concat(y).ToArray();

        var result = _handler.Shuffle(nums, 500);

        Assert.Equal(1000, result.Length);
    }

    [Fact]
    public void LargeInput_FirstElementIsX1()
    {
        var nums = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.Shuffle(nums, 500);

        Assert.Equal(1, result[0]);
    }

    [Fact]
    public void LargeInput_SecondElementIsY1()
    {
        var nums = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.Shuffle(nums, 500);

        Assert.Equal(501, result[1]);
    }

    [Fact]
    public void InterleavingIsCorrectForMiddleElements()
    {
        int[] nums = [1, 2, 3, 4, 5, 6];
        var result = _handler.Shuffle(nums, 3);

        Assert.Equal(3, result[4]); // x3
        Assert.Equal(6, result[5]); // y3
    }

    [Fact]
    public void HandlesRandomValues()
    {
        Assert.Equal(
            [7, 4, 9, 1, 3, 8],
            _handler.Shuffle([7, 9, 3, 4, 1, 8], 3)
        );
    }

    [Fact]
    public void SequentialDuplicates()
    {
        Assert.Equal(
            [2, 3, 2, 3, 2, 3, 2, 3],
            _handler.Shuffle([2, 2, 2, 2, 3, 3, 3, 3], 4)
        );
    }

    [Fact]
    public void nEqualsArrayHalfLength()
    {
        int[] nums = [1, 2, 3, 4];
        var result = _handler.Shuffle(nums, nums.Length / 2);

        Assert.Equal([1, 3, 2, 4], result);
    }

    [Fact]
    public void ShuffledOrderIsExactlyXThenYRepeated()
    {
        var nums = new[] { 8, 6, 4, 1, 3, 5 };
        var result = _handler.Shuffle(nums, 3);

        Assert.Equal([8, 1, 6, 3, 4, 5], result);
    }

    [Fact]
    public void MinimalEdgeCaseWithMaxValues()
    {
        Assert.Equal(
            [1000, 1000],
            _handler.Shuffle([1000, 1000], 1)
        );
    }
}
