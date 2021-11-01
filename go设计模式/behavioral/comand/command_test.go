package comand

import (
	"testing"
)

/**
命令模式就是一个人接到命令后继续往下传，最终让这个命令执行
*/
func TestCommand_Execute(t *testing.T) {
	laowang := NewPerson("wang", NewCommand(nil, nil))
	laozhang := NewPerson("zhang", NewCommand(&laowang, laowang.Listen))
	loafeng := NewPerson("feng", NewCommand(&laozhang, laozhang.Buy))
	laoding := NewPerson("ding", NewCommand(&loafeng, loafeng.Cook))
	laoli := NewPerson("li", NewCommand(&laoding, laoding.Wash))

	laoli.Talk()

}
