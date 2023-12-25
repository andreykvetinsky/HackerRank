package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "sync"
)

/*
 * Complete the 'steadyGene' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING gene as parameter.
 */

func steadyGene(n int32, gene string) int32 {
	// Write your code here
	g := []byte(gene)
	c := int(n) / 4
	s := []int{0, 0, 0, 0}
	cm := make(map[int][]int, n)
	for i, v := range g {
		if v == 65 {
			s[0]++
		} else if v == 84 {
			s[1]++
		} else if v == 67 {
			s[2]++
		} else if v == 71 {
			s[3]++
		}
		cm[i] = []int{s[0], s[1], s[2], s[3]}
	}
	s[0], s[1], s[2], s[3] = c-s[0], c-s[1], c-s[2], c-s[3]

	var total int
	for i, v := range s {
		if v > 0 {
			total += v
			s[i] = 0
		} else {
			s[i] = s[i] * (-1)
		}
	}
	if total == 0 {
		return 0
	} else if total == 1 {
		return 1
	}

	m1 := make(map[byte]int, 4)
	t1 := total - 1
	m1[65], m1[84], m1[67], m1[71] = cm[t1][0], cm[t1][1], cm[t1][2], cm[t1][3]
	start, res := 0, int(n)
	for i := total; i < int(n); i++ {
		// fmt.Println(m1, total, start, res)
		for m1[65] >= s[0] && m1[84] >= s[1] && m1[67] >= s[2] && m1[71] >= s[3] {
			m1[g[start]]--
			start++
			if i-(start-1) < res {
				res = i - (start - 1)
			}
		}

		// fmt.Println(m1, total, start, res)
		m1[g[i]]++
	}

	return int32(res)
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

	gene := readLine(reader)

	result := steadyGene(n, gene)

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
