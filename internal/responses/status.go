package responses

type statusString string

const (
	StatusOK    statusString = "ok"
	StatusError statusString = "error"
)

type Status struct {
	Status statusString `json:"status"`
}
