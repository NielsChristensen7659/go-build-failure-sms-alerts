package main

import "testing"

func TestAlertText(t *testing.T) {
	cases := []struct {
		name, status string
		want         bool
	}{{"failed", "failed", true}, {"passed", "passed", false}, {"queued", "queued", false}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, got := AlertText(BuildEvent{Project: "api", Commit: "abc", Status: tc.status})
			if got != tc.want {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
		})
	}
}
