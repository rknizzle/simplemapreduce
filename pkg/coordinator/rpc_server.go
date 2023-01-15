package coordinator

// coordinatorWorkerFacing defines the methods that the Coordinator needs to implement for the
// worker-facing RPC API
type coordinatorWorkerFacing interface {
	completeTask()
	// registerNewWorker()
}

type RPCserver struct {
	c coordinatorWorkerFacing
}

type CompleteTaskArgs struct{}
type CompleteTaskReply struct{}

func (server RPCserver) handleCompleteTask(args *CompleteTaskArgs, reply *CompleteTaskReply) {
	server.c.completeTask()
}
