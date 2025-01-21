package gograph

const (
	BaseGraphUrl = "https://graph.microsoft.com/v1.0/"
)

const (
	Success int = iota << 4 // Successful Exit Code

	ApiRequestCreationFailure // Business Exception Exit Codes
	ApiCallFailure
	HttpResponseReadFailure
	ModelCreationFailure

	UnimplementedFeature int = iota << 8 // System Failure Exit Codes
	CliInitialiseFailure
	InteractiveModeFailure
	ArgParseFailure
	EnvFileLoadFailure
	EnvFileSaveFailure
	EnvVariableLoadFailure
)
