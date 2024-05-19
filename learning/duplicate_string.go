package learning

import (
	"log"
	"sort"
	"strconv"
	//"strconv"
)

func DuplicateFollderName(strArray []string) []string{
	counts := DuplicateString(strArray)
	var files []string
	//log.Println("*********Debug********counts : ",counts)
	for str,count := range counts{
		if count > 1 {
			count = count -1
			//log.Println("the file names are : ",str + "", str + strconv.Itoa(count) )
			files = append(files, str+"", str + strconv.Itoa(count))
		}else{
			//log.Println("the file name : ",str +"" )
			files = append(files, str+"")
		}
	}
	sort.Strings(files)
	log.Printf("files are %v",files)
	return files
}

func DuplicateString(stringFiles []string) map[string]int {
	counts := make(map[string]int)

	for _, str := range stringFiles {
		counts[str]++
		log.Printf("counts : %v", counts)
	}
	for str, count := range counts {
		log.Println(str, count)
	}
	return counts
}
