package coordinator

// WorkerRPCclient is a client for the Coordinator to communicate with the worker nodes via RPC
type workerRPCclient struct{}

func (w workerRPCclient) callStartTask() {
	// send RPC to the worker to start working on a specific task
	// call(TODO)
}
