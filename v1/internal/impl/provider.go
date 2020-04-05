package impl

import I "github.com/Foxcapades/Argonaut/v1/pkg/argo"

var provider = NewProvider()

func GetProvider() I.Provider {
	return provider
}

func SetProvider(p I.Provider) {
	if p == nil {
		panic("attempted to set a nil Provider")
	}
	provider = p
}

func NewProvider() I.Provider {
	return &Provider{
		arg:  NewArgBuilder,
		com:  NewCommandBuilder,
		flag: NewFlagBuilder,
		fgp:  NewFlagGroupBuilder,
	}
}

type Provider struct {
	arg  I.ArgumentProvider
	com  I.CommandProvider
	flag I.FlagProvider
	fgp  I.FlagGroupProvider
}

func (p *Provider) NewArg() I.ArgumentBuilder {
	return p.arg()
}

func (p *Provider) NewCommand() I.CommandBuilder {
	return p.com()
}

func (p *Provider) NewFlag() I.FlagBuilder {
	return p.flag()
}

func (p *Provider) NewFlagGroup() I.FlagGroupBuilder {
	return p.fgp()
}

func (p *Provider) ArgumentProvider(provider I.ArgumentProvider) (this I.Provider) {
	if provider == nil {
		panic("attempted to set a nil ArgumentProvider")
	}
	p.arg = provider
	return p
}

func (p *Provider) CommandProvider(provider I.CommandProvider) (this I.Provider) {
	if provider == nil {
		panic("attempted to set a nil CommandProvider")
	}
	p.com = provider
	return p
}

func (p *Provider) FlagProvider(provider I.FlagProvider) (this I.Provider) {
	if provider == nil {
		panic("attempted to set a nil FlagProvider")
	}
	p.flag = provider
	return p
}

func (p *Provider) FlagGroupProvider(provider I.FlagGroupProvider) (this I.Provider) {
	if provider == nil {
		panic("attempted to set a nil FlagGroupProvider")
	}
	p.fgp = provider
	return p
}
