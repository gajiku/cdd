package scaffold_mysql

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplUseCaseImpl = template.Must(template.New("usecase-impl").Funcs(template.FuncMap{
		// "GetPrimaryKeyAsString": getPrimaryKeyAsString,
	}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package {{.GetCrudPackageName}}
	
	import (
		"{{.GoModuleName}}/entity"
		_ "github.com/jinzhu/gorm/dialects/mysql"
	)
	
	type usecase struct {
		repo Repository
	}
	
	func NewUsecase(repo Repository) UseCase {
		return &usecase{
			repo: repo,
		}
	}

	func (uc *usecase) GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true true }}) (*entity.{{.GetName}}, error) {
		return uc.repo.GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true false }})
	}

	func (uc *usecase) GetAll() ([]*entity.{{.GetName}}, error) {
		return uc.repo.GetAll()
	}

	func (uc *usecase) Create(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		return uc.repo.Create(in)
	}
	
	func (uc *usecase) Update(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		return uc.repo.Update(in)
	}
	
	func (uc *usecase) Delete({{ .GetPrimaryKeyAsString "" "" "," true true }}) error {
		return uc.repo.Delete({{ .GetPrimaryKeyAsString "" "" "," true false }})
	}
	`))
)

func applyTemplateUseCaseImpl(sm ScaffoldMysql) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		ScaffoldMysql
	}{
		sm,
	}

	if err := tmplUseCaseImpl.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "usecase/" + sm.GetCrudPackageName() + "/usecase-impl.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}
