package errors

// Errors is a slice of errors
type Errors []error

// Error will return a string representing the errors,
// multiple errors will be joined comma separated.
// This is the implementation of the error interface.
func (errs Errors) Error() string {
	if len(errs) > 0 {
		result := errs[0].Error()
		for _, err := range errs[1:] {
			result = result + ", " + err.Error()
		}
		return result
	}
	return ""
}

// Add will append a error to the stack
func (errs *Errors) Add(err error) {
	*errs = append(*errs, err)
}

// HasString checks if there slice has an error who's message matches the string
func (errs Errors) HasString(err string) bool {
	for _, e := range errs {
		if e.Error() == err {
			return true
		}
	}
	return false
}

// HasError checks if the errors exists in the slice
func (errs Errors) HasError(err error) bool {
	for _, e := range errs {
		if e == err {
			return true
		}
	}
	return false
}