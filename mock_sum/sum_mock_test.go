package mock_sum

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/internal/grpctest"
	"../sum"
	"testing"
	"time"
)

type grpcTester struct {
	grpctest.Tester
}

func Test(t *testing.T) {
	grpctest.RunSubTests(t, grpcTester{})
}

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func (grpcTester) TestSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSumComputerClient := NewMockSumComputerClient(ctrl)
	req := sum.SumRequest{FirstOperand:10, SecondOperand:34}
	//mockSumComputerClient.EXPECT().SayHello(
	//	gomock.Any(),
	//	&rpcMsg{msg: req},
	//).Return(&helloworld.HelloReply{Message: "Mocked Interface"}, nil)
	//testSayHello(t, mockSumComputerClient)
	mockSumComputerClient.EXPECT()
}

func testSayHello(t *testing.T, client helloworld.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "unit_test"})
	if err != nil || r.Message != "Mocked Interface" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Message)
}