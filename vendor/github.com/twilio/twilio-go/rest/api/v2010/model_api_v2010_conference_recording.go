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

// ApiV2010ConferenceRecording struct for ApiV2010ConferenceRecording
type ApiV2010ConferenceRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the recording
	ApiVersion *string `json:"api_version,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The Conference SID that identifies the conference associated with the recording
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The start time of the recording, given in RFC 2822 format
	StartTime *string `json:"start_time,omitempty"`
	// The length of the recording in seconds
	Duration *string `json:"duration,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The one-time cost of creating the recording.
	Price *string `json:"price,omitempty"`
	// The currency used in the price property.
	PriceUnit *string `json:"price_unit,omitempty"`
	Status    *string `json:"status,omitempty"`
	// The number of channels in the final recording file as an integer
	Channels *int    `json:"channels,omitempty"`
	Source   *string `json:"source,omitempty"`
	// More information about why the recording is missing, if status is `absent`.
	ErrorCode *int `json:"error_code,omitempty"`
	// How to decrypt the recording.
	EncryptionDetails *interface{} `json:"encryption_details,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}