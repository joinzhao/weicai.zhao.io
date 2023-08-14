package errorx

type Error interface {
	Code() int
	Msg() string
	Error() string
	Errors() []error
	WithCode(code int) Error
	WithMsg(msg string) Error
	WithError(err error) Error
}

type errorX struct {
	code int
	msg  string
	errs []error
}

func (x *errorX) Code() int   { return x.code }
func (x *errorX) Msg() string { return x.msg }
func (x *errorX) Error() string {
	if x.errs == nil || len(x.errs) == 0 {
		return ""
	}
	return x.errs[0].Error()
}
func (x *errorX) Errors() []error { return x.errs }

func (x *errorX) WithCode(code int) Error {
	err := x.copy()
	err.code = code
	return err
}
func (x *errorX) WithMsg(msg string) Error {
	err := x.copy()
	err.msg = msg
	return err
}
func (x *errorX) WithError(errX error) Error {
	err := x.copy()
	if err.errs == nil {
		err.errs = make([]error, 0)
	}
	err.errs = append(err.errs, errX)
	return err
}
func (x *errorX) copy() *errorX {
	return &errorX{
		code: x.code,
		msg:  x.msg,
		errs: x.errs,
	}
}

func DoQueue(errs ...Error) Error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
