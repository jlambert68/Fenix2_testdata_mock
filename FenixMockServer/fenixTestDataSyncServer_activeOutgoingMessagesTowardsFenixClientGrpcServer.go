package FenixMockServer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"time"
)

func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) SendMQmlServerIpAndPortForBackendServer() {

	var err error

	// Wait until QML-server has sent information that it has started
	// TODO -Twmp solution for waiting to QML-servers har connected
	for {
		fmt.Println("sleeping...for another 10 seconds")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
		if fenixTestDataSyncServerObject.qmlServerHasConnected == true {
			break
		}
	}

	// Set up connection to QMl Server
	remoteQmlServerConnection, err = grpc.Dial(qmlServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
			"error message":             err,
		}).Error("Did not connect to QML Server via gRPC!")
		os.Exit(0)
	} else {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
		}).Info("gRPC connection OK to QML Server!")

		// Creates a new Clients
		fenixTestDataSyncServerClient := qml_server_grpc_api.NewQmlGrpcServicesClient(remoteQmlServerConnection)

		//messageToQmlServer := &fenixTestDataSyncServerGrpcApi.WorkerInformation{fenixTestDataSyncServerObject.ip, fenixTestDataSyncServerObject.port, fenixTestDataSyncServerObject.uuid, ""}
		messageToQmlServer := &qml_server_grpc_api.BackendServerInformation{
			BackendServerIp:      fenixTestDataSyncServerObject.ip,
			BackendServerPort:    fenixTestDataSyncServerObject.port,
			BackendServerUuid:    fenixTestDataSyncServerObject.uuid,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		ctx := context.Background()
		returnMessage, err := fenixTestDataSyncServerClient.fenixTestDataSyncServerServerIPandPort(ctx, messageToQmlServer)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"client": fenixTestDataSyncServerClient,
				"error":  err,
			}).Fatal("Problem to connect to QML Server")
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"returnMessage: ":                      returnMessage,
			"fenixTestDataSyncServerObject.ip: ":   fenixTestDataSyncServerObject.ip,
			"fenixTestDataSyncServerObject.port: ": fenixTestDataSyncServerObject.port,
			"fenixTestDataSyncServerObject.uuid":   fenixTestDataSyncServerObject.uuid,
		}).Info("Sent IP and Port to QML Server")

	}

}
