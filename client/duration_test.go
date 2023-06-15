package client_test

import (
	"testing"
	"time"

	"github.com/coinpaprika/cq-source-coinpaprika/client"
	"github.com/stretchr/testify/assert"
)

func TestWithCustomDuration(t *testing.T) {
	tt := map[string]time.Duration{
		"1d":  24 * time.Hour,
		"-2d": -2 * 24 * time.Hour,
		"0d":  0,
		"1h": time.Hour,
		"1m": time.Minute,
	}

	for in, out := range tt {
		t.Run(in, func(t *testing.T) {
			praser := client.WithCustomDurations(time.ParseDuration)
			ret, err := praser(in)
			assert.NoError(t, err)
			assert.Equal(t, out, ret)
		})
	}

}
