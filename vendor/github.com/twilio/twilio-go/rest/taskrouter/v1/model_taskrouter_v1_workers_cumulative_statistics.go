/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// TaskrouterV1WorkersCumulativeStatistics struct for TaskrouterV1WorkersCumulativeStatistics
type TaskrouterV1WorkersCumulativeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The beginning of the interval during which these statistics were calculated
	StartTime *time.Time `json:"start_time,omitempty"`
	// The end of the interval during which these statistics were calculated
	EndTime *time.Time `json:"end_time,omitempty"`
	// The minimum, average, maximum, and total time that Workers spent in each Activity
	ActivityDurations *[]interface{} `json:"activity_durations,omitempty"`
	// The total number of Reservations that were created
	ReservationsCreated *int `json:"reservations_created,omitempty"`
	// The total number of Reservations that were accepted
	ReservationsAccepted *int `json:"reservations_accepted,omitempty"`
	// The total number of Reservations that were rejected
	ReservationsRejected *int `json:"reservations_rejected,omitempty"`
	// The total number of Reservations that were timed out
	ReservationsTimedOut *int `json:"reservations_timed_out,omitempty"`
	// The total number of Reservations that were canceled
	ReservationsCanceled *int `json:"reservations_canceled,omitempty"`
	// The total number of Reservations that were rescinded
	ReservationsRescinded *int `json:"reservations_rescinded,omitempty"`
	// The SID of the Workspace that contains the Workers
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The absolute URL of the Workers statistics resource
	Url *string `json:"url,omitempty"`
}
