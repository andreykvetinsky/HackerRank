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
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */
func merge_two_list(a, b []int32) []int32 {
	r := make([]int32, 0, len(a)+len(b))
	// *c++
	var i, j int

	for i < len(a) && j < len(b) {
		//  *c+= i

		if a[i] <= b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++

		}
	}
	if i < len(a) {
		r = append(r, a[i:]...)
	}
	if j < len(b) {
		r = append(r, b[j:]...)
	}
	return r
}

func mergeSort(s []int32) []int32 {
	if len(s) == 1 {
		return s
	}
	middle := len(s) / 2
	left := mergeSort(s[:middle])
	right := mergeSort(s[middle:])
	return merge_two_list(left, right)
}
func binarSearchLeft(s []int32, I int32) int {
	L := 0
	R := len(s)
	for L < R {
		m := (L + R) / 2
		if s[m] < I {
			L = m + 1
		} else {
			R = m
		}
	}
	return L
}
func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	var res int32
	ch := make(chan int32, 10)
	s := expenditure[:d]
	s = mergeSort(s)
	go func() {
		for i := int(d); i < len(expenditure); i++ {

			if d%2 == 0 {
				ch <- s[int(d/2)] + s[int(d/2)-1]
			} else {
				ch <- s[int(d/2)] * 2
			}
			i1 := binarSearchLeft(s, expenditure[i-int(d)])
			s = append(s[:i1+1], s[i1+2:]...)

			i2 := binarSearchLeft(s, expenditure[i])
			// fmt.Println(i2, s, expenditure[i])
			if i2 == len(s) {
				s = append(s, expenditure[i])
			} else {
				s = append(s[:i2+1], s[i2:]...)
				s[i2+1] = expenditure[i]
			}

			// fmt.Println(i2, s, expenditure[i])

		}
		close(ch)
	}()
	i := d
	for v := range ch {
		if expenditure[i] >= v {
			res++
		}
		i++
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

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
