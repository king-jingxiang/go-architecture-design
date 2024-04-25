package _1_proxy

type Proxy struct {
	realService *Service
}

func NewProxy() *Proxy {
	return &Proxy{
		realService: &Service{},
	}
}

func (p *Proxy) Execute(access string) string {
	if access == "yes" {
		return p.realService.Execute(access)
	}
	return access
}
