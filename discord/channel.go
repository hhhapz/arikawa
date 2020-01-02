package discord

type Channel struct {
	ID   Snowflake   `json:"id,string"`
	Type ChannelType `json:"type"`

	// Fields below may not appear

	GuildID Snowflake `json:"guild_id,string,omitempty"`

	Position int    `json:"position,omitempty"`
	Name     string `json:"name,omitempty"`  // 2-100 chars
	Topic    string `json:"topic,omitempty"` // 0-1024 chars
	NSFW     bool   `json:"nsfw"`

	Icon Hash `json:"icon,omitempty"`

	// Direct Messaging fields
	DMOwnerID    Snowflake `json:"owner_id,omitempty"`
	DMRecipients []User    `json:"recipients,omitempty"`

	// AppID of the group DM creator if it's bot-created
	AppID Snowflake `json:"application_id,omitempty"`

	// ID of the category the channel is in, if any.
	CategoryID Snowflake `json:"parent_id,omitempty"`

	LastPinTime Timestamp `json:"last_pin_timestamp,omitempty"`

	// Explicit permission overrides for members and roles.
	Permissions []Overwrite `json:"permission_overwrites,omitempty"`
	// ID of the last message, may not point to a valid one.
	LastMessageID Snowflake `json:"last_message_id,omitempty"`

	// Slow mode duration. Bots and people with "manage_messages" or
	// "manage_channel" permissions are unaffected.
	UserRateLimit Seconds `json:"rate_limit_per_user,omitempty"`

	// Voice, so GuildVoice only
	VoiceBitrate   int `json:"bitrate,omitempty"`
	VoiceUserLimit int `json:"user_limit,omitempty"`
}

type ChannelType uint8

const (
	GuildText ChannelType = iota
	DirectMessage
	GuildVoice
	GroupDM
	GuildCategory
	GuildNews
	GuildStore
)

type Overwrite struct {
	ID    Snowflake     `json:"id,omitempty"`
	Type  OverwriteType `json:"type"`
	Allow uint64        `json:"allow"`
	Deny  uint64        `json:"deny"`
}

type OverwriteType string

const (
	OverwriteRole   OverwriteType = "role"
	OverwriteMember OverwriteType = "member"
)
