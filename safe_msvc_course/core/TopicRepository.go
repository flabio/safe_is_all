package core

import (
	"sync"

	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/insfractruture/utils"
)

func NewTopicRepository() uicore.UITopicCore {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (c *OpenConnection) GetTopicFindAll() ([]entities.Topic, error) {
	var topicEntities []entities.Topic
	c.mux.Lock()
	result := c.connection.Preload("Course").Order(utils.DB_ORDER_DESC).Find(&topicEntities)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return topicEntities, result.Error
}
func (c *OpenConnection) GetTopicFindById(id uint) (entities.Topic, error) {
	var topic entities.Topic
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&topic)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return topic, result.Error
}

func (c *OpenConnection) CreateTopic(topic entities.Topic) (entities.Topic, error) {
	c.mux.Lock()
	err := c.connection.Create(&topic).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return topic, err
}
func (c *OpenConnection) UpdateTopic(id uint, topic entities.Topic) (entities.Topic, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&topic).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return topic, err
}

func (c *OpenConnection) DeleteTopic(id uint) (bool, error) {
	c.mux.Lock()
	var topic entities.Topic
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&topic).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}
