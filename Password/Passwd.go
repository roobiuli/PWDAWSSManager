package Password

import "os"

type Passwd struct {
	Name string
	Value *string
}

func (p *Passwd) WithPass(s *string) *Passwd {
	p.Value = s
	return p
}


func (p *Passwd) SetEnv() *Passwd {

	err := os.Setenv(p.Name, *p.Value)
	if err != nil {
		panic(err)
	}
	return p
}

