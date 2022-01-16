package FenixClientServer

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

type fenixClientTestDataSyncServerObject_struct struct {
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

var fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct

// Global connection constants
var localServerEngineLocalPort = common_config.FenixClientTestDataSyncServer_initial_port

var (
	registerfenixClientTestDataSyncServerServer *grpc.Server
	lis                                         net.Listener
)

var (
	// Standard gRPC Clientr
	remoteFenixTestDataSyncServerConnection *grpc.ClientConn
	gRpcClientForFenixTestDataSyncServer    fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient

	fenixTestDataSyncServer_address_to_dial string = common_config.FenixTestDataSyncServer_address + ":" + strconv.Itoa(common_config.FenixTestDataSyncServer_port)

	fenixTestDataSyncServerClient fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixClientTestDataGrpcServicesServer struct {
	fenixClientTestDataSyncServerGrpcApi.UnimplementedFenixClientTestDataGrpcServicesServer
}

//TODO FIXA DENNA PATH, HMMM borde köra i DB framöver
// For now hardcoded MerklePath
var merkleFilterPath string = "AccountEnvironment/ClientJuristictionCountryCode/MarketSubType/MarketName/" //SecurityType/"

var testFile = "../../data/FenixRawTestdata_14rows_211216.csv"
