package learning

import (
	"fmt"
	"sort"
	"strings"
)

// Longest subString
func LangestSubString() {
	str := "abcabcbb"
	longestSubStr := make(map[string]string)
	appendStrg := ""
	for _, c := range str {
		s := string(c)
		if !strings.Contains(appendStrg, s) {
			appendStrg = appendStrg + s
			longestSubStr[appendStrg] = appendStrg
		} else {
			appendStrg = s	
			//longestSubStr[appendStrg] = appendStrg		
		}
	}
	uniqueStrLength := 0
	result := ""
	for k, v := range longestSubStr {
		if len(k) > uniqueStrLength {
			result = v
			uniqueStrLength = len(k)
		}
	}
	fmt.Println("***Longest substring is :", result)                      //, " \n******map******:", longestSubStr)
	fmt.Println("***The length of Longest substring is : ", uniqueStrLength)
}

type Employee struct {
	EmployeeName string
	Salary       float32
	Id           int
}

func SortEmployeeStructByName() {
	employees := []Employee{
		{EmployeeName: "Satya", Salary: 500000, Id: 901},
		{EmployeeName: "Alekhya", Salary: 200000, Id: 401},
		{EmployeeName: "Loukya", Salary: 600000, Id: 301},
		{EmployeeName: "Adhvith", Salary: 700000, Id: 101},
	}
	sort.Slice(employees,
		func(i, j int) bool {
			return employees[i].EmployeeName < employees[j].EmployeeName
		})
	for _, emp := range employees {
		fmt.Printf("empName : %v   salary : %v\n", emp.EmployeeName, emp.Salary)
	}
}

func SortEmployeeStructById() {
	employees := []Employee{
		{EmployeeName: "Satya", Salary: 500000, Id: 901},
		{EmployeeName: "Alekhya", Salary: 200000, Id: 401},
		{EmployeeName: "Loukya", Salary: 600000, Id: 301},
		{EmployeeName: "Adhvith", Salary: 700000, Id: 101},
	}
	sort.Slice(employees,
		func(i, j int) bool {
			return employees[i].Id < employees[j].Id
		})
	for _, emp := range employees {
		fmt.Printf("employee Id : %v  ,employee name : %v, salary : %v\n", emp.Id, emp.EmployeeName, emp.Salary)
	}
}

// sort string array with sort.strings method
func SortStringArray() {
	strArray := []string{"Satya", "Sindhu", "Alekhya", "Loukya", "Adhvith"}
	sort.Strings(strArray)
	fmt.Println("sorted name (ascending) : ", strArray)
}

// sort string array without predefined methods using bubble sort
func SortStringArrayBubbleSort() {
	arr := []string{"Satya", "Sindhu", "Alekhya", "Loukya", "Adhvith"}
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Flag to check if any swaps were made in this pass
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements and swap if needed
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	for i, arrVal := range arr {
		fmt.Println("Sorted names", i, "th index :\t", arrVal)
	}
}

// Unique string
func UniqueStringArray() {
	strArray := []string{"hi","how","are","doing","i","am","doing","good","how","are","i","good","are","am"}
	mapStr := make(map[string]int)
	fmt.Println("string array strings : ", strArray)
	for _, str := range strArray {
		mapStr[str]++
	}
	for str, count := range mapStr {
		if count < 2 {
			fmt.Println("******unique string******:", str,",******count",count)
		}else{
		fmt.Println("dublicate string :",str,",count : ",count)
		}
	}
}

//Remove dublicate characters from string
func UniqueString(){
	str := "alekhyanmufghbfgalekhya"
	var result string
	mapStr := make(map[rune]bool)
	for _,char := range str{
		if !mapStr[char] {
			result += string(char)
			mapStr[char] = true
		}
	}
	fmt.Println("result :",result)
}

func UniqueStringUsingMap(){
	str := "abca"
	var result string
	mapStr := make(map[string]string)
	for _,ch := range str{
		s := string(ch)
		if _, isValueExists := mapStr[s]; !isValueExists {
			result += s
			mapStr[s] = s
			fmt.Println("Value does not exist :: ", s)
		}
	}
	fmt.Println("result :",result)
}