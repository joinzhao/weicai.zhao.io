package errorx

var (
	ParamVerifyError    Error = &errorX{code: 10001, msg: "param verify error", errs: nil}
	RecordNotFoundError Error = &errorX{code: 10002, msg: "record not found", errs: nil}

	AuthorizationError Error = &errorX{code: 40001, msg: "access denied", errs: nil}

	UnknownError Error = &errorX{code: 50000, msg: "unknown error", errs: nil}
	MysqlError   Error = &errorX{code: 50001, msg: "unknown error", errs: nil}
)

type Option func() Error

func Do(ops ...Option) Error {
	for i := 0; i < len(ops); i++ {
		if err := ops[i](); err != nil {
			return err
		}
	}
	return nil
}
