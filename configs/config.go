package configs

// Config ...
type Config struct {
	TgBotToken  string `env:"BOT_TOKEN"`
	TgChatID    int64  `env:"TG_CHAT_ID"`
	TargetURL   string `env:"TARGET_URL"`
	TargetText  string `env:"TARGET_TEXT"`
	MinDuration int    `env:"MIN_DURATION" envDefault:"5"`
	MaxDuration int    `env:"MIN_DURATION" envDefault:"15"`
}
