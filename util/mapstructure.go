package util

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/cast"

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

func ToMapString(input any) map[string]string {
	out := map[string]string{}
	for k, v := range ToMap[any](input) {
		out[k] = cast.ToString(v)
	}
	return out
}
