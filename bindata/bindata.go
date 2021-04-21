// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package bindata generated by go-bindata.// sources:
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

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xdd\x6e\xd3\x30\x14\xc7\xaf\xc9\x53\x1c\x72\x31\x25\x5b\xe4\x49\x80\xb8\xd8\x54\x21\xd8\x06\x54\x1a\x45\x6b\xe1\x6a\x9a\xa6\x2c\x39\x2d\x11\xae\xdd\xda\x0e\x19\xb2\xf2\x02\x3c\x05\x6f\xc1\x05\xcf\x03\xcf\x81\x6c\x67\x5d\x9a\xba\x94\x8d\x22\x7c\x33\x25\x3d\x3e\x1f\xff\xdf\x39\x27\x0b\xcc\x19\x97\x2c\x8b\x0a\x85\x53\xd8\xd5\x9a\x8c\x94\x28\x33\xf5\x2e\xbd\xa2\x38\x48\xa7\x58\xd7\x31\x8c\xb2\x94\x45\x82\x57\x90\x5f\x91\x21\xaf\x62\x88\x50\x08\x40\x21\xb8\x88\x75\x00\x00\x50\x8c\xcd\x23\xf4\x40\xf0\x8a\x58\x73\xfb\xda\x1c\xad\x45\xca\x26\x08\x64\x50\x52\xfa\xb2\x40\x9a\xcb\x3e\x1b\xf3\xba\xde\x31\x21\x89\xd6\xe4\x75\x39\x9d\xb9\x50\x49\xe0\x6e\x20\xcb\xeb\x3a\x3e\xb4\x3e\x1f\xf6\x80\x15\x14\xf4\xc2\xa1\x40\x55\x0a\x66\x1f\xeb\xa0\xf5\xa2\x0e\x82\xfd\x7d\xf8\xf9\xed\xfb\x8f\x2f\x5f\x6d\x4d\x90\xe6\xb9\xaf\x20\x88\x32\x75\x0d\x19\x67\x0a\xaf\x15\x39\x72\x7f\x13\xb0\x0a\x78\x05\x88\x68\x2a\x55\x3f\x87\x82\xa9\xa7\x4f\x12\xe8\x16\x9f\x71\xc6\xe0\xa0\x67\xd4\x79\x85\xca\x38\x4f\x20\xd4\x9a\x1c\xbf\x70\xf7\xc3\xd8\x9a\xc9\x39\x35\x56\x61\xc1\x24\x0a\x65\x9c\x71\x13\xee\xad\x28\x26\x05\x5b\x4a\x4f\x6b\xd2\xb7\x46\x56\xae\xd3\x42\x2a\x43\x21\x84\xbd\x85\x06\xe6\x84\x9f\x52\x5a\xa2\x6c\x99\xbf\x49\xc5\xc7\xba\x8e\x43\x6b\x36\x6f\x52\x1a\x9d\x9d\xba\x5f\xa3\xd0\x17\x2d\x4c\x4c\x62\x71\x23\xa4\x74\xd5\x1d\xf4\x6c\x51\xe4\xe4\x1a\xb3\x46\x1f\x57\xd6\x3c\x59\xe5\xea\xbc\x3b\xa6\xeb\x90\xb6\xb1\xb6\x1b\x66\x33\x5c\x27\x7d\x02\x97\xa6\xb9\x50\x92\x53\xf3\xec\x22\xe6\x51\xdc\xe1\x6f\xa9\xe7\x48\xef\x42\xdd\x62\x45\x31\x4e\x33\xd4\xb5\xa7\xb5\xef\x48\x37\x47\x8a\x0a\x61\x2c\xf8\xd4\x4f\xb7\xfa\x80\x02\x4d\xd8\x1e\x3c\x5b\x21\x75\x6c\x6f\x6f\x26\x75\x99\x34\xf3\xb6\x0e\x13\x14\xb9\x57\x9c\x09\xaa\x3b\x88\xe3\x72\x95\x4a\x14\x6c\x92\x40\x2a\x26\x12\xce\x2f\xda\x72\xd9\xad\xe0\x73\xf8\xd7\x53\x22\x91\x62\xa6\x8c\xef\xe7\xcd\xda\x70\x73\xf0\x1b\x65\x43\xd8\x73\x19\x77\x55\x1d\x59\x5f\x9b\x55\x95\x73\x3a\xe4\xd5\xa2\xfb\xcf\x4a\x14\x9f\x87\xbc\xea\x4a\x6b\x74\x20\x84\xc4\x6b\x36\x9f\xf3\x12\x1f\x9a\xe0\x7d\x39\xe0\x43\x5e\xc9\x13\x21\x4c\x63\xc5\x4d\xa7\x3f\x70\x57\x58\x41\xbd\x6b\xcc\x92\x2a\x67\x79\xaa\xd0\xa7\xad\x9f\x95\xb3\x1f\x29\xb1\xe0\xb5\x99\xde\x2d\x23\xb8\x17\x24\x17\xd3\x4f\x43\xa2\xb2\x44\x6e\xf3\xda\x83\x70\x3d\xa3\xf7\xd6\x6c\x1b\x9d\xbf\x84\xa7\x23\x2a\x2d\xe4\x56\xfb\x1f\xcc\x00\xc8\x21\xca\x92\x2a\x38\xbf\xd8\x34\x09\xf7\x53\xf9\x7f\x8c\x82\x29\x6b\xf9\x5b\x60\xa7\xe1\x4f\x46\x61\x65\xa7\x7b\xf7\x7a\x8e\x63\x14\x36\x0c\x39\xa2\x5c\x62\xb3\xcb\xc7\xbc\x79\x39\x30\x61\xe2\xae\x17\x37\x9d\x3e\x99\x75\xbd\x64\xb9\x3a\x96\xc6\xeb\xba\x7f\x29\x3c\x59\xde\x9c\xba\x9b\xc0\x0d\xed\x1e\xa4\xb3\x19\xb2\xbc\xd5\x01\x89\xf9\x3d\x6e\x15\xb9\x48\x41\x12\xb3\x01\xbc\x2d\x99\xf1\x92\x6d\x77\x27\x2b\xae\x52\x6a\xbe\x6a\x8f\x1f\x6d\x6b\x0d\xdb\x24\xa3\xdd\xf8\x5f\xb6\x5c\x6b\xac\x37\x6d\x5e\x47\x74\xc7\x16\xda\x15\x35\x08\x82\x5f\x01\x00\x00\xff\xff\x87\xba\xbb\x2d\xd6\x0a\x00\x00")

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

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 2774, mode: os.FileMode(420), modTime: time.Unix(1618990753, 0)}
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

	info := bindataFileInfo{name: "assets/tpl/curd2.md", size: 4731, mode: os.FileMode(420), modTime: time.Unix(1610531769, 0)}
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

	info := bindataFileInfo{name: "assets/tpl/mc.tpl", size: 1098, mode: os.FileMode(420), modTime: time.Unix(1617161061, 0)}
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
	"assets/tpl/mc.tpl":   assetsTplMcTpl,
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
			"mc.tpl":   &bintree{assetsTplMcTpl, map[string]*bintree{}},
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
