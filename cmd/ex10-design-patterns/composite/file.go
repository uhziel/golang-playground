package main

type File struct {
	Name string
}

func (f *File) String() string {
	return f.Name
}

func (f *File) Clone() Inode {
	return &File{
		Name: f.Name + "_copy",
	}
}
