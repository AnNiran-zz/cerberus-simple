package main

import (
	institution "cerberus/api/accounts/institution"
	person "cerberus/api/accounts/person"
	"cerberus/services/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Log("api_endpoints")
	logger.LogEvent("Listening for api requests", nil)

	// gin
	router := gin.New()

	personAccnts := router.Group("/person")
	institutionAccnts := router.Group("/institution")

	// http://159.203.35.156:8000/person/get/account/data
	// request params keys: id
	personAccnts.GET("/get/account/data", person.GetPersonAccountData) // ok

	// http://159.203.35.156:8000/person/get/account/history
	// request params keys: id
	personAccnts.GET("/get/account/history", person.GetPersonAccountHistory) // ok

	// http://159.203.35.156:8000/person/get/document/data
	// request params keys: id. documentname
	personAccnts.GET("/get/document/data", person.GetPersonAccountDocumentData) // ok

	// http://159.203.35.156:8000/person/get/document/version
	// request params keys: id, documentname, version
	personAccnts.GET("/get/document/version", person.GetPersonAccountDocumentVersion) // ok

	// http://159.203.35.156:8000/person/get/document/versions
	// request params keys: id, documentname
	personAccnts.GET("/get/document/versions", person.GetPersonAccountDocumentVersions) // ok

	// person: query multiple accounts data

	// http://159.203.35.156:8000/person/get/accounts/email
	// request params keys: email
	personAccnts.GET("/get/accounts/email", person.GetPersonAccountsByEmail) // ok

	// http://159.203.35.156:8000/person/get/accounts/firstname
	// request params keys: firstname
	personAccnts.GET("/get/accounts/firstname", person.GetPersonAccountsByFirstName) // ok

	// http://159.203.35.156:8000/person/get/accounts/lastname
	// request params keys: lastname
	personAccnts.GET("/get/accounts/lastname", person.GetPersonAccountsByLastName)

	// person: invoke
	// http://159.203.35.156:8000/person/create/account
	// request params keys: firstname, lastname, email, phone
	personAccnts.POST("/create/account", person.PostCreatePersonAccount) // ok

	// http://159.203.35.156:8000/person/update/account/selector
	// request params keys: id, selector, value 
	// -> selector can be: FirstName, LastName, Email, Phone -> follow exact case
	personAccnts.POST("/update/account/selector", person.PostUpdatePersonAccountBySelector) // ok

	// http://159.203.35.156:8000/person/update/account/firstname
	// request params keys: id, firstname
	personAccnts.POST("/update/account/firstname", person.PostUpdatePersonAccountFirstName) // ok

	// http://159.203.35.156:8000/person/update/account/lastname
	// request earams keys: id, lastname
	personAccnts.POST("/update/account/lastname", person.PostUpdatePersonAccountLastName) // ok

	// http://159.203.35.156:8000/person/update/account/phone
	// request params keys: id, phone
	personAccnts.POST("/update/account/phone", person.PostUpdatePersonAccountPhone) // ok

	// http://159.203.35.156:8000/person/update/account/email
	// request params keys: id, email
	personAccnts.POST("/update/account/email", person.PostUpdatePersonAccountEmail) // ok

	// http://159.203.35.156:8000/person/delete/account
	// request params keys: id
	personAccnts.POST("/delete/account", person.PostDeletePersonAccount)

	// http://159.203.35.156:8000/person/create/document/new
	// request params keys: id, documentname, holder, countryissue
	personAccnts.POST("/create/document/new", person.PostCreatePersonAccountDocument) // ok

	// http://159.203.35.156:8000/person/create/document/version
	// request params keys: id, documentname, filename, rsalink
	personAccnts.POST("/create/document/version", person.PostCreatePersonAccountDocumentVersion) // ok

	// http://159.203.35.156:8000/person/update/document/countryissue
	// request params keys: id, documentname, country
	personAccnts.POST("/update/document/countryissue", person.PostUpdatePersonAccountDocumentCountryIssue) // ok

	// http://159.203.35.156:8000/person/update/document/holder
	// request params keys: id, documentname, holder
	personAccnts.POST("/update/document/holder", person.PostUpdatePersonAccountDocumentHolderName) // ok

	// http://159.203.35.156:8000/person/delete/document
	// request params keys: id, documentname
	personAccnts.POST("/delete/document", person.PostDeletePersonAccountDocument) // ok

	// http://159.203.35.156:8000/person/delete/document/version
	// request params keys: id, documentname, version
	personAccnts.POST("/delete/document/version", person.PostDeletePersonAccountDocumentVersion) // ok

	// requests

	// http://159.203.35.156:8000/person/create/request/account
	// request params keys: requesterid, recipientid, arguments -> arguments is an array without spaces - cannot be empty
	// arguments: firstName, lastName, email, phone
	personAccnts.POST("/create/request/account", person.PostCreatePersonAccountDataRequest) // ok

	// http://159.203.35.156:8000/person/create/request/document
	// request params keys: requesterid, recipientid, documentname, copy, arguments -> arguments is an array without spaces - can be empty
	// arguments: holder, countryIssue, documentName
	personAccnts.POST("/create/request/document", person.PostCreatePersonDocumentDataRequest) // ok

	// http://159.203.35.156:8000/person/accept/request/account
	// request params keys: recipientid, requestid, arguments -> arguments is an array without spaces
	personAccnts.POST("/accept/request/account", person.PostAcceptPersonAccountDataRequest) // ok

	// http://159.203.35.156:8000/person/accept/request/document
	// request params keys: recipientid, requestid, copy, arguments -> arguments is an array without spaces
	// documentcopy can be empty string if the recipient would not provide it, otherwise - it should match the document version that is to be provided
	personAccnts.POST("/accept/request/document", person.PostAcceptPersonDocumentDataRequest) // ok

	// http://159.203.35.156:8000/person/reject/request/account
	// request params keys: recipientid, requestid
	personAccnts.POST("/reject/request/account", person.PostRejectPersonAccountDataRequest) // ok

	// http://159.203.35.156:8000/person/reject/request/document
	// request params keys: recipientid, requestid
	personAccnts.POST("/reject/request/document", person.PostRejectPersonDocumentDataRequest) // ok

	// http://159.203.35.156:8000/person/get/requests/data/selector
	// request params keys: type, selector, value
	// selectors: requesterPublicId, recipientPublicId, status
 	personAccnts.GET("/get/requests/data/selector", person.GetPersonAccountRequestsBySelector) // ok

	// http://159.203.35.156:8000/person/get/requests/publicids/selector
	// request params keys: type, selector, value
	// type: accountData, documentData, any
	// selectors: requesterPublicId, recipientPublicId, status
	personAccnts.GET("/get/requests/publicids/selector", person.GetPersonAccountRequestsPublicIdsBySelector) // ok

	// http:159.203.35.156:8000/person/get/requests/data/recipient
	// request params keys: query, type, recipientid
	// query: objects, publicIds
	// type: accountData, documentData, any
	personAccnts.GET("/get/requests/data/recipient", person.GetPersonAccountRequestsByRecipient) // ok

	// http://159.203.35.156:8000/person/get/requests/data/requester
	// request params keys: query, type, requesterId
	// query: objects, publicIds
	// type: accountData, documentData, any
	personAccnts.GET("/get/requests/requester", person.GetPersonAccountRequestsByRequester) // ok

	// http://159.203.35.156:8000/person/get/requests/data/documentname
	// request params keys: query, documentname
	// query: objects, publicIds
	personAccnts.GET("/get/requests/documentname", person.GetPersonAccountRequestsByDocumentName) // ok

	// http://159.203.35.156:8000/person/get/requests/status
	// request params keys: query, type, status
	// query: objects, publicIds
	// type: accountData, documentData, any
	personAccnts.GET("/get/requests/status", person.GetPersonAccountRequestsByStatus) // ok

	// institution: query single account data
	institutionAccnts.GET("/get/account/data", institution.GetInstitutionAccountData)
	institutionAccnts.GET("/get/account/history", institution.GetInstitutionAccountHistory)
	institutionAccnts.GET("/get/document/data", institution.GetInstitutionAccountDocumentData)
	institutionAccnts.GET("/get/document/version", institution.GetInstitutionAccountDocumentVersion)
	institutionAccnts.GET("/get/document/versions", institution.GetInstitutionAccountDocumentVersions)

	// institution: query multiple accounts data
	institutionAccnts.GET("/get/accounts/email", institution.GetInstitutionAccountsByEmail)
	institutionAccnts.GET("/get/accounts/name", institution.GetInstitutionAccountsByOrgName)
	institutionAccnts.GET("/get/accounts/contact", institution.GetInstitutionAccountsByContactPerson)
	institutionAccnts.GET("/get/accounts/selector", institution.GetInstitutionAccountsBySelector)

	// institution: invoke
	institutionAccnts.POST("/create/account", institution.PostCreateInstitutionAccount)
	institutionAccnts.POST("/update/account/selector", institution.PostUpdateInstitutionAccountBySelector)
	institutionAccnts.POST("/update/account/name", institution.PostUpdateInstitutionAccountName)
	institutionAccnts.POST("/update/account/email", institution.PostUpdateInstitutionAccountEmail)
	institutionAccnts.POST("/update/account/phone", institution.PostUpdateInstitutionAccountPhone)
	institutionAccnts.POST("/update/account/contact", institution.PostUpdateInstitutionAccountContactPerson)
	institutionAccnts.POST("/update/account/address", institution.PostUpdateInstitutionAccountAddress)
	institutionAccnts.POST("/delete/account", institution.PostDeleteInstitutionAccount)

	institutionAccnts.POST("/create/document/new", institution.PostCreateInstitutionAccountDocument)
	institutionAccnts.POST("/create/document/version", institution.PostCreateInstitutionAccountDocumentVersion)
	institutionAccnts.POST("/update/document/countryissue", institution.PostUpdateInstitutionAccountDocumentCountryIssue)
	institutionAccnts.POST("/update/document/holder", institution.PostUpdateInstitutionAccountDocumentHolderName)
	institutionAccnts.POST("/delete/document", institution.PostDeleteInstitutionAccountDocument)
	institutionAccnts.POST("/delete/document/version", institution.PostDeleteInstitutionAccountDocumentVersion)

	err := router.Run(":8000")

	if err != nil {
		logger.LogEvent("Server stopped because of error", err)
	}
}
