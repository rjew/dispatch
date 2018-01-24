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

// AuthHandlerFunc turns a function with the right signature into a auth handler
type AuthHandlerFunc func(AuthParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthHandlerFunc) Handle(params AuthParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AuthHandler interface for that can handle valid auth params
type AuthHandler interface {
	Handle(AuthParams, interface{}) middleware.Responder
}

// NewAuth creates a new http.Handler for the auth operation
func NewAuth(ctx *middleware.Context, handler AuthHandler) *Auth {
	return &Auth{Context: ctx, Handler: handler}
}

/*Auth swagger:route GET /v1/iam/auth auth

handles authentication

*/
type Auth struct {
	Context *middleware.Context
	Handler AuthHandler
}

func (o *Auth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAuthParams()

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
