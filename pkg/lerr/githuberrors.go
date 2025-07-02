package lerr

import (
	"fmt"
	"github.com/pkg/errors"
)

func NewWrapError(err, railErr error, msg string, req ...any) error {
	if len(req) == 0 {
		return errors.Wrapf(err, msg+" err is %v", railErr)
	}
	s := msg + " err is " + railErr.Error()
	for i := 0; i < len(req); i++ {
		s = s + " req" + fmt.Sprintf("%v", req[i])
	}
	return errors.Wrapf(err, s)
}
