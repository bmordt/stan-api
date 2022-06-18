package models

type StanRequest struct {
	Payload      []StanPayload `json:"payload"`
	Skip         int           `json:"skip"`
	Take         int           `json:"take"`
	TotalRecords int           `json:"totalRecords"`
}

type StanPayload struct {
	Country       string        `json:"country,omitempty"`
	Description   string        `json:"description,omitempty"`
	Drm           bool          `json:"drm,omitempty"`
	EpisodeCount  int           `json:"episodeCount,omitempty"`
	Genre         string        `json:"genre,omitempty"`
	Image         StanImage     `json:"image,omitempty"`
	Language      string        `json:"language,omitempty"`
	NextEpisode   interface{}   `json:"nextEpisode,omitempty"`
	PrimaryColour string        `json:"primaryColour,omitempty"`
	Seasons       []StanSeasons `json:"seasons,omitempty"`
	Slug          string        `json:"slug"`
	Title         string        `json:"title"`
	TvChannel     string        `json:"tvChannel"`
}

type StanImage struct {
	ShowImage string `json:"showImage"`
}

type StanNextEpisode struct {
	Channel     interface{} `json:"channel"`
	ChannelLogo string      `json:"channelLogo"`
	Date        interface{} `json:"date"`
	HTML        string      `json:"html"`
	URL         string      `json:"url"`
}

type StanSeasons struct {
	Slug string `json:"slug"`
}
