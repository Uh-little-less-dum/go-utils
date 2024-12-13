package file_handlers_package_json

import (
	"encoding/json"
	"os"

	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	schemas_package_json "github.com/Uh-little-less-dum/go-utils/pkg/schemastructs/packageJson"
)

type PackageJsonHandler struct {
	File           file.File
	hasInitialized bool
	path           string
	data           schemas_package_json.PackageJsonSchema
}

func (p *PackageJsonHandler) GetData() (*schemas_package_json.PackageJsonSchema, error) {
	if p.hasInitialized {
		return &p.data, nil
	}
	b, err := os.ReadFile(p.path)
	if err != nil {
		return &p.data, err
	}
	err = json.Unmarshal(b, &p.data)
	if err != nil {
		return &p.data, err
	}
	p.hasInitialized = true
	return &p.data, nil
}

func NewPackageJsonHandler(filePath string) PackageJsonHandler {
	return PackageJsonHandler{
		path: filePath,
		File: file.NewFileItem(filePath),
	}
}
