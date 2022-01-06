package FenixMockServer

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/qml_server_grpc_api"
	"net"
	"time"
)

type TestInstructionBackendObject_struct struct {
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

var testInstructionBackendObject *TestInstructionBackendObject_struct

// Global connection constants
var localServerEngineLocalPort = common_config.TestInstructionBackendServer_initial_port

var (
	registerTestInstructionBackendServer *grpc.Server
	lis                                  net.Listener
)

var (
	// Standard gRPC Server
	remoteQmlServerConnection *grpc.ClientConn
	gRpcClientForQmlServer    qml_server_grpc_api.QmlGrpcServicesClient

	qmlServer_address_to_dial string = common_config.QmlServer_address + common_config.QmlServer_port
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type TestInstructionBackendServer struct{}
