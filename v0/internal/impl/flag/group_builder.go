package flag

import (
	"github.com/Foxcapades/Argonaut/v0/internal/impl/trait"
	A "github.com/Foxcapades/Argonaut/v0/pkg/argo"
)

func NewFlagGroupBuilder(A.Provider) A.FlagGroupBuilder {
	return new(GBuilder)
}

type iFgb = A.FlagGroupBuilder

type GBuilder struct {
	ParentNode  A.Command
	NameTxt     trait.Named
	DescTxt     trait.Described
	FlagNodes   []A.FlagBuilder
	WarningVals []string
}

//
// Getters
//

func (f *GBuilder) GetName() string           { return f.NameTxt.NameTxt }
func (f *GBuilder) GetDescription() string    { return f.DescTxt.DescTxt }
func (f *GBuilder) GetFlags() []A.FlagBuilder { return f.FlagNodes }

//
// Setters
//

func (f *GBuilder) Parent(com A.Command) iFgb    { f.ParentNode = com; return f }
func (f *GBuilder) Name(name string) iFgb        { f.NameTxt.NameTxt = name; return f }
func (f *GBuilder) Description(desc string) iFgb { f.DescTxt.DescTxt = desc; return f }

//
// Operations
//

func (f *GBuilder) Flag(flag A.FlagBuilder) iFgb {
	if flag == nil {
		f.WarningVals = append(f.WarningVals, "FlagGroupBuilder: nil value passed to Flag()")
	} else {
		f.FlagNodes = append(f.FlagNodes, flag)
	}
	return f
}

func (f *GBuilder) Build() (out A.FlagGroup, err error) {
	flags := make([]A.Flag, len(f.FlagNodes))

	out = &Group{ParentElement: f.ParentNode, Described: f.DescTxt, Named: f.NameTxt, FlagElements: flags}

	for i, fb := range f.FlagNodes {
		fb.Parent(out)
		if flag, err := fb.Build(); err != nil {
			return nil, err
		} else {
			flags[i] = flag
		}
	}

	return
}

func (f *GBuilder) MustBuild() A.FlagGroup {
	if out, err := f.Build(); err != nil {
		panic(err)
	} else {
		return out
	}
}
