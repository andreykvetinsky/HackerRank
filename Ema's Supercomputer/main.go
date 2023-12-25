package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoPluses' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func twoPluses(grid []string) int32 {
	// Write your code here
	var resX []int
	var grids [][]string = make([][]string, len(grid))
	for i, v := range grid {
		grids[i] = strings.Split(v, "")

	}
	ch1 := make(chan []int, 4)
	ch2 := make(chan int, 4)
	go func() {
		for i, v := range grids {
			for j, k := range v {
				if k == "G" {
					var counter int
					in := 1
					ch2 <- 1
					ch1 <- []int{(i*100 + j)}

					mn := []int{}
					mn = append(mn, i*100+j)
					for {
						if i+in < len(grids) && i-in >= 0 && j+in < len(v) && j-in >= 0 {
							if grids[i-in][j] == "G" && grids[i+in][j] == "G" &&
								grids[i][j-in] == "G" && grids[i][j+in] == "G" {
								mn = append(mn, (i-in)*100+j)
								mn = append(mn, (i+in)*100+j)
								mn = append(mn, (i)*100+(j-in))
								mn = append(mn, (i)*100+(j+in))
								counter += 4
								ch2 <- counter + 1
								ch1 <- mn
								in++
							} else {
								break
							}
						} else {
							break
						}
					}
					mn = nil
				}
			}

		}
		close(ch1)
		close(ch2)
	}()

	for mxl := range ch1 {
		mxl2 := map[int]int{}
		for i1, v := range mxl {
			mxl2[v] = i1
		}
		kxl := <-ch2
		//fmt.Println(kxl, mxl2)
		for i, v := range grids {
			for j, k := range v {
				_, ok := mxl2[i*100+j]
				if k == "G" && !ok {
					_, oks := mxl2[i*100+j]
					if !oks {
						resX = append(resX, 1*kxl)
					}
					var counter int
					in := 1
					for {
						iin, i_in, jin, j_in := i+in, i-in, j+in, j-in
						_, ok1 := mxl2[(i-in)*100+j]
						_, ok2 := mxl2[(i+in)*100+j]
						_, ok3 := mxl2[(i)*100+(j-in)]
						_, ok4 := mxl2[(i)*100+(j+in)]
						if iin < len(grids) && i_in >= 0 && jin < len(v) && j_in >= 0 && !ok1 && !ok2 && !ok3 && !ok4 {
							if grids[i_in][j] == "G" && grids[iin][j] == "G" &&
								grids[i][j_in] == "G" && grids[i][jin] == "G" {
								counter += 4
								resX = append(resX, ((counter + 1) * kxl))
								in++
							} else {
								break
							}
						} else {
							break
						}
					}
				}
			}
		}

	}
	//    fmt.Println(resX)
	sort.Ints(resX)
	return int32(resX[len(resX)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	// mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// m := int32(mTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := twoPluses(grid)

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
