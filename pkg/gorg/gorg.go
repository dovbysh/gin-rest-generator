package gorg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Param struct {
	Name    string
	Tags    string
	Comment string
}

type Pager struct {
	Exists     bool
	LimitName  string
	OffsetName string
	MaxLimit   int
}

type Gorg struct {
	NoParamsRequired bool
	NoParamsRequest  bool
	Params           map[string]Param
	Pager            Pager
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
		gorgLineRemainder := strings.TrimSpace(lineRemainder[len(gorgAttributes[0]):])
		switch strings.ToLower(gorgAttributes[0]) {
		case "param":
			return g.parseParams(gorgLineRemainder)
		case "pager":
			return g.parsePager(gorgLineRemainder)
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
	name := strings.TrimSpace(fields[0])
	p, ex := g.Params[name]
	if !ex {
		p.Name = name
	}
	if t := strings.TrimSpace(fields[1]); t == "tags" {
		p.Tags = strings.TrimSpace(s[strings.Index(s, "tags")+len("tags"):])
	}
	if t := strings.TrimSpace(fields[1]); t == "comment" {
		p.Comment = strings.TrimSpace(s[strings.Index(s, "comment")+len("comment"):])
	}

	g.Params[name] = p
	return nil
}

func (g *Gorg) parsePager(s string) error {
	fields := strings.Fields(s)
	if len(fields) < 3 {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "unable to parse Pager")
	}

	limitName := strings.TrimSpace(fields[0])
	offsetName := strings.TrimSpace(fields[1])
	maxLimitString := strings.TrimSpace(fields[2])
	if maxLimitString == "" {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "unable to parse Pager no MaxLimit")
	}
	maxLimit, err := strconv.Atoi(maxLimitString)
	if err != nil {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), fmt.Sprintf("unable to parse Pager no MaxLimit: %#v", err))
	}
	g.Pager.Exists = true
	g.Pager.MaxLimit = maxLimit
	g.Pager.LimitName = limitName
	g.Pager.OffsetName = offsetName
	return nil
}
