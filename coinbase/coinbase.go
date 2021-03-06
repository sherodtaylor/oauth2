// Copyright 2021 The oauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package coinbase provides constants for using OAuth2 to access Coinbase.
package coinbase

import (
	"golang.org/x/oauth2"
)

// Endpoint is Coinbase's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.coinbase.com/oauth/authorize",
	TokenURL: "http://www.coinbase.com/oauth/token",
}
