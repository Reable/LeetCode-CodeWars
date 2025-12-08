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
}
