package FenixMockServer

import (
	"Fenix2_testdata_mock/common_config"
	fenixClientTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixClientTestDataSyncServerGrpcApi/proto"
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
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
var localServerEngineLocalPort = common_config.FenixTestDataSyncServer_port

var (
	registerfenixTestDataSyncServerServer *grpc.Server
	lis                                   net.Listener
)

var (
	// Standard gRPC Clientr
	remoteFenixClientTestDataSyncServerConnection *grpc.ClientConn
	gRpcClientForFenixClientTestDataSyncServer    fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient

	fenixClientTestDataSyncServer_address_to_dial string = common_config.FenixClientTestDataSyncServer_address + ":" + strconv.Itoa(common_config.FenixClientTestDataSyncServer_initial_port)

	fenixClientTestDataSyncServerClient fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixTestDataGrpcServicesServer struct {
	fenixTestDataSyncServerGrpcApi.UnimplementedFenixTestDataGrpcServicesServer
}

// Channels which takes incoming gRPC-messages and pass them to 'message process engine'
// 'TestDataClientInformationMessage' from 'gRPC-RegisterTestDataClient'
var TestDataClientInformationMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage

// 'MerkleHashMessage' from 'gRPC-SendMerkleHash'
var MerkleHashMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleHashMessage

// 'MerkleTreeMessage' from 'gRPC-SendMerkleTree'
var MerkleTreeMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

// 'TestDataHeaderMessage' from 'gRPC-SendTestDataHeaders'
var TestDataHeaderMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage

// 'MerkleTreeMessage' from 'gRPC-SendTestDataRows'
var MerkleTreeMessageMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

var highestFenixProtoFileVersion int32 = -1
var highestClientProtoFileVersion int32 = -1
