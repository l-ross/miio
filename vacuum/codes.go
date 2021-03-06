package vacuum

type StatusCode int

const (
	StatusCodeUnknown       StatusCode = -1
	StatusCodeInitiating    StatusCode = 1
	StatusCodeSleeping      StatusCode = 2
	StatusCodeIdle          StatusCode = 3
	StatusCodeRemoteControl StatusCode = 4
	StatusCodeCleaning      StatusCode = 5
	StatusCodeReturningDock StatusCode = 6
	StatusCodeManualMode    StatusCode = 7
	StatusCodeCharging      StatusCode = 8
	StatusCodeChargingError StatusCode = 9
	StatusCodePaused        StatusCode = 10
	StatusCodeSpotCleaning  StatusCode = 11
	StatusCodeInError       StatusCode = 12
	StatusCodeShuttingDown  StatusCode = 13
	StatusCodeUpdating      StatusCode = 14
	StatusCodeDocking       StatusCode = 15
	StatusCodeGoTo          StatusCode = 16
	StatusCodeZoneClean     StatusCode = 17
	StatusCodeRoomClean     StatusCode = 18
	StatusCodeFullyCharged  StatusCode = 100
)

type ErrorCode int

const (
	ErrorCodeUnknown                    ErrorCode = -1
	ErrorCodeNoError                    ErrorCode = 0
	ErrorCodeLaserSensorFault           ErrorCode = 1
	ErrorCodeCollisionSensorFault       ErrorCode = 2
	ErrorCodeWheelFloating              ErrorCode = 3
	ErrorCodeCliffSensorFault           ErrorCode = 4
	ErrorCodeMainBrushBlocked           ErrorCode = 5
	ErrorCodeSideBrushBlocked           ErrorCode = 6
	ErrorCodeWheelBlocked               ErrorCode = 7
	ErrorCodeDeviceStuck                ErrorCode = 8
	ErrorCodeDustBinMissing             ErrorCode = 9
	ErrorCodeFilterBlocked              ErrorCode = 10
	ErrorCodeMagneticFieldDetected      ErrorCode = 11
	ErrorCodeLowBattery                 ErrorCode = 12
	ErrorCodeChargingProblem            ErrorCode = 13
	ErrorCodeBatteryFailure             ErrorCode = 14
	ErrorCodeWallSensorFault            ErrorCode = 15
	ErrorCodeUnevenSurface              ErrorCode = 16
	ErrorCodeSideBrushFailure           ErrorCode = 17
	ErrorCodeSuctionFanFailure          ErrorCode = 18
	ErrorCodeUnpoweredChargingStation   ErrorCode = 19
	ErrorCodeLaserPressureSensorProblem ErrorCode = 21
	ErrorCodeChargeSensorProblem        ErrorCode = 22
	ErrorCodeDockProblem                ErrorCode = 23
	ErrorCodeInvisibleWallDetected      ErrorCode = 24
	ErrorCodeBinFull                    ErrorCode = 254
	ErrorCodeInternalError              ErrorCode = 255
)

type FeatureCode int

const (
	FeatureCodeClean               FeatureCode = 103
	FeatureCodeFSEndPoint          FeatureCode = 111
	FeatureCodeAutoSplitSegments   FeatureCode = 112
	FeatureCodeDeleteMap           FeatureCode = 113
	FeatureCodeOrderSegmentClean   FeatureCode = 114
	FeatureCodeSpotClean           FeatureCode = 115
	FeatureCodeMapSegment          FeatureCode = 116
	FeatureCodeLEDStatusSwitch     FeatureCode = 119
	FeatureCodeMultiFloorSupported FeatureCode = 120
	FeatureCodeFetchTimerSummary   FeatureCode = 122
	FeatureCodeOrdersClean         FeatureCode = 123
	FeatureCodeAnalysis            FeatureCode = 124
	FeatureCodeRemote              FeatureCode = 125
)
