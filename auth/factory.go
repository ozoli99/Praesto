package auth

import "errors"

func NewAuthAdapterFactory(provider string, config interface{}) (AuthAdapter, error) {
	switch provider {
		case "auth0":
			auth0Config, ok := config.(Auth0Config)
			if !ok {
				return nil, errors.New("invalid configuration for Auth0 adapter")
			}
			return NewAuth0Adapter(auth0Config)
		case "clerk":
			clerkConfig, ok := config.(ClerkConfig)
			if !ok {
				return nil, errors.New("invalid configuration for Clerk adapter")
			}
			return NewClerkAdapter(clerkConfig)
		default:
			return nil, errors.New("unknown auth provider")
	}
}