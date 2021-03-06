// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crioto/swagger-playground/restapi/operations/cars"
	"github.com/crioto/swagger-playground/restapi/operations/user"
)

// NewSwaggerPgAPI creates a new SwaggerPg instance
func NewSwaggerPgAPI(spec *loads.Document) *SwaggerPgAPI {
	return &SwaggerPgAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		UserDeleteUserUsernameHandler: user.DeleteUserUsernameHandlerFunc(func(params user.DeleteUserUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation UserDeleteUserUsername has not yet been implemented")
		}),
		CarsGetCarsIDHandler: cars.GetCarsIDHandlerFunc(func(params cars.GetCarsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation CarsGetCarsID has not yet been implemented")
		}),
		UserGetUserUsernameHandler: user.GetUserUsernameHandlerFunc(func(params user.GetUserUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation UserGetUserUsername has not yet been implemented")
		}),
		CarsPostCarsHandler: cars.PostCarsHandlerFunc(func(params cars.PostCarsParams) middleware.Responder {
			return middleware.NotImplemented("operation CarsPostCars has not yet been implemented")
		}),
		UserPostUserHandler: user.PostUserHandlerFunc(func(params user.PostUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserPostUser has not yet been implemented")
		}),
		UserPutUserUsernameHandler: user.PutUserUsernameHandlerFunc(func(params user.PutUserUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation UserPutUserUsername has not yet been implemented")
		}),
	}
}

/*SwaggerPgAPI This is example API documentation */
type SwaggerPgAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// UserDeleteUserUsernameHandler sets the operation handler for the delete user username operation
	UserDeleteUserUsernameHandler user.DeleteUserUsernameHandler
	// CarsGetCarsIDHandler sets the operation handler for the get cars ID operation
	CarsGetCarsIDHandler cars.GetCarsIDHandler
	// UserGetUserUsernameHandler sets the operation handler for the get user username operation
	UserGetUserUsernameHandler user.GetUserUsernameHandler
	// CarsPostCarsHandler sets the operation handler for the post cars operation
	CarsPostCarsHandler cars.PostCarsHandler
	// UserPostUserHandler sets the operation handler for the post user operation
	UserPostUserHandler user.PostUserHandler
	// UserPutUserUsernameHandler sets the operation handler for the put user username operation
	UserPutUserUsernameHandler user.PutUserUsernameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerPgAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerPgAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerPgAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerPgAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerPgAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerPgAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerPgAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerPgAPI
func (o *SwaggerPgAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.UserDeleteUserUsernameHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserUsernameHandler")
	}

	if o.CarsGetCarsIDHandler == nil {
		unregistered = append(unregistered, "cars.GetCarsIDHandler")
	}

	if o.UserGetUserUsernameHandler == nil {
		unregistered = append(unregistered, "user.GetUserUsernameHandler")
	}

	if o.CarsPostCarsHandler == nil {
		unregistered = append(unregistered, "cars.PostCarsHandler")
	}

	if o.UserPostUserHandler == nil {
		unregistered = append(unregistered, "user.PostUserHandler")
	}

	if o.UserPutUserUsernameHandler == nil {
		unregistered = append(unregistered, "user.PutUserUsernameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerPgAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerPgAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *SwaggerPgAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SwaggerPgAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SwaggerPgAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SwaggerPgAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the swagger pg API
func (o *SwaggerPgAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerPgAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/{username}"] = user.NewDeleteUserUsername(o.context, o.UserDeleteUserUsernameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cars/{id}"] = cars.NewGetCarsID(o.context, o.CarsGetCarsIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{username}"] = user.NewGetUserUsername(o.context, o.UserGetUserUsernameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cars"] = cars.NewPostCars(o.context, o.CarsPostCarsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = user.NewPostUser(o.context, o.UserPostUserHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/{username}"] = user.NewPutUserUsername(o.context, o.UserPutUserUsernameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerPgAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SwaggerPgAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerPgAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerPgAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
