package log

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/rs/xid"
)

type D struct {
	Index    string   `json:"@index"`
	Type     string   `json:"@type"`
	Level    string   `json:"level"`
	Time     xtime    `json:"datetime"`
	UniqueID string   `json:"unique_id"`
	UID      int      `json:"uid"`
	Info     *logInfo `json:"info"`
}

type logInfo struct {
	Host    string `json:"host"`
	Extra   string `json:"extra"`
	Message string `json:"message"`
	Context string `json:"context"`
}

type agentHandler struct {
	conf *AgentConfig
	conn net.Conn
}

func agentInit(conf *AgentConfig) *agentHandler {
	conn, err := net.Dial(conf.Proto, conf.Addr)
	if err != nil {
		panic(err)
	}
	sh := new(agentHandler)
	sh.conn = conn
	sh.conf = conf
	return sh
}

func (ah *agentHandler) info(format string, args ...interface{}) {
	ah.logf("info", format, args...)
}

func (ah *agentHandler) warn(format string, args ...interface{}) {
	ah.logf("warn", format, args...)
}

func (ah *agentHandler) debug(format string, args ...interface{}) {
	ah.logf("debug", format, args...)
}

func (ah *agentHandler) error(format string, args ...interface{}) {
	ah.logf("error", format, args...)
}

func (ah *agentHandler) logf(l string, format string, args ...interface{}) {
	di := new(logInfo)
	di.Message = fmt.Sprintf(format, args...)
	di.Host, _ = os.Hostname()
	data := D{
		Index:    ah.conf.TaskID,
		Time:     xtime(time.Now()),
		Level:    l,
		UniqueID: ah.uniqueID(),
		Info:     di,
	}
	b, _ := json.Marshal(data)
	if _, err := ah.conn.Write(b); err != nil {
	}
}

func (ah *agentHandler) uniqueID() string {
	return xid.New().String()
}