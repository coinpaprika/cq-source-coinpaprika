package client

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
	"go.uber.org/ratelimit"
)

func NewHttpClient(logger zerolog.Logger, debug bool, rateNumber int, rateDuration time.Duration) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Timeout = 5 * time.Second
	retryClient.RetryMax = 3
	retryClient.RetryWaitMin = 1 * time.Second
	retryClient.RetryWaitMax = 5 * time.Second
	retryClient.Logger = nil

	if debug {
		retryClient.RequestLogHook = func(l retryablehttp.Logger, r *http.Request, i int) {
			logger.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("attempt", i).
				Msg("Send request")
		}
	}

	client := retryClient.StandardClient()
	client.Transport = newLimitedTransport(logger, client.Transport, rateNumber, rateDuration)
	return client
}

type transport struct {
	logger    zerolog.Logger
	wrappedRT http.RoundTripper
	limiter   ratelimit.Limiter
}

func newLimitedTransport(logger zerolog.Logger, t http.RoundTripper, rateNumber int, rateDuration time.Duration) http.RoundTripper {
	if t == nil {
		t = http.DefaultTransport
	}
	return &transport{
		logger:    logger,
		wrappedRT: t,
		limiter:   ratelimit.New(rateNumber, ratelimit.Per(rateDuration)),
	}
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	t.limiter.Take()
	return t.wrappedRT.RoundTrip(r)
}
