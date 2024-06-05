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
 * Complete the 'countSort' function below.
 *
 * The function accepts 2D_STRING_ARRAY arr as parameter.
 */

func countSort(arr [][]string) {
	// Write your code here
	l := len(arr)
	// s := make([]string, 100)
	res := make([][]string, 100)

	for k, v := range arr {
		i, _ := strconv.Atoi(v[0])
		if k < l/2 {
			res[i] = append(res[i], "-")
		} else {
			res[i] = append(res[i], v[1])
		}
	}
	for i := range res {
		for _, v := range res[i] {
			fmt.Fprint(os.Stdout, v+" ")
		}
	}
	// r := strings.TrimLeft(strings.Join(res, ""), " ")
	// fmt.Println(0)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]string
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []string
		for _, arrRowItem := range arrRowTemp {
			arrRow = append(arrRow, arrRowItem)
		}

		if len(arrRow) != 2 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	countSort(arr)
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
