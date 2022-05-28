// Code AutoGenerated; DO NOT EDIT.

package types

// Message Represents a message.
type Message struct {
	Animation *Animation `json:"animation,omitempty"`
	Audio *Audio `json:"audio,omitempty"`
	AuthorSignature string `json:"author_signature,omitempty"`
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	Chat Chat `json:"chat"`
	ConnectedWebsite string `json:"connected_website,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Date int64 `json:"date"`
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	Dice *Dice `json:"dice,omitempty"`
	Document *Document `json:"document,omitempty"`
	EditDate int64 `json:"edit_date,omitempty"`
	Entities []MessageEntity `json:"entities,omitempty"`
	ForwardDate int64 `json:"forward_date,omitempty"`
	ForwardFrom *User `json:"forward_from,omitempty"`
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID int64 `json:"forward_from_message_id,omitempty"`
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	ForwardSignature string `json:"forward_signature,omitempty"`
	From *User `json:"from,omitempty"`
	Game *Game `json:"game,omitempty"`
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	Invoice *Invoice `json:"invoice,omitempty"`
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	Location *Location `json:"location,omitempty"`
	MediaGroupID string `json:"media_group_id,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MessageID int64 `json:"message_id"`
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
	NewChatTitle string `json:"new_chat_title,omitempty"`
	PassportData *PassportData `json:"passport_data,omitempty"`
	Photo []PhotoSize `json:"photo,omitempty"`
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	Poll *Poll `json:"poll,omitempty"`
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	SenderChat *Chat `json:"sender_chat,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	Text string `json:"text,omitempty"`
	Venue *Venue `json:"venue,omitempty"`
	ViaBot *User `json:"via_bot,omitempty"`
	Video *Video `json:"video,omitempty"`
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
	VideoNote *VideoNote `json:"video_note,omitempty"`
	Voice *Voice `json:"voice,omitempty"`
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
}
