// Code AutoGenerated; DO NOT EDIT.

package types

// WebhookInfo Contains information about the current status of a webhook.
type WebhookInfo struct {
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
	HasCustomCertificate bool `json:"has_custom_certificate"`
	IPAddress string `json:"ip_address,omitempty"`
	LastErrorDate int64 `json:"last_error_date,omitempty"`
	LastErrorMessage string `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date,omitempty"`
	MaxConnections int `json:"max_connections,omitempty"`
	PendingUpdateCount int64 `json:"pending_update_count"`
	URL string `json:"url"`
}
