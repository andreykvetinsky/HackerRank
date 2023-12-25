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
 * Complete the 'quickestWayUp' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY ladders
 *  2. 2D_INTEGER_ARRAY snakes
 */
type Step struct {
	Number int32
	Steps  []*Step
}

func findRes(step *Step, res *int32, m map[int32]int32, counter int32, rs int32, cs int32) {

	if step.Number == 100 {
		if counter >= 0 {
			rs++
		}
		if *res > rs {
			*res = rs
		}

	} else {
		counter++
		_, ok := m[step.Number]
		if !ok && cs < 5 {
			m[step.Number] = 1
			if counter == 6 {
				rs++
				counter = 0
			}
			if len(step.Steps) == 1 {
				cs = 0
			} else {
				if step.Steps[1].Number < step.Number {
					cs++
				}
			}
			for _, v := range step.Steps {
				c := counter

				mi := make(map[int32]int32, len(m))

				for k, v := range m {
					mi[k] = v
				}

				if v.Number != step.Number+1 {
					if c > 0 {
						rs++
					}
					c = -1
				}

				findRes(v, res, mi, c, rs, cs)
			}
		}
	}
}

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	play := make(map[int32]*Step, 100)
	var i int32
	play[1] = &Step{Number: 1}
	for i = 2; i <= 100; i++ {
		play[i] = &Step{Number: i}
		play[i-1].Steps = []*Step{play[i]}
	}
	for _, v := range ladders {
		play[v[0]].Steps = append(play[v[0]].Steps, play[v[1]])
	}

	for _, v := range snakes {
		play[v[0]].Steps = append(play[v[0]].Steps, play[v[1]])
	}
	var res int32 = 100 //:= make([]int32, 0, 1000)
	m := make(map[int32]int32, 100)
	var c, rs, cs int32
	m[1] = 1
	findRes(play[2], &res, m, c, rs, cs)
	// fmt.Println(res)
	if res == 100 {
		return -1
	}

	return res
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
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var ladders [][]int32
		for i := 0; i < int(n); i++ {
			laddersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var laddersRow []int32
			for _, laddersRowItem := range laddersRowTemp {
				laddersItemTemp, err := strconv.ParseInt(laddersRowItem, 10, 64)
				checkError(err)
				laddersItem := int32(laddersItemTemp)
				laddersRow = append(laddersRow, laddersItem)
			}

			if len(laddersRow) != 2 {
				panic("Bad input")
			}

			ladders = append(ladders, laddersRow)
		}

		mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		m := int32(mTemp)

		var snakes [][]int32
		for i := 0; i < int(m); i++ {
			snakesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var snakesRow []int32
			for _, snakesRowItem := range snakesRowTemp {
				snakesItemTemp, err := strconv.ParseInt(snakesRowItem, 10, 64)
				checkError(err)
				snakesItem := int32(snakesItemTemp)
				snakesRow = append(snakesRow, snakesItem)
			}

			if len(snakesRow) != 2 {
				panic("Bad input")
			}

			snakes = append(snakes, snakesRow)
		}

		result := quickestWayUp(ladders, snakes)

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
