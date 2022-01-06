package FenixMockServer

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
)

// Used for checking if Backend Server is alive
func (s *TestInstructionBackendServer) AreYouAlive(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.AckNackResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Debug("Incoming: 'AreYouAlive'")

	// Temp-solution for handling that QML-server is up and running
	testInstructionBackendObject.qmlServerHasConnected = true

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Debug("Leaving 'AreYouAlive'")
	return &backend_server_grpc_api.AckNackResponse{Acknack: true, Comments: "I'am alive, from " + testInstructionBackendObject.uuid}, nil
}

// *********************************************************************
// Generates a guid in string format to be sent front end
func (s *TestInstructionBackendServer) GenerateGuid(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.GuidResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'GenerateGuid'")

	var returnGuid *backend_server_grpc_api.GuidResponse
	newGuid, _ := uuid.NewUUID()

	returnGuid = &backend_server_grpc_api.GuidResponse{
		Guid:                 newGuid.String(),
		Acknack:              true,
		Comments:             "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id":         "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
		"returnGuid": returnGuid,
	}).Debug("Leaving 'GenerateGuid'")

	return returnGuid, nil

}

// *********************************************************************
// Load stored data of type found in ingoing message and send back towards frontend, via QML-server
func (s *TestInstructionBackendServer) LoadFromServer(ctx context.Context, qmlModelTypeToAndFromBackendMessage *backend_server_grpc_api.QMLModelTypeToAndFromBackendMessage) (*backend_server_grpc_api.QmlModeToLoadlFromServerResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id":                               "649a3945-d24d-42fb-b8b3-c01bf16c552a",
		"QMLModelTypeToAndFromBackendType": qmlModelTypeToAndFromBackendMessage.QMLModelTypeToAndFromBackendType,
	}).Debug("Incoming 'LoadFromServer'")

	var returnMessage *backend_server_grpc_api.QmlModeToLoadlFromServerResponse

	type ListElementStruct struct {
		Name        string `json:"name"`
		GUID        string `json:"guid"`
		ReadyForUse bool   `json:"readyForUse"`
		Activated   bool   `json:"activated"`
		Description string `json:"description"`
	}

	type dataForPlugInStruct struct {
		ListElement []ListElementStruct `json:"ListElement"`
	}

	var dataForPlugin dataForPlugInStruct
	var listElementData []ListElementStruct

	switch qmlModelTypeToAndFromBackendMessage.QMLModelTypeToAndFromBackendType {
	case backend_server_grpc_api.QMLModelTypeToAndFromBackend_PluginModel:

		// Append First Mock-Element
		listElement1 := &ListElementStruct{
			Name:        "Send UTR",
			GUID:        "2243085a-feee-4ae7-8ccf-03f69c0704a4",
			ReadyForUse: true,
			Activated:   true,
			Description: "Sends UTRs via MQ towards Custody Cash",
		}
		listElementData = append(listElementData, *listElement1)

		// Appned Second Mock-ELement
		listElement2 := &ListElementStruct{
			Name:        "Validate New Pacs008",
			GUID:        "d456499c-2ad1-4677-8e1d-909a7ecab560",
			ReadyForUse: true,
			Activated:   true,
			Description: "Validates that a newly created Pacs008 has been sent to CMaaS",
		}
		listElementData = append(listElementData, *listElement2)

	case backend_server_grpc_api.QMLModelTypeToAndFromBackend_DomainModel:
		// Append First Mock-ELement
		listElement1 := &ListElementStruct{
			Name:        "Custody Cash",
			GUID:        "2243085a-feee-4ae7-8ccf-03f69c0704a4",
			ReadyForUse: true,
			Activated:   true,
			Description: "All test regarding Custody Cash",
		}
		listElementData = append(listElementData, *listElement1)

		// Appned Second Mock-Element
		listElement2 := &ListElementStruct{
			Name:        "Custody Arrangement",
			GUID:        "d456499c-2ad1-4677-8e1d-909a7ecab560",
			ReadyForUse: false,
			Activated:   false,
			Description: "All tests regarding Custody Arrangement",
		}
		listElementData = append(listElementData, *listElement2)
	default:
		returnMessage = &backend_server_grpc_api.QmlModeToLoadlFromServerResponse{
			JsonStringForPluginQmlModel:  "",
			QmlModelTypeToAndFromBackend: qmlModelTypeToAndFromBackendMessage.QMLModelTypeToAndFromBackendType,
			Acknack:                      false,
			Comments:                     "Unknown QmlModelType",
		}

		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"id":         "2fddd953-ce75-4d11-a04a-307c34ea49c7",
			"returnGuid": returnMessage,
		}).Debug("Unknown QmlModelType")

		return returnMessage, nil
	}

	dataForPlugin.ListElement = listElementData

	jsonToSend, _ := json.MarshalIndent(dataForPlugin, "", " ")
	jsonToSendAsString := string(jsonToSend)

	returnMessage = &backend_server_grpc_api.QmlModeToLoadlFromServerResponse{
		JsonStringForPluginQmlModel:  jsonToSendAsString,
		QmlModelTypeToAndFromBackend: qmlModelTypeToAndFromBackendMessage.QMLModelTypeToAndFromBackendType,
		Acknack:                      true,
		Comments:                     "",
	}

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id":         "c2c5fdd3-1c31-4694-8428-b8efd448c561",
		"returnGuid": returnMessage,
	}).Debug("Leaving 'LoadPluginModelFromServer'")

	return returnMessage, nil

}

// *********************************************************************
// Save data, of type found in in data message, that was sent from GUI, via QML-server
func (s *TestInstructionBackendServer) SaveToServer(ctx context.Context, pluginQmlModelToServerRequest *backend_server_grpc_api.QmlModelToSaveAtServerRequest) (*backend_server_grpc_api.AckNackResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id": "761ab856-6c7f-41ef-b951-534ecfee3b58",
	}).Debug("Incoming 'SavePluginModelToServer'")

	var returnMessage *backend_server_grpc_api.AckNackResponse

	returnMessage = &backend_server_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "",
	}

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"id":         "5ffb910b-ab1f-4b1a-b488-1e01d3ab834d",
		"returnGuid": returnMessage,
	}).Debug("Leaving 'SavePluginModelToServer'")

	return returnMessage, nil

}
