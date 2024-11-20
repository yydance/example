package audit

// 操作审计
/*
格式：操作人、操作时间、操作内容、操作结果（操作IP、操作设备、操作浏览器、操作系统）
操作人：除了登录，其他来自于jwt token
操作内容：ctx.Method() + ctx.Path()，根据path转换成对应的功能名称
操作结果：ctx.Status()
例子：bob 2022-01-01 12:00:00 登录系统 成功 （192.168.1.1 Windows 10 Chrome 96.0.4664.45）
*/
