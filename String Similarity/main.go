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
 * Complete the 'stringSimilarity' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func stringSimilarity(s string) int {
	// Write your code here
	var result, length = len(s), len(s)
	var right, left = 0, 0
	var z []int = make([]int, length)

	for i := 1; i < length; i++ {

		if i <= right {
			if right-i+1 < z[i-left] {
				z[i] = right - i + 1
			} else {
				z[i] = z[i-left]
			}
		}

		for i+z[i] < length && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		if i+z[i]-1 > right {
			left = i
			right = i + z[i] - 1
		}
		result += z[i]
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := stringSimilarity(s)

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
