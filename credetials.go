package adyen

// APICredentials basic API settings
//
// Description:
//
//     - Env - Environment for next API calls
//	   - APIKey - API Key issued from Adyen
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type APICredentials struct {
	Env  Environment
	HMAC string

	APIKey   string
	Username string
	Password string
}

// makeCredentials create new APICredentials
func makeCredentials(env Environment, apiKey string) APICredentials {
	return APICredentials{
		Env:    env,
		APIKey: apiKey,
	}
}
