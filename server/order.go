package server

import (
	"context"
	"errors"
	"github.com/bill-lhr/grpc-demo/proto/order"
	"time"
)

// GetOrderInfo 具体的方法实现
func (s *Server) GetOrderInfo(ctx context.Context, request *order.OrderRequest) (*order.OrderInfo, error) {
	orderMap := map[string]order.OrderInfo{
		"201907300001": {OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": {OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": {OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	var response order.OrderInfo
	current := time.Now().Unix()
	if request.Timestamp > current {
		response = order.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result, ok := orderMap[request.OrderId]
		if ok && result.OrderId != "" {
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return &response, nil
}
