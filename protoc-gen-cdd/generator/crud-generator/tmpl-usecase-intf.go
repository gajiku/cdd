package crudgenerator

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplUseCaseIntf = template.Must(template.New("usecase-intf").Funcs(template.FuncMap{
		"GetPrimaryKeyAsString": getPrimaryKeyAsString,
	}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package {{.PackageName}}
	
	import (
		"{{.PackagePath}}/entity"
	)
	
	type Repository interface {
		additional_repository
		GetByPrimaryKey({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) (entity.{{.GetName}}, error)
		GetAll() ([]entity.{{.GetName}}, error)
		Create(in entity.{{.GetName}}) (entity.{{.GetName}}, error)
		Update(in entity.{{.GetName}}) (entity.{{.GetName}}, error)
		Delete({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) error
	}

	type UseCase interface {
		additional_usecase
		GetByPrimaryKey({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) (entity.{{.GetName}}, error)
		GetAll() ([]entity.{{.GetName}}, error)
		Create(in entity.{{.GetName}}) (entity.{{.GetName}}, error)
		Update(in entity.{{.GetName}}) (entity.{{.GetName}}, error)
		Delete({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) error
	}
	
	// Please write code below in interface.go
	/*
	package {{.PackageName}}

	type additional_repository interface{
		// AdditionalFunc1()
		// AdditionalFunc2()
	}

	type additional_usecase interface {
		// AdditionalFunc1()
		// AdditionalFunc2()
	}
	*/

	`))
)

func applyTemplateUseCaseIntf(mext *descriptor.MessageDescriptorExt, pkgpath string) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	packageName := strcase.ToKebab(strings.ToLower("crud-" + mext.DBSchema.TableName))
	var tmplData = struct {
		*descriptor.MessageDescriptorExt
		PackageName string
		PackagePath string
	}{
		mext,
		strings.Replace(packageName, "-", "_", -1),
		pkgpath,
	}

	if err := tmplUseCaseIntf.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "usecase/" + packageName + "/interface.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}
