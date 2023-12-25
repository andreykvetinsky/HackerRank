package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Parent *TreeNode
	Child  []*TreeNode
}

func buildTreeNew(n int32, tn []int32, ptn []int32) (map[int32]*TreeNode, *TreeNode) {
	Tree := make(map[int32]*TreeNode, n)
	findRoot := make(map[int32]int, n)
	for i := 1; i <= int(n); i++ {
		Tree[int32(i)] = &TreeNode{Child: []*TreeNode{}}
		findRoot[int32(i)] = 0
	}
	var root *TreeNode
	for i := 0; i < len(tn); i++ {
		delete(findRoot, tn[i])
		i1, i2 := tn[i], ptn[i]
		s := Tree[i2].Child
		Tree[i1].Parent = Tree[i2]
		s = append(s, Tree[i1])
		Tree[i2].Child = s
	}
	for k, _ := range findRoot {
		root = Tree[k]
	}
	return Tree, root
}

func CountEdges(v *TreeNode, count *int32, res *int32) {
	var sum []int32
	if v.Child != nil {
		for _, j := range v.Child {
			var s int32
			CountEdges(j, &s, res)
			sum = append(sum, s)
		}
	}
	for _, v := range sum {
		*count += v
	}
	*count += 1
	if *count%2 == 0 && *count > 0 && v.Parent != nil {
		*res++
		*count = 0
	}
}

// Complete the evenForest function below.
func evenForest(t_nodes int32, t_edges int32, t_from []int32, t_to []int32) int32 {
	_, root := buildTreeNew(t_nodes, t_from, t_to)
	var count, res int32
	CountEdges(root, &count, &res)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tNodesEdges := strings.Split(readLine(reader), " ")

	tNodes, err := strconv.ParseInt(tNodesEdges[0], 10, 64)
	checkError(err)

	tEdges, err := strconv.ParseInt(tNodesEdges[1], 10, 64)
	checkError(err)

	var tFrom, tTo []int32

	for i := 0; i < int(tEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")

		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		tFrom = append(tFrom, int32(edgeFrom))
		tTo = append(tTo, int32(edgeTo))
	}

	res := evenForest(int32(tNodes), int32(tEdges), tFrom, tTo)

	fmt.Fprintf(writer, "%d\n", res)

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
