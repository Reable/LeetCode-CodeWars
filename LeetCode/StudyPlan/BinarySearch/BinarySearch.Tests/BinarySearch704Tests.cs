namespace BinarySearch.Tests;

public class BinarySearch704Tests
{
    [Fact]
    public void Search_LargeArray_ExistsAtStartMiddleEnd()
    {
        int[] nums = Enumerable.Range(0, 100_000).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(0, bs.Search(nums, 0));
        Assert.Equal(50_000, bs.Search(nums, 50_000));
        Assert.Equal(99_999, bs.Search(nums, 99_999));
        Assert.Equal(-1, bs.Search(nums, 100_000));
    }

    [Fact]
    public void Search_LargeArray_DoesNotExist()
    {
        int[] nums = Enumerable.Range(0, 100_000).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(-1, bs.Search(nums, 100_000));
        Assert.Equal(-1, bs.Search(nums, -1));
    }

    [Fact]
    public void Search_LargeNegativeArray()
    {
        int[] nums = Enumerable.Range(-100_000, 100_000).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(0, bs.Search(nums, -100_000));
        Assert.Equal(99_999, bs.Search(nums, -1));
        Assert.Equal(-1, bs.Search(nums, 0));
    }

    [Fact]
    public void Search_ArrayWithMinAndMaxInt()
    {
        int[] nums = new[] { int.MinValue, 0, int.MaxValue };
        var bs = new BinarySearch704();
        Assert.Equal(0, bs.Search(nums, int.MinValue));
        Assert.Equal(1, bs.Search(nums, 0));
        Assert.Equal(2, bs.Search(nums, int.MaxValue));
        Assert.Equal(-1, bs.Search(nums, int.MaxValue - 1));
    }

    [Fact]
    public void Search_ArrayWithRepeatingBlocks()
    {
        int[] nums = Enumerable.Repeat(new[] { 1, 2, 3, 4, 5 }, 10_000).SelectMany(x => x).ToArray();
        var bs = new BinarySearch704();
        int r = bs.Search(nums, 3);
        Assert.True(r % 5 == 2);
        Assert.Equal(-1, bs.Search(nums, 6));
    }

    [Fact]
    public void Search_ArrayWithSkippedElements()
    {
        int[] nums = Enumerable.Range(0, 20_000).Where(x => x % 2 == 0).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(5000, bs.Search(nums, 10_000));
        Assert.Equal(-1, bs.Search(nums, 10001));
    }

    [Fact]
    public void Search_MixedNegativeZeroPositive()
    {
        int[] nums = Enumerable.Range(-5000, 10001).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(5000, bs.Search(nums, 0));
        Assert.Equal(0, bs.Search(nums, -5000));
        Assert.Equal(10000, bs.Search(nums, 5000));
    }

    [Fact]
    public void Search_HugeSequentialArray_TargetExistsEvery1000()
    {
        int[] nums = Enumerable.Range(0, 1_000_000).Select(x => x * 1000).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(123, bs.Search(nums, 123_000));
        Assert.Equal(-1, bs.Search(nums, 123_456));
    }

    //[Fact]
    //public void Search_ArrayWithAlternatingSignNumbers()
    //{
    //    int[] nums = Enumerable.Range(0, 10_000).Select(x => x % 2 == 0 ? x : -x).ToArray();
    //    var bs = new BinarySearch704();
    //    Assert.Equal(1234, bs.Search(nums, 1234));
    //    Assert.Equal(1235, bs.Search(nums, -1235));
    //    Assert.Equal(-1, bs.Search(nums, 1233));
    //}

    [Fact]
    public void Search_VerySparseArray()
    {
        int[] nums = Enumerable.Range(0, 100_000).Select(x => x * 10_000).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(5, bs.Search(nums, 50_000));
        Assert.Equal(-1, bs.Search(nums, 55_000));
    }

    [Fact]
    public void Search_EmptyArray_ReturnsMinusOne()
    {
        int[] nums = Array.Empty<int>();
        var bs = new BinarySearch704();
        Assert.Equal(-1, bs.Search(nums, 0));
    }

    [Fact]
    public void Search_SingleElement()
    {
        var bs = new BinarySearch704();

        int[] nums1 = { 42 };
        Assert.Equal(0, bs.Search(nums1, 42));
        Assert.Equal(-1, bs.Search(nums1, 1));

        int[] nums2 = { 1, 5 };
        Assert.Equal(0, bs.Search(nums2, 1));
        Assert.Equal(1, bs.Search(nums2, 5));
        Assert.Equal(-1, bs.Search(nums2, 3));
    }

    [Fact]
    public void Search_ThreeToFiveElements()
    {
        var bs = new BinarySearch704();

        int[] nums3 = { -1, 0, 1 };
        Assert.Equal(1, bs.Search(nums3, 0));

        int[] nums5 = { -2, -1, 0, 1, 2 };
        for (int i = -2; i <= 2; i++)
        {
            Assert.Equal(i + 2, bs.Search(nums5, i));
        }
        Assert.Equal(-1, bs.Search(nums5, 3));
    }

    [Fact]
    public void Search_MediumArray_EveryElement()
    {
        int[] nums = Enumerable.Range(0, 50).ToArray();
        var bs = new BinarySearch704();
        for (int i = 0; i < 50; i++)
        {
            Assert.Equal(i, bs.Search(nums, i));
        }
        Assert.Equal(-1, bs.Search(nums, 100));
    }

    [Fact]
    public void Search_ArrayWithDuplicates()
    {
        int[] nums = { 1, 2, 2, 2, 3 };
        var bs = new BinarySearch704();
        int r = bs.Search(nums, 2);
        Assert.True(r >= 1 && r <= 3);
        Assert.Equal(-1, bs.Search(nums, 4));
    }

    [Fact]
    public void Search_AllSameElements()
    {
        int[] nums = Enumerable.Repeat(5, 100).ToArray();
        var bs = new BinarySearch704();
        Assert.InRange(bs.Search(nums, 5), 0, 99);
        Assert.Equal(-1, bs.Search(nums, 6));
    }

    [Fact]
    public void Search_NegativeAndPositiveNumbers()
    {
        int[] nums = Enumerable.Range(-50, 101).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(50, bs.Search(nums, 0));
        Assert.Equal(0, bs.Search(nums, -50));
        Assert.Equal(100, bs.Search(nums, 50));
        Assert.Equal(-1, bs.Search(nums, 51));
    }

    [Fact]
    public void Search_LargeRandomArray()
    {
        int[] nums = Enumerable.Range(0, 100_000).OrderBy(x => x).ToArray();
        var bs = new BinarySearch704();
        Assert.Equal(12345, bs.Search(nums, 12345));
        Assert.Equal(-1, bs.Search(nums, 100_001));
    }

    [Fact]
    public void Search_CheckAllIndexesSmallArray()
    {
        int[] nums = Enumerable.Range(0, 100).ToArray();
        var bs = new BinarySearch704();
        for (int i = 0; i < nums.Length; i++)
        {
            Assert.Equal(i, bs.Search(nums, i));
        }
    }
}
