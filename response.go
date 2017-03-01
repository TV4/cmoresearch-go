package search

import (
	"net/http"
	"time"
)

// Response represents the result as received from the search service.
type Response struct {
	TotalHits int    `json:"total_hits"`
	Hits      []*Hit `json:"assets"`
	Meta      Meta   `json:"-"`
}

// Meta contains request/response meta information
type Meta struct {
	StatusCode int
	Header     http.Header
}

// Hit is an item returned by the search service.
type Hit struct {
	HitCommon
	Arena             string            `json:"arena"`
	AwayTeam          Team              `json:"awayteam"`
	Brand             Brand             `json:"brand"`
	ContentSource     string            `json:"content_source"`
	Credits           []Credit          `json:"credits"`
	DRMRestrictions   bool              `json:"drm_restrictions"`
	Duration          int               `json:"duration"`
	EpisodeNumber     int               `json:"episode_number"`
	Events            []Event           `json:"events"`
	HomeTeam          Team              `json:"hometeam"`
	ItemsPublished    bool              `json:"items_published"`
	KeywordsDa        []string          `json:"keywords_dk"`
	KeywordsFi        []string          `json:"keywords_fi"`
	KeywordsNb        []string          `json:"keywords_nb"`
	KeywordsSv        []string          `json:"keywords_sv"`
	Live              bool              `json:"live"`
	LiveEventEnd      time.Time         `json:"live_event_end"`
	LogoAwayTeam      Image             `json:"logoawayteam"`
	LogoHomeTeam      Image             `json:"logohometeam"`
	MLTNIDs           []string          `json:"mlt_nids"`
	MLTTags           string            `json:"mlt_tags"`
	OriginalTitle     OriginalTitle     `json:"original_title"`
	ParentalRatings   []ParentalRating  `json:"parental_ratings"`
	ProductionYear    string            `json:"production_year"`
	PublicationRights PublicationRights `json:"publication_rights"`
	Season            Season            `json:"season"`
	SpokenLanguages   []string          `json:"spoken_languages"`
	Tags              Tags              `json:"tags"`
	Timestamp         string            `json:"timestamp"`
	Type              string            `json:"type"`
	VMANID            string            `json:"vman_id"`
	VideoID           string            `json:"video_id"`
}

// HitCommon the common archetype of all hit types.
type HitCommon struct {
	Cinemascope        Image               `json:"cinemascope"`
	Country            []string            `json:"country"`
	ExternalReferences []ExternalReference `json:"external_references"`
	Genres             []Genre             `json:"genres"`
	Landscape          Image               `json:"landscape"`
	Poster             Image               `json:"poster"`
	Studio             string              `json:"studio"`

	TitleDa string `json:"title_da"`
	TitleFi string `json:"title_fi"`
	TitleNb string `json:"title_nb"`
	TitleSv string `json:"title_sv"`

	DescriptionExtendedDa string `json:"description_extended_da"`
	DescriptionExtendedFi string `json:"description_extended_fi"`
	DescriptionExtendedNb string `json:"description_extended_nb"`
	DescriptionExtendedSv string `json:"description_extended_sv"`

	DescriptionLongDa string `json:"description_long_da"`
	DescriptionLongFi string `json:"description_long_fi"`
	DescriptionLongNb string `json:"description_long_nb"`
	DescriptionLongSv string `json:"description_long_sv"`

	DescriptionMediumDa string `json:"description_medium_da"`
	DescriptionMediumFi string `json:"description_medium_fi"`
	DescriptionMediumNb string `json:"description_medium_nb"`
	DescriptionMediumSv string `json:"description_medium_sv"`

	DescriptionShortDa string `json:"description_short_da"`
	DescriptionShortFi string `json:"description_short_fi"`
	DescriptionShortNb string `json:"description_short_nb"`
	DescriptionShortSv string `json:"description_short_sv"`

	DescriptionTinyDa string `json:"description_tiny_da"`
	DescriptionTinyFi string `json:"description_tiny_fi"`
	DescriptionTinyNb string `json:"description_tiny_nb"`
	DescriptionTinySv string `json:"description_tiny_sv"`
}

// Brand is the brand of an asset, e.g. Idol or Harry Potter.
type Brand struct {
	HitCommon
	ID string `json:"id"`
}

// Credit represents one entry in the credit list for an asset.
type Credit struct {
	Function string `json:"function"`
	NID      string `json:"nid"`
	Name     string `json:"name"`
	Rolename string `json:"rolename"`
}

// Event contains publication rights for an asset.
type Event struct {
	Site        string    `json:"site"`
	DeviceTypes []string  `json:"device_types"`
	Products    []string  `json:"products"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	PublishTime time.Time `json:"publish_time"`
}

// ExternalReference is a reference to additional information contained in
// a different system.
type ExternalReference struct {
	Locator string `json:"locator"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

// Genre is the main and sub genre inforamtion for an asset, e.g. Main:
// Horror, Sub [Action, Drama, Romance]
type Genre struct {
	Main string   `json:"main"`
	Sub  []string `json:"sub"`
}

// Image is the image attribute for an asset. It may contain localizations.
type Image struct {
	Caption       string           `json:"caption"`
	Copyright     string           `json:"copyright"`
	Localizations []LocalizedImage `json:"localizations"`
	URL           string           `json:"url"`
}

// LocalizedImage is a localized image.
type LocalizedImage struct {
	Caption   string `json:"caption"`
	Copyright string `json:"copyright"`
	Language  string `json:"language"`
	URL       string `json:"url"`
}

// LocationRestrictions contains restrictions by location.
type LocationRestrictions struct {
	IncludeCountries []string `json:"include_countries"`
}

// LocationRights contains rights based on location.
type LocationRights struct {
	LocationRestrictions LocationRestrictions `json:"location_restrictions"`
	Product              string               `json:"product"`
}

// OriginalTitle is the title of an asset in the original language.
type OriginalTitle struct {
	Language string `json:"language"`
	Text     string `json:"text"`
	Type     string `json:"type"`
}

// ParentalRating is a parental rating of an asset for a given country and
// rating system.
type ParentalRating struct {
	Country string `json:"country"`
	System  string `json:"system"`
	Value   string `json:"value"`
}

// PublicationRights contain location rights for an asset.
type PublicationRights struct {
	LocationRights LocationRights `json:"location_rights"`
}

// Season is a season of an asset, e.g. "Idol season 2".
type Season struct {
	HitCommon
	ID               string `json:"id"`
	Number           int    `json:"season_number"`
	NumberOfEpisodes int    `json:"number_of_episodes"`
}

// Tags bind otherwise unrelated assets.
type Tags map[string][]string

// Team represents e.g. home team for a sports asset.
type Team struct {
	Name string `json:"name"`
	NID  string `json:"nid"`
}
