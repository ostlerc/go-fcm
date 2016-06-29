// Program fcm-logger logs and echoes as a FCM "server".
package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/aliafshar/toylog"
	"github.com/mcilloni/go-fcm"
)

var (
	serverKey = kingpin.Flag("server_key", "The server key to use for FCM.").Short('k').Required().String()
	senderId  = kingpin.Flag("sender_id", "The sender ID to use for FCM.").Short('s').Required().String()
)

// onMessage receives messages, logs them, and echoes a response.
func onMessage(cm fcm.CcsMessage) error {
	toylog.Infoln("Message, from:", cm.From, "with:", cm.Data)
	// Echo the message with a tag.
	cm.Data["echoed"] = true
	m := fcm.HttpMessage{To: cm.From, Data: cm.Data}
	r, err := fcm.SendHttp(*serverKey, m)
	if err != nil {
		toylog.Errorln("Error sending message.", err)
		return err
	}
	toylog.Infof("Sent message. %+v -> %+v", m, r)
	return nil
}

func main() {
	toylog.Infoln("FCM Logger, starting.")
	kingpin.Parse()
	fcm.Listen(*senderId, *serverKey, onMessage, nil)
}
