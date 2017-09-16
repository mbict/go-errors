package errors

import (
	"fmt"
	"sort"
)

// ErrorHash holds a slices with errors based on a string key
type ErrorHash map[string]Errors

// NewErrorHash constructor
func NewErrorHash() ErrorHash {
	return make(ErrorHash, 0)
}

// Error returns a string with all errors with the keys sorted alphabetical
// Error is the implementation of the error interface.
func (eh ErrorHash) Error() string {
	//sort by key
	index := make([]string, 0, len(eh))
	for key, _ := range eh {
		index = append(index, key)
	}
	sort.Strings(index)

	result := ""
	for _, name := range index {
		if len(eh[name]) > 0 {
			if len(result) > 0 {
				result = fmt.Sprintf("%s, %s:[%s]", result, name, eh[name].Error())
			} else {
				result = fmt.Sprintf("%s:[%s]", name, eh[name].Error())
			}
		}
	}
	return result
}

// Add will append a error for the specified id/key
func (eh ErrorHash) Add(id string, err error) {
	eh[id] = append(eh[id], err)
}

// Has is a helper function to check if there is a error in the ErrorMap for the
// corresponding id/key name. Handy for use with the template funcion map
func (eh ErrorHash) Has(id string) bool {
	errs, ok := eh[id]
	return ok && len(errs) > 0
}

// HasString checks if there is a specific error string message in the ErrorHash
// for the corresponding id/key
func (eh ErrorHash) HasString(id string, err string) bool {
	errs, ok := eh[id]
	if !ok {
		return false
	}

	for _, e := range errs {
		if e.Error() == err {
			return true
		}
	}
	return false
}

// HasError checks if there is a specific error in the ErrorHash
// for the corresponding id/key
func (eh ErrorHash) HasError(id string, err error) bool {
	errs, ok := eh[id]
	if !ok {
		return false
	}

	for _, e := range errs {
		if e == err {
			return true
		}
	}
	return false
}
