package settings

type Settings struct {
	Host string
	Port int
}

var DefaultSettings = Settings{
	Host: "127.0.0.1",
	Port: 8080,
}
