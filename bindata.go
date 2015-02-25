package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _data_publish_gh_pages_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x57\x6d\x73\xda\x48\x12\xfe\xae\x5f\xd1\xc1\x54\x6c\x27\xbc\xc4\xf9\x78\x59\x92\x53\xb0\x6c\xab\x0e\x83\x4b\xe0\xf5\xa5\xd6\x7b\x41\x48\x23\x4b\x17\xa1\xd1\x6a\x46\xd8\xdc\x26\xff\xfd\x9e\x9e\x11\x18\x58\xd7\x5d\x5c\x2e\x40\x33\xfd\xfe\xf2\x74\xeb\xe8\x55\xbf\x56\x55\x7f\x91\x15\x7d\x51\xac\x68\x11\xaa\xd4\x71\x8e\x68\x28\xcb\x75\x95\x3d\xa4\x9a\x4e\xa2\x53\x7a\xff\xee\xec\x8c\xfc\xb0\xa0\xeb\x30\x1a\x09\x19\xd3\x2f\x59\x58\xfc\xbd\x10\xab\xac\xea\x15\x42\x7f\x74\x8e\xc0\x72\x23\xaa\x65\xa6\x54\x26\x0b\xca\x14\xa5\xa2\x12\x8b\x35\x3d\x54\x61\xa1\x45\xdc\xa1\xa4\x12\x82\x64\x42\x51\x1a\x56\x0f\xa2\x43\x5a\x52\x58\xac\xa9\x14\x95\x02\x83\x5c\xe8\x30\x2b\xb2\xe2\x81\x42\x8a\xa0\x9a\x29\x75\x0a\x31\x4a\x26\xfa\x31\xac\x04\x88\x63\xe8\x08\x95\x92\x51\x16\x42\x22\xc5\x32\xaa\x97\xa2\xd0\xa1\x66\x8d\x49\x96\x0b\x45\x27\x3a\x15\xd4\x9a\x36\x3c\xad\x53\xa3\x26\x16\x61\x4e\x59\x41\x7c\xb7\xb9\xa2\xc7\x4c\xa7\xb2\xd6\x54\x09\xa5\xab\x2c\x62\x19\x1d\xc8\xcf\x8a\x28\xaf\x63\xb6\x63\x43\x90\x67\xcb\xac\xd1\xc1\x02\x4c\x4c\x14\x8b\xad\x15\xbc\x60\x5b\x3b\xb4\x94\x71\x96\xf0\xb7\x30\xae\x95\xf5\x22\xcf\x54\xda\xa1\x38\x63\xe1\x8b\x5a\x0b\x96\xad\xf8\x38\x12\x05\xf3\xc1\x9b\xbe\xac\x48\x89\x3c\x67\x19\x19\x6c\x37\x1e\x3f\x5b\x68\x68\x58\x4f\xc9\x61\xd5\x4d\xa0\x8c\xe6\xc7\x54\x2e\xf7\xbd\xc9\x14\xe4\x27\x75\x55\x40\xad\x30\x5c\xb1\x44\xe8\x3a\xac\xf3\xdf\x22\xd2\x7c\xc2\x0c\x89\xcc\x73\xf9\xc8\xee\x45\xb2\x88\x33\xf6\x4a\xfd\xcd\x24\x6f\x86\xdb\x70\x21\x57\xc2\xb8\x64\x33\x5f\x48\x0d\x7b\xad\x1d\x9c\x8b\xf2\x39\xc1\xcd\x95\x4a\x43\x38\xb0\x10\x4d\xdc\xa0\x1a\x71\x0e\x77\x7c\xaa\xac\xdf\x4a\xa3\x0a\x32\xa4\xa1\x94\x95\x51\x7a\xe8\x6d\xcf\x1a\x71\xe5\xd1\x74\x72\x31\xbb\x73\x03\x8f\xfc\x29\xdd\x04\x93\x5f\xfd\x73\xef\x9c\x5a\xee\x14\xcf\xad\x0e\xdd\xf9\xb3\xab\xc9\xed\x8c\x40\x11\xb8\xe3\xd9\x17\x9a\x5c\x90\x3b\xfe\x42\xff\xf0\xc7\xe7\x1d\xf2\xfe\x79\x13\x78\xd3\x29\x4d\x02\xf2\xaf\x6f\x46\xbe\x87\x33\x7f\x3c\x1c\xdd\x9e\xfb\xe3\x4b\xfa\x7c\x3b\x83\x8e\xf1\x64\x46\x23\xff\xda\x9f\x41\xec\x6c\x62\x54\x36\xc2\x7c\x6f\xca\xe2\xae\xbd\x60\x78\x85\x47\xf7\xb3\x3f\xf2\x67\x5f\x3a\x74\xe1\xcf\xc6\x2c\xf5\x02\x62\x5d\xba\x71\x83\x99\x3f\xbc\x1d\xb9\x01\xdd\xdc\x06\x37\x93\xa9\x07\x03\xce\x8d\xe0\xb1\x3f\xbe\x08\xa0\xc9\xbb\xf6\xc6\xb3\x1e\x34\xe3\x8c\xbc\x5f\xf1\x40\xd3\x2b\x77\x34\x32\xca\xdc\x5b\x78\x10\x18\x1b\x87\x93\x9b\x2f\x81\x7f\x79\x35\xa3\xab\xc9\xe8\xdc\xc3\xe1\x67\x0f\xb6\xb9\x9f\x47\x9e\x55\x06\xc7\x86\x23\xd7\xbf\xe6\xea\x39\x77\xaf\xdd\x4b\xcf\xf0\x4d\x20\x27\x30\x84\x8d\x85\x77\x57\x9e\x39\x82\x46\x17\xff\xc3\x99\x3f\x19\xb3\x2b\xc3\xc9\x78\x16\xe0\xb1\x03\x4f\x83\xd9\x96\xf5\xce\x9f\x7a\x1d\x72\x03\x7f\xca\x61\xb9\x08\x26\xd7\x1d\x9a\x98\xe0\x80\x67\x62\xc4\x80\x73\xec\x59\x39\x1c\xf2\xfd\xcc\x80\x84\x9f\x6f\xa7\xde\xb3\x35\xe7\x9e\x3b\x82\xb4\x29\x33\xef\x12\xf7\x18\x4c\x46\x72\x85\x9a\xcb\xd7\xa4\xc3\x6f\x02\xbd\x5a\xa1\x7e\x1f\xd0\x5f\xf5\xa2\x17\xc9\x65\xdf\xe0\x48\xff\xa1\x92\x11\xa8\x95\xd0\xd4\x15\x74\x44\x53\x2d\x4b\x6a\xba\x2e\xc9\x2a\xa5\x29\x09\xb3\xbc\x46\xb5\xeb\x34\xd4\x24\xa3\xa8\xae\x94\xe3\x9c\x4f\x86\xd3\xaf\x37\xee\xec\x6a\xd0\x83\xc8\xbe\xaa\xa2\x48\x76\xf5\xb2\x74\x66\x6e\x70\xe9\xcd\xbe\x7e\x46\x6a\x87\x57\x83\x87\xb4\x5b\x86\x0f\x42\x39\xbf\xfd\x46\xed\x33\xfa\xfd\x77\x7a\xfd\x9a\x1a\x92\xc0\xbb\x9e\xcc\xbc\x01\x8e\xbf\x7f\x3f\x38\x93\xe8\x84\xac\x60\x1f\x2e\xd1\x83\xaa\xcc\xd0\xfb\x8c\x0a\x28\x68\x5d\x2b\x54\x7b\x22\xab\xa5\x45\x07\xfc\xb7\x95\x8e\x45\x55\xd9\xd6\x7d\x14\x68\xc2\xe2\x58\xd3\x23\x6a\x9f\x1b\xb0\x12\x79\xb8\xb6\xc6\x87\x0a\x24\x04\x52\x60\x80\x6d\x4d\x68\x00\xa0\x54\x3d\x42\x4b\x30\xeb\x63\x15\x96\x1c\x23\x23\x0a\xcd\x6c\x69\x53\x3c\xe5\xdc\xbe\x12\xbe\x8b\x7c\x25\x54\xaf\xd7\x73\xc4\x93\x88\xbe\x82\xf6\xe4\x94\xfe\x74\x88\x80\xaf\x6a\x70\x7c\x8c\x5f\x30\x8e\x4e\x4e\x28\xa3\x01\x9d\x7d\xc0\xd7\x2f\x03\x6a\x1f\xe1\xc7\xdb\xb7\x74\x7a\xfa\x01\x62\x41\x43\x24\x56\x68\x4c\x30\x0d\xee\xdb\xed\xcc\x9c\x64\x09\x71\x98\x70\x46\x83\x01\xbd\xb9\xa7\x37\x08\xd8\x07\x36\xb3\x30\xf7\x84\xf4\xfc\xa0\x3b\xc1\x70\x0c\x0c\xb6\x3e\xc5\x12\xf8\x26\xe8\x8f\x5a\x6a\x74\xfe\x63\x06\x14\x00\x4a\x30\x38\xc0\xb1\x98\x5d\x2e\xc3\x4a\x73\xe3\x43\xae\x41\x6e\x82\xe2\x1a\x2e\x34\x32\x8d\xdd\x2d\xd6\xaa\xe8\xde\x7c\xdf\xb7\x5a\xd6\xc2\x5c\x89\x17\x88\xf8\xd3\x12\x24\x6c\x37\xa2\x2d\x1c\x7c\x73\x05\xbd\x65\x7a\x36\xd2\x5b\x09\x2e\x21\x59\x3f\xa4\xdb\xb0\x72\x45\x6d\x6c\xe0\x1c\x36\x36\x77\xcc\xf0\xe3\xd9\xa5\xca\x3c\xd3\xda\xc4\xb9\x00\xda\x66\xb8\x2c\xc3\xc8\x0e\x8d\xac\x40\x8e\xee\xd2\xf5\x27\x28\xc0\xf0\xa8\x73\x3d\x98\x9b\x08\x72\xb2\xac\x5d\xef\x3f\xbe\x3e\x9b\xb3\x21\xa6\x48\x06\xed\x4f\x8d\x51\x5d\x63\x5e\x13\xdc\xa6\x82\xba\x85\xa0\x77\xfb\xd1\x15\x51\x2a\xa9\xd5\xb6\xc2\x5b\xf4\xf1\xf5\x7b\x7b\xfc\xc4\x0a\x2c\x9b\x63\x5c\x76\x0e\x69\x8d\x49\x1a\x63\x80\xde\x39\x3f\x1c\xc7\x6a\x9a\xb3\x61\x1b\x6d\x6a\x4e\xaf\x06\xd4\x6a\xed\x6a\xb4\x32\x6e\x72\x11\x2a\x46\xff\x25\x4f\x1b\x9e\x4d\x9a\xa3\xb1\x46\xa5\xf1\xcc\x2e\xd0\x3b\xc8\x25\x2a\x4a\x6c\xe6\x1b\xc7\x67\x7f\x08\xa3\x94\x6d\x5b\xbf\xda\x98\x6d\x8c\x3e\x73\xd8\xd8\xe1\x6d\x10\x00\x0a\x37\x0d\x69\xcc\x5a\x60\x3b\x88\x52\x04\xac\x1f\x8b\x55\xbf\xa8\xf3\xfc\x3b\x71\xb1\x74\x0b\x3a\xee\xff\xeb\xfe\x4d\x5f\xf1\x27\xf5\xfb\xe5\xf1\x7c\x2b\x60\x38\xb9\x06\x7c\x5b\x01\x95\x58\xa1\xaf\xd1\x0b\x74\xe5\xb9\xe7\x73\xc7\xb4\xf7\xfb\xa6\xbd\x2d\xe1\xd7\x6b\x80\x37\xa0\x73\x80\x73\xf4\xf7\xc1\x61\xeb\x52\x14\xa2\x7a\x69\xa1\x40\x08\xda\xfb\x2a\x5b\x8c\x03\x37\x88\xb5\xa8\x56\xc2\x54\x51\x59\x49\x1e\xaf\xc7\x8a\x18\x7c\xb2\x87\x82\xc3\xa3\xa4\x6d\x88\x2d\x00\x44\xa9\x88\xbe\x71\x9d\x41\xa4\x04\x5b\xf5\x98\xc1\x60\x15\x55\xe2\x91\xea\x92\xd2\x2c\x8e\x45\xb3\xc1\x34\x39\x03\x00\xee\x08\xdc\xc9\x55\x54\xee\x5e\xb4\xb7\xe0\xd7\x37\x21\xde\x49\x78\x13\xd9\x6e\xb7\x90\xdd\x48\xe6\xd0\xfc\x1d\xbb\x98\x28\xa9\x45\xed\x3d\x64\x6c\xcd\xb9\xc7\xf7\x2b\x02\xd3\x06\xdb\x19\x25\x42\x43\x84\x41\x6b\x76\x56\xf3\xe2\xc6\x21\x5f\xa2\x5b\x38\xd7\x0a\x1b\x1d\x54\x66\x0c\x71\x8a\xe0\x8e\x09\xa3\xbd\xcf\xd7\x26\xfb\x16\x98\x1a\x49\xed\x3d\x6c\x75\x1a\x45\xa8\x2b\x50\xa0\x50\x94\xde\xf2\x7e\xda\xf6\xc9\x9e\x33\xe1\x8b\xfe\x58\x26\xd5\xdf\x17\xdf\xff\xff\x6e\x6e\x4a\x7f\x2c\xe9\x78\x9f\xfa\x78\xa3\xd2\x98\xa5\xd0\xf1\x43\xf6\xce\x22\x82\x68\x40\x69\xe3\x9c\x5a\x2f\x17\x12\x9b\x5d\xb7\x12\x89\x29\x43\x58\x94\xa8\x7e\x2a\xc2\x58\x1d\x58\x61\x18\xab\xa5\xc9\x61\x3f\x2b\x62\xf1\xe4\x38\x16\x4e\xb7\x65\x65\x53\x1b\xdb\x72\x00\x24\x61\xd8\x2c\x31\x35\xb1\x45\x99\xc9\x27\xd6\xc7\xbc\x0c\x47\xba\xc6\x9a\xb5\xde\x50\xbf\xda\x81\xed\xee\x5e\x5d\xbc\x5c\x46\xfc\x87\x52\x7a\x99\xee\xf9\x67\x43\xb9\xf5\x34\x8c\xe3\xc3\x5b\x0b\x43\x3b\x34\x11\x70\xa4\xa0\x6e\x12\xff\xe1\xec\x00\xf7\x8b\x53\x75\x2f\x03\xb9\x8c\x00\xa2\x4d\xd8\x0f\xd3\xd1\xb1\x3d\x64\x27\x9f\x7e\xbe\x7e\x31\xd5\xc7\x76\x63\xad\x42\xcb\xc0\xbd\x78\x90\x31\x23\x8c\x05\x75\x17\x07\xdd\x40\xff\x53\x72\x03\xbb\x8d\x57\x7f\x15\x77\x40\xcd\xb4\x47\x3c\x26\x37\x1b\xc0\x37\x81\x82\x05\x0e\x00\x64\xcb\x5c\xa0\x89\xd4\xba\x88\xe8\x24\x16\x78\x30\xbe\xe5\x06\x85\x90\x76\xd4\x77\x54\xd5\x89\xb6\xdd\x57\x02\xe6\x32\x09\xfc\xde\x87\x28\xa8\x2c\x6b\x7d\xba\x9d\xff\x94\xab\xae\x7d\x0f\xfa\x4e\x4f\x66\x12\xa1\xd4\xba\x89\xd3\xac\xf7\x5b\x29\x4a\xe6\xb5\x11\xb0\x10\x78\x1b\x30\x6b\x95\x79\xaf\x30\x2f\x5e\xbd\x58\xea\x6e\x53\x7c\x8c\x55\x16\xf2\x81\x65\x60\xc9\xb3\xff\xf0\xa2\x32\x4f\x50\xba\xf3\x0e\x4f\x46\xe4\x8a\xdf\xac\x78\xe4\x84\x39\x80\x6f\x19\xc6\x16\x19\xe7\x59\x32\xef\xf2\xd4\x11\x66\xd2\xcb\x05\x94\xc2\x63\xbb\xe7\x7f\xdc\x14\x60\x37\x48\x76\x8b\xf0\x0d\xf5\xb6\xb7\x3f\x5b\xcc\x1b\xfa\x9f\x2b\xe8\x0d\x35\x52\x73\x84\x2f\x7c\xb8\xb9\x16\x55\x81\x80\xae\xc4\x73\x60\x6a\xc5\xe9\x98\x03\xf1\xe6\x3b\x06\x9f\x50\x14\xef\xa8\xa0\x0f\x0c\x89\x80\x24\xf1\x14\xe1\xdd\xc7\xa8\xa1\x28\xa1\x2e\xf5\xe8\x14\x29\xe0\xcb\xa7\x6f\xe5\x6a\x85\x23\x48\xe1\xa0\xed\x72\x77\x97\xe1\x53\x2c\x4a\x9d\xd2\x19\xe6\x1d\x36\x25\x0c\x32\x3c\xec\x52\x70\x1d\xbf\x78\x65\xbc\x43\x70\x90\xf8\x4d\x18\x5b\x7f\xfe\x68\x41\xf1\xfd\x07\x4e\x37\xe0\x1b\x6c\x66\x42\xd7\x05\x52\xa9\x38\x81\x5b\xd0\x48\x43\x7e\xcd\xb3\x33\xfd\xa7\x97\x84\x5d\x0c\xe8\xba\x7b\xe5\x6f\x57\x86\xee\x12\x8b\xc8\xfe\x74\x6d\xed\x92\x95\xb5\x3a\x1c\x01\x2f\x37\xcc\xd0\x20\x08\x06\x63\x98\x20\x35\xcf\xab\xae\xc3\xd5\x5c\xed\x16\x8c\xe3\xfc\xb5\x07\x5b\xed\xfd\x45\xa3\xe5\xfc\x37\x00\x00\xff\xff\x5c\x24\xc6\x2a\xe2\x10\x00\x00")

func data_publish_gh_pages_sh_bytes() ([]byte, error) {
	return bindata_read(
		_data_publish_gh_pages_sh,
		"data/publish-gh-pages.sh",
	)
}

func data_publish_gh_pages_sh() (*asset, error) {
	bytes, err := data_publish_gh_pages_sh_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/publish-gh-pages.sh", size: 4322, mode: os.FileMode(420), modTime: time.Unix(1424899082, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_srcco_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\xdd\x92\xda\x36\x14\xbe\xdf\xa7\xd0\x6c\x27\x33\xd9\x8c\xed\x08\x30\xb0\x0b\x37\x4d\xb7\xe9\xf4\xa2\xe9\xf4\xa6\xbd\x97\x2d\x81\x55\x84\xe4\x91\x64\xc0\xc9\xe4\xdd\x7b\x24\xcb\x3f\x60\x9c\x69\x02\xcb\x8e\x2c\x9d\xef\xfc\x7e\xe7\xc8\x85\x3d\x0a\xf4\xe5\x01\xc1\xa7\x60\x7c\x5f\xd8\x0d\x22\x95\x55\x5b\xbf\x73\xe4\x32\x6e\x77\x67\x18\xbf\xd9\x3e\x7c\x7d\xc8\x14\xad\x03\xe0\xfd\x3b\xb4\x53\xd2\xc6\x3b\x72\xe4\xa2\xde\xa0\xc7\xdf\x99\x38\x31\xcb\x73\x82\xfe\x64\x15\x7b\x8c\xba\xe7\xe8\x83\xe6\x44\x44\x86\x48\x13\x1b\xa6\xf9\x6e\x8b\xde\xbd\xf7\x3a\xfe\xa7\x82\xc6\x9f\x8c\xe4\x87\xbd\x56\x95\xa4\x71\xae\x84\xd2\x1b\xa4\xf7\xd9\xdb\x17\x1c\xa1\xe6\xf7\xe4\x1c\x4c\xf6\x9a\xd3\x9b\x90\x1a\xe7\x7d\x48\x44\xef\xb9\xdc\x20\x5c\x5e\x06\x81\x96\x84\x52\x2e\xf7\xcd\xf6\x6c\x59\x5e\xa6\xec\xfd\xf4\xdb\xdc\x7d\xfb\xf4\x9c\x39\xb5\xc5\x06\xad\x30\xbe\xb4\x06\x2e\xed\xe6\x6c\x8e\xb1\x53\x05\x3e\x69\x75\x0e\x2e\xa9\x13\xd3\x3b\xa1\xce\x1b\x54\x70\x4a\x99\xbc\x71\x60\x76\xed\x41\xd0\xd4\x78\x0a\x7a\xa8\xca\x83\x1e\x9f\x38\xc3\x3f\x33\xc0\xac\xcb\xcb\x0f\x44\xd7\x86\x94\xa6\x69\xb3\x01\x6e\x11\x48\x96\x60\x3b\x7b\x65\x7d\xd9\x85\xe1\xcc\x97\x9a\xdd\x71\x61\x7e\xe3\x71\xc0\x0c\x23\x8e\x2f\x43\x6e\x75\xbb\x75\x9f\x09\xb0\x90\x2b\x7a\xa5\xbe\xa5\xc6\x27\x26\x85\x8a\x3e\x29\x49\x72\x15\xbd\x2a\x69\x94\x20\x26\x7a\x7c\x55\x95\xe6\x4c\x03\x5d\xce\x8f\xd1\x51\x49\x65\x4a\x92\xb3\xed\x37\xdc\x2b\xb8\x65\xb1\x97\xda\xb8\x50\xe2\xb3\x26\xe5\xf6\x1b\x95\xb9\x9f\xd2\x4c\x5d\x62\x53\x10\xea\xa4\xdd\x89\xcf\xac\xfb\x07\x84\x24\x6f\x81\x8c\xe1\x2f\x99\x3d\x21\x2e\x0d\xb3\x2d\x4c\x53\xa6\x63\x4d\x28\xaf\x4c\x53\xec\x49\xa6\x2d\xd6\x8b\x97\xc5\x7a\xc4\xb4\xb4\x4f\xec\x80\x6a\xcf\xfd\xee\x98\xf3\x57\xe4\xea\xb2\xbc\x29\x5c\xc0\x88\x84\x6c\xc7\x67\x96\x1d\xb8\x8d\x89\xe4\x47\x62\xb9\x82\x90\x6b\x26\x5c\x81\x76\x04\x4a\x82\x93\x95\x41\x8c\x18\x16\x83\x2b\xaa\xb2\x68\xd6\xe8\x8e\x8f\xea\xf3\xf7\x62\xbe\x43\xbc\x65\x04\xe9\x47\x4e\x1f\x0d\xa4\xdb\xd5\xb5\x1d\x25\x70\xd6\x16\x2b\x76\x67\x71\x7b\x38\x64\x3b\x97\x05\xcc\x9f\x50\x0e\xcb\x2e\x36\xa6\x2c\x57\x3a\xb8\x23\x95\x64\x77\x2b\xb5\x18\x17\x0a\x1a\x85\x4b\x46\x74\xbc\x77\x42\x4c\xda\xb7\x56\x01\xcc\x5a\x75\x8c\x1a\x1a\xcc\x97\xcb\xa8\xfd\x25\x8b\xe5\x13\xc2\x6f\xa2\xf1\x01\x86\x03\x57\xab\xa7\x41\xb4\xa1\x36\x5f\x7e\xd8\x24\x8e\x92\x35\x1e\x1b\xc4\x11\x4e\x16\x78\x60\xef\x01\x92\x16\x77\x1f\x24\xc9\x69\xf0\x78\xff\x03\x09\x4d\xac\xca\x4d\x5f\x11\xca\x0d\xc9\x04\x83\x86\xd3\x90\xc1\x73\x5b\x0f\xd8\x2e\x05\xa9\xc7\x59\xdd\x20\xe8\x5e\x98\xd1\xae\x48\x99\x80\xe8\xae\x52\x6e\x55\x39\x84\x94\xca\xf0\xa6\x38\x3b\x7e\x61\x74\x5c\x85\x86\x42\xe3\x66\xed\x07\x9d\xd7\x88\x9b\xb5\x6e\x9a\x03\x5f\x0d\xab\xd9\xb8\x7b\xe6\x6d\xab\xb8\x50\x13\x92\x5b\x7e\x6a\xe7\x52\x00\x2d\xf0\x50\x04\x5d\xcb\x74\xb1\x67\x42\xb9\xf8\x1a\xa9\x58\x92\x23\xbb\x53\x55\x28\x67\x3d\x31\x82\x87\x11\xe1\xd1\x8c\xed\x22\x74\x44\x26\x82\xef\x41\x32\x07\x5e\x30\x3d\x1d\x8d\x77\xe2\x6e\x44\xb3\xe5\x40\x2c\x1c\xf5\xe9\x27\x19\x54\xad\xb2\x6c\x3c\xbe\x4d\xae\x95\x10\xa3\xe9\x73\xe3\xed\xa2\xdf\xf1\x05\x99\x77\x8f\x37\x44\x01\xfb\x12\xba\x20\x1e\xbc\x69\x8c\x25\x9c\x87\x82\x87\x63\xc1\x0d\x4c\x7a\x5b\x0b\x16\xdb\xba\x64\x43\xfa\x74\xf9\xbb\x61\x3b\x32\xb5\xb4\xe4\x02\xa3\x7e\x5f\x08\xe7\x32\x8c\x94\x7b\xd4\x07\x2e\x03\xe8\x2f\xcd\xac\xad\xe1\xbe\xe0\xd2\x0b\x7a\x5b\x26\x41\x7f\x1b\x46\x21\x3e\x5b\xb8\xab\xc4\x5a\xbe\xab\x93\x7f\x61\xbb\x01\xfd\xc3\x8f\xc8\x54\x32\xab\xb4\xb1\xc8\x16\x0c\x4a\x9f\xd5\xe8\x57\x72\x02\xea\xff\xc1\x78\xa6\x4e\x3c\xf7\xe3\x29\x31\x56\x47\xc8\x77\xbe\x5b\xa2\x2f\xdd\xd5\xbc\x5a\xfe\x82\xd3\xf9\x16\x7d\x75\x6d\x06\x47\xce\x36\x8a\x81\x2f\x8c\x49\x07\x4d\x0e\x67\xda\x22\x61\x39\x40\x7e\x9c\x3f\xbf\xac\xd2\x80\x3c\xb0\xfa\x0c\xcd\x05\x48\x4a\xf4\x01\x95\x5c\x1e\x3c\x3a\x77\x73\xa3\x41\xc3\x72\x80\xfe\xf0\xd1\x7d\xb7\xe1\x0a\x75\xc1\xc2\xe4\xb4\x40\xb0\x3c\x68\x04\xf1\x23\xf0\xcc\xfb\x42\x6a\xaf\x0c\x32\xdf\x2a\x83\xe5\x40\xd9\xf3\x4b\x46\x77\xbb\x00\x74\xf5\x01\x94\x4f\x39\x34\x47\xc5\x3c\x56\x70\xdb\x62\x61\x39\xc0\x2e\x16\xcf\xeb\xd7\xd7\x80\x85\x23\xa6\x89\x00\x78\x07\x2c\x2b\xd9\x02\x61\x39\x00\xee\x3a\x8b\xb0\x9f\xdb\xca\x0f\x78\x40\xfa\xfb\xbf\x81\x8a\x1e\x2a\x26\xa0\x82\x40\xc1\xa1\xb7\xae\x80\x96\xec\xbb\x40\xc9\x7e\x2a\x50\xf7\x62\xfd\xfe\x02\x2f\xd7\x4e\xc6\x5d\x94\xb7\x31\x13\xdb\xd9\x87\xe5\x40\x4d\x46\xb3\xf5\x2a\xbb\x55\x43\x2c\x94\x3f\x83\xfe\x43\x7e\x88\x80\xbe\x43\x41\x0e\x3c\xa8\x3a\xf5\xaa\x4e\x53\xfc\xb9\xa3\xea\x44\x9c\x37\x43\x42\xc1\x6d\xd8\xaa\x82\xe5\x54\x25\xe0\x08\xae\xf0\x61\x25\x7e\x3e\x32\xca\x49\xd3\x21\xbe\x2b\xa7\x49\x8d\x57\x18\xd4\x38\x91\x49\xf6\x62\xbc\x0a\xe4\x3b\x87\x79\x92\x29\x41\x03\x6a\x92\xb5\xf0\x12\x3e\x41\x59\x07\x9b\xe4\x67\x8a\xd3\x69\x63\x93\xcc\xc4\x69\x1a\x44\x26\x39\x98\xa6\x6d\xa0\x93\x5c\xc3\xb8\x15\x99\x64\xd5\x37\x73\x31\x49\x22\x1f\x54\x10\x99\x20\x47\xa8\xc3\xd7\x87\xff\x02\x00\x00\xff\xff\xd9\xfb\x6b\xc6\x04\x0e\x00\x00")

func data_srcco_css_bytes() ([]byte, error) {
	return bindata_read(
		_data_srcco_css,
		"data/srcco.css",
	)
}

func data_srcco_css() (*asset, error) {
	bytes, err := data_srcco_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/srcco.css", size: 3588, mode: os.FileMode(420), modTime: time.Unix(1424900150, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_srcco_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x4d\x6f\x13\x31\x10\xbd\xf7\x57\x98\xe5\x10\x47\x69\xb6\xdc\x03\x42\xa8\xea\x01\x09\x01\x52\x11\x17\xc4\xc1\xb5\x27\xbb\x06\xc7\x0e\xf6\xec\x46\x11\xca\x7f\x67\xec\x4d\x52\xe7\x63\x93\xad\xc4\x1e\x5a\xc5\x7e\xf3\xfc\xde\xf3\x78\x56\xda\x2a\xb7\x2a\x9d\x35\x4e\x28\xf6\x8e\xcd\x1b\x2b\x51\x3b\xcb\xf8\x98\xfd\xbd\x61\xf4\xdd\xdd\x31\x29\x2c\xfb\xc8\x94\x63\x58\xeb\xc0\x56\x1a\x6b\xd7\x20\xfb\x25\x5a\x11\xa4\xd7\x4b\x7c\x9f\x80\xad\xf0\x4c\x3a\x05\x81\x68\x94\x93\xcd\x02\x2c\x96\x7f\x1a\xf0\xeb\x47\x30\x20\xd1\xf9\x0f\xc6\xf0\xa2\x8c\x98\x62\x3c\x4b\x35\x73\xe7\x19\x8f\x85\x9a\x8a\xde\xcc\xe8\xdf\xdb\x8e\xa3\x34\x60\x2b\xac\x69\x65\x32\xd9\x29\xd9\x1d\x12\x70\x6d\x80\xf0\x5b\xed\x15\xe0\xbd\x5b\x2c\x1b\x04\xf5\x18\x77\x78\x22\xf8\xa1\x7f\x96\x4b\xe1\x49\xc3\x83\x81\x28\xe5\x96\xd9\xc6\x98\xed\xb9\xf1\xdb\xc3\x12\x5f\x59\x83\xae\x6a\x24\xda\xee\x27\xb1\x7e\xf5\x6e\x09\x1e\xd7\xdf\x85\x69\x80\x17\x1d\x60\xa7\x7c\xb3\xf7\x6c\xc5\xe2\xaa\x67\x74\x72\x1a\x71\x97\x7c\x27\x9e\x3e\xdf\xdd\xa6\x46\x58\x70\x3d\x2e\x85\x52\x0f\x2d\x1d\xf5\x49\x07\x04\x0b\x9e\x17\xd2\x68\xf9\xbb\xb8\xdd\xdf\x1f\x87\x36\x2f\x8f\x1f\x7a\x5d\x55\xe0\xbf\x7d\xb9\xa7\xcd\x12\x85\x27\x8b\xa5\x56\x59\x22\x9b\x13\x6f\xa8\xd1\x5c\x35\x67\x29\xc8\x69\x42\x5e\xb2\xd7\x51\xf5\xf9\xdb\xee\xfe\x0f\x83\x9f\x49\x4e\xe6\xb0\x6b\x82\xb8\x78\xd6\xe9\x66\x76\x73\xb3\x6f\xfa\x2c\x22\x0a\x66\x4b\x1f\x7d\x78\xe7\xb0\x37\x05\x5e\xbc\x2e\xd8\x84\x69\x45\x7f\x8a\x29\x5d\xf5\x2e\x06\x3d\x67\xfc\x55\x2c\x2d\xa5\x11\x21\x44\x33\xd4\xfd\x16\x85\xb6\x81\x17\x82\xce\x6c\x29\xb2\xdc\xc6\x11\x98\x62\x78\xc6\xcd\x0e\x51\xcf\xbe\x5e\x5e\x70\x24\x3f\x6b\xcf\x81\x5c\xc7\x04\xa9\x05\x9e\x9c\x5a\x5f\x63\xd8\x30\x30\x01\xfa\x1d\x7b\x58\xb8\x16\x5e\x66\x7a\x78\xcd\x20\xdf\xfd\x74\x0a\xd2\x9a\x40\x88\x8d\xdf\x1b\x04\x5d\xe9\xbe\xbb\xb2\xe6\x3a\xac\x8e\xc0\xbc\xc3\xec\x76\x72\xda\x13\x9d\xd9\x23\xeb\x12\xbe\x30\x42\x2e\x8d\x4e\xbb\x9b\x77\xd7\xac\x1e\xca\xce\x5f\xd5\xb1\xe8\xa8\xe7\xac\xe6\x73\x82\xd3\x6b\x88\x0b\x83\x5f\x03\xed\x06\x47\xb3\xd8\xb8\x8a\x8f\x6a\x3d\xca\xae\xe2\x88\x67\x48\xa7\x1d\xb0\x3d\xad\x61\xd4\x7b\xb3\x89\xfc\x64\x6e\xc4\x58\xfe\x05\x00\x00\xff\xff\x5a\x32\x8a\x2d\x2f\x07\x00\x00")

func data_srcco_js_bytes() ([]byte, error) {
	return bindata_read(
		_data_srcco_js,
		"data/srcco.js",
	)
}

func data_srcco_js() (*asset, error) {
	bytes, err := data_srcco_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/srcco.js", size: 1839, mode: os.FileMode(420), modTime: time.Unix(1424901399, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_view_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xbd\x4e\xc3\x30\x10\xde\x79\x8a\x23\x03\x5b\x93\x07\xc0\xe9\xd2\x82\x18\x40\xad\x68\x16\xc6\xd4\xbe\x34\x06\xd7\xae\x6c\x17\xa8\x22\xbf\x3b\xe7\x24\xad\x8c\x40\x81\xa9\xd7\xfb\x7e\xee\xf3\xe9\xc2\xae\x97\xab\x45\xf5\xb2\xbe\x83\xd6\xef\xd5\xfc\x8a\x0d\x3f\x00\xac\xc5\x5a\xc4\x82\x4a\x2f\xbd\xc2\x79\xd7\xe5\x55\x2c\x42\x60\xc5\xd0\x19\x50\x25\xf5\x1b\x58\x54\x65\xe6\xfc\x49\xa1\x6b\x11\x7d\x06\xad\xc5\xa6\xcc\x48\xf3\x8c\xce\x1c\x2d\xc7\x35\x35\xe4\x67\x08\xce\x72\x6e\x72\xee\x5c\x36\xea\x1d\xb7\xf2\xe0\x81\xfa\x13\xfc\x57\xa2\xb3\x62\xa0\xf6\xf1\x8a\x73\x3e\xb6\x35\xe2\x34\x5a\x09\xf9\x0e\x5c\xd5\xce\x95\x99\x37\xfc\x3c\x61\x04\xa4\x28\xb3\x46\x52\xc0\x2c\xe1\xcc\x74\xbd\xc7\x0b\x0f\xa0\x27\x9c\x55\x05\xc9\x7e\x58\x08\x6c\x26\x1d\x22\x3e\x69\xd0\x8f\x98\x91\x32\x75\x49\x0c\x68\x07\xf7\x44\xa9\xea\xad\xc2\x55\xb3\x30\xda\xa3\xf6\x2e\x84\x3f\x53\x4d\x7b\x6e\xbc\x3d\x72\x7f\xb4\x28\xfe\xe1\x9c\x96\xc9\x52\x77\x56\x8a\x8b\x6b\xd7\x81\xad\xf5\x0e\x21\xdf\xe0\x6e\xff\xdd\x29\xd1\x58\xf3\x91\x04\x49\x11\x11\x23\x76\x9d\x6c\x20\x5f\x1a\xfe\x50\x3d\x3d\x86\x40\x41\x93\x1a\x95\xa3\x73\xbb\xd1\x5b\x77\xb8\xa5\x7f\x5a\xc4\xdb\x4b\x9e\x1f\x43\x44\xf9\xc2\x08\x1c\x34\xa9\x3f\xa7\x6e\x1c\x90\xc2\xbd\x7a\xb4\xfa\x6d\xa1\xf4\x28\xc2\x60\x04\x2f\x10\x2b\x86\x2b\xa3\xb3\xeb\xbf\x8f\xaf\x00\x00\x00\xff\xff\xeb\x6c\x9b\x5d\x37\x03\x00\x00")

func data_view_html_bytes() ([]byte, error) {
	return bindata_read(
		_data_view_html,
		"data/view.html",
	)
}

func data_view_html() (*asset, error) {
	bytes, err := data_view_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/view.html", size: 823, mode: os.FileMode(420), modTime: time.Unix(1424899082, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"data/publish-gh-pages.sh": data_publish_gh_pages_sh,
	"data/srcco.css": data_srcco_css,
	"data/srcco.js": data_srcco_js,
	"data/view.html": data_view_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"publish-gh-pages.sh": &_bintree_t{data_publish_gh_pages_sh, map[string]*_bintree_t{
		}},
		"srcco.css": &_bintree_t{data_srcco_css, map[string]*_bintree_t{
		}},
		"srcco.js": &_bintree_t{data_srcco_js, map[string]*_bintree_t{
		}},
		"view.html": &_bintree_t{data_view_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

