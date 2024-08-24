package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("=====String=====")
	// fmt.Printf("1. Đảo ngược chuỗi '%s': %s \n", "Hello World", preservedString("Hello World"))
	// fmt.Printf("2. Đếm số lần xuất hiện của một ký tự '%s' trong chuỗi '%s': %d \n", "l", "hello", countAppearancesCharacter("hello", "l"))
	// fmt.Printf("3. Kiểm tra chuỗi đối xứng của '%s': %t \n", "madam", checkPalindrome("madam"))
	// fmt.Printf("4. Loại bỏ khoảng trắng trong chuỗi '%s': %s \n", "hello  world", removeSpace("hello world"))
	// fmt.Printf("5. Tìm chuỗi con '%s' trong chuỗi '%s': %t \n", "ell", "hello", isSubstring("hello", "ell"))

	// fmt.Println("=====Numbers=====")
	// fmt.Printf("1. Kiểm tra số nguyên tố của %d: %t \n", 7, isPrime(7))
	// fmt.Printf("2. Tính giai thừa của %d: %d \n", 5, factorial(5))
	// fmt.Printf("3. Tìm số Fibonacci thứ %d: %d \n", 7, fibonacci(7))
	// fmt.Printf("4. Kiểm tra số hoàn hảo của %d: %t \n", 6, isPerfectNumber(6))
	// fmt.Printf("5. Tính tổng các chữ số của %d: %d \n", 123, sumOfDigits(123))

	// fmt.Println("=====Array=====")
	// arr := []int{3, 5, 7, 2, 8}
	// max, min := findMaxMin(arr)
	// fmt.Printf("1. Giá trị lớn nhất và nhỏ nhất trong mảng %v: max = %d, min = %d \n", arr, max, min)
	// fmt.Printf("2. Sắp xếp mảng %v theo thứ tự tăng dần: %v \n", arr, sortArray(arr))
	// fmt.Printf("3. Số lần xuất hiện của phần tử %d trong mảng %v: %d \n", 2, []int{1, 2, 3, 2, 2, 4}, countOccurrences([]int{1, 2, 3, 2, 2, 4}, 2))
	// fmt.Printf("4. Xóa phần tử %d khỏi mảng %v: %v \n", 2, []int{1, 2, 3, 4}, removeElement([]int{1, 2, 3, 4}, 2))
	// fmt.Printf("5. Gộp hai mảng %v và %v: %v \n", []int{1, 2}, []int{3, 4}, mergeArrays([]int{1, 2}, []int{3, 4}))

	fmt.Println("=====LeetCode=====")
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(isValid("(]"))
	//fmt.Println(romanToInt("IV"))
	fmt.Println(intToRoman(16))
}

//STRINGS
func preservedString(input string) string {
	var arr = strings.Split(input, "")
	var result = ""

	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i]
	}
	return result
}

func countAppearancesCharacter(input string, key string) int {
	var arr = strings.Split(input, "")
	var count int = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			count += 1
		}
	}

	return count
}

func checkPalindrome(input string) bool {
	return input == preservedString(input)
}

func removeSpace(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func isSubstring(input1 string, input2 string) bool {
	return strings.Contains(input1, input2)
}

//NUMBERS
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

//ARRAYS
func findMaxMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}

func sortArray(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	for i := 0; i < len(sortedArr)-1; i++ {
		for j := i + 1; j < len(sortedArr); j++ {
			if sortedArr[i] > sortedArr[j] {
				sortedArr[i], sortedArr[j] = sortedArr[j], sortedArr[i]
			}
		}
	}
	return sortedArr
}

func countOccurrences(arr []int, element int) int {
	count := 0
	for _, num := range arr {
		if num == element {
			count++
		}
	}
	return count
}

func removeElement(arr []int, element int) []int {
	result := []int{}
	for _, num := range arr {
		if num != element {
			result = append(result, num)
		}
	}
	return result
}

func mergeArrays(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}

//LEETCODE
func lengthOfLongestSubstring(s string) int {
	result := 0
	stringSplit := strings.Split(s, "")
	tempString := ""
	for _, value := range stringSplit {
		if strings.Contains(tempString, value) {
			// Get max count
			if len(tempString) > result {
				result = len(tempString)
			}
			// Get exist substring index
			index := strings.Index(tempString, value)
			// Remove exist substring in tempString
			tempString = tempString[index+1:]
		}
		tempString += value
	}

	return result
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []rune{}

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != char {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func romanToInt(s string) int {
	romanTable := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result, lastValue, currentValue int
	for i := len(s) - 1; i >= 0; i-- {
		currentValue = romanTable[s[i]]
		if currentValue < lastValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		lastValue = currentValue
	}

	return result
}

func intToRoman(num int) string {
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

		var result string

		for i := 0; i < len(values); i++ {
			for num > 0 && num >= values[i] {
				result += symbols[i]
				num -= values[i]
			}
		}
		return result
}