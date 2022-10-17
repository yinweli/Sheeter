package buildoo

// Initialize 初始化
func Initialize(config *Config) *Context {
	context := &Context{
		Global: &config.Global,
	}

	for _, itor := range config.Elements {
		context.element = append(context.element, &initializeElement{
			Global:  &config.Global,
			Element: itor,
		})
	} // for

	return context
}
