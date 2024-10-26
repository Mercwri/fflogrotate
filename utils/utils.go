package utils

type FFLogError struct{}

func (e *FFLogError) Error() string {
	return "Misc FFLogError, Goodluck!"
}
