package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 防护策略检测模块开关选项，如是否开启Web基础防护等
type PolicyOption struct {

	// 基础防护是否开启
	Webattack *bool `json:"webattack,omitempty" xml:"webattack"`

	// 常规检测是否开启
	Common *bool `json:"common,omitempty" xml:"common"`

	// 预留参数，改参数值一直为true，用户可忽略该参数值
	Crawler *bool `json:"crawler,omitempty" xml:"crawler"`

	// 搜索engine是否开启
	CrawlerEngine *bool `json:"crawler_engine,omitempty" xml:"crawler_engine"`

	// 反爬虫检测是否开启
	CrawlerScanner *bool `json:"crawler_scanner,omitempty" xml:"crawler_scanner"`

	// 脚本反爬虫是否开启
	CrawlerScript *bool `json:"crawler_script,omitempty" xml:"crawler_script"`

	// 其他爬虫是否开启
	CrawlerOther *bool `json:"crawler_other,omitempty" xml:"crawler_other"`

	// Webshell检测是否开启
	Webshell *bool `json:"webshell,omitempty" xml:"webshell"`

	// cc规则是否开启
	Cc *bool `json:"cc,omitempty" xml:"cc"`

	// 精准防护是否开启
	Custom *bool `json:"custom,omitempty" xml:"custom"`

	// 黑白名单防护是否开启
	Whiteblackip *bool `json:"whiteblackip,omitempty" xml:"whiteblackip"`

	// 地理位置访问控制规则是否开启
	Geoip *bool `json:"geoip,omitempty" xml:"geoip"`

	// 误报屏蔽是否开启
	Ignore *bool `json:"ignore,omitempty" xml:"ignore"`

	// 隐私屏蔽是否开启
	Privacy *bool `json:"privacy,omitempty" xml:"privacy"`

	// 网页防篡改规则是否开启
	Antitamper *bool `json:"antitamper,omitempty" xml:"antitamper"`

	// 防敏感信息泄露规则是否开启
	Antileakage *bool `json:"antileakage,omitempty" xml:"antileakage"`

	// 网站反爬虫总开关是否开启
	BotEnable *bool `json:"bot_enable,omitempty" xml:"bot_enable"`

	// modulex智能cc防护是否开启，该特性是公测特性，在公测期间，只支持仅记录模式。
	ModulexEnabled *bool `json:"modulex_enabled,omitempty" xml:"modulex_enabled"`
}

func (o PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyOption struct{}"
	}

	return strings.Join([]string{"PolicyOption", string(data)}, " ")
}
