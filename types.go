package main

// Config items in config.json
type Config struct {
	WebhookURL   string
	Port         int
	RichMessages bool
}

// ElasticWebhook is the main webhook payload from Elastic
// See sample payload at the bottom of this file
type ElasticWebhook struct {
	AlertID         string         `json:"alertId"`
	AlertName       string         `json:"alertName"`
	SpaceID         string         `json:"spaceId"`
	Tags            string         `json:"tags"`
	AlertInstanceID string         `json:"alertInstanceId"`
	Context         ElasticContext `json:"context"`
	State           ElasticState   `json:"state"`
}

// ElasticContext is a subfield in the main webhook payload
type ElasticContext struct {
	Message             string `json:"message"`
	DownMonitorsWithGeo string `json:"downMonitorsWithGeo"`
}

// ElasticState is a subfield in the main webhook payload
type ElasticState struct {
	FirstCheckedAt        string `json:"firstCheckedAt"`
	FirstTriggeredAt      string `json:"firstTriggeredAt"`
	CurrentTriggerStarted string `json:"currentTriggerStarted"`
	IsTriggered           string `json:"isTriggered"`
	LastCheckedAt         string `json:"lastCheckedAt"`
	LastResolvedAt        string `json:"lastResolvedAt"`
	LastTriggeredAt       string `json:"lastTriggeredAt"`
}

// DiscordWebhook is a partial struct for https://discord.com/developers/docs/resources/webhook#execute-webhook
type DiscordWebhook struct {
	Content string         `json:"content"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

// DiscordEmbed is a partial struct for https://discord.com/developers/docs/resources/channel#embed-object
type DiscordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	URL         string              `json:"url"`
	Timestamp   string              `json:"timestamp"` // needs to be ISO8601
	Color       int                 `json:"color"`     // 0xRRGGBB
	Fields      []DiscordEmbedField `json:"fields"`
}

// DiscordEmbedField is a struct for https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

/*
Sample Elastic webhook payload:

{
  "alertId": "e5328eb2-24cc-41c3-8dfb-748dea19457b",
  "alertName": "Challenge Uptime",
  "spaceId": "default",
  "tags": "",
  "alertInstanceId": "xpack.uptime.alerts.actionGroups.monitorStatus",
  "context": {
    "message": "Down monitor: test-chal-healthcheck",
    "downMonitorsWithGeo": "test-chal-healthcheck from us-west1; "
  },
  "state": {
    "firstCheckedAt": "2020-09-23T04:31:57.031Z",
    "firstTriggeredAt": "2020-09-23T04:32:59.771Z",
    "currentTriggerStarted": "2020-09-23T08:10:37.358Z",
    "isTriggered": "true",
    "lastCheckedAt": "2020-09-23T08:10:37.358Z",
    "lastResolvedAt": "2020-09-23T08:07:58.299Z",
    "lastTriggeredAt": "2020-09-23T08:10:37.358Z"
  }
}
*/
