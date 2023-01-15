package coordinator

type workerClient interface {
	callStartTask(task *Task) error
}

// workerState stores the state of a single worker node from the perspective of the Coordinator
type WorkerState struct {
	isHealthy   bool
	isBusy      bool
	currentTask *Task

	client workerClient
}

func (w WorkerState) sendTaskToWorker(task *Task) error {
	err := w.client.callStartTask(task)
	if err != nil {
		return err
	}

	return nil
}
