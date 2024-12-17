package file_handlers_package_json

import (
	"fmt"
	"os"

	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PackageJsonHandler struct {
	File           file.File
	hasInitialized bool
	path           string
	data           []byte
}

func (p *PackageJsonHandler) Bytes() []byte {
	if p.hasInitialized {
		return p.data
	}
	b, err := os.ReadFile(p.path)
	if err != nil {
		log.Fatalf("Failed to parse package.json file at %s", p.File.Path())
	}
	p.data = b
	return b
}

func (p *PackageJsonHandler) Json() gjson.Result {
	b := p.Bytes()
	j := gjson.ParseBytes(b)
	p.hasInitialized = true
	return j
}

func (p *PackageJsonHandler) ApplyDependencies(deps []types.Dependency) {
	b := p.Bytes()
	for _, dep := range deps {
		d, err := sjson.SetBytes(b, fmt.Sprintf("%s.%s", dep.Type(), dep.Name()), dep.Version())
		p.data = d
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (p *PackageJsonHandler) Write() {
	if !p.hasInitialized {
		return
	}
	err := os.WriteFile(p.path, p.data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func NewPackageJsonHandler(filePath string) PackageJsonHandler {
	return PackageJsonHandler{
		path: filePath,
		File: file.NewFileItem(filePath),
	}
}
