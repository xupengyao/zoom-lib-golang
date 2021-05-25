package zoom

import "fmt"

// GetUserPath - v2 path for getting a specific user
const GetUserSettingsPath = "/users/%s/settings"

// GetUserOpts contains options for GetUser
type GetUserSettingsOpts struct {
	EmailOrID string         `url:"-"`
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using the default client
func GetUserSettings(opts GetUserSettingsOpts) (UserSettings, error) {
	return defaultClient.GetUserSettings(opts)
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetUserSettings(opts GetUserSettingsOpts) (UserSettings, error) {
	var ret = UserSettings{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetUserSettingsPath, opts.EmailOrID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
