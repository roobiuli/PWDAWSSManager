package Common

import (
	"github.com/roobiuli/PWDKMREP/AwsPwService"
	"github.com/roobiuli/PWDKMREP/Password"
)

type Iterat interface {
	IsNext() bool
	ProcessPassword()
}

type PassProcessor struct {
	index int
	Pws []*Password.Passwd
	SM *AwsPwService.PWService
}

func (p *PassProcessor) WithSecretManager(s *AwsPwService.PWService) *PassProcessor {
	p.SM = s
	return p
}

func (p *PassProcessor) WithPassColletion(c []*Password.Passwd) *PassProcessor {
	p.Pws = c
	return p
}


func (p *PassProcessor) IsNext() bool {
	if p.index < len(p.Pws) {
		return true
	}
	return false
}

func (p *PassProcessor) ProcessPassword()  {
		pa := p.Pws[p.index]
		data, _ := p.SM.GetPassword(pa)

		pa.WithPass(data).SetEnv()
		p.index++

}





