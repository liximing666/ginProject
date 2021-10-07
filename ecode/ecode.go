package ecode

var codeMap = map[int]bool{}

type ErrInfo struct {
	ecode int
	msg string
}

func NewErrInfo(ecode int, msg string) ErrInfo {
	return ErrInfo{ecode: ecode, msg: msg}
}

func (e *ErrInfo) GetEcode() int {
	return e.ecode
}

func (e *ErrInfo) SetEcode(ecode int) {
	e.ecode = ecode
}

func (e *ErrInfo) GetMsg() string {
	return e.msg
}

func (e *ErrInfo) SetMsg(msg string) {
	e.msg = msg
}

func Add(code int) ErrInfo {
	_, ok := codeMap[code]
	if ok {
		panic("this code already have")
	}

	codeMap[code] = true

	return ErrInfo{
		ecode: code,
	}

}