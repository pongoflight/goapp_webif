// modelParameter
package main

func InitlizeParaOptMgr() *OptionMgr {
	OptMgr := CreateOptionMgr(32)
	OptMgr.NewOption("ETH0", "IP地址", "172.16.10.100")
	OptMgr.NewOption("ETH0", "子网掩码", "255.255.0.0")
	OptMgr.NewOption("ETH1", "IP地址", "172.16.100.10")
	OptMgr.NewOption("ETH1", "子网掩码", "255.255.0.0")
	go OptMgr.OptionManagerServ()
	return OptMgr
}
