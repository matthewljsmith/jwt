package jwt_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/matthewljsmith/jwt/v3"
	"github.com/matthewljsmith/jwt/v3/internal"
)

func TestTimeMarshalJSON(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		tt   jwt.Time
		want int64
	}{
		{jwt.Time{}, 0},
		{jwt.Time{now}, now.Unix()},
		{jwt.Time{now.Add(24 * time.Hour)}, now.Add(24 * time.Hour).Unix()},
		{jwt.Time{now.Add(24 * 30 * 12 * time.Hour)}, now.Add(24 * 30 * 12 * time.Hour).Unix()},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			b, err := tc.tt.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			var n int64
			if err = json.Unmarshal(b, &n); err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, n; got != want {
				t.Errorf("jwt.Time.Marshal mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		n     int64
		want  jwt.Time
		isNil bool
	}{
		{now.Unix(), jwt.Time{now}, false},
		{internal.Epoch.Unix() - 0xDEAD, jwt.Time{internal.Epoch}, false},
		{internal.Epoch.Unix(), jwt.Time{internal.Epoch}, false},
		{internal.Epoch.Unix() + 0xDEAD, jwt.Time{internal.Epoch.Add(0xDEAD * time.Second)}, false},
		{0, jwt.Time{}, true},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			var n *int64
			if !tc.isNil {
				n = &tc.n
			}
			b, err := json.Marshal(n)
			if err != nil {
				t.Fatal(err)
			}
			var tt jwt.Time
			if err = tt.UnmarshalJSON(b); err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want.Unix(), tt.Unix(); got != want {
				t.Errorf("jwt.Time.Unmarshal mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
func TestTimeStringUnmarshalJSON(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		n     string
		want  jwt.Time
		isNil bool
	}{
		{fmt.Sprintf("%d", now.Unix()), jwt.Time{now}, false},
		{fmt.Sprintf("%d", internal.Epoch.Unix()-0xDEAD), jwt.Time{internal.Epoch}, false},
		{fmt.Sprintf("%d", internal.Epoch.Unix()), jwt.Time{internal.Epoch}, false},
		{fmt.Sprintf("%d", internal.Epoch.Unix()+0xDEAD), jwt.Time{internal.Epoch.Add(0xDEAD * time.Second)}, false},
		{"", jwt.Time{}, true},
		{"0", jwt.Time{}, true},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			var n *string
			if !tc.isNil {
				n = &tc.n
			}
			b, err := json.Marshal(n)
			if err != nil {
				t.Fatal(err)
			}
			var tt jwt.Time
			if err = tt.UnmarshalJSON(b); err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want.Unix(), tt.Unix(); got != want {
				t.Errorf("jwt.Time.Unmarshal mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
