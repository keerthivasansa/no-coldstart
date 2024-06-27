package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	Link  string
	Delay int
}

func (j Job) exec() {
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

func (j Job) schedule(wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(j.Delay))
	defer ticker.Stop()
	defer wg.Done()

	j.exec()

	for range ticker.C {
		j.exec()
	}
}
