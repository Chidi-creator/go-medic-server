package utils

type Specialty string

const (
	PEDIATRICIAN         Specialty = "pediatrician"
	OPTOMETRIST          Specialty = "optometrist"
	DENTIST              Specialty = "dentist"
	SURGEON              Specialty = "surgeon"
	CARDIOLOGIST         Specialty = "cardiologist"
	DERMATOLOGIST        Specialty = "dermatologist"
	GENERAL_PRACTITIONER Specialty = "general_practitioner"
)

var ValidSpecialties = []Specialty{PEDIATRICIAN, OPTOMETRIST, DENTIST, SURGEON, CARDIOLOGIST, DERMATOLOGIST, GENERAL_PRACTITIONER}

type InviteStatus string

const (
	PENDING  InviteStatus = "pending"
	ACCEPTED InviteStatus = "accepted"
	REJECTED InviteStatus = "rejected"
)

type Status string

const (
	WAITING Status = "waiting"
	ONGOING Status = "ongoing"
	DONE    Status = "done"
)

type Roles string

const (
	CUSTOMER Roles = "customer"
	DOCTOR   Roles = "doctor"
	HOSPITAL Roles = "hospital_owner"
	ADMIN    Roles = "admin"
)

var ValidRoles = []Roles{CUSTOMER, DOCTOR, HOSPITAL, ADMIN}

// structuring the response manager
type ApiResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}


type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1"`
}
