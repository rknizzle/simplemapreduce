package worker

import (
	"time"

	"github.com/rknizzle/simplemapreduce/pkg/storage"
)

type Worker struct {
	ID          string
	storage     storage.Storage
	coordClient coordinatorClient
}

// coordinatorClient defines how the worker should communicate with the coordinator node
type coordinatorClient interface {
	callCompleteTask(output string, jobID string, taskID string) error
}

func (w *Worker) completeTask(output string, jobID string, taskID string) error {
	err := w.coordClient.callCompleteTask(output, jobID, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) startTask(jobID string, taskID string, plugin string, inputs []string, isMapTask bool, isReduceTask bool) {
	// download input(s)
	// download plugin
	// if reduce task: sort & combine inputs
	// call plugin code on input
	// write the output to a file
	// w.storage.Write()
	time.Sleep(1 * time.Second)
	// TODO: does the output all get combined into 1 output file at the end? How does that combining
	// happen? Append to the result?
	output := "TODO"
	w.completeTask(output, jobID, taskID)
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
