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
 * Complete the 'almostSorted' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func almostSorted(arr []int32) {
	// Write your code here
	var counter int32
	var different []int
	for i, v := range arr {
		if i != 0 {
			if v < arr[i-1] {
				different = append(different, i)
				counter++
			}
		}
	}
	if counter == 0 {
		fmt.Println("yes")
		return
	}

	if len(arr) == len(different)+1 && len(arr) > 2 {
		fmt.Println("yes")
		fmt.Println("reverse", 0+1, len(arr)-1+1)
		return
	} else if len(different) == 1 && len(arr) == 2 {
		fmt.Println("yes")
		fmt.Println("swap", 0+1, 1+1)
		return
	} else if len(different) == 2 {
		c := 0
		arr[different[0]-1], arr[different[1]] = arr[different[1]], arr[different[0]-1]
		for i, v := range arr {
			if i != 0 {
				if v < arr[i-1] {
					c++
				}
			}
		}
		if c == 0 {
			fmt.Println("yes")
			fmt.Println("swap", different[0], different[1]+1)
			return
		} else {
			fmt.Println("no")
			return
		}
	}

	counter = 0
	if len(different) > 2 {
		for i, v := range different {
			if i != 0 {
				if v != different[i-1]+1 {
					counter++
				}
			}
		}
	}

	if counter != 0 {
		fmt.Println("no")
		return
	} else {
		sch := different[0] - 1
		end := different[len(different)-1]
		if sch == 0 && arr[end+1] >= arr[0] {
			fmt.Println("yes")
			fmt.Println("reverse", 0+1, end+1)
			return
		} else if end == len(arr)-1 && arr[sch-1] <= arr[len(arr)-1] {
			fmt.Println("yes")
			fmt.Println("reverse", sch+1, len(arr)-1+1)
			return
		} else if sch != 0 && end < len(arr)-1 && arr[sch-1] <= arr[end] && arr[end+1] >= arr[sch] {
			if len(different) > 1 {
				fmt.Println("yes")
				fmt.Println("reverse", sch+1, end+1)
				return
			} else {
				fmt.Println("yes")
				fmt.Println("swap", sch+1, end+1)
				return
			}

		} else {
			fmt.Println("no")
			return
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	almostSorted(arr)
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
