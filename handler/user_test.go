package handler

import (
	"encoding/json"
	"go-echo-starter/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.GET, "/api/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	assert.NoError(t, h.getUsers(c))

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res usersResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(res.Users))
	}
}

func TestGetUserSuccess(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.GET, "/api/users/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.getUser(c))

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res userResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "User1", res.User.Name)
	}
}

func TestGetUserNotFound(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.GET, "/api/users/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	assert.NoError(t, h.getUser(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestRegisterUserSuccess(t *testing.T) {
	setup()
	reqJSON := `{"name":"NewUser","token":"ValidToken"}`
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.registerUser(c))

	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var res userResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "NewUser", res.User.Name)
	}
}

func TestRegisterUserDuplicateUid(t *testing.T) {
	setup()
	reqJSON := `{"name":"NewUser","token":"ValidToken"}`
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.registerUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	reqJSON2 := `{"name":"NewUser","token":"ValidToken"}`
	req2 := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON2))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()

	c2 := e.NewContext(req2, rec2)

	assert.NoError(t, h.registerUser(c2))
	assert.Equal(t, http.StatusUnprocessableEntity, rec2.Code)
}

func TestRegisterUserInvalidToken(t *testing.T) {
	setup()
	reqJSON := `{"name":"NewUser","token":"InValidToken"}`
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.registerUser(c))
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestRegisterUserBadRequest(t *testing.T) {
	setup()
	reqJSON := `{"_name":"NewUser","token":"ValidToken"}`
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.registerUser(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLoginUserSuccess(t *testing.T) {
	setup()
	reqJSON := `{"token":"ValidToken1"}`
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.loginUser(c))

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res userResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "User1", res.User.Name)
	}
}

func TestLoginUserNotFound(t *testing.T) {
	setup()
	reqJSON := `{"token":"ValidToken"}`
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, h.loginUser(c))
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

}

func TestUpdateUserSuccess(t *testing.T) {
	setup()
	reqJSON := `{"name":"UpdateUser1"}`
	authMiddleware := middleware.AuthMiddleware(h.authStore, h.dataStore)
	req := httptest.NewRequest(echo.PUT, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer: ValidToken1")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := authMiddleware(func(c echo.Context) error {
		return h.updateUser(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res userResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "UpdateUser1", res.User.Name)
	}
}

func TestUpdateUserUnauthorized(t *testing.T) {
	setup()
	reqJSON := `{"name":"UpdateUser1"}`
	authMiddleware := middleware.AuthMiddleware(h.authStore, h.dataStore)
	req := httptest.NewRequest(echo.PUT, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer: InValidToken")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := authMiddleware(func(c echo.Context) error {
		return h.updateUser(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
