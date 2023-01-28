package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Graph struct {
	Start    string
	End      string
	Edges    map[string][]string
	Vertices []string
}

type Ants struct {
	NumberOfAnts int
}

type Paths struct {
	AllPaths             [][]string
	NonInterceptingPaths [][][]string
	SortedCombinations   map[int][][]string
	BestCombination      [][]string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: missing filename")
		fmt.Println("USAGE: go run main.go example00.txt")
		return
	}

	filePath := os.Args[1]
	runProgram(filePath)
}

func sendAnts(ants Ants, data []string, bestCombination Paths) []string {
	inputData(data)
	queue := assignAnts(ants, bestCombination)
	determineOrder(queue)
	return calculateSteps(queue, bestCombination)
}

func assignAnts(ants Ants, bestCombination Paths) [][]string {
	queue := make([][]string, len(bestCombination.BestCombination))

	for i := 1; i <= ants.NumberOfAnts; i++ {
		ant := strconv.Itoa(i)
		minSteps := len(bestCombination.BestCombination[0]) + len(queue[0])
		minIndex := 0

		for j, path := range bestCombination.BestCombination {
			steps := len(path) + len(queue[j])
			if steps < minSteps {
				minSteps = steps
				minIndex = j
			}
		}
		queue[minIndex] = append(queue[minIndex], ant)
	}
	return queue
}

func determineOrder(queue [][]string) []int {
	order := []int{}
	longest := len(queue[0])
	for i := 0; i < len(queue); i++ {
		if len(queue[i]) > longest {
			longest = len(queue[i])
		}
	}

	for j := 0; j < longest; j++ {
		for i := 0; i < len(queue); i++ {
			if j < len(queue[i]) {
				x, _ := strconv.Atoi(queue[i][j])
				order = append(order, x)
			}
		}
	}
	return order
}

func calculateSteps(queue [][]string, bestCombination Paths) []string {
	container := make([][][]string, len(queue))
	for i, path := range queue {
		for _, ant := range path {
			adder := []string{}
			for _, vertex := range bestCombination.BestCombination[i] {
				str := "L" + ant + "-" + vertex
				adder = append(adder, str)
			}
			container[i] = append(container[i], adder)
		}
	}

	finalMoves := []string{}
	for _, paths := range container {
		for j, moves := range paths {
			for k, vertex := range moves {
				if j+k > len(finalMoves)-1 {
					finalMoves = append(finalMoves, vertex+" ")
				} else {
					finalMoves[j+k] = finalMoves[j+k] + vertex + " "
				}
			}
		}
	}
	return finalMoves
}

func bestCombination(ants Ants, pathGroups Paths) Paths {
	minLevel := math.MaxInt32
	var bestCombination [][]string

	for _, combination := range pathGroups.SortedCombinations {
		spaceInPath := 0
		totalPathLength := 0
		levelOfAnts := 0
		longestPath := len(combination[len(combination)-1])
		for i := 1; i < len(combination); i++ {
			spaceInPath = spaceInPath + longestPath - len(combination[i])
			totalPathLength = totalPathLength + len(combination[i])
		}
		levelOfAnts = (ants.NumberOfAnts-spaceInPath)/len(combination) + longestPath

		if levelOfAnts < minLevel {
			minLevel = levelOfAnts
			bestCombination = combination
		}
	}

	return Paths{BestCombination: bestCombination}
}

func sortCombinations(combinations Paths) Paths {
	pathGroups := make(map[int][][]string)
	for _, combination := range combinations.NonInterceptingPaths {
		category := len(combination)
		currentCombLength := getCombinationLength(combination)
		if _, ok := pathGroups[category]; ok {
			valueInMap := pathGroups[category]
			if currentCombLength < getCombinationLength(valueInMap) {
				pathGroups[category] = combination
			}
		} else {
			pathGroups[category] = combination
		}
	}
	return Paths{SortedCombinations: pathGroups}
}

func getCombinationLength(combination [][]string) int {
	length := 0
	for _, path := range combination {
		length = length + len(path)
	}
	return length
}

func findNonInterceptingPaths(paths Paths) Paths {
	sortByLength(paths)
	var result [][][]string
	for i, path := range paths.AllPaths {
		var nonInterceptingPaths [][]string
		nonInterceptingPaths = append(nonInterceptingPaths, path)
		result = append(result, nonInterceptingPaths)
		for j := i + 1; j < len(paths.AllPaths); j++ {
			if !hasInterception(nonInterceptingPaths, paths.AllPaths[j]) {
				nonInterceptingPaths = append(nonInterceptingPaths, paths.AllPaths[j])
				result = append(result, nonInterceptingPaths)
			}
		}
	}
	return Paths{NonInterceptingPaths: result}
}

func hasInterception(nonInterceptingPaths [][]string, path []string) bool {
	for _, paths := range nonInterceptingPaths {
		for _, vertex1 := range paths[:len(paths)-1] {
			for _, vertex2 := range path[:len(path)-1] {
				if vertex1 == vertex2 {
					return true
				}
			}
		}
	}
	return false
}

func sortByLength(paths Paths) {
	for i := 0; i < len(paths.AllPaths)-1; i++ {
		for j := 0; j < len(paths.AllPaths)-i-1; j++ {
			if len(paths.AllPaths[j]) > len(paths.AllPaths[j+1]) {
				paths.AllPaths[j+1], paths.AllPaths[j] = paths.AllPaths[j], paths.AllPaths[j+1]
			}
		}
	}
}

func findPaths(graph Graph) Paths {
	if _, ok := graph.Edges[graph.Start]; !ok {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}
	if _, ok := graph.Edges[graph.End]; !ok {
		fmt.Println("ERROR: invalid data format, no end room found")
		os.Exit(0)
	}
	if graph.Start == "" {
		fmt.Println("ERROR: invalid data format, no start room specified")
		os.Exit(0)
	}
	if graph.End == "" {
		fmt.Println("ERROR: invalid data format, no end room specified")
		os.Exit(0)
	}

	var paths [][]string
	visited := make(map[string]bool)
	var path []string

	var findPaths func(string, string, map[string][]string, map[string]bool, []string)
	findPaths = func(start, end string, edges map[string][]string, visited map[string]bool, path []string) {
		visited[start] = true
		path = append(path, start)

		if start == end {
			temp := make([]string, len(path[1:]))
			copy(temp, path[1:])
			paths = append(paths, temp)
		} else {
			for _, vertex := range edges[start] {
				if !visited[vertex] {
					findPaths(vertex, end, edges, visited, path)
				}
			}
		}

		path = path[:len(path)-1]
		visited[start] = false
	}
	findPaths(graph.Start, graph.End, graph.Edges, visited, path)
	return Paths{AllPaths: paths}
}

func parseData(data []string) (Graph, Ants) {
	var graph Graph
	var ants Ants

	graph.Edges = make(map[string][]string)

	numberOfAnts, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Printf("Error converting number of ants to integer: %v", err)
		os.Exit(0)
	}
	ants.NumberOfAnts = numberOfAnts
	if ants.NumberOfAnts < 1 {
		fmt.Printf("ERROR: invalid data format, invalid number of Ants\n")
		os.Exit(0)
	}
	if ants.NumberOfAnts > 1000_000_000 {
		fmt.Printf("ERROR: invalid data format, invalid number of Ants\n")
		os.Exit(0)
	}

	var i int
	for i = 1; i < len(data); i++ {
		if data[i] == "##start" {
			fields := strings.Fields(data[i+1])
			start := fields[0]
			graph.Start = string(start)
		} else if data[i] == "##end" {
			fields := strings.Fields(data[i+1])
			end := fields[0]
			graph.End = string(end)
		} else if strings.Contains(data[i], "-") {
			edge := strings.Split(data[i], "-")
			from := edge[0]
			to := edge[1]
			if from == "" || to == "" {
				fmt.Printf("Invalid edge format")
				os.Exit(0)
			}
			graph.Edges[from] = append(graph.Edges[from], to)
			graph.Edges[to] = append(graph.Edges[to], from)
		} else if strings.Contains(data[i], " ") {
			vertex := strings.Split(data[i], " ")
			v := vertex[0]
			if strings.HasPrefix(v, "L") || strings.HasPrefix(v, "#") {
				fmt.Printf("Invalid vertex name: %s", v)
				os.Exit(0)
			}
			graph.Vertices = append(graph.Vertices, v)
		}
	}
	return graph, ants
}

func inputData(data []string) {
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println()
}

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	if len(fileLines) == 0 {
		fmt.Println("The file is empty.")
		os.Exit(0)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
		os.Exit(0)
	}
	return fileLines
}

func runProgram(filePath string) {
	data := readFile("examples/" + filePath)
	graph, ants := parseData(data)

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	start := time.Now()

	paths := findPaths(graph)
	nonInterceptingPaths := findNonInterceptingPaths(paths)
	sortedCombinations := sortCombinations(nonInterceptingPaths)
	bestCombination := bestCombination(ants, sortedCombinations)
	sendAnts := sendAnts(ants, data, bestCombination)

	turns := 0
	for _, v := range sendAnts {
		fmt.Println(v)
		turns++
	}
	fmt.Println()

	fmt.Printf("Found %v paths in %v.\n", len(paths.AllPaths), time.Since(start))
	fmt.Printf("Used quickest path possible with %v turns.\n", turns)
}
