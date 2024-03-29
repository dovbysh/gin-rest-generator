package gorg

import (
	"encoding/json"
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
	Exists       bool
	UseOffset    bool
	LimitName    string
	OffsetName   string
	MaxLimit     int
	DefaultLimit int
}

type SuccessCallback struct {
	Package string
	Cb      string
}

type Gorg struct {
	NoParamsRequired bool
	NoParamsRequest  bool
	Params           map[string]Param
	Pager            Pager
	Vars             map[string]string
	SuccessCallbacks struct {
		Key string
		Cbs map[string]SuccessCallback
	}
	RequestType         string
	RequestAfterParseCb string
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
		case "var":
			return g.parseVar(gorgLineRemainder)
		case "success_cb":
			return g.parseSuccessCallbacks(gorgLineRemainder)
		case "request_type":
			return g.parseRequestType(gorgLineRemainder)
		case "request_after_parse_cb":
			return g.parseRequestAfterParseCb(gorgLineRemainder)
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
	var useOffset bool
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
	if len(fields) >= 4 {
		for i := 3; i < len(fields); i++ {
			v := strings.ToLower(strings.TrimSpace(fields[i]))
			switch v {
			case "useoffset":
				useOffset = true
			}
			z, err := strconv.Atoi(v)
			if err == nil {
				g.Pager.DefaultLimit = z
			}
		}
	}
	g.Pager.Exists = true
	g.Pager.MaxLimit = maxLimit
	g.Pager.LimitName = limitName
	g.Pager.OffsetName = offsetName
	g.Pager.UseOffset = useOffset
	return nil
}

func (g *Gorg) parseVar(s string) error {
	fields := strings.Fields(s)
	if len(fields) < 2 {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "unable to parse Var")
	}

	name := strings.TrimSpace(fields[0])
	value := strings.TrimSpace(fields[1])
	g.Vars[name] = value
	return nil
}

func (g *Gorg) parseSuccessCallbacks(s string) error {
	if g.SuccessCallbacks.Cbs != nil {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), fmt.Sprintf("success_cb already defined, Key: %s", g.SuccessCallbacks.Key))
	}

	g.SuccessCallbacks.Cbs = make(map[string]SuccessCallback)
	fields := strings.SplitN(s, " ", 2)
	if len(fields) < 2 {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), "unable to parse success_cb")
	}
	g.SuccessCallbacks.Key = strings.TrimSpace(fields[0])
	err := json.Unmarshal([]byte(fields[1]), &g.SuccessCallbacks.Cbs)
	if err != nil {
		return errors.Wrap(errors.Wrap(NoGorgTagsParsed, s), fmt.Sprintf("unable to parse success_cb json error: %s", err))
	}
	return nil
}

func (g *Gorg) parseRequestType(s string) error {
	g.RequestType = s
	return nil
}

func (g *Gorg) parseRequestAfterParseCb(s string) error {
	g.RequestAfterParseCb = s
	return nil
}
