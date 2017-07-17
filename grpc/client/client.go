package main

import (
	"context"

	"go_test/grpc/logger"
	"go_test/grpc/protoc/order"

	"google.golang.org/grpc"
)

func GetOrders(addr string) (rslt *order.OrderRslt, err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Error("failed to connect order server:", err)
		return
	}

	log.Debug("connect order server success")

	c := order.NewOrderServiceClient(conn)
	rslt, err = c.GetOrderDetail(context.Background(), &order.OrderReq{})
	if err != nil {
		log.Error("failed to getOrderDetail:", err)
		return
	}

	log.Debug("get order detail from order server:", rslt)
	return
}

func main() {
	GetOrders("127.0.0.1:8089")
}
