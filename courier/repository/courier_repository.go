package repository

import (
	"errors"

	"github.com/agstyogottulen/clean-arc-lion/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(req *models.Courier) (*models.Courier, error)
	Read(id int) (*models.Courier, error)
	ReadAll() ([]*models.Courier, error)
	Update(id int, req *models.Courier) (*models.Courier, error)
	Delete(id int) (*models.Courier, error)
}

type CourierRepository struct {
	Conn *gorm.DB
}

func (c *CourierRepository) Create(req *models.Courier) (*models.Courier, error) {
	if err := c.Conn.Table("couriers").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("add courier data: error")
	}

	return req, nil
}

func (c *CourierRepository) Read(id int) (*models.Courier, error) {
	courier := new(models.Courier)

	if err := c.Conn.Table("couriers").Where("id = ?", id).First(&courier).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("get courier data: error")
	}

	return courier, nil
}

func (c *CourierRepository) ReadAll() ([]*models.Courier, error) {
	courierList := make([]*models.Courier, 0)

	if err := c.Conn.Table("couriers").Find(&courierList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("get courier list data: error")
	}

	return courierList, nil
}

func (c *CourierRepository) Update(id int, req *models.Courier) (*models.Courier, error) {
	courier := new(models.Courier)

	if err := c.Conn.Table("couriers").Where("id = ?", id).First(&courier).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("update courier data: error")
	}

	return courier, nil
}

func (c *CourierRepository) Delete(id int) (*models.Courier, error) {
	if err := c.Conn.Table("couriers").Where("id = ?", id).Delete(&models.Courier{}).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("delete courier data: error")
	}

	return nil, nil
}

func NewCourierRepository(conn *gorm.DB) Repository {
	return &CourierRepository{
		Conn: conn,
	}
}
