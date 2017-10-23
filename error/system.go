package error

import(
	"fmt"
)

type SystemError struct {
	Message string
	LogMessage string
}

func (err *SystemError) Error() string {
	return fmt.Sprintf("%s", err.Message)
}
