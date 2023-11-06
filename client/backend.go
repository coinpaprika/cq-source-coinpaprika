package client

import "github.com/cloudquery/plugin-sdk/v4/state"

type Backend interface {
	state.Client
}
