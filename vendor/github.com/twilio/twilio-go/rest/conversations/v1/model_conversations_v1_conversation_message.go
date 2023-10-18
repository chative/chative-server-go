/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ConversationsV1ConversationMessage struct for ConversationsV1ConversationMessage
type ConversationsV1ConversationMessage struct {
	// The unique ID of the Account responsible for this message.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the Conversation for this message.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The index of the message within the Conversation.
	Index *int `json:"index,omitempty"`
	// The channel specific identifier of the message's author.
	Author *string `json:"author,omitempty"`
	// The content of the message.
	Body *string `json:"body,omitempty"`
	// An array of objects that describe the Message's media if attached, otherwise, null.
	Media *[]interface{} `json:"media,omitempty"`
	// A string metadata field you can use to store any data you wish.
	Attributes *string `json:"attributes,omitempty"`
	// The unique ID of messages's author participant.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An absolute API URL for this message.
	Url *string `json:"url,omitempty"`
	// An object that contains the summary of delivery statuses for the message to non-chat participants.
	Delivery *interface{} `json:"delivery,omitempty"`
	// Absolute URL to access the receipts of this message.
	Links *map[string]interface{} `json:"links,omitempty"`
}