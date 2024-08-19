package dict

// 微信接口
const (
	WxLogin  = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WxUnid   = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
	WXreqUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
)

const (
	// 彩票小程序
	LibrarayId    = "wxe863950d7a4f8ac0"
	LibrarySecret = "2f12460a03b3701369020bd36a0279a1"
)
