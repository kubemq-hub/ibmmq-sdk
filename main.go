package main

import (
	"github.com/kubemq-hub/ibmmq-sdk/mq-golang-jms20/mqjms"
	"log"
)

func main() {
	cf := mqjms.ConnectionFactoryImpl{
		QMName:           "QM1",
		Hostname:         "localhost",
		PortNumber:       1414,
		ChannelName:      "DEV.APP.SVRCONN",
		UserName:         "admin",
		TransportType:    mqjms.TransportType_CLIENT,
		TLSClientAuth:    mqjms.TLSClientAuth_NONE,
		KeyRepository:    "passw0rd",
		CertificateLabel: "",
	}

	jmsContext, err := cf.CreateContext()
	if err != nil {
		log.Fatal(err)
	}


	queue := jmsContext.CreateQueue("test")
	if queue == nil {
		log.Fatal(queue)
	}
}
