package main

import (
	"math/rand"
	"net"
	"time"

	log "go_test/grpc/logger"
	"go_test/grpc/protoc/common"
	"go_test/grpc/protoc/data/ordervo"
	"go_test/grpc/protoc/order"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

var config struct {
	net     string
	address string
}

var order1 = &order.Order{
	Order: &ordervo.OrderVo{
		OrderId:    "0000000001",
		MemberCode: 00000001,
	},
	Suborders: []*ordervo.SuborderVo{
		{
			SuborderId:    "A000000001",
			OrderId:       "0000000001",
			MemberCode:    00000001,
			Channel:       "froadmall",
			MerchantId:    "000001",
			MerchantName:  "test1",
			SuborderPrice: 6.12,
			Freight:       12.4,
			CreateTime:    int32(time.Now().Unix()),
			UpdateTime:    int32(time.Now().Unix()),
		},
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s server) GetOrderDetail(ctx context.Context, req *order.OrderReq) (rslt *order.OrderRslt, err error) {
	log.Debug("got GetOrderDetail req:", *req)

	rslt = &order.OrderRslt{
		Rslt:   &common.Rslt{},
		Orders: make([]*order.Order, 0, 12),
	}

	//if rand.Intn(100) >= 90 { // annotate it in benchmark
	//	err = errors.New("something not predictable happened")
	//	log.Error(err)
	//	return rslt, err
	//}

	rslt.Orders = append(rslt.Orders, order1)
	order1.Order.MemberCode += 1
	log.Info("result: ", rslt.Orders)

	return rslt, nil
}

func Config(net, address string) error {
	config.net, config.address = net, address
	return nil
}

func Serve() {
	l, err := net.Listen(config.net, config.address)
	if err != nil {
		log.Error("failed to listen: ", err)
	}
	s := grpc.NewServer()
	order.RegisterOrderServiceServer(s, server{})
	s.Serve(l)
}

func main() {
	config.net, config.address = "tcp", "127.0.0.1:8089"
	Serve()
}
