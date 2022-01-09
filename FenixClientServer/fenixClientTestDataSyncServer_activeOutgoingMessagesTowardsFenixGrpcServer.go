package FenixClientServer

import (
	"Fenix2_testdata_mock/common_config"
	fenixClientTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixClientTestDataSyncServerGrpcApi/proto"
	fenixTestDataSyncServerGrpcApi "Fenix2_testdata_mock/grpc_api/fenixTestDataSyncServerGrpcApi/proto"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Set upp connection and Dial to FenixTestDataSyncServer
func SetConnectionToFenixTestDataSyncServer() {

	var err error

	// Set up connection to FenixTestDataSyncServer
	remoteFenixTestDataSyncServerConnection, err = grpc.Dial(fenixTestDataSyncServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixTestDataSyncServer_address_to_dial": fenixTestDataSyncServer_address_to_dial,
			"error message": err,
		}).Error("Did not connect to FenixTestDataSyncServer via gRPC")
		//os.Exit(0)
	} else {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixTestDataSyncServer_address_to_dial": fenixTestDataSyncServer_address_to_dial,
		}).Info("gRPC connection OK to FenixTestDataSyncServer")

		// Creates a new Clients
		fenixTestDataSyncServerClient = fenixTestDataSyncServerGrpcApi.NewFenixTestDataGrpcServicesClient(remoteFenixTestDataSyncServerConnection)

	}
}

// Get the highest ProtoFileVersionEnumeration
func getHighestProtoFileVersion() int32 {

	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestDataSyncServerGrpcApi.CurrentTestDataProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

// Generate the current MerkleTree for Testdata supported by the client
func getCurrentTestDataMerkleTree() fenixTestDataSyncServerGrpcApi.MerkleTreeMessage {

	var merkleTreeMessage fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

	return merkleTreeMessage
}

func (fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct) RegisterTestDataClient() fenixTestDataSyncServerGrpcApi.AckNackResponse {

	// Set up variables to be sent to FenixTestDataSyncServer
	TestDataClientInformationMessage := fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage{
		TestDataClientGuid:           common_config.FenicClientTestDataSyncServer_TestDataClientGuid,
		TestDomainGuid:               common_config.FenicClientTestDataSyncServer_DomainGuid,
		TestDomainName:               common_config.FenicClientTestDataSyncServer_DomainName,
		TestDataClientIpAddress:      common_config.FenixClientTestDataSyncServer_address,
		TestDataClientPort:           string(common_config.FenixClientTestDataSyncServer_initial_port),
		ProtoFileVersionUsedByCLient: fenixTestDataSyncServerGrpcApi.CurrentTestDataProtoFileVersionEnum(getHighestProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixTestDataSyncServerClient.RegisterTestDataClient(ctx, &TestDataClientInformationMessage)

	// Shouldn't happen
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "6b080a23-4e06-4d16-8295-a67ba7115a56",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixTestDataSyncServer for 'RegisterTestDataClient'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "44671efb-e24d-450a-acba-006cc248d058",
				"Message from FenixTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixTestDataSyncServer for 'RegisterTestDataClient'")
		}
	}
	return *returnMessage
}

func (fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct) SendMerkleHash() fenixTestDataSyncServerGrpcApi.AckNackResponse {

	merkleRootHash, _ := common_config.LoadAndProcessFile("data/FenixRawTestdata_14rows_211216.csv")

	// Set up variables to be sent to FenixTestDataSyncServer
	merkleHashMessage := fenixTestDataSyncServerGrpcApi.MerkleHashMessage{
		TestDataClientGuid: common_config.FenicClientTestDataSyncServer_TestDataClientGuid,
		MerkleHash:         merkleRootHash}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixTestDataSyncServerClient.SendMerkleHash(ctx, &merkleHashMessage)

	// Shouldn't happen
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "69a62788-b798-471a-bb8d-7fa1cec0f485",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendMerkleHash'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "fb923a55-136e-481e-9c30-d7d7019e17e3",
				"Message from FenixTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendMerkleHash'")
		}
	}
	return *returnMessage
}

func (fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct) SendMerkleTree() fenixTestDataSyncServerGrpcApi.AckNackResponse {

	var merkleTreeNodeMessages []*fenixTestDataSyncServerGrpcApi.MerkleTreeNodeMessage

	// Set up variables to be sent to FenixTestDataSyncServer
	_, merkleTree := common_config.LoadAndProcessFile("data/FenixRawTestdata_14rows_211216.csv")

	merkleTreeNRows := merkleTree.Nrow()
	for rowCounter := 0; rowCounter < merkleTreeNRows; rowCounter++ {
		merkleLevel, _ := merkleTree.Elem(rowCounter, 0).Int()
		merkleTreeNodeMessage := &fenixTestDataSyncServerGrpcApi.MerkleTreeNodeMessage{
			MerkleLevel:     int64(merkleLevel),
			MerklePath:      merkleTree.Elem(rowCounter, 1).String(),
			MerkleHash:      merkleTree.Elem(rowCounter, 2).String(),
			MerkleChildHash: merkleTree.Elem(rowCounter, 3).String(),
		}
		merkleTreeNodeMessages = append(merkleTreeNodeMessages, merkleTreeNodeMessage)
	}
	merkleTreeMessage := &fenixTestDataSyncServerGrpcApi.MerkleTreeMessage{
		TestDataClientGuid: common_config.FenicClientTestDataSyncServer_TestDataClientGuid,
		MerkleTreeNodes:    merkleTreeNodeMessages}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixTestDataSyncServerClient.SendMerkleTree(ctx, merkleTreeMessage)

	// Shouldn't happen
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "c8a66468-17ca-4e0a-942b-a9ec9b246c82",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendMerkleTree'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "d8225481-d28c-426c-9cdb-986678001e5c",
				"Message from FenixTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendMerkleTree'")
		}
	}
	return *returnMessage
}

func (fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct) SendTestDataHeaders() fenixTestDataSyncServerGrpcApi.AckNackResponse {

	var testDataHeaderItemsMessage []*fenixTestDataSyncServerGrpcApi.TestDataHeaderItemMessage
	_, merkleTree := common_config.LoadAndProcessFile("data/FenixRawTestdata_14rows_211216.csv")

	// Get all headers as a string array
	headers := merkleTree.Names()

	// Hash the header into a single hash
	headerHash := common_config.HashValues(headers)

	// Create variables to be sent to FenixTestDataSyncServer
	for _, header := range headers {
		testDataHeaderItemMessage := &fenixTestDataSyncServerGrpcApi.TestDataHeaderItemMessage{HeaderValueAsString: header}
		testDataHeaderItemsMessage = append(testDataHeaderItemsMessage, testDataHeaderItemMessage)
	}

	// Header message to be set to  TestDataSyncServer
	testDataHeaderMessage := &fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage{
		TestDataClientGuid:  common_config.FenicClientTestDataSyncServer_TestDataClientGuid,
		HeaderRowHash:       headerHash,
		TestDataHeaderItems: testDataHeaderItemsMessage,
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixTestDataSyncServerClient.SendTestDataHeaders(ctx, testDataHeaderMessage)

	// Shouldn't happen
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "5644eeb0-7e95-4b42-ae2a-1fafdf926f9d",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendTestDataHeaders'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "3902e0d2-d28a-40e4-8aa8-553d31ac3b78",
				"Message from FenixTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendTestDataHeaders'")
		}
	}
	return *returnMessage
}

func (fenixClientTestDataSyncServerObject *fenixClientTestDataSyncServerObject_struct) SendTestDataRows(merklePathsMessage *fenixClientTestDataSyncServerGrpcApi.MerklePathsMessage) fenixTestDataSyncServerGrpcApi.AckNackResponse {

	var merkleTreeNodeMessages []*fenixTestDataSyncServerGrpcApi.MerkleTreeNodeMessage

	_, merkleTree := common_config.LoadAndProcessFile("data/FenixRawTestdata_14rows_211216.csv")

	// Filter out only rows with correct MerklePath
	MerkleTreeRowsToSendToTestDataServer := merkleTree.Filter(
		dataframe.F{
			Colname:    "MerklePath",
			Comparator: series.In,
			Comparando: merklePathsMessage.MerklePath})

	// Create data to be sent to TestDataSyncServer
	merkleTreeNRows := MerkleTreeRowsToSendToTestDataServer.Nrow()
	for rowCounter := 0; rowCounter < merkleTreeNRows; rowCounter++ {
		merkleLevel, _ := MerkleTreeRowsToSendToTestDataServer.Elem(rowCounter, 0).Int()
		merkleTreeNodeMessage := &fenixTestDataSyncServerGrpcApi.MerkleTreeNodeMessage{
			MerkleLevel:     int64(merkleLevel),
			MerklePath:      MerkleTreeRowsToSendToTestDataServer.Elem(rowCounter, 1).String(),
			MerkleHash:      MerkleTreeRowsToSendToTestDataServer.Elem(rowCounter, 2).String(),
			MerkleChildHash: MerkleTreeRowsToSendToTestDataServer.Elem(rowCounter, 3).String(),
		}
		merkleTreeNodeMessages = append(merkleTreeNodeMessages, merkleTreeNodeMessage)
	}
	merkleTreeMessage := &fenixTestDataSyncServerGrpcApi.MerkleTreeMessage{
		TestDataClientGuid: common_config.FenicClientTestDataSyncServer_TestDataClientGuid,
		MerkleTreeNodes:    merkleTreeNodeMessages}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixTestDataSyncServerClient.SendTestDataRows(ctx, merkleTreeMessage)

	// Shouldn't happen
	if err != nil {
		fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "b457b233-41f9-4b3d-9f1e-00782b467045",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendTestDataRows'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixClientTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "c1f6a351-fb7e-4759-81a7-04ec61b74e59",
				"Message from FenixTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixTestDataSyncServer for 'SendTestDataRows'")
		}
	}
	return *returnMessage
}
