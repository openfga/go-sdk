package telemetry

import (
	"testing"
)

type MockCheckRequestTupleKey struct {
	user *string
}

func (m *MockCheckRequestTupleKey) GetUser() *string {
	return m.user
}

type MockRequestAuthorizationModelId struct {
	authorizationModelId *string
}

func (m *MockRequestAuthorizationModelId) GetAuthorizationModelId() *string {
	return m.authorizationModelId
}

type MockCheckRequest struct {
	MockCheckRequestTupleKey
	MockRequestAuthorizationModelId
}

func (m *MockCheckRequest) GetTupleKey() *MockCheckRequestTupleKey {
	return &m.MockCheckRequestTupleKey
}

func TestCheckRequestInterfaceImplementation(t *testing.T) {
	user := "test-user"
	modelId := "test-model-id"

	mockCheckRequest := &MockCheckRequest{
		MockCheckRequestTupleKey:        MockCheckRequestTupleKey{user: &user},
		MockRequestAuthorizationModelId: MockRequestAuthorizationModelId{authorizationModelId: &modelId},
	}

	if mockCheckRequest.GetTupleKey().GetUser() != &user {
		t.Errorf("Expected GetUser to return '%s', but got '%s'", user, *mockCheckRequest.GetTupleKey().GetUser())
	}

	if mockCheckRequest.GetAuthorizationModelId() != &modelId {
		t.Errorf("Expected GetAuthorizationModelId to return '%s', but got '%s'", modelId, *mockCheckRequest.GetAuthorizationModelId())
	}
}

func TestCheckRequestInterfaceNilValues(t *testing.T) {
	mockCheckRequest := &MockCheckRequest{}

	if mockCheckRequest.GetTupleKey().GetUser() != nil {
		t.Errorf("Expected GetUser to return nil, but got '%s'", *mockCheckRequest.GetTupleKey().GetUser())
	}

	if mockCheckRequest.GetAuthorizationModelId() != nil {
		t.Errorf("Expected GetAuthorizationModelId to return nil, but got '%s'", *mockCheckRequest.GetAuthorizationModelId())
	}
}
