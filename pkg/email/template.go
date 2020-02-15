package email

import (
	"fmt"
	model "github.com/HFO4/cloudreve/models"
	"github.com/HFO4/cloudreve/pkg/util"
)

// NewOveruseNotification 新建超额提醒邮件
func NewOveruseNotification(userName, reason string) (string, string) {
	options := model.GetSettingByNames("siteName", "siteURL", "siteTitle", "over_used_template")
	replace := map[string]string{
		"{siteTitle}":    options["siteName"],
		"{userName}":     userName,
		"{notifyReason}": reason,
		"{siteUrl}":      options["siteURL"],
		"{siteSecTitle}": options["siteTitle"],
	}
	return fmt.Sprintf("【%s】空间容量超额提醒", options["siteName"]),
		util.Replace(replace, options["over_used_template"])
}
