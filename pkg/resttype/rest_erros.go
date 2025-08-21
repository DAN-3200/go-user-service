package resttype

import "strings"

type ErrorRes struct {
	Status    int        `json:"status"`
	ErrType   string     `json:"err_type"`
	Path      string     `json:"path"`
	Message   string     `json:"message"`
	Details   ErrorsList `json:"details"`
	TimeStamp string     `json:"timeStamp"`
}

func (it ErrorRes) Error() string {
	return it.ErrType
}

type ErrorsList []string

func (it ErrorsList) Error() string {
	return "Errors: " + strings.Join(it, "; ")
}
