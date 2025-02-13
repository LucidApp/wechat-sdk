package context

import (
	"github.com/lucidapp/wechatsdkgo/v2/credential"
	"github.com/lucidapp/wechatsdkgo/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
