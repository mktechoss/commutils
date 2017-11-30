package servicenotify

import "testing"

func TestEmailNotify_Send(t *testing.T) {
	var notifyer Notify

	notifyer = NewEmailInstance("13424028418@163.com", "ht101996",
		"smtp.163.com:25", "675509312@qq.com")
	notifyer.Send("From mktechoss 报警通知", "当前服务器挂了!!")
}
