package net.away0x.lib_order_center.data.protocol

/*
    修改收货地址
 */
data class EditShipAddressReq(val id:Int,val shipUserName:String,val shipUserMobile:String,val shipAddress:String,val shipIsDefault:Int)
