package search

import (
	"net/http"
	"net/url"
	"time"
)

// Response represents the result as received from the search service.
type Response struct {
	TotalHits int
	Hits      []Hit
	Meta      Meta
}

// Hit is a search hit. It holds e.g. Asset or Series.
type Hit interface {
	Subset() *HitSubset
}

// HitSubset contains the subset of fields in common between different hit types (*Series and *Asset for now)
type HitSubset struct {
	ID                    string              `json:"id"`
	Type                  string              `json:"type"`
	Cinemascope           Image               `json:"cinemascope"`
	ContentSource         string              `json:"content_source"`
	Country               []string            `json:"country"`
	Credits               []Credit            `json:"credits"`
	DescriptionExtendedDa string              `json:"description_extended_da"`
	DescriptionExtendedFi string              `json:"description_extended_fi"`
	DescriptionExtendedNb string              `json:"description_extended_nb"`
	DescriptionExtendedSv string              `json:"description_extended_sv"`
	DescriptionLongDa     string              `json:"description_long_da"`
	DescriptionLongFi     string              `json:"description_long_fi"`
	DescriptionLongNb     string              `json:"description_long_nb"`
	DescriptionLongSv     string              `json:"description_long_sv"`
	DescriptionMediumDa   string              `json:"description_medium_da"`
	DescriptionMediumFi   string              `json:"description_medium_fi"`
	DescriptionMediumNb   string              `json:"description_medium_nb"`
	DescriptionMediumSv   string              `json:"description_medium_sv"`
	DescriptionShortDa    string              `json:"description_short_da"`
	DescriptionShortFi    string              `json:"description_short_fi"`
	DescriptionShortNb    string              `json:"description_short_nb"`
	DescriptionShortSv    string              `json:"description_short_sv"`
	DescriptionTinyDa     string              `json:"description_tiny_da"`
	DescriptionTinyFi     string              `json:"description_tiny_fi"`
	DescriptionTinyNb     string              `json:"description_tiny_nb"`
	DescriptionTinySv     string              `json:"description_tiny_sv"`
	Events                []Event             `json:"events"`
	ExternalReferences    []ExternalReference `json:"external_references"`
	FifteenBySeven        Image               `json:"fifteen_by_seven"`
	FourByThree           Image               `json:"four_by_three"`
	GenreDescriptionDa    string              `json:"genre_description_da"`
	GenreDescriptionFi    string              `json:"genre_description_fi"`
	GenreDescriptionNb    string              `json:"genre_description_nb"`
	GenreDescriptionSv    string              `json:"genre_description_sv"`
	Genres                []Genre             `json:"genres"`
	KeywordsDa            []Keyword           `json:"keywords_da"`
	KeywordsFi            []Keyword           `json:"keywords_fi"`
	KeywordsNb            []Keyword           `json:"keywords_nb"`
	KeywordsSv            []Keyword           `json:"keywords_sv"`
	Landscape             Image               `json:"landscape"`
	Poster                Image               `json:"poster"`
	SpokenLanguages       []string            `json:"spoken_languages"`
	Studio                string              `json:"studio"`
	Tags                  Tags                `json:"tags"`
	Timestamp             string              `json:"timestamp"`
	TitleDa               string              `json:"title_da"`
	TitleFi               string              `json:"title_fi"`
	TitleNb               string              `json:"title_nb"`
	TitleSv               string              `json:"title_sv"`
}

// Meta contains request/response meta information
type Meta struct {
	StatusCode int
	Header     http.Header
	RequestURL *url.URL
}

// Asset is an asset hit returned by the search service.
type Asset struct {
	hitSubset *HitSubset

	Arena                 string              `json:"arena"`
	AwayTeam              Team                `json:"awayteam"`
	Brand                 Brand               `json:"brand"`
	Cinemascope           Image               `json:"cinemascope"`
	ContentSource         string              `json:"content_source"`
	Country               []string            `json:"country"`
	Credits               []Credit            `json:"credits"`
	DRMRestrictions       bool                `json:"drm_restrictions"`
	DescriptionExtendedDa string              `json:"description_extended_da"`
	DescriptionExtendedFi string              `json:"description_extended_fi"`
	DescriptionExtendedNb string              `json:"description_extended_nb"`
	DescriptionExtendedSv string              `json:"description_extended_sv"`
	DescriptionLongDa     string              `json:"description_long_da"`
	DescriptionLongFi     string              `json:"description_long_fi"`
	DescriptionLongNb     string              `json:"description_long_nb"`
	DescriptionLongSv     string              `json:"description_long_sv"`
	DescriptionMediumDa   string              `json:"description_medium_da"`
	DescriptionMediumFi   string              `json:"description_medium_fi"`
	DescriptionMediumNb   string              `json:"description_medium_nb"`
	DescriptionMediumSv   string              `json:"description_medium_sv"`
	DescriptionShortDa    string              `json:"description_short_da"`
	DescriptionShortFi    string              `json:"description_short_fi"`
	DescriptionShortNb    string              `json:"description_short_nb"`
	DescriptionShortSv    string              `json:"description_short_sv"`
	DescriptionTinyDa     string              `json:"description_tiny_da"`
	DescriptionTinyFi     string              `json:"description_tiny_fi"`
	DescriptionTinyNb     string              `json:"description_tiny_nb"`
	DescriptionTinySv     string              `json:"description_tiny_sv"`
	Duration              int                 `json:"duration"`
	EpisodeNumber         int                 `json:"episode_number"`
	Events                []Event             `json:"events"`
	ExternalReferences    []ExternalReference `json:"external_references"`
	FifteenBySeven        Image               `json:"fifteen_by_seven"`
	FourByThree           Image               `json:"four_by_three"`
	GenreDescriptionDa    string              `json:"genre_description_da"`
	GenreDescriptionFi    string              `json:"genre_description_fi"`
	GenreDescriptionNb    string              `json:"genre_description_nb"`
	GenreDescriptionSv    string              `json:"genre_description_sv"`
	Genres                []Genre             `json:"genres"`
	HomeTeam              Team                `json:"hometeam"`
	ItemsPublished        bool                `json:"items_published"`
	KeywordsDa            []Keyword           `json:"keywords_da"`
	KeywordsFi            []Keyword           `json:"keywords_fi"`
	KeywordsNb            []Keyword           `json:"keywords_nb"`
	KeywordsSv            []Keyword           `json:"keywords_sv"`
	Landscape             Image               `json:"landscape"`
	League                string              `json:"league"`
	LeagueDa              string              `json:"league_da"`
	LeagueFi              string              `json:"league_fi"`
	LeagueNb              string              `json:"league_nb"`
	LeagueSv              string              `json:"league_sv"`
	Live                  bool                `json:"live"`
	LiveEventEnd          time.Time           `json:"live_event_end"`
	LogoAwayTeam          Image               `json:"logoawayteam"`
	LogoHomeTeam          Image               `json:"logohometeam"`
	MLTNIDs               []string            `json:"mlt_nids"`
	OriginalTitle         OriginalTitle       `json:"original_title"`
	ParentalRatings       []ParentalRating    `json:"parental_ratings"`
	Poster                Image               `json:"poster"`
	ProductionYear        string              `json:"production_year"`
	PublicationRights     PublicationRights   `json:"publication_rights"`
	Season                Season              `json:"season"`
	SpokenLanguages       []string            `json:"spoken_languages"`
	Studio                string              `json:"studio"`
	Tags                  Tags                `json:"tags"`
	Timestamp             string              `json:"timestamp"`
	TitleDa               string              `json:"title_da"`
	TitleFi               string              `json:"title_fi"`
	TitleNb               string              `json:"title_nb"`
	TitleSv               string              `json:"title_sv"`
	Type                  string              `json:"type"`
	VMANID                string              `json:"vman_id"`
	VideoID               string              `json:"video_id"`
}

// Subset returns the HitSubset for an *Asset
func (a *Asset) Subset() *HitSubset {
	if a.hitSubset != nil {
		return a.hitSubset
	}
	a.hitSubset = &HitSubset{
		ID:                    a.VideoID,
		Type:                  a.Type,
		Cinemascope:           a.Cinemascope,
		ContentSource:         a.ContentSource,
		Country:               a.Country,
		Credits:               a.Credits,
		DescriptionExtendedDa: a.DescriptionExtendedDa,
		DescriptionExtendedFi: a.DescriptionExtendedFi,
		DescriptionExtendedNb: a.DescriptionExtendedNb,
		DescriptionExtendedSv: a.DescriptionExtendedSv,
		DescriptionLongDa:     a.DescriptionLongDa,
		DescriptionLongFi:     a.DescriptionLongFi,
		DescriptionLongNb:     a.DescriptionLongNb,
		DescriptionLongSv:     a.DescriptionLongSv,
		DescriptionMediumDa:   a.DescriptionMediumDa,
		DescriptionMediumFi:   a.DescriptionMediumFi,
		DescriptionMediumNb:   a.DescriptionMediumNb,
		DescriptionMediumSv:   a.DescriptionMediumSv,
		DescriptionShortDa:    a.DescriptionShortDa,
		DescriptionShortFi:    a.DescriptionShortFi,
		DescriptionShortNb:    a.DescriptionShortNb,
		DescriptionShortSv:    a.DescriptionShortSv,
		DescriptionTinyDa:     a.DescriptionTinyDa,
		DescriptionTinyFi:     a.DescriptionTinyFi,
		DescriptionTinyNb:     a.DescriptionTinyNb,
		DescriptionTinySv:     a.DescriptionTinySv,
		Events:                a.Events,
		ExternalReferences:    a.ExternalReferences,
		FifteenBySeven:        a.FifteenBySeven,
		FourByThree:           a.FourByThree,
		GenreDescriptionDa:    a.GenreDescriptionDa,
		GenreDescriptionFi:    a.GenreDescriptionFi,
		GenreDescriptionNb:    a.GenreDescriptionNb,
		GenreDescriptionSv:    a.GenreDescriptionSv,
		Genres:                a.Genres,
		KeywordsDa:            a.KeywordsDa,
		KeywordsFi:            a.KeywordsFi,
		KeywordsNb:            a.KeywordsNb,
		KeywordsSv:            a.KeywordsSv,
		Landscape:             a.Landscape,
		Poster:                a.Poster,
		SpokenLanguages:       a.SpokenLanguages,
		Studio:                a.Studio,
		Tags:                  a.Tags,
		Timestamp:             a.Timestamp,
		TitleDa:               a.TitleDa,
		TitleFi:               a.TitleFi,
		TitleNb:               a.TitleNb,
		TitleSv:               a.TitleSv,
	}
	return a.hitSubset
}

// Series is an series hit returned by the search service.
type Series struct {
	hitSubset *HitSubset

	BrandID               string              `json:"brand_id"`
	Cinemascope           Image               `json:"cinemascope"`
	ContentSource         string              `json:"content_source"`
	Country               []string            `json:"country"`
	Credits               []Credit            `json:"credits"`
	DescriptionExtendedDa string              `json:"description_extended_da"`
	DescriptionExtendedFi string              `json:"description_extended_fi"`
	DescriptionExtendedNb string              `json:"description_extended_nb"`
	DescriptionExtendedSv string              `json:"description_extended_sv"`
	DescriptionLongDa     string              `json:"description_long_da"`
	DescriptionLongFi     string              `json:"description_long_fi"`
	DescriptionLongNb     string              `json:"description_long_nb"`
	DescriptionLongSv     string              `json:"description_long_sv"`
	DescriptionMediumDa   string              `json:"description_medium_da"`
	DescriptionMediumFi   string              `json:"description_medium_fi"`
	DescriptionMediumNb   string              `json:"description_medium_nb"`
	DescriptionMediumSv   string              `json:"description_medium_sv"`
	DescriptionShortDa    string              `json:"description_short_da"`
	DescriptionShortFi    string              `json:"description_short_fi"`
	DescriptionShortNb    string              `json:"description_short_nb"`
	DescriptionShortSv    string              `json:"description_short_sv"`
	DescriptionTinyDa     string              `json:"description_tiny_da"`
	DescriptionTinyFi     string              `json:"description_tiny_fi"`
	DescriptionTinyNb     string              `json:"description_tiny_nb"`
	DescriptionTinySv     string              `json:"description_tiny_sv"`
	Events                []Event             `json:"events"`
	ExternalReferences    []ExternalReference `json:"external_references"`
	FifteenBySeven        Image               `json:"fifteen_by_seven"`
	FourByThree           Image               `json:"four_by_three"`
	GenreDescriptionDa    string              `json:"genre_description_da"`
	GenreDescriptionFi    string              `json:"genre_description_fi"`
	GenreDescriptionNb    string              `json:"genre_description_nb"`
	GenreDescriptionSv    string              `json:"genre_description_sv"`
	Genres                []Genre             `json:"genres"`
	ID                    string              `json:"id"`
	KeywordsDa            []Keyword           `json:"keywords_da"`
	KeywordsFi            []Keyword           `json:"keywords_fi"`
	KeywordsNb            []Keyword           `json:"keywords_nb"`
	KeywordsSv            []Keyword           `json:"keywords_sv"`
	Landscape             Image               `json:"landscape"`
	Poster                Image               `json:"poster"`
	Seasons               []int               `json:"seasons"`
	SpokenLanguages       []string            `json:"spoken_languages"`
	Studio                string              `json:"studio"`
	Tags                  Tags                `json:"tags"`
	Timestamp             string              `json:"timestamp"`
	TitleDa               string              `json:"title_da"`
	TitleFi               string              `json:"title_fi"`
	TitleNb               string              `json:"title_nb"`
	TitleSv               string              `json:"title_sv"`
	Type                  string              `json:"type"`
}

// Subset returns the HitSubset for a *Series
func (s *Series) Subset() *HitSubset {
	if s.hitSubset != nil {
		return s.hitSubset
	}
	s.hitSubset = &HitSubset{
		ID:                    s.BrandID,
		Type:                  s.Type,
		Cinemascope:           s.Cinemascope,
		ContentSource:         s.ContentSource,
		Country:               s.Country,
		Credits:               s.Credits,
		DescriptionExtendedDa: s.DescriptionExtendedDa,
		DescriptionExtendedFi: s.DescriptionExtendedFi,
		DescriptionExtendedNb: s.DescriptionExtendedNb,
		DescriptionExtendedSv: s.DescriptionExtendedSv,
		DescriptionLongDa:     s.DescriptionLongDa,
		DescriptionLongFi:     s.DescriptionLongFi,
		DescriptionLongNb:     s.DescriptionLongNb,
		DescriptionLongSv:     s.DescriptionLongSv,
		DescriptionMediumDa:   s.DescriptionMediumDa,
		DescriptionMediumFi:   s.DescriptionMediumFi,
		DescriptionMediumNb:   s.DescriptionMediumNb,
		DescriptionMediumSv:   s.DescriptionMediumSv,
		DescriptionShortDa:    s.DescriptionShortDa,
		DescriptionShortFi:    s.DescriptionShortFi,
		DescriptionShortNb:    s.DescriptionShortNb,
		DescriptionShortSv:    s.DescriptionShortSv,
		DescriptionTinyDa:     s.DescriptionTinyDa,
		DescriptionTinyFi:     s.DescriptionTinyFi,
		DescriptionTinyNb:     s.DescriptionTinyNb,
		DescriptionTinySv:     s.DescriptionTinySv,
		Events:                s.Events,
		ExternalReferences:    s.ExternalReferences,
		FifteenBySeven:        s.FifteenBySeven,
		FourByThree:           s.FourByThree,
		GenreDescriptionDa:    s.GenreDescriptionDa,
		GenreDescriptionFi:    s.GenreDescriptionFi,
		GenreDescriptionNb:    s.GenreDescriptionNb,
		GenreDescriptionSv:    s.GenreDescriptionSv,
		Genres:                s.Genres,
		KeywordsDa:            s.KeywordsDa,
		KeywordsFi:            s.KeywordsFi,
		KeywordsNb:            s.KeywordsNb,
		KeywordsSv:            s.KeywordsSv,
		Landscape:             s.Landscape,
		Poster:                s.Poster,
		SpokenLanguages:       s.SpokenLanguages,
		Studio:                s.Studio,
		Tags:                  s.Tags,
		Timestamp:             s.Timestamp,
		TitleDa:               s.TitleDa,
		TitleFi:               s.TitleFi,
		TitleNb:               s.TitleNb,
		TitleSv:               s.TitleSv,
	}
	return s.hitSubset
}

// Brand is the brand of an asset, e.g. Idol or Harry Potter.
type Brand struct {
	Cinemascope           Image               `json:"cinemascope"`
	Country               []string            `json:"country"`
	DescriptionExtendedDa string              `json:"description_extended_da"`
	DescriptionExtendedFi string              `json:"description_extended_fi"`
	DescriptionExtendedNb string              `json:"description_extended_nb"`
	DescriptionExtendedSv string              `json:"description_extended_sv"`
	DescriptionLongDa     string              `json:"description_long_da"`
	DescriptionLongFi     string              `json:"description_long_fi"`
	DescriptionLongNb     string              `json:"description_long_nb"`
	DescriptionLongSv     string              `json:"description_long_sv"`
	DescriptionMediumDa   string              `json:"description_medium_da"`
	DescriptionMediumFi   string              `json:"description_medium_fi"`
	DescriptionMediumNb   string              `json:"description_medium_nb"`
	DescriptionMediumSv   string              `json:"description_medium_sv"`
	DescriptionShortDa    string              `json:"description_short_da"`
	DescriptionShortFi    string              `json:"description_short_fi"`
	DescriptionShortNb    string              `json:"description_short_nb"`
	DescriptionShortSv    string              `json:"description_short_sv"`
	DescriptionTinyDa     string              `json:"description_tiny_da"`
	DescriptionTinyFi     string              `json:"description_tiny_fi"`
	DescriptionTinyNb     string              `json:"description_tiny_nb"`
	DescriptionTinySv     string              `json:"description_tiny_sv"`
	ExternalReferences    []ExternalReference `json:"external_references"`
	FifteenBySeven        Image               `json:"fifteen_by_seven"`
	FourByThree           Image               `json:"four_by_three"`
	GenreDescriptionDa    string              `json:"genre_description_da"`
	GenreDescriptionFi    string              `json:"genre_description_fi"`
	GenreDescriptionNb    string              `json:"genre_description_nb"`
	GenreDescriptionSv    string              `json:"genre_description_sv"`
	Genres                []Genre             `json:"genres"`
	ID                    string              `json:"id"`
	Landscape             Image               `json:"landscape"`
	Poster                Image               `json:"poster"`
	Studio                string              `json:"studio"`
	TitleDa               string              `json:"title_da"`
	TitleFi               string              `json:"title_fi"`
	TitleNb               string              `json:"title_nb"`
	TitleSv               string              `json:"title_sv"`
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

// Keyword holds a URL-friendly ID and a human-friendly text for a keyword.
type Keyword struct {
	NID  string `json:"nid"`
	Text string `json:"text"`
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
	Cinemascope           Image               `json:"cinemascope"`
	Country               []string            `json:"country"`
	DescriptionExtendedDa string              `json:"description_extended_da"`
	DescriptionExtendedFi string              `json:"description_extended_fi"`
	DescriptionExtendedNb string              `json:"description_extended_nb"`
	DescriptionExtendedSv string              `json:"description_extended_sv"`
	DescriptionLongDa     string              `json:"description_long_da"`
	DescriptionLongFi     string              `json:"description_long_fi"`
	DescriptionLongNb     string              `json:"description_long_nb"`
	DescriptionLongSv     string              `json:"description_long_sv"`
	DescriptionMediumDa   string              `json:"description_medium_da"`
	DescriptionMediumFi   string              `json:"description_medium_fi"`
	DescriptionMediumNb   string              `json:"description_medium_nb"`
	DescriptionMediumSv   string              `json:"description_medium_sv"`
	DescriptionShortDa    string              `json:"description_short_da"`
	DescriptionShortFi    string              `json:"description_short_fi"`
	DescriptionShortNb    string              `json:"description_short_nb"`
	DescriptionShortSv    string              `json:"description_short_sv"`
	DescriptionTinyDa     string              `json:"description_tiny_da"`
	DescriptionTinyFi     string              `json:"description_tiny_fi"`
	DescriptionTinyNb     string              `json:"description_tiny_nb"`
	DescriptionTinySv     string              `json:"description_tiny_sv"`
	ExternalReferences    []ExternalReference `json:"external_references"`
	FifteenBySeven        Image               `json:"fifteen_by_seven"`
	FourByThree           Image               `json:"four_by_three"`
	GenreDescriptionDa    string              `json:"genre_description_da"`
	GenreDescriptionFi    string              `json:"genre_description_fi"`
	GenreDescriptionNb    string              `json:"genre_description_nb"`
	GenreDescriptionSv    string              `json:"genre_description_sv"`
	Genres                []Genre             `json:"genres"`
	ID                    string              `json:"id"`
	Landscape             Image               `json:"landscape"`
	Number                int                 `json:"season_number"`
	NumberOfEpisodes      int                 `json:"number_of_episodes"`
	Poster                Image               `json:"poster"`
	Studio                string              `json:"studio"`
	TitleDa               string              `json:"title_da"`
	TitleFi               string              `json:"title_fi"`
	TitleNb               string              `json:"title_nb"`
	TitleSv               string              `json:"title_sv"`
}

// Tags bind otherwise unrelated assets.
type Tags map[string][]string

// Team represents e.g. home team for a sports asset.
type Team struct {
	Name string `json:"name"`
	NID  string `json:"nid"`
}
