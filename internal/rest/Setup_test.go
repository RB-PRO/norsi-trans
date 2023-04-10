package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ready-steady/assert"
)

func TestSetup(t *testing.T) {

	// Создаём экзепляр API
	API, ErrorSetupAPI := setup()
	if ErrorSetupAPI != nil {
		t.Error(ErrorSetupAPI)
	}

	w := httptest.NewRecorder()

	// Ping
	ReqPing, ErrorReq := http.NewRequest("GET", "/ping", nil)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	API.R.ServeHTTP(w, ReqPing)
	assert.Equal(200, w.Code, t)
	assert.Equal("pong", w.Body.String(), t)
	w.Body.Reset()

	// Setup test Data
	NewUser := "roman"
	NewPass := "password"

	// ADD
	ReqPost, ErrorReq := http.NewRequest("POST", "/add?user="+NewUser+"&pass="+NewPass, nil)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	API.R.ServeHTTP(w, ReqPost)
	assert.Equal(200, w.Code, t)
	assert.Equal(`{"pass":"`+NewPass+`","user":"`+NewUser+`"}`, w.Body.String(), t)
	w.Body.Reset()

	// GET
	ReqGet, ErrorReq := http.NewRequest("GET", "/user/"+NewUser, nil)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	API.R.ServeHTTP(w, ReqGet)
	assert.Equal(200, w.Code, t)
	assert.Equal(`{"pass":"`+NewPass+`","user":"`+NewUser+`"}`, w.Body.String(), t)
	w.Body.Reset()

	// Delete
	ReqDelete, ErrorReq := http.NewRequest("DELETE", "/remove/"+NewUser, nil)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	API.R.ServeHTTP(w, ReqDelete)
	assert.Equal(200, w.Code, t)
	assert.Equal(`{"user":"`+NewUser+`"}`, w.Body.String(), t)
	w.Body.Reset()

	// Get
	API.R.ServeHTTP(w, ReqGet)
	assert.Equal(200, w.Code, t)
	assert.Equal(`{"status":"Get: not found this user","user":"`+NewUser+`"}`, w.Body.String(), t)
	w.Body.Reset()
}
