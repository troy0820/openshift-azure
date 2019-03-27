// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x51\x6f\x13\x39\x10\x7e\xc6\xbf\x62\xd8\x14\x4a\x7b\x75\xb6\x2d\x3c\xa0\x40\x90\x4a\x29\x52\xa5\x1e\x45\x2d\x27\x1e\x00\x55\xce\x7a\x76\xd7\x64\xd7\xde\xb3\xc7\x49\x4b\xc9\x7f\x3f\xd9\x9b\xa4\x49\x9a\xe6\x2a\xc1\x43\xd8\xda\x9f\x67\xc6\xdf\xf7\xd9\x9e\xce\xd3\x74\xa0\x74\x3a\x10\xae\x04\x8e\xd7\x8c\xa9\x1c\x9e\x42\x61\xb1\x81\x74\x24\x6c\x5a\xa9\x41\x2a\x4d\x36\x44\x0b\x29\x52\x96\xe6\x8e\xc4\xe0\x0d\x50\x89\x9a\x01\xb8\x1b\x47\x58\x67\x54\x81\x23\xd3\x40\x0b\xec\x3a\xb4\x23\x95\x21\x03\xa8\x87\xb9\xeb\x5e\xe7\x0e\x78\x0e\xa9\xc4\x51\x2a\x95\x1b\xa6\xe2\x97\xb7\x98\x5a\x74\xc6\xdb\x0c\x79\x23\x2c\x1d\x30\x00\xcc\x4a\x03\xdb\x9b\x61\x70\xaf\x2a\x08\xe1\xa1\xb0\xcd\xbf\xde\x90\x00\xd8\x87\xfd\x6d\x78\xf7\xee\xae\xd8\x50\x86\xf1\x9a\x56\x57\x32\x00\x8b\x8e\x8c\xc5\xcc\x68\xe0\x17\x6b\xe6\x33\x41\xd0\x46\x6a\x87\x52\x29\xb0\x36\xba\xfb\xd3\x19\x0d\x6f\xdf\x6e\x9f\x9c\x7f\xdc\x66\xb7\x0c\x20\xa9\x4c\xc1\xa5\x55\x23\xb4\x49\x0f\x92\x9f\xc6\x5b\x2d\x2a\x99\xb0\x09\x3b\x39\xff\xb8\x42\x94\xb0\xb4\xca\x54\xae\x18\x9b\xee\xa7\xf1\x55\x05\xb7\xb7\xd0\x3d\x36\x3a\x57\x45\xf7\xb4\x16\x05\xba\xee\x27\x23\x11\x26\x13\x78\xfe\x2e\x12\xa4\x03\xea\xf9\x5a\xb5\x90\x32\xb9\x4e\xab\xb9\x16\xab\x0c\xbb\xcc\xa9\x83\xb4\xf2\x7a\x1f\x7e\xff\x06\xb2\x1e\x1f\x14\x63\x01\xba\x92\xb0\x95\x41\x62\x2e\x7c\x45\xee\x51\x32\x84\x75\x0f\x8b\x10\x67\x03\x2f\xb9\xb1\x20\x1d\x81\xd2\x40\x59\xb3\xf7\xfa\xd5\xab\x57\x6f\x40\x1a\xf6\xa4\xb1\x86\x4c\x7f\xeb\x56\x3a\x7a\xf6\x6c\x6f\x77\xc2\x9e\x34\xc6\x52\x3b\xd0\xe9\xec\xee\x4d\xd8\x13\xd5\x90\x18\x54\xe8\x80\x1f\xc1\xf9\xe5\xd5\xc7\xd3\x8b\x93\xaf\x47\x67\x67\x57\x47\x67\x67\xe7\x5f\x81\x37\xb0\x15\x83\x00\xaf\x83\x2e\x84\xc0\x79\xfb\xff\xa7\x93\xaf\x61\x70\x36\xcd\x65\x08\x0d\x5b\xf1\x97\xff\x84\xa3\xe3\xe3\x93\xcf\x5f\x98\x34\x1a\x19\x9b\x25\xe1\x4e\x8c\x70\xea\x17\x77\xe3\xb2\xa8\x5f\x3a\x9b\x65\xac\x03\xe3\x12\x75\x6b\x00\xa5\x0b\xd0\x41\xd2\xb1\x10\x05\x6a\x02\xa1\x25\x68\xa4\xb1\xb1\x43\xf0\xa4\x2a\x45\x0a\x1d\x14\x06\x1d\x28\x4d\x06\xac\xc8\x10\x32\xa3\xa5\x22\x65\x74\x97\x75\x40\xe5\xf3\xc5\xd6\x6b\x07\x03\xcc\x8d\x45\x90\xda\x81\x72\x30\xd4\x66\xac\x81\x4c\x50\x7f\x9a\x09\x01\xb5\x04\xdf\xc0\x58\x51\x09\x58\x37\x74\x03\x8e\xac\xd2\x05\x1b\x97\xaa\x42\xf8\xf6\x0d\xb6\x5e\x94\xc6\x91\x16\x35\x02\x97\x3b\xd0\xef\x43\x92\xc0\x8f\x1f\x81\x70\x70\x15\x62\x03\x07\xe1\x3b\x6c\xbb\x5d\xf3\x14\x36\xfb\xf6\x32\xec\xd6\x37\x30\x99\x44\xd1\x60\x16\xa5\xe5\xce\x21\xc1\x5f\xd7\x0c\xaf\x23\xb1\x97\x47\x97\xff\x5c\x9c\xf6\xb7\x17\xa2\xfc\x2d\x1c\xa1\x9d\x06\x69\xe7\x61\x32\xd9\x8e\x0b\xf9\xf5\xec\xcc\x58\xaf\x81\xf3\xc6\xaa\x91\xaa\xb0\x40\x09\x9c\xdb\x1a\x38\x9f\x11\x1a\xf6\x04\x7c\x04\x69\x2f\x0d\x9f\xbd\x5f\xc0\x71\x9a\x6d\x63\xc9\xcc\xeb\x90\xa8\x45\x32\xe6\x1b\x29\x08\x79\x26\x38\x59\xef\x88\x31\x17\x52\x29\xe0\x16\x21\x71\x9d\x17\xb0\x1b\xce\x33\xda\x1e\xec\x74\x77\x3b\xdf\x0f\x4a\xa2\xc6\xf5\xd2\xf4\x8e\xd4\x9d\x4e\xd2\x1e\x4d\x63\x55\xa1\x74\x5a\xc7\xed\xa5\xa6\x41\xed\x4a\x95\x13\x6f\x07\xba\x43\x3f\xc0\xd6\x3f\x7f\x9e\x23\x88\x1f\x7f\x16\xa3\xb2\xdb\x5b\x1e\x1c\xa4\x11\xb6\xba\xef\x45\x36\xf4\xcd\xfb\xca\x0c\x3e\x05\xe1\x93\x24\x6c\xbd\x32\x45\x81\x16\x38\x41\x5b\x13\x77\x2d\x2d\x5d\x57\x42\x32\xf7\x70\x38\xb7\x23\xb4\x37\x60\xf4\x82\x77\x76\x92\xe0\x76\x47\x41\x68\x28\x90\xa2\x09\x07\x31\x0b\x0b\xc2\x5c\xe4\xcb\xe7\x3c\xdd\x65\x84\x75\x13\xea\xf8\xa0\x6c\x7f\x79\x6e\xba\xae\x1e\x4a\x65\x61\x6b\x01\xc7\x36\xd7\x28\xcd\x58\x57\x46\xc8\x50\x66\x1b\x23\x79\xa4\x67\x4f\x28\x93\x2d\x27\x0f\xd8\x76\xc9\x76\xf7\x9d\xf6\x9d\x41\x74\xdb\x3d\xa1\x7b\xf7\x87\xd6\x81\xb3\xca\x78\xd9\x58\x33\x52\x12\x6d\xda\x4b\xaf\xa4\x20\x91\x5e\x19\x3f\x0f\xbd\x48\x43\x2f\x35\x3e\x58\x3a\x4c\xfd\xcf\x5e\xda\xe5\x83\xca\x0c\x82\x4a\xfd\x80\x5e\xd1\x7e\x86\x91\xe8\x48\x69\x11\xee\x9a\x7e\x88\x3f\x55\xa1\x2b\x07\x30\xe3\x75\x33\xfb\x2d\x7e\x0e\x46\x79\x67\x89\xc3\xd9\x95\xbf\x39\x42\x0b\x0a\xea\x39\x2d\x1a\x57\x1a\x7a\xac\x7e\xed\xa5\x11\x76\xfe\xe7\xfa\x05\x0b\xf6\xe6\x5f\xf3\xa9\x45\x87\xf6\x96\xff\x6a\x95\xe0\x08\x27\x5f\x8e\x3f\x1c\x7f\x39\xbb\x3a\xfa\x7c\xda\x4f\x5e\x26\x0f\x08\xb4\x54\x6c\xc4\x84\x28\xb1\x53\x98\x6e\x7b\x46\xd7\x92\xea\x0b\x82\xc4\x74\x3c\x78\x84\x87\x23\xb2\x7c\x7a\x34\x8e\xa7\x80\x78\xa9\x2f\x9c\xd1\xe9\xb0\xd2\x8a\x94\xa8\x78\x56\xf9\xe8\xc7\x64\x2a\xc5\x7e\xfc\xd7\x9f\xdd\x2f\x4b\xa3\xbd\xc3\x97\xaf\xf7\xf7\x16\x87\x0e\xd6\x02\x0f\xee\x03\x0f\xd7\x02\x0f\x23\x30\x59\x5f\x12\x27\x33\x44\x1d\x69\xe1\xb9\xb1\x3c\xb6\x22\x2b\x50\x21\x47\x68\x49\x39\xe4\x0d\xa2\xe5\xde\x56\x0e\xd6\x5c\x8d\x31\x0d\x63\xf5\xe8\x3e\x4b\xe9\xee\xca\x58\xbc\xa9\xec\xea\x4d\x15\xf8\x5c\xba\x84\x96\xda\x97\x95\xb8\x8f\x31\x38\xc6\xe7\x34\x89\x17\x72\x78\x9f\x27\x13\xc6\xc8\x6b\x94\x5c\xc8\x1a\x1a\x6b\xf2\x60\xf9\xbb\x17\x22\x33\x9a\xac\xa9\x78\x53\x89\xf0\x0a\x77\x40\x1b\xc2\x1e\x08\x32\xb5\xca\xf8\x1d\x2e\xbe\xf9\x99\x0d\x7d\x7d\x65\x4c\xe3\xc0\x6b\x52\xd5\xb4\x8e\xd0\x21\xf8\x86\xdd\x75\xa5\xa8\x43\x93\xb2\x3e\xca\xbc\x4b\x5d\x6d\x62\x37\xa2\x43\x77\xda\x01\xa9\x9c\x18\x54\xf1\xa5\x70\x37\xae\x32\x05\x38\xa5\xb3\xd8\x8c\xd4\x42\x8b\x02\x01\xc3\xf3\x41\x65\x80\x50\x69\x8d\x2f\x4a\x98\xb5\xcf\x0b\x09\xdb\x38\x38\x8b\xb2\xb6\x24\xd3\xdc\x9b\xfe\x2f\x00\x00\xff\xff\x29\xb1\x96\x00\xde\x0c\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4f\x8f\xdb\xb6\x13\xbd\xf3\x53\xbc\xf5\xfe\x10\xff\x8a\x82\x56\x72\xdd\x64\x17\x28\x8a\x04\xc8\xa5\x01\x76\x51\xf4\x10\xe4\x40\x4b\x23\x89\x31\x35\xa3\x90\x43\xdb\x1b\x43\xdf\xbd\xa0\x65\x37\xdb\xdd\xc0\xbd\x09\xe0\xfb\x33\xf3\xf8\xa8\xeb\xab\x6a\xed\xb9\x5a\xbb\xd4\xc3\xd2\xde\x18\xdf\xe2\x0a\x5d\xa4\x11\xd5\xd6\xc5\x2a\xf8\x75\xd5\x48\xbd\xa1\x88\x8a\xb4\xae\xda\xa4\x6e\xfd\x16\xda\x13\x1b\x20\x3d\x26\xa5\xa1\xd6\x80\xa4\x32\x62\x06\xae\x12\xc5\xad\xaf\xc9\x00\xc3\xa6\x4d\xab\x7d\x9b\x60\x5b\x54\x0d\x6d\xab\xc6\xa7\x4d\xe5\xbe\xe7\x48\x55\xa4\x24\x39\xd6\x64\x47\x17\xf5\x8d\x01\xa8\xee\x05\xcb\xcb\x30\xbc\x98\x0a\x45\x1e\x5d\x1c\xbf\x65\x51\x07\xbc\xc6\xeb\x25\xee\xee\x7e\x0c\x5b\xc6\x90\xcc\xfa\x9c\x69\x80\x48\x49\x25\x52\x2d\x0c\x7b\xff\xe2\xfc\x70\xb0\xf0\x2d\xe8\x1b\x56\xf7\x12\x08\x0b\xcf\x6d\x74\x0b\x4c\x93\x01\x6a\xa7\x98\x4d\x66\x74\xd5\x38\x1a\x84\x57\x5f\x93\x30\xde\xbd\x5b\xbe\xff\xf4\x61\x69\x0e\x06\x58\x04\xe9\x6c\x13\xfd\x96\xe2\xe2\x06\x8b\xaf\x92\x23\xbb\xd0\x2c\xcc\x64\xde\x7f\xfa\x70\x34\x21\x6e\x66\xd1\xa7\x71\xba\xa8\xcf\xf3\x6c\xbd\x31\xa7\xad\xc7\x1c\x02\x0e\x07\xac\x7e\x17\x6e\x7d\xb7\xfa\x38\xb8\x8e\xd2\xea\x0f\x69\x08\xd3\x84\x57\x77\xc7\x18\xb9\xa0\x5e\x19\x73\x8d\x5d\x4f\x3c\x8b\x7a\xee\xc0\x05\xb6\x73\xae\x23\x56\x38\x6e\xc0\xa4\x3b\x89\x1b\x64\xf5\xc1\xab\xa7\x84\x4e\x28\xc1\xb3\x0a\xa2\xab\x09\xb5\x70\xe3\xd5\x0b\xaf\xcc\x75\x09\xe5\x4c\x8e\x99\x13\xd6\xd4\x4a\x24\x34\x9c\xe0\x13\x36\x2c\x3b\x86\x4a\xe9\xc8\xc9\x89\x8e\x2b\xe6\x11\x3b\xaf\x3d\x68\x18\xf5\x11\x49\xa3\xe7\xce\xec\x7a\x1f\x08\x9f\x3f\xe3\x7f\xff\xef\x25\x29\xbb\x81\x60\x9b\x5f\x70\x7b\x8b\xc5\x02\x5f\xbe\xbc\x45\x23\x48\x81\x68\xc4\x9b\xf2\xcd\x64\x4e\x9c\x2b\x5c\xce\xe2\xa1\x6c\x9b\x47\x4c\x53\xe1\x95\x74\x67\x15\x73\x14\x49\xa4\xf8\x75\x6f\x68\x3f\x4a\x54\x3c\xfc\xf6\xf0\xe7\xfd\xc7\xdb\xe5\x13\x95\xbf\x24\x6e\x28\x9e\x44\xe6\x73\x4c\xd3\xf2\x48\xb4\xfb\xf3\x3d\xc4\xcc\xb0\x76\x8c\x7e\xeb\x03\x75\xd4\xc0\xda\x38\xc0\xda\x73\xa0\x65\x27\xd8\x2d\xaa\x9b\xaa\x7c\xde\x7c\x87\xa5\x93\xdb\xc5\x91\x4d\xe6\x62\x34\x23\x8d\xc9\x63\xe3\x94\x6c\xed\xac\xc6\x9c\xd4\x5c\xea\xa6\x66\xa6\xc6\xba\x66\xc0\x18\xa5\x2d\x49\xc9\x48\x9c\x7a\xdf\xaa\xad\x85\x35\x4a\xb0\x63\x70\x4c\x73\xf7\x42\xa2\xff\x62\x95\x4b\x7c\x5a\x54\x73\x0d\x16\xa5\x1b\x38\x95\xc1\xd7\xf6\xdf\x48\xd4\xb1\xfc\x4c\x82\xc8\x98\x90\x59\x7d\xc0\xe0\x92\x52\x2c\xe5\xc8\xa3\xf9\x51\x72\x62\xb7\x0e\xf4\x73\x95\x7f\x4a\xff\xfc\x4d\x5c\x44\xcf\x65\x6f\x7c\x72\xeb\x50\x8a\x1e\xd3\x63\x0a\xd2\x21\x79\xae\x8f\x3d\x1c\x1c\xbb\x8e\x40\x5b\x8a\x8f\xda\x17\x88\xf6\x51\x72\xd7\xe3\xfc\x30\x9f\x18\xce\x3a\x74\x56\xf9\xe9\x48\x32\xbe\x38\xfe\x3b\x00\x00\xff\xff\x6b\x17\x8a\x82\x53\x05\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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