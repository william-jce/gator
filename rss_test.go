package main

import (
	"testing"
	"time"
)

func TestParsePubDate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   time.Time
		wantOK bool
	}{
		{
			name:   "RFC1123 with +0000 offset",
			input:  "Tue, 23 Nov 2026 14:00:00 +0000",
			want:   time.Date(2026, time.November, 23, 14, 0, 0, 0, time.UTC),
			wantOK: true,
		},
		{
			name:   "empty string",
			input:  "",
			wantOK: false,
		},
		{
			name:   "single-digit day",
			input:  "Tue, 3 Nov 2026 14:00:00 +0000",
			want:   time.Date(2026, time.November, 3, 14, 0, 0, 0, time.UTC),
			wantOK: true,
		},
		{
			name:   "GMT zone name",
			input:  "Tue, 03 Nov 2026 14:00:00 GMT",
			want:   time.Date(2026, time.November, 3, 14, 0, 0, 0, time.UTC),
			wantOK: true,
		},
		{
			name:   "non-zero offset",
			input:  "Tue, 03 Nov 2026 09:00:00 -0500",
			want:   time.Date(2026, time.November, 3, 14, 0, 0, 0, time.UTC), // same instant as 09:00 -0500
			wantOK: true,
		},
		{
			name:   "two-digit year is not supported",
			input:  "Tue, 03 Nov 26 14:00:00 +0000",
			wantOK: false,
		},
		{
			name:   "surrounding whitespace",
			input:  "\n    Tue, 03 Nov 2026 14:00:00 +0000\n  ",
			want:   time.Date(2026, time.November, 3, 14, 0, 0, 0, time.UTC),
			wantOK: true,
		},
		{
			name:   "garbage input",
			input:  "not a date",
			wantOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := parsePubDate(tc.input)
			if ok != tc.wantOK {
				t.Fatalf("parsePubDate(%q) ok = %v, want %v", tc.input, ok, tc.wantOK)
			}
			if !got.Equal(tc.want) {
				t.Errorf("parsePubDate(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
