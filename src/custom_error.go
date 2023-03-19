package src

// Custom error message.
// Font: https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
type StatusCodeError struct{}

func (s *StatusCodeError) Error() string {
	return "Error: Request status code above or equal to 300."
}
