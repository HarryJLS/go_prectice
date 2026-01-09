package service

// 为了检查方法是否都实现
var _ HelloService = (*helloService)(nil) // HelloService 定义服务接口
type HelloService interface {
	GetHelloMessage() string

	TestDemo() (string, error)
}

type helloService struct {
	workerService *WorkerService
}

func NewHelloService(workerService *WorkerService) HelloService {

	return &helloService{
		workerService: workerService,
	}
}

func (s *helloService) GetHelloMessage() string {
	return "Hello World"
}

func (s *helloService) TestDemo() (string, error) {
	return "Hello World", nil
}
