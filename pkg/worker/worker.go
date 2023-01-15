package worker

import (
	"github.com/rknizzle/simplemapreduce/pkg/storage"
)

type Worker struct {
	ID          string
	storage     storage.Storage
	coordClient coordinatorClient
}

// coordinatorClient defines how the worker should communicate with the coordinator node
type coordinatorClient interface {
	callCompleteTask()
}

func (w *Worker) completeTask() {
	w.coordClient.callCompleteTask()
}

func (w *Worker) startTask() {
	// TODO:
}

type keyValue struct {
	key   string
	value string
}

// load the map and reduce functions from the application plugin
func loadPlugin(filename string) (func(string, string) []keyValue, func(string, []string) string, error) {
	// TODO
	return nil, nil, nil
}
