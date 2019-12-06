package cerr

import "time"

type CustomErr struct {
	when time.Time
	err  error
}

func (e *CustomErr) Error() string {
	return e.when.Format(time.UnixDate) + ": " + e.err.Error()
}
func NewCustomErr(err error) *CustomErr {
	return &CustomErr{
		when: time.Now(),
		err:  err,
	}
}
