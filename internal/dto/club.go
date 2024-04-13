package dto

type Club struct {
	Id              int64    `json:"id"`
	ResourceState   string   `json:"resouce_state"`
	Name            string   `json:"name"`
	ProfileMedium   string   `json:"profile_medium"`
	Profile         string   `json:"string"`
	CoverPhoto      string   `json:"cover_photo"`
	CoverPhotoSmall string   `json:"cover_photo_small"`
	SportType       string   `json:"sport_type"`
	ActivityTypes   []string `json:"activity_types"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Private         bool     `json:"private"`
	MemberCount     int      `json:"member_count"`
	Featured        bool     `json:"featured"`
	Verified        bool     `json:"verified"`
	Url             string   `json:"url"`
	Membership      string   `json:"membership"`
	Admin           bool     `json:"admin"`
	Owner           bool     `json:"owner"`
	Description     string   `json:"description"`
	ClubType        string   `json:"club_type"`
	PostCount       int      `json:"post_count"`
	OwnerId         int64    `json:"owner_id"`
	FollowingCount  int      `json:"following_count"`
}
