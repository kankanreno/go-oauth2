package models

// Client client model
type Client struct {
	ID     string
	Secret string
	Domain string
	Scopes string
	Public bool
}

// GetID client id
func (c *Client) GetID() string {
	return c.ID
}

// GetSecret client secret
func (c *Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *Client) GetDomain() string {
	return c.Domain
}

// GetScopes client scopes
func (c *Client) GetScopes() string {
	return c.Scopes
}

// IsPublic public
func (c *Client) IsPublic() bool {
	return c.Public
}
