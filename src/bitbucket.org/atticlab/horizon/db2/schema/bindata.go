// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_aggregate_expenses_for_accounts.sql
// migrations/4_account_statistics_updated_at_timezone.sql
// DO NOT EDIT!

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x6d\x6f\xdb\xbe\x11\x7f\x9f\x4f\x41\xf4\x8d\x1d\xc0\x0e\x9c\xb4\xe9\x1a\x07\x2d\xe0\x26\xea\x6a\xcc\x91\xdb\x58\x5e\x5b\x0c\x83\x40\x4b\xb4\xad\x55\x16\x55\x51\x4e\x93\x0e\xfb\xee\x3b\x3d\xd9\x7a\x20\x45\xda\x91\xf2\xcf\x9b\xc0\xe2\xf1\xee\xf7\xbb\x23\x8f\x77\x94\xfa\xfd\x93\x7e\x1f\x7d\xa1\x2c\x5c\x05\x64\xf6\x75\x82\x6c\x1c\xe2\x05\x66\x04\xd9\xdb\x8d\x0f\x63\x27\x27\x33\xcd\x40\x2c\xc4\x21\xd9\x10\x2f\x34\x43\x67\x43\xe8\x36\x44\xef\xd1\xe0\x3a\x1e\x72\xa9\xf5\xb3\xfa\xd4\x72\x9d\x48\x9a\x78\x16\xb5\x1d\x6f\x05\x03\x9d\xb9\xf1\xe9\x5d\xe7\x3a\x53\xe7\xd9\x38\xb0\x4d\x8b\x7a\x4b\x1a\x6c\x40\xc2\x64\x61\x00\xff\x18\x48\x52\x2f\xd5\xb1\x26\xa0\x7a\xb9\xf5\xac\xd0\xa1\x9e\xb9\x00\x4d\x24\x1a\x5f\x62\x97\x91\x82\x19\x50\x60\x6e\x08\x63\x78\x15\x0b\xfc\xc6\x81\x07\xba\xae\x53\xec\x04\x07\xd6\xda\xf4\x71\xb8\x86\x31\x7f\xbb\x70\x1d\xab\x87\xfc\x95\x69\x01\x55\x97\x66\x62\x36\x59\xe2\xad\x0b\x04\xf1\xc2\x25\xcc\xc7\x16\x89\x40\x77\x4a\xa3\xbf\x9d\x70\x6d\x52\xc7\xce\xe1\x38\x49\x7c\xa8\xe3\x0d\x19\xa2\x15\x0d\x7c\x80\xb3\x0a\x70\x84\x99\x5d\x23\xe3\xc9\x87\xc7\xc6\xe8\xe3\x44\xbb\x46\x33\xa0\xb4\xc1\xc3\x14\xc4\x35\x9a\xfe\xf6\x48\x30\x44\x7d\x10\xdb\x59\x1d\xa2\xd8\xeb\x37\xf7\xda\xc8\xd0\x92\x89\x65\xad\xa8\x7b\x82\xe0\xcf\xb1\x51\x48\x1e\x43\xa4\x4f\x0d\xa4\xcf\x27\x93\x5e\xfc\x14\xfb\x3e\x38\xc5\x36\x71\x88\xa2\xa8\x80\xab\x37\x3e\x8a\x60\xc7\x3f\xd1\x1f\xea\x91\x93\x53\x40\x5d\x80\xbd\x76\x58\x48\x83\x27\x13\x5b\x16\xdd\x7a\x61\x53\xb8\xcb\x6a\xf7\xc0\x17\xce\xca\xf1\x2a\xd0\x6d\x3b\x80\x28\x42\xe0\x71\x80\xad\x90\x04\xe8\x01\x07\x4f\x10\xc9\xee\xdb\x37\xa7\x62\xd0\x64\xb9\x24\x56\xe3\x98\x53\xad\x29\xe4\x12\x13\x53\x44\x21\x93\xa3\x3e\x49\x82\x25\x94\x7c\x45\x03\x9b\x04\xaf\x10\x8c\x90\x15\x50\x2d\x8e\x86\x40\x45\x30\x64\x93\x10\x3b\x2e\x43\xff\x61\xd4\x5b\x88\xbd\xe2\x12\x1b\xe6\x36\xed\x95\x54\x6b\xea\x15\x46\x7e\x6d\x61\x8b\x8b\x90\x26\xc2\xe6\x1a\xb3\x35\x3f\xa6\x25\x79\x3f\x20\x0f\x0e\xdd\x32\x53\x3a\x31\x75\x52\x80\x3d\x86\x93\xec\x10\x87\x65\x87\xe3\x56\xfb\x34\x9a\x4f\x0c\x34\x28\x59\xd8\x87\x45\x4d\xde\x72\x29\xe3\xed\xa5\x28\xd7\xed\xb6\x53\x79\x4e\x40\x20\x59\xca\x26\x25\xb2\x5b\xdf\x56\x96\xdd\x2d\xa4\xf4\xe7\xc6\xa7\x01\xb8\xc5\x7c\x80\x78\x00\xa3\x0a\x97\xf3\xf2\x92\xa2\x90\xee\x80\xb7\x03\x09\x84\xbb\x22\x97\x84\x98\x3e\xa5\x2e\x7f\x34\x3a\x14\x4c\x10\x11\xc4\x3a\x1e\x86\xdd\x4b\x82\x07\x91\xc8\x06\x3f\x9a\xe1\xa3\xc9\x48\x68\x32\xe7\x4f\x55\x4a\xbc\x96\xf7\x61\xf3\x71\x10\x3a\x96\xe3\xe3\xe6\x93\x14\xdf\xc8\x3e\x65\xf1\x49\xa9\x6f\x78\x79\x0a\x39\xd4\x01\xa0\x02\xbc\xf9\x2b\xf3\xc3\x4c\xfb\x3a\xd7\xf4\x9b\x1a\x57\xe4\xc9\x67\xd2\x6a\x36\x62\x06\x33\x63\x74\x6f\xa0\x6f\x63\xe3\x33\x3a\x8f\x1f\x8c\x75\x50\x76\xa7\xe9\x06\xfa\xf8\x23\x7d\xa4\x4f\xd1\xdd\x58\xff\xe7\x68\x32\xd7\x76\xbf\x47\xdf\xf7\xbf\x6f\x46\x37\x9f\x35\x74\xde\x08\x51\x34\xfd\xa6\x6b\xb7\x60\x5b\xc2\x78\x34\x31\xb4\xfb\x03\x09\xef\x74\x4b\xc4\xcf\x1c\x5b\xca\xa5\xb5\x95\x2a\x3b\x50\xf3\x09\x52\x78\xe8\x46\xf5\x82\x95\x10\x8b\x8f\xa4\x67\x9e\x48\xc9\x23\x46\xb7\x81\x45\xb2\xb5\x2e\xc8\xfe\x59\xa6\xea\x74\x86\xc3\x8a\x84\xc2\xae\xc8\xd3\x6b\x31\x31\x88\xcc\xa8\xa6\x06\x95\x28\x3c\x27\x39\x88\xf0\x35\x9b\x1e\x24\x56\x5e\x2a\x41\x1c\x48\xf6\x99\x29\x42\x62\xad\x9a\x24\x44\x13\x6a\xd2\x44\x6e\x4a\x8b\x2b\x37\x5b\xad\x79\x80\xca\x85\x59\x5a\x8f\x49\xca\x3d\xd5\x4c\x52\x9f\x14\xb8\xb2\x7b\xd3\xe2\xca\x05\x0b\x37\xa2\xa8\xea\xfb\x4b\xea\x36\xa8\x80\x88\xf7\x40\x5c\x00\xc5\x6b\xdf\x60\x18\xaa\x28\x68\x35\x05\x83\x1b\xc8\xb5\x82\xa1\xc8\x0b\xa2\x61\xe6\xac\x3c\x1c\x6e\x41\x35\xc7\xed\x57\x6f\x4f\xff\xf5\xef\x7d\x36\xfe\xef\xff\x78\xf9\x18\x24\x4a\xe5\x1c\xd9\x50\x33\x3e\x15\xaa\xb9\x7b\xa7\xcb\x03\x37\xd4\x66\xf7\xbd\xae\xaa\x9a\x94\x19\xb8\x13\x9a\xff\xad\x07\x4d\x37\x78\xf1\x1d\x2c\xe0\x15\xa7\x85\x85\x0d\x96\x6e\x9e\xd4\xb8\xd2\x8e\x4f\xf6\xcb\x54\x9f\xc8\xce\x79\x94\xc8\xdf\x4c\x27\xf3\x3b\x3d\x8a\x69\x74\x2b\x90\xd1\xf4\xc0\xe1\x0f\xd8\xed\x76\x94\x4a\x0b\xf0\x47\x40\x56\x96\x8b\x19\x6b\x8f\x85\xf0\xd0\x3a\x88\x87\x24\xff\xf1\x99\xdc\x62\x58\x83\x4b\x1a\x28\x5c\x89\xa0\xdb\x91\x31\x92\x50\x1c\xeb\x33\x0d\x4e\x95\xb1\x6e\x4c\x2b\x17\x21\xf1\xb1\x31\x43\xdd\xce\xb9\xe9\x78\x4e\xe8\x40\x83\xc3\x62\x5d\x67\xec\x97\xdb\xe9\xa1\xce\xc5\xe0\xfc\x6d\x7f\xf0\xa6\x7f\x31\x40\x83\xab\xe1\xe5\xe5\xf0\x72\x70\x76\x7e\x79\xf9\xfa\xdd\x9b\xfe\xe0\x6f\x1d\x00\xad\xa4\xfd\x02\xb4\xdb\xe4\xb1\xe8\x82\x05\xb8\x87\x3a\x76\xbd\xa5\xab\xc1\xd5\x45\x6a\x49\xe0\x9e\xda\xab\x17\x15\xff\x48\xf4\xf2\x6e\x47\x1a\x50\xcb\xbb\x5e\x68\x40\xad\x42\xa7\x77\x88\x95\xe7\x34\x17\xb0\x31\x24\x56\x66\xda\x44\xbb\x31\x72\x57\x89\x67\xd0\xdb\x1e\x90\x07\x7a\xe8\xbc\x97\xdc\x22\xca\x97\x87\xa0\x99\x68\xc0\xe5\x4a\x55\xf4\xf1\x4e\x3f\xb4\x60\x6b\xc2\xed\xb2\xb4\x75\x88\xe3\x85\xe5\xd9\xe1\x2e\x29\xe5\x17\xd3\xff\x49\x9e\x32\x95\x37\x53\x7d\x66\xdc\x8f\x20\x0f\x1d\x54\xf6\x55\xf2\x7f\xc9\x46\x7c\x82\x8e\x6e\x6f\x73\xfa\xb9\x30\xd0\x97\xfb\xf1\xdd\xe8\xfe\x07\xfa\x87\xf6\x03\x75\x1d\xfb\xd0\x9b\x88\x36\xa8\xd4\x9b\xe4\x31\x53\x00\xa9\x4c\x54\xb8\x86\xda\xa4\x2a\x32\x5a\x47\xb6\x16\xa8\x94\xee\x62\x77\xf2\x64\x9c\xc6\xfa\xad\xf6\xfd\x98\xde\x23\x9e\x98\x53\x08\xd4\xf8\x9d\xc8\x7c\x36\xd6\xff\x8e\x16\x61\x40\x08\xea\xa6\xc2\xbd\x4a\xa9\xcf\x83\x1a\x75\x2c\xcd\xe1\x8c\xfb\x1f\x25\x90\xe5\xae\x89\x87\x2d\x39\x11\x9b\x43\x97\xe8\x53\xc3\x57\x6a\xd0\x7a\xd5\x5e\x8c\xbb\xce\x4d\x12\x55\x31\xf1\xf8\xb3\x71\xcf\xf5\x31\x64\xf0\x14\x7e\x49\x79\x9e\x44\xf6\x5e\xa5\x80\x9f\x77\x8b\xda\xcb\x5e\x91\x88\xa0\xef\x6b\xe5\x46\x41\x43\x4d\xac\x0a\x77\x7f\x5b\xd3\xe3\x5e\x04\x4b\x28\x50\xdf\xf4\xdb\x61\x91\x6a\xce\x13\x11\xb4\x35\x47\xf1\xe2\xd3\x81\x16\xb4\x25\x3a\xa9\x66\xc1\x5e\x38\x92\x50\xf1\x5a\xae\x4a\x09\x7c\x18\xe5\x08\xda\x00\xa3\x94\xca\x5e\xe3\xb1\x81\xa9\x0f\xc2\xee\xc5\x19\x58\x69\x3c\x0e\x45\xe5\x79\x02\xd9\x3b\xc1\x02\x62\x3e\xbe\xbc\xcf\xdb\x01\x59\xb1\xa0\x96\x40\x79\x70\xc3\x24\x5c\x61\x73\x0b\x60\xaf\xf1\xf8\xa5\x2c\x59\xb6\x49\xa3\x5a\x6e\x2a\x4d\x10\x4e\xdf\xae\x37\xeb\x71\xa9\xb9\x3c\xd1\xdd\x67\x00\xc5\x02\x20\x11\x3c\x80\x49\xd3\xcb\xa6\xce\x92\x1c\xbf\x34\x08\xe9\x11\x12\xe9\x8b\xae\xcb\x1a\x5a\x4c\xb5\x36\xa4\x27\x58\x24\x24\x81\x9d\x6e\xeb\x48\xe5\xee\x35\x78\x2b\xd8\x79\x86\xa4\xf9\x65\x27\xa9\xce\xa2\xdd\x65\x53\x30\x74\x4c\x7a\x14\xab\x2b\xbd\xe9\x6f\x3b\x08\x95\x2f\x0b\xa4\x64\x4a\x13\xd4\xa9\xe5\x3e\xf4\x78\xa1\xd8\xe4\x3f\x2d\x91\xf1\xca\xc9\xaa\x53\xe2\x7d\xc4\xf2\x42\xdc\xb8\xdf\xcf\xc8\x48\xf2\x26\xa9\xb3\xcd\x3a\x8e\x17\x62\xb8\x7b\x0d\x24\x63\x25\x6c\x22\x8b\xaa\xf7\xb7\x6a\xed\x27\x88\xb2\x2d\x6e\x0d\x78\x68\x9a\x28\x2a\x2d\xd6\x06\xad\xe4\x89\x3a\x83\x2a\x8c\x0e\x2a\x5f\x4a\xc6\xda\x3a\x3c\xab\x66\x94\x98\xc8\x8f\xd0\x7c\xbd\xd9\xfe\x02\xab\x5a\x3b\xba\xf6\x05\x61\x9b\xec\x8a\x8a\xac\x8d\x37\x17\x94\xfe\x6c\x28\x02\x35\x16\xa4\xc5\x4b\xb7\x9b\x7d\x71\xd2\xff\xf0\x01\x75\x18\x75\xa1\x10\x60\xd1\x77\x65\x51\x4c\x3a\xc3\x61\xf4\x02\xf4\xf4\xb4\x87\xc4\x82\x16\xb5\xd5\x04\x1d\xc6\xb6\x24\x10\x8b\x2e\xe8\x76\xb5\x0e\x95\xcc\x17\x44\xeb\x01\x14\x44\x4b\x10\x4e\xd1\xb7\xcf\xda\xbd\x96\x2c\x40\xf4\x1e\xbd\x7e\x9d\x0b\x9f\xe8\x4b\x6e\x64\xd1\x8d\xef\x92\x90\xc4\x91\xf8\x7f\x00\x00\x00\xff\xff\xe3\xf4\x95\x59\xf6\x2d\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 11766, mode: os.FileMode(420), modTime: time.Unix(1463228958, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x6d\x6f\xdb\x46\x12\xfe\xee\x5f\xb1\xc8\x17\xc9\x38\xf9\x2e\x41\x0e\x41\xce\x46\x02\x28\x36\x73\x11\x2a\x53\x89\x44\x35\x09\x8a\x82\x58\x91\x2b\x8a\x35\xc9\x65\x76\x49\xbf\xa4\xe8\x7f\xef\x2c\xdf\xdf\x96\xa4\x6c\xd2\x2d\x0a\xb4\xe2\xce\xce\xcc\x33\x33\xfb\xcc\x70\xe9\xb3\x33\xf4\x2f\xd7\xb6\x18\x0e\x08\xda\xfa\x27\x67\x67\xf0\x2f\xfa\x4c\x79\x60\x31\xb2\xf9\xb2\x44\x26\x0e\xf0\x0e\x73\x82\xcc\xd0\x8d\x96\x4f\x36\x8a\x86\x78\x00\xf2\x2e\xf1\x02\x3d\xb0\x5d\x42\xc3\x00\xbd\x43\x2f\x2f\xa2\x25\x87\x1a\x37\xf5\xa7\x86\x63\x0b\x69\xe2\x19\xd4\xb4\x3d\x0b\x16\x26\x5b\xed\xe3\xdb\xc9\x45\xaa\xce\x33\x31\x33\x75\x83\x7a\x7b\xca\x5c\x90\xd0\x79\xc0\xe0\x3f\x1c\x24\xa9\x97\xe8\x38\x10\x50\xbd\x0f\x3d\x23\xb0\xa9\xa7\xef\x40\x13\x11\xeb\x7b\xec\x70\x52\x32\x03\x0a\x74\x97\x70\x8e\xad\x48\xe0\x0e\x33\x0f\x74\x5d\x9c\x24\xf0\x54\xec\x92\x73\xe4\x3b\xbe\xc5\x7f\x38\x17\x48\x7b\xf0\xe1\xa7\xf2\x4d\x53\xd4\xcd\x62\xa5\x5e\xa0\x0d\x58\x72\xf1\x39\x3a\xbb\x40\xab\x3b\x8f\x30\xf8\xbf\x08\xf9\xe5\x5a\x99\x6b\x4a\x2e\x89\x16\x1f\x91\xba\xd2\xe0\xc1\x62\xa3\x6d\x52\x85\xe8\xeb\x42\xfb\x84\x36\x97\x9f\x94\xeb\x39\xf2\x2d\xdd\x80\x08\x3a\x54\x58\x2f\x99\xcf\xb5\x54\x1c\xb9\x5c\x5d\x5f\x2b\xaa\xd6\xe2\x46\x2c\x80\x60\x6b\x4d\x09\x5a\x6c\xd0\xe4\xf3\xf2\x3f\xbe\x25\x92\xe7\x33\x6a\x10\x33\x64\xd8\x41\x0e\xf6\xac\x10\xe2\x31\xa9\xfa\x71\xe0\x01\x65\x64\xb8\x28\xc4\xfa\xca\x41\x08\x77\x8e\x6d\xc8\x03\x50\x76\xe1\x71\xf8\x13\xb3\x02\xbe\x28\x59\x14\x80\x2e\x04\xb5\x84\xc4\x73\x51\x71\x9c\x04\x1c\xd1\x3d\x9a\xde\x90\x87\x19\xba\xc5\x4e\x48\x4e\x91\x8f\x6d\xc6\xa3\x90\x44\x65\x48\x30\x33\x0e\xba\x8f\x83\x03\x54\x4d\xec\xf5\xac\x9c\x42\x21\x66\x92\x3d\x0e\x1d\x28\x7d\xbc\x73\x08\xf7\xb1\x41\x44\x39\x4f\x2a\xab\x77\x76\x70\xd0\xa9\x6d\x16\x2a\xb4\x1c\x77\x5b\x78\xf6\xa0\x63\xc3\xa0\xa1\x17\xf0\x14\xbe\x36\xff\xb0\x54\x72\xf0\x49\xec\xb2\x08\x80\x58\x66\xf6\xbc\x98\x8f\x68\x5f\x4d\x2b\x9a\x9e\x20\xf8\xc7\x36\xd1\xce\xb6\x6c\x2f\x88\x32\xa5\x6e\x97\xcb\x59\xf4\x1c\x9b\x26\x83\x73\x02\x47\x0b\x33\x6c\x04\x84\x41\x60\xd8\x03\x84\x6b\xfa\xe6\xbf\xa7\x27\xa7\xb5\x5a\x49\xb4\x93\xfd\x9e\x18\x43\xbb\x9c\x28\x4d\x3c\xae\x00\xd1\x65\x08\x52\x39\xea\x13\xe0\x30\xc1\x0b\x32\xc9\x17\x94\x99\x84\xbd\x40\xb0\x42\x2c\x40\x5a\x5e\x8d\xea\xa5\x79\xc9\x24\x01\xb6\x1d\x8e\xfe\xe0\xd4\xdb\xc9\x83\xe2\x10\x13\xf6\x0e\x1c\x94\x44\x69\x12\x14\x4e\x7e\x84\x40\xa1\x32\x47\x63\x61\xfd\x80\xf9\xa1\x39\xa3\x15\x79\x9f\x91\x5b\x9b\x86\x5c\xef\xdc\x98\xc4\x88\x61\x8f\xe3\x98\x7d\xa3\xac\x64\x7e\x5c\x29\x1f\xe7\xdb\xa5\x86\x5e\x56\x2c\xe4\x59\xe9\x27\x6f\x38\x94\x13\x53\xc7\x01\x12\x1d\x04\xda\x82\xeb\x23\x71\x90\x44\x2f\x11\x4f\xd0\x4f\xea\x91\xea\x1e\x46\xa0\x19\x75\x6d\x8a\x65\x43\xdf\xec\x2d\x9b\xd5\x51\xf2\xd3\xf5\x29\x83\xb0\xe8\xb7\x90\x0f\x40\x54\xc3\xf2\xaa\x5a\x51\x14\x48\x03\x70\xdb\x1e\x6f\x2e\xc8\x3d\x21\xba\x4f\xa9\xd3\xbc\x2a\x9a\xae\x0e\x22\x92\x5c\x47\xcb\x70\x76\x09\xbb\x95\x89\xb8\xf8\x5e\x0f\xee\x75\x20\x3e\x9d\xdb\x3f\xeb\x52\xf2\x52\xce\xd3\xe6\x63\x16\xd8\x86\xed\xe3\xc1\x19\xaa\xd9\x46\xce\x57\xcd\x98\xfa\x1f\xf7\x6e\x02\x39\x16\x3f\xa8\x80\x60\xfe\x48\xc3\xb0\x51\xbe\x6c\x15\xf5\xb2\x25\x12\x45\xf0\xa9\x74\x3f\x1b\x11\x82\x8d\x36\x5f\x6b\x71\x23\x7d\x15\x3d\x58\xa8\xa0\x2c\x6a\x7d\x1f\xbe\x27\x8f\xd4\x15\xba\x5e\xa8\xbf\xce\x97\x5b\x25\xfb\x3d\xff\x96\xff\xbe\x9c\x43\x0b\x46\xaf\x06\x01\x8a\x56\x5f\x55\xe5\x0a\x6c\x77\x20\x9e\x2f\x35\x65\x7d\x24\xe0\x4c\x77\x87\xf8\xbf\x6d\xb3\x13\xcb\x58\x85\xda\xd5\x4c\x8b\xf4\x28\x6d\xb8\xbe\x0f\x3e\xc4\xb8\xa2\x7e\xf4\xc4\x76\x14\x3f\xe2\x34\x64\x06\x49\x4b\x5d\xc2\xfd\x29\x4f\x4d\x26\xe7\xe7\x35\x89\x1e\x87\xa2\x08\x6f\x3c\x5a\x90\x59\x89\x62\x2f\xa1\x85\xa6\xbd\xcd\x09\x78\x0a\x29\xc8\x3c\x1b\x96\x16\x3a\xac\x3c\x17\x31\x1c\x09\xf6\x89\xd4\xd0\x61\xad\x4e\x0e\xb2\x0d\x2d\xf4\x50\xd8\x32\x5e\xc9\xa6\x14\x51\xf4\xaf\xf7\x38\x96\x4c\x61\x1d\x43\x5e\x5f\x06\x69\x27\x83\x46\xd9\xdc\xb4\x7c\x5e\xc1\xd2\xd6\x2c\x9b\xf5\xfe\x91\x69\x0d\xe6\x1e\xe2\xdd\x12\x07\x9c\x42\x01\xb9\xaf\x51\xf5\xbd\x98\x9d\xe0\x35\x4d\xb2\xe8\x12\xf1\x0a\xd9\xb8\x24\xa2\x20\x5b\xe6\xb6\xe5\xe1\x20\x04\xd5\x0d\x61\xff\xdf\x9b\xd3\xdf\x7e\xcf\x59\xf8\xcf\xbf\x9a\x78\x18\x24\x2a\x43\x1c\x71\xa9\x1e\x75\x83\x3a\x67\x67\xba\x3c\x08\x43\x2b\xab\xe7\xba\xea\x6a\x12\x64\x10\x4e\x7d\x07\x89\x83\x17\x56\x88\xe2\x5b\x28\x60\x8b\x44\x64\x58\x3c\x4c\x70\xbc\x92\xa3\x93\xd8\xee\x75\xde\xe3\xe3\xb2\x52\x97\x5d\xdd\x1d\xc5\xf2\x97\xab\xe5\xf6\x5a\x15\x29\x15\x2f\xd4\x29\x4a\x0f\xe2\x0d\xaf\xed\xd3\x49\xaf\x81\x02\xc2\xc1\x88\x65\x38\x98\xf3\x1a\xa3\x0f\x86\x42\xda\xac\x8e\xc2\xd1\xc1\x7e\x6d\x48\x3a\x42\xe1\xdf\x90\x87\xfc\x5a\x45\xdd\x68\xeb\xf9\x42\x6d\x41\x5b\x27\xbc\x23\x13\x18\x95\xd2\xfc\xea\xaa\x60\xad\x8f\x8f\xe8\xf3\x7a\x71\x3d\x5f\x7f\x47\xbf\x28\xdf\xd1\xd4\x36\x8f\xef\xc1\x23\x22\x95\xd9\x6c\xc3\xda\xea\x67\x27\xda\x5d\x36\xa0\xa4\x90\x16\xea\x95\xf2\xed\x11\x8d\x2a\xda\x57\xd0\x27\xee\xcc\x1a\xdb\xd6\x76\xb3\x50\xff\x8f\x76\x01\x83\x17\xce\x69\x22\x3c\xab\xf5\x85\x26\x4f\x45\x7b\x1b\xcc\xcd\xa8\x57\xf6\xf2\xb1\xda\x61\x9b\x5c\x8b\x1b\xea\x60\xce\xc5\xea\xfa\xb9\x57\xe9\xe5\xb3\x7a\xdb\x6e\xac\x71\x1d\x38\xf8\x21\x5e\x7f\xaa\xdb\x5b\x75\x01\x53\x56\xe2\x7d\x45\x77\x11\x43\x7a\xed\x56\x72\xbf\xe9\x35\x7b\x96\xde\xa0\xc9\x3c\xcf\x69\x75\x48\x9f\x81\x3d\xfb\x7a\x9b\x4f\xf5\xb3\xc6\x8b\x82\x0e\x04\xd4\xd7\xfd\x51\x40\x24\x8a\x8b\x38\x24\xfd\xef\x51\xb0\xea\x68\xb2\x1b\x3d\x48\xf8\xd0\x80\xca\xba\x8b\x98\xd2\xbb\xca\x12\x88\x66\xf7\x8a\xa7\x77\x14\x1f\x6b\x06\xfa\x1d\xdb\x06\x6f\x6d\xcf\x24\xf7\x7a\xf5\x5e\x5d\x07\xbd\xc9\xe5\xf9\xa0\xae\x77\x5a\x2b\xe2\xc8\x2e\xf9\xcb\xec\x1d\x0b\x1e\x01\x64\xe0\xf0\xb7\x19\xea\x76\xbf\x33\x05\x09\x05\x08\x7d\x62\x2e\x1e\x86\xde\x5b\x4d\x74\x12\x90\x10\xea\xf0\x3a\x39\x1c\x42\x65\x76\xc9\x3d\x86\xeb\x4d\x76\x3a\x0f\x69\x26\xd9\x1f\xc4\xa8\x35\x53\xb2\xf3\x18\x8a\x91\xab\xab\xdc\xe2\x8f\x9c\x82\xda\x47\x83\x4e\x2c\x95\x0d\xfd\x91\x15\xbe\xe1\x3c\x4f\x66\x8a\x1f\x8d\xba\x60\x15\x64\xfb\x23\x6a\xfa\x3c\xf5\x3c\xd0\x1a\x3f\x8c\x75\x61\x6c\xda\xd4\x1f\x6c\x3a\x29\x3e\x0f\xc0\xec\xa2\xa7\x0b\x94\x74\xf2\x2f\xab\xce\xef\xc8\x47\xe7\x86\xaa\xa9\xc6\xa9\xea\x58\x86\x28\x2b\x2d\xdf\x23\x8f\x41\x11\x6d\xf6\xfa\x00\x2a\xef\x38\x0e\xdc\x48\x3d\xb3\x6e\xa5\x17\x90\xa6\xce\x19\x0d\xcd\xc1\xfd\x48\xd3\x78\xa2\x58\x32\x10\x3e\x72\x1e\xaf\x27\x44\x9e\x8f\xe2\xf8\x39\xfa\x71\xa9\x1b\x7b\xf4\x24\x0c\xc2\x26\xc9\x66\xa3\xf4\x5d\x52\xdf\x51\x7a\x33\x4c\x41\xb5\x18\xe8\x1c\xc1\xa6\xd3\xf4\xbb\xd8\xd9\xfb\xf7\x68\xc2\xa9\x03\xf3\x0c\x17\xdf\xbe\x45\x89\x4d\xce\xcf\xc5\x75\xed\xe9\xe9\x0c\xc9\x05\x0d\x6a\xf6\x13\xb4\x39\x0f\x09\x93\x8b\xee\x68\x68\x1d\x82\x5e\xe6\x4b\xa2\xed\x0e\x94\x44\x2b\x2e\x9c\xa2\xaf\x9f\x94\xb5\x12\x9f\x27\xf4\x0e\xbd\x7e\x5d\xc8\x9e\xec\xaf\xf9\x90\x41\x5d\xdf\x21\x01\x89\x32\x51\xfc\x43\xc0\x2b\x7a\xe7\x9d\x98\x8c\xfa\x28\xfa\x1b\xa7\xe6\x72\x31\x30\x37\x20\x5f\x17\x1d\x82\xe5\x03\xd5\xb6\xa9\xc0\x11\xbd\xc4\xfa\x6b\x4e\x5b\x5b\x9b\x4c\x5a\x55\x6d\x32\xd9\x1b\x4b\x26\xf4\x77\x00\x00\x00\xff\xff\x5d\xb2\x1f\x7d\x3f\x29\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1463228958, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_index_participants_by_toidSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xb1\x0a\xc2\x30\x10\x86\xf7\x7b\x8a\x1b\x15\xe9\x13\x74\x12\x1b\xa4\x4b\x2a\xd5\x82\x5b\x48\xdb\x60\x6e\x30\x17\x92\x03\xe9\xdb\x2b\x3a\xd8\xda\xc5\xf5\xf8\xf8\xfe\xfb\x8a\x02\x77\x77\xba\x25\x2b\x0e\xbb\x08\x70\x68\xd5\xfe\xa2\xb0\xd6\x95\xba\xa2\xe7\x68\xfa\xc9\x78\xa6\x11\x1b\x8d\x9e\xb2\x70\x9a\x0c\x47\xf7\xe2\x89\x83\x89\x36\x09\x0d\x14\x6d\x90\x8c\xdd\xb9\xd6\x47\xec\x25\x39\x87\x9b\x35\x4b\xe3\xb6\xfc\xd1\xcb\x47\x2f\x4b\xbd\x24\x1b\xb2\x1d\xfe\x1c\x98\xd3\xef\x09\x98\x27\x55\xfc\x08\x00\x55\xdb\x9c\xd6\x49\xe5\xe2\xfe\xfd\xa5\x84\x67\x00\x00\x00\xff\xff\x33\xec\x54\x7a\x15\x01\x00\x00")

func migrations2_index_participants_by_toidSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_index_participants_by_toidSql,
		"migrations/2_index_participants_by_toid.sql",
	)
}

func migrations2_index_participants_by_toidSql() (*asset, error) {
	bytes, err := migrations2_index_participants_by_toidSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1463228958, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_aggregate_expenses_for_accountsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x4b\xc3\x30\x1c\x47\xcf\xcb\xa7\xf8\x1f\x37\xdc\x40\x45\xbc\xec\x54\x6d\x84\x61\xed\x46\xe9\xc0\x9d\x42\x4c\x42\x1b\x5c\x93\xd2\xfc\x6b\x9d\x9f\xde\x6c\x2b\xa5\xe8\xb4\xcd\x31\xbc\xdf\x23\xb4\x6f\xb1\x80\xab\x42\x67\x15\x47\x05\xdb\x92\x90\xc7\x84\x06\x29\x85\x34\x78\x88\x28\x70\x21\x6c\x6d\x90\x39\xe4\xa8\x1d\x6a\xe1\x60\x4a\xc0\x1f\x2e\x65\xa5\x9c\x83\xfe\x11\x39\xaf\xb8\x40\x55\xc1\x07\xaf\x0e\xda\x64\xd3\xfb\xbb\x19\xc4\xeb\x14\xe2\x6d\x14\xcd\xcf\x3b\xe7\x14\x32\x61\xa5\xfa\x6f\x77\x73\xdb\xdb\x79\xe2\x34\x95\x5c\xef\x0f\x4c\x1b\x61\x0b\x05\x93\xc9\x9b\xce\xb4\xc1\x0e\x83\x90\x3e\x05\xdb\x28\x85\xeb\x79\x8f\xb6\x35\x8e\xc3\x1b\xa5\xde\x7f\xdb\x27\x03\x78\xab\x1f\xb4\x17\xd6\x60\xde\xe9\x47\xe3\xdd\xeb\x07\x78\x6e\x4c\xcd\xf7\x63\xed\x2d\x3d\xf6\xed\x75\x29\x7d\x1a\x92\x71\xf4\x9f\xe5\x78\x81\xba\x50\xbe\x87\xa2\x84\x46\x63\xee\x35\xa7\x1b\xf8\xb2\x46\xfd\xf8\xd9\x9b\x64\xf5\x12\x24\x3b\x78\xa6\xbb\x69\x1b\xcc\xbc\x57\xc0\x8c\xcc\x96\x5d\x6f\xab\x38\xa4\xaf\x17\x7a\x63\xed\x90\x69\xf9\x09\xeb\xf8\x62\x91\x2d\x72\xb4\xf5\x6b\x0e\x6d\x63\x08\x09\x93\xf5\x66\x94\x7d\x79\x46\xff\x0a\x7f\x49\xbe\x03\x00\x00\xff\xff\x90\xed\xa7\x46\x2a\x03\x00\x00")

func migrations3_aggregate_expenses_for_accountsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_aggregate_expenses_for_accountsSql,
		"migrations/3_aggregate_expenses_for_accounts.sql",
	)
}

func migrations3_aggregate_expenses_for_accountsSql() (*asset, error) {
	bytes, err := migrations3_aggregate_expenses_for_accountsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_aggregate_expenses_for_accounts.sql", size: 810, mode: os.FileMode(420), modTime: time.Unix(1463669016, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations4_account_statistics_updated_at_timezoneSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x89\x2f\x2e\x49\x2c\xc9\x2c\x2e\xc9\x4c\x2e\x56\x80\x48\x3b\xfb\xfb\x84\xfa\xfa\x29\x28\x94\x16\xa4\x00\xf5\xa5\xc4\x27\x96\x28\x84\x44\x06\xb8\x2a\x94\x64\xe6\xa6\x02\x15\xe7\x16\x28\x94\x67\x96\x64\x80\xb9\x0a\x55\xf9\x79\xa9\xd6\x5c\x5c\xc8\xf6\xb8\xe4\x97\xe7\x51\xd5\xa6\xfc\xd2\x12\x64\xcb\x00\x01\x00\x00\xff\xff\x99\xce\x1e\x1a\xd4\x00\x00\x00")

func migrations4_account_statistics_updated_at_timezoneSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations4_account_statistics_updated_at_timezoneSql,
		"migrations/4_account_statistics_updated_at_timezone.sql",
	)
}

func migrations4_account_statistics_updated_at_timezoneSql() (*asset, error) {
	bytes, err := migrations4_account_statistics_updated_at_timezoneSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/4_account_statistics_updated_at_timezone.sql", size: 212, mode: os.FileMode(420), modTime: time.Unix(1463750432, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_aggregate_expenses_for_accounts.sql": migrations3_aggregate_expenses_for_accountsSql,
	"migrations/4_account_statistics_updated_at_timezone.sql": migrations4_account_statistics_updated_at_timezoneSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_aggregate_expenses_for_accounts.sql": &bintree{migrations3_aggregate_expenses_for_accountsSql, map[string]*bintree{}},
		"4_account_statistics_updated_at_timezone.sql": &bintree{migrations4_account_statistics_updated_at_timezoneSql, map[string]*bintree{}},
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

