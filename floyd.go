package graph

func Floyd(graph *Mgraph) (*[][]int, *[][]int) {
	n := graph.N
	dist := initializeArray2d(makeArray2d(n, n), INF)
	path := initializeArray2d(makeArray2d(n, n), -1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			(*dist)[i][j] = graph.Edges[i][j]
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if (*dist)[i][k]+(*dist)[k][j] < (*dist)[i][j] {
					(*dist)[i][j] = (*dist)[i][k] + (*dist)[k][j]
					(*path)[i][j] = k
				}
			}
		}
	}

	return dist, path
}
