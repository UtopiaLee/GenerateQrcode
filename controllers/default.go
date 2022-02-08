package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	qrcode "github.com/skip2/go-qrcode"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetQrcode() {
   png,_:= qrcode.Encode(c.GetString("k"), qrcode.Medium, 256)
   c.Ctx.Output.ContentType("png")
   c.Ctx.Output.Body(png)
}

func (c *MainController) Get() {
	key :=c.GetString("k")
	if key=="" {
		key="123123123"
	}
	png,_:= qrcode.Encode(key, qrcode.Medium, 256)
	c.Ctx.Output.ContentType("png")
	c.Ctx.Output.Body(png)
 }
 
