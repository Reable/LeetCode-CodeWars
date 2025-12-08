namespace Solutions.Tests;

public class ShuffleTheArrayTests
{
    private ShuffleTheArray _handler = new ();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        Assert.Equal([2,3,5,4,1,7], _handler.Shuffle([2,5,1,3,4,7], 3));
        Assert.Equal([1,4,2,3,3,2,4,1], _handler.Shuffle([1,2,3,4,4,3,2,1], 4));
        Assert.Equal([1,2,1,2], _handler.Shuffle([1,1,2,2], 2));
    }

    [Fact]
    public void BasicShuffle_n3()
    {
        int[] nums = { 2, 5, 1, 3, 4, 7 };
        int n = 3;
        int[] expected = { 2, 3, 5, 4, 1, 7 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void BasicShuffle_n4()
    {
        int[] nums = { 1, 2, 3, 4, 4, 3, 2, 1 };
        int n = 4;
        int[] expected = { 1, 4, 2, 3, 3, 2, 4, 1 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void Shuffle_TwoElements_n1()
    {
        int[] nums = { 10, 20 };
        int n = 1;
        int[] expected = { 10, 20 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void Shuffle_FourElements_n2()
    {
        int[] nums = { 1, 1, 2, 2 };
        int n = 2;
        int[] expected = { 1, 2, 1, 2 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void AllElementsSame_n2()
    {
        int[] nums = { 5, 5, 5, 5 };
        int n = 2;
        int[] expected = { 5, 5, 5, 5 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void AlternatingValues_n3()
    {
        int[] nums = { 1, 2, 3, 10, 20, 30 };
        int n = 3;
        int[] expected = { 1, 10, 2, 20, 3, 30 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void FirstHalfZero_n2()
    {
        int[] nums = { 0, 0, 1, 2 };
        int n = 2;
        int[] expected = { 0, 1, 0, 2 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void SecondHalfZero_n2()
    {
        int[] nums = { 1, 2, 0, 0 };
        int n = 2;
        int[] expected = { 1, 0, 2, 0 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void MixedPositiveNegative_n2()
    {
        int[] nums = { -1, -2, 3, 4 };
        int n = 2;
        int[] expected = { -1, 3, -2, 4 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void LargeNumbers_n3()
    {
        int[] nums = { 1000, 2000, 3000, 9999, 8888, 7777 };
        int n = 3;
        int[] expected = { 1000, 9999, 2000, 8888, 3000, 7777 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void SinglePair_n1_LargeValues()
    {
        int[] nums = { int.MaxValue, int.MinValue };
        int n = 1;
        int[] expected = { int.MaxValue, int.MinValue };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void IncreasingSequence_n4()
    {
        int[] nums = { 1, 2, 3, 4, 5, 6, 7, 8 };
        int n = 4;
        int[] expected = { 1, 5, 2, 6, 3, 7, 4, 8 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void DecreasingSequence_n3()
    {
        int[] nums = { 6, 5, 4, 3, 2, 1 };
        int n = 3;
        int[] expected = { 6, 3, 5, 2, 4, 1 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void PalindromeLike_n2()
    {
        int[] nums = { 1, 2, 2, 1 };
        int n = 2;
        int[] expected = { 1, 2, 2, 1 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void RepeatingPattern_n3()
    {
        int[] nums = { 1, 1, 1, 2, 2, 2 };
        int n = 3;
        int[] expected = { 1, 2, 1, 2, 1, 2 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void EdgeCase_n0_EmptyArray()
    {
        int[] nums = { };
        int n = 0;
        int[] expected = { };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void LargeInput_n5()
    {
        int[] nums = { 1, 2, 3, 4, 5, 10, 20, 30, 40, 50 };
        int n = 5;
        int[] expected = { 1, 10, 2, 20, 3, 30, 4, 40, 5, 50 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void OddEvenInterleave_n4()
    {
        int[] nums = { 1, 3, 5, 7, 2, 4, 6, 8 };
        int n = 4;
        int[] expected = { 1, 2, 3, 4, 5, 6, 7, 8 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void NegativeAndPositive_n2()
    {
        int[] nums = { -5, 10, -3, 8 };
        int n = 2;
        int[] expected = { -5, -3, 10, 8 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void MaxArraySize_n10()
    {
        // Создаём массив из 20 элементов (n=10)
        int[] nums = new int[20];
        for (int i = 0; i < 10; i++)
        {
            nums[i] = i + 1;           // x1..x10: 1..10
            nums[i + 10] = 100 + i;   // y1..y10: 100..109
        }

        int n = 10;
        int[] expected = new int[20];
        for (int i = 0; i < 10; i++)
        {
            expected[2 * i] = i + 1;     // x1, x2, ..., x10 на чётных позициях
            expected[2 * i + 1] = 100 + i; // y1, y2, ..., y10 на нечётных позициях
        }

        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void AllZeros_n3()
    {
        int[] nums = { 0, 0, 0, 0, 0, 0 };
        int n = 3;
        int[] expected = { 0, 0, 0, 0, 0, 0 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void AlternatingSigns_n2()
    {
        int[] nums = { -1, -2, 1, 2 };
        int n = 2;
        int[] expected = { -1, 1, -2, 2 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void LargeNegativeNumbers_n2()
    {
        int[] nums = { -999, -888, -777, -666 };
        int n = 2;
        int[] expected = { -999, -777, -888, -666 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void SingleElementRepeated_n2()
    {
        int[] nums = { 42, 42, 42, 42 };
        int n = 2;
        int[] expected = { 42, 42, 42, 42 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void IncreasingThenDecreasing_n3()
    {
        int[] nums = { 1, 2, 3, 6, 5, 4 };
        int n = 3;
        int[] expected = { 1, 6, 2, 5, 3, 4 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void MirroredHalves_n2()
    {
        int[] nums = { 1, 2, 2, 1 };
        int n = 2;
        int[] expected = { 1, 2, 2, 1 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void FirstHalfAscending_SecondHalfDescending_n3()
    {
        int[] nums = { 1, 2, 3, 9, 8, 7 };
        int n = 3;
        int[] expected = { 1, 9, 2, 8, 3, 7 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void MixedSmallLarge_n2()
    {
        int[] nums = { 1, 999, 2, 888 };
        int n = 2;
        int[] expected = { 1, 2, 999, 888 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }

    [Fact]
    public void IdenticalPairs_n3()
    {
        int[] nums = { 5, 5, 6, 6, 7, 7 }; // 2 6 5 7 6 7
        int n = 3;
        int[] expected = { 5, 6, 5, 7, 6, 7 };
        Assert.Equal(expected, _handler.Shuffle(nums, n));
    }
}
