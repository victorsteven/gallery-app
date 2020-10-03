package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ImageInfo struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	//ImgUrl    string     `gorm:"text;not null;" json:"imgUrl" validate:"required"`
	ImageId    uint64     `gorm:"size:100;not null;" json:"imageId" validate:"required"`
	UserIp    string     `gorm:"size:100;not null;" json:"userIp" validate:"required"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type service struct {
	db *gorm.DB
}

type ImageService interface {
	SaveImageInfo(info *ImageInfo) (*ImageInfo, error)
	GetImageInfo(imageId uint64) (*ImageInfo, error)
	DeleteImageInfo(imageId uint64) error
}


func NewImageService(db *gorm.DB) *service {
	return &service{db}
}

var _ ImageService = &service{}

func (r *service) SaveImageInfo(info *ImageInfo) (*ImageInfo, error) {
	err := r.db.Debug().Create(&info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (r *service) GetImageInfo(imageId uint64) (*ImageInfo, error) {

	var imgInfo = &ImageInfo{}
	err := r.db.Debug().Where("image_id = ?", imageId).Take(&imgInfo).Error
	if err != nil {
		return nil, err
	}
	return imgInfo, nil
}

func (r *service) DeleteImageInfo(imageId uint64) error{

	var imgInfo ImageInfo
	err := r.db.Debug().Where("image_id = ?", imageId).Delete(&imgInfo).Error
	if err != nil {
		return err
	}
	return nil
}



