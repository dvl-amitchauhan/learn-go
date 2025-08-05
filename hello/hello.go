package hello

const hindi = "Hindi"
const (
	englishHelloPrefix = "Hello, "
	hindiHelloPrefix   = "नमस्ते, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
