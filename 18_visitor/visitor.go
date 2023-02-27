package _8_visitor

import (
	"fmt"
	"path"
)

/*
访问者模式：表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
访问者模式（Visitor）是一种操作一组对象的操作，它的目的是不改变对象的定义，但允许新增不同的访问者，来定义新的操作。
*/

// Visitor 访问者
type Visitor interface {
	Visit(IResourceFile) error
}

// IResourceFile
type IResourceFile interface {
	Accept(visitor Visitor) error
}

func NewResourceFile(filePath string) (IResourceFile, error) {
	switch path.Ext(filePath) {
	case ".ppt":
		return &PPTFile{path: filePath}, nil
	case ".pdf":
		return &PdfFile{path: filePath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filePath)
	}
}

type PdfFile struct {
	path string
}

func (f *PdfFile) Accept(visitor Visitor) error {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}

func (p *PPTFile) Accept(visitor Visitor) error {
	return visitor.Visit(p)
}

// Compressor 实现压缩功能
type Compressor struct {
}

func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.VisitPPTFile(f)
	case *PdfFile:
		return c.VisitPDFFile(f)
	default:
		return fmt.Errorf("not found resource typr: %#v", r)
	}
}

// VisitPPTFile VisitPPTFile
func (c *Compressor) VisitPPTFile(f *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}

// VisitPDFFile VisitPDFFile
func (c *Compressor) VisitPDFFile(f *PdfFile) error {
	fmt.Println("this is pdf file")
	return nil
}
