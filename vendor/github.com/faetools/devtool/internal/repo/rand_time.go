package repo

import (
	"crypto/sha512"
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

const (
	// MinutesInDay number of minutes in one full 24 hour period.
	minutesInDay = 1440

	// HourMinuteFormat 24 hour time format HH:mm.
	hourMinuteFormat = "15:04"
)

// randTime generates a random time of day so that
// dependabot updates are spread out throughout the day
// but are consistent for any given service.
func randTime(name string) string {
	h := sha512.New()
	_, err := h.Write([]byte(name))
	cobra.CheckErr(err)

	seed := int64(binary.BigEndian.Uint64(h.Sum(nil)))

	r := rand.New(rand.NewSource(seed))

	at := time.Time{}.Add(time.Duration(r.Intn(minutesInDay)) * time.Minute)

	return at.Format(hourMinuteFormat)
}
