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
