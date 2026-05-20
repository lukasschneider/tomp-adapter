package tomperr

type TompModule int

const (
	OperatorInformation TompModule = 1
	Planning            TompModule = 2
	Booking             TompModule = 3
	TripExecution       TompModule = 4
	Support             TompModule = 5
	Payment             TompModule = 6
	General             TompModule = 7
	CustomerManagement  TompModule = 8
)

// Technical errors
// https://github.com/TOMP-WG/TOMP-API/wiki/Error-handling-in-TOMP
const (
	// - Errorcode: x001
	// - Type: Missing
	// - Title: Field: {}, Reason: {}
	Missing = 1
	// - Errorcode: x002
	// - Type: Invalid
	// - Title: Field: {}, Reason: {}
	Invalid = 2
	// - Errorcode: x004
	// - Type: Illegal operation
	// - Title: Operation {} is illegal in current status.
	IllegalOperation = 4
	// - Errorcode: x005
	// - Type: Technical issue
	// - Title: Internal technical problem, contact support.
	TechnicalIssue = 5
)

// Functional errors
// https://github.com/TOMP-WG/TOMP-API/wiki/Error-handling-in-TOMP
const (
	// - Errorcode: x202
	// - Type: Expired
	// - Example detail: Availability of booking {} expired.
	// - Description: When the pending timer is finished and a new booking must be made in order to book this asset.
	Expired = 202
	// - Errorcode: x203
	// - Type: In use
	// - Example detail: Booking {} has started.
	// - Description: It is no longer possible to edit the details of this booking since the asset is already in use.
	InUse = 203
	// - Errorcode: x204
	// - Type: Not found
	// - Example detail: Booking {} not found.
	// - Description: The specified booking, leg, customer etc. cannot be found.
	NotFound = 204
)
