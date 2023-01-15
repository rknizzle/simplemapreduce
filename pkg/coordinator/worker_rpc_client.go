package coordinator

import (
	"net/rpc"

	"github.com/rknizzle/simplemapreduce/pkg/worker"
)

// WorkerRPCclient is a client for the Coordinator to communicate with the worker nodes via RPC
type workerRPCclient struct {
	// endpoint is usually either a unix socket when running locally or a URL when calling across the network
	endpoint     string
	isUnixSocket bool
	isTCP        bool
}

func (w workerRPCclient) callStartTask(task *Task) error {
	args := &worker.StartTaskArgs{
		JobID:        task.JobID,
		TaskID:       task.ID,
		Plugin:       task.Plugin,
		Inputs:       task.Inputs,
		IsMapTask:    task.IsMapTask,
		IsReduceTask: task.IsReduceTask,
	}
	reply := &worker.StartTaskReply{}

	err := call("worker.RPCserver.HandleStartTask", w.endpoint, args, reply)
	if err != nil {
		return err
	}

	return nil
}

// send RPC
func call(rpcname string, endpoint string, args interface{}, reply interface{}) error {
	// TODO: have the 'unix' be configurable to 'tcp' when running over the network
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	c, err := rpc.DialHTTP("unix", endpoint)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err != nil {
		return err
	}

	return nil
}
