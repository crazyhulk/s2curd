// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/tpl/curd-struct.tpl
// assets/tpl/curd.tpl
// assets/tpl/curd2.md
// assets/tpl/mc.tpl
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

var _assetsTplCurdStructTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xcf\x4a\x33\x31\x14\xc5\xd7\xc9\x53\x1c\x0a\xdf\x47\x52\xcb\x74\x5f\x98\x9d\x2e\x75\x23\xba\x29\x5d\xc4\xf4\x8e\x1d\x9d\x66\xe4\x26\xd3\x52\x86\xbc\x42\x7d\x03\x37\xe2\x56\x7c\xa9\xea\x6b\x48\xa6\x7f\x6c\x05\x57\x33\x84\xf3\x3b\x97\x1f\xa7\x68\x9c\x85\xf2\xe8\xb7\x6d\x76\x65\xe6\x14\xa3\xc6\xb5\x35\x4e\x79\xb6\x28\x5d\x20\x2e\x8c\xa5\x36\x6a\x28\x62\x06\x31\xd7\xac\xd1\x4a\xd1\xf7\xc8\x71\x80\xda\x28\x45\x59\x20\x41\x79\x0e\x57\x56\x29\x22\x98\x42\xc3\x4e\x8a\x28\x85\x5f\x96\xc1\xce\x52\xe0\xdc\x04\x83\x51\x9e\x7e\x33\x15\x56\x4f\xd4\xd5\x59\xe3\x09\x3e\x70\xe9\xee\x47\x52\x88\x74\x2b\xc7\x83\xaf\x5d\x76\xe3\xe6\x86\xfd\xcc\x54\x6a\x3c\xb9\x5b\x05\x52\xbb\x0e\x3d\xc0\x7f\xaf\x8f\xae\x74\x15\xdb\xcc\x9f\x15\x3b\xf6\x37\x3a\xa5\xc2\x34\x55\xf8\xc1\x3a\x4f\x9f\x5d\xa4\x4f\xa1\x7a\xe3\x83\xe8\x04\x8d\x7b\x74\xf5\xd2\xc1\x5b\xe3\xe0\xeb\x86\x2d\x8d\xf0\xef\x6c\xd1\x1b\x24\x25\xdd\xd9\xee\x7a\xa3\x94\xc3\x21\x6e\x4d\xd5\x10\x36\xef\x2f\x5f\xeb\x0f\x4c\xb9\x5c\x10\x67\xdd\x1b\xe3\x73\xfd\xb6\x79\x7e\x95\xfb\x11\x8e\x36\xe8\x02\x4a\x43\x2d\x4e\x90\x01\x4e\x57\xd8\x1e\xda\x5a\x5e\xee\x1d\xb5\x8c\xf2\x3b\x00\x00\xff\xff\xc2\x77\x1d\xab\xd8\x01\x00\x00")

func assetsTplCurdStructTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurdStructTpl,
		"assets/tpl/curd-struct.tpl",
	)
}

func assetsTplCurdStructTpl() (*asset, error) {
	bytes, err := assetsTplCurdStructTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd-struct.tpl", size: 472, mode: os.FileMode(420), modTime: time.Unix(1640660313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5d\x6e\xdb\x38\x10\x7e\xd7\x29\x66\x05\x23\x90\xb2\x02\x03\xec\x2e\xf6\x21\x85\x11\xb4\x49\xda\x06\x48\x53\x24\x6e\x9f\x82\xa0\x50\xa4\x91\x4b\x94\x26\x1d\x92\xaa\x53\x10\xbc\x40\x4f\xd1\x5b\xf4\xa1\xe7\x69\xcf\x51\x90\x94\x5d\x59\x96\xf2\xaf\x87\xc4\xa2\x87\x33\xdf\x7c\xdf\xc7\xa1\x23\xf7\x54\x35\x2f\x12\xaa\x71\x06\xdb\xc6\x90\x89\x96\x75\xa1\xdf\xe5\x97\x0c\x4f\xf2\x19\x5a\x9b\xc2\xa4\xc8\x79\x22\xc5\x02\xca\x4b\x72\x26\x16\x29\x24\x28\x25\xa0\x94\x42\xa6\x26\x02\x00\xf7\x19\xc6\x20\xc5\x82\xf8\x58\xbf\xe6\x1e\x63\x64\xce\xa7\x08\xe4\xa4\x66\xec\x25\x45\x56\xaa\x23\x5e\x09\x6b\xb7\x5c\x3d\x62\x0c\x79\x5d\xcf\xe6\xa1\x4e\x16\x85\x1d\xc8\x4b\x6b\xd3\xc8\xbf\x49\xd4\xb5\xe4\x91\x8d\xa2\x9d\x1d\xf8\xf5\xfd\xc7\xcf\xaf\xdf\x3c\x5c\xc8\xcb\xb2\x0f\x2b\x24\x85\xbe\x86\x42\x70\x8d\xd7\x9a\xec\x87\xff\x19\xf8\xe6\x7a\x7b\x4b\x58\xae\xf4\x51\x09\x94\xeb\xff\xff\xcb\xa0\xdb\x57\x21\x38\x87\xdd\xb1\x6b\xfc\x15\x6a\x97\x3c\x83\xd8\x18\x72\xf0\x22\xec\x8f\x53\x1f\xa6\xae\x98\x8b\x8a\x29\x57\x28\xb5\x4b\x26\x5c\xb9\xb7\x92\x4e\x29\x5f\x83\x17\xc3\xdf\x2b\x72\xda\x04\x8d\x68\x36\x42\x9f\x84\x1c\xf9\x24\x9e\xac\x63\xaa\xb4\xb5\xc6\xd0\x0a\x46\xd4\xda\x2c\xee\x6e\x46\xa6\xd0\x05\x78\xce\x62\x63\x46\xc8\xfe\xbc\xae\x07\xc7\x29\x7c\xce\x59\x8d\x0a\x12\x63\x9a\x22\x6f\x72\xf9\xc9\xda\x34\xf6\x81\x57\x4d\xa3\x93\xd3\xe3\xf0\x6d\x12\xf7\xf5\x10\x67\xae\xdd\x95\x3e\x2a\x90\xb6\x3b\xf6\x5c\x91\xc3\x6b\x2c\x1a\xda\x03\x5b\x57\xd9\xa6\x19\x42\xfa\x60\x84\x21\x1f\xac\x79\xc1\xbd\xd0\xca\x17\xfa\x6b\x0c\x9c\x32\x30\xab\xa0\xc6\x22\xee\xa3\x0d\xa0\x82\xa4\xd9\xd2\x94\xa8\xc8\xb1\x5b\x09\x45\xcb\x64\xc3\x5a\xde\x50\x25\xb2\xfb\x18\xca\x3b\x06\x65\x95\x17\x68\x6c\xcf\x81\xb8\xa7\x71\x4a\x64\xa8\x11\x2a\x29\x66\xfd\xc6\x59\x7c\x44\x89\xae\xec\x18\xf6\x36\xe4\x3a\xf0\xbb\xef\x20\xd7\x87\x25\x29\x43\x5a\x01\x2d\xfb\xe9\xa9\xe7\x65\xae\xb1\x8f\xa1\x7e\x82\x02\x5e\xa5\x25\xe5\xd3\x0c\x72\x39\x55\x70\x7e\xd1\xa6\xac\xc5\x18\x3c\x88\xb2\x80\xa8\x9f\x2d\x85\x1a\xba\x27\xed\xd1\xa7\xed\x86\x13\x37\xde\x6b\x16\x5c\xd5\xd0\x7b\x57\xa3\xf7\x1e\xee\x93\x68\xe4\xd8\x24\x84\x6c\x08\x15\x94\x62\x54\xe9\x7b\x38\xf9\x56\xa1\xc0\x4d\x7e\x75\x86\xaa\x66\x1a\xce\x2f\xfa\x52\xb7\xe7\xe6\xc3\xc4\x54\xc8\xb0\xd0\xd0\xc3\xf8\x80\x6c\xcf\x9b\xeb\xe4\x0e\x9a\xdd\x3e\x29\x87\xea\xc6\xc3\xe7\x71\x50\xe7\x89\xef\xe4\x2e\xa3\x53\x2c\x3a\xb3\xf3\xb4\x46\xf9\x65\x50\xec\x9b\x66\xe0\xc6\x1c\x74\x7f\x4b\xac\x50\xfa\x32\x64\x9f\x09\x85\xcb\xc1\x57\x89\x66\xf5\xc4\xd5\x49\xbb\x69\xc4\xc2\x01\xea\x93\xd9\xd8\xb5\xc8\x06\x4c\xeb\xda\x77\x59\xd3\x67\xc3\x10\x3b\x30\x97\x8f\xed\x02\x58\xba\x6d\x0c\xf9\x7c\x8e\xbc\x6c\x39\x30\x73\xdf\xa7\xed\x69\xbf\xc2\xa0\xc8\xa1\x94\x03\xd3\xbd\x10\x35\x7f\xd2\x53\x91\x68\xa1\x73\xe6\x2e\x81\x7f\xff\x79\xf4\xcf\x86\xc6\xfd\x1e\x64\xb2\x9d\xde\x70\x0d\x3c\xde\x76\xad\xd9\xe2\xfd\x76\x26\x16\x43\x96\x0b\xa2\x6e\xf9\x4e\x37\x87\x4d\x14\xfd\x0e\x00\x00\xff\xff\xad\x9d\x71\xeb\x3d\x0a\x00\x00")

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

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 2621, mode: os.FileMode(420), modTime: time.Unix(1640576872, 0)}
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

	info := bindataFileInfo{name: "assets/tpl/curd2.md", size: 4731, mode: os.FileMode(420), modTime: time.Unix(1640576872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplMcTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x53\xbd\x8e\xd4\x30\x10\xae\x9d\xa7\x18\x5c\xa0\x44\x8a\xb2\x27\x21\x9a\x48\x5b\xdd\xc2\xe9\xb4\x3a\x0a\x16\xd1\x22\xaf\x77\xb2\x67\x92\xd8\x91\x33\x91\x12\xad\xd2\x22\x6a\x68\xe8\x78\x02\xde\x80\xd7\x39\x21\xde\x02\x39\x76\xf6\x76\x0f\x8e\x0a\x09\x37\xf1\x78\xfe\xbe\x6f\xbe\x49\x23\x64\x29\xf6\x08\x7d\xdf\x47\x91\xaa\x1b\x63\x09\xe2\x08\x00\x80\x4b\xa3\x09\x7b\xe2\xde\x42\x2d\xcd\x4e\xe9\xfd\xe2\x7d\x6b\x74\x78\x2b\x6a\xe2\x91\xbf\xb6\x5a\x35\x68\x17\x1d\xa9\x6a\x21\x85\xbc\xc5\x12\x07\xfe\xbb\xab\x96\x3c\x4a\xa2\x48\x1a\xdd\xce\x7d\x16\x0b\x78\xa9\x7a\xb8\xc1\xc9\x3a\x1c\xb2\x0d\xd9\x4e\xd2\x1b\xb1\xad\xf0\x95\xa8\x71\x1c\xd7\x38\xc0\x12\x78\xef\xce\x94\x5e\x74\x5a\x82\xd2\x8a\xe2\xe4\x30\x65\xcd\x1d\xb3\xd7\xb8\x57\x2d\xa1\x8d\x8f\x2f\x6b\x1c\xae\x75\x61\x7c\x1c\x63\xe7\xdd\x18\x73\x1d\xf2\xc7\xba\xa6\x21\x68\x65\x64\x0e\xc0\x79\xb0\xc7\x24\x1a\x03\x88\x0d\xd2\x9f\x72\x2f\x5d\xf7\x58\x52\x0f\x61\x88\xd9\xa5\xff\xa6\xa0\x76\xa0\x34\xa1\x2d\x84\xc4\xc3\x98\xc4\x68\x2d\xa0\xb5\xc6\xce\x54\x20\x5f\x42\x2d\xb3\x2b\x24\x57\x20\x05\xbe\xc3\x42\x74\x15\xf1\x64\xf2\x97\x38\xb8\x88\xa2\xa6\x6c\xd3\x58\xa5\x29\x7e\x04\x7c\xe2\x95\xd9\xa6\xae\xbc\x4b\x71\xc2\x65\x37\xc2\xb6\xb7\xa2\x8a\xd5\xce\x97\x53\xc5\xe4\x7e\xb2\x04\xad\x2a\xf0\x10\xdc\xb1\x48\x9d\xd5\x93\x39\xfa\x42\x2e\x6c\x09\x32\xdb\xcc\xc0\x9e\xd6\x32\xbb\x26\xac\xef\x93\xd6\x38\xe4\xf3\xbd\x9c\xc7\xe7\xce\x5b\x51\x75\x18\x5c\xdb\xfb\xe7\x17\x7d\xa3\xac\x20\x65\x74\x0e\xcf\x2e\x2e\x52\xb7\x0b\x3f\xbe\x7f\xbe\xfb\xf6\x05\x9e\xc3\xdd\xc7\x0f\x3f\x3f\x7d\xf5\x08\x02\x97\x00\x6a\x1e\xfe\xd5\x3f\x1f\xfe\xf9\x76\xfc\x4d\x0a\x16\x74\xe0\xd3\x4e\x46\xec\x5d\x7a\x1c\xd0\x31\xbc\x74\x22\xb0\x87\x13\x66\x2c\xd0\x60\xe3\x03\x52\x9e\xd5\x0a\xab\xff\xc6\xea\x64\xc1\xb8\xff\xdb\x4e\x95\x5f\x61\x85\x84\x27\xdc\xce\xe0\xff\x0a\x00\x00\xff\xff\x8f\x89\x42\xbd\x4a\x04\x00\x00")

func assetsTplMcTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplMcTpl,
		"assets/tpl/mc.tpl",
	)
}

func assetsTplMcTpl() (*asset, error) {
	bytes, err := assetsTplMcTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/mc.tpl", size: 1098, mode: os.FileMode(420), modTime: time.Unix(1640576872, 0)}
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
	"assets/tpl/curd-struct.tpl": assetsTplCurdStructTpl,
	"assets/tpl/curd.tpl":        assetsTplCurdTpl,
	"assets/tpl/curd2.md":        assetsTplCurd2Md,
	"assets/tpl/mc.tpl":          assetsTplMcTpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"tpl": &bintree{nil, map[string]*bintree{
			"curd-struct.tpl": &bintree{assetsTplCurdStructTpl, map[string]*bintree{}},
			"curd.tpl":        &bintree{assetsTplCurdTpl, map[string]*bintree{}},
			"curd2.md":        &bintree{assetsTplCurd2Md, map[string]*bintree{}},
			"mc.tpl":          &bintree{assetsTplMcTpl, map[string]*bintree{}},
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
