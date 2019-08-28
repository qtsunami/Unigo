package main

import (
	"Unigo/tools/go_worker_pool/pool"
	"Unigo/tools/go_worker_pool/work"
	"log"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}