/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010Queue struct for ApiV2010Queue
type ApiV2010Queue struct {
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The number of calls currently in the queue.
	CurrentSize *int `json:"current_size,omitempty"`
	// A string that you assigned to describe this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Average wait time of members in the queue
	AverageWaitTime *int `json:"average_wait_time,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The max number of calls allowed in the queue
	MaxSize *int `json:"max_size,omitempty"`
}
