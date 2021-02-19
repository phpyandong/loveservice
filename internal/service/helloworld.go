
package service

import(
	"context"

	pb "loveservice/api/helloworld"
	"loveservice/internal/biz"
	"helloworld/api/helloworld/errors"
	errors2 "github.com/go-kratos/kratos/v2/errors"
)

type HelloworldService struct {
	pb.UnimplementedHelloworldServer
}

func NewHelloworldService() pb.HelloworldServer {
	return &HelloworldService{}
}

func (s *HelloworldService) CreateHelloworld(ctx context.Context, req *pb.CreateHelloworldRequest) (*pb.CreateHelloworldReply, error) {
	//dto -> do
	//
	o := new(biz.Greeter)
	o.Hello = req.hello
	//调用biz的方法

	//使用自动生成的IsXXX 判断是否是某个错误
	err := new(errors2.StatusError)
	err.Reason = errors.Errors_MissingName

	errors.IsMissingName(err)

	return &pb.CreateHelloworldReply{}, nil
}
func (s *HelloworldService) UpdateHelloworld(ctx context.Context, req *pb.UpdateHelloworldRequest) (*pb.UpdateHelloworldReply, error) {
	return &pb.UpdateHelloworldReply{}, nil
}
func (s *HelloworldService) DeleteHelloworld(ctx context.Context, req *pb.DeleteHelloworldRequest) (*pb.DeleteHelloworldReply, error) {
	return &pb.DeleteHelloworldReply{}, nil
}
func (s *HelloworldService) GetHelloworld(ctx context.Context, req *pb.GetHelloworldRequest) (*pb.GetHelloworldReply, error) {
	return &pb.GetHelloworldReply{}, nil
}
func (s *HelloworldService) ListHelloworld(ctx context.Context, req *pb.ListHelloworldRequest) (*pb.ListHelloworldReply, error) {
	return &pb.ListHelloworldReply{}, nil
}
