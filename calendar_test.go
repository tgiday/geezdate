package geezdate

import (
	"testing"
)

var tests = []struct {
	in   string
	want string
}{
	/*
		{"2024-01-09", "2016-04-30 00:00:00 +0000 UTC"},
		{"1991-05-24", "1983-9-16 00:00:00 +0000 UTC"},
		{"2023-09-10", "2015-13-05 00:00:00 +0000 UTC"},
		{"2023-09-11", "pag 2015-13-06"},
		{"2023-09-12", "2016-01-01 00:00:00 +0000 UTC"},
		{"2023-09-19", "2016-01-08 00:00:00 +0000 UTC"},
		{"2023-12-01", "2016-03-21 00:00:00 +0000 UTC"},
		{"2023-12-31", "2016-04-21 00:00:00 +0000 UTC"},
		{"2024-01-01", "2016-04-22 00:00:00 +0000 UTC"},
		{"2024-02-29", "2016-06-21 00:00:00 +0000 UTC"},
		{"2024-03-01", "2016-06-22 00:00:00 +0000 UTC"},
		{"2024-04-01", "2016-07-23 00:00:00 +0000 UTC"},
		{"2024-09-10", "2016-13-05 00:00:00 +0000 UTC"},
		{"2024-09-11", "2017-01-01 00:00:00 +0000 UTC"},
	*/
	{"2023-09-10", "5-13-2015"},
	{"2023-09-11", "6-13-2015"}, //
	{"2023-09-12", "1-1-2016"},
	{"2024-09-10", "5-13-2016"},
	{"2024-09-11", "1-1-2017"},
	{"2024-09-12", "2-1-2017"},
	{"2024-12-31", "22-4-2017"},
	{"2025-01-01", "23-4-2017"},
	{"2025-09-10", "5-13-2017"},
	{"2025-09-11", "1-1-2018"},
	{"2025-09-12", "2-1-2018"},
	{"2026-01-01", "23-4-2018"},
	{"2026-09-10", "5-13-2018"},
	{"2026-09-11", "1-1-2019"},
	{"2026-09-12", "2-1-2019"},
	{"2027-09-10", "5-13-2019"},
	{"2027-09-11", "6-13-2019"}, //
	{"2027-09-12", "1-1-2020"},
	{"2028-09-10", "5-13-2020"},
	{"2028-09-11", "1-1-2021"},
	{"2028-09-12", "2-1-2021"},
}

func TestConvert(t *testing.T) {
	for _, v := range tests {
		got := Convert(v.in)
		if got != v.want {
			t.Errorf("got %s want %s", got, v.want)
		}
	}

}
