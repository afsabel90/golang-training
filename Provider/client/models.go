package air

type Application struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	State                    string `json:"state"`
	Owner                    string `json:"owner"`
	SubmissionDate           string `json:"submissionDate"`
	ConfigurationItemId      string `json:"configurationItemId"`
	Identifier               string `json:"identifier"`
	PendingReviewConnections int64  `json:"pendingReviewConnections"`
	Category                 string `json:"category"`
}
