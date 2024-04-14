package learning

import (
	"fmt"
	"strings"
)

// func longestSubstring(s string) string {
// 	var start, maxLen, maxStart int

// 	// Create a map to store the last seen index of each character
// 	lastSeen := make(map[rune]int)

// 	for i, char := range s {
// 		// Check if the character has been seen before in the current substring
// 		if lastIndex, ok := lastSeen[char]; ok && lastIndex >= start {
// 			// If yes, update the start index to the next character after the last occurrence of the current character
// 			start = lastIndex + 1
// 		}

// 		// Update the last seen index of the current character
// 		lastSeen[char] = i

// 		// Update the maximum length and start index of the longest substring
// 		if i-start+1 > maxLen {
// 			maxLen = i - start + 1
// 			maxStart = start
// 		}
// 	}

// 	return s[maxStart : maxStart+maxLen]
// }

func LangestSubString() {
	//s := "aabcabcbbz12345yz" //result abc
	//s := "ssdvnscawerchjk" //sdnc, scawer, dvnscawer
	s := "advbcdabcvdvcs"
	longest := LongestSubstring1(s)
	fmt.Println("Longest substring without repeating characters:", longest)
}

func LongestSubstring1(str string) string {
	//var strg string
	 longestSubStr := make(map[string]string)
	 appendStrg := ""
	for _, c := range str {
		//fmt.Println("strg : ", i, string(c),string(c+1))
		s := string(c)

		if !strings.Contains(appendStrg, s) {
			appendStrg = appendStrg +s
		}else{
			longestSubStr[appendStrg] = appendStrg
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

	fmt.Println("****Final result : ", result)
	//fmt.Println("longestSubStr : ", longestSubStr)
	return ""
}
