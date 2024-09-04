package util

import (
	"github.com/go-viper/mapstructure/v2"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func ToMap[T any](input any) map[string]T {
	out := map[string]T{}
	cfg := &mapstructure.DecoderConfig{Squash: true, Result: &out, TagName: ""}
	decoder, err := mapstructure.NewDecoder(cfg)
	osutil.PanicErr(err)
	err = decoder.Decode(input)
	osutil.PanicErr(err)
	return out
}
