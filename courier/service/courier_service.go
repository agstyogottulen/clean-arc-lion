package service

import (
	"github.com/agstyogottulen/clean-arc-lion/common"
	"github.com/agstyogottulen/clean-arc-lion/courier/repository"
	"github.com/agstyogottulen/clean-arc-lion/models"
)

type Service interface {
	Create(req *models.Courier) (map[string]interface{}, error)
	Read(id int) (map[string]interface{}, error)
	ReadAll() (map[string]interface{}, error)
	Update(id int, req *models.Courier) (map[string]interface{}, error)
	Delete(id int) (map[string]interface{}, error)
}

type CourierService struct {
	CourierRepository repository.Repository
}

func (c *CourierService) Create(req *models.Courier) (map[string]interface{}, error) {
	response, err := c.CourierRepository.Create(req)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "create courier data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (c *CourierService) Read(id int) (map[string]interface{}, error) {
	response, err := c.CourierRepository.Read(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "read courier data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (c *CourierService) ReadAll() (map[string]interface{}, error) {
	response, err := c.CourierRepository.ReadAll()
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "read all courier data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (c *CourierService) Update(id int, req *models.Courier) (map[string]interface{}, error) {
	response, err := c.CourierRepository.Update(id, req)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "update courier data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (c *CourierService) Delete(id int) (map[string]interface{}, error) {
	_, err := c.CourierRepository.Delete(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "delete courier data: success")
	return mapResponse, nil
}

func NewCourierService(courierRepository repository.Repository) Service {
	return &CourierService{
		CourierRepository: courierRepository,
	}
}
