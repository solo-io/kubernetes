// Code generated by go-bindata.
// sources:
// instance-listener/Godeps/Godeps.json
// instance-listener/Godeps/Readme
// instance-listener/main.go
// DO NOT EDIT!

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

var _instanceListenerGodepsGodepsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcc\xbd\x0a\xc2\x30\x14\x40\xe1\x39\x79\x0a\xb9\xb3\xb9\x41\x84\x0e\xce\x82\xb8\x75\x72\x91\x0e\x31\xb9\xb4\xa1\xe6\x87\x24\xcd\x22\xbe\xbb\x6d\x70\xe8\xfa\x1d\x38\x1f\xce\xe0\xee\x62\x48\xa5\x57\x65\x82\xcb\x01\x46\x5b\xa6\xe5\x85\x3a\x38\x49\x4e\x0b\x65\xaa\xf2\x9a\x8c\x30\x54\xe5\xe2\xed\x2c\xad\xcf\x65\x23\xf1\xb6\xb9\x90\xa7\x04\xc7\x75\x72\x0b\x0f\x4a\xd9\x06\xdf\x1e\xe1\x84\xdd\x9f\x0d\xc5\x5d\xa9\xdd\xb9\x79\xaf\xf4\xac\x46\xca\xab\x3d\x39\x63\x80\x12\x11\x81\xb3\x61\x8b\x57\x8a\x2d\x0c\xfc\xcb\x7f\x01\x00\x00\xff\xff\x2a\x8c\x59\x34\xa1\x00\x00\x00")

func instanceListenerGodepsGodepsJsonBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerGodepsGodepsJson,
		"instance-listener/Godeps/Godeps.json",
	)
}

func instanceListenerGodepsGodepsJson() (*asset, error) {
	bytes, err := instanceListenerGodepsGodepsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/Godeps/Godeps.json", size: 161, mode: os.FileMode(420), modTime: time.Unix(1465354508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _instanceListenerGodepsReadme = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\x8d\xb1\x0d\x03\x21\x10\x04\xf3\xaf\xe2\x2a\x80\xdc\x55\x58\xb2\x1b\xe0\xe1\x0c\x48\xc0\xbe\x8e\xfd\x80\xee\x8d\x9d\xad\x46\x9a\x9d\x77\xa9\x53\x52\x35\x8d\x84\x2d\xa1\xa9\xca\x26\x59\x87\x5a\xa0\x26\x09\x37\xd1\x03\x6b\x0c\xad\x2d\x39\x97\x64\x24\xbd\xdc\x71\x3c\x9b\x86\xa9\x92\x20\x03\x14\x4d\x95\x1b\xbe\xb6\x5e\xc8\x6b\x3e\xbc\xcf\x95\xe5\x3e\x5d\x44\xf7\x04\xda\xf4\x7f\x53\x3e\x30\xe9\xb0\x9d\x19\x7b\xfe\xae\x31\xdc\xf1\x0d\x00\x00\xff\xff\x52\x08\x9f\x0e\x88\x00\x00\x00")

func instanceListenerGodepsReadmeBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerGodepsReadme,
		"instance-listener/Godeps/Readme",
	)
}

func instanceListenerGodepsReadme() (*asset, error) {
	bytes, err := instanceListenerGodepsReadmeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/Godeps/Readme", size: 136, mode: os.FileMode(420), modTime: time.Unix(1465354508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _instanceListenerMainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\x6d\x6f\xdb\x38\x12\xfe\x6c\xff\x0a\x56\xc0\x16\x72\xe2\xca\x59\x5c\x76\x17\x9b\x6e\x3e\xa4\xdb\x97\x35\xd0\xb4\x3e\xa7\xc5\x7e\x48\x7b\x59\x46\xa2\x13\x5d\x64\x51\x15\x29\xb7\x41\x2f\xff\xfd\x9e\x19\x52\x6f\x91\x9d\x2d\xee\x0d\x57\x14\x91\xcc\x19\x0e\x67\x1e\xce\x0c\x1f\xaa\x90\xf1\x8d\xbc\x52\x62\x2d\xd3\x7c\x3c\x4e\xd7\x85\x2e\xad\x08\xc7\xa3\x40\xe5\xb1\x4e\xd2\xfc\x6a\xf6\x77\xa3\xf3\x80\x06\xca\x52\x97\x86\xde\x56\x99\xbc\xe2\xe7\xda\xd2\x23\xd5\xb3\x54\x57\x36\xcd\xe8\x47\xa6\x59\x94\x2b\xeb\x1f\xb3\x6b\x6b\x0b\x7a\xd7\x3c\xd7\xd8\x12\x46\xdd\xeb\x6d\x1e\xd3\xd3\xa6\x6b\x15\x8c\x27\xe3\x71\xac\x73\x63\x85\xb1\xd2\xaa\x55\x9a\x29\x71\x2c\x82\x59\x22\xad\x9c\x35\x43\x91\x73\x66\x6c\x6f\x0b\xe5\x14\xf1\xb7\xac\x62\x2b\xbe\x8e\x47\xa7\x32\x9e\x17\xa7\xb2\x10\x08\xa6\x38\x77\x0b\x7d\x74\x0f\xd1\xf9\xf7\x07\xd9\x38\x0a\xe6\x85\x09\xfe\xe0\x49\x2f\xf2\x0d\xcd\xea\x4c\x1a\xce\xf7\x93\xa0\x4a\xb3\xee\xc6\xe3\x55\x95\xc7\x0c\x5a\x38\xa1\xb5\x65\x79\x65\xc4\xd1\xb1\xd0\x26\x3a\xc1\xeb\x78\xb4\xd2\xa5\x48\xa7\x02\xe3\x34\x5c\xca\x1c\x18\xb3\x12\x94\x47\x00\x29\x5a\xc0\xae\x5d\x85\x01\x69\x7c\xb7\x39\x12\xdf\x99\x60\xea\x67\x4c\xc6\xa3\xbb\xf1\x88\x22\x5f\x94\x08\xfb\x0b\x59\x20\xcc\xa3\x33\x76\x26\x0c\x0a\x1e\x86\x7e\x50\xe5\xe9\xcd\x05\xbd\xb8\x21\x41\xcb\xd2\x44\x61\x54\x6e\xc5\x26\x95\xa2\x4a\x0a\x11\xb2\x37\x09\x86\xd2\x55\x1a\x4b\x9b\xea\x5c\x14\x55\x59\x68\xa3\x4c\x80\xd5\x54\x2e\x2f\x33\xb5\x50\xa5\x49\x8d\xc5\xc6\xab\x66\xc5\x67\x5a\x67\x61\x30\x90\x63\x45\xc0\xae\xb0\xae\x34\xa6\x5a\x23\x34\x51\xd4\x52\x2c\xab\x33\x1a\x4b\x8d\x58\xeb\x2a\xb7\x2a\x11\x56\x0b\xde\x49\x5a\x8c\xed\x2e\x64\x69\x54\x38\xd9\x8e\x13\x6b\x10\x8c\x8c\xed\x3d\xbc\x48\xb8\x13\xb3\x74\x25\xf6\x3a\xb8\x1d\x1f\xd7\x08\x0d\x60\x7f\xb1\x5c\xbe\x5d\x1e\x89\x75\x85\x8c\x2b\x4a\xbd\x01\x3a\xe2\x89\xc7\x15\xb6\x46\xa5\xb2\x55\x99\xef\x32\xba\xd3\x9e\x37\x21\x62\x99\xe7\xda\x8a\x4b\x25\x3e\x04\x1f\x82\x2d\x16\x1f\xed\x0d\x41\x27\x9b\x48\xa0\xd3\x9b\x24\x2d\x4f\x32\xe0\xee\x30\x9b\x8a\x83\x9f\x7e\xf8\xc1\x07\x48\x09\xfe\x5a\xc7\x37\x84\x16\x15\x50\xb4\xfc\xfd\xb4\xb2\xea\xcb\xd7\x3b\xda\xc6\xcd\x03\x52\x23\x37\xaa\x27\x6b\x24\x1b\x59\x0a\xe3\xca\x09\x6a\x51\x53\x47\xc7\x48\xf0\x1b\x15\x0e\xaa\x61\xe2\xb5\x7c\xe1\x0c\xd5\xb6\xcc\x70\xf9\x3c\x15\xe8\x22\xe4\x80\x6b\x18\xd1\x52\xc9\xe4\x25\xea\x3a\x6c\x2a\x7c\xc2\xe8\x90\xd6\xa3\x63\x91\xa7\xd9\x00\xe8\x58\x57\x59\x22\x08\xdc\x12\x93\xdb\x6e\x31\x85\x17\xb7\xc0\xdb\x5e\x23\xef\xf0\x7f\x95\x96\xd8\xda\x4b\xad\xed\x91\x08\xc4\x3e\x99\x8c\x5e\x50\x0b\x0b\x27\x04\xa5\x50\x99\x71\x88\xfb\xe5\xe0\x14\xd5\x78\xf4\x3e\x5f\x23\x37\xaf\x65\x16\x3a\x87\x1f\x9b\xc9\xd3\xfb\xfe\xf4\x13\x52\x62\x75\x4e\xf1\x82\x92\xda\xb7\x25\xee\x17\x5b\x16\xa6\x4d\x44\xfb\x18\x65\xbc\xe9\xaa\x9c\x17\x0d\x26\x57\xca\x62\x7f\x64\x36\x2f\xc2\x3f\x45\xc1\xa7\x5b\xbb\x36\x26\x8b\x8c\x66\x8b\xb4\x40\x61\x6c\x02\x36\xdb\xcf\xbb\x71\xcf\xc4\x99\x95\xa5\xa5\xde\x46\x15\x22\x92\xd4\xc4\x7a\xa3\xca\x5b\x11\x52\xc7\xb8\x56\x10\x5e\x2a\x09\x04\x4b\x2d\x93\x58\x1a\x3b\x11\x9f\x53\x7b\x0d\xf3\xae\xec\xda\x00\xea\xb6\x44\xe1\xa5\xf9\x4a\x53\x2c\xe7\x1f\x2f\x6f\xad\x0a\xbb\x75\xb3\x2f\x82\x23\xc2\x63\xfb\xc4\x76\xf4\x54\x1a\xce\xd1\x8e\xde\x73\xb5\x92\x55\x66\x49\x42\xd0\x3c\x5b\xbe\x3d\x79\xfe\xeb\xc9\xd9\xbb\x8b\xf9\x62\x73\xc8\x7d\x43\xc1\x75\xa3\x58\xa1\x8b\x6c\xdf\xaa\x43\xf5\xde\x4a\xc7\x0f\x02\x5c\x2b\x3f\x71\xa8\x3e\xed\x0e\x3c\x59\x63\xbe\x1b\x6d\x93\x12\x3b\x10\x57\x19\x65\x40\x03\x9c\x90\x49\x52\x2a\xd3\x07\x6d\x8b\x6f\xdd\x0e\x61\x50\xa8\xca\x36\xa9\x81\x73\x34\x7a\x9e\xca\xec\xfd\xf3\x45\x18\x60\x7f\x0e\x61\x0b\x5e\x23\x3b\x49\x82\xd1\x13\x2c\x41\x21\xcc\x17\x47\x38\xe2\xfa\x00\x4d\x31\xbe\xc0\xb1\x7e\x24\x7e\xfe\xf9\xc7\x9f\xf0\xeb\xee\xcf\xd2\x0b\x07\x7b\x74\x56\xf4\x91\x68\xc2\x69\xa0\x68\x93\x2f\x81\x6b\x7c\xd0\xb4\x31\xe3\x2c\xcf\x55\x4c\x27\x0d\x5c\xed\xfb\x33\xb9\x17\xeb\x95\x16\x74\x9a\xba\x83\xb4\xb7\x09\x8d\x39\x64\x49\x14\x45\xdc\x46\xe9\xc0\xe0\x1a\xbc\x70\xe8\xa0\x99\x31\x56\xd1\xef\x65\x8a\x8c\xa3\x04\x24\xb5\x2d\x01\x6e\xdd\xe0\x9d\x61\x7d\x86\x39\xaa\x0e\x84\xd7\x46\x45\x31\xba\xe5\x50\xdb\xfb\x9d\xca\x1e\xc4\xc8\xeb\xd5\x41\x72\xd5\x8f\x46\x44\x74\xa2\xb3\x4c\xa9\x22\xfc\xfe\xe0\xe0\x40\xec\x09\x1e\x39\x4d\x33\xe4\x82\x02\x62\x49\xdd\x20\x28\xcb\xd7\xb4\xf1\x44\x9e\xa2\x37\xea\xf3\x99\x2a\x37\xea\xb4\xfa\xc2\x82\xe8\x37\x99\x27\x99\x7a\x49\xa0\x05\xb3\x52\x5d\x51\x26\x95\xc0\x99\x61\x44\xaa\xb9\x69\x4b\x65\x0a\x30\x2a\xc5\xc0\x94\x53\x14\xc9\x27\xb1\xe7\x25\x9f\x2a\x45\x15\xed\xfb\x1f\x24\xd1\xa9\xb2\xd7\x3a\x21\xbc\x82\xc5\xdb\xb3\x77\xee\x84\x43\x04\xc6\x01\xfb\x1b\x9a\xad\x2a\x43\x9e\x8e\xce\x61\x2b\xf3\x46\xdb\x97\x38\xde\xd9\xe7\x36\x52\x0a\xd4\x14\x59\x6a\x29\x23\xf9\xa8\x71\x84\x0f\xf9\x84\xc1\x90\x56\x5a\xaa\xb5\xb6\x8a\xe4\x53\x6a\x09\x13\xe7\x43\xa6\xf2\xb0\x99\x38\x11\xbf\x88\xef\x87\xad\xb6\x3f\x9b\xb6\x6b\xda\x29\x3d\xd7\x7e\x4b\x96\x73\xcd\x09\xcc\xd2\x68\x58\x33\x26\xb5\xb1\x5e\x5f\xa6\xb9\xf4\x29\xd9\x37\x35\x8c\x21\x05\x15\x95\x38\x94\xe7\x05\x07\x51\x3b\x76\x7e\xf0\x11\xc2\xb5\x8c\x4f\x5c\x4d\xbb\xe6\xf3\x29\x7a\xbf\x7c\x1d\xfd\xb5\x42\xfb\x0c\x27\xd1\x2b\x65\xc3\x00\x2a\x17\x75\xdd\x4f\xee\xa5\xf5\xdc\xdb\x16\xf5\xd6\xa9\x64\xa0\xd3\x74\xf1\xd6\x91\x9d\x2a\xad\x3b\xa4\x32\x9b\xe1\x77\xdd\x73\x50\x1e\xf6\x5a\x35\x46\x40\x02\xf1\x2e\x36\xa6\xb8\xc6\xaa\xb3\xcd\xa5\xfe\x82\x19\xfd\x02\x6c\x39\x46\x44\x7f\x42\xc6\x26\x51\x2b\x05\x3c\x1b\xc9\xfb\x3c\x6b\x64\x2d\x63\x38\x6f\x3d\xf9\x88\x95\x5b\xd7\x49\x0d\xab\x10\x05\x09\xcd\x54\xd4\x54\x84\xd3\x9d\x8d\xb4\xd4\x25\x5a\x36\xab\xba\x45\xbb\xa2\xce\xb2\x18\x9e\x0a\xed\xd8\x4c\x4b\x46\xba\x1e\xb8\xbc\x7a\x04\x1d\x0e\x0b\x13\x1e\x62\x34\x7d\x6c\x61\x66\x2b\xb8\x3d\x25\x58\xf4\x4a\x1d\x0f\x06\x5a\xb9\xa6\x10\x40\xc8\x2d\xb3\xf3\x1a\x14\x4a\xc1\x22\xbb\xa5\x16\xc3\xe7\xaa\x5a\x17\xf6\x96\x2e\x21\x81\x6f\x02\x7d\xc2\xc4\xdc\xe4\xd4\x33\x13\xd8\x9b\x8c\xb7\xb6\xb8\xed\x4c\xc9\x33\x1a\xf6\x03\x15\xb1\x93\x9b\xf4\x4b\xa0\x5f\x79\xd4\x4b\x92\xc6\x5b\xf2\xcd\xf3\x6f\x7a\xe5\xc6\x8c\x43\xe3\xa5\x3f\x34\xa0\x8e\xda\xee\x8a\xef\x86\x8d\x0b\x88\x5c\xd4\x68\x5c\xc0\xb5\xff\x9f\x06\xf6\x2f\xd4\xf7\x2e\x76\x4b\x3c\x9e\x0c\x3c\xd3\xc9\xed\x8e\x3d\x7b\xc0\xc5\x39\x6e\x4f\x65\x2e\x33\x6e\xff\x25\xef\xd4\xa4\x37\x25\xf4\x44\xab\xbb\x91\xc3\x78\x5c\x21\xd5\x6e\x44\xbf\x66\xda\x5d\xbe\x98\xf9\x53\x52\x0c\x0a\xe2\x1b\x78\x31\x25\xe1\x90\x19\xff\xf7\x83\xe9\x1d\xe2\x5b\x0a\x6b\xd0\x4a\x77\xd5\xf2\xf6\x52\xf6\xb5\xd5\xe9\x3a\xbb\xfb\x51\xa7\x1d\xed\xe8\x41\xe8\x38\xd0\x1f\xef\xec\x7e\x0f\xc0\x75\x12\xc7\xaa\xc0\xe5\x79\x7b\xf1\xd4\xd1\x9a\xff\x5c\xd5\xbc\x7a\xf1\x6f\x15\x4d\x7b\x38\xdc\xef\xe0\x1d\x49\x07\xb1\xdd\xed\xad\x3d\x50\xfe\x87\x15\x73\xd7\xdd\x8d\xb0\xd3\xb5\xba\x69\xe2\x38\x3b\x75\x41\xfa\x82\x42\xac\xe2\x2f\x20\x72\x94\x71\xec\xc2\x6b\x16\x9f\xe4\x09\xaf\x1e\x06\x47\x2c\x45\xde\x4d\x9a\x4f\x46\xdd\x4b\x9e\x08\x89\xbd\xcf\x17\x8c\x02\xfc\xa4\xe0\xd2\x95\x8c\xa9\x79\x76\x88\x3f\x87\xc4\xc3\xdb\xef\x85\x6e\x17\x9c\xea\xe2\xeb\x9d\xb7\x66\x88\x36\x52\xe3\x46\x3d\xab\x0d\xb9\x0c\x85\xcf\xba\xbc\x21\x56\xe4\xed\x6d\xb9\x17\xbb\x2f\x32\xe0\xd6\xec\x48\xfb\x4d\xc6\xf9\x35\x60\xe8\x2b\xca\x05\x21\xf3\xd6\x28\x95\xd1\x07\x62\x58\x3c\x83\x90\xa5\x36\xd9\x46\xc4\xc3\x11\xd5\x07\x87\xb3\x6d\x7b\xc1\x84\x41\xbd\x2b\xe5\x77\xc5\x3b\x24\x3d\xa1\xf4\xdf\xd2\xc8\xc0\xf0\xcc\x43\x59\x14\x74\xed\x40\xb8\xbe\x3b\xfb\xb2\x96\x35\xcb\x33\x38\xbf\xe2\x6b\xb1\x21\x53\x34\x18\x85\xf4\x31\xd1\x53\x1f\x90\x7c\x25\xf6\x1c\x92\x6f\x40\xf2\x99\xc5\x13\x89\xd8\x60\x20\x9a\x9b\xd7\x5a\x17\x97\x92\x12\x58\x3c\x7e\x2c\xfc\xe0\xab\x4c\x5f\xe2\x6a\x96\xa7\x74\x45\xe8\x48\xde\xe9\x43\xfc\xea\x5d\x40\xea\xbd\x6a\xe5\x7c\x8f\x63\xd9\x5d\x7d\x4f\xe0\x2f\x04\x0f\xef\x6a\x7b\xf5\x5a\xa5\xc0\x1f\x77\x72\xe4\xa3\xdb\x22\x77\xb0\xf7\x6e\x70\x8e\x47\xb2\x74\xc2\x99\x38\x9b\x89\x65\x7b\x63\x16\x6e\x29\xc3\x4c\x11\x90\xe1\xa6\x2d\xf4\x4a\xd0\xed\x96\x2f\x40\x18\x9d\x2f\x1a\x42\x89\xb5\x98\x01\x90\x38\x72\x39\xdd\xbd\x7d\x43\x5c\x27\x35\x69\xf8\x1f\x7c\xc3\xf5\xef\x04\x45\xce\x17\x7c\x30\xfe\xb4\x70\x29\x9d\x13\x4e\x34\xb0\x66\xcd\x6e\x5a\x13\x3c\xc0\x03\x47\x2a\x4d\x62\xfe\x56\xaf\x90\xd7\x1f\x10\x49\x72\xf0\x14\xcf\x5f\x44\x8e\xc7\xfe\xbe\xfb\x96\x56\xd9\xf3\x94\x89\x68\x41\xcf\x7f\x88\xf0\x6f\x64\x1e\xef\x93\x2e\xc4\x50\x6b\xca\xd3\xf5\x68\xf7\x0d\x07\x57\xf4\xce\x87\x32\x5f\x9d\x75\x1a\x7b\xc2\xcc\x9b\xe2\xaa\xe2\xfe\x39\x91\x45\xdf\xd6\xec\x76\xf6\x38\x76\x0e\xc3\x75\x7f\x5d\xdd\xa3\x17\xdc\xa9\xfa\x5f\xcf\x1c\xdd\x9a\x8a\x83\x1f\x0f\x0f\xb7\x1e\xce\xf7\x6c\xf6\x30\x0e\x87\x33\xb6\x7f\xe7\x22\x90\xfc\x67\x2e\xfe\x4e\xcf\x3c\xaf\xfb\x05\xef\x0e\x78\xfe\x33\x00\x00\xff\xff\xb2\xe6\x4f\xab\x59\x18\x00\x00")

func instanceListenerMainGoBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerMainGo,
		"instance-listener/main.go",
	)
}

func instanceListenerMainGo() (*asset, error) {
	bytes, err := instanceListenerMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/main.go", size: 6233, mode: os.FileMode(420), modTime: time.Unix(1474463118, 0)}
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
	"instance-listener/Godeps/Godeps.json": instanceListenerGodepsGodepsJson,
	"instance-listener/Godeps/Readme":      instanceListenerGodepsReadme,
	"instance-listener/main.go":            instanceListenerMainGo,
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
	"instance-listener": &bintree{nil, map[string]*bintree{
		"Godeps": &bintree{nil, map[string]*bintree{
			"Godeps.json": &bintree{instanceListenerGodepsGodepsJson, map[string]*bintree{}},
			"Readme":      &bintree{instanceListenerGodepsReadme, map[string]*bintree{}},
		}},
		"main.go": &bintree{instanceListenerMainGo, map[string]*bintree{}},
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