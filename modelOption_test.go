// modelOption_test.go
package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRun(t *testing.T) {
	mgr := CreateOptionMgr(32)
	mgr.NewOption("root", "name", "Mike")
	mgr.NewOption("root", "age", "18")
	mgr.NewOption("root", "tall", "180")
	go mgr.OptionManagerServ()
	for i := 0; i < 100000; i++ {
		opt1 := mgr.GetOptionByIndex(i % 3)
		opt2 := mgr.GetOptionByToken(opt1.Token)
		opt2.Value = fmt.Sprintf("ChangeTo%d", i)
		opt3 := mgr.PutOption(opt2)
		fmt.Print(i, *opt1, *opt2, *opt3, "\n")
		runtime.Gosched()
	}
}
