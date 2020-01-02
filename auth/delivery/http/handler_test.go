package http

import (
	"BotLeha/Oleksii-bot/auth/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestSignUp(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Username, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("Post", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSignIn(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignIn", signUpBody.Username, signUpBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}\n", w.Body.String())
}
