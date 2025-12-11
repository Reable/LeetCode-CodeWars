namespace Solutions.Tests;

public class CountCharactersInStringTests
{
    [Fact]
    public void Count_ReturnsEmptyDictionary_WhenStringIsEmpty()
    {
        var result = CountCharactersInString.Count("");
        Assert.Empty(result);
    }

    [Fact]
    public void Count_ReturnsSingleEntry_ForSingleCharacter()
    {
        var result = CountCharactersInString.Count("a");
        Assert.Equal(1, result['a']);
    }

    [Fact]
    public void Count_HandlesTwoDifferentCharacters()
    {
        var result = CountCharactersInString.Count("ab");
        Assert.Equal(1, result['a']);
        Assert.Equal(1, result['b']);
    }

    [Fact]
    public void Count_HandlesRepeatedCharacters()
    {
        var result = CountCharactersInString.Count("aaa");
        Assert.Equal(3, result['a']);
    }

    [Fact]
    public void Count_HandlesMixedCounts()
    {
        var result = CountCharactersInString.Count("aabbbcc");
        Assert.Equal(2, result['a']);
        Assert.Equal(3, result['b']);
        Assert.Equal(2, result['c']);
    }

    [Fact]
    public void Count_HandlesSpacesCorrectly()
    {
        var result = CountCharactersInString.Count("a a ");
        Assert.Equal(2, result['a']);
        Assert.Equal(2, result[' ']);
    }

    [Fact]
    public void Count_DistinguishesUpperAndLowerCase()
    {
        var result = CountCharactersInString.Count("aA");
        Assert.Equal(1, result['a']);
        Assert.Equal(1, result['A']);
    }

    [Fact]
    public void Count_HandlesDigitsInsideString()
    {
        var result = CountCharactersInString.Count("a1a1");
        Assert.Equal(2, result['a']);
        Assert.Equal(2, result['1']);
    }

    [Fact]
    public void Count_CountsOnlyDigits()
    {
        var result = CountCharactersInString.Count("112233");
        Assert.Equal(2, result['1']);
        Assert.Equal(2, result['2']);
        Assert.Equal(2, result['3']);
    }

    [Fact]
    public void Count_HandlesPunctuationCharacters()
    {
        var result = CountCharactersInString.Count("!!??.");
        Assert.Equal(2, result['!']);
        Assert.Equal(2, result['?']);
        Assert.Equal(1, result['.']);
    }

    [Fact]
    public void Count_HandlesUnicodeCharacters()
    {
        var result = CountCharactersInString.Count("абвгг");
        Assert.Equal(1, result['а']);
        Assert.Equal(1, result['б']);
        Assert.Equal(1, result['в']);
        Assert.Equal(2, result['г']);
    }

    [Fact]
    public void Count_HandlesNewLineCharacter()
    {
        var result = CountCharactersInString.Count("a\na");
        Assert.Equal(2, result['a']);
        Assert.Equal(1, result['\n']);
    }

    [Fact]
    public void Count_HandlesTabCharacter()
    {
        var result = CountCharactersInString.Count("a\ta");
        Assert.Equal(2, result['a']);
        Assert.Equal(1, result['\t']);
    }

    [Fact]
    public void Count_HandlesStringWithOnlySpaces()
    {
        var result = CountCharactersInString.Count("   ");
        Assert.Equal(3, result[' ']);
    }

    [Fact]
    public void Count_HandlesLongString()
    {
        var result = CountCharactersInString.Count("aaabbbcccddd");
        Assert.Equal(3, result['a']);
        Assert.Equal(3, result['b']);
        Assert.Equal(3, result['c']);
        Assert.Equal(3, result['d']);
    }

    [Fact]
    public void Count_HandlesStringWithSpecialSymbols()
    {
        var result = CountCharactersInString.Count("@@@###");
        Assert.Equal(3, result['@']);
        Assert.Equal(3, result['#']);
    }

    [Fact]
    public void Count_HandlesMixedLettersDigitsSymbols()
    {
        var result = CountCharactersInString.Count("a1!a1!");
        Assert.Equal(2, result['a']);
        Assert.Equal(2, result['1']);
        Assert.Equal(2, result['!']);
    }

    [Fact]
    public void Count_DoesNotCollapseSimilarButDifferentCharacters()
    {
        var result = CountCharactersInString.Count("o0");
        Assert.Equal(1, result['o']);
        Assert.Equal(1, result['0']);
    }

    [Fact]
    public void Count_HandlesStringWithLeadingAndTrailingSpaces()
    {
        var result = CountCharactersInString.Count(" a ");
        Assert.Equal(1, result['a']);
        Assert.Equal(2, result[' ']);
    }
}
