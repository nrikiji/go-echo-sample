package error

import(
	"fmt"
)

type BusinessError struct {
	Message string
	LogMessage string
}

func (err *BusinessError) Error() string {
	return fmt.Sprintf("%s", err.Message)
}
