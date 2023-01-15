package coordinator

// job represents a MapReduce problem subimitted by a user and it consists of many tasks
type Job struct {
	ID string

	plugin string

	isMapPhaseComplete bool

	mapTasks    []*Task
	reduceTasks []*Task

	// number of map tasks is based on number of inputs or based on size of input and config for how
	// small to break up the input
	numMapTasks          int
	numMapTasksCompleted int

	// num reduce tasks is specified by numReducers. Theyre redundant
	numReduceTasks          int
	numReduceTasksCompleted int

	// this is specified by the user when creating the job. This basically says how many output files
	// each mapper should write in this job
	numReducers int
}

func NewJob(plugin string, inputs []string) (*Job, error) {
	jobID, err := generateUUID()
	if err != nil {
		return nil, err
	}

	// turn each input into a map task
	tasks := make([]*Task, len(inputs))
	for i, input := range inputs {
		t, err := NewTask(jobID, plugin, []string{input})
		if err != nil {
			return nil, err
		}

		tasks[i] = t
	}

	return &Job{
		ID:     jobID,
		plugin: plugin,

		isMapPhaseComplete: false,
		mapTasks:           tasks,
		numMapTasks:        len(tasks),

		reduceTasks: nil,
	}, nil
}

func (j *Job) HasTaskAvailable() bool {
	if !j.isMapPhaseComplete {
		for _, t := range j.mapTasks {
			if t.IsBeingProcessed == false && t.IsComplete == false {
				return true
			}
		}
	} else {
		for _, t := range j.reduceTasks {
			if t.IsBeingProcessed == false && t.IsComplete == false {
				return true
			}
		}
	}

	return false
}

func (j *Job) AvailableTask() *Task {
	if !j.isMapPhaseComplete {
		for _, t := range j.mapTasks {
			if t.IsBeingProcessed == false && t.IsComplete == false {
				return t
			}
		}
	} else {
		for _, t := range j.reduceTasks {
			if t.IsBeingProcessed == false && t.IsComplete == false {
				return t
			}
		}
	}

	return nil
}
