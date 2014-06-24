package ptor

import (
	"regexp"
)

var (
	parts_matcher = regexp.MustCompile(`\/+\w+|\/\:\w+(\[.+?\])?|\(.+?\)`)
	part_matcher  = regexp.MustCompile(`(\:?\w+)(\[.+?\])?`)
)

type Param struct {
	Name     string
	Regexp   string
	Required bool
}

func NewParam(name, regex string, required bool) *Param {
	this := new(Param)
	this.Name = name
	this.Regexp = regex
	this.Required = required

	return this
}

func PathToRegexp(path string, sensitive, end bool) (*regexp.Regexp, []*Param) {
	var (
		pattern string
		params  []*Param
	)
	
	if sensitive == false {
		pattern += "(?i)"
	}
	pattern += "^"
	
	parts := parts_matcher.FindAllString(path, -1)

	for i := range parts {
		part := parts[i]
		if len(part) <= 0 {
			continue
		}

		if string(part[0]) == "(" {
			pattern += "(?:\\" + string(part[1])
			partParts := part_matcher.FindAllStringSubmatch(part, -1)[0]
			part = partParts[1]

			if string(part[0]) == ":" {
				regex := partParts[2]
				if regex == "" {
					regex = "[a-zA-Z0-9-_]"
				}

				pattern += "(" + regex + "+)"
				params = append(params, NewParam(part[1:], regex, false))
			} else {
				pattern += part
			}

			pattern += ")?"
		} else {
			pattern += "\\" + string(part[0])
			partParts := part_matcher.FindAllStringSubmatch(part, -1)[0]
			part = partParts[1]

			if string(part[0]) == ":" {
				regex := partParts[2]
				if regex == "" {
					regex = "[a-zA-Z0-9-_]"
				}

				pattern += "(" + regex + "+)"
				params = append(params, NewParam(part[1:], regex, true))
			} else {
				pattern += part
			}
		}
	}

	if end == true {
		pattern += "\\/?$"
	} else {
		pattern += "\\/?|$"
	}
	return regexp.MustCompile(pattern), params
}