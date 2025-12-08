namespace BinarySearch.Tests;

public class BinarySearch704Tests
{
    private readonly BinarySearch _handler = new ();
    
    [Fact]
    public void Search_LargeArray_ExistsAtStartMiddleEnd()
    {
        int[] nums = Enumerable.Range(0, 100_000).ToArray();
        Assert.Equal(0, _handler.Search(nums, 0));
        Assert.Equal(50_000, _handler.Search(nums, 50_000));
        Assert.Equal(99_999, _handler.Search(nums, 99_999));
        Assert.Equal(-1, _handler.Search(nums, 100_000));
    }

    [Fact]
    public void Search_LargeArray_DoesNotExist()
    {
        int[] nums = Enumerable.Range(0, 100_000).ToArray();
        Assert.Equal(-1, _handler.Search(nums, 100_000));
        Assert.Equal(-1, _handler.Search(nums, -1));
    }

    [Fact]
    public void Search_LargeNegativeArray()
    {
        int[] nums = Enumerable.Range(-100_000, 100_000).ToArray();
        Assert.Equal(0, _handler.Search(nums, -100_000));
        Assert.Equal(99_999, _handler.Search(nums, -1));
        Assert.Equal(-1, _handler.Search(nums, 0));
    }

    [Fact]
    public void Search_ArrayWithMinAndMaxInt()
    {
        int[] nums = new[] { int.MinValue, 0, int.MaxValue };
        Assert.Equal(0, _handler.Search(nums, int.MinValue));
        Assert.Equal(1, _handler.Search(nums, 0));
        Assert.Equal(2, _handler.Search(nums, int.MaxValue));
        Assert.Equal(-1, _handler.Search(nums, int.MaxValue - 1));
    }

    [Fact]
    public void Search_ArrayWithRepeatingBlocks()
    {
        int[] nums = Enumerable.Repeat(new[] { 1, 2, 3, 4, 5 }, 10_000).SelectMany(x => x).ToArray();
        int r = _handler.Search(nums, 3);
        Assert.True(r % 5 == 2);
        Assert.Equal(-1, _handler.Search(nums, 6));
    }

    [Fact]
    public void Search_ArrayWithSkippedElements()
    {
        int[] nums = Enumerable.Range(0, 20_000).Where(x => x % 2 == 0).ToArray();
        Assert.Equal(5000, _handler.Search(nums, 10_000));
        Assert.Equal(-1, _handler.Search(nums, 10001));
    }

    [Fact]
    public void Search_MixedNegativeZeroPositive()
    {
        int[] nums = Enumerable.Range(-5000, 10001).ToArray();
        Assert.Equal(5000, _handler.Search(nums, 0));
        Assert.Equal(0, _handler.Search(nums, -5000));
        Assert.Equal(10000, _handler.Search(nums, 5000));
    }

    [Fact]
    public void Search_HugeSequentialArray_TargetExistsEvery1000()
    {
        int[] nums = Enumerable.Range(0, 1_000_000).Select(x => x * 1000).ToArray();
        Assert.Equal(123, _handler.Search(nums, 123_000));
        Assert.Equal(-1, _handler.Search(nums, 123_456));
    }

    [Fact]
    public void Search_VerySparseArray()
    {
        int[] nums = Enumerable.Range(0, 100_000).Select(x => x * 10_000).ToArray();
        Assert.Equal(5, _handler.Search(nums, 50_000));
        Assert.Equal(-1, _handler.Search(nums, 55_000));
    }

    [Fact]
    public void Search_EmptyArray_ReturnsMinusOne()
    {
        int[] nums = Array.Empty<int>();
        Assert.Equal(-1, _handler.Search(nums, 0));
    }

    [Fact]
    public void Search_SingleElement()
    {
        int[] nums1 = { 42 };
        Assert.Equal(0, _handler.Search(nums1, 42));
        Assert.Equal(-1, _handler.Search(nums1, 1));

        int[] nums2 = { 1, 5 };
        Assert.Equal(0, _handler.Search(nums2, 1));
        Assert.Equal(1, _handler.Search(nums2, 5));
        Assert.Equal(-1, _handler.Search(nums2, 3));
    }

    [Fact]
    public void Search_ThreeToFiveElements()
    {
        int[] nums3 = { -1, 0, 1 };
        Assert.Equal(1, _handler.Search(nums3, 0));

        int[] nums5 = { -2, -1, 0, 1, 2 };
        for (int i = -2; i <= 2; i++)
        {
            Assert.Equal(i + 2, _handler.Search(nums5, i));
        }
        Assert.Equal(-1, _handler.Search(nums5, 3));
    }

    [Fact]
    public void Search_MediumArray_EveryElement()
    {
        int[] nums = Enumerable.Range(0, 50).ToArray();
        for (int i = 0; i < 50; i++)
        {
            Assert.Equal(i, _handler.Search(nums, i));
        }
        Assert.Equal(-1, _handler.Search(nums, 100));
    }

    [Fact]
    public void Search_ArrayWithDuplicates()
    {
        int[] nums = { 1, 2, 2, 2, 3 };
        int r = _handler.Search(nums, 2);
        Assert.True(r >= 1 && r <= 3);
        Assert.Equal(-1, _handler.Search(nums, 4));
    }

    [Fact]
    public void Search_AllSameElements()
    {
        int[] nums = Enumerable.Repeat(5, 100).ToArray();
        Assert.InRange(_handler.Search(nums, 5), 0, 99);
        Assert.Equal(-1, _handler.Search(nums, 6));
    }

    [Fact]
    public void Search_NegativeAndPositiveNumbers()
    {
        int[] nums = Enumerable.Range(-50, 101).ToArray();
        Assert.Equal(50, _handler.Search(nums, 0));
        Assert.Equal(0, _handler.Search(nums, -50));
        Assert.Equal(100, _handler.Search(nums, 50));
        Assert.Equal(-1, _handler.Search(nums, 51));
    }

    [Fact]
    public void Search_LargeRandomArray()
    {
        int[] nums = Enumerable.Range(0, 100_000).OrderBy(x => x).ToArray();
        Assert.Equal(12345, _handler.Search(nums, 12345));
        Assert.Equal(-1, _handler.Search(nums, 100_001));
    }

    [Fact]
    public void Search_CheckAllIndexesSmallArray()
    {
        int[] nums = Enumerable.Range(0, 100).ToArray();
        for (int i = 0; i < nums.Length; i++)
        {
            Assert.Equal(i, _handler.Search(nums, i));
        }
    }
}
