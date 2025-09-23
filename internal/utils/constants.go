package utils

type Specialty string

const (
	PEDIATRICIAN Specialty = "pediatrician"
	OPTOMETRIST  Specialty = "optometrist"
	DENTIST      Specialty = "dentist"
)

type Status string

const (
	WAITING Status = "waiting"
	ONGOING Status = "ongoing"
	DONE    Status = "done"
)

//structuring the response manager
type ApiResponse struct {
	Success bool      `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
