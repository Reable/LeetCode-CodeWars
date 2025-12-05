namespace BinarySearch.Tests;

public class FindSmallestLetterGreaterThanTarget744Tests
{
    private readonly FindSmallestLetterGreaterThanTarget744 _handler = new();
    
    [Fact]
    public void LeetCodeSimpleTest()
    {
        char[] letters1 = ['c', 'f', 'j'];
        char[] letters2 = ['x','x','y','y'];
        char[] letters3 = ['e','e','e','e','e','e','n','n','n','n'];
        
        Assert.Equal('c', _handler.NextGreatestLetter(letters1, 'a'));
        Assert.Equal('f', _handler.NextGreatestLetter(letters1, 'c'));
        Assert.Equal('x', _handler.NextGreatestLetter(letters2, 'z'));
        Assert.Equal('n', _handler.NextGreatestLetter(letters3, 'e'));
    }

    // 1. Базовые случаи: очевидное следующее значение
    [Fact]
    public void BasicNextLetterCases()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'c'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'b', 'c'], 'b'));
        Assert.Equal('d', _handler.NextGreatestLetter(['a', 'c', 'd'], 'c'));
    }

    // 2. Когда целевого символа нет, возвращаем первый
    [Fact]
    public void TargetLargerThanAllReturnsFirst()
    {
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'b', 'c'], 'z'));
        Assert.Equal('x', _handler.NextGreatestLetter(['x', 'y', 'z'], 'z'));
        Assert.Equal('m', _handler.NextGreatestLetter(['m', 'n', 'o'], 'p'));
    }

    // 3. Повторяющиеся символы в массиве
    [Fact]
    public void DuplicatesInLetters()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'a', 'b', 'b'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'b', 'b', 'c', 'c'], 'b'));
        Assert.Equal('x', _handler.NextGreatestLetter(['w', 'w', 'x', 'x'], 'w'));
    }

    // 4. Целевой символ равен одному из элементов
    [Fact]
    public void TargetEqualsSomeElement()
    {
        Assert.Equal('f', _handler.NextGreatestLetter(['c', 'f', 'j'], 'c'));
        Assert.Equal('j', _handler.NextGreatestLetter(['c', 'f', 'j'], 'f'));
        Assert.Equal('n', _handler.NextGreatestLetter(['e', 'e', 'n', 'n'], 'e'));
    }

    // 5. Крайние значения: 'a' и 'z'
    [Fact]
    public void EdgeLettersAndTargets()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b'], 'a')); // min letter
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'z'], 'z')); // max target
        Assert.Equal('y', _handler.NextGreatestLetter(['x', 'y', 'z'], 'x'));
    }

    // 6. Короткие массивы (минимальная длина 2)
    [Fact]
    public void MinimalLengthArrays()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b'], 'a'));
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'b'], 'b')); // z > b
        Assert.Equal('c', _handler.NextGreatestLetter(['b', 'c'], 'b'));
    }

    // 7. Большие массивы с множеством повторов
    [Fact]
    public void LargeArraysWithDuplicates()
    {
        char[] letters = new char[100];
        for (int i = 0; i < 50; i++) letters[i] = 'a';
        for (int i = 50; i < 75; i++) letters[i] = 'm';
        for (int i = 75; i < 100; i++) letters[i] = 'z';

        Assert.Equal('m', _handler.NextGreatestLetter(letters, 'a'));
        Assert.Equal('z', _handler.NextGreatestLetter(letters, 'm'));
        Assert.Equal('a', _handler.NextGreatestLetter(letters, 'z')); // wrap around
    }

    // 8. Все символы одинаковые (кроме одного)
    [Fact]
    public void MostlySameCharacters()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'a', 'a', 'b'], 'a'));
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'a', 'b', 'b'], 'b'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'a', 'c', 'c'], 'a'));
    }

    // 9. Целевой символ между группами повторов
    [Fact]
    public void TargetBetweenDuplicateGroups()
    {
        Assert.Equal('m', _handler.NextGreatestLetter(['a', 'a', 'm', 'm'], 'g')); // g между a и m
        Assert.Equal('z', _handler.NextGreatestLetter(['k', 'k', 'z', 'z'], 'm')); // m между k и z
        Assert.Equal('t', _handler.NextGreatestLetter(['p', 'p', 't', 't'], 'r')); // r между p и t
    }

    // 10. Проверка на устойчивость: все буквы алфавита
    [Fact]
    public void FullAlphabetSpan()
    {
        char[] alphabet = "abcdefghijklmnopqrstuvwxyz".ToCharArray();
        
        Assert.Equal('b', _handler.NextGreatestLetter(alphabet, 'a'));
        Assert.Equal('y', _handler.NextGreatestLetter(alphabet, 'x'));
        Assert.Equal('a', _handler.NextGreatestLetter(alphabet, 'z')); // wrap
    }

    // 11. Только две уникальные буквы
    [Fact]
    public void OnlyTwoDistinctLetters()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'a', 'b', 'b'], 'a'));
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'a', 'b', 'b'], 'b'));
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'b', 'b'], 'a'));
    }

    // 12. Целевой символ отсутствует в массиве
    [Fact]
    public void TargetNotInLetters()
    {
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'b', 'c'], 'g')); // g нет в массиве
        Assert.Equal('i', _handler.NextGreatestLetter(['i', 'j', 'k'], 'l')); // l нет
        Assert.Equal('o', _handler.NextGreatestLetter(['o', 'p', 'q'], 'r'));
    }

    // 13. Массив с одним «островом» значений
    [Fact]
    public void SingleValueCluster()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'b', 'b'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'a', 'c', 'c'], 'b'));
        Assert.Equal('d', _handler.NextGreatestLetter(['b', 'b', 'd', 'd'], 'c'));
    }

    // 14. Символы идут подряд без пропусков
    [Fact]
    public void ConsecutiveLetters()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'c'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'b', 'c'], 'b'));
        Assert.Equal('d', _handler.NextGreatestLetter(['b', 'c', 'd'], 'c'));
    }

    // 15. Пропуски в алфавитном порядке
    [Fact]
    public void SkippedLettersInArray()
    {
        Assert.Equal('e', _handler.NextGreatestLetter(['a', 'e', 'i'], 'a'));
        Assert.Equal('i', _handler.NextGreatestLetter(['a', 'e', 'i'], 'e'));
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'e', 'i'], 'i')); // wrap
    }

    // 16. Целевой символ меньше всех элементов массива (включая первый)
    [Fact]
    public void TargetSmallerThanFirst()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['b', 'c', 'd'], 'a'));
        Assert.Equal('n', _handler.NextGreatestLetter(['n', 'o', 'p'], 'm'));
        Assert.Equal('y', _handler.NextGreatestLetter(['y', 'z'], 'x'));
        Assert.Equal('h', _handler.NextGreatestLetter(['h', 'h', 'i', 'i'], 'g'));
        Assert.Equal('l', _handler.NextGreatestLetter(['l', 'l', 'm', 'n', 'o'], 'k'));
    }
    
    // 17. Целевой символ точно равен последнему элементу
    [Fact]
    public void TargetEqualsLastElement()
    {
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'b'], 'b'));
        Assert.Equal('w', _handler.NextGreatestLetter(['w', 'x'], 'x'));
        Assert.Equal('k', _handler.NextGreatestLetter(['k', 'l', 'm'], 'm'));
    }

    // 18. Массив из двух элементов, целевой между ними
    [Fact]
    public void TwoElementsTargetBetween()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b'], 'a')); // a < target < b
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'c'], 'b')); // b между a и c
        Assert.Equal('z', _handler.NextGreatestLetter(['y', 'z'], 'y'));
    }

    // 19. Целевой символ больше всех, кроме последнего
    [Fact]
    public void TargetGreaterThanAllExceptLast()
    {
        Assert.Equal('z', _handler.NextGreatestLetter(['a', 'y', 'z'], 'y'));
        Assert.Equal('d', _handler.NextGreatestLetter(['a', 'b', 'd'], 'b'));
        Assert.Equal('f', _handler.NextGreatestLetter(['c', 'e', 'f'], 'e'));
    }

    // 20. Все символы, кроме первого, одинаковые
    [Fact]
    public void FirstUniqueThenSame()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'b', 'b'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'c', 'c'], 'a'));
        Assert.Equal('d', _handler.NextGreatestLetter(['b', 'd', 'd', 'd'], 'b'));
    }

    // 21. Все символы, кроме последнего, одинаковые
    [Fact]
    public void SameThenUniqueLast()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'a', 'b'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['b', 'b', 'c'], 'b'));
        Assert.Equal('z', _handler.NextGreatestLetter(['y', 'y', 'z'], 'y'));
    }

    // 22. Целевой символ совпадает с несколькими элементами
    [Fact]
    public void TargetMatchesMultipleElements()
    {
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'a', 'c', 'c'], 'a'));
        Assert.Equal('d', _handler.NextGreatestLetter(['b', 'b', 'd', 'd'], 'b'));
        Assert.Equal('f', _handler.NextGreatestLetter(['e', 'e', 'f', 'f'], 'e'));
    }

    // 23. Массив с одним уникальным «пиком»
    [Fact]
    public void SinglePeakInArray()
    {
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'a', 'c', 'a'], 'a'));
        Assert.Equal('d', _handler.NextGreatestLetter(['b', 'b', 'd', 'b'], 'b'));
        Assert.Equal('e', _handler.NextGreatestLetter(['c', 'c', 'e', 'c'], 'c'));
    }

    // 24. Целевой символ — минимальная возможная буква ('a')
    [Fact]
    public void TargetIsMinLetterA()
    {
        Assert.Equal('b', _handler.NextGreatestLetter(['a', 'b', 'c'], 'a'));
        Assert.Equal('c', _handler.NextGreatestLetter(['a', 'c', 'd'], 'a'));
        Assert.Equal('x', _handler.NextGreatestLetter(['a', 'x', 'y'], 'a'));
    }

    // 25. Целевой символ — максимальная возможная буква ('z')
    [Fact]
    public void TargetIsMaxLetterZ()
    {
        Assert.Equal('a', _handler.NextGreatestLetter(['a', 'b', 'c'], 'z'));
        Assert.Equal('x', _handler.NextGreatestLetter(['x', 'y', 'z'], 'z'));
        Assert.Equal('m', _handler.NextGreatestLetter(['m', 'n', 'o'], 'z'));
    }
    
}