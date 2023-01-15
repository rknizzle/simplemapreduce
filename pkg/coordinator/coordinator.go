package coordinator

import (
	"time"

	"github.com/rknizzle/simplemapreduce/pkg/storage"
)

// Coordinator contains the logic of the Coordinator server that keeps track of jobs and
// distributing the tasks of a job to all of the worker nodes
type Coordinator struct {
	Workers []*WorkerState
	Jobs    []*Job
	Storage storage.Storage
}

func NewCoordinator(s storage.Storage) Coordinator {
	return Coordinator{Storage: s}
}

func (c *Coordinator) StartJob(plugin string, inputs []string) (*Job, error) {
	job, err := NewJob(plugin, inputs)
	if err != nil {
		return nil, err
	}

	c.Jobs = append(c.Jobs, job)
	return job, nil
}

func (c Coordinator) ProcessJobLoop() {
	for {

		for _, job := range c.Jobs {

			for job.HasTaskAvailable() {

				worker := c.workerAvailable()
				if worker == nil {
					break
				}

				task := job.AvailableTask()

				worker.sendTaskToWorker(task)

				task.IsBeingProcessed = true
			}
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (c *Coordinator) workerAvailable() *WorkerState {
	for _, workerState := range c.Workers {
		if workerState.isHealthy && !workerState.isBusy {
			return workerState
		}
	}

	return nil
}
