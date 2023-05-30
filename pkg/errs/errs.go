package errs

var ErrFmt = "err: [%v]"
var ErrFmtMore = "err: [%v - %v]"

// common
var (
	ErrLoadingTimeZone  = _new("CMN000", "error loading timezone data")
	ErrMarshalingJson   = _new("CMN001", "error marshaling json")
	ErrUnmarshalingJson = _new("CMN002", "error unmarshaling json")
)

// pkg/config
var (
	ErrCreatingParamStore    = _new("CFG002", "unable to create param store service")
	ErrUnknownConfigProvider = _new("CFG003", "error unknown config provider")
	ErrGettingEnv            = _new("CFG004", "error unknown get env variables")

	ErrGettingEnvWebSocketURL  = _new("CFG005", "error unknown get env variable Web Socket URL")
	ErrGettingEnvWebSocketPort = _new("CFG006", "error unknown get env variable Web Socket Port")
	ErrGettingEnvTopic         = _new("CFG007", "error unknown get env variable Topic")
)

// pkg/stream
var (
	ErrNatsConnection = _new("NAT001", "error starting nats connection")
)

// pkg/websocket
var (
	ErrWebSocketConnection   = _new("WBS001", "error starting websocket connection")
	ErrWebSocketWriteMessage = _new("WBS002", "error websocket write message")
	ErrWebSocketReadMessage  = _new("WBS003", "error websocket read message")
	ErrWebSocketUpgrader     = _new("WBS004", "error websocket upgrader connection")
	ErrWebSocketSendingMsg   = _new("WBS005", "error send message to websocket")
)

// validations
var (
	ErrValidation = _new("VAL000", "error on validation")
)
