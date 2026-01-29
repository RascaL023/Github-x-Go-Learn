package cava

import "try1/register"

type CavaProcessor struct {}

func (p CavaProcessor) Name() string { return "cava"; }

func (p CavaProcessor) Resolve(name any) (any, error) {
	return name, nil;
}

func init() { register.Register(CavaProcessor{}); }

