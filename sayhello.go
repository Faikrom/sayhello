package sayhello

func SayHello(name string, umur int) string {
	return "Hello World"
}

func Random(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		return v + "Hehe"
	case int:
		return v + 5
	case bool:
		return !v
	default:
		return nil
	}
}
