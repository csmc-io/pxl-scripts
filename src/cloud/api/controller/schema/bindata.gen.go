// Package schema Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package schema

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x51\x73\xe3\xb8\x0d\x7e\xd7\xaf\x40\x2e\x0f\x97\xcc\x24\x79\xe8\xb4\xf7\xe0\xa7\xba\xb6\xf7\x56\xdd\xc4\x49\x63\x67\xb7\xd7\x9b\x4c\x86\x16\x61\x8b\x13\x89\xd4\x92\x90\x13\xb7\xb3\xff\xbd\x03\x92\x92\x45\x3b\xe9\xce\x76\xe6\x9e\x62\x91\x20\xf0\x11\xc0\x07\x80\x39\x85\x65\xa9\x1c\xac\x55\x85\x20\xd1\x15\x56\xad\xd0\x01\x95\x08\xae\x28\xb1\x16\xb0\xb6\xa6\xf6\xdf\xe3\xbb\x1c\x1c\xda\xad\x2a\xf0\x2a\x3b\xcd\x4e\x21\xa7\x9f\x1d\x68\x43\xa0\x24\x8a\xea\x02\x56\x2d\xc1\x0b\x82\x46\x94\x40\x06\x6a\xa1\x5b\x51\x55\x3b\xd8\xa0\x46\x2b\x08\x81\x76\x0d\x3a\x58\x1b\xeb\xf5\x2d\x77\x0d\x2e\x0a\xab\x1a\x82\x87\x3c\x3b\x85\x97\x12\x35\x50\x0f\x46\x39\x68\x1b\x29\x08\xe5\x55\x80\x58\x08\x0d\x2b\x04\x69\x34\xc2\x6a\x07\xb6\xd5\x5a\xe9\xcd\x28\x3b\x05\xd8\x58\xd1\x94\x5f\xab\xcb\x00\xf9\xd2\xdb\x09\x9a\x3b\xdb\x97\xe4\xe2\x85\xae\xa2\x30\x5c\x5e\x9a\x96\x9a\x96\xba\x75\x79\x45\xce\xc3\x50\x45\x09\x2f\xaa\xaa\x06\xc0\x4b\x84\x28\xcc\xba\x03\x40\x2a\x05\x05\xb9\x15\x42\xa3\x8a\x67\x94\xd0\x36\x0c\x8d\xc5\x1f\xf2\xab\x2c\xfa\x76\xa0\xdf\x9f\x74\xe0\x4a\xd3\x56\x12\xf0\x55\x39\x02\xa5\x83\xbb\x45\x8d\x20\x95\xc5\x82\x8c\xdd\x81\x18\x06\xa1\xc7\xcc\xc7\xaf\xb2\x2c\x86\xe6\x3f\x19\xc0\xd7\x16\xed\x6e\x04\xff\xe0\x3f\x19\x40\xdd\x92\x20\x65\xf4\x08\x6e\xe2\xaf\xec\x5b\x96\x79\xd0\x0f\x0e\x6d\xae\xd7\xc6\x1f\x53\x72\x04\xf9\xf4\x24\x03\xd0\xa2\xc6\x11\x2c\xc8\x2a\xbd\xe1\x6f\xac\x85\xaa\x86\x0b\x8d\x2a\xa8\xb5\x89\x8c\xb1\x9b\x79\x72\xec\x5b\x96\xa1\x6e\x6b\x18\x5b\x52\x6b\x51\x10\xc7\xd6\xdb\x01\x18\x2f\x9f\x1e\xe6\x9f\xe6\xb7\x5f\xe6\xdd\xe7\x75\x3e\x7f\xf8\xe7\xd3\xf8\x66\xfa\xcb\x9f\xbb\xa5\xe9\xf8\xfe\x4b\x3e\x4f\xd7\x26\xb7\xf3\xe5\x38\x9f\xcf\xee\x9f\x16\xb3\xe5\xd3\x6f\xe3\x9b\xeb\xc5\xdb\x5b\x43\x7d\x3d\x90\x96\x4c\x61\xea\xa6\x42\xc2\x71\xc1\x7e\xe8\x21\x8d\x13\x44\xa7\x30\xd6\x80\x52\x11\x08\x2f\x06\xa6\x28\x5a\xeb\x40\xad\x41\x40\xeb\xd0\x42\x29\x1c\xd4\x46\xaa\xb5\xe2\xbc\x2e\x11\x94\xf6\x89\x80\xaf\xc4\xc1\x56\xda\xa1\x25\xa5\x37\x60\x2c\x48\xac\xd0\xff\x2e\x4a\x61\x45\x41\x68\xdd\x95\x37\xe2\x13\x41\xe9\xa2\x6a\x25\xd3\x6b\xd7\xf8\x03\x21\xf2\xcf\xb8\x5b\x19\x61\x25\x08\x2d\xa1\x11\x2e\x28\x30\x75\x2d\xb4\xf4\xc7\x19\xf1\x6c\x9a\x2f\x03\x5c\x70\x58\x61\xb1\xc7\xab\xab\xdd\xdb\xa0\x8b\xd2\x38\xd4\x20\x34\x88\x81\x37\xc0\xb5\x9b\x0d\x3a\x3e\x7b\xd5\xc1\x92\xaa\x10\xc4\xb8\x8c\x37\xc1\xa0\x92\x23\x3e\xd5\x15\x75\x79\x5b\x9b\x6d\xe0\x04\x9b\xfa\xd9\x01\xdb\x66\x52\x1b\xbf\xa8\xd9\x31\xa2\x69\xac\x69\xac\xf2\xec\x11\xab\xee\x16\x8b\xd9\xf5\x6c\xb2\x7c\x33\x4a\x33\x4d\x8a\x76\x9f\x94\x96\x21\x4a\xb3\x4f\x83\x28\xf1\xd7\xdd\xed\x34\xfe\x5a\x7c\x9e\x74\xbf\x26\xf7\xf9\xdd\x32\x7e\xcc\xc7\x37\xb3\xc5\xdd\x78\x32\xeb\x53\x7e\x8a\x4d\x65\x76\x35\x6a\xfa\x84\xbb\x83\xbc\x7f\xc6\xdd\x30\xa5\x0b\x8b\x5c\x6b\xc6\x74\xe3\x46\xf0\xa1\x32\x82\x78\x95\x2b\x62\x92\xe5\x5e\xad\x27\x9b\x57\xc7\x0e\x18\xf5\xcc\x3a\x09\xa1\xbe\x9d\xde\x9e\x71\x52\x58\xa5\xcd\xf9\x08\x6e\xc4\x33\x42\x3e\x05\x8b\x5f\x5b\x65\x51\x82\xd1\x05\xd7\x07\xef\x46\x07\x66\x8b\xde\x75\x75\x5b\x91\xba\x2c\xaa\xd6\x11\x5a\x70\x6d\xd3\x18\x4b\xec\xb7\xb8\x74\x16\xa0\x9f\x8f\x60\x12\x16\x3a\x8b\x71\xdf\x8d\xe0\xf7\xe1\xce\xe3\x1f\x8a\x66\x62\xb4\x46\x9f\x80\x47\xb8\xf6\x5b\x7b\x84\xaa\x2b\x0c\x67\x62\x50\x21\x46\x49\xbd\x60\x0d\xd7\x79\xb7\xc2\xe7\x3a\x59\xd7\x9f\x1a\x56\x9d\xf3\xfd\x71\xd7\x59\x1a\x66\xed\x99\xe7\x69\x27\x7d\x11\xb3\xf4\xce\xb8\x11\xe4\x9a\x2e\x22\x7f\x46\xef\x94\x8a\xf3\x74\xe3\x1e\x5d\x5b\xd1\xa1\x89\x0f\x0a\x2b\x79\x68\x67\xcd\x8b\xf1\x7a\x6f\xe6\xf7\x85\x2f\x63\x5d\x00\xc6\x76\xc3\xc2\x1c\xbe\xb7\xc5\x1f\xcf\x47\xc9\xce\xa2\xe7\xef\x63\xe6\x43\x1c\x9a\x68\xbd\xb1\x80\x5a\x36\x46\x69\x72\x17\x60\x71\x1d\x22\x29\x4d\xc1\x14\x87\xa2\x32\xad\x14\x8d\xba\x6a\xac\xf1\x3c\xaf\xd4\x16\x3f\x2b\x7c\x61\xcb\xd7\xf1\xf7\x0d\x92\x90\x82\x44\xc8\x9e\x4e\x62\x62\x34\xa1\x26\x17\x43\x7d\x72\x3e\x82\xeb\x83\x2d\x16\x0f\x2d\x97\xd5\x05\x44\xa9\xb2\xb0\xfb\x86\xaa\x45\xb2\x71\x12\xee\x14\x98\xcb\x24\x75\x9e\x86\x03\x1e\xb3\x81\x84\xd8\x41\x7f\x22\x33\x50\x9f\x8a\xf6\x1c\x3e\x0e\xae\x27\x34\x57\x6a\xe4\x11\xa5\x16\x44\x28\x63\xad\x57\x6e\x50\xf8\x5d\x8c\x73\x18\x14\xb8\xd0\xae\x10\x35\x34\xc2\x3a\x94\x5d\xfb\x4f\xcb\xa7\xe9\x6b\x6c\xa8\xaf\x62\xb5\x20\xd3\x40\x63\x9c\xe2\x38\xfa\x22\xdf\xdb\xcc\x87\xe9\xe4\xe5\xbf\x94\x48\x25\xda\x23\x0c\x8c\x4b\xc0\x56\x54\x4a\x5e\x00\xbe\x62\xd1\x92\x58\x55\xd8\xf5\x0e\xd6\xaa\xdc\xac\x5f\x1f\xc1\xdf\x8c\xa9\x50\xe8\xd0\x47\xaa\x6a\xd0\x0a\xc2\x58\x86\xa2\x28\xc1\xac\xbd\xa1\x08\xd2\x63\xe3\xdf\x7b\xd1\x11\xfc\xbe\x1c\x2e\x3c\xf6\x4e\x4d\x96\x07\xfe\x54\x5a\xe2\xeb\x40\x71\x68\x28\x54\xa2\xc3\x04\x83\xb0\xde\xf7\xd1\x64\xce\xa7\x3c\x59\x13\x2f\x84\xf6\xc7\xd7\x17\x83\xc3\x71\xac\xe4\x48\x89\x55\x34\xe8\x87\xb3\x9a\x0b\x1e\xdb\x8d\x5e\x19\x38\x8a\xed\xec\xbf\xc6\x6b\x42\xbb\xf0\xca\x87\x9e\x72\xc9\xc5\xdf\x23\xe2\x5b\x69\x75\xe0\x8a\x67\xa5\xe5\x7b\x25\xe1\x60\x0e\x8b\x9d\x87\x79\xe1\x4b\x54\xbf\x5a\x0b\x2a\x4a\x4e\x11\x89\xaf\xbe\x64\xe4\x9a\x1e\xfb\x76\x1a\x0b\xf0\x82\x04\xb5\x6e\xe0\x7e\x89\x6b\xc1\x09\xee\x88\xdb\xb1\x5a\xf3\xd0\x5e\xc6\xfc\x79\xd6\xe6\x45\xb3\x23\x3e\xff\xeb\x69\x91\x0e\x46\x7c\x34\x1e\x71\x50\xa2\xa8\xa8\xdc\xf1\xe9\x12\x85\xa5\x15\x0a\x0a\x01\xb3\x58\xa0\xda\xfa\x4e\x02\x16\x37\x6d\x25\x2c\x28\x4d\x68\xb7\xa2\x72\x7e\xa6\xa1\x32\xe4\x7d\xd7\x4e\x94\x03\x8b\xae\x31\x5a\x32\x08\x32\xbe\x16\xa2\x23\xb7\xc7\xf1\x71\x36\xbe\x5e\x7e\xfc\xed\x00\x47\x98\xca\x8d\x2f\x6b\xca\x15\xa1\xd1\x30\x4b\x43\x66\xfd\x7a\x7f\x37\x81\xa2\x6f\x3f\xb0\xb2\x28\x9e\xdd\x95\x57\x50\x9a\x06\x03\x8f\x05\xf5\x43\x4e\x07\xc8\xeb\x2d\x4c\x8d\xb0\x12\xc5\x33\x8f\x54\x4a\xa3\x87\x6e\xd1\xb5\x35\x27\x30\x44\x44\x01\xc9\x1e\xe8\x34\x5f\x4c\x6e\xe7\xf3\xd9\x64\x39\x9b\x1e\x7b\xcd\xbf\x60\xf8\x92\xf1\x71\x83\x43\x1f\xc4\xc1\xbf\xb1\xa6\x40\xe7\x98\x1e\x9d\xf8\x20\x1e\x77\xd3\xf1\x32\x9f\xff\xda\xab\xde\xaa\x7f\xab\x6e\xbe\xeb\xee\x1f\x9e\x5e\xbc\xc4\xaf\x31\x87\x9a\x40\xe8\x1d\x18\x4f\x97\x75\x6b\x03\x6d\x42\x56\x84\x37\x95\x03\xb1\x32\x6d\x70\xc4\x4b\xe4\x95\xa2\x61\x9c\x8d\x7d\x03\xcd\xf1\x4d\x23\x9c\x17\x7e\xae\xd8\x5d\x0c\x67\xb0\x11\x50\xad\x85\xaa\x30\x8c\xb6\x8a\xf1\xbd\xf0\xb5\x05\xac\x84\x3c\xf4\xa4\xbf\xea\xec\xe9\xc3\x38\xbf\x9e\x4d\x7b\x42\x7d\xf6\x06\x26\x46\xaf\xd5\xc6\xa7\x74\x23\x9c\xa3\xd2\x9a\x76\x53\xce\x34\xf3\x56\xee\xd9\xda\x1d\xe2\x66\x22\x94\x4e\xa8\x70\xf8\xce\xf1\xd6\x87\x0b\x35\x3a\x27\x36\x43\x0e\x5a\x14\x6e\x40\xbf\x4e\xfb\x9d\x91\xdf\xd1\xdb\xba\x1f\x52\x0c\x1c\xca\x00\xd8\x4f\x71\x29\x7a\x6e\x6f\xfd\xc5\xf6\x03\xde\xc1\x38\xdb\x99\x4d\x6a\x80\x6f\xe2\xc2\xd1\xc7\x8e\xb1\xc9\x70\xbb\x1d\xb8\x76\x94\x38\x7a\xbf\xfb\x19\xad\x4b\x4b\x50\x4c\xe0\x77\x37\xe6\x69\x1d\x6b\x2c\x12\xed\x26\x6f\xee\xc5\x13\x0f\xf9\xf4\xc0\x15\xd6\x54\x77\x95\xd0\xd8\x7b\xda\xd7\xba\xfe\x2b\x34\x7c\xdd\xd6\x73\x23\x31\x8c\x73\x71\x21\xd7\x8e\x6c\xcb\xcd\x1e\xe5\x70\xf3\xc0\x7f\xe9\x88\x1a\x3c\xd9\x8c\xa5\xb4\xe8\x92\xc8\x91\x79\x46\x7d\x3c\xff\x77\x4f\x6b\x7f\x70\xe2\x1f\x0e\x51\x71\x32\x9c\xc3\x5f\x25\x36\x16\xb9\xed\xcb\xb3\x2e\xe4\x3f\x45\x81\x50\x3b\x99\x0f\xf1\xe5\x01\x5b\x25\xa0\x79\x8d\x73\xcc\x4f\xe7\x19\xc0\x83\xe7\xd1\x30\x30\x67\xd1\x65\xec\xb1\x7c\x7a\x72\xf1\xbf\xd8\x70\xde\xff\x3a\xe9\x61\x26\xc3\xd0\xd1\x6c\x04\x30\xe5\xf7\x6b\x2a\x35\x18\xa5\x7a\x75\xbd\x3b\xf7\xf3\x7a\x7c\xed\xb7\x36\xf9\xdf\x01\x80\x2b\xc5\x9f\xfe\xf2\xcb\xb1\x0f\x93\xd1\x3d\x44\x80\xb0\xf6\x2d\x36\xee\x3c\x1e\xc9\x7a\xb1\x6d\x9a\x78\xfe\x7d\x51\x0a\xbd\xc1\xca\x6c\x92\xd8\xa9\x1a\x1d\x89\xba\x19\xe4\xfc\xb7\x2c\x3b\x85\xfb\xef\x4c\xc6\xde\xe4\xe1\x40\xfc\x9d\x7f\x9a\x1c\xbd\x13\x7f\xd0\x4c\x37\xfd\x7a\x33\x75\xb4\x39\x3a\x42\xe1\xff\x1d\xf3\x5a\x75\xd2\x43\x04\x5b\xe5\xfe\xbe\xb8\x9d\xff\x3f\x20\xd2\x69\xfd\x87\x6e\x0a\xdc\x75\x3a\x94\x69\x82\xfc\x90\xf1\x77\xee\x7f\xf0\x8e\x88\xd5\x21\xbd\xfa\xb7\xec\xbf\x01\x00\x00\xff\xff\x16\xb8\xa3\x63\xc3\x14\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 5315, mode: os.FileMode(436), modTime: time.Unix(1594793405, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
