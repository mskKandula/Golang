package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	maxWorkers int
	taskQueue  chan Task
	wg         sync.WaitGroup
}

type Task struct {
	ID int
	// Additional task data
}

func NewWorkerPool(maxWorkers, maxTasks int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		taskQueue:  make(chan Task, maxTasks),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for task := range wp.taskQueue {
				fmt.Printf("Worker %d: processing task %d\n", workerID, task.ID)
				time.Sleep(1 * time.Second) // Simulating task processing
			}
		}(i)
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	select {
	case wp.taskQueue <- task:
		// Task added to the queue
	default:
		// All workers are busy, create a new goroutine to handle the task
		go func() {
			fmt.Printf("Creating new goroutine to handle task %d\n", task.ID)
			fmt.Printf("Goroutine: processing task %d\n", task.ID)
			time.Sleep(1 * time.Second) // Simulating task processing
		}()
	}
}

func (wp *WorkerPool) Wait() {
	close(wp.taskQueue)
	wp.wg.Wait()
}

func main() {
	maxWorkers := 3
	maxTasks := 5

	workerPool := NewWorkerPool(maxWorkers, maxTasks)
	workerPool.Start()

	for i := 1; i <= 10; i++ {
		task := Task{ID: i}
		workerPool.AddTask(task)
	}

	workerPool.Wait()
}
