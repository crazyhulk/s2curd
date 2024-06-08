// Code generated by go-bindata.
// sources:
// ../template/orm_model.tpl
// DO NOT EDIT!

package gencode

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _TemplateOrm_modelTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcf\x6f\x9b\x4e\x10\xc5\xcf\xec\x5f\xf1\xc4\x09\xa2\xef\x17\x5f\xaa\x1e\x22\xa1\x2a\x0e\x6e\x95\x43\xdc\x28\xc6\xea\x31\xac\xd9\xc1\x1e\x95\x1f\xee\xb2\xb4\xa2\x88\xff\xbd\x5a\x4c\x52\x88\x2b\x59\xe6\xc2\x2c\x3b\xf3\x3e\x4f\xf3\x30\xed\x91\xd0\x75\xc1\xc6\xe8\x26\x35\x6b\x59\x50\xdf\xa3\x1e\x0e\xe8\x84\x13\x49\x23\x77\xb2\x26\x00\x37\xfb\x4a\x17\x41\xb4\x04\x12\x5b\xdd\xba\xff\xbb\xa8\x7f\xe4\xf6\x9d\x60\xb1\xc0\x81\x15\xc1\x1c\xb8\x46\xc6\x94\x2b\xe1\x3c\x28\xbc\x3e\x5c\x9a\x8f\x1f\x86\x6a\x9c\x3d\x6a\x2e\xa4\x6e\x5f\xbe\x53\xeb\x26\xc2\xd9\x1e\x95\x34\x14\x73\x41\x30\x5c\x50\x30\x54\xc9\xa0\xae\x28\x93\x4d\x6e\x6e\xef\xb7\xcf\xcf\xab\x75\xfc\x12\x3f\x3c\xae\x36\xf1\xdd\xe3\x13\xbe\xae\xb1\x7d\x8a\xee\xe2\x15\xce\xee\xac\xe6\xbd\xa6\xab\x35\xdd\x44\xf4\x42\x64\x4d\x99\xc2\xab\x29\xcf\x70\xf3\x6e\x35\x3e\x62\xb9\xcb\xc9\x1e\x3c\xdf\xee\x89\xcb\xbd\xdd\xd3\x62\x81\xb6\x6a\x50\xc8\x16\xe9\x41\x96\x7b\x02\x1b\xe1\x68\x32\x8d\x2e\xe1\x76\x5d\xf0\x36\xd6\xf7\xee\x45\x46\xb4\xf4\xfc\xbf\xfb\xee\x84\xc3\x19\x6c\x6b\xf0\x16\x47\x18\xa2\xe4\xdc\x5e\x4d\xd1\xa4\xd8\x80\x0d\x4c\x85\x63\xc5\xe5\x50\xa4\x4d\x6d\xaa\x82\x7f\x93\x82\xda\x81\xcb\xda\xc8\x32\x25\xe1\xbc\x9a\x53\x3b\xe1\xf4\xa0\xbc\xa6\x41\x6d\xfc\x3a\xa3\x09\xa7\xbf\x68\xf9\x0b\x19\x8f\xd5\x29\x68\x1f\xa4\x75\xa5\xad\xde\x4c\x6e\xe9\xf9\xc1\xb7\x03\x69\xf2\x5c\x56\x08\xf1\xc9\xfd\x0f\xac\xfc\xe0\x33\xeb\xda\x0c\xc2\x7e\xb0\xb2\x93\x17\x69\xa7\x6c\xbd\x09\x68\x20\x4c\x22\x0f\x4f\xa1\xaf\xab\x5f\x9e\x3f\xde\x4e\x7e\xb2\x10\xef\xfa\xff\xe1\x74\x64\x5c\x63\xeb\x44\x38\xb3\x35\x03\x4f\x6d\x9d\x31\x37\xf2\xe7\x75\xc4\x88\x72\x9a\x13\xcf\x34\xc7\x96\xb9\xea\x9f\x00\x00\x00\xff\xff\x96\x13\xd1\x69\xf7\x03\x00\x00")

func TemplateOrm_modelTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplateOrm_modelTpl,
		"../template/orm_model.tpl",
	)
}

func TemplateOrm_modelTpl() (*asset, error) {
	bytes, err := TemplateOrm_modelTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../template/orm_model.tpl", size: 1015, mode: os.FileMode(420), modTime: time.Unix(1524487516, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../template/orm_model.tpl": TemplateOrm_modelTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"template": &bintree{nil, map[string]*bintree{
			"orm_model.tpl": &bintree{TemplateOrm_modelTpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
