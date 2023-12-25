package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'pSequences' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pSequences(n, p int32) int32 {
	var MOD int = int(math.Pow10(9)) + 7
	var sqrt int = 1
	for sqrt*sqrt <= int(p) {
		sqrt++
	}
	sqrt--
	var interlen []int = make([]int, sqrt*2)
	for i := 0; i < int(sqrt); i++ {
		interlen[i] = 1
		interlen[i+sqrt] = int(p)/(sqrt-i) - int(p)/(sqrt-i+1)
	}
	interlen[sqrt] = int(p)/sqrt - sqrt

	var currnum []int = make([]int, sqrt*2)

	currnum[0] = 1
	for i := 0; i < int(n)+1; i++ {
		var total int
		var nextnum []int = make([]int, 2*sqrt)
		for j := 2*sqrt - 1; j >= 0; j-- {
			total = (total + currnum[2*sqrt-j-1]) % MOD
			nextnum[j] = total
		}
		for j := 0; j < 2*sqrt; j++ {
			currnum[j] = (nextnum[j] * interlen[j]) % MOD
		}
	}
	return int32(currnum[0])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pSequences(n, p)

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
