package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'commonChild' function below.
 * --LCS--
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1    FGHQIOPLA
 *  2. STRING s2    FGHAIOPLQ
 */

func commonChild(s1 string, s2 string) int32 {
	//Write your code here
	m, n := len(s1), len(s2)
	var s [][]int = make([][]int, m+1)
	for i, _ := range s {
		s[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				s[i][j] = s[i-1][j-1] + 1
			} else {
				if s[i-1][j] > s[i][j-1] {
					s[i][j] = s[i-1][j]
				} else {
					s[i][j] = s[i][j-1]
				}

			}
		}
	}
	return int32(s[n][m])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

	fmt.Fprintf(writer, "%d\n", result)

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
