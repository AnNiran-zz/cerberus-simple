package api

import (
	"cerberus/app/person"
	"cerberus/services/logger"

	"github.com/gin-gonic/gin"
)

func GetPersonAccountRequestsBySelector(c *gin.Context) {

	requestType := c.Request.URL.Query().Get("type")
	selectorKey := c.Request.URL.Query().Get("selector")
	value := c.Request.URL.Query().Get("value")

	data, err := person.GetRequestsObjectsBySelector(requestType, selectorKey, value)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestsPublicIdsBySelector(c *gin.Context) {

	requestType := c.Request.URL.Query().Get("type")
	selectorKey := c.Request.URL.Query().Get("selector")
	value := c.Request.URL.Query().Get("value")

	data, err := person.GetRequestsPublicIdsBySelector(requestType, selectorKey, value)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestsByRecipient(c *gin.Context) {

	queryType := c.Request.URL.Query().Get("query")
	requestType := c.Request.URL.Query().Get("type")
	recipientId := c.Request.URL.Query().Get("recipientid")

	data, err := person.GetRequestsByRecipient(queryType, requestType, recipientId)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestsByRequester(c *gin.Context) {

	queryType := c.Request.URL.Query().Get("query")
	requestType := c.Request.URL.Query().Get("type")
	requesterId := c.Request.URL.Query().Get("requesterid")

	data, err := person.GetRequestsByRequester(queryType, requestType, requesterId)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestsByDocumentName(c *gin.Context) {

	queryType := c.Request.URL.Query().Get("query")
	documentName := c.Request.URL.Query().Get("documentname")

	data, err := person.GetRequestsByDocumentName(queryType, documentName)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestsByStatus(c *gin.Context) {

	queryType := c.Request.URL.Query().Get("query")
	requestType := c.Request.URL.Query().Get("type")
	status := c.Request.URL.Query().Get("status")

	data, err := person.GetRequestsByStatus(queryType, requestType, status)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestObject(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("id")
	idType := c.Request.URL.Query().Get("type")

	data, err := person.GetRequestObject(idType, requestId)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetPersonAccountRequestPublicId(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("id")
	idType := c.Request.URL.Query().Get("type")

	data, err := person.GetRequestPublicId(requestId, idType)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}
