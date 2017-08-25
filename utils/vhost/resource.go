// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatedier/frp/utils/version"
)

const (
	NotFound = `<!doctype html><html lang="en-US"><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"><meta name="keywords" content="frp,frp穿透,内网穿透,免费frp,frp服务器,ngrok,花生壳,绑定域名"/>
<meta name="description" content="提供高性能的反向代理应用，类似于花生壳、ngrok的工具，可以帮助您轻松地进行内网穿透，对外网提供服务，支持 tcp, udp, http, https 等协议类型，并且 web 服务支持根据域名进行路由转发，可绑定自己的顶级域名。"/><title>404 - 对不起，您查找的页面不存在！_boke112导航</title><style>a,body,div,h1,h2,html,p,span{margin:0;padding:0;border:0;font-size:100%;font:inherit;vertical-align:baseline;outline:0}body{background:#f2f2f2;font-family:"Microsoft YaHei",Helvetica,Arial,Lucida Grande,Tahoma,sans-serif;-webkit-font-smoothing:antialiased;overflow:hidden}.right{float:right}#main{position:relative;width:100%;max-width:600px;margin:0 auto;padding-top:8%}h1{position:relative;display:block;font:72px "Microsoft YaHei",Helvetica,Arial,Lucida Grande,Tahoma,sans-serif;color:#16a085;text-shadow:2px 2px #f7f7f7;text-align:center}.sub{position:relative;font-size:21px;top:-20px;padding:0 10px;font-style:italic}@media screen and (max-width:374px){.sub{display:none}}.icon{position:relative;display:inline-block;top:-6px;margin:0 10px 5px 0;background:#16a085;width:50px;height:50px;-moz-box-shadow:1px 2px #fff;-webkit-box-shadow:1px 2px #fff;box-shadow:1px 2px #fff;-webkit-border-radius:50px;-moz-border-radius:50px;border-radius:50px;color:#dfdfdf;font-size:46px;line-height:48px;font-weight:700;text-align:center;text-shadow:0 0}#content{position:relative;width:100%;max-width:600px;background:#fff;-webkit-border-radius:5px;-moz-border-radius:5px;border-radius:5px;z-index:5}h2{background:url(data:image/jpeg;base64,/9j/4QAYRXhpZgAASUkqAAgAAAAAAAAAAAAAAP/sABFEdWNreQABAAQAAABkAAD/4QMraHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLwA8P3hwYWNrZXQgYmVnaW49Iu+7vyIgaWQ9Ilc1TTBNcENlaGlIenJlU3pOVGN6a2M5ZCI/PiA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJBZG9iZSBYTVAgQ29yZSA1LjAtYzA2MCA2MS4xMzQ3NzcsIDIwMTAvMDIvMTItMTc6MzI6MDAgICAgICAgICI+IDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+IDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bXA6Q3JlYXRvclRvb2w9IkFkb2JlIFBob3Rvc2hvcCBDUzUgTWFjaW50b3NoIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjA4MTc1NUYyNDI3NjExRTE4MkFCOThDNzMzMDg1MzRFIiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjA4MTc1NUYzNDI3NjExRTE4MkFCOThDNzMzMDg1MzRFIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6MDgxNzU1RjA0Mjc2MTFFMTgyQUI5OEM3MzMwODUzNEUiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6MDgxNzU1RjE0Mjc2MTFFMTgyQUI5OEM3MzMwODUzNEUiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz7/7gAOQWRvYmUAZMAAAAAB/9sAhAABAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAgICAgICAgICAgIDAwMDAwMDAwMDAQEBAQEBAQIBAQICAgECAgMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwP/wAARCAAKAlIDAREAAhEBAxEB/8QAeAABAQEBAQAAAAAAAAAAAAAAAAIBAwoBAQAAAAAAAAAAAAAAAAAAAAAQAAIAAQYHDAkFAQAAAAAAAAABAhExkgMEBYGRwdFTBhbwIVFx4RJSQ1QVBxdhobHxQoLiE0RBorIjFNIRAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwEAAhEDEQA/APfwAASSzgTzIeLd6QJdXC+WR5AJ+zDuUnsYGOog4FLxcoEOywudLdgAz/JBwLdgAh2GrcySxgT3fVcEOICHdtT0Vx72UCHddQ99wfxAl3RUP4FiUuPnAc+5qhzwQ+oCe47O56uECHcFmfwL1ZJAJer1lc9XDiWcCNm7Jo4cUIEvVixOX+uGXhkQHN6q2N77q4MSzAQ9VLC56qDj5uWQCNkbBooaP0gQ9T7A5f6oKLzAQ9TLvc9TC/kXIBOxV2ueohooDm9SLufUQ4YU8gEvUW7XPUQUQJeod2PqIKKypgTsDdfZ4aKzAS/D+63NUQL5EBHl5dWggoIDH4d3S57PBQQEvw5uh/jwUFv40wJ8t7o7NV0UvZCBL8Nrof49XhhlyIDPLW6OzVVDkAeWt0dmqqHIAXhrdC/HqsEEmQCvLe6OzwYl/wAgUvDi6V+PAuKBb/7UBXl1dHZ4KAFeXl1aCCggNXh7damqIOLmKQC9gLqU1ngorMBS1CuxdRDRWYCthLt7PBRArYe7tBBRWYDotSbtXULDCnmApamXdoVRQFrU6waKHDBmAtao2BKT7UNH6WBa1TsK6qF8cP0gXsrY9HBuwAWtV7Gp6uDAlyAdNmrGurhxIClq7ZF1cOFL3AVs/ZuhDiQHTuOzaOF4siQFdy1HQh9T9oFK56hTwLEnlQHTuqo6HsWUCldtSvg/iBau+qXwpeh7mgKVgq10cQFqx1amhW7ABv8Alg4Fj5ALVnh/VLF7gK+xDulzgUqqH3JLOBvMXpApJKZAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/9k=) no-repeat;background-position:bottom;padding:12px 0 22px 0;font:20px "Microsoft YaHei",Helvetica,Arial,Lucida Grande,Tahoma,sans-serif;text-align:center}p{position:relative;padding:20px;font-size:14px;line-height:25px}.utilities{padding:0 20px 20px}.utilities .button{display:inline-block;height:34px;margin:0 0 0 6px;padding:0 18px;background:#16a085;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px;font-size:14px;line-height:34px;color:#fff;font-weight:700;text-decoration:none}</style></head><body><div id="wrapper"> <div id="main"> <header id="header"> <h1><span class="icon">!</span>404<span class="sub">page not found</span></h1> </header> <div id="content"> <h2>您打开的这个的页面不存在！</h2> <p>当您看到这个页面，表示您的访问出错，这个错误是您打开的页面不存在，请确认您输入的地址是正确的，如果是在本站点击后出现这个页面，请联系站长进行处理，感谢您的支持!</p> <div class="utilities"> <a class="button right" href="http://www.zishuo.net">技术支持</a> <a class="button" href="http://wpa.qq.com/msgrd?V=3&uin=43509704&Site=QQ&Menu=yes">联系站长</a> <div class="clear"></div> </div> </div> </div></div></html>
`
)

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")
	res := &http.Response{
		Status:     "Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(strings.NewReader(NotFound)),
	}
	return res
}
