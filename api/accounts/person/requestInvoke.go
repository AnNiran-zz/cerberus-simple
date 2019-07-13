package api

import (
	"cerberus/app/person"
	"cerberus/services/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostCreatePersonAccountDataRequest(c *gin.Context) {

	requesterId := c.Request.URL.Query().Get("requesterid")
	recipientId := c.Request.URL.Query().Get("recipientid")

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	_, data, err := person.CreateAccountDataRequest(requesterId, recipientId, argumentsAsArray)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, string(data))
}

func PostCreatePersonDocumentDataRequest(c *gin.Context) {

	requesterId := c.Request.URL.Query().Get("requesterid")
	recipientId := c.Request.URL.Query().Get("recipientid")
	documentName := c.Request.URL.Query().Get("documentname")
	documentCopy := c.Request.URL.Query().Get("copy")

	var copy bool
	if documentCopy == "true" {
		copy = true
	} else {
		copy = false
	}

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	_, data, err := person.CreateDocumentDataRequest(requesterId, recipientId, documentName, argumentsAsArray, copy)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, string(data))
}

func PostAcceptPersonAccountDataRequest(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("requestid")
	recipientId := c.Request.URL.Query().Get("recipientid")

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	_, data, err := person.AcceptAccountDataRequest(recipientId, requestId, argumentsAsArray)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func PostAcceptPersonDocumentDataRequest(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("requestid")
	recipientId := c.Request.URL.Query().Get("recipientid")
	documentCopy := c.Request.URL.Query().Get("copy")

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	_, data, _, err := person.AcceptDocumentDataRequest(recipientId, requestId, documentCopy, argumentsAsArray)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func PostRejectPersonAccountDataRequest(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("requestid")
	recipientId := c.Request.URL.Query().Get("recipientid")

	_, data, err := person.RejectAccountDataRequest(recipientId, requestId)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func PostRejectPersonDocumentDataRequest(c *gin.Context) {

	requestId := c.Request.URL.Query().Get("requestid")
	recipientId := c.Request.URL.Query().Get("recipientid")

	_, data, err := person.RejectDocumentDataRequest(recipientId, requestId)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func PostUpdatePersonAccountDataRequest(c *gin.Context) {

	requesterId := c.Request.URL.Query().Get("requesterid")
	recipientId := c.Request.URL.Query().Get("recipientid")
	requestId := c.Request.URL.Query().Get("requestid")

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	_, data, err := person.UpdateAccountDataRequest(requesterId, recipientId, requestId, argumentsAsArray)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, string(data))
}

func PostUpdatePersonDocumentDataRequest(c *gin.Context) {

	requesterId := c.Request.URL.Query().Get("requesterid")
	recipientId := c.Request.URL.Query().Get("recipientid")
	requestId := c.Request.URL.Query().Get("requestid")
	documentName := c.Request.URL.Query().Get("documentName")
	documentCopy := c.Request.URL.Query().Get("copy")

	arguments := c.Request.URL.Query().Get("arguments")
	argumentsAsArray := strings.Split(arguments, ",")

	var copy bool
	if documentCopy == "true" {
		copy = true
	} else {
		copy = false
	}

	_, data, err := person.UpdateDocumentDataRequest(requesterId, recipientId, requestId, documentName, argumentsAsArray, copy)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, string(data))
}
