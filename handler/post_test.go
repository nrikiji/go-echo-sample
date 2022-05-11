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

func TestGetPosts(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.GET, "/api/posts", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	assert.NoError(t, h.getPosts(c))

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res postsResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(res.Posts))
	}
}

func TestUpdatePostSuccess(t *testing.T) {
	setup()

	reqJSON := `{"title":"NewTitle", "body":"NewBody"}`
	authMiddleware := middleware.AuthMiddleware(h.authStore, h.userStore)

	req := httptest.NewRequest(echo.PATCH, "/api/posts/:id", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer: ValidToken1")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := authMiddleware(func(c echo.Context) error {
		return h.updatePost(c)
	})(c)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var res postResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "NewTitle", res.Post.Title)
		assert.Equal(t, "NewBody", res.Post.Body)
	}
}

func TestUpdatePostForbidden(t *testing.T) {
	setup()

	reqJSON := `{"title":"NewTitle", "body":"NewBody"}`
	authMiddleware := middleware.AuthMiddleware(h.authStore, h.userStore)

	req := httptest.NewRequest(echo.PATCH, "/api/posts/:id", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer: ValidToken1")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := authMiddleware(func(c echo.Context) error {
		return h.updatePost(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}
