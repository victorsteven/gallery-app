package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gallery/server/domain"
	"gallery/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type fakeImageInfoService struct {
	SaveImageInfoFn func(info *domain.ImageInfo) (*domain.ImageInfo, error)
	GetImageInfoFn func(imageId uint64) (*domain.ImageInfo, error)
	DeleteImageInfoFn func(id uint64) error
}

func (u *fakeImageInfoService) SaveImageInfo(info *domain.ImageInfo) (*domain.ImageInfo, error) {
	return u.SaveImageInfoFn(info)
}

func (u *fakeImageInfoService) GetImageInfo(imageId uint64) (*domain.ImageInfo, error) {
	return u.GetImageInfoFn(imageId)
}

func (u *fakeImageInfoService) DeleteImageInfo(id uint64) error {
	return u.DeleteImageInfoFn(id)
}

var (
	fakeImageInfo fakeImageInfoService

	h = handler.NewHandlerService(&fakeImageInfo)
)

type FailureResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type SaveImageInfoResponse struct {
	Status int    `json:"status"`
	Body   domain.ImageInfo `json:"body"`
}

type DeleteImageInfoResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type GetImageInfoResponse struct {
	Status int    `json:"status"`
	Body   bool `json:"body"`
}


//We dont need to mock the domain layer, because we will never get there.
func TestSaveImageInfo_WrongInput(t *testing.T) {

	inputJSON := `{"imageId": 1, "userIp": 123}` //integer is given as the userIp instead of string

	req, err := http.NewRequest(http.MethodPost, "/image_info", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.POST("/image_info", h.SaveImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response FailureResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 400)
	assert.EqualValues(t, response.Status, 400)
	assert.EqualValues(t, response.Body, "please provide valid inputs")
}

func TestSaveImageInfo_RequiredNotCorrect(t *testing.T) {

	inputJSON := `{"imageId": 1, "userIp": ""}` //the userIp is required, but left empty

	req, err := http.NewRequest(http.MethodPost, "/image_info", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.POST("/image_info", h.SaveImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = FailureResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 422)
	assert.EqualValues(t, response.Status, 422)
	assert.EqualValues(t, response.Body, "The userIp is required")
}


func TestSaveImageInfo_Success(t *testing.T) {

	//fake the domain method
	fakeImageInfo.SaveImageInfoFn = func(cred *domain.ImageInfo) (*domain.ImageInfo, error) {
		return cred, nil
	}
	
	inputJSON := `{"imageId": 1, "userIp": "127.0.0.1"}` //correct details

	req, err := http.NewRequest(http.MethodPost, "/image_info", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.POST("/image_info", h.SaveImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = SaveImageInfoResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Status, 200)
	assert.EqualValues(t, response.Body.ImageId, 1)
	assert.EqualValues(t, response.Body.UserIp, "127.0.0.1")

}

func TestSaveImageInfo_Failure(t *testing.T) {

	//Lets assume that the database is down
	fakeImageInfo.SaveImageInfoFn = func(cred *domain.ImageInfo) (*domain.ImageInfo, error) {
		return nil, errors.New("something went wrong saving image info")
	}

	inputJSON := `{"imageId": 1, "userIp": "127.0.0.1"}` //correct details

	req, err := http.NewRequest(http.MethodPost, "/image_info", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.POST("/image_info", h.SaveImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = FailureResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 500)
	assert.EqualValues(t, response.Status, 500)
	assert.EqualValues(t, response.Body, "something went wrong saving image info")
}


func TestDeleteImageInfo_RequiredNotCorrect(t *testing.T) {

	id := "abc"

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/image_info/%s", id), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.DELETE("/image_info/:id", h.DeleteImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = FailureResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 400)
	assert.EqualValues(t, response.Status, 400)
	assert.EqualValues(t, response.Body, "Invalid id provided")
}

func TestGetImageInfo_True(t *testing.T) {

	imgDomain := &domain.ImageInfo{
		ID:        1,
		ImageId:   1,
		UserIp:    "127.0.0.1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//fake the domain method
	fakeImageInfo.GetImageInfoFn = func(imageId uint64) (*domain.ImageInfo, error) {
		return  imgDomain, nil
	}

	imageId := 1

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/image_info/%d", imageId), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/image_info/:imageId", h.GetImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = GetImageInfoResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Status, 200)
	assert.EqualValues(t, response.Body, true)
}


func TestGetImageInfo_False(t *testing.T) {

	//fake the domain method
	fakeImageInfo.GetImageInfoFn = func(imageId uint64) (*domain.ImageInfo, error) {
		return  nil, nil
	}

	imageId := 1

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/image_info/%d", imageId), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/image_info/:imageId", h.GetImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = GetImageInfoResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Status, 200)
	assert.EqualValues(t, response.Body, false)
}


func TestDeleteImageInfo_Success(t *testing.T) {

	//fake the domain method
	fakeImageInfo.DeleteImageInfoFn = func(id uint64) error {
		return  nil
	}

	imageId := 1

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/image_info/%d", imageId), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.DELETE("/image_info/:imageId", h.DeleteImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = DeleteImageInfoResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Status, 200)
	assert.EqualValues(t, response.Body, "success")
}

func TestDeleteImageInfo_Failure(t *testing.T) {

	//Lets assume that the database is down
	fakeImageInfo.DeleteImageInfoFn = func(id uint64)  error {
		return errors.New("something went wrong deleting image info")
	}

	imageId := 1

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/image_info/%d", imageId), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.DELETE("/image_info/:imageId", h.DeleteImageInfo)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = FailureResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 500)
	assert.EqualValues(t, response.Status, 500)
	assert.EqualValues(t, response.Body, "something went wrong deleting image info")
}



