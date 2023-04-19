package client

import (
	_context "context"
	"encoding/json"
	"fmt"
	"github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/credentials"
	"golang.org/x/sync/errgroup"
	"math"
	_nethttp "net/http"
)

var (
	_ _context.Context
)

var DEFAULT_MAX_METHOD_PARALLEL_REQS = int32(15)

type ClientConfiguration struct {
	openfga.Configuration
	ApiScheme            string                   `json:"apiScheme,omitempty"`
	ApiHost              string                   `json:"apiHost,omitempty"`
	StoreId              string                   `json:"storeId,omitempty"`
	AuthorizationModelId *string                  `json:"authorization_model_id,omitempty"`
	Credentials          *credentials.Credentials `json:"credentials,omitempty"`
	DefaultHeaders       map[string]string        `json:"defaultHeader,omitempty"`
	UserAgent            string                   `json:"userAgent,omitempty"`
	Debug                bool                     `json:"debug,omitempty"`
	HTTPClient           *_nethttp.Client
	RetryParams          *openfga.RetryParams
}

func newClientConfiguration(cfg *openfga.Configuration) ClientConfiguration {
	return ClientConfiguration{
		ApiScheme:      cfg.ApiScheme,
		ApiHost:        cfg.ApiHost,
		StoreId:        cfg.StoreId,
		Credentials:    cfg.Credentials,
		DefaultHeaders: cfg.DefaultHeaders,
		UserAgent:      cfg.UserAgent,
		Debug:          cfg.Debug,
		RetryParams:    cfg.RetryParams,
	}
}

type OpenFgaClient struct {
	Config ClientConfiguration
	SdkClient
	openfga.APIClient
}

func NewSdkClient(cfg *ClientConfiguration) (*OpenFgaClient, error) {
	apiConfiguration, err := openfga.NewConfiguration(openfga.Configuration{
		ApiScheme:      cfg.ApiScheme,
		ApiHost:        cfg.ApiHost,
		StoreId:        cfg.StoreId,
		Credentials:    cfg.Credentials,
		DefaultHeaders: make(map[string]string),
		UserAgent:      openfga.GetSdkUserAgent(),
		Debug:          cfg.Debug,
		RetryParams:    cfg.RetryParams,
	})

	if err != nil {
		return nil, err
	}

	clientConfig := newClientConfiguration(apiConfiguration)
	clientConfig.AuthorizationModelId = cfg.AuthorizationModelId

	apiClient := openfga.NewAPIClient(apiConfiguration)

	return &OpenFgaClient{
		Config:    clientConfig,
		APIClient: *apiClient,
	}, nil
}

type ClientRequestOptions struct {
	MaxRetry    *int `json:"max_retry,omitempty"`
	MinWaitInMs *int `json:"min_wait_in_ms,omitempty"`
}

type AuthorizationModelIdOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

type ClientRequestOptionsWithAuthZModelId struct {
	ClientRequestOptions
	AuthorizationModelIdOptions
}

type ClientTupleKey struct {
	Object   string `json:"object,omitempty"`
	Relation string `json:"relation,omitempty"`
	User     string `json:"user,omitempty"`
}

func (tupleKey ClientTupleKey) ToTupleKey() openfga.TupleKey {
	return openfga.TupleKey{
		User:     openfga.PtrString(tupleKey.User),
		Relation: openfga.PtrString(tupleKey.Relation),
		Object:   openfga.PtrString(tupleKey.Object),
	}
}

type ClientPaginationOptions struct {
	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
}

func getPageSizeFromRequest(options *ClientPaginationOptions) *int32 {
	if options == nil {
		return nil
	}
	return options.PageSize
}
func getContinuationTokenFromRequest(options *ClientPaginationOptions) *string {
	if options == nil {
		return nil
	}
	return options.ContinuationToken
}

type SdkClient interface {
	/* Stores */
	ListStores(ctx _context.Context) SdkClientListStoresRequest
	ListStoresExecute(request SdkClientListStoresRequest) (openfga.ListStoresResponse, *_nethttp.Response, error)
	CreateStore(ctx _context.Context) SdkClientCreateStoreRequest
	CreateStoreExecute(request SdkClientCreateStoreRequest) (openfga.CreateStoreResponse, *_nethttp.Response, error)
	GetStore(ctx _context.Context) SdkClientGetStoreRequest
	GetStoreExecute(request SdkClientGetStoreRequest) (openfga.GetStoreResponse, *_nethttp.Response, error)
	DeleteStore(ctx _context.Context) SdkClientDeleteStoreRequest
	DeleteStoreExecute(request SdkClientDeleteStoreRequest) (ClientDeleteStoreResponse, *_nethttp.Response, error)

	/* Authorization Models */
	ReadAuthorizationModels(ctx _context.Context) SdkClientReadAuthorizationModelsRequest
	ReadAuthorizationModelsExecute(request SdkClientReadAuthorizationModelsRequest) (openfga.ReadAuthorizationModelsResponse, *_nethttp.Response, error)
	WriteAuthorizationModel(ctx _context.Context) SdkClientWriteAuthorizationModelRequest
	WriteAuthorizationModelExecute(request SdkClientWriteAuthorizationModelRequest) (openfga.WriteAuthorizationModelResponse, *_nethttp.Response, error)
	ReadAuthorizationModel(ctx _context.Context) SdkClientReadAuthorizationModelRequest
	ReadAuthorizationModelExecute(request SdkClientReadAuthorizationModelRequest) (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error)
	ReadLatestAuthorizationModel(ctx _context.Context) SdkClientReadLatestAuthorizationModelRequest
	ReadLatestAuthorizationModelExecute(request SdkClientReadLatestAuthorizationModelRequest) (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error)

	/* Relationship Tuples */
	ReadChanges(ctx _context.Context) SdkClientReadChangesRequest
	ReadChangesExecute(request SdkClientReadChangesRequest) (openfga.ReadChangesResponse, *_nethttp.Response, error)
	Read(ctx _context.Context) SdkClientReadRequest
	ReadExecute(request SdkClientReadRequest) (openfga.ReadResponse, *_nethttp.Response, error)
	Write(ctx _context.Context) SdkClientWriteRequest
	WriteExecute(request SdkClientWriteRequest) (ClientWriteResponse, error)
	WriteTuples(ctx _context.Context) SdkClientWriteTuplesRequest
	WriteTuplesExecute(request SdkClientWriteTuplesRequest) (ClientWriteResponse, error)
	DeleteTuples(ctx _context.Context) SdkClientDeleteTuplesRequest
	DeleteTuplesExecute(request SdkClientDeleteTuplesRequest) (ClientWriteResponse, error)

	/* Relationship Queries */
	Check(ctx _context.Context) SdkClientCheckRequest
	CheckExecute(request SdkClientCheckRequest) (openfga.CheckResponse, *_nethttp.Response, error)
	BatchCheck(ctx _context.Context) SdkClientBatchCheckRequest
	BatchCheckExecute(request SdkClientBatchCheckRequest) (ClientBatchCheckResponse, error)
	Expand(ctx _context.Context) SdkClientExpandRequest
	ExpandExecute(request SdkClientExpandRequest) (openfga.ExpandResponse, *_nethttp.Response, error)
	ListObjects(ctx _context.Context) SdkClientListObjectsRequest
	ListObjectsExecute(request SdkClientListObjectsRequest) (openfga.ListObjectsResponse, *_nethttp.Response, error)
	ListRelations(ctx _context.Context) SdkClientListRelationsRequest
	ListRelationsExecute(request SdkClientListRelationsRequest) (*ClientListRelationsResponse, error)

	/* Assertions */
	ReadAssertions(ctx _context.Context) SdkClientReadAssertionsRequest
	ReadAssertionsExecute(request SdkClientReadAssertionsRequest) (openfga.ReadAssertionsResponse, *_nethttp.Response, error)
	WriteAssertions(ctx _context.Context) SdkClientWriteAssertionsRequest
	WriteAssertionsExecute(request SdkClientWriteAssertionsRequest) (*_nethttp.Response, error)
}

func (client *OpenFgaClient) getAuthorizationModelId(authorizationModelId *string) *string {
	if authorizationModelId != nil {
		return authorizationModelId
	}
	return client.Config.AuthorizationModelId
}

/* Stores */

// / ListStores
type SdkClientListStoresRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientListStoresOptions
}

type ClientListStoresOptions struct {
	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
}

func (request SdkClientListStoresRequest) Options(options ClientListStoresOptions) SdkClientListStoresRequest {
	request.options = &options
	return request
}

func (request SdkClientListStoresRequest) Execute() (openfga.ListStoresResponse, *_nethttp.Response, error) {
	return request.Client.ListStoresExecute(request)
}

func (client *OpenFgaClient) ListStoresExecute(request SdkClientListStoresRequest) (openfga.ListStoresResponse, *_nethttp.Response, error) {
	req := client.OpenFgaApi.ListStores(request.ctx)
	pageSize := getPageSizeFromRequest((*ClientPaginationOptions)(request.options))
	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	continuationToken := getContinuationTokenFromRequest((*ClientPaginationOptions)(request.options))
	if continuationToken != nil {
		req.ContinuationToken(*continuationToken)
	}
	return req.Execute()
}

func (client *OpenFgaClient) ListStores(ctx _context.Context) SdkClientListStoresRequest {
	return SdkClientListStoresRequest{
		ctx:    ctx,
		Client: *client,
	}
}

// / CreateStore
type SdkClientCreateStoreRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientCreateStoreRequest
	options *ClientCreateStoreOptions
}

type ClientCreateStoreRequest struct {
	Name string `json:"name"`
}

type ClientCreateStoreOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (request SdkClientCreateStoreRequest) Options(options ClientCreateStoreOptions) SdkClientCreateStoreRequest {
	request.options = &options
	return request
}

func (request SdkClientCreateStoreRequest) Body(body ClientCreateStoreRequest) SdkClientCreateStoreRequest {
	request.body = &body
	return request
}

func (request SdkClientCreateStoreRequest) Execute() (openfga.CreateStoreResponse, *_nethttp.Response, error) {
	return request.Client.CreateStoreExecute(request)
}

func (client *OpenFgaClient) CreateStoreExecute(request SdkClientCreateStoreRequest) (openfga.CreateStoreResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.CreateStore(request.ctx).Body(openfga.CreateStoreRequest{
		Name: request.body.Name,
	}).Execute()
}

func (client *OpenFgaClient) CreateStore(ctx _context.Context) SdkClientCreateStoreRequest {
	return SdkClientCreateStoreRequest{
		Client: *client,
		ctx:    ctx,
	}
}

// / GetStore
type SdkClientGetStoreRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientGetStoreOptions
}

type ClientGetStoreOptions struct {
}

func (request SdkClientGetStoreRequest) Options(options ClientGetStoreOptions) SdkClientGetStoreRequest {
	request.options = &options
	return request
}

func (request SdkClientGetStoreRequest) Execute() (openfga.GetStoreResponse, *_nethttp.Response, error) {
	return request.Client.GetStoreExecute(request)
}

func (client *OpenFgaClient) GetStoreExecute(request SdkClientGetStoreRequest) (openfga.GetStoreResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.GetStore(request.ctx).Execute()
}

func (client *OpenFgaClient) GetStore(ctx _context.Context) SdkClientGetStoreRequest {
	return SdkClientGetStoreRequest{
		Client: *client,
		ctx:    ctx,
	}
}

// / DeleteStore
type SdkClientDeleteStoreRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientDeleteStoreOptions
}
type ClientDeleteStoreOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

type ClientDeleteStoreResponse struct{}

func (request SdkClientDeleteStoreRequest) Options(options ClientDeleteStoreOptions) SdkClientDeleteStoreRequest {
	request.options = &options
	return request
}

func (request SdkClientDeleteStoreRequest) Execute() (ClientDeleteStoreResponse, *_nethttp.Response, error) {
	return request.Client.DeleteStoreExecute(request)
}

func (client *OpenFgaClient) DeleteStoreExecute(request SdkClientDeleteStoreRequest) (ClientDeleteStoreResponse, *_nethttp.Response, error) {
	httpResponse, err := client.OpenFgaApi.DeleteStore(request.ctx).Execute()
	return ClientDeleteStoreResponse{}, httpResponse, err
}

func (client *OpenFgaClient) DeleteStore(ctx _context.Context) SdkClientDeleteStoreRequest {
	return SdkClientDeleteStoreRequest{
		Client: *client,
		ctx:    ctx,
	}
}

/* Authorization Models */

// / ReadAuthorizationModels
type SdkClientReadAuthorizationModelsRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientReadAuthorizationModelsOptions
}

type ClientReadAuthorizationModelsOptions struct {
	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
}

func (request SdkClientReadAuthorizationModelsRequest) Options(options ClientReadAuthorizationModelsOptions) SdkClientReadAuthorizationModelsRequest {
	request.options = &options
	return request
}

func (request SdkClientReadAuthorizationModelsRequest) Execute() (openfga.ReadAuthorizationModelsResponse, *_nethttp.Response, error) {
	return request.Client.ReadAuthorizationModelsExecute(request)
}

func (client *OpenFgaClient) ReadAuthorizationModelsExecute(request SdkClientReadAuthorizationModelsRequest) (openfga.ReadAuthorizationModelsResponse, *_nethttp.Response, error) {
	req := client.OpenFgaApi.ReadAuthorizationModels(request.ctx)
	pageSize := getPageSizeFromRequest((*ClientPaginationOptions)(request.options))
	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	continuationToken := getContinuationTokenFromRequest((*ClientPaginationOptions)(request.options))
	if continuationToken != nil {
		req.ContinuationToken(*continuationToken)
	}
	return req.Execute()
}

func (client *OpenFgaClient) ReadAuthorizationModels(ctx _context.Context) SdkClientReadAuthorizationModelsRequest {
	return SdkClientReadAuthorizationModelsRequest{
		Client: *client,
		ctx:    ctx,
	}
}

// / WriteAuthorizationModel
type SdkClientWriteAuthorizationModelRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientWriteAuthorizationModelRequest
	options *ClientWriteAuthorizationModelOptions
}

type ClientWriteAuthorizationModelRequest struct {
	TypeDefinitions []openfga.TypeDefinition `json:"type_definitions"`
	SchemaVersion   string                   `json:"schema_version,omitempty"`
}

type ClientWriteAuthorizationModelOptions struct {
}

func (request SdkClientWriteAuthorizationModelRequest) Options(options ClientWriteAuthorizationModelOptions) SdkClientWriteAuthorizationModelRequest {
	request.options = &options
	return request
}

func (request SdkClientWriteAuthorizationModelRequest) Body(body ClientWriteAuthorizationModelRequest) SdkClientWriteAuthorizationModelRequest {
	request.body = &body
	return request
}

func (request SdkClientWriteAuthorizationModelRequest) Execute() (openfga.WriteAuthorizationModelResponse, *_nethttp.Response, error) {
	return request.Client.WriteAuthorizationModelExecute(request)
}

func (client *OpenFgaClient) WriteAuthorizationModelExecute(request SdkClientWriteAuthorizationModelRequest) (openfga.WriteAuthorizationModelResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.WriteAuthorizationModel(request.ctx).Body(openfga.WriteAuthorizationModelRequest{
		TypeDefinitions: request.body.TypeDefinitions,
		SchemaVersion:   openfga.PtrString(request.body.SchemaVersion),
	}).Execute()
}

func (client *OpenFgaClient) WriteAuthorizationModel(ctx _context.Context) SdkClientWriteAuthorizationModelRequest {
	return SdkClientWriteAuthorizationModelRequest{
		Client: *client,
		ctx:    ctx,
	}
}

// / ReadAuthorizationModel
type SdkClientReadAuthorizationModelRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientReadAuthorizationModelRequest
	options *ClientReadAuthorizationModelOptions
}

type ClientReadAuthorizationModelRequest struct {
}

type ClientReadAuthorizationModelOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (request SdkClientReadAuthorizationModelRequest) Options(options ClientReadAuthorizationModelOptions) SdkClientReadAuthorizationModelRequest {
	request.options = &options
	return request
}

func (request SdkClientReadAuthorizationModelRequest) Body(body ClientReadAuthorizationModelRequest) SdkClientReadAuthorizationModelRequest {
	request.body = &body
	return request
}

func (request SdkClientReadAuthorizationModelRequest) Execute() (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error) {
	return request.Client.ReadAuthorizationModelExecute(request)
}

func (client *OpenFgaClient) ReadAuthorizationModelExecute(request SdkClientReadAuthorizationModelRequest) (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.ReadAuthorizationModel(request.ctx, *request.options.AuthorizationModelId).Execute()
}

func (client *OpenFgaClient) ReadAuthorizationModel(ctx _context.Context) SdkClientReadAuthorizationModelRequest {
	return SdkClientReadAuthorizationModelRequest{
		Client: *client,
		ctx:    ctx,
	}
}

// / ReadLatestAuthorizationModel
type SdkClientReadLatestAuthorizationModelRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientReadLatestAuthorizationModelOptions
}

type ClientReadLatestAuthorizationModelOptions struct {
}

func (client *OpenFgaClient) ReadLatestAuthorizationModel(ctx _context.Context) SdkClientReadLatestAuthorizationModelRequest {
	return SdkClientReadLatestAuthorizationModelRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientReadLatestAuthorizationModelRequest) Options(options ClientReadLatestAuthorizationModelOptions) SdkClientReadLatestAuthorizationModelRequest {
	request.options = &options
	return request
}

func (request SdkClientReadLatestAuthorizationModelRequest) Execute() (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error) {
	return request.Client.ReadLatestAuthorizationModelExecute(request)
}

func (client *OpenFgaClient) ReadLatestAuthorizationModelExecute(request SdkClientReadLatestAuthorizationModelRequest) (openfga.ReadAuthorizationModelResponse, *_nethttp.Response, error) {
	response, httpResponse, err := client.ReadAuthorizationModels(request.ctx).Options(ClientReadAuthorizationModelsOptions{
		PageSize: openfga.PtrInt32(1),
	}).Execute()

	var authorizationModel *openfga.AuthorizationModel

	if err == nil && len(*response.AuthorizationModels) > 0 {
		authorizationModels := *response.AuthorizationModels
		authorizationModel = &(authorizationModels)[0]
	}

	return openfga.ReadAuthorizationModelResponse{
		AuthorizationModel: authorizationModel,
	}, httpResponse, err
}

/* Relationship Tuples */

// / ReadChanges
type SdkClientReadChangesRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientReadChangesRequest
	options *ClientReadChangesOptions
}

type ClientReadChangesRequest struct {
	Type string `json:"type,omitempty"`
}

type ClientReadChangesOptions struct {
	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
}

func (client *OpenFgaClient) ReadChanges(ctx _context.Context) SdkClientReadChangesRequest {
	return SdkClientReadChangesRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientReadChangesRequest) Options(options ClientReadChangesOptions) SdkClientReadChangesRequest {
	request.options = &options
	return request
}

func (request SdkClientReadChangesRequest) Body(body ClientReadChangesRequest) SdkClientReadChangesRequest {
	request.body = &body
	return request
}

func (request SdkClientReadChangesRequest) Execute() (openfga.ReadChangesResponse, *_nethttp.Response, error) {
	return request.Client.ReadChangesExecute(request)
}

func (client *OpenFgaClient) ReadChangesExecute(request SdkClientReadChangesRequest) (openfga.ReadChangesResponse, *_nethttp.Response, error) {
	req := client.OpenFgaApi.ReadChanges(request.ctx)
	pageSize := getPageSizeFromRequest((*ClientPaginationOptions)(request.options))
	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	continuationToken := getContinuationTokenFromRequest((*ClientPaginationOptions)(request.options))
	if continuationToken != nil {
		req.ContinuationToken(*continuationToken)
	}
	return req.Execute()
}

// / Read
type SdkClientReadRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientReadRequest
	options *ClientReadOptions
}

type ClientReadRequest struct {
	User     *string `json:"user,omitempty"`
	Relation *string `json:"relation,omitempty"`
	Object   *string `json:"object,omitempty"`
}

type ClientReadOptions struct {
	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
}

func (client *OpenFgaClient) Read(ctx _context.Context) SdkClientReadRequest {
	return SdkClientReadRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientReadRequest) Options(options ClientReadOptions) SdkClientReadRequest {
	request.options = &options
	return request
}

func (request SdkClientReadRequest) Body(body ClientReadRequest) SdkClientReadRequest {
	request.body = &body
	return request
}

func (request SdkClientReadRequest) Execute() (openfga.ReadResponse, *_nethttp.Response, error) {
	return request.Client.ReadExecute(request)
}

func (client *OpenFgaClient) ReadExecute(request SdkClientReadRequest) (openfga.ReadResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.Read(request.ctx).Body(openfga.ReadRequest{
		TupleKey: &openfga.TupleKey{
			User:     request.body.User,
			Relation: request.body.Relation,
			Object:   request.body.Object,
		},
		PageSize:          getPageSizeFromRequest((*ClientPaginationOptions)(request.options)),
		ContinuationToken: getContinuationTokenFromRequest((*ClientPaginationOptions)(request.options)),
	}).Execute()
}

// / Write
type SdkClientWriteRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientWriteRequest
	options *ClientWriteOptions
}

type ClientWriteRequest struct {
	Writes  *[]ClientTupleKey
	Deletes *[]ClientTupleKey
}

type TransactionOptions struct {
	// If set to true will disable running in transaction mode (transaction mode means everything is sent in a single transaction to the server)
	Disable bool `json:"disable,omitempty"`
	// When transaction mode is disabled, the requests are chunked and sent separately and each chunk is a transaction (default = 1)
	MaxPerChunk int32 `json:"max_per_chunk,omitempty"`
	// Number of requests to issue in parallel
	MaxParallelRequests int32 `json:"max_parallel_requests,omitempty"`
}

type ClientWriteOptions struct {
	AuthorizationModelId *string             `json:"authorization_model_id,omitempty"`
	Transaction          *TransactionOptions `json:"transaction_options,omitempty"`
}

type ClientWriteStatus string

// List of ClientWriteStatus
const (
	SUCCESS ClientWriteStatus = "CLIENT_WRITE_STATUS_SUCCESS"
	FAILURE ClientWriteStatus = "CLIENT_WRITE_STATUS_FAILURE"
)

//var allowedTupleOperationEnumValues = []ClientWriteStatus{
//	SUCCESS,
//	FAILURE,
//}

type ClientWriteSingleResponse struct {
	TupleKey     ClientTupleKey     `json:"tuple_key,omitempty"`
	Status       ClientWriteStatus  `json:"status,omitempty"`
	HttpResponse *_nethttp.Response `json:"http_response,omitempty"`
	Error        error              `json:"error,omitempty"`
}

func (o ClientWriteSingleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	toSerialize["status"] = o.Status
	if o.HttpResponse != nil {
		toSerialize["http_response"] = o.HttpResponse
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type ClientWriteResponse struct {
	Writes  []ClientWriteSingleResponse `json:"writes,omitempty"`
	Deletes []ClientWriteSingleResponse `json:"deletes,omitempty"`
}

func (o ClientWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Writes != nil {
		toSerialize["writes"] = o.Writes
	}
	if o.Deletes != nil {
		toSerialize["writes"] = o.Deletes
	}
	return json.Marshal(toSerialize)
}

func (client *OpenFgaClient) Write(ctx _context.Context) SdkClientWriteRequest {
	return SdkClientWriteRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientWriteRequest) Options(options ClientWriteOptions) SdkClientWriteRequest {
	request.options = &options
	return request
}

func (request SdkClientWriteRequest) Body(body ClientWriteRequest) SdkClientWriteRequest {
	request.body = &body
	return request
}

func (request SdkClientWriteRequest) Execute() (ClientWriteResponse, error) {
	return request.Client.WriteExecute(request)
}

func (client *OpenFgaClient) WriteExecute(request SdkClientWriteRequest) (ClientWriteResponse, error) {
	var maxPerChunk = int32(1) // 1 has to be the default otherwise the chunks will be sent in transactions
	if request.options != nil && request.options.Transaction != nil {
		maxPerChunk = request.options.Transaction.MaxPerChunk
	}
	var maxParallelReqs = DEFAULT_MAX_METHOD_PARALLEL_REQS
	if request.options != nil && request.options.Transaction != nil {
		maxParallelReqs = request.options.Transaction.MaxParallelRequests
	}

	response := ClientWriteResponse{
		Writes:  []ClientWriteSingleResponse{},
		Deletes: []ClientWriteSingleResponse{},
	}

	// Unless explicitly disabled, transaction mode is enabled
	// In transaction mode, the client will send the request to the server as is
	if request.options == nil || request.options.Transaction == nil || !request.options.Transaction.Disable {
		writeRequest := openfga.WriteRequest{
			AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
		}
		if request.body.Writes != nil && len(*request.body.Writes) > 0 {
			writes := openfga.TupleKeys{}
			for index := 0; index < len(*request.body.Writes); index++ {
				writes.TupleKeys = append(writes.TupleKeys, (*request.body.Writes)[index].ToTupleKey())
			}
			writeRequest.Writes = &writes
		}
		if request.body.Deletes != nil && len(*request.body.Deletes) > 0 {
			deletes := openfga.TupleKeys{}
			for index := 0; index < len(*request.body.Deletes); index++ {
				deletes.TupleKeys = append(deletes.TupleKeys, (*request.body.Deletes)[index].ToTupleKey())
			}
			writeRequest.Deletes = &deletes
		}
		_, httpResponse, err := client.OpenFgaApi.Write(request.ctx).Body(writeRequest).Execute()

		clientWriteStatus := SUCCESS
		if err != nil {
			clientWriteStatus = FAILURE
		}

		if request.body.Writes != nil {
			writeRequestTupleKeys := *request.body.Writes
			for index := 0; index < len(writeRequestTupleKeys); index++ {
				response.Writes = append(response.Writes, ClientWriteSingleResponse{
					TupleKey:     writeRequestTupleKeys[index],
					HttpResponse: httpResponse,
					Status:       clientWriteStatus,
					Error:        err,
				})
			}
		}

		if request.body.Deletes != nil {
			deleteRequestTupleKeys := *request.body.Deletes
			for index := 0; index < len(deleteRequestTupleKeys); index++ {
				response.Deletes = append(response.Deletes, ClientWriteSingleResponse{
					TupleKey:     deleteRequestTupleKeys[index],
					HttpResponse: httpResponse,
					Status:       clientWriteStatus,
					Error:        err,
				})
			}
		}

		return response, err
	}

	// If the transaction mode is disabled:
	// - the client will attempt to chunk the writes and deletes into multiple requests
	// - each request is a transaction
	// - the max items in each request are based on maxPerChunk (default=1)
	var writeChunkSize = int(maxPerChunk)
	var writeChunks [][]ClientTupleKey
	for i := 0; i < len(*request.body.Writes); i += writeChunkSize {
		end := int(math.Min(float64(i+writeChunkSize), float64(len(*request.body.Writes))))

		writeChunks = append(writeChunks, (*request.body.Writes)[i:end])
	}

	writeGroup, ctx := errgroup.WithContext(request.ctx)
	writeGroup.SetLimit(int(maxParallelReqs))
	writeResponses := make([]ClientWriteResponse, len(writeChunks))
	for index, writeBody := range writeChunks {
		index, writeBody := index, writeBody
		writeGroup.Go(func() error {
			singleResponse, _ := client.WriteExecute(SdkClientWriteRequest{
				ctx:    ctx,
				Client: *client,
				body: &ClientWriteRequest{
					Writes: &writeBody,
				},
				options: &ClientWriteOptions{
					AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
				},
			})

			writeResponses[index] = singleResponse

			return nil
		})
	}

	_ = writeGroup.Wait()

	var deleteChunkSize = int(maxPerChunk)
	var deleteChunks [][]ClientTupleKey
	for i := 0; i < len(*request.body.Deletes); i += deleteChunkSize {
		end := int(math.Min(float64(i+writeChunkSize), float64(len(*request.body.Deletes))))

		deleteChunks = append(deleteChunks, (*request.body.Deletes)[i:end])
	}

	deleteGroup, ctx := errgroup.WithContext(request.ctx)
	deleteGroup.SetLimit(int(maxParallelReqs))
	deleteResponses := make([]ClientWriteResponse, len(deleteChunks))
	for index, deleteBody := range deleteChunks {
		index, deleteBody := index, deleteBody
		deleteGroup.Go(func() error {
			singleResponse, _ := client.WriteExecute(SdkClientWriteRequest{
				ctx:    ctx,
				Client: *client,
				body: &ClientWriteRequest{
					Deletes: &deleteBody,
				},
				options: &ClientWriteOptions{
					AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
				},
			})

			deleteResponses[index] = singleResponse

			return nil
		})
	}

	_ = deleteGroup.Wait()

	for _, writeResponse := range writeResponses {
		for _, writeSingleResponse := range writeResponse.Writes {
			response.Writes = append(response.Writes, writeSingleResponse)
		}
	}

	for _, deleteResponse := range deleteResponses {
		for _, deleteSingleResponse := range deleteResponse.Deletes {
			response.Deletes = append(response.Deletes, deleteSingleResponse)
		}
	}

	return response, nil
}

// / WriteTuples
type SdkClientWriteTuplesRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientWriteTuplesBody
	options *ClientWriteOptions
}

type ClientWriteTuplesBody = []ClientTupleKey

func (client *OpenFgaClient) WriteTuples(ctx _context.Context) SdkClientWriteTuplesRequest {
	return SdkClientWriteTuplesRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientWriteTuplesRequest) Options(options ClientWriteOptions) SdkClientWriteTuplesRequest {
	request.options = &options
	return request
}

func (request SdkClientWriteTuplesRequest) Body(body ClientWriteTuplesBody) SdkClientWriteTuplesRequest {
	request.body = &body
	return request
}

func (request SdkClientWriteTuplesRequest) Execute() (ClientWriteResponse, error) {
	return request.Client.WriteTuplesExecute(request)
}

func (client *OpenFgaClient) WriteTuplesExecute(request SdkClientWriteTuplesRequest) (ClientWriteResponse, error) {
	return client.Write(request.ctx).Body(ClientWriteRequest{
		Writes: request.body,
	}).Options(*request.options).Execute()
}

// / DeleteTuples
type SdkClientDeleteTuplesRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientDeleteTuplesBody
	options *ClientWriteOptions
}

type ClientDeleteTuplesBody = []ClientTupleKey

func (client *OpenFgaClient) DeleteTuples(ctx _context.Context) SdkClientDeleteTuplesRequest {
	return SdkClientDeleteTuplesRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientDeleteTuplesRequest) Options(options ClientWriteOptions) SdkClientDeleteTuplesRequest {
	request.options = &options
	return request
}

func (request SdkClientDeleteTuplesRequest) Body(body ClientDeleteTuplesBody) SdkClientDeleteTuplesRequest {
	request.body = &body
	return request
}

func (request SdkClientDeleteTuplesRequest) Execute() (ClientWriteResponse, error) {
	return request.Client.DeleteTuplesExecute(request)
}

func (client *OpenFgaClient) DeleteTuplesExecute(request SdkClientDeleteTuplesRequest) (ClientWriteResponse, error) {
	return client.Write(request.ctx).Body(ClientWriteRequest{
		Deletes: request.body,
	}).Options(*request.options).Execute()
}

/* Relationship Queries */

/// Check

type SdkClientCheckRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientCheckRequest
	options *ClientCheckOptions
}

type ClientCheckRequest struct {
	User             string            `json:"user,omitempty"`
	Relation         string            `json:"relation,omitempty"`
	Object           string            `json:"object,omitempty"`
	ContextualTuples *[]ClientTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientCheckOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (client *OpenFgaClient) Check(ctx _context.Context) SdkClientCheckRequest {
	return SdkClientCheckRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientCheckRequest) Options(options ClientCheckOptions) SdkClientCheckRequest {
	request.options = &options
	return request
}

func (request SdkClientCheckRequest) Body(body ClientCheckRequest) SdkClientCheckRequest {
	request.body = &body
	return request
}

func (request SdkClientCheckRequest) Execute() (openfga.CheckResponse, *_nethttp.Response, error) {
	return request.Client.CheckExecute(request)
}

func (client *OpenFgaClient) CheckExecute(request SdkClientCheckRequest) (openfga.CheckResponse, *_nethttp.Response, error) {
	var contextualTuples []openfga.TupleKey
	if request.body.ContextualTuples != nil {
		for index := 0; index < len(*request.body.ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (*request.body.ContextualTuples)[index].ToTupleKey())
		}
	}
	requestBody := openfga.CheckRequest{
		TupleKey: openfga.TupleKey{
			User:     openfga.PtrString(request.body.User),
			Relation: openfga.PtrString(request.body.Relation),
			Object:   openfga.PtrString(request.body.Object),
		},
		ContextualTuples:     openfga.NewContextualTupleKeys(contextualTuples),
		AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
	}

	return client.OpenFgaApi.Check(request.ctx).Body(requestBody).Execute()
}

/// BatchCheck

type SdkClientBatchCheckRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientBatchCheckBody
	options *ClientBatchCheckOptions
}

type ClientBatchCheckBody = []ClientCheckRequest

type ClientBatchCheckOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
	MaxParallelRequests  *int32  `json:"max_parallel_requests,omitempty"`
}

type ClientBatchCheckSingleResponse struct {
	openfga.CheckResponse
	Request      ClientCheckRequest
	HttpResponse *_nethttp.Response
	Error        error
}

type ClientBatchCheckResponse = []ClientBatchCheckSingleResponse

func (client *OpenFgaClient) BatchCheck(ctx _context.Context) SdkClientBatchCheckRequest {
	return SdkClientBatchCheckRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientBatchCheckRequest) Options(options ClientBatchCheckOptions) SdkClientBatchCheckRequest {
	request.options = &options
	return request
}

func (request SdkClientBatchCheckRequest) Body(body ClientBatchCheckBody) SdkClientBatchCheckRequest {
	request.body = &body
	return request
}

func (request SdkClientBatchCheckRequest) Execute() (ClientBatchCheckResponse, error) {
	return request.Client.BatchCheckExecute(request)
}

func (client *OpenFgaClient) BatchCheckExecute(request SdkClientBatchCheckRequest) (ClientBatchCheckResponse, error) {
	group, ctx := errgroup.WithContext(request.ctx)
	var maxParallelReqs int
	if request.options == nil || request.options.MaxParallelRequests == nil {
		maxParallelReqs = int(DEFAULT_MAX_METHOD_PARALLEL_REQS)
	} else {
		maxParallelReqs = int(*request.options.MaxParallelRequests)
	}
	group.SetLimit(maxParallelReqs)
	var numOfChecks = len(*request.body)
	response := make(ClientBatchCheckResponse, numOfChecks)
	for index, checkBody := range *request.body {
		index, checkBody := index, checkBody
		group.Go(func() error {
			singleResponse, httpResponse, err := client.CheckExecute(SdkClientCheckRequest{
				ctx:    ctx,
				Client: *client,
				body:   &checkBody,
				options: &ClientCheckOptions{
					AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
				},
			})

			response[index] = ClientBatchCheckSingleResponse{
				Request:       checkBody,
				CheckResponse: singleResponse,
				HttpResponse:  httpResponse,
				Error:         err,
			}

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	}

	return response, nil
}

// / Expand
type SdkClientExpandRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientExpandRequest
	options *ClientExpandOptions
}

type ClientExpandRequest struct {
	Relation string `json:"relation,omitempty"`
	Object   string `json:"object,omitempty"`
}

type ClientExpandOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (client *OpenFgaClient) Expand(ctx _context.Context) SdkClientExpandRequest {
	return SdkClientExpandRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientExpandRequest) Options(options ClientExpandOptions) SdkClientExpandRequest {
	request.options = &options
	return request
}

func (request SdkClientExpandRequest) Body(body ClientExpandRequest) SdkClientExpandRequest {
	request.body = &body
	return request
}

func (request SdkClientExpandRequest) Execute() (openfga.ExpandResponse, *_nethttp.Response, error) {
	return request.Client.ExpandExecute(request)
}

func (client *OpenFgaClient) ExpandExecute(request SdkClientExpandRequest) (openfga.ExpandResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.Expand(request.ctx).Body(openfga.ExpandRequest{
		TupleKey: openfga.TupleKey{
			Relation: &request.body.Relation,
			Object:   &request.body.Object,
		},
		AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
	}).Execute()
}

// / ListObjects
type SdkClientListObjectsRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientListObjectsRequest
	options *ClientListObjectsOptions
}

type ClientListObjectsRequest struct {
	User             string            `json:"user,omitempty"`
	Relation         string            `json:"relation,omitempty"`
	Type             string            `json:"type,omitempty"`
	ContextualTuples *[]ClientTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientListObjectsOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (client *OpenFgaClient) ListObjects(ctx _context.Context) SdkClientListObjectsRequest {
	return SdkClientListObjectsRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientListObjectsRequest) Options(options ClientListObjectsOptions) SdkClientListObjectsRequest {
	request.options = &options
	return request
}

func (request SdkClientListObjectsRequest) Body(body ClientListObjectsRequest) SdkClientListObjectsRequest {
	request.body = &body
	return request
}

func (request SdkClientListObjectsRequest) Execute() (openfga.ListObjectsResponse, *_nethttp.Response, error) {
	return request.Client.ListObjectsExecute(request)
}

func (client *OpenFgaClient) ListObjectsExecute(request SdkClientListObjectsRequest) (openfga.ListObjectsResponse, *_nethttp.Response, error) {
	var contextualTuples []openfga.TupleKey
	if request.body.ContextualTuples != nil {
		for index := 0; index < len(*request.body.ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (*request.body.ContextualTuples)[index].ToTupleKey())
		}
	}
	return client.OpenFgaApi.ListObjects(request.ctx).Body(openfga.ListObjectsRequest{
		User:                 request.body.User,
		Relation:             request.body.Relation,
		Type:                 request.body.Type,
		ContextualTuples:     openfga.NewContextualTupleKeys(contextualTuples),
		AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
	}).Execute()
}

/// ListRelations

type SdkClientListRelationsRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientListRelationsRequest
	options *ClientListRelationsOptions
}

type ClientListRelationsRequest struct {
	User             string            `json:"user,omitempty"`
	Object           string            `json:"object,omitempty"`
	Relations        []string          `json:"relations,omitempty"`
	ContextualTuples *[]ClientTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientListRelationsOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

type ClientListRelationsResponse struct {
	Relations []string `json:"response,omitempty"`
}

func (o ClientListRelationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relations"] = o.Relations
	return json.Marshal(toSerialize)
}

func (client *OpenFgaClient) ListRelations(ctx _context.Context) SdkClientListRelationsRequest {
	return SdkClientListRelationsRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientListRelationsRequest) Options(options ClientListRelationsOptions) SdkClientListRelationsRequest {
	request.options = &options
	return request
}

func (request SdkClientListRelationsRequest) Body(body ClientListRelationsRequest) SdkClientListRelationsRequest {
	request.body = &body
	return request
}

func (request SdkClientListRelationsRequest) Execute() (*ClientListRelationsResponse, error) {
	return request.Client.ListRelationsExecute(request)
}

func (client *OpenFgaClient) ListRelationsExecute(request SdkClientListRelationsRequest) (*ClientListRelationsResponse, error) {
	if len(request.body.Relations) <= 0 {
		return nil, fmt.Errorf("ListRelations - expected len(Relations) > 0")
	}

	batchRequestBody := ClientBatchCheckBody{}
	for index := 0; index < len(request.body.Relations); index++ {
		batchRequestBody = append(batchRequestBody, ClientCheckRequest{
			User:             request.body.User,
			Relation:         request.body.Relations[index],
			Object:           request.body.Object,
			ContextualTuples: request.body.ContextualTuples,
		})
	}

	batchResponse, err := client.BatchCheckExecute(SdkClientBatchCheckRequest{
		ctx:    request.ctx,
		Client: *client,
		body:   &batchRequestBody,
		options: &ClientBatchCheckOptions{
			AuthorizationModelId: client.getAuthorizationModelId(request.options.AuthorizationModelId),
		},
	})

	if err != nil {
		return &ClientListRelationsResponse{}, err
	}

	var relations []string
	for index := 0; index < len(batchResponse); index++ {
		if batchResponse[index].GetAllowed() {
			relations = append(relations, batchResponse[index].Request.Relation)
		}
	}

	return &ClientListRelationsResponse{Relations: relations}, err
}

// / ReadAssertions
type SdkClientReadAssertionsRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientReadAssertionsOptions
}

type ClientReadAssertionsOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (client *OpenFgaClient) ReadAssertions(ctx _context.Context) SdkClientReadAssertionsRequest {
	return SdkClientReadAssertionsRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientReadAssertionsRequest) Options(options ClientReadAssertionsOptions) SdkClientReadAssertionsRequest {
	request.options = &options
	return request
}

func (request SdkClientReadAssertionsRequest) Execute() (openfga.ReadAssertionsResponse, *_nethttp.Response, error) {
	return request.Client.ReadAssertionsExecute(request)
}

func (client *OpenFgaClient) ReadAssertionsExecute(request SdkClientReadAssertionsRequest) (openfga.ReadAssertionsResponse, *_nethttp.Response, error) {
	return client.OpenFgaApi.ReadAssertions(request.ctx, *client.getAuthorizationModelId(request.options.AuthorizationModelId)).Execute()
}

// / WriteAssertions
type SdkClientWriteAssertionsRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	body    *ClientWriteAssertionsRequest
	options *ClientWriteAssertionsOptions
}

type ClientAssertion struct {
	User        string `json:"user,omitempty"`
	Relation    string `json:"relation,omitempty"`
	Object      string `json:"object,omitempty"`
	Expectation bool   `json:"expectation,omitempty"`
}

type ClientWriteAssertionsRequest = []ClientAssertion

func (clientAssertion ClientAssertion) ToAssertion() openfga.Assertion {
	return openfga.Assertion{
		TupleKey: openfga.TupleKey{
			User:     openfga.PtrString(clientAssertion.User),
			Relation: openfga.PtrString(clientAssertion.Relation),
			Object:   openfga.PtrString(clientAssertion.Object),
		},
		Expectation: clientAssertion.Expectation,
	}
}

type ClientWriteAssertionsOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

func (client *OpenFgaClient) WriteAssertions(ctx _context.Context) SdkClientWriteAssertionsRequest {
	return SdkClientWriteAssertionsRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request SdkClientWriteAssertionsRequest) Options(options ClientWriteAssertionsOptions) SdkClientWriteAssertionsRequest {
	request.options = &options
	return request
}

func (request SdkClientWriteAssertionsRequest) Body(body ClientWriteAssertionsRequest) SdkClientWriteAssertionsRequest {
	request.body = &body
	return request
}

func (request SdkClientWriteAssertionsRequest) Execute() (*_nethttp.Response, error) {
	return request.Client.WriteAssertionsExecute(request)
}

func (client *OpenFgaClient) WriteAssertionsExecute(request SdkClientWriteAssertionsRequest) (*_nethttp.Response, error) {
	writeAssertionsRequest := openfga.WriteAssertionsRequest{}
	for index := 0; index < len(*request.body); index++ {
		clientAssertion := (*request.body)[index]
		writeAssertionsRequest.Assertions = append(writeAssertionsRequest.Assertions, clientAssertion.ToAssertion())
	}
	return client.OpenFgaApi.WriteAssertions(request.ctx, *client.getAuthorizationModelId(request.options.AuthorizationModelId)).Body(writeAssertionsRequest).Execute()
}
