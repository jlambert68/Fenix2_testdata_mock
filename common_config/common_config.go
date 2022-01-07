package common_config

import "github.com/sirupsen/logrus"

// gRPC-ports
const QmlServer_address = "127.0.0.1"
const QmlServer_port = 5998

const TestInstructionBackendServer_address = "127.0.0.1"
const TestInstructionBackendServer_initial_port = 6660

// Logrus debug level

//const LoggingLevel = logrus.DebugLevel
//const LoggingLevel = logrus.InfoLevel
const LoggingLevel = logrus.DebugLevel // InfoLevel
