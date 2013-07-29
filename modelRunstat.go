// modelRunstat.go
package main

func InitlizeStatOptMgr() *OptionMgr {
	OptMgr := CreateOptionMgr(32)
	OptMgr.NewOption("root", "name", "Mike")
	OptMgr.NewOption("root", "age", "18")
	OptMgr.NewOption("root", "tall", "180")
	go OptMgr.OptionManagerServ()
	return OptMgr
}
