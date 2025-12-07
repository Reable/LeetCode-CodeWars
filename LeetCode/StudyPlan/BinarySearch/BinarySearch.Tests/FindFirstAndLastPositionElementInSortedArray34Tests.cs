namespace BinarySearch.Tests;

public class FindFirstAndLastPositionElementInSortedArray34Tests
{
    
    private readonly FindFirstAndLastPositionElementInSortedArray34 _handler = new ();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        int[] nums = [5, 7, 7, 8, 8, 10];
        Assert.Equal([3, 4], _handler.SearchRange(nums, 8));
        Assert.Equal([-1, -1], _handler.SearchRange(nums, 6));
        Assert.Equal([-1, -1], _handler.SearchRange(nums, 0));
        Assert.Equal([1, 2], _handler.SearchRange(nums, 7));
    }
    
    [Fact]
    public void EmptyArray()
    {
        Assert.Equal([-1, -1], _handler.SearchRange([], 5));
        Assert.Equal([-1, -1], _handler.SearchRange([], -1));
        Assert.Equal([-1, -1], _handler.SearchRange([], 0));
        Assert.Equal([-1, -1], _handler.SearchRange([], 100));
    }

    [Fact]
    public void SingleElementArray()
    {
        Assert.Equal([0, 0], _handler.SearchRange([5], 5));
        Assert.Equal([-1, -1], _handler.SearchRange([5], 4));
        Assert.Equal([-1, -1], _handler.SearchRange([5], 6));
        Assert.Equal([-1, -1], _handler.SearchRange([5], -5));
    }

    [Fact]
    public void TwoElements()
    {
        Assert.Equal([0, 1], _handler.SearchRange([2, 2], 2));
        Assert.Equal([1, 1], _handler.SearchRange([1, 2], 2));
        Assert.Equal([-1, -1], _handler.SearchRange([1, 3], 2));
        Assert.Equal([0, 0], _handler.SearchRange([2, 3], 2));
    }

    [Fact]
    public void AllElementsSame()
    {
        Assert.Equal([0, 3], _handler.SearchRange([7, 7, 7, 7], 7));
        Assert.Equal([-1, -1], _handler.SearchRange([7, 7, 7, 7], 6));
        Assert.Equal([-1, -1], _handler.SearchRange([7, 7, 7, 7], 8));
        Assert.Equal([0, 4], _handler.SearchRange([3, 3, 3, 3, 3], 3));
    }

    [Fact]
    public void TargetAtBeginning()
    {
        Assert.Equal([0, 1], _handler.SearchRange([1, 1, 2, 3, 4], 1));
        Assert.Equal([0, 2], _handler.SearchRange([2, 2, 2, 3, 4], 2));
        Assert.Equal([0, 0], _handler.SearchRange([5, 6, 7], 5));
        Assert.Equal([0, 0], _handler.SearchRange([9, 10, 11], 9));
    }

    [Fact]
    public void TargetAtEnd()
    {
        Assert.Equal([3, 4], _handler.SearchRange([1, 2, 3, 4, 4], 4));
        Assert.Equal([2, 3], _handler.SearchRange([1, 1, 2, 2], 2));
        Assert.Equal([2, 2], _handler.SearchRange([0, 0, 1], 1));
        Assert.Equal([4, 4], _handler.SearchRange([1, 2, 3, 4, 5], 5));
    }

    [Fact]
    public void TargetInMiddle()
    {
        Assert.Equal([2, 3], _handler.SearchRange([1, 2, 3, 3, 4], 3));
        Assert.Equal([3, 5], _handler.SearchRange([0, 1, 2, 5, 5, 5], 5));
        Assert.Equal([1, 2], _handler.SearchRange([0, 3, 3, 4], 3));
        Assert.Equal([2, 4], _handler.SearchRange([1, 2, 7, 7, 7], 7));
    }

    [Fact]
    public void NoTarget()
    {
        Assert.Equal([-1, -1], _handler.SearchRange([1, 2, 3, 4], 5));
        Assert.Equal([-1, -1], _handler.SearchRange([2, 3, 4, 5], 1));
        Assert.Equal([-1, -1], _handler.SearchRange([0, 2, 4, 6], 3));
        Assert.Equal([-1, -1], _handler.SearchRange([-5, -4, -3], -2));
    }

    [Fact]
    public void NegativeNumbers()
    {
        Assert.Equal([1, 2], _handler.SearchRange([-5, -3, -3, 0, 1], -3));
        Assert.Equal([0, 0], _handler.SearchRange([-2, -1, 0], -2));
        Assert.Equal([3, 4], _handler.SearchRange([-10, -9, -8, -5, -5], -5));
        Assert.Equal([-1, -1], _handler.SearchRange([-4, -3, -2], 0));
    }

    [Fact]
    public void LargeNumbers()
    {
        Assert.Equal([1, 2], _handler.SearchRange([100, 200, 200, 300], 200));
        Assert.Equal([0, 0], _handler.SearchRange([999999], 999999));
        Assert.Equal([-1, -1], _handler.SearchRange([1_000_000, 2_000_000], 3_000_000));
        Assert.Equal([2, 3], _handler.SearchRange([0, 100, 500, 500, 1000], 500));
    }

    [Fact]
    public void LongArraySmallTarget()
    {
        Assert.Equal([0, 2], _handler.SearchRange([0, 0, 0, 1, 2, 3], 0));
        Assert.Equal([-1, -1], _handler.SearchRange([1, 1, 1, 1], 0));
        Assert.Equal([-1, -1], _handler.SearchRange([1, 2, 3, 4], -1));
        Assert.Equal([1, 1], _handler.SearchRange([0, 1, 2, 3, 4], 1));
    }

    [Fact]
    public void LongArrayLargeTarget()
    {
        Assert.Equal([3, 5], _handler.SearchRange([1, 2, 3, 9, 9, 9], 9));
        Assert.Equal([-1, -1], _handler.SearchRange([5, 6, 7], 8));
        Assert.Equal([2, 2], _handler.SearchRange([2, 4, 6, 8], 6));
        Assert.Equal([-1, -1], _handler.SearchRange([10, 11, 12], 99));
    }

    [Fact]
    public void LargeDuplicatesBlocks()
    {
        Assert.Equal([0, 3], _handler.SearchRange([1, 1, 1, 1, 2, 3], 1));
        Assert.Equal([4, 5], _handler.SearchRange([1, 2, 3, 4, 7, 7], 7));
        Assert.Equal([2, 4], _handler.SearchRange([0, 1, 2, 2, 2, 3], 2));
        Assert.Equal([5, 7], _handler.SearchRange([3, 4, 5, 6, 7, 8, 8, 8], 8));
    }

    [Fact]
    public void SparseDistribution()
    {
        Assert.Equal([1, 1], _handler.SearchRange([0, 2, 5, 9], 2));
        Assert.Equal([3, 3], _handler.SearchRange([1, 2, 4, 8, 16], 8));
        Assert.Equal([-1, -1], _handler.SearchRange([10, 20, 30], 25));
        Assert.Equal([0, 0], _handler.SearchRange([4, 7, 9], 4));
    }

    [Fact]
    public void OddLengthArrays()
    {
        Assert.Equal([2, 3], _handler.SearchRange([1, 1, 3, 3, 5], 3));
        Assert.Equal([0, 1], _handler.SearchRange([5, 5, 7, 8], 5));
        Assert.Equal([1, 1], _handler.SearchRange([1, 2, 3], 2));
        Assert.Equal([-1, -1], _handler.SearchRange([2, 4, 6, 8], 1));
    }

    [Fact]
    public void EvenLengthArrays()
    {
        Assert.Equal([2, 3], _handler.SearchRange([0, 1, 2, 2], 2));
        Assert.Equal([0, 0], _handler.SearchRange([3, 4, 5, 6], 3));
        Assert.Equal([-1, -1], _handler.SearchRange([3, 4, 5, 6], 7));
        Assert.Equal([1, 2], _handler.SearchRange([1, 3, 3, 9], 3));
    }

    [Fact]
    public void RandomizedSmallSets()
    {
        Assert.Equal([-1, -1], _handler.SearchRange([1, 3, 5], 2));
        Assert.Equal([1, 3], _handler.SearchRange([2, 4, 4, 4, 6], 4));
        Assert.Equal([3, 3], _handler.SearchRange([1, 2, 3, 9], 9));
        Assert.Equal([0, 1], _handler.SearchRange([7, 7, 9], 7));
    }

    [Fact]
    public void RandomizedLargeSets()
    {
        Assert.Equal([0, 0], _handler.SearchRange([1, 5, 8], 1));
        Assert.Equal([2, 4], _handler.SearchRange([3, 4, 7, 7, 7, 9], 7));
        Assert.Equal([-1, -1], _handler.SearchRange([1, 2, 3, 4, 5], 0));
        Assert.Equal([4, 5], _handler.SearchRange([0, 1, 2, 3, 9, 9], 9));
    }

    [Fact]
    public void BigMixedNegatives()
    {
        Assert.Equal([2, 3], _handler.SearchRange([-10, -8, -5, -5, 0, 2], -5));
        Assert.Equal([-1, -1], _handler.SearchRange([-9, -7, -5], -6));
        Assert.Equal([0, 1], _handler.SearchRange([-3, -3, -1], -3));
        Assert.Equal([3, 3], _handler.SearchRange([-5, -4, -3, 7], 7));
    }

    [Fact]
    public void MixedZeroCases()
    {
        Assert.Equal([1, 2], _handler.SearchRange([-1, 0, 0, 1], 0));
        Assert.Equal([0, 0], _handler.SearchRange([0, 2, 3, 4], 0));
        Assert.Equal([-1, -1], _handler.SearchRange([1, 2, 3], 0));
        Assert.Equal([2, 3], _handler.SearchRange([-2, -1, 5, 5], 5));
    }

    [Fact]
    public void ManyTargets()
    {
        Assert.Equal([1, 5], _handler.SearchRange([1, 2, 2, 2, 2, 2, 3], 2));
        Assert.Equal([0, 2], _handler.SearchRange([3, 3, 3, 5], 3));
        Assert.Equal([3, 6], _handler.SearchRange([1, 2, 3, 6, 6, 6, 6], 6));
        Assert.Equal([4, 5], _handler.SearchRange([0, 1, 2, 3, 8, 8, 9], 8));
    }

    [Fact]
    public void TargetCloseToMiddle()
    {
        Assert.Equal([2, 4], _handler.SearchRange([0, 1, 5, 5, 5, 7], 5));
        Assert.Equal([1, 1], _handler.SearchRange([2, 4, 6], 4));
        Assert.Equal([0, 0], _handler.SearchRange([3, 5, 9], 3));
        Assert.Equal([-1, -1], _handler.SearchRange([4, 6, 8], 5));
    }

    [Fact]
    public void SmallEdgeTargets()
    {
        Assert.Equal([0, 0], _handler.SearchRange([1, 2, 3], 1));
        Assert.Equal([2, 2], _handler.SearchRange([1, 2, 3], 3));
        Assert.Equal([-1, -1], _handler.SearchRange([2, 3, 4], 1));
        Assert.Equal([-1, -1], _handler.SearchRange([5, 6, 7], 4));
    }

    [Fact]
    public void ComplexPatterns()
    {
        Assert.Equal([1, 3], _handler.SearchRange([1, 4, 4, 4, 9], 4));
        Assert.Equal([0, 1], _handler.SearchRange([8, 8, 10, 12], 8));
        Assert.Equal([2, 4], _handler.SearchRange([3, 5, 7, 7, 7, 9], 7));
        Assert.Equal([-1, -1], _handler.SearchRange([5, 6, 8, 10], 7));
    }
}