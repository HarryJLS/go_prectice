package service

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) GetHelloMessage() string {
	return "Hello World"
}