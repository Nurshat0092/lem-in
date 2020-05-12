package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Edge struct {
	from, to  int
	cap, flow int64
}

type Dinic struct {
	N     int
	edges []Edge
	g     [][]int
	d, pt []int
}
type Ant struct {
	ID      int
	pathIdx int
	// pathLen int
	step int
}
type PathAnswer struct {
	lvl  int
	path []int
	ants []Ant
}

var ants [][]Ant
var paths [][]int
var outputPaths []PathAnswer
var answerPaths [][]int
var antsOutput []Ant
var antIDCounter int = 1
var end int
var stepsNum int = 999999999999999

func quickSort(answerPaths [][]int) [][]int {
	if len(answerPaths) < 2 {
		return answerPaths
	}
	left, right := 0, len(answerPaths)-1
	pivot := rand.Int() % len(answerPaths)
	answerPaths[pivot], answerPaths[right] = answerPaths[right], answerPaths[pivot]
	for i := range answerPaths {
		if len(answerPaths[i]) < len(answerPaths[right]) {
			answerPaths[left], answerPaths[i] = answerPaths[i], answerPaths[left]
			left++
		}
	}
	answerPaths[left], answerPaths[right] = answerPaths[right], answerPaths[left]
	quickSort(answerPaths[:left])
	quickSort(answerPaths[left+1:])
	return answerPaths
}

func antsQuickSort(antArr []Ant) []Ant {
	if len(antArr) < 2 {
		return antArr
	}
	left, right := 0, len(antArr)-1
	pivot := rand.Int() % len(antArr)
	antArr[pivot], antArr[right] = antArr[right], antArr[pivot]
	for i := range antArr {
		if antArr[i].ID < antArr[right].ID {
			antArr[left], antArr[i] = antArr[i], antArr[left]
			left++
		}
	}
	antArr[left], antArr[right] = antArr[right], antArr[left]
	antsQuickSort(antArr[:left])
	antsQuickSort(antArr[left+1:])
	return antArr
}

func runAnts() {
	for d.antNum != 0 {
		for i := 0; i < len(answerPaths); {
			outputPaths[i].ants = append(outputPaths[i].ants, Ant{ID: antIDCounter, pathIdx: i})
			outputPaths[i].lvl++
			if i == len(answerPaths)-1 {
				i++
			} else if outputPaths[i].lvl > outputPaths[i+1].lvl {
				i++
			}
			antIDCounter++
			d.antNum--
			if d.antNum == 0 {
				break
			}
		}
	}
	for i := 0; i < stepsNum; i++ {
		strOutput := ""
		for i := 0; i < len(outputPaths); i++ {
			if len(outputPaths[i].ants) > 0 {
				antsOutput = append(antsOutput, outputPaths[i].ants[0])
				outputPaths[i].ants = outputPaths[i].ants[1:]
			}
		}
		antsOutput = antsQuickSort(antsOutput)
		// fmt.Println(antsOutput)
		for i := 0; i < len(antsOutput); i++ {
			strOutput = strOutput + fmt.Sprintf("L%d-%s ", antsOutput[i].ID, d.idxName[outputPaths[antsOutput[i].pathIdx].path[antsOutput[i].step]])
			antsOutput[i].step++
			if antsOutput[i].step == len(outputPaths[antsOutput[i].pathIdx].path) {
				antsOutput = append(antsOutput[:i], antsOutput[i+1:]...)
				i--
			}
		}

		fmt.Println(strOutput)
	}
	// for i := 0; i < stepsNum; i++ {
	// 	strOutput := ""
	// 	for i := 0; i < len(outputPaths); i++ {
	// 		if len(outputPaths[i].ants) > 0 {
	// 			antsOutput = append(antsOutput, outputPaths[i].ants[0])
	// 			outputPaths[i].ants = outputPaths[i].ants[1:]
	// 		}
	// 	}
	// 	for i := 0; i < len(antsOutput); i++ {
	// 		// fmt.Println(antsOutput, antsOutput[i])
	// 		strOutput = strOutput + fmt.Sprintf("L%d-%s ", antsOutput[i].ID, d.idxName[outputPaths[antsOutput[i].pathIdx].path[antsOutput[i].step]])
	// 		antsOutput[i].step++
	// 		// fmt.Println(antsOutput[i])
	// 		// fmt.Println("q", len(outputPaths[antsOutput[i].pathIdx].path), antsOutput[i])
	// 		for {
	// 			fmt.Println("q", len(outputPaths[antsOutput[i].pathIdx].path), antsOutput[i])
	// 			if antsOutput[i].step > len(outputPaths[antsOutput[i].pathIdx].path)-1 {
	// 				antsOutput = antsOutput[1:]

	// 			} else {
	// 				i--
	// 				break
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(strOutput)
}

func main() {
	args := os.Args[1:]
	d := getData(args[0])
	g := NewDinic(d.vertexNumber)
	for _, w := range d.conns {
		g.AddEdge(d.nameIdx[w[0]], d.nameIdx[w[1]], 1)
	}
	end = d.end
	g.MaxFlow(d.start, d.end)
	answerPaths = quickSort(answerPaths)
	for i := range answerPaths {
		answerPaths[i] = append(answerPaths[i], d.end)
		outputPaths = append(outputPaths, PathAnswer{lvl: len(answerPaths[i]), path: answerPaths[i]})
	}
	runAnts()
}

func NewDinic(N int) *Dinic {
	return &Dinic{N, make([]Edge, 0), make([][]int, N), make([]int, N), make([]int, N)}
}

func (g *Dinic) AddEdge(from, to, cap int) {
	if from == d.start && to == d.end {
		for i := 1; i <= d.antNum; i++ {
			fmt.Printf("L%d-%s ", i, d.idxName[d.end])
		}
		log.Fatalln()
	}
	g.edges = append(g.edges, Edge{from, to, int64(cap), int64(0)})
	g.g[from] = append(g.g[from], len(g.edges)-1)
	g.edges = append(g.edges, Edge{to, from, int64(cap), int64(0)})
	g.g[to] = append(g.g[to], len(g.edges)-1)
}

func (g *Dinic) bfs(s, t int) bool {
	cur := []int{s}
	next := []int{}

	inf := g.N + 10
	for i := range g.d {
		g.d[i] = inf
	}
	g.d[s] = 0

	for len(cur) > 0 {
		for _, node := range cur {
			for _, id := range g.g[node] {
				e := &g.edges[id]
				if e.flow < e.cap && g.d[e.to] > g.d[e.from]+1 {
					g.d[e.to] = g.d[e.from] + 1
					next = append(next, e.to)
				}
			}
		}
		cur = next
		next = []int{}
	}

	return g.d[t] != inf
}

func (g *Dinic) dfs(node, T int, flow int64) int64 {
	if node == T || flow == 0 {
		return flow
	}
	for ; g.pt[node] < len(g.g[node]); g.pt[node]++ {
		id := g.g[node][g.pt[node]]
		e := &g.edges[id]
		oe := &g.edges[id^1]
		if g.d[e.from] == g.d[e.to]-1 {
			amt := min64(e.cap-e.flow, flow)
			if pushed := g.dfs(e.to, T, amt); pushed > 0 {
				e.flow += pushed
				oe.flow -= pushed
				return pushed
			}

		}
	}
	return 0
}

func (g *Dinic) MaxFlow(source, sink int) int64 {
	total := int64(0)
	for g.bfs(source, sink) {
		for i := range g.pt {
			g.pt[i] = 0
		}
		for flow := g.dfs(source, sink, int64(1<<60)); flow > 0; {
			total += flow
			paths = make([][]int, 0)

			some_slice := []int{}
			g.get(source, some_slice)
			flow = g.dfs(source, sink, int64(1<<60))
			maxHeight := 0
			sumHeight := 0
			for _, w := range paths {
				sumHeight += len(w)
			}
			if stepsNum > max(maxHeight, maxHeight+(d.antNum-(maxHeight*len(paths)-sumHeight)+len(paths)-1)/len(paths)) {
				stepsNum = max(maxHeight, maxHeight+(d.antNum-(maxHeight*len(paths)-sumHeight)+len(paths)-1)/len(paths))
				answerPaths = make([][]int, 0)
				for _, w := range paths {
					answerPaths = append(answerPaths, w)
				}
			}
		}
	}
	return total
}

func (g *Dinic) get(v int, cur_path []int) {
	for _, i := range g.g[v] {
		to := g.edges[i].to
		if g.edges[i].flow == 1 {
			if to != end {
				cur_path = append(cur_path, to)
				g.get(to, cur_path)
				cur_path = cur_path[:len(cur_path)-1]
			} else {
				paths = append(paths, cur_path)
				return
			}
		}
	}
	return
}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
