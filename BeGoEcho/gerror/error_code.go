package gerror

type Error struct {
	Error error
	Code  uint32
	Line  string
}

func New(code uint32, err error, line string) *Error {
	return &Error{
		Error: err,
		Code:  code,
		Line:  line,
	}
}

/** Client-side Error code */
const (
	ErrorBindData  uint32 = 40000
	ErrorValidData uint32 = 40001
)

/** Client-side Error Code */
const (
	ErrorConnect     uint32 = 50000
	ErrorSaveData    uint32 = 50001
	ErrorRetriveData uint32 = 50002
	ErrorNotFound    uint32 = 50003
	ErrorInternal    uint32 = 50004
	ErrorBadRequest  uint32 = 50005
)
