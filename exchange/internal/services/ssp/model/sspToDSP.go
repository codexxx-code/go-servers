package model

// SSPToDSP - Связка между SSP и DSP
type SSPToDSP struct {
	SSPSlug string `db:"ssp_slug"`
	DSPSlug string `db:"dsp_slug"`
}
