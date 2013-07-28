// modelOption
package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Option struct {
	Token string
	Group string
	Name  string
	Value string
}

func (self *Option) ToJson() []byte {
	buf, _ := json.Marshal(self)
	return buf
}

type OptionWithChan struct {
	Opt    Option
	Action string
	Index  int
	Chan   chan Option
}

type OptionMgr struct {
	OptDict   map[string]*Option
	OptList   []*Option
	OptLength int
	Chan      chan OptionWithChan
}

func genToken() string {
	inbuf := make([]byte, 32)
	outbuf := make([]byte, sha1.Size*2)
	var n = 0
	for n < len(inbuf) {
		var nn int
		nn, _ = rand.Reader.Read(inbuf[n:])
		n += nn
	}
	hashObj := sha1.New()
	hashObj.Write(inbuf)
	hash := hashObj.Sum(nil)
	hex.Encode(outbuf, hash)
	return string(outbuf)
}

func CreateOptionMgr(optCount int) *OptionMgr {
	return &OptionMgr{
		OptDict:   make(map[string]*Option),
		OptList:   make([]*Option, optCount),
		OptLength: 0,
		Chan:      make(chan OptionWithChan, 16)}
}

func (self *OptionMgr) NewOption(group, name, value string) *Option {
	token := genToken()
	opt := &Option{Token: token, Group: group, Name: name, Value: value}
	self.OptList[self.OptLength] = opt
	self.OptDict[token] = opt
	self.OptLength++
	return opt
}

func (self *OptionMgr) GetOptionCount() int {
	return self.OptLength
}

func (self *OptionMgr) GetOptionByIndex(index int) *Option {
	recvChan := make(chan Option)
	defer close(recvChan)
	self.Chan <- OptionWithChan{
		Opt:    Option{},
		Action: "GET", Index: index, Chan: recvChan}
	recvopt := <-recvChan
	return &recvopt
}

func (self *OptionMgr) GetOptionByToken(token string) *Option {
	recvChan := make(chan Option)
	defer close(recvChan)
	self.Chan <- OptionWithChan{
		Opt:    Option{Token: token},
		Action: "GET", Chan: recvChan}
	recvopt := <-recvChan
	return &recvopt
}

func (self *OptionMgr) PutOption(opt *Option) *Option {
	recvChan := make(chan Option)
	defer close(recvChan)
	self.Chan <- OptionWithChan{
		Opt:    *opt,
		Action: "PUT", Chan: recvChan}
	recvopt := <-recvChan
	return &recvopt
}

func (self *OptionMgr) PrintOptions() {
	for i := 0; i < self.OptLength; i++ {
		fmt.Println(self.OptList[i])
	}
}

func (self *OptionMgr) OptionManagerServ() {
	for {
		optchan := <-self.Chan
		if optchan.Action == "GET" {
			var opt *Option
			var ok bool
			if optchan.Opt.Token == "" {
				ok = (optchan.Index >= 0)
				if ok {
					idx := optchan.Index % self.OptLength
					opt = self.OptList[idx]
				}
			} else {
				opt, ok = self.OptDict[optchan.Opt.Token]
			}
			if ok {
				optchan.Opt.Token = opt.Token
				optchan.Opt.Group = opt.Group
				optchan.Opt.Name = opt.Name
				optchan.Opt.Value = opt.Value
			} else {
				optchan.Opt.Token = ""
			}
			optchan.Chan <- optchan.Opt
		}
		if optchan.Action == "PUT" {
			opt, ok := self.OptDict[optchan.Opt.Token]
			if ok {
				opt.Group = optchan.Opt.Group
				opt.Name = optchan.Opt.Name
				opt.Value = optchan.Opt.Value
			} else {
				optchan.Opt.Token = ""
			}
			optchan.Chan <- optchan.Opt
		}
	}
}
