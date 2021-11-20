package parser

import (
	"github.com/schetininl/availability_parser/internal/pkg/bot"
)

// Target ...
type Target struct {
	URL  string
	Text string
}

// Parser ...
type Parser struct {
	Bot    bot.Bot
	Target Target
}

type parserOption func(*Parser)

// WithTarget ...
func WithTarget(TargetURL, TargetText string) parserOption {
	return func(p *Parser) {
		p.Target = Target{URL: TargetURL, Text: TargetText}
	}
}

// WithBot ...
func WithBot(bot bot.Bot) parserOption {
	return func(p *Parser) {
		p.Bot = bot
	}
}

// New ...
func New(opts ...parserOption) *Parser {
	p := &Parser{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
