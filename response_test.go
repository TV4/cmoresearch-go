package search

import "testing"

func TestSubset(t *testing.T) {
	var (
		a = &Asset{
			VideoID: "video-id-123",
			Type:    "movie",
		}

		s = &Series{
			BrandID: "brand-id-345",
			Type:    "series",
		}
	)

	t.Run("Asset", func(t *testing.T) {
		sub := Hit(a).Subset()

		if got, want := sub.ID, a.VideoID; got != want {
			t.Fatalf("sub.ID = %q, want %q", got, want)
		}

		if got, want := sub.Type, a.Type; got != want {
			t.Fatalf("sub.Type = %q, want %q", got, want)
		}
	})

	t.Run("Series", func(t *testing.T) {
		sub := Hit(s).Subset()

		if got, want := sub.ID, s.BrandID; got != want {
			t.Fatalf("sub.ID = %q, want %q", got, want)
		}

		if got, want := sub.Type, s.Type; got != want {
			t.Fatalf("sub.Type = %q, want %q", got, want)
		}
	})
}
