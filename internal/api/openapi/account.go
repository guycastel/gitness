// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package openapi

import (
	"net/http"

	"github.com/bradrydzewski/my-app/internal/api/render"
	"github.com/bradrydzewski/my-app/types"

	"github.com/swaggest/openapi-go/openapi3"
)

// request to login to an account.
type loginRequest struct {
	Username string `formData:"username"`
	Password string `formData:"password"`
}

// request to register an account.
type registerRequest struct {
	Username string `formData:"username"`
	Password string `formData:"password"`
}

// helper function that constructs the openapi specification
// for the account registration and login endpoints.
func buildAccount(reflector *openapi3.Reflector) {

	onLogin := openapi3.Operation{}
	onLogin.WithTags("account")
	onLogin.WithMapOfAnything(map[string]interface{}{"operationId": "onLogin"})
	reflector.SetRequest(&onLogin, new(loginRequest), http.MethodPost)
	reflector.SetJSONResponse(&onLogin, new(types.Token), http.StatusOK)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPost, "/login", onLogin)

	onRegister := openapi3.Operation{}
	onRegister.WithTags("account")
	onRegister.WithMapOfAnything(map[string]interface{}{"operationId": "onRegister"})
	reflector.SetRequest(&onRegister, new(registerRequest), http.MethodPost)
	reflector.SetJSONResponse(&onRegister, new(types.Token), http.StatusOK)
	reflector.SetJSONResponse(&onRegister, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&onRegister, new(render.Error), http.StatusBadRequest)
	reflector.Spec.AddOperation(http.MethodPost, "/register", onRegister)
}
