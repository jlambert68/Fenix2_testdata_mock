package FenixMockServer

import (
	"Fenix2_testdata_mock/common_config"
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

type fenixTestDataSyncServerObject_struct struct {
	logger                *logrus.Logger
	iAmBusy               bool
	uuid                  string
	startTime             time.Time
	timeBeforeFinish      int32
	currentTaskuuid       string
	currentTaskName       string
	ip                    string
	port                  string
	qmlServerHasConnected bool
}

var fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct

// Global connection constants
var localServerEngineLocalPort = common_config.TestInstructionBackendServer_initial_port

var (
	registerfenixTestDataSyncServerServer *grpc.Server
	lis                                   net.Listener
)

/*
var (
	// Standard gRPC Clientr
	remoteQmlServerConnection *grpc.ClientConn
	gRpcClientForQmlServer    qml_server_grpc_api.QmlGrpcServicesClient

	qmlServer_address_to_dial string = common_config.QmlServer_address + common_config.QmlServer_port
)
*/
// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixTestDataGrpcServicesServer struct {
	fenixTestDataSyncServerGrpcApi.UnimplementedFenixTestDataGrpcServicesServer
}
