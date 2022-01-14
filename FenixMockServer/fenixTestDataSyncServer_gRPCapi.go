package FenixMockServer

import (
	"Fenix2_testdata_mock/common_config"
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// *********************************************************************
//Fenix client can check if Fenix Testdata sync server is alive with this service
func (s *FenixTestDataGrpcServicesServer) AreYouAlive(ctx context.Context, emptyParameter *fenixTestDataSyncServerGrpcApi.EmptyParameter) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
	}).Debug("Incoming 'AreYouAlive'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84",
	}).Debug("Outgoing 'AreYouAlive'")

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: "I'am alive"}, nil
}

// *********************************************************************
// Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) SendMerkleHash(ctx context.Context, merkleHashMessage *fenixTestDataSyncServerGrpcApi.MerkleHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "a55f9c82-1d74-44a5-8662-058b8bc9e48f",
	}).Debug("Incoming 'SendMerkleHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "27fb45fe-3266-41aa-a6af-958513977e28",
	}).Debug("Outgoing 'SendMerkleHash'")

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForClient(*merkleHashMessage)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleTree(ctx context.Context, merkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cffc25f0-b0e6-407a-942a-71fc74f831ac",
	}).Debug("Incoming 'SendMerkleTree'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "61e2c28d-b091-442a-b7f8-d2502d9547cf",
	}).Debug("Outgoing 'SendMerkleTree'")

	// Convert the merkleTree into a DataFrame object
	merkleTreeAsDataFrame := fenixTestDataSyncServerObject.convertgRpcMerkleTreeMessageToDataframe(*merkleTreeMessage)

	// Verify MerkleTree
	clientsMerkleRootHash := common_config.ExtractMerkleRootHashFromMerkleTree(merkleTreeAsDataFrame)
	recalculatedMerkleRootHash := common_config.CalculateMerkleHashFromMerkleTree(merkleTreeAsDataFrame)

	// Something is wrong with clients hash computation
	if clientsMerkleRootHash != recalculatedMerkleRootHash {
		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: false, Comments: "There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes"}, nil
	}

	// Save the MerkleTree Dataframe message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleTreeForClient(merkleTreeMessage.TestDataClientGuid, merkleTreeAsDataFrame)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaders(ctx context.Context, testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aee48999-12ad-4bb7-bc8a-96b62a8eeedf",
	}).Debug("Incoming 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "ca0b58a8-6d56-4392-8751-45906670e86b",
	}).Debug("Outgoing 'SendTestDataHeaders'")

	// Convert gRPC-message into other 'format'
	headerHash, headerItems := fenixTestDataSyncServerObject.convertgRpcHeaderMessageToStringArray(*testDataHeaderMessage)

	// Validate HeaderHash
	computedHeaderHash := common_config.HashValues(headerItems)
	if computedHeaderHash != headerHash {
		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: false, Comments: "Header hash is not correct computed. Expected '" + computedHeaderHash + "' as HeaderHash"}, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashsForClient(testDataHeaderMessage.TestDataClientGuid, headerHash)
	_ = fenixTestDataSyncServerObject.saveCurrentHeadersForClient(testDataHeaderMessage.TestDataClientGuid, headerItems)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil

}

// *********************************************************************
// Fenix client can send TestData rows to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataRows(ctx context.Context, testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2b1c8752-eb84-4c15-b8a7-22e2464e5168",
	}).Debug("Incoming 'SendTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "755e8b4f-f184-4277-ad41-e041714c2ca8",
	}).Debug("Outgoing 'SendTestDataRows'")

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) RegisterTestDataClient(ctx context.Context, testDataClientInformationMessage *fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5133b80b-6f3a-4562-9e62-1b3ceb169cc1",
	}).Debug("Incoming 'RegisterTestDataClient'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "316dcd7e-2229-4a82-b15b-0f808c2dd8aa",
	}).Debug("Outgoing 'RegisterTestDataClient'")

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

/*
func (s *FenixTestDataGrpcServicesServer) mustEmbedUnimplementedFenixTestDataGrpcServicesServer() {
	//TODO implement me
	panic("implement me")
}


*/
