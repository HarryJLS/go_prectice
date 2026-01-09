package service

// HelloServiceInterface 定义服务接口
type HelloServiceInterface interface {
	GetHelloMessage() string

	TestDemo() (string, error)
}

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) GetHelloMessage() string {
	return "Hello World"
}

func (s *HelloService) TestDemo() (string, error) {
	return "Hello World", nil
}
