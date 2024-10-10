package log

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(Disable)
	if errorLog.Writer() == os.Stdout || infoLog.Writer() == os.Stdout {
		t.Fatal("关闭日志失败")
	}
	SetLevel(ErrorLevel)
	if errorLog.Writer() != os.Stdout || infoLog.Writer() == os.Stdout {
		t.Fatal("设置为仅显示错误日志失败")
	}
	SetLevel(InfoLevel)
	if errorLog.Writer() != os.Stdout || infoLog.Writer() != os.Stdout {
		t.Fatal("设置为显示所有日志失败")
	}
}
