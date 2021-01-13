// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package bindata generated by go-bindata.// sources:
// assets/tpl/curd.tpl
// assets/tpl/curd2.md
package bindata

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4f\x6f\xe3\x44\x14\x3f\xe3\x4f\xf1\xf0\x61\xe5\xd9\x58\x5e\x09\x10\x87\x5d\x45\x68\xb5\x04\x88\x14\x82\x9a\xa4\x5c\xaa\xaa\x72\xed\x97\x60\x31\x99\x49\x67\xc6\x24\xc8\xf2\x1d\x21\x01\xa7\x96\x03\x42\xa2\x37\x4e\xa8\x07\xa4\x1e\x10\x1f\x27\x55\x3f\x06\x9a\x19\xd7\x75\x12\xa7\x4e\x4a\xeb\x4b\x6b\xcf\x9b\xf7\xe7\xf7\xfb\xbd\xf7\xe2\x38\x8e\x33\x4e\x59\x04\x13\x54\x59\x16\x0c\x95\x48\x23\x35\x0a\x4f\x29\xf6\xc3\x29\xe6\x79\xf9\xaf\x47\x40\x2a\x91\xb0\x09\x64\x0e\x00\x80\x40\x95\x0a\x06\x6e\x96\x05\x15\x73\xd7\xc9\x1d\x27\xe2\x4c\x2a\x48\x98\x44\x51\xeb\xf3\xb3\x04\x69\x2c\xa1\x6d\x2e\x77\x8d\x99\xf9\xd4\x4b\xa4\xd2\x2e\x0a\x07\x34\x91\xcd\xd7\xdf\x52\x5a\x77\x37\x3e\xad\xbb\xa9\xef\xa4\x12\x85\x6b\x6b\xf6\x12\x85\x53\x78\x59\x67\x49\x60\x18\x85\xcc\x13\x7c\x0e\xf1\x69\x30\xe0\x73\x02\x1e\x0a\x01\x28\x04\x17\xc4\x22\x90\x8c\xf5\x2b\xb4\x41\xf0\x79\x60\xcc\xcd\x67\xfd\x64\x99\x08\xd9\x04\x21\xe8\xa7\x45\x7e\xb2\xcb\xc6\x3c\xcf\x5f\xe8\x90\x41\x96\x05\x5f\xa4\xd3\x99\x0d\xe5\x3b\xf6\x06\xb2\x38\xcf\xc9\x1b\xe3\xf3\xfd\x36\xb0\x84\x16\x48\xdf\xa3\x6d\x5e\xf3\x0a\xfc\x1a\xed\x57\xaf\xe0\xe6\xfa\x9f\xe5\x8f\x7f\x58\x1e\xc3\x38\xae\x2d\xdd\x8b\xd4\x02\x22\xce\x14\x2e\x54\xf0\xce\xfe\xf5\xe1\xbb\x90\xa6\x08\xb5\x08\x78\x34\x94\xaa\x1b\x43\xc2\xd4\xc7\x1f\xf9\xb0\x5e\x7d\xc4\x19\x83\xd7\x6d\x0d\xcf\xe7\xa8\xb4\x77\xbf\x00\x97\x98\x73\x79\x46\xf5\xb1\xdb\xed\x0f\x3b\x83\x11\x74\xfb\xa3\xaf\xc0\x85\xd6\x2e\x3a\x6b\x81\x0b\x9e\xb6\x6d\x94\x50\x0b\x5c\xa2\xbd\x96\x38\xe9\xc7\xfd\xfa\x6d\xef\xb0\x33\x04\xaf\x54\xd7\x97\xa1\xf8\x36\xcf\x89\x6b\xcc\xce\x8a\xac\x87\x07\x3d\x7b\xea\x35\xa7\xe4\xeb\x72\x48\x01\xbc\xb4\x60\xbc\x6e\x1b\x0c\x82\xce\x02\xa3\x02\x4f\x8b\xc2\x99\xbf\xa9\x03\x1b\xc9\x6a\xc0\x60\x5e\xa7\x81\xaa\x0e\xaa\x0a\x6b\x56\x83\xa5\xca\x87\x13\xad\x46\x94\x41\x4f\xbf\xdb\x90\xb1\x47\xd6\x04\x63\x64\x12\x23\xdd\x43\x26\x89\x91\x01\x8a\x71\x18\x61\x96\xd7\xf4\xc2\xae\x6a\x88\x91\xa2\x42\x18\x0b\x3e\xdd\x47\x0d\xf3\x6f\x50\xa0\xce\xa2\x0d\x9f\x6c\xb0\xf8\xa9\xf1\xb9\x1f\x8b\x27\x7e\xd1\xbb\xdb\x28\x84\x24\x5e\xc7\x4d\x77\xda\xed\xcf\xd7\xcb\x5f\x2e\x96\x3f\x9d\xdf\xfc\x7e\x79\xfb\xd7\xd5\xf2\xdf\xf3\x07\xa7\xe7\x16\x38\x6d\x39\x76\x9c\xfa\x10\x8a\x89\x84\xa3\xe3\x2a\xc0\x66\xf0\xd4\x39\xdc\xad\x0f\xeb\xa7\x5f\xc9\xc4\x08\x17\xca\xb0\x21\x91\x62\xa4\x0c\x11\x0d\xc3\x56\x93\xb0\x2f\x69\xda\xd4\x54\xba\x4e\xd8\xd0\x84\xdd\x95\x30\x9d\x6d\x99\xfa\x80\xcf\xcb\xc6\x3b\x48\x51\x7c\x3f\xe0\xf3\x75\xe6\x34\x9e\x41\x10\x90\x2d\x43\xda\x7a\x21\x6f\x74\x32\x5d\xd9\xe7\x03\x3e\x97\x1d\x21\xb4\xa4\x49\xd1\x63\xef\xd9\x2b\x2c\xa1\x5b\x27\xee\x6f\x7f\xdf\x5c\x5c\x59\xee\xd3\x59\x1c\x2a\xac\x2b\xa5\x9e\x7d\x6b\x3f\x54\xa2\x54\x40\xb3\x1e\xee\x59\x87\xff\x4f\xbb\x4d\x60\x1f\x2a\x25\x5a\x95\xdc\xa7\xfe\x20\xc1\x87\xc6\x6c\x7f\x82\x9b\xbb\x72\x85\xdb\x15\x46\x96\x3f\xfc\x7a\x7b\xf9\xa7\x65\x64\x9b\x98\x1f\xdb\x8e\xa0\xfb\x51\x0e\x50\xa6\x54\xc1\xd1\x71\x53\x63\x3e\x01\x45\x8f\xe9\xcc\x56\x33\xde\xad\x67\xe8\x4a\x0d\xcd\xea\x46\x34\x8d\xb9\x4b\x57\x6e\x2c\xb6\xda\xe5\x16\xe3\x18\x85\x09\x13\xbc\xa3\x5c\x62\xb1\xd0\xc6\xbc\xf8\xd8\xd7\x61\xc8\xba\x17\x3b\x28\xea\x4a\xc9\xf2\x15\xcb\xcd\x09\xa1\xbd\x6e\xfb\x21\x56\x93\xe5\xdd\x93\xaf\x27\x70\xa7\x98\x36\x84\xb3\x19\xb2\xb8\xa2\x22\x5f\x9f\x93\x4a\x91\x65\x0a\x32\xd0\xc3\x68\x63\xf7\x58\x61\x47\x3c\x65\x4f\xbb\x68\x14\x57\x21\xd5\xcb\xfd\xc3\x0f\x9e\x65\xb7\x98\x8c\xbd\x97\xe4\x4e\xa3\x3b\x4c\x9c\xe7\x10\x69\x65\xac\x34\xad\x0d\xab\x81\x17\x06\x98\x4d\x1a\x9c\xff\x02\x00\x00\xff\xff\x75\x53\x88\x65\x31\x0d\x00\x00")

func assetsTplCurdTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurdTpl,
		"assets/tpl/curd.tpl",
	)
}

func assetsTplCurdTpl() (*asset, error) {
	bytes, err := assetsTplCurdTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 3377, mode: os.FileMode(420), modTime: time.Unix(1610528210, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplCurd2Md = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4f\x6f\xe3\x44\x14\x3f\xc7\x9f\xe2\xe1\xc3\xca\xde\x5a\x8e\x04\x88\xc3\xa2\x08\x2d\x6d\x59\x22\x85\xa0\x6d\xba\x5c\x56\xab\xca\xb5\x5f\x8a\xc5\x64\xa6\x9d\x19\x93\x22\xcb\xf7\x15\x12\x70\xda\x45\x08\x21\xed\x4a\x1c\x38\xa1\x3d\x20\xed\x01\xf1\x71\x52\x95\x6f\x81\x66\xc6\xc9\xba\xce\xb8\x76\x43\x0a\x1c\xf0\xa5\x75\xfc\x32\xf3\xde\xef\xcf\xbc\xe7\x38\x8e\xe3\x4c\x33\x1a\xc3\x09\xca\x3c\x0f\x27\x92\x67\xb1\x3c\x8c\x8e\x09\x8e\xa3\x19\x16\xc5\xea\x5f\xcf\x07\x21\x79\x4a\x4f\x20\x77\x00\x00\x38\xca\x8c\x53\x70\xf3\x3c\xac\x84\xbb\x4e\xe1\x38\x31\xa3\x42\x42\x4a\x05\x72\xeb\x9a\x1f\xa5\x48\x12\x01\x03\x70\xdd\x65\x30\x49\x45\xd7\xd0\xe4\xd8\x16\xa8\x42\x32\x81\xdc\x35\xe5\x78\xa9\xc4\x19\xdc\xb5\x45\xfa\x30\x89\x23\xea\x71\x36\x87\xe4\x38\x3c\x60\x73\x1f\x3c\xe4\x1c\x90\x73\xc6\x7d\x53\x5c\x3a\x55\xb7\x30\x00\xce\xe6\xa1\x0e\xd7\x1f\xab\x2b\xcf\x79\x44\x4f\x10\xc2\x71\x46\x88\x49\x6f\x48\xa7\xac\x28\x56\x11\x77\xd4\x97\xf2\x3c\xfc\x38\x9b\x9d\x9a\x2d\x83\xca\xb7\x91\x26\x65\xac\xff\xbe\xde\xe4\xad\x01\xd0\x94\x94\xa8\xbe\x41\x56\xdf\x16\x15\xa8\x15\xb2\xfd\x3e\x5c\xbc\xfe\x7d\xf1\xf5\x0b\xc3\x59\x94\x24\x5e\x2c\xcf\x21\x66\x54\xe2\xb9\x0c\x77\xcd\xdf\x00\xbe\x8c\x48\x86\x60\xab\xde\x23\x91\x90\xc3\x04\x52\x2a\xdf\x7b\x37\x80\x7a\xe1\x31\xa3\x14\xee\x0d\x14\x32\x0f\x50\xaa\xc5\x83\x12\x57\x5f\x3f\x17\x67\x44\x3d\x76\x87\xe3\xc9\xfe\xc1\x21\x0c\xc7\x87\x9f\x82\x0b\x3b\x5d\xd4\xb3\x03\x2e\x78\x79\x1e\x0e\xb5\x2e\x34\x72\xa3\x54\x48\x45\x88\x0b\x3b\xab\xea\xd5\xe5\x7e\x76\x7f\xf4\x68\x7f\x52\x09\xff\x24\xe2\x5f\x14\x85\xef\xea\xb0\xb3\x32\xc3\xc9\xc3\x91\x79\xea\xb5\x6f\x1f\xa8\xd4\xfd\x12\x4e\x61\x0a\xbf\x37\xd0\xf5\x86\xfb\xe7\x18\x97\xd0\x99\x8a\xcf\x82\x75\xba\xcd\x4e\x86\x6a\x0d\x6f\x0b\xc5\x7e\x55\x48\xed\x1c\x1b\x5a\x02\x38\x52\xa2\x43\x11\x8e\xd4\xbd\xd9\x32\xf1\xfc\x9a\x0c\x34\xf9\x09\x12\x3b\xf9\xa9\x66\x17\xf9\x34\x8a\x31\x2f\x2c\xea\xee\x4a\x72\x82\x04\x25\xc2\x94\xb3\xd9\x4d\x48\x9e\x7f\x8e\x1c\x55\x16\x03\xf8\x60\x8d\xb0\x3d\xbd\xe6\xcd\x08\x3b\x0a\x4a\x37\x36\xb1\x05\x69\x52\x87\x68\x75\xae\xdd\x4f\x12\xab\x0f\x6e\xe8\x1b\xbf\xd1\x38\xd0\x0c\xaa\x6d\xa1\xa3\xbd\x0f\xff\x29\x2b\xfd\x6f\xa3\xeb\x6d\xd4\xef\xc3\xe2\xe9\x8b\x3f\x7f\xf8\x79\xf1\xcd\xb3\x8b\x9f\x5e\x5e\xfe\xfa\x6a\xf1\xc7\x33\xa3\x9a\x3d\x24\xdd\x55\x63\x04\x6f\xba\x63\x00\x11\x3f\x11\xf0\xf8\x49\x93\x05\x37\x91\x4b\xcf\xe6\xc8\xab\x8d\x57\x4b\x47\x27\xe2\xf4\xd6\xed\x56\xef\xd2\x25\x8d\x4e\xaf\xdd\x5a\xaa\x9c\x30\x0c\x2d\xd8\x5d\x7e\xfb\x7a\xf1\xdd\xf3\x75\xec\x1e\xd8\x55\x64\xc7\x6e\x6a\x1a\xfc\x12\xbc\x56\x28\x75\xdf\xb6\x2d\xdf\xad\x97\xb5\x39\xf2\x10\xcf\xa5\x46\x5a\x20\xc1\x58\x6a\x54\xcb\x14\x95\xe9\xae\x07\xbe\x66\xaa\x89\x5e\xc2\x8e\xbd\xda\xc7\x6c\x5a\x81\xff\x61\x86\xfc\xab\x03\x36\x6f\xa2\xa0\x36\x89\xa8\xab\xd7\x6b\x9c\x47\x2c\x63\x48\xbf\xaf\x92\xdf\x65\xb3\x19\x52\x59\x99\x58\x96\x2b\x55\x1c\x77\x75\xe6\xf8\xf1\xb7\x8b\xe7\xaf\x0c\xbb\x8f\x4e\x93\x48\x62\x77\x82\x33\x1d\x3f\x91\xfc\x06\x1c\x6f\xe5\x70\x5d\x51\x69\x32\xa8\xd3\x26\xd0\xb0\xfb\x26\x3f\x45\x70\x23\x99\xa6\xec\x36\x32\xff\x8e\x9d\x16\x4f\xbf\xbf\x7c\xf9\x8b\x01\x59\x8d\xc3\x1b\x9f\x35\x9c\xcd\xc5\x01\x8a\x8c\x48\x78\xfc\xa4\xcd\x2a\xdb\x00\xb8\xe2\x95\x96\x39\xfe\x96\x3c\xa4\x2a\xbe\xda\x8b\xb4\x95\xae\xc5\xbe\xa9\xa5\x58\xdb\x4a\x82\x53\xe4\x7a\x9b\x70\x97\x30\x81\x65\x2b\x99\xb2\xf2\xc3\xb1\xda\xc6\xaf\xaf\xc2\xe6\x2a\x1f\x1b\x1c\xf9\x55\xef\xad\xbf\x6e\xa8\x55\x9b\xde\x0f\x2c\x59\x2e\xaf\xa2\x9e\xc0\x52\x08\x03\x88\x4e\x4f\x91\x26\x15\x71\x04\xea\xb9\x5f\x29\x72\x95\x82\x08\xf7\x39\x5f\xeb\x96\x75\x8d\x8e\x1a\xa8\xde\xde\x39\xff\xaf\xa9\xf8\x76\x4f\xfc\x6d\xa9\xb5\xd7\xab\x28\xa0\xd7\x28\xd4\x16\xa5\x76\x53\x69\x55\xa1\xc2\xd2\x87\xb6\xd5\x85\xac\x1d\x69\x15\x71\x9d\x23\x6a\x6e\x28\x9c\x0d\x5c\x60\xb7\x41\xdd\x07\x5a\xfd\x31\xcb\xe8\xa6\x47\xb4\x27\x99\x8c\x88\x7a\x93\x78\xe7\xed\x6e\x63\x8b\xfd\x37\x8f\x46\x0d\x9b\xe4\xee\xfa\xff\x89\x79\xe5\x8e\xae\xd6\xf2\x7e\xf6\x57\x00\x00\x00\xff\xff\xea\x88\x00\xbd\x7b\x12\x00\x00")

func assetsTplCurd2MdBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurd2Md,
		"assets/tpl/curd2.md",
	)
}

func assetsTplCurd2Md() (*asset, error) {
	bytes, err := assetsTplCurd2MdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd2.md", size: 4731, mode: os.FileMode(420), modTime: time.Unix(1610523290, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"assets/tpl/curd.tpl": assetsTplCurdTpl,
	"assets/tpl/curd2.md": assetsTplCurd2Md,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"assets": &bintree{nil, map[string]*bintree{
		"tpl": &bintree{nil, map[string]*bintree{
			"curd.tpl": &bintree{assetsTplCurdTpl, map[string]*bintree{}},
			"curd2.md": &bintree{assetsTplCurd2Md, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}