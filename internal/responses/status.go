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

// Healthz is a struct defining the status of dependent services
type Healthz struct {
	Status            string `json:"status"`
	GRPCStarterStatus string `json:"grpcStarterStatus"`
}
