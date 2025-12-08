namespace BinarySearch.Tests;

public class CountNegativeNumbersInSortedMatrixTests
{
    
    private readonly CountNegativeNumbersInSortedMatrix _handler = new ();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        int[][] matrix = [
            [ 4,  3,  2, -1],
            [ 3,  2,  1, -1],
            [ 1,  1, -1, -2],
            [-1, -1, -2, -3]
        ];

        Assert.Equal(8, _handler.CountNegatives(matrix));
        Assert.Equal(0, _handler.CountNegatives([[3, 2], [1, 0]]));
    }
    
    [Fact]
    public void TestSingleElements()
    {
        // Положительный элемент
        Assert.Equal(0, _handler.CountNegatives(new int[][] { new int[] { 5 } }));
        
        // Отрицательный элемент
        Assert.Equal(1, _handler.CountNegatives(new int[][] { new int[] { -5 } }));
        
        // Ноль
        Assert.Equal(0, _handler.CountNegatives(new int[][] { new int[] { 0 } }));
    }

    [Fact]
    public void TestAllSameSign()
    {
        // Все положительные
        Assert.Equal(0, _handler.CountNegatives(new int[][] 
        {
            new int[] { 1, 2, 3 },
            new int[] { 4, 5, 6 }
        }));
        
        // Все отрицательные
        Assert.Equal(6, _handler.CountNegatives(new int[][]
        {
            new int[] { -1, -2, -3 },
            new int[] { -4, -5, -6 }
        }));
    }

    [Fact]
    public void TestEdgeRows()
    {
        // Последняя строка отрицательная
        Assert.Equal(2, _handler.CountNegatives(new int[][]
        {
            new int[] { 3, 2, 1 },
            new int[] { 0, -1, -2 }
        }));
    }

    [Fact]
    public void TestEdgeColumns()
    {
        // Последний столбец отрицательный
        Assert.Equal(3, _handler.CountNegatives(new int[][]
        {
            new int[] { 1, -3 },
            new int[] { 2, -2 },
            new int[] { 3, -1 }
        }));
    }

    [Fact]
    public void TestMixedPatterns()
    {
        // Ступенчатые отрицательные
        Assert.Equal(6, _handler.CountNegatives(new int[][]
        {
            new int[] { 5, 4, 3, 2 },
            new int[] { 4, 3, 2, -1 },
            new int[] { 3, 2, -1, -2 },
            new int[] { 2, -1, -2, -3 }
        }));
    }

    [Fact]
    public void TestRectangularMatrices()
    {
        // Широкая матрица
        Assert.Equal(3, _handler.CountNegatives(new int[][]
        {
            new int[] { 5, 4, 3, 2, 1 },
            new int[] { 4, 3, 2, 1, -1 },
            new int[] { 3, 2, 1, -1, -2 }
        }));
        
        // Высокая матрица
        Assert.Equal(4, _handler.CountNegatives(new int[][]
        {
            new int[] { 3, 2 },
            new int[] { 2, 1 },
            new int[] { 1, -1 },
            new int[] { 0, -2 },
            new int[] { -1, -3 }
        }));
    }

    [Fact]
    public void TestOneDimensional()
    {
        // Одна строка, все положительные
        Assert.Equal(0, _handler.CountNegatives(new int[][] { new int[] { 1, 2, 3, 4 } }));
        
        // Одна строка, смешанные
        Assert.Equal(2, _handler.CountNegatives(new int[][] { new int[] { 3, 2, -1, -2 } }));
        
        // Один столбец, все положительные
        Assert.Equal(0, _handler.CountNegatives(new int[][]
        {
            new int[] { 4 },
            new int[] { 3 },
            new int[] { 2 },
            new int[] { 1 }
        }));
        
        // Один столбец, смешанные
        Assert.Equal(2, _handler.CountNegatives(new int[][]
        {
            new int[] { 2 },
            new int[] { 1 },
            new int[] { -1 },
            new int[] { -2 }
        }));
    }

    [Fact]
    public void TestZerosAndBounds()
    {
        // Только нули
        Assert.Equal(0, _handler.CountNegatives(new int[][]
        {
            new int[] { 0, 0 },
            new int[] { 0, 0 }
        }));
        
        // Граничные значения (0 и -1)
        Assert.Equal(2, _handler.CountNegatives(new int[][]
        {
            new int[] { 0, 0 },
            new int[] { -1, -1 }
        }));
    }

    [Fact]
    public void TestLargeUniform()
    {
        // Большая матрица, все положительные
        int[][] largePositive = new int[10][];
        for (int i = 0; i < 10; i++)
        {
            largePositive[i] = new int[10];
            for (int j = 0; j < 10; j++)
                largePositive[i][j] = i + j + 1;
        }
        Assert.Equal(0, _handler.CountNegatives(largePositive));

        // Большая матрица, все отрицательные
        int[][] largeNegative = new int[5][];
        for (int i = 0; i < 5; i++)
        {
            largeNegative[i] = new int[5];
            for (int j = 0; j < 5; j++)
                largeNegative[i][j] = -i - j - 1;
        }
        Assert.Equal(25, _handler.CountNegatives(largeNegative));
    }

    [Fact]
    public void TestLeetCodeAndSimple()
    {
        // Стандартный тестовый кейс LeetCode
        int[][] leetCodeMatrix = new int[][]
        {
            new int[] { 4, 3, 2, -1 },
            new int[] { 3, 2, 1, -1 },
            new int[] { 1, 1, -1, -2 },
            new int[] { -1, -1, -2, -3 }
        };
        Assert.Equal(8, _handler.CountNegatives(leetCodeMatrix));

        // Простой случай 2x2
        Assert.Equal(0, _handler.CountNegatives(new int[][]
        {
            new int[] { 3, 2 },
            new int[] { 1, 0 }
        }));
    }
}