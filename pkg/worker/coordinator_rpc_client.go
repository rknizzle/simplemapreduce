package worker

// coordinatorRPCclient is a client for a Worker to communicate with the coordinator node via RPC
type coordinatorRPCclient struct{}

func (c coordinatorRPCclient) callCompleteTask() {
	// send RPC with the results
	// call(TODO)
}
