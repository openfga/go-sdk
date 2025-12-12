package client

import (
	_context "context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	_nethttp "net/http"
	"time"

	"github.com/sourcegraph/conc/pool"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/internal/constants"
	internalutils "github.com/openfga/go-sdk/internal/utils"
	"github.com/openfga/go-sdk/telemetry"
)

var (
	_ _context.Context
	// Ensure the SdkClient fits OpenFgaClient interface
	_ SdkClient = (*OpenFgaClient)(nil)
)

var DEFAULT_MAX_METHOD_PARALLEL_REQS = int32(constants.ClientMaxMethodParallelRequests)

type ClientConfiguration struct {
	fgaSdk.Configuration
	// ApiScheme - defines the scheme for the API: http or https
	// Deprecated: use ApiUrl instead of ApiScheme and ApiHost
	ApiScheme string `json:"api_scheme"`
	// ApiHost - defines the host for the API without the scheme e.g. (api.fga.example)
	// Deprecated: use ApiUrl instead of ApiScheme and ApiHost
	ApiHost              string                   `json:"api_host"`
	ApiUrl               string                   `json:"api_url"`
	StoreId              string                   `json:"store_id"`
	AuthorizationModelId string                   `json:"authorization_model_id"`
	Credentials          *credentials.Credentials `json:"credentials"`
	DefaultHeaders       map[string]string        `json:"default_headers"`
	UserAgent            string                   `json:"user_agent"`
	Debug                bool                     `json:"debug"`
	HTTPClient           *_nethttp.Client
	RetryParams          *fgaSdk.RetryParams
	Telemetry            *telemetry.Configuration `json:"telemetry,omitempty"`
}

func newClientConfiguration(cfg *fgaSdk.Configuration) ClientConfiguration {
	return ClientConfiguration{
		ApiScheme:      cfg.ApiScheme,
		ApiHost:        cfg.ApiHost,
		ApiUrl:         cfg.ApiUrl,
		Credentials:    cfg.Credentials,
		DefaultHeaders: cfg.DefaultHeaders,
		UserAgent:      cfg.UserAgent,
		Debug:          cfg.Debug,
		HTTPClient:     cfg.HTTPClient,
		RetryParams:    cfg.RetryParams,
		Telemetry:      cfg.Telemetry,
	}
}

type OpenFgaClient struct {
	config ClientConfiguration
	SdkClient
	fgaSdk.APIClient
}

func NewSdkClient(cfg *ClientConfiguration) (*OpenFgaClient, error) {
	apiConfiguration, err := fgaSdk.NewConfiguration(fgaSdk.Configuration{
		ApiScheme:      cfg.ApiScheme,
		ApiHost:        cfg.ApiHost,
		ApiUrl:         cfg.ApiUrl,
		Credentials:    cfg.Credentials,
		DefaultHeaders: cfg.DefaultHeaders,
		UserAgent:      cfg.UserAgent,
		Debug:          cfg.Debug,
		HTTPClient:     cfg.HTTPClient,
		RetryParams:    cfg.RetryParams,
		Telemetry:      cfg.Telemetry,
	})

	if err != nil {
		return nil, err
	}

	clientConfig := newClientConfiguration(apiConfiguration)
	clientConfig.AuthorizationModelId = cfg.AuthorizationModelId
	clientConfig.StoreId = cfg.StoreId

	// store id is already validate as part of configuration validation

	if cfg.AuthorizationModelId != "" && !internalutils.IsWellFormedUlidString(cfg.AuthorizationModelId) {
		return nil, FgaInvalidError{param: "AuthorizationModelId", description: "Expected ULID format"}
	}

	if cfg.StoreId != "" && !internalutils.IsWellFormedUlidString(cfg.StoreId) {
		return nil, FgaInvalidError{param: "StoreId", description: "Expected ULID format"}
	}

	apiClient := fgaSdk.NewAPIClient(apiConfiguration)

	return &OpenFgaClient{
		config:    clientConfig,
		APIClient: *apiClient,
	}, nil
}

type RequestOptions = fgaSdk.RequestOptions

type AuthorizationModelIdOptions struct {
	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
}

type ClientRequestOptionsWithAuthZModelId struct {
	RequestOptions
	AuthorizationModelIdOptions
}

type ClientTupleKey = fgaSdk.TupleKey
type ClientTupleKeyWithoutCondition = fgaSdk.TupleKeyWithoutCondition
type ClientCheckRequestTupleKey = fgaSdk.CheckRequestTupleKey
type ClientReadRequestTupleKey = fgaSdk.ReadRequestTupleKey
type ClientExpandRequestTupleKey = fgaSdk.ExpandRequestTupleKey
type ClientContextualTupleKey = ClientTupleKey

// ClientBatchCheckItem represents a flattened check item for batch check operations
type ClientBatchCheckItem struct {
	User             string                     `json:"user"`
	Relation         string                     `json:"relation"`
	Object           string                     `json:"object"`
	CorrelationId    string                     `json:"correlation_id"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
}

// ClientBatchCheckRequest represents a request for batch check operations
type ClientBatchCheckRequest struct {
	Checks []ClientBatchCheckItem `json:"checks"`
}

// BatchCheckOptions represents options for server-side batch check operations
type BatchCheckOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	MaxParallelRequests  *int32                        `json:"max_parallel_requests,omitempty"`
	MaxBatchSize         *int32                        `json:"max_batch_size,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
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

	/*
	 * ListStores Get a paginated list of stores.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientListStoresRequestInterface
	 */
	ListStores(ctx _context.Context) SdkClientListStoresRequestInterface

	/*
	 * ListStoresExecute executes the ListStores request
	 * @return *ClientListStoresResponse
	 */
	ListStoresExecute(request SdkClientListStoresRequestInterface) (*ClientListStoresResponse, error)

	/*
	 * CreateStore Create and initialize a store
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientCreateStoreRequestInterface
	 */
	CreateStore(ctx _context.Context) SdkClientCreateStoreRequestInterface

	/*
	 * CreateStoreExecute executes the CreateStore request
	 * @return *ClientCreateStoreResponse
	 */
	CreateStoreExecute(request SdkClientCreateStoreRequestInterface) (*ClientCreateStoreResponse, error)

	/*
	 * GetStore Get information about the current store.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientGetStoreRequestInterface
	 */
	GetStore(ctx _context.Context) SdkClientGetStoreRequestInterface

	/*
	 * GetStoreExecute executes the GetStore request
	 * @return *ClientGetStoreResponse
	 */
	GetStoreExecute(request SdkClientGetStoreRequestInterface) (*ClientGetStoreResponse, error)

	/*
	 * DeleteStore Delete a store.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientDeleteStoreRequestInterface
	 */
	DeleteStore(ctx _context.Context) SdkClientDeleteStoreRequestInterface

	/*
	 * DeleteStoreExecute executes the DeleteStore request
	 * @return *ClientDeleteStoreResponse
	 */
	DeleteStoreExecute(request SdkClientDeleteStoreRequestInterface) (*ClientDeleteStoreResponse, error)

	/* Authorization Models */

	/*
	 * ReadAuthorizationModels Read all authorization models in the store.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadAuthorizationModelsRequest
	 */
	ReadAuthorizationModels(ctx _context.Context) SdkClientReadAuthorizationModelsRequestInterface

	/*
	 * ReadAuthorizationModelsExecute executes the ReadAuthorizationModels request
	 * @return *ClientReadAuthorizationModelsResponse
	 */
	ReadAuthorizationModelsExecute(request SdkClientReadAuthorizationModelsRequestInterface) (*ClientReadAuthorizationModelsResponse, error)

	/*
	 * WriteAuthorizationModel Create a new authorization model.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientWriteAuthorizationModelRequest
	 */
	WriteAuthorizationModel(ctx _context.Context) SdkClientWriteAuthorizationModelRequestInterface

	/*
	 * WriteAuthorizationModelExecute executes the WriteAuthorizationModel request
	 * @return *ClientWriteAuthorizationModelResponse
	 */
	WriteAuthorizationModelExecute(request SdkClientWriteAuthorizationModelRequestInterface) (*ClientWriteAuthorizationModelResponse, error)

	/*
	 * ReadAuthorizationModel Read a particular authorization model.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadAuthorizationModelRequestInterface
	 */
	ReadAuthorizationModel(ctx _context.Context) SdkClientReadAuthorizationModelRequestInterface

	/*
	 * ReadAuthorizationModelExecute executes the ReadAuthorizationModel request
	 * @return *ClientReadAuthorizationModelResponse
	 */
	ReadAuthorizationModelExecute(request SdkClientReadAuthorizationModelRequestInterface) (*ClientReadAuthorizationModelResponse, error)

	/*
	 * ReadLatestAuthorizationModel Reads the latest authorization model (note: this ignores the model id in configuration).
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadLatestAuthorizationModelRequestInterface
	 */
	ReadLatestAuthorizationModel(ctx _context.Context) SdkClientReadLatestAuthorizationModelRequestInterface

	/*
	 * ReadLatestAuthorizationModelExecute executes the ReadLatestAuthorizationModel request
	 * @return *ClientReadAuthorizationModelResponse
	 */
	ReadLatestAuthorizationModelExecute(request SdkClientReadLatestAuthorizationModelRequestInterface) (*ClientReadAuthorizationModelResponse, error)

	/* Relationship Tuples */

	/*
	 * ReadChanges Reads the list of historical relationship tuple writes and deletes.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadChangesRequestInterface
	 */
	ReadChanges(ctx _context.Context) SdkClientReadChangesRequestInterface

	/*
	 * ReadChangesExecute executes the ReadChanges request
	 * @return *ClientReadChangesResponse
	 */
	ReadChangesExecute(request SdkClientReadChangesRequestInterface) (*ClientReadChangesResponse, error)

	/*
	 * Read Reads the relationship tuples stored in the database. It does not evaluate nor exclude invalid tuples according to the authorization model.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadRequestInterface
	 */
	Read(ctx _context.Context) SdkClientReadRequestInterface

	/*
	 * ReadExecute executes the Read request
	 * @return *ClientReadResponse
	 */
	ReadExecute(request SdkClientReadRequestInterface) (*ClientReadResponse, error)

	/*
	 * Write Create and/or delete relationship tuples to update the system state.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientWriteRequestInterface
	 */
	Write(ctx _context.Context) SdkClientWriteRequestInterface

	/*
	 * WriteExecute executes the Write request
	 * @return *ClientWriteResponse
	 */
	WriteExecute(request SdkClientWriteRequestInterface) (*ClientWriteResponse, error)

	/*
	 * WriteTuples Utility method around Write
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientWriteTuplesRequestInterface
	 */
	WriteTuples(ctx _context.Context) SdkClientWriteTuplesRequestInterface

	/*
	 * WriteTuplesExecute executes the WriteTuples request
	 * @return *ClientWriteResponse
	 */
	WriteTuplesExecute(request SdkClientWriteTuplesRequestInterface) (*ClientWriteResponse, error)

	/*
	 * DeleteTuples Utility method around Write
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientDeleteTuplesRequestInterface
	 */
	DeleteTuples(ctx _context.Context) SdkClientDeleteTuplesRequestInterface

	/*
	 * DeleteTuplesExecute executes the DeleteTuples request
	 * @return *ClientWriteResponse
	 */
	DeleteTuplesExecute(request SdkClientDeleteTuplesRequestInterface) (*ClientWriteResponse, error)

	/* Relationship Queries */

	/*
	 * Check Check if a user has a particular relation with an object.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientCheckRequestInterface
	 */
	Check(ctx _context.Context) SdkClientCheckRequestInterface

	/*
	 * CheckExecute executes the Check request
	 * @return *ClientCheckResponse
	 */
	CheckExecute(request SdkClientCheckRequestInterface) (*ClientCheckResponse, error)

	/*
	 * ClientBatchCheck Run a set of [checks](#check). Batch Check will return `allowed: false` if it encounters an error, and will return the error in the body.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientBatchCheckRequestInterface
	 */
	ClientBatchCheck(ctx _context.Context) SdkClientBatchCheckClientRequestInterface

	/*
	 * ClientBatchCheckExecute executes the BatchCheck request
	 * @return *ClientBatchCheckResponse
	 */
	ClientBatchCheckExecute(request SdkClientBatchCheckClientRequestInterface) (*ClientBatchCheckClientResponse, error)

	/*
	 * BatchCheck Run a set of checks on the server. Server-side batch check allows for more efficient checking of multiple tuples.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientBatchCheckRequestInterface
	 */
	BatchCheck(ctx _context.Context) SdkClientBatchCheckRequestInterface

	/*
	 * BatchCheckExecute executes the server-side BatchCheck request
	 * @return *BatchCheckResponse
	 */
	BatchCheckExecute(request SdkClientBatchCheckRequestInterface) (*fgaSdk.BatchCheckResponse, error)

	/*
	 * Expand Expands the relationships in userset tree format.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientExpandRequestInterface
	 */
	Expand(ctx _context.Context) SdkClientExpandRequestInterface

	/*
	 * ExpandExecute executes the Expand request
	 * @return *ClientExpandResponse
	 */
	ExpandExecute(request SdkClientExpandRequestInterface) (*ClientExpandResponse, error)

	/*
	 * ListObjects List the objects of a particular type a user has access to.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientListObjectsRequestInterface
	 */
	ListObjects(ctx _context.Context) SdkClientListObjectsRequestInterface

	/*
	 * ListObjectsExecute executes the ListObjects request
	 * @return *ClientListObjectsResponse
	 */
	ListObjectsExecute(request SdkClientListObjectsRequestInterface) (*ClientListObjectsResponse, error)

	/*
	 * ListRelations List the relations a user has on an object.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientListRelationsRequestInterface
	 */
	ListRelations(ctx _context.Context) SdkClientListRelationsRequestInterface

	/*
	 * ListRelationsExecute executes the ListRelations request
	 * @return *ClientListRelationsResponse
	 */
	ListRelationsExecute(request SdkClientListRelationsRequestInterface) (*ClientListRelationsResponse, error)

	/*
	 * ListUsers List all users of the given type that the object has a relation with
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListUsersRequest
	 */
	ListUsers(ctx _context.Context) SdkClientListUsersRequestInterface

	/*
	 * ListUsersExecute executes the request
	 * @return ListUsersResponse
	 */
	ListUsersExecute(r SdkClientListUsersRequestInterface) (*ClientListUsersResponse, error)

	/* Assertions */

	/*
	 * ReadAssertions Read assertions for a particular authorization model.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientReadAssertionsRequestInterface
	 */
	ReadAssertions(ctx _context.Context) SdkClientReadAssertionsRequestInterface

	/*
	 * ReadAssertionsExecute executes the ReadAssertions request
	 * @return *ClientReadAssertionsResponse
	 */
	ReadAssertionsExecute(request SdkClientReadAssertionsRequestInterface) (*ClientReadAssertionsResponse, error)

	/*
	 * WriteAssertions Update the assertions for a particular authorization model.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SdkClientWriteAssertionsRequestInterface
	 */
	WriteAssertions(ctx _context.Context) SdkClientWriteAssertionsRequestInterface

	/*
	 * WriteAssertionsExecute executes the WriteAssertions request
	 * @return *ClientWriteAssertionsResponse
	 */
	WriteAssertionsExecute(request SdkClientWriteAssertionsRequestInterface) (*ClientWriteAssertionsResponse, error)

	/*
	 * SetAuthorizationModelId allows setting the Authorization Model ID for an OpenFgaClient.
	 * @param string authorizationModelId - The Authorization Model ID to set.
	 */
	SetAuthorizationModelId(authorizationModelId string) error
	/*
	 * GetAuthorizationModelId retrieves the Authorization Model ID for an OpenFgaClient.
	 * @return string
	 */
	GetAuthorizationModelId() (string, error)
	/*
	 * SetStoreId allows setting the Store ID for an OpenFgaClient.
	 * @param string storeId - The Store ID to set.
	 */
	SetStoreId(storeId string) error
	/*
	 * GetStoreId retrieves the Store ID set in the OpenFgaClient.
	 * @return string
	 */
	GetStoreId() (string, error)
}

func (client *OpenFgaClient) SetAuthorizationModelId(authorizationModelId string) error {
	if authorizationModelId != "" && !internalutils.IsWellFormedUlidString(authorizationModelId) {
		return FgaInvalidError{param: "AuthorizationModelId", description: "Expected ULID format"}
	}

	client.config.AuthorizationModelId = authorizationModelId

	return nil
}

func (client *OpenFgaClient) GetAuthorizationModelId() (string, error) {
	modelId := client.config.AuthorizationModelId
	if modelId != "" && !internalutils.IsWellFormedUlidString(modelId) {
		return "", FgaInvalidError{param: "AuthorizationModelId", description: "Expected ULID format"}
	}

	return modelId, nil
}

func (client *OpenFgaClient) getAuthorizationModelId(authorizationModelId *string) (*string, error) {
	modelId := client.config.AuthorizationModelId
	if authorizationModelId != nil && *authorizationModelId != "" {
		modelId = *authorizationModelId
	}

	if modelId != "" && !internalutils.IsWellFormedUlidString(modelId) {
		return nil, FgaInvalidError{param: "AuthorizationModelId", description: "Expected ULID format"}
	}
	return &modelId, nil
}

func (client *OpenFgaClient) SetStoreId(storeId string) error {
	if storeId != "" && !internalutils.IsWellFormedUlidString(storeId) {
		return FgaInvalidError{param: "StoreId", description: "Expected ULID format"}
	}
	client.config.StoreId = storeId
	return nil
}

func (client *OpenFgaClient) GetStoreId() (string, error) {
	storeId := client.config.StoreId
	if storeId != "" && !internalutils.IsWellFormedUlidString(storeId) {
		return "", FgaInvalidError{param: "StoreId", description: "Expected ULID format"}
	}
	return storeId, nil
}

func (client *OpenFgaClient) getStoreId(storeId *string) (*string, error) {
	store := client.config.StoreId
	if storeId != nil && *storeId != "" {
		store = *storeId
	}
	if store != "" && !internalutils.IsWellFormedUlidString(store) {
		return nil, FgaInvalidError{param: "StoreId", description: "Expected ULID format"}
	}

	return &store, nil
}

/* Stores */

// / ListStores
type SdkClientListStoresRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	options *ClientListStoresOptions
}

type SdkClientListStoresRequestInterface interface {
	Options(options ClientListStoresOptions) SdkClientListStoresRequestInterface
	Execute() (*ClientListStoresResponse, error)
	GetContext() _context.Context
	GetOptions() *ClientListStoresOptions
}

type ClientListStoresOptions struct {
	RequestOptions

	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
	Name              *string `json:"name,omitempty"`
}

type ClientListStoresResponse = fgaSdk.ListStoresResponse

func (request *SdkClientListStoresRequest) Options(options ClientListStoresOptions) SdkClientListStoresRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientListStoresRequest) Execute() (*ClientListStoresResponse, error) {
	return request.Client.ListStoresExecute(request)
}

func (request *SdkClientListStoresRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientListStoresRequest) GetOptions() *ClientListStoresOptions {
	return request.options
}

func (client *OpenFgaClient) ListStoresExecute(request SdkClientListStoresRequestInterface) (*ClientListStoresResponse, error) {
	req := client.OpenFgaApi.ListStores(request.GetContext())
	options := request.GetOptions()
	if options != nil {
		req = req.Options(options.RequestOptions)
		if options.PageSize != nil {
			req = req.PageSize(*options.PageSize)
		}
		if options.ContinuationToken != nil {
			req = req.ContinuationToken(*options.ContinuationToken)
		}
		if options.Name != nil {
			req = req.Name(*options.Name)
		}
	}

	data, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) ListStores(ctx _context.Context) SdkClientListStoresRequestInterface {
	return &SdkClientListStoresRequest{
		ctx:    ctx,
		Client: client,
	}
}

// / CreateStore
type SdkClientCreateStoreRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientCreateStoreRequest
	options *ClientCreateStoreOptions
}

type SdkClientCreateStoreRequestInterface interface {
	Options(options ClientCreateStoreOptions) SdkClientCreateStoreRequestInterface
	Body(body ClientCreateStoreRequest) SdkClientCreateStoreRequestInterface
	Execute() (*ClientCreateStoreResponse, error)

	GetContext() _context.Context
	GetOptions() *ClientCreateStoreOptions
	GetBody() *ClientCreateStoreRequest
}

type ClientCreateStoreRequest struct {
	Name string `json:"name"`
}

type ClientCreateStoreOptions struct {
	RequestOptions
}

type ClientCreateStoreResponse = fgaSdk.CreateStoreResponse

func (request *SdkClientCreateStoreRequest) Options(options ClientCreateStoreOptions) SdkClientCreateStoreRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientCreateStoreRequest) Body(body ClientCreateStoreRequest) SdkClientCreateStoreRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientCreateStoreRequest) Execute() (*ClientCreateStoreResponse, error) {
	return request.Client.CreateStoreExecute(request)
}

func (request *SdkClientCreateStoreRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientCreateStoreRequest) GetOptions() *ClientCreateStoreOptions {
	return request.options
}

func (request *SdkClientCreateStoreRequest) GetBody() *ClientCreateStoreRequest {
	return request.body
}

func (client *OpenFgaClient) CreateStoreExecute(request SdkClientCreateStoreRequestInterface) (*ClientCreateStoreResponse, error) {
	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	requestBody := fgaSdk.CreateStoreRequest{
		Name: request.GetBody().Name,
	}

	data, _, err := client.OpenFgaApi.
		CreateStore(request.GetContext()).
		Body(requestBody).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) CreateStore(ctx _context.Context) SdkClientCreateStoreRequestInterface {
	return &SdkClientCreateStoreRequest{
		Client: client,
		ctx:    ctx,
	}
}

// / GetStore
type SdkClientGetStoreRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	options *ClientGetStoreOptions
}

type SdkClientGetStoreRequestInterface interface {
	Options(options ClientGetStoreOptions) SdkClientGetStoreRequestInterface
	Execute() (*ClientGetStoreResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientGetStoreOptions
}

type ClientGetStoreOptions struct {
	RequestOptions

	StoreId *string `json:"store_id,omitempty"`
}

type ClientGetStoreResponse = fgaSdk.GetStoreResponse

func (request *SdkClientGetStoreRequest) Options(options ClientGetStoreOptions) SdkClientGetStoreRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientGetStoreRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientGetStoreRequest) Execute() (*ClientGetStoreResponse, error) {
	return request.Client.GetStoreExecute(request)
}

func (request *SdkClientGetStoreRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientGetStoreRequest) GetOptions() *ClientGetStoreOptions {
	return request.options
}

func (client *OpenFgaClient) GetStoreExecute(request SdkClientGetStoreRequestInterface) (*ClientGetStoreResponse, error) {
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	data, _, err := client.OpenFgaApi.
		GetStore(request.GetContext(), *storeId).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) GetStore(ctx _context.Context) SdkClientGetStoreRequestInterface {
	return &SdkClientGetStoreRequest{
		Client: client,
		ctx:    ctx,
	}
}

// / DeleteStore
type SdkClientDeleteStoreRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	options *ClientDeleteStoreOptions
}

type SdkClientDeleteStoreRequestInterface interface {
	Options(options ClientDeleteStoreOptions) SdkClientDeleteStoreRequestInterface
	Execute() (*ClientDeleteStoreResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientDeleteStoreOptions
}

type ClientDeleteStoreOptions struct {
	RequestOptions

	StoreId *string `json:"store_id,omitempty"`
}

type ClientDeleteStoreResponse struct{}

func (request *SdkClientDeleteStoreRequest) Options(options ClientDeleteStoreOptions) SdkClientDeleteStoreRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientDeleteStoreRequest) Execute() (*ClientDeleteStoreResponse, error) {
	return request.Client.DeleteStoreExecute(request)
}

func (request *SdkClientDeleteStoreRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientDeleteStoreRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientDeleteStoreRequest) GetOptions() *ClientDeleteStoreOptions {
	return request.options
}

func (client *OpenFgaClient) DeleteStoreExecute(request SdkClientDeleteStoreRequestInterface) (*ClientDeleteStoreResponse, error) {
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	_, err = client.OpenFgaApi.
		DeleteStore(request.GetContext(), *storeId).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &ClientDeleteStoreResponse{}, nil
}

func (client *OpenFgaClient) DeleteStore(ctx _context.Context) SdkClientDeleteStoreRequestInterface {
	return &SdkClientDeleteStoreRequest{
		Client: client,
		ctx:    ctx,
	}
}

/* Authorization Models */

// / ReadAuthorizationModels
type SdkClientReadAuthorizationModelsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	options *ClientReadAuthorizationModelsOptions
}

type SdkClientReadAuthorizationModelsRequestInterface interface {
	Options(options ClientReadAuthorizationModelsOptions) SdkClientReadAuthorizationModelsRequestInterface
	Execute() (*ClientReadAuthorizationModelsResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientReadAuthorizationModelsOptions
}

type ClientReadAuthorizationModelsOptions struct {
	RequestOptions

	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
	StoreId           *string `json:"store_id,omitempty"`
}

type ClientReadAuthorizationModelsResponse = fgaSdk.ReadAuthorizationModelsResponse

func (request *SdkClientReadAuthorizationModelsRequest) Options(options ClientReadAuthorizationModelsOptions) SdkClientReadAuthorizationModelsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadAuthorizationModelsRequest) Execute() (*ClientReadAuthorizationModelsResponse, error) {
	return request.Client.ReadAuthorizationModelsExecute(request)
}

func (request *SdkClientReadAuthorizationModelsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadAuthorizationModelsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientReadAuthorizationModelsRequest) GetOptions() *ClientReadAuthorizationModelsOptions {
	return request.options
}

func (client *OpenFgaClient) ReadAuthorizationModelsExecute(request SdkClientReadAuthorizationModelsRequestInterface) (*ClientReadAuthorizationModelsResponse, error) {
	pagingOpts := ClientPaginationOptions{}
	if request.GetOptions() != nil {
		pagingOpts.PageSize = request.GetOptions().PageSize
		pagingOpts.ContinuationToken = request.GetOptions().ContinuationToken
	}

	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	req := client.OpenFgaApi.
		ReadAuthorizationModels(request.GetContext(), *storeId).
		Options(requestOptions)

	pageSize := getPageSizeFromRequest(&pagingOpts)
	if pageSize != nil {
		req = req.PageSize(*pageSize)
	}
	continuationToken := getContinuationTokenFromRequest(&pagingOpts)
	if continuationToken != nil {
		req = req.ContinuationToken(*continuationToken)
	}
	data, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) ReadAuthorizationModels(ctx _context.Context) SdkClientReadAuthorizationModelsRequestInterface {
	return &SdkClientReadAuthorizationModelsRequest{
		Client: client,
		ctx:    ctx,
	}
}

// / WriteAuthorizationModel
type SdkClientWriteAuthorizationModelRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientWriteAuthorizationModelRequest
	options *ClientWriteAuthorizationModelOptions
}

type SdkClientWriteAuthorizationModelRequestInterface interface {
	Options(options ClientWriteAuthorizationModelOptions) SdkClientWriteAuthorizationModelRequestInterface
	Body(body ClientWriteAuthorizationModelRequest) SdkClientWriteAuthorizationModelRequestInterface
	Execute() (*ClientWriteAuthorizationModelResponse, error)
	GetStoreIdOverride() *string

	GetBody() *ClientWriteAuthorizationModelRequest
	GetOptions() *ClientWriteAuthorizationModelOptions
	GetContext() _context.Context
}

type ClientWriteAuthorizationModelRequest = fgaSdk.WriteAuthorizationModelRequest

type ClientWriteAuthorizationModelOptions struct {
	RequestOptions

	StoreId *string `json:"store_id,omitempty"`
}

type ClientWriteAuthorizationModelResponse = fgaSdk.WriteAuthorizationModelResponse

func (request *SdkClientWriteAuthorizationModelRequest) Options(options ClientWriteAuthorizationModelOptions) SdkClientWriteAuthorizationModelRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientWriteAuthorizationModelRequest) Body(body ClientWriteAuthorizationModelRequest) SdkClientWriteAuthorizationModelRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientWriteAuthorizationModelRequest) Execute() (*ClientWriteAuthorizationModelResponse, error) {
	return request.Client.WriteAuthorizationModelExecute(request)
}

func (request *SdkClientWriteAuthorizationModelRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientWriteAuthorizationModelRequest) GetBody() *ClientWriteAuthorizationModelRequest {
	return request.body
}

func (request *SdkClientWriteAuthorizationModelRequest) GetOptions() *ClientWriteAuthorizationModelOptions {
	return request.options
}

func (request *SdkClientWriteAuthorizationModelRequest) GetContext() _context.Context {
	return request.ctx
}

func (client *OpenFgaClient) WriteAuthorizationModelExecute(request SdkClientWriteAuthorizationModelRequestInterface) (*ClientWriteAuthorizationModelResponse, error) {
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	data, _, err := client.OpenFgaApi.
		WriteAuthorizationModel(request.GetContext(), *storeId).
		Body(*request.GetBody()).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) WriteAuthorizationModel(ctx _context.Context) SdkClientWriteAuthorizationModelRequestInterface {
	return &SdkClientWriteAuthorizationModelRequest{
		Client: client,
		ctx:    ctx,
	}
}

// / ReadAuthorizationModel
type SdkClientReadAuthorizationModelRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientReadAuthorizationModelRequest
	options *ClientReadAuthorizationModelOptions
}

type SdkClientReadAuthorizationModelRequestInterface interface {
	Options(options ClientReadAuthorizationModelOptions) SdkClientReadAuthorizationModelRequestInterface
	Body(body ClientReadAuthorizationModelRequest) SdkClientReadAuthorizationModelRequestInterface
	Execute() (*ClientReadAuthorizationModelResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientReadAuthorizationModelRequest
	GetOptions() *ClientReadAuthorizationModelOptions
}

type ClientReadAuthorizationModelRequest struct {
}

type ClientReadAuthorizationModelOptions struct {
	RequestOptions

	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
	StoreId              *string `json:"store_id,omitempty"`
}

type ClientReadAuthorizationModelResponse = fgaSdk.ReadAuthorizationModelResponse

func (request *SdkClientReadAuthorizationModelRequest) Options(options ClientReadAuthorizationModelOptions) SdkClientReadAuthorizationModelRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadAuthorizationModelRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientReadAuthorizationModelRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadAuthorizationModelRequest) Body(body ClientReadAuthorizationModelRequest) SdkClientReadAuthorizationModelRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientReadAuthorizationModelRequest) Execute() (*ClientReadAuthorizationModelResponse, error) {
	return request.Client.ReadAuthorizationModelExecute(request)
}

func (request *SdkClientReadAuthorizationModelRequest) GetBody() *ClientReadAuthorizationModelRequest {
	return request.body
}

func (request *SdkClientReadAuthorizationModelRequest) GetOptions() *ClientReadAuthorizationModelOptions {
	return request.options
}

func (request *SdkClientReadAuthorizationModelRequest) GetContext() _context.Context {
	return request.ctx
}

func (client *OpenFgaClient) ReadAuthorizationModelExecute(request SdkClientReadAuthorizationModelRequestInterface) (*ClientReadAuthorizationModelResponse, error) {
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	if authorizationModelId == nil || *authorizationModelId == "" {
		return nil, FgaRequiredParamError{param: "AuthorizationModelId"}
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	data, _, err := client.OpenFgaApi.
		ReadAuthorizationModel(request.GetContext(), *storeId, *authorizationModelId).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *OpenFgaClient) ReadAuthorizationModel(ctx _context.Context) SdkClientReadAuthorizationModelRequestInterface {
	return &SdkClientReadAuthorizationModelRequest{
		Client: client,
		ctx:    ctx,
	}
}

// / ReadLatestAuthorizationModel
type SdkClientReadLatestAuthorizationModelRequest struct {
	ctx    _context.Context
	Client OpenFgaClient

	options *ClientReadLatestAuthorizationModelOptions
}

type SdkClientReadLatestAuthorizationModelRequestInterface interface {
	Options(options ClientReadLatestAuthorizationModelOptions) SdkClientReadLatestAuthorizationModelRequestInterface
	Execute() (*ClientReadAuthorizationModelResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientReadLatestAuthorizationModelOptions
}

type ClientReadLatestAuthorizationModelOptions struct {
	RequestOptions

	StoreId *string `json:"store_id,omitempty"`
}

func (client *OpenFgaClient) ReadLatestAuthorizationModel(ctx _context.Context) SdkClientReadLatestAuthorizationModelRequestInterface {
	return &SdkClientReadLatestAuthorizationModelRequest{
		Client: *client,
		ctx:    ctx,
	}
}

func (request *SdkClientReadLatestAuthorizationModelRequest) Options(options ClientReadLatestAuthorizationModelOptions) SdkClientReadLatestAuthorizationModelRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadLatestAuthorizationModelRequest) Execute() (*ClientReadAuthorizationModelResponse, error) {
	return request.Client.ReadLatestAuthorizationModelExecute(request)
}

func (request *SdkClientReadLatestAuthorizationModelRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadLatestAuthorizationModelRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientReadLatestAuthorizationModelRequest) GetOptions() *ClientReadLatestAuthorizationModelOptions {
	return request.options
}

func (client *OpenFgaClient) ReadLatestAuthorizationModelExecute(request SdkClientReadLatestAuthorizationModelRequestInterface) (*ClientReadAuthorizationModelResponse, error) {
	opts := ClientReadAuthorizationModelsOptions{
		PageSize: fgaSdk.PtrInt32(1),
	}
	if request.GetOptions() != nil {
		opts.StoreId = request.GetOptions().StoreId
		opts.RequestOptions = request.GetOptions().RequestOptions
	}
	req := client.ReadAuthorizationModels(request.GetContext()).Options(opts)

	response, err := req.Execute()
	if err != nil {
		return nil, err
	}

	var authorizationModel *fgaSdk.AuthorizationModel

	if len(response.AuthorizationModels) > 0 {
		authorizationModels := response.AuthorizationModels
		authorizationModel = &(authorizationModels)[0]
	}

	return &fgaSdk.ReadAuthorizationModelResponse{
		AuthorizationModel: authorizationModel,
	}, nil
}

/* Relationship Tuples */

// / ReadChanges
type SdkClientReadChangesRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientReadChangesRequest
	options *ClientReadChangesOptions
}

type SdkClientReadChangesRequestInterface interface {
	Options(options ClientReadChangesOptions) SdkClientReadChangesRequestInterface
	Body(body ClientReadChangesRequest) SdkClientReadChangesRequestInterface
	Execute() (*ClientReadChangesResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientReadChangesRequest
	GetOptions() *ClientReadChangesOptions
}

type ClientReadChangesRequest struct {
	Type      string    `json:"type,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

type ClientReadChangesOptions struct {
	RequestOptions

	PageSize          *int32  `json:"page_size,omitempty"`
	ContinuationToken *string `json:"continuation_token,omitempty"`
	StoreId           *string `json:"store_id"`
}

type ClientReadChangesResponse = fgaSdk.ReadChangesResponse

func (client *OpenFgaClient) ReadChanges(ctx _context.Context) SdkClientReadChangesRequestInterface {
	return &SdkClientReadChangesRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientReadChangesRequest) Options(options ClientReadChangesOptions) SdkClientReadChangesRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadChangesRequest) Body(body ClientReadChangesRequest) SdkClientReadChangesRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientReadChangesRequest) Execute() (*ClientReadChangesResponse, error) {
	return request.Client.ReadChangesExecute(request)
}

func (request *SdkClientReadChangesRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadChangesRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientReadChangesRequest) GetBody() *ClientReadChangesRequest {
	return request.body
}

func (request *SdkClientReadChangesRequest) GetOptions() *ClientReadChangesOptions {
	return request.options
}

func (client *OpenFgaClient) ReadChangesExecute(request SdkClientReadChangesRequestInterface) (*ClientReadChangesResponse, error) {
	pagingOpts := ClientPaginationOptions{}
	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		pagingOpts.PageSize = request.GetOptions().PageSize
		pagingOpts.ContinuationToken = request.GetOptions().ContinuationToken
	}

	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	req := client.OpenFgaApi.
		ReadChanges(request.GetContext(), *storeId).
		Options(requestOptions)
	pageSize := getPageSizeFromRequest(&pagingOpts)
	if pageSize != nil {
		req = req.PageSize(*pageSize)
	}
	continuationToken := getContinuationTokenFromRequest(&pagingOpts)
	if continuationToken != nil {
		req = req.ContinuationToken(*continuationToken)
	}
	requestBody := request.GetBody()
	if requestBody != nil && requestBody.Type != "" {
		req = req.Type_(requestBody.Type)
	}
	if requestBody != nil && !requestBody.StartTime.IsZero() {
		req = req.StartTime(requestBody.StartTime)
	}

	data, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// / Read
type SdkClientReadRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientReadRequest
	options *ClientReadOptions
}

type SdkClientReadRequestInterface interface {
	Options(options ClientReadOptions) SdkClientReadRequestInterface
	Body(body ClientReadRequest) SdkClientReadRequestInterface
	Execute() (*ClientReadResponse, error)
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientReadRequest
	GetOptions() *ClientReadOptions
}

type ClientReadRequest struct {
	User     *string `json:"user,omitempty"`
	Relation *string `json:"relation,omitempty"`
	Object   *string `json:"object,omitempty"`
}

type ClientReadOptions struct {
	RequestOptions

	PageSize          *int32                        `json:"page_size,omitempty"`
	ContinuationToken *string                       `json:"continuation_token,omitempty"`
	StoreId           *string                       `json:"store_id,omitempty"`
	Consistency       *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientReadResponse = fgaSdk.ReadResponse

func (client *OpenFgaClient) Read(ctx _context.Context) SdkClientReadRequestInterface {
	return &SdkClientReadRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientReadRequest) Options(options ClientReadOptions) SdkClientReadRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadRequest) Body(body ClientReadRequest) SdkClientReadRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientReadRequest) Execute() (*ClientReadResponse, error) {
	return request.Client.ReadExecute(request)
}

func (request *SdkClientReadRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientReadRequest) GetBody() *ClientReadRequest {
	return request.body
}

func (request *SdkClientReadRequest) GetOptions() *ClientReadOptions {
	return request.options
}

func (client *OpenFgaClient) ReadExecute(request SdkClientReadRequestInterface) (*ClientReadResponse, error) {
	pagingOpts := ClientPaginationOptions{}
	requestOptions := RequestOptions{}
	var consistency *fgaSdk.ConsistencyPreference
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		pagingOpts.PageSize = request.GetOptions().PageSize
		pagingOpts.ContinuationToken = request.GetOptions().ContinuationToken
		consistency = request.GetOptions().Consistency
	}

	body := fgaSdk.ReadRequest{
		PageSize:          getPageSizeFromRequest(&pagingOpts),
		ContinuationToken: getContinuationTokenFromRequest(&pagingOpts),
		Consistency:       consistency,
	}
	if request.GetBody() != nil && (request.GetBody().User != nil || request.GetBody().Relation != nil || request.GetBody().Object != nil) {
		body.TupleKey = &fgaSdk.ReadRequestTupleKey{
			User:     request.GetBody().User,
			Relation: request.GetBody().Relation,
			Object:   request.GetBody().Object,
		}
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	data, _, err := client.OpenFgaApi.
		Read(request.GetContext(), *storeId).
		Body(body).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// / Write
type SdkClientWriteRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientWriteRequest
	options *ClientWriteOptions
}

type SdkClientWriteRequestInterface interface {
	Options(options ClientWriteOptions) SdkClientWriteRequestInterface
	Body(body ClientWriteRequest) SdkClientWriteRequestInterface
	Execute() (*ClientWriteResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientWriteOptions
	GetBody() *ClientWriteRequest
}

type ClientWriteRequest struct {
	Writes  []ClientTupleKey
	Deletes []ClientTupleKeyWithoutCondition
}

type TransactionOptions struct {
	// If set to true will disable running in transaction mode (transaction mode means everything is sent in a single transaction to the server)
	Disable bool `json:"disable,omitempty"`
	// When transaction mode is disabled, the requests are chunked and sent separately and each chunk is a transaction (default = 1)
	MaxPerChunk int32 `json:"max_per_chunk,omitempty"`
	// Number of requests to issue in parallel
	MaxParallelRequests int32 `json:"max_parallel_requests,omitempty"`
}

// ClientWriteRequestOnDuplicateWrites indicates what to do when a write conflicts with an existing tuple
type ClientWriteRequestOnDuplicateWrites string

func (w *ClientWriteRequestOnDuplicateWrites) ToString() *string {
	if w == nil {
		return nil
	}

	str := string(*w)

	return &str
}

const (
	// CLIENT_WRITE_REQUEST_ON_DUPLICATE_WRITES_ERROR returns an error if a write conflicts with an existing tuple (default)
	CLIENT_WRITE_REQUEST_ON_DUPLICATE_WRITES_ERROR ClientWriteRequestOnDuplicateWrites = "error"
	// CLIENT_WRITE_REQUEST_ON_DUPLICATE_WRITES_IGNORE ignores writes that conflict with existing tuples (they must match exactly, including conditions)
	CLIENT_WRITE_REQUEST_ON_DUPLICATE_WRITES_IGNORE ClientWriteRequestOnDuplicateWrites = "ignore"
)

// ClientWriteRequestOnMissingDeletes indicates what to do when a delete is issued for a tuple that does not exist
type ClientWriteRequestOnMissingDeletes string

func (d *ClientWriteRequestOnMissingDeletes) ToString() *string {
	if d == nil {
		return nil
	}

	str := string(*d)

	return &str
}

const (
	// CLIENT_WRITE_REQUEST_ON_MISSING_DELETES_ERROR returns an error if a delete is issued for a tuple that does not exist (default)
	CLIENT_WRITE_REQUEST_ON_MISSING_DELETES_ERROR ClientWriteRequestOnMissingDeletes = "error"
	// CLIENT_WRITE_REQUEST_ON_MISSING_DELETES_IGNORE ignores deletes for tuples that do not exist
	CLIENT_WRITE_REQUEST_ON_MISSING_DELETES_IGNORE ClientWriteRequestOnMissingDeletes = "ignore"
)

type ClientWriteConflictOptions struct {
	// OnDuplicateWrites defines what to do when a write conflicts with an existing tuple
	// Options are: "error" (default) or "ignore"
	OnDuplicateWrites ClientWriteRequestOnDuplicateWrites `json:"on_duplicate_writes,omitempty"`
	// OnMissingDeletes defines what to do when a delete is issued for a tuple that does not exist
	// Options are: "error" (default) or "ignore"
	OnMissingDeletes ClientWriteRequestOnMissingDeletes `json:"on_missing_deletes,omitempty"`
}

type ClientWriteOptions struct {
	RequestOptions

	AuthorizationModelId *string             `json:"authorization_model_id,omitempty"`
	StoreId              *string             `json:"store_id,omitempty"`
	Transaction          *TransactionOptions `json:"transaction_options,omitempty"`
	Conflict             ClientWriteConflictOptions
}

type ClientWriteStatus string

// List of ClientWriteStatus
const (
	SUCCESS ClientWriteStatus = "CLIENT_WRITE_STATUS_SUCCESS"
	FAILURE ClientWriteStatus = "CLIENT_WRITE_STATUS_FAILURE"
)

type ClientWriteRequestWriteResponse struct {
	TupleKey     ClientTupleKey     `json:"tuple_key,omitempty"`
	Status       ClientWriteStatus  `json:"status,omitempty"`
	HttpResponse *_nethttp.Response `json:"http_response,omitempty"`
	Error        error              `json:"error,omitempty"`
}

func (o ClientWriteRequestWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	toSerialize["status"] = o.Status
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type ClientWriteRequestDeleteResponse struct {
	TupleKey     ClientTupleKeyWithoutCondition `json:"tuple_key,omitempty"`
	Status       ClientWriteStatus              `json:"status,omitempty"`
	HttpResponse *_nethttp.Response             `json:"http_response,omitempty"`
	Error        error                          `json:"error,omitempty"`
}

func (o ClientWriteRequestDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	toSerialize["status"] = o.Status
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type ClientWriteResponse struct {
	Writes  []ClientWriteRequestWriteResponse  `json:"writes,omitempty"`
	Deletes []ClientWriteRequestDeleteResponse `json:"deletes,omitempty"`
}

func (o ClientWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Writes != nil {
		toSerialize["writes"] = o.Writes
	}
	if o.Deletes != nil {
		toSerialize["deletes"] = o.Deletes
	}
	return json.Marshal(toSerialize)
}

func (client *OpenFgaClient) Write(ctx _context.Context) SdkClientWriteRequestInterface {
	return &SdkClientWriteRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientWriteRequest) Options(options ClientWriteOptions) SdkClientWriteRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientWriteRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientWriteRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientWriteRequest) Body(body ClientWriteRequest) SdkClientWriteRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientWriteRequest) Execute() (*ClientWriteResponse, error) {
	return request.Client.WriteExecute(request)
}

func (request *SdkClientWriteRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientWriteRequest) GetOptions() *ClientWriteOptions {
	return request.options
}

func (request *SdkClientWriteRequest) GetBody() *ClientWriteRequest {
	return request.body
}

func (client *OpenFgaClient) WriteExecute(request SdkClientWriteRequestInterface) (*ClientWriteResponse, error) {
	options := request.GetOptions()
	transactionOptionsSet := options != nil && options.Transaction != nil
	response := ClientWriteResponse{
		Writes:  []ClientWriteRequestWriteResponse{},
		Deletes: []ClientWriteRequestDeleteResponse{},
	}
	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}

	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	// Unless explicitly disabled, transaction mode is enabled
	// In transaction mode, the client will send the request to the server as is
	if !transactionOptionsSet || !options.Transaction.Disable {
		writeRequest := fgaSdk.WriteRequest{
			AuthorizationModelId: authorizationModelId,
		}
		if len(request.GetBody().Writes) > 0 {
			writes := fgaSdk.WriteRequestWrites{}
			if options != nil {
				writes.OnDuplicate = options.Conflict.OnDuplicateWrites.ToString()
			}
			for index := 0; index < len(request.GetBody().Writes); index++ {
				writes.TupleKeys = append(writes.TupleKeys, (request.GetBody().Writes)[index])
			}
			writeRequest.Writes = &writes
		}
		if len(request.GetBody().Deletes) > 0 {
			deletes := fgaSdk.WriteRequestDeletes{}
			if options != nil {
				deletes.OnMissing = options.Conflict.OnMissingDeletes.ToString()
			}
			for index := 0; index < len(request.GetBody().Deletes); index++ {
				deletes.TupleKeys = append(deletes.TupleKeys, (request.GetBody().Deletes)[index])
			}
			writeRequest.Deletes = &deletes
		}

		_, httpResponse, err := client.OpenFgaApi.
			Write(request.GetContext(), *storeId).
			Body(writeRequest).
			Options(requestOptions).
			Execute()

		clientWriteStatus := SUCCESS
		if err != nil {
			clientWriteStatus = FAILURE
		}

		if request.GetBody() != nil && request.GetBody().Writes != nil {
			writeRequestTupleKeys := request.GetBody().Writes
			for index := 0; index < len(writeRequestTupleKeys); index++ {
				response.Writes = append(response.Writes, ClientWriteRequestWriteResponse{
					TupleKey:     writeRequestTupleKeys[index],
					HttpResponse: httpResponse,
					Status:       clientWriteStatus,
					Error:        err,
				})
			}
		}

		if request.GetBody() != nil && request.GetBody().Deletes != nil {
			deleteRequestTupleKeys := request.GetBody().Deletes
			for index := 0; index < len(deleteRequestTupleKeys); index++ {
				response.Deletes = append(response.Deletes, ClientWriteRequestDeleteResponse{
					TupleKey:     deleteRequestTupleKeys[index],
					HttpResponse: httpResponse,
					Status:       clientWriteStatus,
					Error:        err,
				})
			}
		}

		return &response, err
	}

	maxPerChunk := int32(1) // 1 has to be the default otherwise the chunks will be sent in transactions
	if options.Transaction.MaxPerChunk > 0 {
		maxPerChunk = options.Transaction.MaxPerChunk
	}

	maxParallelReqs := DEFAULT_MAX_METHOD_PARALLEL_REQS
	if options.Transaction.MaxParallelRequests > 0 {
		maxParallelReqs = options.Transaction.MaxParallelRequests
	}

	// If the transaction mode is disabled:
	// - the client will attempt to chunk the writes and deletes into multiple requests
	// - each request is a transaction
	// - the max items in each request are based on maxPerChunk (default=1)
	var writeChunkSize = int(maxPerChunk)
	var writeChunks [][]ClientTupleKey
	if request.GetBody() != nil {
		for i := 0; i < len(request.GetBody().Writes); i += writeChunkSize {
			end := int(math.Min(float64(i+writeChunkSize), float64(len(request.GetBody().Writes))))
			writeChunks = append(writeChunks, (request.GetBody().Writes)[i:end])
		}
	}

	writePool := pool.NewWithResults[*ClientWriteResponse]().WithContext(request.GetContext()).WithMaxGoroutines(int(maxParallelReqs))
	for _, writeBody := range writeChunks {
		writeBody := writeBody
		writePool.Go(func(ctx _context.Context) (*ClientWriteResponse, error) {
			singleResponse, err := client.WriteExecute(&SdkClientWriteRequest{
				ctx:    ctx,
				Client: client,
				body: &ClientWriteRequest{
					Writes: writeBody,
				},
				options: &ClientWriteOptions{
					RequestOptions:       options.RequestOptions,
					AuthorizationModelId: authorizationModelId,
					StoreId:              request.GetStoreIdOverride(),
					Conflict:             options.Conflict,
				},
			})
			var authErr fgaSdk.FgaApiAuthenticationError
			// If an error was returned then it will be an authentication error so we want to return
			if errors.As(err, &authErr) {
				return nil, err
			}

			return singleResponse, nil
		})
	}
	writeResponses, err := writePool.Wait()
	if err != nil {
		return &response, err
	}

	var deleteChunkSize = int(maxPerChunk)
	var deleteChunks [][]ClientTupleKeyWithoutCondition
	if request.GetBody() != nil {
		for i := 0; i < len(request.GetBody().Deletes); i += deleteChunkSize {
			end := int(math.Min(float64(i+writeChunkSize), float64(len(request.GetBody().Deletes))))

			deleteChunks = append(deleteChunks, (request.GetBody().Deletes)[i:end])
		}
	}

	deletePool := pool.NewWithResults[*ClientWriteResponse]().WithContext(request.GetContext()).WithMaxGoroutines(int(maxParallelReqs))
	for _, deleteBody := range deleteChunks {
		deleteBody := deleteBody
		deletePool.Go(func(ctx _context.Context) (*ClientWriteResponse, error) {
			singleResponse, err := client.WriteExecute(&SdkClientWriteRequest{
				ctx:    ctx,
				Client: client,
				body: &ClientWriteRequest{
					Deletes: deleteBody,
				},
				options: &ClientWriteOptions{
					RequestOptions:       options.RequestOptions,
					AuthorizationModelId: authorizationModelId,
					StoreId:              request.GetStoreIdOverride(),
					Conflict:             options.Conflict,
				},
			})

			var authErr fgaSdk.FgaApiAuthenticationError
			if errors.As(err, &authErr) {
				return nil, err
			}
			return singleResponse, nil
		})
	}

	deleteResponses, err := deletePool.Wait()
	// If an error was returned then it will be an authentication error so we want to return
	if err != nil {
		return &response, err
	}

	for _, writeResponse := range writeResponses {
		response.Writes = append(response.Writes, writeResponse.Writes...)
	}

	for _, deleteResponse := range deleteResponses {
		response.Deletes = append(response.Deletes, deleteResponse.Deletes...)
	}

	return &response, nil
}

// / WriteTuples
type SdkClientWriteTuplesRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientWriteTuplesBody
	options *ClientWriteOptions
}

type SdkClientWriteTuplesRequestInterface interface {
	Options(options ClientWriteOptions) SdkClientWriteTuplesRequestInterface
	Body(body ClientWriteTuplesBody) SdkClientWriteTuplesRequestInterface
	Execute() (*ClientWriteResponse, error)

	GetContext() _context.Context
	GetBody() *ClientWriteTuplesBody
	GetOptions() *ClientWriteOptions
}

type ClientWriteTuplesBody = []ClientTupleKey

func (client *OpenFgaClient) WriteTuples(ctx _context.Context) SdkClientWriteTuplesRequestInterface {
	return &SdkClientWriteTuplesRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientWriteTuplesRequest) Options(options ClientWriteOptions) SdkClientWriteTuplesRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientWriteTuplesRequest) Body(body ClientWriteTuplesBody) SdkClientWriteTuplesRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientWriteTuplesRequest) Execute() (*ClientWriteResponse, error) {
	return request.Client.WriteTuplesExecute(request)
}

func (request *SdkClientWriteTuplesRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientWriteTuplesRequest) GetBody() *ClientWriteTuplesBody {
	return request.body
}

func (request *SdkClientWriteTuplesRequest) GetOptions() *ClientWriteOptions {
	return request.options
}

func (client *OpenFgaClient) WriteTuplesExecute(request SdkClientWriteTuplesRequestInterface) (*ClientWriteResponse, error) {
	baseReq := client.Write(request.GetContext()).Body(ClientWriteRequest{
		Writes: *request.GetBody(),
	})
	if request.GetOptions() != nil {
		baseReq = baseReq.Options(*request.GetOptions())
	}
	return baseReq.Execute()
}

// / DeleteTuples
type SdkClientDeleteTuplesRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientDeleteTuplesBody
	options *ClientWriteOptions
}

type SdkClientDeleteTuplesRequestInterface interface {
	Options(options ClientWriteOptions) SdkClientDeleteTuplesRequestInterface
	Body(body ClientDeleteTuplesBody) SdkClientDeleteTuplesRequestInterface
	Execute() (*ClientWriteResponse, error)

	GetContext() _context.Context
	GetBody() *ClientDeleteTuplesBody
	GetOptions() *ClientWriteOptions
}

type ClientDeleteTuplesBody = []ClientTupleKeyWithoutCondition

func (client *OpenFgaClient) DeleteTuples(ctx _context.Context) SdkClientDeleteTuplesRequestInterface {
	return &SdkClientDeleteTuplesRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientDeleteTuplesRequest) Options(options ClientWriteOptions) SdkClientDeleteTuplesRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientDeleteTuplesRequest) Body(body ClientDeleteTuplesBody) SdkClientDeleteTuplesRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientDeleteTuplesRequest) Execute() (*ClientWriteResponse, error) {
	return request.Client.DeleteTuplesExecute(request)
}

func (request *SdkClientDeleteTuplesRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientDeleteTuplesRequest) GetBody() *ClientDeleteTuplesBody {
	return request.body
}

func (request *SdkClientDeleteTuplesRequest) GetOptions() *ClientWriteOptions {
	return request.options
}

func (client *OpenFgaClient) DeleteTuplesExecute(request SdkClientDeleteTuplesRequestInterface) (*ClientWriteResponse, error) {
	baseReq := client.Write(request.GetContext()).Body(ClientWriteRequest{
		Deletes: *request.GetBody(),
	})
	if request.GetOptions() != nil {
		baseReq = baseReq.Options(*request.GetOptions())
	}
	return baseReq.Execute()
}

/* Relationship Queries */

/// Check

type SdkClientCheckRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientCheckRequest
	options *ClientCheckOptions
}

type SdkClientCheckRequestInterface interface {
	Options(options ClientCheckOptions) SdkClientCheckRequestInterface
	Body(body ClientCheckRequest) SdkClientCheckRequestInterface
	Execute() (*ClientCheckResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientCheckRequest
	GetOptions() *ClientCheckOptions
}

type ClientCheckRequest struct {
	User             string                     `json:"user,omitempty"`
	Relation         string                     `json:"relation,omitempty"`
	Object           string                     `json:"object,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientCheckOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientCheckResponse struct {
	fgaSdk.CheckResponse
	HttpResponse *_nethttp.Response
}

func (client *OpenFgaClient) Check(ctx _context.Context) SdkClientCheckRequestInterface {
	return &SdkClientCheckRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientCheckRequest) Options(options ClientCheckOptions) SdkClientCheckRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientCheckRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientCheckRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientCheckRequest) Body(body ClientCheckRequest) SdkClientCheckRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientCheckRequest) Execute() (*ClientCheckResponse, error) {
	return request.Client.CheckExecute(request)
}

func (request *SdkClientCheckRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientCheckRequest) GetBody() *ClientCheckRequest {
	return request.body
}

func (request *SdkClientCheckRequest) GetOptions() *ClientCheckOptions {
	return request.options
}

func (client *OpenFgaClient) CheckExecute(request SdkClientCheckRequestInterface) (*ClientCheckResponse, error) {
	if request.GetBody() == nil {
		return nil, FgaRequiredParamError{param: "body"}
	}

	var contextualTuples []ClientContextualTupleKey
	if request.GetBody().ContextualTuples != nil {
		for index := 0; index < len(request.GetBody().ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (request.GetBody().ContextualTuples)[index])
		}
	}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}
	requestBody := fgaSdk.CheckRequest{
		TupleKey: fgaSdk.CheckRequestTupleKey{
			User:     request.GetBody().User,
			Relation: request.GetBody().Relation,
			Object:   request.GetBody().Object,
		},
		Context:              request.GetBody().Context,
		ContextualTuples:     fgaSdk.NewContextualTupleKeys(contextualTuples),
		AuthorizationModelId: authorizationModelId,
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		requestBody.Consistency = request.GetOptions().Consistency
	}

	data, httpResponse, err := client.OpenFgaApi.
		Check(request.GetContext(), *storeId).
		Body(requestBody).
		Options(requestOptions).
		Execute()
	return &ClientCheckResponse{CheckResponse: data, HttpResponse: httpResponse}, err
}

/// ClientBatchCheck

type SdkClientBatchCheckClientRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientBatchCheckClientBody
	options *ClientBatchCheckClientOptions
}

type SdkClientBatchCheckClientRequestInterface interface {
	Options(options ClientBatchCheckClientOptions) SdkClientBatchCheckClientRequestInterface
	Body(body ClientBatchCheckClientBody) SdkClientBatchCheckClientRequestInterface
	Execute() (*ClientBatchCheckClientResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientBatchCheckClientBody
	GetOptions() *ClientBatchCheckClientOptions
}

type ClientBatchCheckClientBody = []ClientCheckRequest

type ClientBatchCheckClientOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	MaxParallelRequests  *int32                        `json:"max_parallel_requests,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientBatchCheckClientSingleResponse struct {
	ClientCheckResponse
	Request ClientCheckRequest
	Error   error
}

type ClientBatchCheckClientResponse = []ClientBatchCheckClientSingleResponse

func (client *OpenFgaClient) ClientBatchCheck(ctx _context.Context) SdkClientBatchCheckClientRequestInterface {
	return &SdkClientBatchCheckClientRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientBatchCheckClientRequest) Options(options ClientBatchCheckClientOptions) SdkClientBatchCheckClientRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientBatchCheckClientRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientBatchCheckClientRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientBatchCheckClientRequest) Body(body ClientBatchCheckClientBody) SdkClientBatchCheckClientRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientBatchCheckClientRequest) Execute() (*ClientBatchCheckClientResponse, error) {
	return request.Client.ClientBatchCheckExecute(request)
}

func (request *SdkClientBatchCheckClientRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientBatchCheckClientRequest) GetBody() *ClientBatchCheckClientBody {
	return request.body
}

func (request *SdkClientBatchCheckClientRequest) GetOptions() *ClientBatchCheckClientOptions {
	return request.options
}

func (client *OpenFgaClient) ClientBatchCheckExecute(request SdkClientBatchCheckClientRequestInterface) (*ClientBatchCheckClientResponse, error) {
	ctx := request.GetContext()
	requestOptions := RequestOptions{}
	maxParallelReqs := int(DEFAULT_MAX_METHOD_PARALLEL_REQS)
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		if request.GetOptions().MaxParallelRequests != nil {
			maxParallelReqs = int(*request.GetOptions().MaxParallelRequests)
		}
	}

	var numOfChecks = len(*request.GetBody())
	response := make(ClientBatchCheckClientResponse, numOfChecks)
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}

	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	checkOptions := &ClientCheckOptions{
		RequestOptions: requestOptions,

		AuthorizationModelId: authorizationModelId,
		StoreId:              storeId,
	}

	if request.GetOptions() != nil && request.GetOptions().Consistency != nil {
		checkOptions.Consistency = request.GetOptions().Consistency
	}

	type batchCheckResult struct {
		Index    int
		Response ClientBatchCheckClientSingleResponse
	}

	checkPool := pool.NewWithResults[*batchCheckResult]().WithContext(ctx).WithMaxGoroutines(maxParallelReqs)
	for index, checkBody := range *request.GetBody() {
		index, checkBody := index, checkBody
		checkPool.Go(func(ctx _context.Context) (*batchCheckResult, error) {
			singleResponse, err := client.CheckExecute(&SdkClientCheckRequest{
				ctx:     ctx,
				Client:  client,
				body:    &checkBody,
				options: checkOptions,
			})

			var authErr fgaSdk.FgaApiAuthenticationError
			// If an error was returned then it will be an authentication error so we want to return
			if errors.As(err, &authErr) {
				return nil, err
			}

			return &batchCheckResult{
				Index: index,
				Response: ClientBatchCheckClientSingleResponse{
					Request:             checkBody,
					ClientCheckResponse: *singleResponse,
					Error:               err,
				},
			}, nil
		})
	}

	results, err := checkPool.Wait()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		response[result.Index] = result.Response
	}

	return &response, nil
}

// Server-side BatchCheck implementation

// SdkClientBatchCheckRequest represents a server-side batch check request
type SdkClientBatchCheckRequest struct {
	ctx     _context.Context
	client  *OpenFgaClient
	body    *ClientBatchCheckRequest
	options *BatchCheckOptions
}

// SdkClientBatchCheckRequestInterface defines the interface for server-side batch check requests
type SdkClientBatchCheckRequestInterface interface {
	Body(body ClientBatchCheckRequest) SdkClientBatchCheckRequestInterface
	Options(options BatchCheckOptions) SdkClientBatchCheckRequestInterface
	Execute() (*fgaSdk.BatchCheckResponse, error)
	GetContext() _context.Context
	GetBody() *ClientBatchCheckRequest
	GetOptions() *BatchCheckOptions
}

func (r *SdkClientBatchCheckRequest) Body(body ClientBatchCheckRequest) SdkClientBatchCheckRequestInterface {
	r.body = &body
	return r
}

func (r *SdkClientBatchCheckRequest) Options(options BatchCheckOptions) SdkClientBatchCheckRequestInterface {
	r.options = &options
	return r
}

func (r *SdkClientBatchCheckRequest) Execute() (*fgaSdk.BatchCheckResponse, error) {
	return r.client.BatchCheckExecute(r)
}

func (r *SdkClientBatchCheckRequest) GetContext() _context.Context {
	return r.ctx
}

func (r *SdkClientBatchCheckRequest) GetBody() *ClientBatchCheckRequest {
	return r.body
}

func (r *SdkClientBatchCheckRequest) GetOptions() *BatchCheckOptions {
	return r.options
}

// BatchCheck initializes a new batch check request
func (client *OpenFgaClient) BatchCheck(ctx _context.Context) SdkClientBatchCheckRequestInterface {
	return &SdkClientBatchCheckRequest{
		ctx:    ctx,
		client: client,
	}
}

/*
 * BatchCheckExecute executes the server-side BatchCheck request
 * @param request SdkClientBatchCheckRequestInterface - the request interface
 * @return *fgaSdk.BatchCheckResponse
 */
func (client *OpenFgaClient) BatchCheckExecute(request SdkClientBatchCheckRequestInterface) (*fgaSdk.BatchCheckResponse, error) {
	ctx := request.GetContext()
	body := request.GetBody()
	options := request.GetOptions()

	if body == nil || len(body.Checks) == 0 {
		return nil, FgaRequiredParamError{param: "checks"}
	}

	if options == nil {
		options = &BatchCheckOptions{}
	}

	maxParallelRequests := DEFAULT_MAX_METHOD_PARALLEL_REQS
	if options.MaxParallelRequests != nil {
		maxParallelRequests = *options.MaxParallelRequests
	}

	maxBatchSize := int32(constants.ClientMaxBatchSize)
	if options.MaxBatchSize != nil {
		maxBatchSize = *options.MaxBatchSize
	}

	_, err := client.getStoreId(options.StoreId)
	if err != nil {
		return nil, err
	}

	authorizationModelId, err := client.getAuthorizationModelId(options.AuthorizationModelId)
	if err != nil {
		return nil, err
	}

	chunks := chunkClientBatchCheckItems(body.Checks, int(maxBatchSize))

	p := pool.NewWithResults[*fgaSdk.BatchCheckResponse]().WithContext(ctx).WithMaxGoroutines(int(maxParallelRequests))

	for _, chunk := range chunks {
		chunkCopy := chunk

		p.Go(func(ctx _context.Context) (*fgaSdk.BatchCheckResponse, error) {
			batchCheckRequest := createBatchCheckRequest(chunkCopy, authorizationModelId, options.Consistency)
			return client.singleBatchCheck(ctx, batchCheckRequest, options)
		})
	}

	responses, err := p.Wait()
	if err != nil {
		return nil, err
	}

	combinedResult := make(map[string]fgaSdk.BatchCheckSingleResult)

	for _, response := range responses {
		for correlationID, result := range response.GetResult() {
			combinedResult[correlationID] = result
		}
	}

	combinedResponse := fgaSdk.NewBatchCheckResponse()
	combinedResponse.SetResult(combinedResult)

	return combinedResponse, nil
}

/*
 * singleBatchCheck performs a single batch check request to the API
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 * @param body fgaSdk.BatchCheckRequest - the request body
 * @param options *BatchCheckOptions - options for the request
 * @return *fgaSdk.BatchCheckResponse
 */
func (client *OpenFgaClient) singleBatchCheck(ctx _context.Context, body fgaSdk.BatchCheckRequest, options *BatchCheckOptions) (*fgaSdk.BatchCheckResponse, error) {
	storeId, err := client.getStoreId(options.StoreId)
	if err != nil {
		return nil, err
	}

	req := client.OpenFgaApi.
		BatchCheck(ctx, *storeId).
		Body(body).
		Options(options.RequestOptions)
	response, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return &response, nil
}

/*
 * chunkClientBatchCheckItems splits a list of check items into chunks of specified size
 * @param items []ClientBatchCheckItem - the items to chunk
 * @param chunkSize int - the maximum size of each chunk
 * @return [][]ClientBatchCheckItem - the chunked items
 */
func chunkClientBatchCheckItems(items []ClientBatchCheckItem, chunkSize int) [][]ClientBatchCheckItem {
	if len(items) == 0 {
		return [][]ClientBatchCheckItem{}
	}

	chunks := make([][]ClientBatchCheckItem, 0, (len(items)+chunkSize-1)/chunkSize)

	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}

	return chunks
}

/*
 * createBatchCheckRequest creates a BatchCheckRequest from ClientBatchCheckItems
 * @param items []ClientBatchCheckItem - the client batch check items
 * @param authorizationModelId *string - optional authorization model ID
 * @param consistency *fgaSdk.ConsistencyPreference - optional consistency preference
 * @return fgaSdk.BatchCheckRequest - the created request
 */
func createBatchCheckRequest(items []ClientBatchCheckItem, authorizationModelId *string, consistency *fgaSdk.ConsistencyPreference) fgaSdk.BatchCheckRequest {
	batchCheckItems := make([]fgaSdk.BatchCheckItem, 0, len(items))

	for _, item := range items {
		tupleKey := fgaSdk.CheckRequestTupleKey{
			User:     item.User,
			Relation: item.Relation,
			Object:   item.Object,
		}

		batchCheckItem := fgaSdk.BatchCheckItem{
			TupleKey:      tupleKey,
			CorrelationId: item.CorrelationId,
		}

		if len(item.ContextualTuples) > 0 {
			contextualTuples := &fgaSdk.ContextualTupleKeys{
				TupleKeys: []fgaSdk.TupleKey{},
			}

			contextualTuples.TupleKeys = append(contextualTuples.TupleKeys, item.ContextualTuples...)

			batchCheckItem.ContextualTuples = contextualTuples
		}

		if item.Context != nil {
			batchCheckItem.Context = item.Context
		}

		batchCheckItems = append(batchCheckItems, batchCheckItem)
	}

	batchCheckRequest := fgaSdk.BatchCheckRequest{
		Checks: batchCheckItems,
	}

	if authorizationModelId != nil && *authorizationModelId != "" {
		batchCheckRequest.AuthorizationModelId = authorizationModelId
	}

	if consistency != nil {
		batchCheckRequest.Consistency = consistency
	}

	return batchCheckRequest
}

// / Expand
type SdkClientExpandRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientExpandRequest
	options *ClientExpandOptions
}

type SdkClientExpandRequestInterface interface {
	Options(options ClientExpandOptions) SdkClientExpandRequestInterface
	Body(body ClientExpandRequest) SdkClientExpandRequestInterface
	Execute() (*ClientExpandResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientExpandRequest
	GetOptions() *ClientExpandOptions
}

type ClientExpandRequest struct {
	Relation         string                     `json:"relation,omitempty"`
	Object           string                     `json:"object,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientExpandOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientExpandResponse = fgaSdk.ExpandResponse

func (client *OpenFgaClient) Expand(ctx _context.Context) SdkClientExpandRequestInterface {
	return &SdkClientExpandRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientExpandRequest) Options(options ClientExpandOptions) SdkClientExpandRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientExpandRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientExpandRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientExpandRequest) Body(body ClientExpandRequest) SdkClientExpandRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientExpandRequest) Execute() (*ClientExpandResponse, error) {
	return request.Client.ExpandExecute(request)
}

func (request *SdkClientExpandRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientExpandRequest) GetBody() *ClientExpandRequest {
	return request.body
}

func (request *SdkClientExpandRequest) GetOptions() *ClientExpandOptions {
	return request.options
}

func (client *OpenFgaClient) ExpandExecute(request SdkClientExpandRequestInterface) (*ClientExpandResponse, error) {
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	var contextualTuples []ClientContextualTupleKey
	if request.GetBody().ContextualTuples != nil {
		for index := 0; index < len(request.GetBody().ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (request.GetBody().ContextualTuples)[index])
		}
	}

	body := fgaSdk.ExpandRequest{
		TupleKey: fgaSdk.ExpandRequestTupleKey{
			Relation: request.GetBody().Relation,
			Object:   request.GetBody().Object,
		},
		ContextualTuples:     fgaSdk.NewContextualTupleKeys(contextualTuples),
		AuthorizationModelId: authorizationModelId,
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		body.Consistency = request.GetOptions().Consistency
	}

	data, _, err := client.OpenFgaApi.
		Expand(request.GetContext(), *storeId).
		Body(body).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// / ListObjects
type SdkClientListObjectsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientListObjectsRequest
	options *ClientListObjectsOptions
}

type SdkClientListObjectsRequestInterface interface {
	Options(options ClientListObjectsOptions) SdkClientListObjectsRequestInterface
	Body(body ClientListObjectsRequest) SdkClientListObjectsRequestInterface
	Execute() (*ClientListObjectsResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientListObjectsRequest
	GetOptions() *ClientListObjectsOptions
}

type ClientListObjectsRequest struct {
	User             string                     `json:"user,omitempty"`
	Relation         string                     `json:"relation,omitempty"`
	Type             string                     `json:"type,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientListObjectsOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientListObjectsResponse = fgaSdk.ListObjectsResponse

func (client *OpenFgaClient) ListObjects(ctx _context.Context) SdkClientListObjectsRequestInterface {
	return &SdkClientListObjectsRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientListObjectsRequest) Options(options ClientListObjectsOptions) SdkClientListObjectsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientListObjectsRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientListObjectsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientListObjectsRequest) Body(body ClientListObjectsRequest) SdkClientListObjectsRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientListObjectsRequest) Execute() (*ClientListObjectsResponse, error) {
	return request.Client.ListObjectsExecute(request)
}

func (request *SdkClientListObjectsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientListObjectsRequest) GetBody() *ClientListObjectsRequest {
	return request.body
}

func (request *SdkClientListObjectsRequest) GetOptions() *ClientListObjectsOptions {
	return request.options
}

func (client *OpenFgaClient) ListObjectsExecute(request SdkClientListObjectsRequestInterface) (*ClientListObjectsResponse, error) {
	var contextualTuples []ClientContextualTupleKey
	if request.GetBody().ContextualTuples != nil {
		for index := 0; index < len(request.GetBody().ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (request.GetBody().ContextualTuples)[index])
		}
	}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}
	body := fgaSdk.ListObjectsRequest{
		User:                 request.GetBody().User,
		Relation:             request.GetBody().Relation,
		Type:                 request.GetBody().Type,
		ContextualTuples:     fgaSdk.NewContextualTupleKeys(contextualTuples),
		Context:              request.GetBody().Context,
		AuthorizationModelId: authorizationModelId,
	}
	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		body.Consistency = request.GetOptions().Consistency
	}
	data, _, err := client.OpenFgaApi.
		ListObjects(request.GetContext(), *storeId).
		Body(body).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

/// ListRelations

type SdkClientListRelationsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientListRelationsRequest
	options *ClientListRelationsOptions
}

type SdkClientListRelationsRequestInterface interface {
	Options(options ClientListRelationsOptions) SdkClientListRelationsRequestInterface
	Body(body ClientListRelationsRequest) SdkClientListRelationsRequestInterface
	Execute() (*ClientListRelationsResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientListRelationsRequest
	GetOptions() *ClientListRelationsOptions
}

type ClientListRelationsRequest struct {
	User             string                     `json:"user,omitempty"`
	Object           string                     `json:"object,omitempty"`
	Relations        []string                   `json:"relations,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientListRelationsOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	MaxParallelRequests  *int32                        `json:"max_parallel_requests,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientListRelationsResponse struct {
	Relations []string `json:"response,omitempty"`
}

func (o ClientListRelationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relations"] = o.Relations
	return json.Marshal(toSerialize)
}

func (client *OpenFgaClient) ListRelations(ctx _context.Context) SdkClientListRelationsRequestInterface {
	return &SdkClientListRelationsRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientListRelationsRequest) Options(options ClientListRelationsOptions) SdkClientListRelationsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientListRelationsRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientListRelationsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientListRelationsRequest) Body(body ClientListRelationsRequest) SdkClientListRelationsRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientListRelationsRequest) Execute() (*ClientListRelationsResponse, error) {
	return request.Client.ListRelationsExecute(request)
}

func (request *SdkClientListRelationsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientListRelationsRequest) GetBody() *ClientListRelationsRequest {
	return request.body
}

func (request *SdkClientListRelationsRequest) GetOptions() *ClientListRelationsOptions {
	return request.options
}

func (client *OpenFgaClient) ListRelationsExecute(request SdkClientListRelationsRequestInterface) (*ClientListRelationsResponse, error) {
	if len(request.GetBody().Relations) <= 0 {
		return nil, fmt.Errorf("ListRelations - expected len(Relations) > 0")
	}

	batchRequestBody := ClientBatchCheckClientBody{}
	for index := 0; index < len(request.GetBody().Relations); index++ {
		batchRequestBody = append(batchRequestBody, ClientCheckRequest{
			User:             request.GetBody().User,
			Relation:         request.GetBody().Relations[index],
			Object:           request.GetBody().Object,
			Context:          request.GetBody().Context,
			ContextualTuples: request.GetBody().ContextualTuples,
		})
	}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	options := &ClientBatchCheckClientOptions{
		AuthorizationModelId: authorizationModelId,
		StoreId:              storeId,
	}
	if request.GetOptions() != nil {
		options.RequestOptions = request.GetOptions().RequestOptions
		options.Consistency = request.GetOptions().Consistency
		options.MaxParallelRequests = request.GetOptions().MaxParallelRequests
	}

	batchResponse, err := client.ClientBatchCheckExecute(&SdkClientBatchCheckClientRequest{
		ctx:     request.GetContext(),
		Client:  client,
		body:    &batchRequestBody,
		options: options,
	})

	if err != nil {
		return nil, err
	}

	var relations []string
	for index := 0; index < len(*batchResponse); index++ {
		// If any check encountered an error, return immediately
		if (*batchResponse)[index].Error != nil {
			return nil, (*batchResponse)[index].Error
		}
		if (*batchResponse)[index].GetAllowed() {
			relations = append(relations, (*batchResponse)[index].Request.Relation)
		}
	}

	return &ClientListRelationsResponse{Relations: relations}, nil
}

// / ListUsers
type SdkClientListUsersRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientListUsersRequest
	options *ClientListUsersOptions
}

type SdkClientListUsersRequestInterface interface {
	Options(options ClientListUsersOptions) SdkClientListUsersRequestInterface
	Body(body ClientListUsersRequest) SdkClientListUsersRequestInterface
	Execute() (*ClientListUsersResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientListUsersRequest
	GetOptions() *ClientListUsersOptions
}

type ClientListUsersRequest struct {
	Object           fgaSdk.FgaObject           `json:"object" yaml:"object"`
	Relation         string                     `json:"relation" yaml:"relation"`
	UserFilters      []fgaSdk.UserTypeFilter    `json:"user_filters" yaml:"user_filters"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
	// Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation.
	Context *map[string]interface{} `json:"context,omitempty" yaml:"context,omitempty"`
}

type ClientListUsersOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
}

type ClientListUsersResponse = fgaSdk.ListUsersResponse

func (client *OpenFgaClient) ListUsers(ctx _context.Context) SdkClientListUsersRequestInterface {
	return &SdkClientListUsersRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientListUsersRequest) Options(options ClientListUsersOptions) SdkClientListUsersRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientListUsersRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientListUsersRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientListUsersRequest) Body(body ClientListUsersRequest) SdkClientListUsersRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientListUsersRequest) Execute() (*ClientListUsersResponse, error) {
	return request.Client.ListUsersExecute(request)
}

func (request *SdkClientListUsersRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientListUsersRequest) GetBody() *ClientListUsersRequest {
	return request.body
}

func (request *SdkClientListUsersRequest) GetOptions() *ClientListUsersOptions {
	return request.options
}

func (client *OpenFgaClient) ListUsersExecute(request SdkClientListUsersRequestInterface) (*ClientListUsersResponse, error) {
	var contextualTuples []ClientContextualTupleKey
	if request.GetBody().ContextualTuples != nil {
		for index := 0; index < len(request.GetBody().ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (request.GetBody().ContextualTuples)[index])
		}
	}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}
	body := fgaSdk.ListUsersRequest{
		Object:               request.GetBody().Object,
		Relation:             request.GetBody().Relation,
		UserFilters:          request.GetBody().UserFilters,
		ContextualTuples:     &fgaSdk.NewContextualTupleKeys(contextualTuples).TupleKeys,
		Context:              request.GetBody().Context,
		AuthorizationModelId: authorizationModelId,
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		body.Consistency = request.GetOptions().Consistency
	}

	data, _, err := client.OpenFgaApi.
		ListUsers(request.GetContext(), *storeId).
		Body(body).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// / ReadAssertions
type SdkClientReadAssertionsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	options *ClientReadAssertionsOptions
}

type SdkClientReadAssertionsRequestInterface interface {
	Options(options ClientReadAssertionsOptions) SdkClientReadAssertionsRequestInterface
	Execute() (*ClientReadAssertionsResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetOptions() *ClientReadAssertionsOptions
}

type ClientReadAssertionsOptions struct {
	RequestOptions

	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
	StoreId              *string `json:"store_id,omitempty"`
}

type ClientReadAssertionsResponse = fgaSdk.ReadAssertionsResponse

func (client *OpenFgaClient) ReadAssertions(ctx _context.Context) SdkClientReadAssertionsRequestInterface {
	return &SdkClientReadAssertionsRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientReadAssertionsRequest) Options(options ClientReadAssertionsOptions) SdkClientReadAssertionsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientReadAssertionsRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientReadAssertionsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientReadAssertionsRequest) Execute() (*ClientReadAssertionsResponse, error) {
	return request.Client.ReadAssertionsExecute(request)
}

func (request *SdkClientReadAssertionsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientReadAssertionsRequest) GetOptions() *ClientReadAssertionsOptions {
	return request.options
}

func (client *OpenFgaClient) ReadAssertionsExecute(request SdkClientReadAssertionsRequestInterface) (*ClientReadAssertionsResponse, error) {
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	if authorizationModelId == nil || *authorizationModelId == "" {
		return nil, FgaRequiredParamError{param: "AuthorizationModelId"}
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	data, _, err := client.OpenFgaApi.
		ReadAssertions(request.GetContext(), *storeId, *authorizationModelId).
		Options(requestOptions).
		Execute()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// / WriteAssertions
type SdkClientWriteAssertionsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientWriteAssertionsRequest
	options *ClientWriteAssertionsOptions
}

type SdkClientWriteAssertionsRequestInterface interface {
	Options(options ClientWriteAssertionsOptions) SdkClientWriteAssertionsRequestInterface
	Body(body ClientWriteAssertionsRequest) SdkClientWriteAssertionsRequestInterface
	Execute() (*ClientWriteAssertionsResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientWriteAssertionsRequest
	GetOptions() *ClientWriteAssertionsOptions
}

type ClientAssertion struct {
	User             string                     `json:"user,omitempty"`
	Relation         string                     `json:"relation,omitempty"`
	Object           string                     `json:"object,omitempty"`
	Expectation      bool                       `json:"expectation,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientWriteAssertionsRequest = []ClientAssertion

func (clientAssertion ClientAssertion) ToAssertion() fgaSdk.Assertion {
	assertion := fgaSdk.Assertion{
		TupleKey: fgaSdk.AssertionTupleKey{
			User:     clientAssertion.User,
			Relation: clientAssertion.Relation,
			Object:   clientAssertion.Object,
		},
		Expectation: clientAssertion.Expectation,
	}
	if clientAssertion.Context != nil {
		assertion.Context = clientAssertion.Context
	}
	if clientAssertion.ContextualTuples != nil {
		assertion.ContextualTuples = &clientAssertion.ContextualTuples
	}
	return assertion
}

type ClientWriteAssertionsOptions struct {
	RequestOptions

	AuthorizationModelId *string `json:"authorization_model_id,omitempty"`
	StoreId              *string `json:"store_id,omitempty"`
}

type ClientWriteAssertionsResponse struct {
}

func (client *OpenFgaClient) WriteAssertions(ctx _context.Context) SdkClientWriteAssertionsRequestInterface {
	return &SdkClientWriteAssertionsRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientWriteAssertionsRequest) Options(options ClientWriteAssertionsOptions) SdkClientWriteAssertionsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientWriteAssertionsRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientWriteAssertionsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientWriteAssertionsRequest) Body(body ClientWriteAssertionsRequest) SdkClientWriteAssertionsRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientWriteAssertionsRequest) Execute() (*ClientWriteAssertionsResponse, error) {
	return request.Client.WriteAssertionsExecute(request)
}

func (request *SdkClientWriteAssertionsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientWriteAssertionsRequest) GetBody() *ClientWriteAssertionsRequest {
	return request.body
}

func (request *SdkClientWriteAssertionsRequest) GetOptions() *ClientWriteAssertionsOptions {
	return request.options
}

func (client *OpenFgaClient) WriteAssertionsExecute(request SdkClientWriteAssertionsRequestInterface) (*ClientWriteAssertionsResponse, error) {
	writeAssertionsRequest := fgaSdk.WriteAssertionsRequest{}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	if authorizationModelId == nil || *authorizationModelId == "" {
		return nil, FgaRequiredParamError{param: "AuthorizationModelId"}
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}
	for index := 0; index < len(*request.GetBody()); index++ {
		clientAssertion := (*request.GetBody())[index]
		writeAssertionsRequest.Assertions = append(writeAssertionsRequest.Assertions, clientAssertion.ToAssertion())
	}

	requestOptions := RequestOptions{}
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
	}

	_, err = client.OpenFgaApi.
		WriteAssertions(request.GetContext(), *storeId, *authorizationModelId).
		Body(writeAssertionsRequest).
		Options(requestOptions).
		Execute()

	if err != nil {
		return nil, err
	}
	return &ClientWriteAssertionsResponse{}, nil
}

type SdkClientStreamedListObjectsRequest struct {
	ctx    _context.Context
	Client *OpenFgaClient

	body    *ClientStreamedListObjectsRequest
	options *ClientStreamedListObjectsOptions
}

type SdkClientStreamedListObjectsRequestInterface interface {
	Options(options ClientStreamedListObjectsOptions) SdkClientStreamedListObjectsRequestInterface
	Body(body ClientStreamedListObjectsRequest) SdkClientStreamedListObjectsRequestInterface
	Execute() (*ClientStreamedListObjectsResponse, error)
	GetAuthorizationModelIdOverride() *string
	GetStoreIdOverride() *string

	GetContext() _context.Context
	GetBody() *ClientStreamedListObjectsRequest
	GetOptions() *ClientStreamedListObjectsOptions
}

type ClientStreamedListObjectsRequest struct {
	User             string                     `json:"user,omitempty"`
	Relation         string                     `json:"relation,omitempty"`
	Type             string                     `json:"type,omitempty"`
	Context          *map[string]interface{}    `json:"context,omitempty"`
	ContextualTuples []ClientContextualTupleKey `json:"contextual_tuples,omitempty"`
}

type ClientStreamedListObjectsOptions struct {
	RequestOptions

	AuthorizationModelId *string                       `json:"authorization_model_id,omitempty"`
	StoreId              *string                       `json:"store_id,omitempty"`
	Consistency          *fgaSdk.ConsistencyPreference `json:"consistency,omitempty"`
	// StreamBufferSize configures the buffer size for streaming response channels.
	// A larger buffer improves throughput for high-volume streams but increases memory usage.
	// A smaller buffer reduces memory usage but may decrease throughput.
	// Defaults to 10 if not specified or if set to 0.
	StreamBufferSize *int `json:"stream_buffer_size,omitempty"`
}

type ClientStreamedListObjectsResponse struct {
	Objects <-chan fgaSdk.StreamedListObjectsResponse
	Errors  <-chan error
	close   func()
}

func (r *ClientStreamedListObjectsResponse) Close() {
	if r.close != nil {
		r.close()
	}
}

func (client *OpenFgaClient) streamedListObjects(ctx _context.Context) SdkClientStreamedListObjectsRequestInterface {
	return &SdkClientStreamedListObjectsRequest{
		Client: client,
		ctx:    ctx,
	}
}

func (request *SdkClientStreamedListObjectsRequest) Options(options ClientStreamedListObjectsOptions) SdkClientStreamedListObjectsRequestInterface {
	request.options = &options
	return request
}

func (request *SdkClientStreamedListObjectsRequest) GetAuthorizationModelIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.AuthorizationModelId
}

func (request *SdkClientStreamedListObjectsRequest) GetStoreIdOverride() *string {
	if request.options == nil {
		return nil
	}
	return request.options.StoreId
}

func (request *SdkClientStreamedListObjectsRequest) Body(body ClientStreamedListObjectsRequest) SdkClientStreamedListObjectsRequestInterface {
	request.body = &body
	return request
}

func (request *SdkClientStreamedListObjectsRequest) Execute() (*ClientStreamedListObjectsResponse, error) {
	return request.Client.streamedListObjectsExecute(request)
}

func (request *SdkClientStreamedListObjectsRequest) GetContext() _context.Context {
	return request.ctx
}

func (request *SdkClientStreamedListObjectsRequest) GetBody() *ClientStreamedListObjectsRequest {
	return request.body
}

func (request *SdkClientStreamedListObjectsRequest) GetOptions() *ClientStreamedListObjectsOptions {
	return request.options
}

func (client *OpenFgaClient) streamedListObjectsExecute(request SdkClientStreamedListObjectsRequestInterface) (*ClientStreamedListObjectsResponse, error) {
	if request.GetBody() == nil {
		return nil, FgaRequiredParamError{param: "body"}
	}
	var contextualTuples []ClientContextualTupleKey
	if request.GetBody().ContextualTuples != nil {
		for index := 0; index < len(request.GetBody().ContextualTuples); index++ {
			contextualTuples = append(contextualTuples, (request.GetBody().ContextualTuples)[index])
		}
	}
	authorizationModelId, err := client.getAuthorizationModelId(request.GetAuthorizationModelIdOverride())
	if err != nil {
		return nil, err
	}
	storeId, err := client.getStoreId(request.GetStoreIdOverride())
	if err != nil {
		return nil, err
	}
	body := fgaSdk.ListObjectsRequest{
		User:                 request.GetBody().User,
		Relation:             request.GetBody().Relation,
		Type:                 request.GetBody().Type,
		ContextualTuples:     fgaSdk.NewContextualTupleKeys(contextualTuples),
		Context:              request.GetBody().Context,
		AuthorizationModelId: authorizationModelId,
	}
	requestOptions := RequestOptions{}
	bufferSize := 0
	if request.GetOptions() != nil {
		requestOptions = request.GetOptions().RequestOptions
		body.Consistency = request.GetOptions().Consistency
		if request.GetOptions().StreamBufferSize != nil {
			bufferSize = *request.GetOptions().StreamBufferSize
		}
	}

	channel, err := fgaSdk.ExecuteStreamedListObjectsWithBufferSize(
		&client.APIClient,
		request.GetContext(),
		*storeId,
		body,
		requestOptions,
		bufferSize,
	)

	if err != nil {
		return nil, err
	}

	return &ClientStreamedListObjectsResponse{
		Objects: channel.Objects,
		Errors:  channel.Errors,
		close:   channel.Close,
	}, nil
}
