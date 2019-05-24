package responses

const (
	// StatusOK defines the OK status
	StatusOK string = "ok"
	// StatusError defines the Error status
	StatusError string = "error"
)

// Status is a struct defining a status response
type Status struct {
	Status string `json:"status"`
}
