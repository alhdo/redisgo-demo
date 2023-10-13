package core

// Define constants for job statuses
const (
	Pending    = "pending"
	InProgress = "in-progress"
	Completed  = "completed"
	Failed     = "failed"
)

type Job struct {
	ID     string                 `json:"id""`
	Task   string                 `json:"task"`
	Config map[string]interface{} `json:"config"`
}

type JobResponse struct {
	Key    string `json:"key"`
	Status string `json:"status"`
}
