package zoom

import "fmt"

// UpdateMeetingOptions are the options to create a meeting with
type UpdateMeetingOptions struct {
	MeetingID      int             `url:"-" json:"-"`
	Topic          string          `json:"topic,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
	StartTime      *Time           `json:"start_time,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Password       string          `json:"password,omitempty"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Agenda         string          `json:"agenda,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Settings       MeetingSettings `json:"settings,omitempty"`
}

// UpdateMeetingPath - v2 create a meeting for a user
const UpdateMeetingPath = "/meetings/%d"

// UpdateMeeting calls PATCH /meetings/{meetingId}
func UpdateMeeting(opts UpdateMeetingOptions) (error) {
	return defaultClient.UpdateMeeting(opts)
}

// UpdateMeeting calls PATCH /meetings/{meetingId}
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingupdate
func (c *Client) UpdateMeeting(opts UpdateMeetingOptions) (error) {
	//var ret = Meeting{}
	return c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateMeetingPath, opts.MeetingID),
		DataParameters: &opts,
		HeadResponse:  true,
	})
}
