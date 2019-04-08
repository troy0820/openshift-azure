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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xdf\x4f\x1b\x39\x10\x7e\xae\xff\x8a\xe9\x86\x96\xc2\xe1\x2c\xd0\x3e\x54\x69\x53\x89\x52\x2a\x21\x71\xa5\x82\x9e\xfa\xd0\x56\xc8\x59\x4f\x36\x6e\x76\x6d\x9f\x3d\x4e\xa0\x34\xff\xfb\xc9\xde\x0d\x24\x21\x70\x48\xe5\x21\x04\xcf\x78\x7e\x7c\xdf\x67\x7b\xe8\x3c\xcd\x07\x4a\xe7\x03\xe1\x47\xc0\xf1\x92\x31\x35\x84\xa7\x50\x3a\xb4\x90\x4f\x84\xcb\x2b\x35\xc8\xa5\x29\xc6\xe8\x20\x47\x2a\xf2\xa1\x27\x31\x78\x03\x34\x42\xcd\x00\xfc\x95\x27\xac\x0b\xaa\xc0\x93\xb1\xd0\x38\x76\x3d\xba\x89\x2a\x90\x01\xd4\xe3\xa1\xef\x5e\x0e\x3d\xf0\x21\xe4\x12\x27\xb9\x54\x7e\x9c\x8b\x5f\xc1\x61\xee\xd0\x9b\xe0\x0a\xe4\x56\x38\xda\x63\x00\x58\x8c\x0c\x6c\x3e\xec\x06\x77\xaa\x82\x18\x1e\x4a\x67\xff\x0d\x86\x04\xc0\x2e\xec\x6e\xc2\xbb\x77\xb7\xc5\xc6\x32\x4c\xd0\xb4\xba\x93\x01\x38\xf4\x64\x1c\x16\x46\x03\x3f\x5b\x63\x2f\x04\x41\x13\xa9\x59\xca\xa5\xc0\xda\xe8\xee\x4f\x6f\x34\xbc\x7d\xbb\x79\x74\xfa\x71\x93\x5d\x33\x80\xac\x32\x25\x97\x4e\x4d\xd0\x65\x3d\xc8\x7e\x9a\xe0\xb4\xa8\x64\xc6\x66\xec\xe8\xf4\xe3\x0a\x50\xc2\xd1\x2a\x52\x43\xc5\x58\xdb\x8f\x0d\x55\x05\xd7\xd7\xd0\x3d\x34\x7a\xa8\xca\xee\x71\x2d\x4a\xf4\xdd\x4f\x46\x22\xcc\x66\xf0\xfc\x5d\x02\x48\x47\xaf\xe7\x6b\xd9\x42\x2a\xe4\x3a\xae\x6e\xb8\x58\x45\xd8\x17\x5e\xed\xe5\x55\xd0\xbb\xf0\xfb\x37\x90\x0b\x78\x2f\x19\x0b\xae\x2b\x09\x1b\x1a\x24\x0e\x45\xa8\xc8\x3f\x8a\x86\xb8\xef\x7e\x12\x92\x35\xe2\x32\x34\x0e\xa4\x27\x50\x1a\xa8\xb0\x3b\xaf\x5f\xbd\x7a\xf5\x06\xa4\x61\x4f\xac\x33\x64\xfa\x1b\xd7\xd2\xd3\xb3\x67\x3b\xdb\x33\xf6\xc4\x1a\x47\xcd\x42\xa7\xb3\xbd\x33\x63\x4f\x94\x25\x31\xa8\xd0\x03\x3f\x80\xd3\xf3\x8b\x8f\xc7\x67\x47\x5f\x0f\x4e\x4e\x2e\x0e\x4e\x4e\x4e\xbf\x02\xb7\xb0\x91\x82\x00\xaf\x23\x2f\x84\xc0\x79\xf3\xfb\xd3\xd1\xd7\xb8\x38\x37\x73\x19\x43\xc3\x46\xfa\xe4\x3f\xe1\xe0\xf0\xf0\xe8\xf3\x17\xe0\x53\x26\x8d\x46\xc6\xe6\x79\xb8\x17\x13\x6c\x25\xe3\xaf\x7c\x91\x28\xcc\xe7\x56\xc6\x3a\x30\x1d\xa1\x6e\x34\xa0\x74\x09\x3a\xb2\x3a\x15\xa2\x44\x4d\x20\xb4\x04\x8d\x34\x35\x6e\x0c\x81\x54\xa5\x48\xa1\x87\xd2\xa0\x07\xa5\xc9\x80\x13\x05\x42\x61\xb4\x54\xa4\x8c\xee\xb2\x0e\xa8\xe1\xcd\x66\x17\xb4\x87\x01\x0e\x8d\x43\x90\xda\x83\xf2\x30\xd6\x66\xaa\x81\x4c\x14\x40\x9b\x09\x01\xb5\x84\x60\x61\xaa\x68\x04\x58\x5b\xba\x02\x4f\x4e\xe9\x92\x4d\x47\xaa\x42\xf8\xf6\x0d\x36\x5e\x8c\x8c\x27\x2d\x6a\x04\x2e\xb7\xa0\xdf\x87\x2c\x83\x1f\x3f\x22\xe6\xe0\x2b\x44\x0b\x7b\xf1\x7b\x6c\xbb\xd9\xf3\x14\x1e\x96\xee\x79\xec\x36\x58\x98\xcd\x12\x6f\x30\x8f\xd2\x60\xe7\x91\xe0\xaf\x4b\x86\x97\x09\xdb\xf3\x83\xf3\x7f\xce\x8e\xfb\x9b\x0b\x51\xfe\x16\x9e\xd0\xb5\x41\x1a\x3b\xcc\x66\x9b\x69\x23\xbf\x9c\x1f\x1b\x17\x34\x70\x6e\x9d\x9a\xa8\x0a\x4b\x94\xc0\xb9\xab\x81\xf3\x39\xa0\xb1\x27\xe0\x13\xc8\x7b\x79\xfc\xda\xfb\x05\x1c\xdb\x6c\xc0\x39\x6a\x72\x57\xd6\x28\x4d\x0d\x37\xc1\x3e\xd8\x07\x0b\x3a\x66\x6f\xb6\x33\x16\xac\x14\x84\xbc\x10\x9c\x5c\xf0\xc4\x98\x8f\xf9\x15\x70\x87\x90\xf9\xce\x0b\xd8\x8e\xe7\x1c\x5d\x0f\xb6\xba\xdb\x9d\xef\x7b\x23\x22\xeb\x7b\x79\x7e\x8b\xf4\x56\x27\x6b\x8e\xac\x71\xaa\x54\x3a\xaf\x53\xcf\xb9\xb1\xa8\xfd\x48\x0d\x89\x37\x0b\xdd\x71\x18\x60\x23\xaa\x3f\xcf\x11\x15\x91\x3e\x16\xa3\xb2\xeb\x6b\x1e\x65\xa5\x11\x36\xba\xef\x45\x31\x0e\xf6\x7d\x65\x06\x9f\xa2\x1a\xb2\x2c\xb6\x5e\x99\xb2\x44\x07\x9c\xa0\xa9\x89\xb7\x80\x75\xfd\x08\xb2\x1b\x61\xc7\xf3\x3c\x41\x77\x05\x46\x2f\x08\x6a\x2b\x8b\x47\xc0\x53\x64\x1f\x4a\xa4\xa4\xcc\x41\xca\xc2\x22\x5b\x67\xc3\xe5\xf3\x9f\x6f\x33\xc2\xda\xc6\x3a\x3e\x28\xd7\x5f\xb6\xb5\xfb\xea\xb1\x54\x0e\x36\x16\xfc\xd8\xc3\x35\x4a\x33\xd5\x95\x11\x32\x96\xd9\xc4\xc8\x1e\x29\xe4\x23\x2a\x64\x83\xc9\x3d\x5a\x5e\xd2\xe2\x5d\xf9\x7d\x67\x90\x24\x78\x87\xe8\xde\xdd\xa5\x75\xce\x45\x65\x82\xb4\xce\x4c\x94\x44\x97\xf7\xf2\x0b\x29\x48\xe4\x17\x26\xdc\x84\x5e\x84\xa1\x97\x9b\x10\x75\x9e\x4c\x8b\x0a\x8f\xe8\x35\x8d\x27\xdb\xff\xf4\xd9\xec\x1f\x54\x66\x10\x19\xec\x47\xef\x15\x5d\xcc\x7d\x24\x7a\x52\x5a\xc4\xcb\xa9\x1f\x73\xb7\x0c\x75\xe5\x00\xe6\x98\x3f\xcc\x4c\x5b\xd4\xdc\x19\xe5\xad\x5c\xf6\xe7\xcf\xc4\xc3\x11\x1a\xa7\xc8\xac\xd7\xc2\xfa\x91\xa1\xc7\x72\xdb\xdc\x32\xb1\xf3\x3f\xe7\x36\x02\xdc\xbb\xf9\x76\x63\x5a\x54\x6f\x6f\xf9\xaf\x96\x25\x84\xa3\x2f\x87\x1f\x0e\xbf\x9c\x5c\x1c\x7c\x3e\xee\x67\x2f\xb3\x7b\x08\x5a\x2a\x36\xf9\xc4\x28\x69\xba\x68\xdb\x9e\xc3\xb5\xa4\x88\x05\x42\x1a\x51\x44\xfd\xf0\x78\x7c\x96\x4f\x96\xc6\x69\xeb\x90\x5e\x81\x85\xf3\xdb\x2e\x2b\xad\x48\x89\x8a\x17\x55\x48\x5a\xcd\x5a\x2a\x76\xd3\x4f\x7f\x7e\xf7\x2c\xad\xf6\xf6\x5f\xbe\xde\xdd\x59\x5c\xda\x5b\xeb\xb8\x77\xd7\x71\x7f\xad\xe3\x7e\x72\xcc\xd6\x97\xc4\xc9\x8c\x51\x27\x58\xf8\xd0\x38\x9e\xc6\x97\x15\x57\x21\x27\xe8\x48\x79\xe4\x16\xd1\xf1\xe0\x2a\x0f\x6b\xae\xcd\x94\x86\xb1\x7a\x72\x17\xa5\x7c\x7b\x65\x2d\xdd\x62\x6e\xf5\x16\x8b\x78\x2e\x5d\x50\x4b\x23\xcf\x4a\xdc\xc7\x08\x1c\xd3\xfb\x9b\xa5\xcb\x3a\x3e\xe8\xb3\x19\x63\x14\x34\x4a\x2e\x64\x0d\xd6\x99\x61\x94\xfc\xed\xeb\x51\x18\x4d\xce\x54\xdc\x56\x22\x3e\xdb\x1d\xd0\x86\xb0\x07\x82\x4c\xad\x0a\x7e\xeb\x97\x86\x84\xc2\xc5\xff\x05\x2a\x63\xac\x87\xa0\x49\x55\x6d\x1d\x71\xa4\x08\x96\xdd\x4e\xb2\xa8\xe3\x54\xb3\x3e\xca\xcd\x64\xbb\x3a\xf8\x3e\xe8\x1d\x27\xda\x0e\x48\xe5\xc5\xa0\x4a\xaf\x88\xbf\xf2\x95\x29\xc1\x2b\x5d\xa4\xe9\xa5\x16\x5a\x94\x08\x18\x9f\x16\x1a\x45\x17\x1a\x39\x13\xca\x11\xcc\x47\xee\x85\x84\x4d\x1c\x9c\x47\x59\x5b\x92\xb1\x77\xcc\xff\x05\x00\x00\xff\xff\x56\xa6\x26\x7f\x12\x0d\x00\x00")

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

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4d\x6f\xdb\x46\x10\xbd\xef\xaf\x78\x96\x8b\xa8\x45\xb1\x62\x72\x75\x62\x03\x45\x91\x00\xb9\x34\x80\x8d\xa2\x87\x20\x87\x15\x39\x24\x37\x22\x67\x36\xbb\xb3\xfa\x88\xa0\xff\x5e\xac\x28\x35\xae\x6d\x28\x37\x02\xfb\xde\x9b\x99\x37\x6f\x78\x7d\x55\x2d\x3d\x57\x4b\x97\x7a\x58\xda\x1a\xe3\x5b\x5c\xa1\x8b\x14\x50\xad\x5d\xac\x06\xbf\xac\x1a\xa9\x57\x14\x51\x91\xd6\x55\x9b\xd4\x2d\xdf\x42\x7b\x62\x03\xa4\x5d\x52\x1a\x6b\x1d\x90\x54\x02\x26\xe0\x22\x51\x5c\xfb\x9a\x0c\x30\xae\xda\xb4\xd8\xb6\x09\xb6\x45\xd5\xd0\xba\x6a\x7c\x5a\x55\xee\x7b\x8e\x54\x45\x4a\x92\x63\x4d\x36\xb8\xa8\x6f\x0c\x40\x75\x2f\x98\x5f\x86\xe1\x59\x57\x28\xf2\xe8\x62\xf8\x96\x45\x1d\xf0\x1a\xaf\xe7\xb8\xbb\xfb\xd1\x6c\x69\x43\x32\xeb\x53\xa6\x01\x22\x25\x95\x48\xb5\x30\xec\xfd\xb3\xf7\xfd\xde\xc2\xb7\xa0\x6f\x58\xdc\xcb\x40\x98\x79\x6e\xa3\x9b\xe1\x70\x30\x40\xed\x14\x53\x91\x09\x5d\x35\x8e\x46\xe1\xc5\xd7\x24\x8c\x77\xef\xe6\xef\x3f\x7d\x98\x9b\xbd\x01\x66\x83\x74\xb6\x89\x7e\x4d\x71\x76\x83\xd9\x57\xc9\x91\xdd\xd0\xcc\xcc\xc1\xbc\xff\xf4\xe1\x58\x84\xb8\x99\x44\x1f\xdb\xe9\xa2\x3e\xf5\xb3\xf5\xc6\x9c\xa6\x0e\x79\x18\xb0\xdf\x63\xf1\xa7\x70\xeb\xbb\xc5\xc7\xd1\x75\x94\x16\x7f\x49\x43\x38\x1c\xf0\xea\xee\x68\x23\x17\xd4\x2b\x63\xae\xb1\xe9\x89\x27\x51\xcf\x1d\xb8\xc0\x36\xce\x75\xc4\x0a\xc7\x0d\x98\x74\x23\x71\x85\xac\x7e\xf0\xea\x29\xa1\x13\x4a\xf0\xac\x82\xe8\x6a\x42\x2d\xdc\x78\xf5\xc2\x0b\x73\x5d\x4c\x39\x93\x63\xe6\x84\x25\xb5\x12\x09\x0d\x27\xf8\x84\x15\xcb\x86\xa1\x52\x32\x72\xaa\x44\xc7\x11\x73\xc0\xc6\x6b\x0f\x1a\x83\xee\x90\x34\x7a\xee\xcc\xa6\xf7\x03\xe1\xf3\x67\xfc\xf2\x6b\x2f\x49\xd9\x8d\x04\xdb\xfc\x86\xdb\x5b\xcc\x66\xf8\xf2\xe5\x2d\x1a\x41\x1a\x88\x02\xde\x94\x6f\x26\x73\xe2\x5c\xe1\xb2\x17\x0f\x65\xda\x1c\x70\x38\x14\x5e\x71\x77\x52\x31\x47\x91\x44\x8a\xdf\xb7\x86\xb6\x41\xa2\xe2\xe1\x8f\x87\xbf\xef\x3f\xde\xce\x1f\xa9\xfc\x23\x71\x45\xf1\x24\x32\xbd\xe3\x70\x98\x1f\x89\x76\x7b\xde\x43\xcc\x0c\x6b\x43\xf4\x6b\x3f\x50\x47\x0d\xac\x8d\x23\xac\x3d\x1b\x5a\x66\x82\x5d\xa3\xba\xa9\xca\xe7\xcd\x77\x58\x3a\x55\x83\xb5\xc4\x1a\x77\x41\x3c\xeb\xb4\x9b\x1c\x2e\xce\x61\x32\x97\xea\x13\xdd\x98\x1c\x1a\xa7\x64\x6b\x67\x35\xe6\xa4\xe6\x52\x60\x35\x33\x35\xd6\x35\x23\x42\x94\xb6\xd8\x27\x81\x38\xf5\xbe\x55\x5b\x0b\x6b\x94\xc1\x86\xc1\x31\x4d\x81\x1c\x12\xfd\x8c\x55\x36\xfb\x38\xbd\xe6\x1a\x2c\x4a\x37\x70\x2a\xa3\xaf\xed\xff\x91\xa8\x63\xf9\xc3\x0c\x22\x21\x21\xb3\xfa\x01\xa3\x4b\x4a\xb1\x24\x26\x07\xf3\x23\xf9\xc4\x6e\x39\xd0\xcb\x2a\xff\x5d\xc2\xd3\x43\xb9\x88\x9e\x2e\xa0\xf1\xc9\x2d\x87\x92\xfe\x98\x76\x69\x90\x0e\xc9\x73\x7d\x0c\xe7\xe8\xd8\x75\x04\x5a\x53\xdc\x69\x5f\x20\xda\x47\xc9\x5d\x8f\xf3\xb5\x3e\x2a\x38\xe9\xd0\x59\xe5\xc5\x96\x24\x3c\x7b\xfe\x37\x00\x00\xff\xff\xdc\x5b\x5f\x47\x68\x05\x00\x00")

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
