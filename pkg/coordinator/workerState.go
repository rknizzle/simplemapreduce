package coordinator

type workerClient interface {
	callStartTask()
}

// workerState stores the state of a single worker node from the perspective of the Coordinator
type WorkerState struct {
	isHealthy   bool
	isBusy      bool
	url         string
	currentTask *Task

	client workerClient
}
