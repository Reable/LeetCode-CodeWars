namespace Solutions.Tests.Quests.Array2;

using Solutions.Quests.Array2;

public class Array_2_1_Tests
{
    private readonly Array_2_1 _handler = new();

    [Fact]
    public void SimpleTest()
    {
        Assert.Equal([2,3], _handler.FindErrorNums([1,2,2,4]));
        Assert.Equal([2,3], _handler.FindErrorNums([1,2,2,4]));
    }

}
