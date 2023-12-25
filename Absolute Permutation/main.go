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
 * Complete the 'absolutePermutation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 */

func absolutePermutation(n int32, k int32) []int32 {
	ns := make([]int32, 0)
	var i int32
	for i = 1; i <= n; i++ {
		ns = append(ns, i)
	}
	if k == 0 {
		return ns
	} else if n%k != 0 || (n/k)%2 != 0 {
		return []int32{-1}
	} else {
		p := make([]int32, 0, len(ns))
		for i := int(k); i <= len(ns); i += int(k * 2) {
			p = append(p, ns[i:i+int(k)]...)
			p = append(p, ns[i-int(k):i]...)
		}
		ns = p
	}
	return ns
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
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		result := absolutePermutation(n, k)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
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
