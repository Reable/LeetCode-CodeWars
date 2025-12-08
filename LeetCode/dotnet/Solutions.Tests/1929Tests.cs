namespace Solutions.Tests;

public class ConcatenationOfArrayTests
{
    
    private ConcatenationOfArray _handler = new ();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        Assert.Equal([1,2,1,1,2,1], _handler.GetConcatenation([1,2,1]));
        Assert.Equal([1,3,2,1,1,3,2,1], _handler.GetConcatenation([1,3,2,1]));
    }
    
    [Fact]
    public void EmptyArray_ReturnsEmptyArray()
    {
        var result = _handler.GetConcatenation(new int[0]);
        Assert.Equal(new int[0], result);
    }

    [Fact]
    public void SingleElement_DuplicatesElement()
    {
        var result = _handler.GetConcatenation(new[] { 42 });
        Assert.Equal(new[] { 42, 42 }, result);
    }

    [Fact]
    public void TwoElements_CorrectConcatenation()
    {
        var result = _handler.GetConcatenation(new[] { 1, 2 });
        Assert.Equal(new[] { 1, 2, 1, 2 }, result);
    }

    [Fact]
    public void ThreeElements_CorrectConcatenation()
    {
        var result = _handler.GetConcatenation(new[] { 1, 2, 3 });
        Assert.Equal(new[] { 1, 2, 3, 1, 2, 3 }, result);
    }

    [Fact]
    public void FourElements_CorrectConcatenation()
    {
        var result = _handler.GetConcatenation(new[] { 10, 20, 30, 40 });
        Assert.Equal(new[] { 10, 20, 30, 40, 10, 20, 30, 40 }, result);
    }

    [Fact]
    public void ArrayWithZeros_PreservesZeros()
    {
        var result = _handler.GetConcatenation(new[] { 0, 0, 1 });
        Assert.Equal(new[] { 0, 0, 1, 0, 0, 1 }, result);
    }

    [Fact]
    public void AllZeros_ReturnsDoubleZeros()
    {
        var result = _handler.GetConcatenation(new[] { 0, 0, 0 });
        Assert.Equal(new[] { 0, 0, 0, 0, 0, 0 }, result);
    }

    [Fact]
    public void NegativeNumbers_HandlesNegativesCorrectly()
    {
        var result = _handler.GetConcatenation(new[] { -1, -2, -3 });
        Assert.Equal(new[] { -1, -2, -3, -1, -2, -3 }, result);
    }

    [Fact]
    public void MixedPositiveAndNegative_CorrectConcatenation()
    {
        var result = _handler.GetConcatenation(new[] { -1, 2, -3, 4 });
        Assert.Equal(new[] { -1, 2, -3, 4, -1, 2, -3, 4 }, result);
    }

    [Fact]
    public void LargeNumbers_NoOverflowIssues()
    {
        var result = _handler.GetConcatenation(new[] { int.MaxValue, int.MinValue });
        Assert.Equal(new[] { int.MaxValue, int.MinValue, int.MaxValue, int.MinValue }, result);
    }

    [Fact]
    public void IdenticalElements_DuplicatesCorrectly()
    {
        var result = _handler.GetConcatenation(new[] { 5, 5, 5 });
        Assert.Equal(new[] { 5, 5, 5, 5, 5, 5 }, result);
    }

    [Fact]
    public void AlternatingSigns_PreservesOrder()
    {
        var result = _handler.GetConcatenation(new[] { 1, -1, 2, -2 });
        Assert.Equal(new[] { 1, -1, 2, -2, 1, -1, 2, -2 }, result);
    }

    [Fact]
    public void OneElementArray_LengthCheck()
    {
        var result = _handler.GetConcatenation(new[] { 99 });
        Assert.True(result.Length == 2);
        Assert.Equal(99, result[0]);
        Assert.Equal(99, result[1]);
    }

    [Fact]
    public void FiveElements_LengthCheck()
    {
        var input = new[] { 1, 2, 3, 4, 5 };
        var result = _handler.GetConcatenation(input);
        Assert.True(result.Length == 10);
        for (int i = 0; i < input.Length; i++)
        {
            Assert.Equal(input[i], result[i]);
            Assert.Equal(input[i], result[i + input.Length]);
        }
    }

    [Fact]
    public void LongArray_PerformanceAndCorrectness()
    {
        var input = Enumerable.Range(1, 1000).ToArray();
        var result = _handler.GetConcatenation(input);

        Assert.True(result.Length == 2000);
        for (int i = 0; i < input.Length; i++)
        {
            Assert.Equal(input[i], result[i]);
            Assert.Equal(input[i], result[i + input.Length]);
        }
    }

    [Fact]
    public void MaxArraySize_EdgeCase()
    {
        // Тестируем близкий к максимальному размер (на практике может вызвать OutOfMemoryException)
        // В реальных тестах лучше использовать меньший размер
        const int size = 10_000;
        var input = new int[size];
        for (int i = 0; i < size; i++) input[i] = i % 10;

        var result = _handler.GetConcatenation(input);
        Assert.True(result.Length == size * 2);
    }

    [Fact]
    public void SingleNegativeElement_CorrectDuplication()
    {
        var result = _handler.GetConcatenation(new[] { -42 });
        Assert.Equal(new[] { -42, -42 }, result);
    }
}
