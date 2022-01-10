package FenixMockServer

import (
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/go-gota/gota/dataframe"
)

/*
// Initiate channels for "incomming gRPC-messages"
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) InitateProcessEngineChannels() {

	make(TestDataClientInformationMessageChannel, fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage, 1)

	// 'MerkleHashMessage' from 'gRPC-SendMerkleHash'
	var MerkleHashMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleHashMessage

	// 'MerkleTreeMessage' from 'gRPC-SendMerkleTree'
	var MerkleTreeMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

	// 'TestDataHeaderMessage' from 'gRPC-SendTestDataHeaders'
	var TestDataHeaderMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage

	// 'MerkleTreeMessage' from 'gRPC-SendTestDataRows'
	var MerkleTreeMessageMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

}
*/

// Convert gRPC-MerkleTree message into a DataFrame object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcMerkleTreeMessageToDataframe(merkleTreeMessage fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) dataframe.DataFrame {

	var myMerkleTree []MerkleTree_struct

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	merkleTreeNodes := merkleTreeMessage.MerkleTreeNodes

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, merkleTreeNode := range merkleTreeNodes {
		myMerkleTreeRow := MerkleTree_struct{
			MerkleLevel:     int(merkleTreeNode.MerkleLevel),
			MerklePath:      merkleTreeNode.MerklePath,
			MerkleHash:      merkleTreeNode.MerkleHash,
			MerkleChildHash: merkleTreeNode.MerkleChildHash,
		}
		myMerkleTree = append(myMerkleTree, myMerkleTreeRow)

	}

	df := dataframe.LoadStructs(myMerkleTree)

	return df
}

// Convert gRPC-Header message into string and string array objects
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcHeaderMessageToStringArray(testDataHeaderMessage fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage) (headerHash string, headersItems []string) {

	// Extract  HeaderHash
	headerHash = testDataHeaderMessage.HeaderRowHash

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	testDataHeaderItems := testDataHeaderMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		headersItems = append(headersItems, headerItem.HeaderValueAsString)

	}

	return headerHash, headersItems
}