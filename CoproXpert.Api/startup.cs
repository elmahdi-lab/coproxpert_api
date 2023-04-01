namespace CoproXpert.Api;

public class Startup
{
    // OddNumbers: receives a list of numbers and returns a list of odd numbers
    public static List<int> OddNumbers(List<int> numbers)
    {
        var oddNumbers = new List<int>();
        foreach (var number in numbers)
            if (number % 2 != 0)
                oddNumbers.Add(number);
        return oddNumbers;
    }

    // EvenNumbers: receives a list of numbers and returns a list of even numbers
    public static List<int> EvenNumbers(List<int> numbers)
    {
        var evenNumbers = new List<int>();
        foreach (var number in numbers)
            if (number % 2 == 0)
                evenNumbers.Add(number);
        return evenNumbers;
    }

    // Create Unit tests for OddNumbers and EvenNumbers
    public static void Main(string[] args)
    {
        var numbers = new List<int> { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        var oddNumbers = OddNumbers(numbers);
        var evenNumbers = EvenNumbers(numbers);
        Console.WriteLine("Odd numbers: " + string.Join(", ", oddNumbers));
        Console.WriteLine("Even numbers: " + string.Join(", ", evenNumbers));
    }
}