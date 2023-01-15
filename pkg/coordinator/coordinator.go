package coordinator

import (
	"github.com/rknizzle/simplemapreduce/pkg/storage"
)

// Coordinator contains the logic of the Coordinator server that keeps track of jobs and
// distributing the tasks of a job to all of the worker nodes
type Coordinator struct {
	workers []*WorkerState
	jobs    []*Job
	storage storage.Storage
}

func (c *Coordinator) startJob() {
	// TODO
}
