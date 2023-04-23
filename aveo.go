package aveo

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type (
	lookuper struct {
		env Env
	}
)

func Process(env Env, ctx context.Context, prefix string, spec interface{}) (err error) {
	lookuper := envconfig.PrefixLookuper(prefix, NewLookuper(env))
	return envconfig.ProcessWith(ctx, spec, lookuper)
}

func NewLookuper(env Env) envconfig.Lookuper {
	return &lookuper{
		env: env,
	}
}

func (l *lookuper) Lookup(key string) (val string, exists bool) {
	return l.env.LookupEnv(key)
}
