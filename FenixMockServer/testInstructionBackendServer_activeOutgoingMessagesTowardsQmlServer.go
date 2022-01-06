package FenixMockServer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/qml_server_grpc_api"
	"os"
	"time"
)

func (testInstructionBackendObject *TestInstructionBackendObject_struct) SendMQmlServerIpAndPortForBackendServer() {

	var err error

	// Wait until QML-server has sent information that it has started
	// TODO -Twmp solution for waiting to QML-servers har connected
	for {
		fmt.Println("sleeping...for another 10 seconds")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
		if testInstructionBackendObject.qmlServerHasConnected == true {
			break
		}
	}

	// Set up connection to QMl Server
	remoteQmlServerConnection, err = grpc.Dial(qmlServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
			"error message":             err,
		}).Error("Did not connect to QML Server via gRPC!")
		os.Exit(0)
	} else {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
		}).Info("gRPC connection OK to QML Server!")

		// Creates a new Clients
		testInstructionBackendClient := qml_server_grpc_api.NewQmlGrpcServicesClient(remoteQmlServerConnection)

		//messageToQmlServer := &fenixClientTestDataSyncServerGrpcApi.WorkerInformation{testInstructionBackendObject.ip, testInstructionBackendObject.port, testInstructionBackendObject.uuid, ""}
		messageToQmlServer := &qml_server_grpc_api.BackendServerInformation{
			BackendServerIp:      testInstructionBackendObject.ip,
			BackendServerPort:    testInstructionBackendObject.port,
			BackendServerUuid:    testInstructionBackendObject.uuid,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		ctx := context.Background()
		returnMessage, err := testInstructionBackendClient.TestInstructionBackendServerIPandPort(ctx, messageToQmlServer)
		if err != nil {
			testInstructionBackendObject.logger.WithFields(logrus.Fields{
				"client": testInstructionBackendClient,
				"error":  err,
			}).Fatal("Problem to connect to QML Server")
		}

		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"returnMessage: ":                     returnMessage,
			"testInstructionBackendObject.ip: ":   testInstructionBackendObject.ip,
			"testInstructionBackendObject.port: ": testInstructionBackendObject.port,
			"testInstructionBackendObject.uuid":   testInstructionBackendObject.uuid,
		}).Info("Sent IP and Port to QML Server")

	}

}
