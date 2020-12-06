package app

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Level             string `env:"LOG_LEVEL" envDefault:"info"`
	BaseUrl           string `env:"TRANSLATOR_URL" envDefault:"https://www.ldoceonline.com/dictionary/"`
	CSVWordInputFile  string `env:"WORD_CSV_INPUT_FILE" envDefault:"words.csv"`
	CSVWordOutputFile string `env:"WORD_CSV_OUT_PUTFILE" envDefault:"words_translated.csv"`
}

func NewConfig() (cfg *Config, err error) {
	cfg = new(Config)
	return cfg, env.Parse(cfg)
}
