// modelRunstat.go
package main

import (
	"encoding/json"
)

type Stat struct {
	Running bool
	Uptime  string
}

func (r *Stat) ToJson() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
