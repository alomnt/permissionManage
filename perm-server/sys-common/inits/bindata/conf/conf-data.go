// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// conf/app.yml
// conf/db.yml
package conf

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

var _confAppYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x41\x4a\xc3\x40\x14\x86\xf7\x03\x73\x87\x07\xdd\xb6\x64\x2a\x56\xe2\xec\xb4\xb6\x62\x69\x6d\xd1\x14\x17\x22\x32\x4d\x9e\x49\x60\x32\xaf\xcc\x4c\x6c\x3d\x80\x88\x0b\x2f\xe0\x11\x5c\x74\xef\xc2\xdb\x28\xea\x2d\x64\xd4\x66\xf7\xfe\xff\x7b\xfc\x7c\x47\xa5\x53\x0b\x8d\x33\xe5\x8b\x3e\x59\x8b\xa9\x2f\xc9\x48\xb8\x51\xda\x21\x67\x03\xb3\x85\x03\x97\xaa\x25\x36\x60\x58\x5a\x9c\xa0\x2f\x28\x3b\x25\x7f\xa0\x35\xad\x30\x6b\xe0\xff\xe6\x21\x65\x77\x7d\x32\xae\xae\x96\x61\x74\x6a\xe6\xa6\x52\xd6\x15\x4a\x37\x9f\x49\x59\xe1\x90\x6c\xa5\xbc\x84\x09\x99\x36\x88\x2e\x8c\x94\x81\x1d\x21\xf6\xa0\xdb\x93\x62\x57\x8a\x1e\x1c\x4f\x12\xce\xfa\x85\xb2\x0e\xbd\x84\x79\x32\xec\xc4\x9c\x71\xd6\xfa\x7a\x78\xf9\xd8\x3c\xbf\xbf\x3e\x7e\xdf\x3f\x7d\xbe\x6d\x38\x9b\xfa\x02\xad\xe4\x0c\x60\x46\xd6\x4b\x88\x45\x1c\x87\x74\x92\x1b\xb2\x38\x3f\x1b\x3b\x09\x97\x51\x1b\xa2\x82\x2a\x6c\x43\x54\x3b\xb4\x91\xa6\xbc\x34\xdb\x60\x31\x2f\x9d\x0f\xec\xb7\xbe\x5e\xae\xb2\xab\xb0\x30\xba\x48\x82\x2b\xd5\x5e\xc2\xbe\x10\xd0\x72\x98\x92\xc9\x02\x1a\x53\x3e\xc6\x5b\xd4\x12\x32\x5c\xd4\x79\xa8\xce\x31\xb5\x41\x75\xbd\x5e\x77\xfe\xee\xe0\xcb\xd9\x4f\x00\x00\x00\xff\xff\x3a\x5d\xab\xd4\x6d\x01\x00\x00")

func confAppYmlBytes() ([]byte, error) {
	return bindataRead(
		_confAppYml,
		"conf/app.yml",
	)
}

func confAppYml() (*asset, error) {
	bytes, err := confAppYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.yml", size: 365, mode: os.FileMode(438), modTime: time.Unix(1554210936, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _confDbYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\xd0\xb1\x4a\xc3\x60\x10\xc0\xf1\x3d\x90\x77\x38\xe8\x5e\x12\x0b\x2a\xb7\x3a\x09\x82\x83\x4f\x70\x49\xce\xa6\xf0\x25\x5f\xfb\xdd\xa5\xd5\xad\x83\x8b\x8b\x0e\x52\x04\x11\x71\xa8\x14\xb7\xe8\x28\x82\x4f\x93\xaf\xfa\x16\xf2\x05\x41\xdf\xc1\xf1\x7e\xdc\xff\x86\x8b\xa3\x8a\x44\xd9\x61\x1c\x01\x14\x13\x32\x9c\x2b\x42\x75\x2e\x33\x13\xa4\x11\x76\x08\xce\x5a\x0d\xd3\x94\x44\x16\xd6\x15\xbf\x52\x5a\x51\x84\x74\x67\x6f\x98\x0c\x93\x61\xda\x2f\x59\xa7\x08\xa3\x51\xb2\xdb\x9f\x24\xa5\x8c\x84\x11\x72\x92\x6c\x52\x07\xcb\x4b\x72\xc2\x8a\xd0\xe8\xe9\x7e\x00\x29\xed\xe2\x64\x66\x10\xd4\x35\x1c\xc0\xd8\xf1\x11\xcf\xd9\x20\x14\x9c\x35\xe3\x40\x15\x9d\x1d\x16\x86\x0f\x6c\x5d\x0b\x42\x9a\xc0\x00\x3e\x3f\x1e\xfc\xd5\x93\x7f\x79\xdc\xde\x5d\x6c\x9f\xdf\xbe\x6e\x5f\xfd\xaa\xed\xd6\x9b\xae\xbd\x8e\xa3\x41\x9f\x1c\x4f\xb9\xfe\x9b\xf8\xfb\x65\xb7\xde\xf8\xcb\x9b\xee\x7d\xf9\x93\xaf\xda\x38\x12\x43\x73\xfe\xcf\x2f\xf8\x0e\x00\x00\xff\xff\xa6\xd8\x6a\xbd\x07\x02\x00\x00")

func confDbYmlBytes() ([]byte, error) {
	return bindataRead(
		_confDbYml,
		"conf/db.yml",
	)
}

func confDbYml() (*asset, error) {
	bytes, err := confDbYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/db.yml", size: 519, mode: os.FileMode(438), modTime: time.Unix(1553698957, 0)}
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
	"conf/app.yml": confAppYml,
	"conf/db.yml":  confDbYml,
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
	"conf": &bintree{nil, map[string]*bintree{
		"app.yml": &bintree{confAppYml, map[string]*bintree{}},
		"db.yml":  &bintree{confDbYml, map[string]*bintree{}},
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
