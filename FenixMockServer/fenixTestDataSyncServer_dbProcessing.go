package FenixMockServer

import (
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/go-gota/gota/dataframe"
)

// ***** Before I implement a DB everthing will be stored in variables *****
/*
var dbCurrentMerkleHashForClient string
var dbCurrentMerkleTreeForClient dataframe.DataFrame
var dbCurrentHeaderHashsForClient string
var dbCurrentHeadersForClient []string
*/

var dbData tempDB

type tempTestDataRowStruct struct {
	rowHash       string
	testDataValue []string
}

type tempDBDataStruct struct {
	merkleHash   string
	merkleTree   dataframe.DataFrame
	headerHash   string
	headers      []string
	testDataRows []tempTestDataRowStruct
}
type tempDB struct {
	serverData tempDBDataStruct
	clientData tempDBDataStruct
}

type MerkleTree_struct struct {
	MerkleLevel     int
	MerklePath      string
	MerkleHash      string
	MerkleChildHash string
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleHashForClient(testDataClientGuid string) string {

	var currentMerkleHashForClient string

	currentMerkleHashForClient = dbData.clientData.merkleHash

	return currentMerkleHashForClient
}

// Save current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleHashForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.clientData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleHashForServer(testDataClientGuid string) string {

	var currentMerkleHashForServer string

	currentMerkleHashForServer = dbData.serverData.merkleHash

	return currentMerkleHashForServer
}

// Save current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleHashForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.serverData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleTreeForClient(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForClient dataframe.DataFrame

	currentMerkleTreeForClient = dbData.clientData.merkleTree

	return currentMerkleTreeForClient
}

// Save current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleTreeForClient(merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	dbData.clientData.merkleTree = merkleTreeDataFrameMessage

	/*
		m := jsonpb.Marshaler{}
		js, err := m.MarshalToString(&merkleTreeMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(js)
		}
	*/

	return true
}

// Retrieve current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleTreeForServer(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForServer dataframe.DataFrame

	currentMerkleTreeForServer = dbData.serverData.merkleTree

	return currentMerkleTreeForServer
}

// Save current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleTreeForServer(merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	dbData.serverData.merkleTree = merkleTreeDataFrameMessage

	return true
}

// Retrieve current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeaderHashsForClient(testDataClientGuid string) string {

	var currentHeaderHashsForClient string

	currentHeaderHashsForClient = dbData.clientData.headerHash

	return currentHeaderHashsForClient
}

// Save current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeaderHashsForClient(currentHeaderHashsForClient string) bool {

	dbData.clientData.headerHash = currentHeaderHashsForClientcurrentHeaderHashsForClient

	return true
}

// Retrieve currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeadersForClient(testDataClientGuid string) []string {

	var currentHeadersForClient []string

	currentHeadersForClient = dbData.clientData.headers

	return currentHeadersForClient
}

// Save currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeadersForClient(testDataHeaderItems []string) bool {

	dbData.clientData.headers = testDataHeaderItems

	return true
}

// Retrieve current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeaderHashForServer(testDataClientGuid string) string {

	var currentHeaderHashsForServer string

	currentHeaderHashsForServer = dbData.serverData.headerHash

	return currentHeaderHashsForServer
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeadersForServer(testDataClientGuid string) []string {

	var currentHeadersForServer []string

	currentHeadersForServer = dbData.serverData.headers

	return currentHeadersForServer
}
