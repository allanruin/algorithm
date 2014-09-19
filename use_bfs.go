package main

import (
	"fmt"
	"me.irix/graph"
)

/*

                +------+
                |      |
                |   0  |
                +---+--+
                    |
                 +--+--+
                 |  1  |
                 |     |
             +---+-----+--+
             |            |
             |            |
         +---+--+      +--+--+
         |  2   |      | 3   |
         |      |      |     |
    +----+------++     +--+--+------+
    |            |        |         |
    |            |        |         |
    |            |        |         |
+---+---+     +--+-+   +--+--+  +---+--+
|       |     |    |   |     |  |      |
|    4  |     |  5 |   | 6   |  | 7    |
|       |     |    |   |     |  |      |
+----+--+     +----+-+ +-----+  +---+--+
     |               | |            |
     |               | |            |
     |               | |            |
     |            +--+-+-+          |
     +------------+   8  +----------+
                  |      |
                  +------+

*/

func main() {
	edges := []string{
		"0-1 1",
		"1-2 1",
		"1-3 2",
		"2-4 3",
		"2-5 4",
		"3-6 5",
		"3-7 6",
		"4-8 7",
		"5-8 8",
		"6-8 9",
		"7-8 10",
	}

	mygraph := graph.BuildGraph(edges)
	seq := graph.BFS_A(mygraph, 0)
	fmt.Printf("seq: %v\n", seq)
	// fmt.Printf("graph: %+v\n", *mygraph)

	// 构建一个邻接矩阵表示的图，并调用邻接矩阵版本的广度优先算法
	mymgraph := graph.BuildMGraph(edges)
	seq2 := graph.BFS_M(mymgraph, 0)
	fmt.Printf("seq: %v\n", seq2)
}
