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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	// Write your code here
	if n%2 == 0 {
		n = 2
	}
	var grids [][]string = make([][]string, len(grid))
	var gridRepead map[string]int = make(map[string]int, 0)
	for i, v := range grid {
		grids[i] = strings.Split(v, "")
	}
	var bom [][]int
	for in := 1; in <= int(n); {
		bom = nil
		if int(n) == in {
			break
		}
		for i, v := range grids {
			for j, k := range v {
				if k == "O" {
					bom = append(bom, []int{i, j})
				} else if k == "." {
					grids[i][j] = "O"
				}
			}
		}
		in++
		if int(n) == in {
			break
		}
		for _, v := range bom {
			i, j := v[0], v[1]
			grids[i][j] = "."
			if i-1 >= 0 {
				grids[i-1][j] = "."
			}
			if i+1 < len(grids) {
				grids[i+1][j] = "."
			}
			if j-1 >= 0 {
				grids[i][j-1] = "."
			}
			if j+1 < len(grids[i]) {
				grids[i][j+1] = "."
			}
		}
		in++
		if int(n) == in {
			break
		}
		var s strings.Builder
		for _, v := range grids {
			for _, k := range v {
				s.WriteString(k)
			}
		}
		i, ok := gridRepead[s.String()]
		if !ok {
			gridRepead[s.String()] = in
		} else {
			gridRepead = make(map[string]int, 0)
			n = (n-int32(i)-1)%(int32(in)-int32(i)) + int32(i) - 1
			in = 1
			continue
		}
	}
	var res []string = make([]string, len(grids))
	for i, v := range grids {
		res[i] = strings.Join(v, "")
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	//cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	//c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
