package telemetry

/*
CheckRequestTupleKeyInterface is a simplified interface that defines the methods the CheckRequestTupleKey struct implements, relevant to the context of the telemetry package.
*/
type CheckRequestTupleKeyInterface interface {
	GetUser() *string
}

/*
CheckRequestInterface is a simplified interface that defines the methods the CheckRequest struct implements, relevant to the context of the telemetry package.
*/
type CheckRequestInterface interface {
	GetTupleKey() CheckRequestTupleKeyInterface
	RequestAuthorizationModelIdInterface
}

/*
RequestAuthorizationModelIdInterface is a generic interface that defines the GetAuthorizationModelId() method a Request struct implements, relevant to the context of the telemetry package.
*/
type RequestAuthorizationModelIdInterface interface {
	GetAuthorizationModelId() *string
}
