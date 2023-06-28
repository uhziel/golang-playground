package main

import "strings"

type Folder struct {
	Name     string
	Children []Inode
}

func (f *Folder) String() string {
	var builder strings.Builder
	builder.WriteString(f.Name)
	builder.Write([]byte("/\n"))
	for _, child := range f.Children {
		builder.WriteByte('\t')

		//childContent := child.String()
		//builder.WriteString(strings.Join(strings.Split(childContent, "\n"), "\n\t"))
		builder.WriteString(strings.ReplaceAll(child.String(), "\n", "\n\t"))

		builder.WriteByte('\n')
	}
	return builder.String()
}

func (f *Folder) Clone() Inode {
	clone := &Folder{
		Name: f.Name + "_copy",
	}

	for _, child := range f.Children {
		clone.Children = append(clone.Children, child.Clone())
	}

	return clone
}
