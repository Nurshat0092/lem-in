package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type data struct {
	antNum       int
	vertexNumber int
	nameIdx      map[string]int
	idxName      map[int]string
	conns        [][]string
	vertexes     []string
	start        int
	end          int
}

type answer struct {
	id, cur_p, cur_i, step int
}

var d data
var counter = 0

func getData(fileName string) data {
	d.nameIdx = make(map[string]int)
	d.idxName = make(map[int]string)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanLines)
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	readFileAndUpdate(arr)
	return d
}

func readFileAndUpdate(arr []string) {
	d.antNum, _ = strconv.Atoi(arr[0])
	if d.antNum < 1 {
		log.Fatalln("ERROR: invalid data format.")
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == "##start" || arr[i] == "##end" {
			isStart, isEnd := true, false
			if arr[i] == "##end" {
				isStart, isEnd = false, true
			}
			i++
			for {
				if i == len(arr)-1 || arr[i] == "" {
					log.Fatalln("ERROR: invalid data format.")
				}
				if arr[i][0] == '#' {
					i++
				} else {
					break
				}
			}
			vertexUpdate(arr[i], isStart, isEnd)
		} else if arr[i] == "" {
			log.Fatalln("ERROR: invalid data format.")
		} else if arr[i][0] == '#' {
			continue
		} else if len(strings.Fields(arr[i])) == 1 {
			edgeUpdate(arr[i])
		} else {
			vertexUpdate(arr[i], false, false)
		}

	}
}

func vertexUpdate(str string, isStart, isEnd bool) {
	d.vertexNumber++
	vertexInfo := strings.Fields(str)
	if len(vertexInfo) != 3 {
		log.Fatalln("ERROR: invalid data format.")
	} else if d.nameIdx[vertexInfo[0]] != 0 && d.idxName[d.nameIdx[vertexInfo[0]]] != "" {
		log.Fatalln("ERROR: invalid data format.")
	}
	d.vertexes = append(d.vertexes, vertexInfo[0])
	d.nameIdx[vertexInfo[0]] = counter
	d.idxName[counter] = vertexInfo[0]
	if isStart {
		d.start = counter
	} else if isEnd {
		d.end = counter
	}
	counter++
}

func edgeUpdate(str string) {
	conn := strings.Split(str, "-")
	if len(conn) != 2 {
		log.Fatalln("ERROR: invalid data format.")
	}
	d.conns = append(d.conns, conn)
}
