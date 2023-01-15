package coordinator

import (
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
