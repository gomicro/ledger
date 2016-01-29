package penname

import ()

type PenName struct {
	Written     []byte
	returnError error
}

func New() *PenName {
	return &PenName{}
}

func (p *PenName) Clear() {
	p.Written = []byte{}
}

func (p *PenName) ReturnError(err error) {
	p.returnError = err
}

func (p *PenName) Write(b []byte) (n int, err error) {
	if p.returnError != nil {
		return 0, p.returnError
	}

	p.Written = b
	return len(p.Written), nil
}
