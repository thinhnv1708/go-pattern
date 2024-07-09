package lib

import (
	"fmt"
	"time"
)

type Job interface {
	Process()
}

type Worker struct {
	WorkerId   int
	Done       chan bool
	JobRunning chan Job
}

func NewWorker(workerID int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerId:   workerID,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run() {
	fmt.Println("Run worker id", w.WorkerId)
	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				job.Process()
			case <-w.Done:
				fmt.Println("Stop worker id", w.WorkerId)
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.Done <- true
}

type JobQueue struct {
	Workers    []*Worker
	JobRunning chan Job
	Done       chan bool
}

func (jq *JobQueue) PushJob(job Job) {
	jq.JobRunning <- job
}

func (jq *JobQueue) Start() {
	fmt.Println("Start JobQueue")
	go func() {
		for i := 0; i < len(jq.Workers); i++ {
			jq.Workers[i].Run()
		}
	}()

	go func() {

		if <-jq.Done {
			for i := 0; i < len(jq.Workers); i++ {
				jq.Workers[i].StopWorker()
			}
		}

	}()
}

func (jq *JobQueue) Stop() {
	jq.Done <- true
}

func NewJobQueue(numOfWorkers int) *JobQueue {
	workers := make([]*Worker, numOfWorkers)
	jobRunning := make(chan Job)

	for i := 0; i < numOfWorkers; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}

	return &JobQueue{
		Workers:    workers,
		JobRunning: jobRunning,
		Done:       make(chan bool),
	}
}

type Sender struct {
	Email string
}

func (s Sender) Process() {
	fmt.Println(s.Email)
}

func JobQueueMain() {

	emails := []string{"thinh", "anh", "lo", "cac", "thinh", "anh", "lo", "cac"}

	jq := NewJobQueue(1)
	jq.Start()

	for _, email := range emails {
		jq.PushJob(Sender{Email: email})
	}

	time.AfterFunc(2*time.Second, func() {
		jq.Stop()

	})

	time.Sleep(time.Second * 3)
}
