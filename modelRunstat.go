// modelRunstat.go
package main

func InitlizeStatOptMgr() *OptionMgr {
	statOptMgr := CreateOptionMgr(32)
	statOptMgr.NewOption("root", "name", "Mike")
	statOptMgr.NewOption("root", "age", "18")
	statOptMgr.NewOption("root", "tall", "180")
	go statOptMgr.OptionManagerServ()
	return statOptMgr
}
