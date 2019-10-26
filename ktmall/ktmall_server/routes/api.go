package routes

import (
	"ktmall/app/context"
	. "ktmall/app/controllers"

	"github.com/labstack/echo/v4"
)

const (
	APIPrefix = "/api"
)

func registerAPI(e *echo.Echo) {
	ee := e.Group(APIPrefix)

	context.RegisterHandler(ee.GET, "/upload_token", CommonGetQiNiuUploadToken)

	user := ee.Group("/user")
	{
		context.RegisterHandler(user.POST, "/register", UserRegister)
		context.RegisterHandler(user.POST, "/login", UserLogin)
		context.RegisterHandler(user.POST, "/forget_pwd", UserForgetPwd)
		context.RegisterHandler(user.POST, "/reset_pwd", UserResetPwd)
		context.RegisterHandler(user.POST, "/edit", UserEdit)
		context.RegisterHandler(user.GET, "/info", UserInfo)
	}

	address := ee.Group("/address")
	{
		context.RegisterHandler(address.POST, "/add", AddressAdd)
		context.RegisterHandler(address.POST, "/delete", AddressDelete)
		context.RegisterHandler(address.POST, "/modify", AddressModify)
		context.RegisterHandler(address.GET, "/list", AddressList)
	}

	cart := ee.Group("/cart")
	{
		context.RegisterHandler(cart.GET, "/list", CartList)
		context.RegisterHandler(cart.POST, "/add", CartAdd)
		context.RegisterHandler(cart.POST, "/delete", CartDelete)
		context.RegisterHandler(cart.POST, "/submit", CartSubmit)
	}

	category := ee.Group("/category")
	{
		context.RegisterHandler(category.GET, "/list", CategoryList)
	}

	goods := ee.Group("/goods")
	{
		context.RegisterHandler(goods.GET, "/list", GoodsList)
		context.RegisterHandler(goods.GET, "/list_by_keyword", GoodsListByKeyword)
		context.RegisterHandler(goods.GET, "/detail", GoodsDetail)
	}

	message := ee.Group("/message")
	{
		context.RegisterHandler(message.GET, "/list", MessageList)
	}

	order := ee.Group("/order")
	{
		context.RegisterHandler(order.POST, "/get_pay_sign", OrderGetPaySign)
		context.RegisterHandler(order.POST, "/pay", OrderPay)
		context.RegisterHandler(order.POST, "/cancel", OrderCancel)
		context.RegisterHandler(order.POST, "/confirm", OrderConfirm)
		context.RegisterHandler(order.GET, "/detail", OrderDetail)
		context.RegisterHandler(order.GET, "/list", OrderList)
		context.RegisterHandler(order.POST, "/submit", OrderSubmit)
	}
}