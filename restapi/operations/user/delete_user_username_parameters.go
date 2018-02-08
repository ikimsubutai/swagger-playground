// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserUsernameParams creates a new DeleteUserUsernameParams object
// with the default values initialized.
func NewDeleteUserUsernameParams() DeleteUserUsernameParams {
	var ()
	return DeleteUserUsernameParams{}
}

// DeleteUserUsernameParams contains all the bound params for the delete user username operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteUserUsername
type DeleteUserUsernameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The name that needs to be deleted
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteUserUsernameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteUserUsernameParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Username = raw

	return nil
}
