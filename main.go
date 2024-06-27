package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var jobs []Job
	var wg sync.WaitGroup

	file, err := os.OpenFile("websites.txt", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cont := scanner.Text()
		parts := strings.Split(cont, " ")
		var j Job
		j.Link = parts[0]
		j.Delay, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		jobs = append(jobs, j)
	}

	// Schedule all jobs
	for _, j := range jobs {
		wg.Add(1)
		go j.schedule(&wg)
	}

	wg.Wait()
}
