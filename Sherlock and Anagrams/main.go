package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	var res int32
	var m map[rune]int

	for k := 0; k < len(s); k++ {
		for j := k + 1; j <= len(s); j++ {
			m = make(map[rune]int)
			s1 := s[k:j]
			for _, v := range s1 {
				m[v]++
			}
			if j+1 > len(s) {
				break
			}
			for i := 1; i <= len(s[j:]); i++ {
				s2 := s[k+i : j+i]
				var m2 map[rune]int = make(map[rune]int)
				for _, v := range s2 {
					m2[v]++
				}
				var f bool = true
				for h, g := range m2 {
					if m[h] != g {
						f = false
						break
					}

				}
				if f {
					res++
				}
			}

		}
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
