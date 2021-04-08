package gorg

import (
	"strings"

	"github.com/pkg/errors"
)

type Param struct {
	Name string
	Tags string
}

type Gorg struct {
	NoParamsRequired bool
	NoParamsRequest  bool
	Params           map[string]Param
}

var NoGorgTags = errors.New("no @gorg tags")
var NoGorgTagsParsed = errors.New("no @gorg tags parsed")

func (g *Gorg) ParseComment(s string) error {
	commentLine := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(s), "//"))
	if len(commentLine) == 0 {
		return errors.Wrap(NoGorgTags, s)
	}
	attribute := strings.Fields(commentLine)[0]
	lineRemainder := strings.TrimSpace(commentLine[len(attribute):])
	lowerAttribute := strings.ToLower(attribute)

	if lowerAttribute == "@gorg" {
		lowerLineReminder := strings.ToLower(lineRemainder)
		switch lowerLineReminder {
		case "noparamsrequired":
			g.NoParamsRequired = true
			break
		case "noparamsrequest":
			g.NoParamsRequest = true
			break
		}
		gorgAttributes := strings.Fields(lineRemainder)
		if len(gorgAttributes) == 0 {
			return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "no gorg Attributes")
		}
		gorgLineRemainder := strings.TrimSpace(commentLine[len(gorgAttributes[0]):])
		switch strings.ToLower(gorgAttributes[0]) {
		case "param":
			return g.parseParams(gorgLineRemainder)
		}
	}

	return errors.Wrap(NoGorgTagsParsed, s)
}

func (g *Gorg) parseParams(s string) error {
	if g.Params == nil {
		g.Params = make(map[string]Param)
	}

	fields := strings.Fields(s)
	if len(fields) < 3 {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "unable to parse Params")
	}
	name := strings.TrimSpace(fields[1])
	p, ex := g.Params[name]
	if !ex {
		p.Name = name
	}
	if t := strings.TrimSpace(fields[2]); t == "tags" {
		p.Tags = strings.TrimSpace(s[strings.Index(s, "tags")+len("tags"):])
	}

	g.Params[name] = p
	return nil
}
