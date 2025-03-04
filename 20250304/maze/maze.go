package maze

import "container/list"

type Point struct {
	x int
	y int
}

func solveMaze(maze [][]int, start, end Point) (int, []Point) {
	return bfs(maze, start, end)
}

func IsValidPoint(maze *[][]int, point Point) bool {
	if point.x < 0 || point.y >= len((*maze)[0]) || (*maze)[point.x][point.y] == 0 {
		return false
	}
	return true
}

func IsVisited(visited *[][]bool, point Point) bool {
	return (*visited)[point.x][point.y]
}

func bfs(maze [][]int, start, end Point) (int, []Point) {
	if IsValidPoint(&maze, start) == false || IsValidPoint(&maze, end) == false {
		return 0, nil
	}

	rows := len(maze)
	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	prePoints := make([][]Point, rows)
	for i := 0; i < rows; i++ {
		prePoints[i] = make([]Point, cols)
	}

	queue := list.New()
	queue.PushBack(start)

	visited[start.x][start.y] = true
	for queue.Len() > 0 {
		cur := queue.Front().Value.(Point)
		queue.Remove(queue.Front())
		if cur.x == end.x && cur.y == end.y {
			path := make([]Point, 0)
			for cur.x != start.x || cur.y != start.y {
				path = append(path, cur)
				cur = prePoints[cur.x][cur.y]
			}
			path = append(path, start)
			return len(path), path
		}
		nextPoints := []Point{
			{cur.x - 1, cur.y},
			{cur.x + 1, cur.y},
			{cur.x, cur.y - 1},
			{cur.x, cur.y + 1},
		}
		for _, next := range nextPoints {
			if IsValidPoint(&maze, next) && !IsVisited(&visited, next) {
				queue.PushBack(next)
				visited[next.x][next.y] = true
				prePoints[next.x][next.y] = cur
			}
		}
	}
	return 0, nil
}
