package _1_proxy

type ServiceInterface interface {
	Execute(access string) string
}

type Service struct {
}

func (t *Service) Execute(access string) string {
	return "Proxy Service-" + access
}
