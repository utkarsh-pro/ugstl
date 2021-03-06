package heap

import (
	"errors"
	"fmt"
)

var (
	ErrIndexOutOfBound   = errors.New("index out of bound")
	ErrInvalidHeapConfig = errors.New("heap can either be of type MIN or MAX")
	ErrInvalidElement    = errors.New("invalid element")
	ErrHeapUnderflow     = errors.New("heap underflow")
)

func GetErrIndexOutOfBound(idx, len int) error {
	return fmt.Errorf("%w: index() = %v, len() = %v", ErrIndexOutOfBound, idx, len)
}
