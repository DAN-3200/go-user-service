package mytypes

import "strings"

type ErrorRes struct {
	Status int
	Error  error
}

type ErrorsList []string

func (it ErrorsList) Error() string {
	return "Errors: " + strings.Join(it, "; ")
}
