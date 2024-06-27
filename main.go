package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Job struct {
	Link  string
	Delay int
}

func execJob(j Job) {
	resp, err := http.Get(j.Link)
	var status string
	if err != nil {
		status = "Get Failed"
		fmt.Print(err)
	} else {
		status = resp.Status
	}

	fmt.Printf("HIT %s - %s\n", j.Link, status)
}

func scheduleJob(j Job, wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(j.Delay))
	defer ticker.Stop()
	defer wg.Done()

	execJob(j)

	for range ticker.C {
		execJob(j)
	}
}

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
		go scheduleJob(j, &wg)
	}

	wg.Wait()
}
