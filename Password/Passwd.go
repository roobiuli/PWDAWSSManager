package Password

import "os"

type Passwd struct {
	Name  string
	Value []byte
}

func (p *Passwd) WithPass(s []byte) *Passwd {
	p.Value = s
	return p
}

func (p *Passwd) SetEnv() *Passwd {
	st := string(p.Value)
	err := os.Setenv(p.Name, st)
	if err != nil {
		panic(err)
	}
	return p
}
