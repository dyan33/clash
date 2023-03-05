package config

func ParseConfig(main []byte, extends ...[]byte) (*Config, error) {
	rawCfg, err := UnmarshalRawConfig(main)
	if err != nil {
		return nil, err
	}
	for _, data := range extends {

		extend, err := UnmarshalRawConfig(data)
		if err != nil {
			return nil, err
		}
		mergeRawConfig(rawCfg, extend)
	}
	return ParseRawConfig(rawCfg)
}

func mergeRawConfig(main *RawConfig, extend *RawConfig) {

	//proxy
	if extend.Proxy != nil {
		main.Proxy = append(main.Proxy, extend.Proxy...)
	}

	if extend.ProxyGroup != nil {
		main.ProxyGroup = append(main.ProxyGroup, extend.ProxyGroup...)
	}

	if extend.Rule != nil {
		main.Rule = append(extend.Rule, main.Rule...)
	}

}
