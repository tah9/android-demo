package net.away0x.lib_pay_center.data.api

import io.reactivex.Observable
import retrofit2.http.Body
import net.away0x.lib_pay_center.data.protocol.GetPaySignReq
import net.away0x.lib_pay_center.data.protocol.PayOrderReq
import net.away0x.lib_base.data.protocol.BaseResp
import retrofit2.http.POST


/*
    支付 接口
 */
interface PayApi {

    /*
        获取支付宝支付签名
     */
    @POST("pay/getPaySign")
    fun getPaySign(@Body req: GetPaySignReq): Observable<BaseResp<String>>

    /*
        刷新订单状态，已支付
     */
    @POST("order/pay")
    fun payOrder(@Body req: PayOrderReq): Observable<BaseResp<String>>

}
