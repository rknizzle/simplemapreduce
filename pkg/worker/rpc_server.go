package worker

import "fmt"

// worker defines the logic that this API needs to use
type worker interface {
	startTask(string, string, string, []string, bool, bool)
}

// RPCserver is a server for the Worker that listens for RPC calls from the coordinator node
type RPCserver struct {
	w worker
}

type StartTaskArgs struct {
	JobID        string
	TaskID       string
	Plugin       string
	Inputs       []string
	IsMapTask    bool
	IsReduceTask bool
}
type StartTaskReply struct{}

func (server RPCserver) HandleStartTask(args *StartTaskArgs, reply *StartTaskReply) {
	fmt.Println("DEBUG: worker server recieved an RPC call to start a task")

	server.w.startTask(
		args.JobID,
		args.TaskID,
		args.Plugin,
		args.Inputs,
		args.IsMapTask,
		args.IsReduceTask,
	)
}

func (server RPCserver) startServer() error {
	// TODO
	return nil
}
