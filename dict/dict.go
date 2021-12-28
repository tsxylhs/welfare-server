package dict

// 微信接口
const (
	WxLogin  = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WxUnid   = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
	WXreqUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
)

const (

	// 彩票小程序
	LibrarayId    = "wxf18ea67803d92fcd"
	LibrarySecret = "95fcccc4a80190e1167558245a82e81b"
)
