namespace Solutions.Tests.Quests.Array1;

public class Array_1_Tests
{
    private readonly Array_1 _handler = new();

    [Fact]
    public void SingleElement_ReturnsElementTwice()
    {
        Assert.Equal([5, 5], _handler.GetConcatenation([5]));
    }

    [Fact]
    public void TwoElements_ReturnsCorrectConcatenation()
    {
        Assert.Equal([1, 2, 1, 2], _handler.GetConcatenation([1, 2]));
    }

    [Fact]
    public void ExampleFromLeetCode_ThreeElements()
    {
        Assert.Equal([1, 2, 1, 1, 2, 1], _handler.GetConcatenation([1, 2, 1]));
    }

    [Fact]
    public void ExampleFromLeetCode_FourElements()
    {
        Assert.Equal([1, 3, 2, 1, 1, 3, 2, 1], _handler.GetConcatenation([1, 3, 2, 1]));
    }

    [Fact]
    public void AllElementsSame()
    {
        Assert.Equal([7, 7, 7, 7, 7, 7], _handler.GetConcatenation([7, 7, 7]));
    }

    [Fact]
    public void IncreasingSequence()
    {
        Assert.Equal([1, 2, 3, 4, 1, 2, 3, 4], _handler.GetConcatenation([1, 2, 3, 4]));
    }

    [Fact]
    public void DecreasingSequence()
    {
        Assert.Equal([4, 3, 2, 1, 4, 3, 2, 1], _handler.GetConcatenation([4, 3, 2, 1]));
    }

    [Fact]
    public void ContainsMaxConstraintValues()
    {
        Assert.Equal([1000, 1, 1000, 1000, 1, 1000], _handler.GetConcatenation([1000, 1, 1000]));
    }

    [Fact]
    public void ContainsMinConstraintValues()
    {
        Assert.Equal([1, 1], _handler.GetConcatenation([1]));
    }

    [Fact]
    public void AlternatingValues()
    {
        Assert.Equal(
            [1, 2, 1, 2, 1, 2, 1, 2],
            _handler.GetConcatenation([1, 2, 1, 2])
        );
    }

    [Fact]
    public void LargeArray_LengthIsDoubled()
    {
        var nums = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.GetConcatenation(nums);

        Assert.Equal(2000, result.Length);
    }

    [Fact]
    public void LargeArray_FirstHalfEqualsOriginal()
    {
        var nums = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.GetConcatenation(nums);

        Assert.Equal(nums, result.Take(1000));
    }

    [Fact]
    public void LargeArray_SecondHalfEqualsOriginal()
    {
        var nums = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.GetConcatenation(nums);

        Assert.Equal(nums, result.Skip(1000));
    }

    [Fact]
    public void ResultIsNewArray_NotSameReference()
    {
        var nums = new[] { 1, 2, 3 };
        var result = _handler.GetConcatenation(nums);

        Assert.NotSame(nums, result);
    }

    [Fact]
    public void OriginalArrayIsNotModified()
    {
        var nums = new[] { 1, 2, 3 };
        _handler.GetConcatenation(nums);

        Assert.Equal([1, 2, 3], nums);
    }

    [Fact]
    public void WorksWithRandomValues()
    {
        Assert.Equal(
            [4, 9, 2, 4, 9, 2],
            _handler.GetConcatenation([4, 9, 2])
        );
    }

    [Fact]
    public void RepeatingPattern()
    {
        Assert.Equal(
            [3, 3, 3, 3, 3, 3],
            _handler.GetConcatenation([3, 3, 3])
        );
    }

    [Fact]
    public void MixedSmallAndLargeValues()
    {
        Assert.Equal(
            [1, 1000, 500, 1, 1000, 500],
            _handler.GetConcatenation([1, 1000, 500])
        );
    }

    [Fact]
    public void OddLengthArray()
    {
        Assert.Equal(
            [1, 2, 3, 4, 5, 1, 2, 3, 4, 5],
            _handler.GetConcatenation([1, 2, 3, 4, 5])
        );
    }

    [Fact]
    public void EvenLengthArray()
    {
        Assert.Equal(
            [6, 7, 8, 9, 6, 7, 8, 9],
            _handler.GetConcatenation([6, 7, 8, 9])
        );
    }

    [Fact]
    public void ArrayWithTwoDifferentValues()
    {
        Assert.Equal(
            [9, 1, 9, 1],
            _handler.GetConcatenation([9, 1])
        );
    }

    [Fact]
    public void ArrayWithSequentialDuplicates()
    {
        Assert.Equal(
            [2, 2, 3, 3, 2, 2, 3, 3],
            _handler.GetConcatenation([2, 2, 3, 3])
        );
    }

    [Fact]
    public void ResultContainsExactlyTwoCopiesInOrder()
    {
        var nums = new[] { 8, 6, 4 };
        var result = _handler.GetConcatenation(nums);

        Assert.Equal(nums.Concat(nums), result);
    }

    [Fact]
    public void MinimalValidInput()
    {
        Assert.Equal([1, 1], _handler.GetConcatenation([1]));
    }

    [Fact]
    public void MaxLengthInput_AllOnes()
    {
        var nums = Enumerable.Repeat(1, 1000).ToArray();
        var result = _handler.GetConcatenation(nums);

        Assert.Equal(2000, result.Length);
        Assert.All(result, x => Assert.Equal(1, x));
    }
}
