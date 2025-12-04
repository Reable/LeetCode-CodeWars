namespace BinarySearch.Tests;

public class SearchInsertPosition35Tests
{
    private readonly SearchInsertPosition35 _sut = new();

    [Fact]
    public void BasicScenarios_TargetFound()
    {
        // Точное совпадение: в середине, в начале, в конце
        Assert.Equal(2, _sut.SearchInsert(new[] { 1, 3, 5, 7, 9 }, 5));
        Assert.Equal(0, _sut.SearchInsert(new[] { 2, 4, 6, 8 }, 2));
        Assert.Equal(3, _sut.SearchInsert(new[] { 10, 20, 30, 40 }, 40));
        
        // Один элемент — совпадение
        Assert.Equal(0, _sut.SearchInsert(new[] { 42 }, 42));
    }

    [Fact]
    public void TargetNotPresent_InsertPositions()
    {
        // Вставка перед первым элементом
        Assert.Equal(0, _sut.SearchInsert(new[] { 5, 10, 15 }, 1));
        
        // Вставка после последнего элемента
        Assert.Equal(3, _sut.SearchInsert(new[] { 3, 6, 9 }, 12));
        
        // Вставка в середину (между элементами)
        Assert.Equal(2, _sut.SearchInsert(new[] { 1, 4, 7, 10 }, 5));
        Assert.Equal(1, _sut.SearchInsert(new[] { 20, 40 }, 30));
    }

    [Fact]
    public void SingleAndTwoElementArrays()
    {
        // Один элемент: меньше, больше, равно
        Assert.Equal(0, _sut.SearchInsert(new[] { 100 }, 50));
        Assert.Equal(1, _sut.SearchInsert(new[] { 50 }, 75));
        Assert.Equal(0, _sut.SearchInsert(new[] { -42 }, -42));

        // Два элемента: найдено в первом/втором, вставка между/до/после
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20 }, 10));
        Assert.Equal(1, _sut.SearchInsert(new[] { 15, 25 }, 25));
        Assert.Equal(1, _sut.SearchInsert(new[] { 20, 40 }, 30));
        Assert.Equal(0, _sut.SearchInsert(new[] { 30, 50 }, 10));
        Assert.Equal(2, _sut.SearchInsert(new[] { 10, 20 }, 30));
    }

    [Fact]
    public void EdgeCases_EmptyAndDuplicates()
    {
        // Пустой массив
        Assert.Equal(0, _sut.SearchInsert(Array.Empty<int>(), 10));

        // Все элементы одинаковые: совпадение и вне диапазона
        Assert.Equal(0, _sut.SearchInsert(new[] { 7, 7, 7, 7 }, 7));
        Assert.Equal(0, _sut.SearchInsert(new[] { 8, 8, 8 }, 5));
        Assert.Equal(3, _sut.SearchInsert(new[] { 6, 6, 6 }, 9));

        // Дубликаты: первое вхождение и вставка между
        Assert.Equal(1, _sut.SearchInsert(new[] { 1, 2, 2, 2, 3 }, 2));
        Assert.Equal(4, _sut.SearchInsert(new[] { 1, 3, 3, 3, 5 }, 4));
    }

    [Fact]
    public void NegativeAndMixedNumbers()
    {
        // Отрицательные числа: найдено, вставка до/после/в середину
        Assert.Equal(1, _sut.SearchInsert(new[] { -10, -5, 0, 5, 10 }, -5));
        Assert.Equal(2, _sut.SearchInsert(new[] { -20, -10, 0, 10 }, -5));
        Assert.Equal(0, _sut.SearchInsert(new[] { -5, -3, -1 }, -10));
        Assert.Equal(3, _sut.SearchInsert(new[] { -100, -50, -25 }, 0));

        // Смешанные знаки: ноль, вставка у границы
        Assert.Equal(2, _sut.SearchInsert(new[] { -3, -1, 0, 2, 4 }, 0));
        Assert.Equal(0, _sut.SearchInsert(new[] { 1, 2, 3 }, 0));
    }

    [Fact]
    public void LargeArrays_BoundaryScenarios()
    {
        int[] large = Enumerable.Range(1, 1000).ToArray();

        // Границы большого массива
        Assert.Equal(0,  _sut.SearchInsert(large, 1));
        Assert.Equal(999, _sut.SearchInsert(large, 1000));
        Assert.Equal(499, _sut.SearchInsert(large, 500));

        // Вставка внутри большого массива
        int[] sub = Enumerable.Range(10, 100).ToArray(); // 10..109
        Assert.Equal(5,  _sut.SearchInsert(sub, 15));
        Assert.Equal(80, _sut.SearchInsert(Enumerable.Range(100, 100).ToArray(), 180));
    }

    [Fact]
    public void ConsecutiveAndGappedArrays()
    {
        // Последовательные числа
        Assert.Equal(2, _sut.SearchInsert(new[] { 1, 2, 3, 4, 5 }, 3));
        Assert.Equal(2, _sut.SearchInsert(new[] { 5, 6, 7, 8 }, 7)); // 6.5 → 6 (целочисленное)

        // Массивы с пропусками
        Assert.Equal(2, _sut.SearchInsert(new[] { 1, 5, 9, 13 }, 7));
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20, 30 }, 5));
        Assert.Equal(3, _sut.SearchInsert(new[] { 2, 8, 14 }, 20));
    }

    [Fact]
    public void MinMaxAndBoundaryValues()
    {
        // Совпадение с min/max
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20, 30, 40 }, 10));
        Assert.Equal(3, _sut.SearchInsert(new[] { 15, 25, 35, 45 }, 45));

        // Вставка на границе ёмкости
        Assert.Equal(1, _sut.SearchInsert(new[] { 100, 200 }, 150));

        // Трёхэлементные сценарии
        Assert.Equal(1, _sut.SearchInsert(new[] { 7, 14, 21 }, 14));
        Assert.Equal(1, _sut.SearchInsert(new[] { 3, 9, 15 }, 6));
        Assert.Equal(2, _sut.SearchInsert(new[] { 22, 44, 66 }, 55));
    }

    [Fact]
    public void IdenticalPairsAndAlternatingSigns()
    {
        // Одинаковые пары: поиск первого вхождения и вставка между парами
        Assert.Equal(2, _sut.SearchInsert(new[] { 2, 2, 4, 4, 6, 6 }, 4));
        Assert.Equal(4, _sut.SearchInsert(new[] { 1, 3, 3, 3, 5 }, 4));

        // Чередующиеся знаки: поиск положительного и отрицательного числа
        Assert.Equal(3, _sut.SearchInsert(new[] { -8, -4, 0, 4, 8 }, 4));
        Assert.Equal(1, _sut.SearchInsert(new[] { -8, -4, 0, 4, 8 }, -4));
    }

    [Fact]
    public void ZeroInArray_Scenarios()
    {
        // Массив с нулём: поиск нуля, вставка до/после нуля
        Assert.Equal(2, _sut.SearchInsert(new[] { -3, -1, 0, 2, 5 }, 0));
        Assert.Equal(2, _sut.SearchInsert(new[] { -5, -2, 0, 3, 7 }, -1));
        Assert.Equal(3, _sut.SearchInsert(new[] { -4, -1, 0, 6, 9 }, 1));
    }

    [Fact]
    public void LargeRangeWithNegativeAndPositive()
    {
        // Большой диапазон от -500 до 500: поиск нуля и границ
        int[] largeRange = Enumerable.Range(-500, 1001).ToArray();
        Assert.Equal(500, _sut.SearchInsert(largeRange, 0));
        Assert.Equal(0,  _sut.SearchInsert(largeRange, -500));
        Assert.Equal(1000, _sut.SearchInsert(largeRange, 500));


        // Диапазон от -1000 до 1000: поиск крайних значений
        int[] widerRange = Enumerable.Range(-1000, 2001).ToArray();
        Assert.Equal(0,   _sut.SearchInsert(widerRange, -1000));
        Assert.Equal(2000, _sut.SearchInsert(widerRange, 1000));
    }

    [Fact]
    public void SmallGapsAndTightRanges()
    {
        // Малый зазор между элементами: вставка в плотный диапазон
        Assert.Equal(1, _sut.SearchInsert(new[] { 1, 3 }, 2));

        // Трёхэлементный массив: различные сценарии вставки
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20, 30 }, 5));
        Assert.Equal(3, _sut.SearchInsert(new[] { 10, 20, 30 }, 35));
        Assert.Equal(1, _sut.SearchInsert(new[] { 5, 15, 25 }, 10));
    }

    [Fact]
    public void SpecialCases_UnusualInputs()
    {
        // Один большой отрицательный/положительный элемент
        Assert.Equal(0, _sut.SearchInsert(new[] { -9999 }, -9999));
        Assert.Equal(0, _sut.SearchInsert(new[] { 9999 }, 9999));

        // Некорректно отсортированный массив (проверка устойчивости)
        int[] unsorted = { 5, 4, 3, 2, 1 };
        int result = _sut.SearchInsert(unsorted, 3);
        Assert.True(result >= 0 && result <= 5); // результат в допустимом диапазоне


        // Вставка ровно между двумя значениями
        Assert.Equal(1, _sut.SearchInsert(new[] { 1, 3 }, 2));
    }

    [Fact]
    public void ComprehensiveEdgeCases()
    {
        // Пустой массив
        Assert.Equal(0, _sut.SearchInsert(Array.Empty<int>(), 42));


        // Один элемент (все варианты)
        Assert.Equal(0, _sut.SearchInsert(new[] { 42 }, 42));  // совпадение
        Assert.Equal(0, _sut.SearchInsert(new[] { 10 }, 5));   // меньше
        Assert.Equal(1, _sut.SearchInsert(new[] { 10 }, 15));  // больше


        // Два элемента (все основные сценарии)
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20 }, 10)); // найдено первое
        Assert.Equal(1, _sut.SearchInsert(new[] { 10, 20 }, 20)); // найдено второе
        Assert.Equal(1, _sut.SearchInsert(new[] { 10, 20 }, 15)); // между
        Assert.Equal(0, _sut.SearchInsert(new[] { 10, 20 }, 5));  // до
        Assert.Equal(2, _sut.SearchInsert(new[] { 10, 20 }, 25)); // после

        // Три элемента (поиск и вставка)
        Assert.Equal(1, _sut.SearchInsert(new[] { 7, 14, 21 }, 14));
        Assert.Equal(1, _sut.SearchInsert(new[] { 3, 9, 15 }, 6));
        Assert.Equal(2, _sut.SearchInsert(new[] { 22, 44, 66 }, 55));
    }
}
