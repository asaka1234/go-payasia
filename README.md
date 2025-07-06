整体流程
==============
1. server负责对请求参数做sign签名. 随后把参数和sign签名一起返回给fe前端
2. 前端构造一个form做action提交. psp收到后会做redirect到收银台页面.
3. 用户在收银台做支付
4. server收到回调


鉴权
==============
1. 后端做一个参数的签名计算. 随后前端把sign当做form里的一个字段一起action提交
2. 是对请求的所有参数一起做签名


回调地址
==============
callback地址是portal里写死的. 静态的


Comment
===============
1. both support deposit && withdrawl
2. 在pre接口中有一个参数notification_url, 这个是充值回调url指定) 
3. 所有接口都是 application/x-www-form-urlencoded 格式的