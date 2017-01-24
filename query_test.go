package search

import "testing"

func TestRawQuery(t *testing.T) {
	for n, tt := range []struct {
		query       *Query
		rawURLQuery string
	}{
		{&Query{}, ""},
		{&Query{VideoIDs: []string{"12345"}}, "video_ids=12345"},
		{
			&Query{
				DeviceType: "device-type",
				Language:   "language-identifier",
				Site:       "site-identifier",

				Episode:  "episode-number",
				PageSize: "page-size",
				Season:   "season-number",
				SeasonID: "season-id",
				Type:     "type-identifier",
				VideoIDs: []string{"video-id-1", "video-id-2"},

				SortBy: "episode_number",
				Order:  "asc",
			},
			"device_type=device-type" +
				"&episode=episode-number" +
				"&lang=language-identifier" +
				"&order=asc" +
				"&page_size=page-size" +
				"&season=season-number" +
				"&season_id=season-id" +
				"&site=site-identifier" +
				"&sort_by=episode_number" +
				"&type=type-identifier" +
				"&video_ids=video-id-1%2Cvideo-id-2",
		},
	} {
		if got, want := tt.query.rawURLQuery(), tt.rawURLQuery; got != want {
			t.Errorf("[%d] got %q, want %q", n, got, want)
		}
	}
}
