package go_configurator

type ConfigProviderData map[string]interface{}

type ConfigProvider interface {
	Encode() (ConfigProviderData, error)
	Decode(ConfigProviderData) error
}
