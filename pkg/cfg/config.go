package cfg

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/goccy/go-yaml"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"

	"github.com/starudream/sign-task/util"
)

//go:embed config_keys
var configKeysRaw []byte

var configKeys = map[string]int{}

func init() {
	sc := bufio.NewScanner(bytes.NewReader(configKeysRaw))
	for i := 1; sc.Scan(); i++ {
		configKeys[sc.Text()] = i
	}
}

func Sort(r string, m map[string]any) (ms yaml.MapSlice) {
	for k, v := range m {
		if m1, ok1 := v.(map[string]any); ok1 {
			ms = append(ms, item(r, k, Sort(key(r, k), m1)))
		} else if s1, ok2 := v.([]any); ok2 {
			s2 := make([]any, len(s1))
			for i := range s1 {
				if m2, ok3 := s1[i].(map[string]any); ok3 {
					s2[i] = Sort(key(r, k), m2)
				} else {
					s2[i] = s1[i]
				}
			}
			ms = append(ms, item(r, k, s2))
		} else {
			ms = append(ms, item(r, k, v))
		}
	}
	slices.SortFunc(ms, func(a, b yaml.MapItem) int {
		i, j := configKeys[key(r, a.Key)], configKeys[key(r, b.Key)]
		if i == 0 {
			return 9999
		} else if j == 0 {
			return -9999
		}
		return i - j
	})
	return
}

func key(r, k any) string {
	return r.(string) + "." + k.(string)
}

func item(r, k string, v any) yaml.MapItem {
	if t := key(r, k); osutil.ArgTest() && configKeys[t] == 0 {
		fmt.Println(t)
	}
	return yaml.MapItem{Key: k, Value: v}
}

func Save() error {
	raw := Sort("$", config.Raw())

	bs, err := yaml.MarshalWithOptions(raw, options...)
	if err != nil {
		return fmt.Errorf("marshal config error: %w", err)
	}

	filename := config.LoadedFile()
	if filename == "" {
		filename = filepath.Join(osutil.ExeDir(), osutil.ExeName()+".yaml")
		slog.Info("no config file specified, use default: %s", filename)
	} else {
		diff := compare(filename, bs)
		if diff == "" {
			slog.Debug("config file not changed")
		} else if diff != "" {
			slog.Debug("config file changed, diff: %s", diff)
		}
	}

	err = os.WriteFile(filename, bs, 0644)
	if err != nil {
		return fmt.Errorf("write config file error: %w", err)
	}

	slog.Info("save config success", slog.String("file", filename))

	return nil
}

func compare(srcPath string, dstData []byte) string {
	srcData, err := os.ReadFile(srcPath)
	if err != nil {
		return ""
	}

	var src, dst any
	err1 := yaml.Unmarshal(srcData, &src)
	err2 := yaml.Unmarshal(dstData, &dst)
	if err1 != nil || err2 != nil {
		return ""
	}

	return util.Diff(src, dst)
}
