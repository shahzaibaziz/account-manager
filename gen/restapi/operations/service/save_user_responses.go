// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shahzaibaziz/account-manager/gen/models"
)

// SaveUserCreatedCode is the HTTP code returned for type SaveUserCreated
const SaveUserCreatedCode int = 201

/*SaveUserCreated successfully save user object into database

swagger:response saveUserCreated
*/
type SaveUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.UserResponse `json:"body,omitempty"`
}

// NewSaveUserCreated creates SaveUserCreated with default headers values
func NewSaveUserCreated() *SaveUserCreated {

	return &SaveUserCreated{}
}

// WithPayload adds the payload to the save user created response
func (o *SaveUserCreated) WithPayload(payload *models.UserResponse) *SaveUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save user created response
func (o *SaveUserCreated) SetPayload(payload *models.UserResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveUserUnauthorizedCode is the HTTP code returned for type SaveUserUnauthorized
const SaveUserUnauthorizedCode int = 401

/*SaveUserUnauthorized Unauthorized

swagger:response saveUserUnauthorized
*/
type SaveUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveUserUnauthorized creates SaveUserUnauthorized with default headers values
func NewSaveUserUnauthorized() *SaveUserUnauthorized {

	return &SaveUserUnauthorized{}
}

// WithPayload adds the payload to the save user unauthorized response
func (o *SaveUserUnauthorized) WithPayload(payload *models.Error) *SaveUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save user unauthorized response
func (o *SaveUserUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveUserConflictCode is the HTTP code returned for type SaveUserConflict
const SaveUserConflictCode int = 409

/*SaveUserConflict Conflict

swagger:response saveUserConflict
*/
type SaveUserConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveUserConflict creates SaveUserConflict with default headers values
func NewSaveUserConflict() *SaveUserConflict {

	return &SaveUserConflict{}
}

// WithPayload adds the payload to the save user conflict response
func (o *SaveUserConflict) WithPayload(payload *models.Error) *SaveUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save user conflict response
func (o *SaveUserConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SaveUserDefault Internal Server Error

swagger:response saveUserDefault
*/
type SaveUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveUserDefault creates SaveUserDefault with default headers values
func NewSaveUserDefault(code int) *SaveUserDefault {
	if code <= 0 {
		code = 500
	}

	return &SaveUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the save user default response
func (o *SaveUserDefault) WithStatusCode(code int) *SaveUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the save user default response
func (o *SaveUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the save user default response
func (o *SaveUserDefault) WithPayload(payload *models.Error) *SaveUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save user default response
func (o *SaveUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
