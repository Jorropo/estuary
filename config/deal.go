package config

type Deal struct {
	FailOnTransferFailure bool `json:",omitempty"`
	Disable               bool `json:",omitempty"`
	Verified              bool `json:",omitempty"`
}
