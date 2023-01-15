package coordinator

// Task is a single operation (either map or reduce) that makes some progress toward the completion
// of a mapReduce job
type Task struct {
	ID       string
	JobID    string
	WorkerID int
	// map tasks will just have one input but reduce tasks will have many inputs
	Inputs           []string
	Plugin           string
	IsBeingProcessed bool
	IsComplete       bool
	Outputs          []string
}

func NewTask(jobID string, plugin string, inputs []string) (*Task, error) {
	taskID, err := generateUUID()
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:     taskID,
		JobID:  jobID,
		Inputs: inputs,
		Plugin: plugin,
	}, nil
}
