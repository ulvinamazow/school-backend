package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	F "github.com/stretchr/testify/assert"
	"github.com/ulvinamazow/studentAPI/config"
	"github.com/ulvinamazow/studentAPI/models"
	"gorm.io/gorm"
)

func set_CreateUser_router(dbase *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	r.POST("/user", CreateUser)

	req, err := http.NewRequest("POST", "/user", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestCreateUser(t *testing.T) {
	a := F.New(t)
	config.ConnnectDatabase()
	db := config.GetDB()

	input_model := models.UserTest{
		Username: "testuser",
		Name:     "Test",
		Surname:  "User",
		Email:    "testuser@gmail.com",
		RoleID:   10,
	}
	reqBody, err := json.Marshal(input_model)
	if err != nil {
		a.Error(err)
	}

	req, w, er := set_CreateUser_router(db, bytes.NewBuffer(reqBody))
	if er != nil {
		a.Error(er)
	}

	a.Equal(http.MethodPost, req.Method, "Request method should be POST")
	body, err := io.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.UserTest{}
	if err = json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	excepted := input_model
	fmt.Printf("Excepted model: %v\n Response model: %v\n", excepted, actual)
	a.Equal(excepted, actual)

}
