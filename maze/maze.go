package main

import (
	"fmt"
	"os"
)

type point struct {
	i int
	j int
}

var direction = [] point {
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(maze [][] int) (int, bool) {
	if p.i < 0 || p.i >= len(maze) {
		return -1, false
	}
	if p.j < 0 || p.j >= len(maze[0]) {
		return -1, false
	}
	return maze[p.i][p.j], true
}

func initialMaze(filePath string) [][] int{
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func walk(maze [][] int, start point, end point) [][] int {
	step := make([][] int, len(maze))
	for i := range maze {
		step[i] = make([] int, len(maze[i]))
	}

	queue := [] point {
		start,
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p == end {
			break
		}
		for _, d := range direction {
			r := p.add(d)
			if r == start {
				continue
			}
			if val, ok := r.at(maze); !ok || val != 0 {
				continue
			}
			if val, ok := r.at(step); !ok || val != 0  {
				continue
			}
			curStep, _ := p.at(step)
			step[r.i][r.j] = curStep + 1
			queue = append(queue, r)
		}

	}

	return step
}

func main() {
	maze := initialMaze("maze/maze.in")
	step := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for i := range step {
		for j := range step[i] {
			fmt.Printf("%3d", step[i][j])
		}
		fmt.Println()
	}
	//fmt.Println(step)
}
