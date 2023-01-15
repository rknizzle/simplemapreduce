package worker

// worker defines the logic that this API needs to use
type worker interface {
	startTask()
}

// RPCserver is a server for the Worker that listens for RPC calls from the coordinator node
type RPCserver struct {
	w worker
}

type StartTaskArgs struct{}
type StartTaskReply struct{}

func (server RPCserver) HandleStartTask(args *StartTaskArgs, reply *StartTaskReply) {
	server.w.startTask()
}

func (server RPCserver) startServer() error {
	// TODO
	return nil
}
