///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// Oauth2HandlerFunc turns a function with the right signature into a oauth2 handler
type Oauth2HandlerFunc func(Oauth2Params, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn Oauth2HandlerFunc) Handle(params Oauth2Params, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// Oauth2Handler interface for that can handle valid oauth2 params
type Oauth2Handler interface {
	Handle(Oauth2Params, interface{}) middleware.Responder
}

// NewOauth2 creates a new http.Handler for the oauth2 operation
func NewOauth2(ctx *middleware.Context, handler Oauth2Handler) *Oauth2 {
	return &Oauth2{Context: ctx, Handler: handler}
}

/*Oauth2 swagger:route GET /v1/iam/auth oauth2

handles oauth2 based authentication

*/
type Oauth2 struct {
	Context *middleware.Context
	Handler Oauth2Handler
}

func (o *Oauth2) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewOauth2Params()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
