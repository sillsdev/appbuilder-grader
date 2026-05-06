package loc

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed locales/*.json
var localesFS embed.FS

var currentLang string = "en"
var translations map[string]map[string]string

// Init loads the specified language file. If it doesn't exist, it falls back to en.
func Init(lang string) error {
	currentLang = lang
	data, err := localesFS.ReadFile(fmt.Sprintf("locales/%s.json", lang))
	if err != nil {
		if lang == "en" {
			return err
		}
		// fallback to en
		currentLang = "en"
		data, err = localesFS.ReadFile("locales/en.json")
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(data, &translations)
}

// T translates a key like "categories.multimedia_name" into the translated string.
// If the key is not found, or there are no translations loaded, it returns the key.
func T(key string, args ...any) string {
	res := key
	if translations != nil {
		parts := strings.SplitN(key, ".", 2)
		if len(parts) == 2 {
			if section, ok := translations[parts[0]]; ok {
				if val, ok := section[parts[1]]; ok {
					res = val
				}
			}
		}
	}

	if len(args) > 0 {
		// Handle both variadic arguments and a single slice of arguments (which comes from the template passing a slice)
		if len(args) == 1 {
			if slice, ok := args[0].([]any); ok {
				return fmt.Sprintf(res, slice...)
			}
		}
		return fmt.Sprintf(res, args...)
	}

	return res
}
