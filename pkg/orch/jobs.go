package orch

const (
	jobs = "/orchestrator/v1/jobs"
	job  = "/orchestrator/v1/job/{job-id}"
)

// Jobs lists all of the jobs known to the orchestrator (GET /jobs)
func (c *Client) Jobs() (*Jobs, error) {
	payload := &Jobs{}
	r, err := c.resty.R().SetResult(&payload).Get(jobs)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// Job lists all details of a given job (GET /job)
func (c *Client) Job(jobID string) (*Job, error) {
	payload := &Job{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(job)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// Jobs contains data about all jobs
type Jobs struct {
	Items      []Job      `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// Job contains data about a single job
type Job struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Command     string      `json:"command"`
	Options     Options     `json:"options"`
	NodeCount   int         `json:"node_count"`
	Owner       Owner       `json:"owner"`
	Description string      `json:"description"`
	Timestamp   string      `json:"timestamp"`
	Environment Environment `json:"environment"`
	Status      []Status    `json:"status"`
	Nodes       Nodes       `json:"nodes"`
	Report      Report      `json:"report"`
}

// Environment in the current job
type Environment struct {
	Name string `json:"name"`
}

// Status of the current job
type Status struct {
	State     string `json:"state"`
	EnterTime string `json:"enter_time"`
	ExitTime  string `json:"exit_time"`
}

// Nodes in the current job
type Nodes struct {
	ID string `json:"id"`
}

// Events in the current job
type Events struct {
	ID string `json:"id"`
}

// Report for the current job
type Report struct {
	ID string `json:"id"`
}

// NodeStates for the current job
type NodeStates struct {
	Finished int `json:"finished"`
	Errored  int `json:"errored"`
	Failed   int `json:"failed"`
	Running  int `json:"running"`
}

// Options for the current job
type Options struct {
	Concurrency        interface{} `json:"concurrency"`
	Noop               bool        `json:"noop"`
	Trace              bool        `json:"trace"`
	Debug              bool        `json:"debug"`
	Scope              Scope       `json:"scope"`
	EnforceEnvironment bool        `json:"enforce_environment"`
	Environment        string      `json:"environment"`
	Evaltrace          bool        `json:"evaltrace"`
	Target             interface{} `json:"target"`
	Description        string      `json:"description"`
}
