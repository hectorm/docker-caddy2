package tomladapter

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/pelletier/go-toml"

	"caddy/utils/envreplacer"
)

func init() {
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

// Adapter adapts TOML to Caddy JSON.
type Adapter struct{}

// Adapt converts the TOML config in body to Caddy JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	bodyReplaced, err := envreplacer.Replace(body)
	if err != nil {
		return nil, nil, err
	}
	tree, err := toml.LoadBytes(bodyReplaced)
	if err != nil {
		return nil, nil, err
	}
	json, err := json.Marshal(tree.ToMap())
	if err != nil {
		return nil, nil, err
	}
	return json, nil, nil
}

// Interface guard.
var _ caddyconfig.Adapter = (*Adapter)(nil)
