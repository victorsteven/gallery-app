package domain_test

import (
	"gallery/server/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveImageInfo_Success(t *testing.T) {
	conn, err := domain.DBConn()
	if err != nil {
		t.Fatal(err)
	}

	var info = domain.ImageInfo{
		ID:        1,
		ImageId:    1,
		UserIp:    "234.3423.34",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repo := domain.NewImageService(conn)

	u, saveErr := repo.SaveImageInfo(&info)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.ImageId, 1)
	assert.EqualValues(t, u.UserIp, "234.3423.34")
}

func TestGetImageInfo_Success(t *testing.T) {

	conn, err := domain.DBConn()
	if err != nil {
		t.Fatal(err)
	}

	imageId := uint64(1)

	info, err := domain.SeedImageInfo(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	repo := domain.NewImageService(conn)

	u, saveErr := repo.GetImageInfo(imageId)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.ImageId, info.ImageId)
	assert.EqualValues(t, u.UserIp, info.UserIp)
}


func TestDeleteImageInfo_Success(t *testing.T) {
	conn, err := domain.DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	info, err := domain.SeedImageInfo(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := domain.NewImageService(conn)
	delErr := repo.DeleteImageInfo(info.ID)

	assert.Nil(t, delErr)
}


