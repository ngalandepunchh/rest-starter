package responses

type statusString string

const (
	// StatusOK defines the OK status
	StatusOK statusString = "ok"
	// StatusError defines the Error status
	StatusError statusString = "error"
)

// Status is a struct defining a status response
type Status struct {
	Status statusString `json:"status"`
}
