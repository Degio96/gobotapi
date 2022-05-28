package types

// LoginUrl Represents a parameter of the inline keyboard button used to automatically authorize a user
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram
// All the user needs to do is tap/click a button and confirm that they want to log in:
// Telegram apps support these buttons as of version 5.7.
type LoginUrl struct {
	BotUsername string `json:"bot_username,omitempty"`
	ForwardText string `json:"forward_text,omitempty"`
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
	Url string `json:"url"`
}
