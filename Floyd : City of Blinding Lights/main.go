package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	N      int32
	Waight map[int32]int32
	Nodes  map[*Node]int32
}

func buildTree(rN int32, rF []int32, rT []int32, rW []int32) map[int32]*Node {
	tree := make(map[int32]*Node, rN)
	mr := make(map[int32]int32)
	var i, j int32
	for i = 1; i <= rN; i++ {
		tree[i] = &Node{N: i, Nodes: make(map[*Node]int32), Waight: make(map[int32]int32)}
		mr[i]++
	}
	for j = 0; j < int32(len(rF)); j++ {
		tree[rF[j]].Nodes[tree[rT[j]]] = rW[j]
		delete(mr, rT[j])
	}

	if len(mr) > 0 {
		for k, _ := range mr {
			var sum int32
			m := make(map[int32]int32)
			sumWeight(tree[k], sum, m)
		}
	} else {
		var sum int32
		m := make(map[int32]int32)
		sumWeight(tree[1], sum, m)

	}
	return tree
}

func sumWeight(root *Node, sum int32, m map[int32]int32) {
	// i := 0
	for k, v := range m {
		b, ok := root.Waight[k]
		if ok {
			if b > v+sum || b == 0 {
				root.Waight[k] = v + sum
				// i++
			}
		} else {
			root.Waight[k] = v + sum
			// i++
		}
	}
	root.Waight[root.N] = 0
	// if i != 0 {
	for d, w := range root.Nodes {
		sumWeight(d, w, root.Waight)
	}
	// }

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	roadNodesEdges := strings.Split(readLine(reader), " ")

	roadNodes, err := strconv.ParseInt(roadNodesEdges[0], 10, 64)
	checkError(err)

	roadEdges, err := strconv.ParseInt(roadNodesEdges[1], 10, 64)
	checkError(err)

	var roadFrom, roadTo, roadWeight []int32

	for i := 0; i < int(roadEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")

		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		edgeWeight, err := strconv.ParseInt(edgeFromToWeight[2], 10, 64)
		checkError(err)

		roadFrom = append(roadFrom, int32(edgeFrom))
		roadTo = append(roadTo, int32(edgeTo))
		roadWeight = append(roadWeight, int32(edgeWeight))
	}

	tree := buildTree(int32(roadNodes), roadFrom, roadTo, roadWeight)
	// fmt.Println(tree[9])
	var bs strings.Builder

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		xTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		v, ok := tree[y].Waight[x]
		// fmt.Println(x, y, v, ok, tree[y].Waight, tree[y] )
		if ok {
			if qItr != int(q-1) {
				fmt.Fprintf(&bs, "%d\n", v)
			} else {
				fmt.Fprintf(&bs, "%d", v)
			}

		} else {
			if qItr != int(q-1) {
				fmt.Fprintf(&bs, "%d\n", -1)
			} else {
				fmt.Fprintf(&bs, "%d", -1)
			}

		}

	}
	fmt.Println(bs.String())
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

// 9 13
// 8 1 7
// 7 8 3
// 1 5 10
// 1 4 3
// 4 5 5
// 5 9 5
// 4 9 20
// 2 5 4
// 2 6 3
// 6 9 7
// 3 9 12
// 3 7 1
// 7 9 3
// 9
// 8 4
// 8 8
// 8 9
// 4 9
// 1 5
// 2 9
// 3 9
// 3 1
// 8 1
