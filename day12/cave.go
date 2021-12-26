package main

import (
	"strings"
)

type Cave struct {
	options []*Cave
	name    string
	big     bool
	start   bool
	end     bool
}

type Path []*Cave

func (c *Cave) Explore(path Path, paths []Path) []Path {
	if c.end {
		path = append(path, c)
		return append(paths, path)
	}
	if !c.big && path.Contains(c.name) {
		return paths
	}
	path = append(path, c)
	for _, option := range c.options {
		newPath := cloneSlice(path)
		paths = option.Explore(newPath, paths)
	}
	return paths
}

func (c Cave) Clone() Cave {
	return Cave{
		options: c.options,
		big:     c.big,
		start:   c.start,
		end:     c.end,
		name:    c.name,
	}
}

func cloneSlice(caves []*Cave) []*Cave {
	clone := make([]*Cave, len(caves))
	for i := 0; i < len(caves); i++ {
		clonedI := caves[i].Clone()
		clone[i] = &clonedI
	}
	return clone
}

func (p Path) String() string {
	builder := strings.Builder{}
	for _, point := range p {
		builder.WriteString(point.name)
		builder.WriteString(",")
	}
	return builder.String()
}

func (p Path) Contains(name string) bool {
	for _, point := range p {
		if point.name == name {
			return true
		}
	}
	return false
}
