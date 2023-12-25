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
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func formingMagicSquare(s [][]int32) int32 {
	var change int32 = 50
	// Write your code here\
	a := []int32{8, 1, 6}
	a1 := []int32{3, 5, 7}
	a11 := []int32{4, 9, 2}
	ai1 := [][]int32{a, a1, a11}

	b := []int32{6, 1, 8}
	b2 := []int32{7, 5, 3}
	b22 := []int32{2, 9, 4}
	ai2 := [][]int32{b, b2, b22}

	c := []int32{8, 3, 4}
	c3 := []int32{1, 5, 9}
	c33 := []int32{6, 7, 2}
	ai3 := [][]int32{c, c3, c33}

	f := []int32{4, 3, 8}
	f1 := []int32{9, 5, 1}
	f2 := []int32{2, 7, 6}
	ai4 := [][]int32{f, f1, f2}

	r1 := []int32{2, 7, 6}
	r2 := []int32{9, 5, 1}
	r3 := []int32{4, 3, 8}
	ai5 := [][]int32{r1, r2, r3}

	p := []int32{6, 7, 2}
	p1 := []int32{1, 5, 9}
	p2 := []int32{8, 3, 4}
	ai6 := [][]int32{p, p1, p2}

	b44 := []int32{4, 9, 2}
	b4 := []int32{3, 5, 7}
	bb := []int32{8, 1, 6}
	ai7 := [][]int32{b44, b4, bb}

	d55 := []int32{2, 9, 4}
	d5 := []int32{7, 5, 3}
	d := []int32{6, 1, 8}
	ai8 := [][]int32{d55, d5, d}

	iii := [][][]int32{ai1, ai2, ai3, ai4, ai5, ai6, ai7, ai8}

	for _, v := range iii {
		var ch int32 = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if v[i][j] == s[i][j] {
					continue
				} else {
					if v[i][j] > s[i][j] {
						ch += v[i][j] - s[i][j]
					} else if v[i][j] < s[i][j] {
						ch += s[i][j] - v[i][j]
					}
				}

			}
		}
		if ch <= change {
			change = ch
		}
	}

	return change
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
