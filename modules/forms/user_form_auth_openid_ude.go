// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package forms

import (
	"net/http"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/web/middleware"
	"gitea.com/go-chi/binding"
)

// SignInOpenIDUDEForm form for signing in with OpenID
type SignInOpenIDUDEForm struct {
	Email    string `binding:"Required;MaxSize(256)"`
	Remember bool
}

// Validate validates the fields
func (f *SignInOpenIDUDEForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// SignUpOpenIDForm form for signin up with OpenID
type SignUpOpenIDUDEForm struct {
	UserName           string `binding:"Required;AlphaDashDot;MaxSize(40)"`
	Email              string `binding:"Required;Email;MaxSize(254)"`
	GRecaptchaResponse string `form:"g-recaptcha-response"`
	HcaptchaResponse   string `form:"h-captcha-response"`
}

// Validate validates the fields
func (f *SignUpOpenIDUDEForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// ConnectOpenIDForm form for connecting an existing account to an OpenID URI
type ConnectOpenIDUDEForm struct {
	UserName string `binding:"Required;MaxSize(254)"`
	Password string `binding:"Required;MaxSize(255)"`
}

// Validate validates the fields
func (f *ConnectOpenIDUDEForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}
