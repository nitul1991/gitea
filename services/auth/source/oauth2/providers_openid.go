// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/openidConnect"
)

// OpenIDProvider is a GothProvider for OpenID
type OpenIDProvider struct {
}

// Name provides the technical name for this provider
func (o *OpenIDProvider) Name() string {
	return "openidconnect"
}

// DisplayName returns the friendly name for this provider
func (o *OpenIDProvider) DisplayName() string {
	return "OpenID Connect"
}

// Image returns an image path for this provider
func (o *OpenIDProvider) Image() string {
	return "/assets/img/auth/openid_connect.svg"
}

// CreateGothProvider creates a GothProvider from this Provider
func (o *OpenIDProvider) CreateGothProvider(providerName, callbackURL string, source *Source) (goth.Provider, error) {
	provider, err := openidConnect.New(source.ClientID, source.ClientSecret, callbackURL, source.OpenIDConnectAutoDiscoveryURL, setting.OAuth2Client.OpenIDConnectScopes...)
	if err != nil {
		log.Warn("Failed to create OpenID Connect Provider with name '%s' with url '%s': %v", providerName, source.OpenIDConnectAutoDiscoveryURL, err)
	}
	return provider, err
}

// CustomURLSettings returns the custom url settings for this provider
func (o *OpenIDProvider) CustomURLSettings() *CustomURLSettings {
	return nil
}

var _ (GothProvider) = &OpenIDProvider{}

func init() {
	RegisterGothProvider(&OpenIDProvider{})
}
