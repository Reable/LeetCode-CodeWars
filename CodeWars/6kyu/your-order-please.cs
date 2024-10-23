using System;
using System.Text.RegularExpressions;
using System.Collections.Generic;

//https://www.codewars.com/kata/55c45be3b2079eccff00010f/train/csharp

internal class Kata
{
    static void Main()
    {
        Console.WriteLine(Order("is2 Thi1s T4est 3a"));
        Console.WriteLine(Order("4of Fo1r pe6ople g3ood th5e the2"));
    }

    public static string Order(string words)
    {
        string[] arr = words.Split(" ");
        List<string> strArr = new List<string>(new string[arr.Length]);

        for(int i = 0; i < arr.Length; i++){
            var match = Regex.Match(arr[i], @"\d");
            
            if(match.Success){
                strArr[Convert.ToInt32(match.Value) - 1] = arr[i];
            }
        }
        
        return string.Join(" ", strArr);
    }

    public static string Order2(string words)
    {
        if(string.IsNullOrEmpty(words)) return words;
        return
            string.Join(" ",
            words.Split(" ")
                .OrderBy(
                    s => s.ToList().Find(c => char.IsDigit(c))
                )
            );
    }
}