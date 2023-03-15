package main

import (
	"flag"
	"fmt"
	"github.com/Dreamacro/clash/config"
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

type ExtendConfig struct {
	ProxyGroup []map[string]any `yaml:"proxy-groups"`
	Rule       []string         `yaml:"rules"`
}

var (
	source string
	target string
)

func init() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	home := current.HomeDir

	flag.StringVar(&source, "source", path.Join(home, ".config/clash/config.yaml"), "源配置")
	flag.StringVar(&target, "target", path.Join(home, ".config/clash/openai.extend.yaml"), "目标配置")
	flag.Parse()
}

func getSourceConfig() *config.RawConfig {
	abs, err := filepath.Abs(source)
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		panic(err)
	}

	cfg, err := config.UnmarshalRawConfig(data)

	if err != nil {
		panic(err)
	}

	return cfg
}

func filterOpenapiProxy(cfg *config.RawConfig) (proxies []string) {

	for _, proxy := range cfg.Proxy {
		name := proxy["name"].(string)
		if strings.Contains(name, "香港") ||
			strings.Contains(name, "台湾") ||
			strings.Contains(name, "新加坡") {
			continue
		}
		proxies = append(proxies, name)
	}
	return
}

func writeExtendConfig(proxies []string) {
	extend, err := yaml.Marshal(ExtendConfig{
		ProxyGroup: []map[string]any{
			{
				"name":    "OPENAPI",
				"type":    "select",
				"proxies": proxies,
			},
		},
		Rule: []string{
			"DOMAIN-SUFFIX,openai.com,OPENAPI"},
	})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(target, extend, 0666)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("生成OPENAPI扩展配置...")
	cfg := getSourceConfig()

	proxies := filterOpenapiProxy(cfg)

	writeExtendConfig(proxies)

}
