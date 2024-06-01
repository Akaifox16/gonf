package cmd

import "github.com/Akaifox16/gonf/config"

var cfg *config.Config

func SetConfig(config *config.Config) {
	cfg = config
}
