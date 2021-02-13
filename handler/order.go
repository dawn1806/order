package handler
import (
	"context"
	"github.com/dawn1806/common"
	"github.com/dawn1806/order/domain/model"
	"github.com/dawn1806/order/domain/service"
	. "github.com/dawn1806/order/proto/order"
)
type Order struct{
     OrderDataService service.IOrderDataService
}

func (o *Order) GetOrderByID(ctx context.Context, req *OrderID, res *OrderInfo) error {
	order, err := o.OrderDataService.FindOrderByID(req.OrderId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(order, res); err != nil {
		return err
	}
	return nil
}

func (o *Order) GetAllOrder(ctx context.Context, req *AllOrderRequest, res *AllOrder) error {
	orderSlice, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}

	for _, v := range orderSlice {
		orderInfo := &OrderInfo{}
		if err := common.SwapTo(v, orderInfo); err != nil {
			return err
		}
		res.OrderInfo = append(res.OrderInfo, orderInfo)
	}
	return nil
}

func (o *Order) CreateOrder(ctx context.Context, req *OrderInfo, res *OrderID) error {
	order := &model.Order{}
	if err := common.SwapTo(req, order); err != nil {
		return err
	}
	orderID, err := o.OrderDataService.AddOrder(order)
	if err != nil {
		return err
	}
	res.OrderId = orderID
	return nil
}

func (o *Order) DeleteOrderByID(ctx context.Context, req *OrderID, res *Response) error {
	if err := o.OrderDataService.DeleteOrder(req.OrderId); err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}

func (o *Order) UpdateOrderPayStatus(ctx context.Context, req *PayStatus, res *Response) error {
	if err := o.OrderDataService.UpdatePayStatus(req.OrderId, req.PayStatus); err != nil {
		return err
	}
	res.Msg = "支付状态更新成功"
	return nil
}

func (o *Order) UpdateOrderShipStatus(ctx context.Context, req *ShipStatus, res *Response) error {
	if err := o.OrderDataService.UpdateShipStatus(req.OrderId, req.ShipStatus); err != nil {
		return err
	}
	res.Msg = "发货状态更新成功"
	return nil
}

func (o *Order) UpdateOrder(ctx context.Context, req *OrderInfo, res *Response) error {
	order := &model.Order{}
	if err := common.SwapTo(req, order); err != nil {
		return err
	}

	if err := o.OrderDataService.UpdateOrder(order); err != nil {
		return err
	}
	res.Msg = "订单更新成功"
	return nil
}
