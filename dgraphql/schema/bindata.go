// Code generated by go-bindata. DO NOT EDIT.
// sources:
// query.graphql (815B)
// query_alpha.graphql (14B)
// registered_token.graphql (208B)
// schema.graphql (59B)
// serum_fill.graphql (1.462kB)
// serum_instruction.graphql (2.797kB)
// serum_market.graphql (110B)
// subscription.graphql (94B)
// subscription_alpha.graphql (21B)
// token.graphql (180B)

package schema

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xcd\x6e\xdb\x4c\x0c\xbc\xeb\x29\xc6\xba\x7c\x09\x10\x04\xf9\x80\xa2\x07\x5f\x03\x14\x6d\x0f\xfd\xb1\xd3\x53\xd1\xc3\x5a\x4b\x59\x84\xd7\x4b\x95\xa4\x6c\x08\x41\xde\xbd\x58\x09\x32\x10\x23\x45\x6f\x24\xc5\x99\xe1\x8c\xd6\xc7\x9e\xf0\x7d\x20\x1d\xf1\x5c\x55\x80\xcb\x81\xb2\xad\xf1\xf3\xa9\x14\x1b\xb2\x5e\xb2\xd1\xea\xd7\xea\xf2\xf1\x26\xc4\xa8\x64\xb6\xc6\xd6\x95\xf3\x7e\x75\xbb\xc6\xab\xe5\xb2\xa9\xb4\x67\x73\x52\x8a\x4f\x0b\xe1\xe6\xf5\xe8\x8a\xfa\x0a\xf0\x96\xc8\x5f\x08\x0a\xba\xae\xeb\x0d\xf9\xa0\x19\x29\x98\xe3\xff\x87\x07\xb4\x9c\x92\xa1\x15\x85\x77\x84\x3d\x9f\x28\xc3\x35\x44\xd2\xff\x0c\xfd\xb0\x4b\xdc\xe0\x40\xe3\x3d\x66\xa0\x81\x55\xe9\x44\x6a\xbc\x4b\x04\xc9\x69\x44\x0c\x1e\x60\x82\xc4\x27\x9a\x9b\x2c\x8e\x91\x1c\x7d\x50\x87\xb4\x08\xb0\x24\x0e\xef\x82\x83\x0d\x9c\x27\x31\x15\x71\x18\x39\xce\x9c\xd2\x84\xd9\x11\x74\x92\xa1\x88\x73\x47\x19\x4d\x48\x89\xe2\x7d\x5d\xd7\x15\x60\xa4\xc3\xf1\x03\xa7\xf4\x91\xcd\x45\xc7\x9b\xf9\xce\x8b\xf7\x3b\x1c\x83\x1e\xc8\x97\xc1\xed\x1a\xdb\x05\xf2\x28\x39\x53\xe3\x2c\x79\x55\xbd\x54\x55\x5d\xd7\x8f\x83\x9a\xa8\x41\xe9\xf7\xc0\x4a\x11\x2e\x68\x24\x3b\xe7\x81\x40\xec\x1d\x69\x89\xe5\x1c\x34\x42\x14\xbb\xd0\x1c\x4a\x6d\x68\x55\x8e\x08\x48\x6c\x93\xb7\x3e\xec\x39\x07\xa7\x08\x4a\x74\xa4\xec\x56\xae\x9d\xde\xcb\xb7\xb0\xa7\x4f\xb9\x95\xe7\x0a\x98\xc2\x6f\x26\xcd\x82\x2a\xfe\x5b\x56\xf3\x05\xb5\x0c\x0b\xed\x1d\x06\x23\xb0\x97\x93\x8c\x82\x36\xdd\x12\x99\xf4\xbd\x18\x3b\x21\xb2\xce\x76\xe6\x68\x00\xf3\xa0\x3e\x7b\xba\x04\x52\xbd\xad\x3b\xfd\xfb\x7f\xc8\x2e\x49\x2c\xf4\x94\xe3\x35\xf9\x4b\x55\x59\x13\x52\x50\xfc\xe0\xec\xef\xdf\x2d\xdd\xe7\xed\xd7\x2f\x7f\x02\x00\x00\xff\xff\x84\x0c\x85\x71\x2f\x03\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 815, mode: os.FileMode(0644), modTime: time.Unix(1608324834, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x26, 0x5f, 0x1c, 0xcf, 0x4c, 0x2c, 0x59, 0x3, 0xcc, 0xf2, 0xab, 0xd1, 0xf6, 0xdc, 0x54, 0xbb, 0xb3, 0x69, 0x49, 0x8e, 0xbe, 0x8c, 0xfd, 0x7, 0x3d, 0x6b, 0xa0, 0xd1, 0x42, 0x91, 0xbb, 0xfa}}
	return a, nil
}

var _query_alphaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x76\xca\x60\x3d\x0e\x00\x00\x00")

func query_alphaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_query_alphaGraphql,
		"query_alpha.graphql",
	)
}

func query_alphaGraphql() (*asset, error) {
	bytes, err := query_alphaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query_alpha.graphql", size: 14, mode: os.FileMode(0644), modTime: time.Unix(1608153826, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x58, 0xf, 0x86, 0xb0, 0x43, 0x44, 0x23, 0x5f, 0x97, 0xd3, 0xde, 0x25, 0xbd, 0x4b, 0x29, 0x22, 0xad, 0x9b, 0x95, 0xef, 0x8, 0x81, 0x45, 0x11, 0x3a, 0x12, 0x62, 0xab, 0x4c, 0x93, 0xf8}}
	return a, nil
}

var _registered_tokenGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8b\x31\x0e\xc2\x30\x0c\x45\xf7\x9e\xc2\xdc\x01\x31\x74\x63\x64\x2d\x70\x80\x96\x7c\x82\x45\x62\x47\xb1\x2b\x14\x10\x77\x47\x30\x65\x60\x7c\xff\xbd\xef\xad\x80\x26\x44\x36\x47\x45\x38\xe9\x1d\x32\xc1\x8a\x8a\x81\x5e\x03\xd1\x1c\x42\x85\xd9\x48\x47\xaf\x2c\x71\x33\x10\x65\x16\xdf\xaf\x7e\xd3\xca\xde\x7a\x71\xad\xc0\x13\x7f\x95\xad\xa5\xa4\x36\xd2\x99\xc5\x77\xdb\xef\x12\x70\xe1\x3c\x27\x1b\xe9\x20\xfe\x4b\x5a\x5e\x34\xf5\x27\x99\x33\x7a\x4e\x1a\xb5\xe7\x07\x16\x63\xef\x92\xf7\xf0\x09\x00\x00\xff\xff\x65\x08\x1e\x7a\xd0\x00\x00\x00")

func registered_tokenGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_registered_tokenGraphql,
		"registered_token.graphql",
	)
}

func registered_tokenGraphql() (*asset, error) {
	bytes, err := registered_tokenGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registered_token.graphql", size: 208, mode: os.FileMode(0644), modTime: time.Unix(1608153758, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0xad, 0xf8, 0x27, 0xfd, 0xab, 0x5a, 0x68, 0x60, 0x15, 0x89, 0x72, 0xbe, 0xcc, 0xcf, 0x4f, 0xcf, 0x8c, 0x4c, 0x3b, 0xe6, 0xcb, 0x89, 0x83, 0xe7, 0x3c, 0x5d, 0x1a, 0xf6, 0xe1, 0xec, 0xb5}}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xe2\xd2\xa4\xe2\xe4\xa2\xcc\x82\x92\xcc\xfc\x3c\x2b\x85\x60\x24\x1e\x57\x2d\x17\x20\x00\x00\xff\xff\x52\xd9\x58\xe5\x3b\x00\x00\x00")

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

	info := bindataFileInfo{name: "schema.graphql", size: 59, mode: os.FileMode(0644), modTime: time.Unix(1608156028, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5b, 0xba, 0x8f, 0x74, 0x29, 0x4e, 0xaf, 0x41, 0x66, 0x2b, 0x4e, 0x31, 0x85, 0x84, 0x19, 0x59, 0x50, 0x81, 0xda, 0x72, 0x50, 0x56, 0xaf, 0xe3, 0xd8, 0xb7, 0x35, 0xbc, 0xd1, 0x85, 0xf3, 0xc6}}
	return a, nil
}

var _serum_fillGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\x4f\x6f\xe3\x36\x10\xc5\xef\xfa\x14\xb3\xba\x6c\x02\xb8\x36\xd0\x6e\xf7\xa0\x9b\x77\xbb\x2d\x82\xd6\x40\x77\xe5\xa2\x87\x20\x07\x4a\x7c\x92\xd8\x50\xa4\x42\x0e\x93\x1a\x45\xbe\x7b\x41\x52\x8e\xff\xc4\xc1\x9e\x68\x93\x33\xbf\x79\x6f\x38\x54\x51\x96\xe5\x76\x00\x7d\xb6\xc6\xa0\x65\x65\x0d\xf1\x6e\x02\x75\xd6\x91\xa0\x1a\x2e\x8c\xbf\x2a\xad\x97\x54\x03\xc4\x03\xe8\xf6\x37\x27\xa6\xe1\xeb\x1f\xd4\x5a\x07\x6a\xad\x69\x31\xb1\xbf\xbb\x1a\x98\x27\x5f\xad\x56\xd2\xb6\x7e\x29\xbb\xe0\xb1\x54\x76\xd5\x07\x25\xe1\x57\x31\xf6\x87\x7d\xec\xaa\x8f\x84\x07\xbd\xba\x2e\xcb\xb2\x48\xd5\x5e\xea\x1c\xc9\xf8\xaf\x20\x2a\xd7\xa4\x95\x67\xb2\x1d\x41\xf6\xf0\xc4\xf6\x10\x5b\x16\x94\x77\x2b\xba\x7d\xd9\xfc\x22\x7b\xdc\xbd\x2b\x62\xee\x8d\xe9\xac\x1b\x45\xf6\x64\x49\x28\x49\x93\xe8\x95\x49\x3b\x31\x79\x12\x3d\x62\x50\x45\x7f\xce\xbf\xde\x15\xcf\x45\xec\x48\xb1\x26\xaf\x4c\xaf\x8f\x94\x11\x34\x46\x98\xa4\x25\xf6\xa1\x3d\x28\x8d\x12\x97\xc5\x6b\x33\x51\x4b\xb6\x51\x96\x5f\x03\xdc\x8e\xda\xe0\xbc\x75\x0b\x12\x5a\xdb\x27\x65\x7a\xda\xd9\x10\xb5\xb5\xd6\xb0\x32\x01\xd4\x81\xdb\x21\x1e\x38\xf8\xa0\xd9\x13\x1e\x61\x48\x74\x0c\x47\xca\x30\x9c\x0b\x53\xaa\x69\xbb\x98\xeb\x8e\x64\x2c\xe8\x49\xf1\x60\x03\xd3\xa8\x7c\x54\x4f\x82\x1a\x08\x5e\x46\x61\x34\x97\xae\xa8\x66\xa7\x4c\x9f\x3b\x94\xaf\xfe\x60\xd1\x36\xff\xa0\xdd\x27\x18\x2b\x51\x1d\x0e\x53\x6b\x4e\xfd\xed\xbd\x45\x88\x75\x12\x8e\x82\x51\x0f\x01\xa4\x24\x0c\xab\x4e\xc1\xa5\x39\x7a\x1a\x54\x3b\x10\x0f\xca\x1f\xe5\x2a\x4f\x8d\x0d\x46\x12\xdb\xb9\x60\x42\xdc\xc8\x4b\x12\xd9\x09\x09\xf7\xde\xd3\x14\x1a\xad\x5a\xba\xc7\x2e\x91\x4f\x99\x0b\x52\x4b\x2c\xd3\xed\x58\x13\xa7\x55\x30\x35\x61\x47\xd6\x91\x87\xd6\x87\x8c\x51\xb8\x7b\xec\x7d\x66\xf6\xa5\xaa\x39\xec\x4d\x0b\xef\xfd\xec\xfa\xb5\x95\x9c\x39\x77\x6f\x93\xfe\xcc\xdc\xbf\x13\xc8\x2b\x89\x0b\x0d\xe9\xe2\x6c\x40\xf1\x00\x47\x9f\x6e\x7e\xa1\xab\x28\x1b\xee\x3a\x3a\x58\xd7\xbf\xd3\x55\x13\x76\x70\xd7\x73\x91\x08\xa9\xa8\x56\x12\xdb\xdd\x84\x93\x76\xd9\x7b\x98\x73\xbe\x67\xe1\xd8\x53\xe7\xec\x38\x03\x1a\xe1\xb1\x8d\xa1\x15\xa5\xe5\x7b\x04\x18\xe9\x0f\x1e\x1f\x82\xe5\xb7\xd2\x4d\x18\x1b\xb8\x38\xa4\xda\xf2\x39\xa6\x53\x5a\x43\xe6\x6b\x4a\xfd\x8b\x73\xbb\x20\x87\xc9\xc1\xc3\xb0\x4f\x47\x13\x5c\x0b\xc3\xa2\xc7\xfe\xc5\xe5\xe0\x33\x98\x54\x72\xd6\xa3\x2d\x7f\xb6\xc1\x70\x45\x7f\x29\xc3\x1f\x3f\x1c\x37\x64\x72\xaa\x05\x4d\xf1\x13\xf0\x32\x04\xf8\x57\xb4\x7c\xf4\x85\xcb\x94\x14\x79\x09\x91\x03\x81\x6d\x9c\x6a\x8d\x47\xe8\x33\x5e\x77\x80\x74\x39\xac\x3a\x49\x4a\x0f\x08\x26\x8c\xa7\xa8\xf8\x86\x3e\xad\xeb\x2f\x05\x51\xfd\x6d\xf3\x63\x5e\x7e\xca\xcb\x87\xbc\xfc\x9c\x97\x8f\x05\xd1\xa6\xfe\xb6\x29\x9e\x8b\xff\x03\x00\x00\xff\xff\x1d\x92\xf8\x16\xb6\x05\x00\x00")

func serum_fillGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_fillGraphql,
		"serum_fill.graphql",
	)
}

func serum_fillGraphql() (*asset, error) {
	bytes, err := serum_fillGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_fill.graphql", size: 1462, mode: os.FileMode(0644), modTime: time.Unix(1608324834, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x29, 0x5f, 0xa5, 0x21, 0x4d, 0x3c, 0x15, 0x79, 0x21, 0x35, 0x1, 0x3c, 0x4, 0xb5, 0x23, 0x23, 0x7f, 0x2c, 0x1e, 0xcd, 0xaf, 0xb5, 0xbf, 0x42, 0xaf, 0xf2, 0xbf, 0x4, 0x10, 0x88, 0xed, 0x74}}
	return a, nil
}

var _serum_instructionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdd\x4e\xe3\x3a\x10\xbe\xcf\x53\x98\x77\x38\x3a\x17\x91\xce\x45\xff\x90\x22\xfa\x47\x53\x0e\x5a\x21\x84\xdc\x64\x68\x2d\x1c\x3b\xd8\x63\xa0\xec\xf2\xee\x2b\x3b\x4d\x63\x27\x69\x97\xdd\x25\x17\x95\xf2\xcd\x78\x32\xf3\x7d\x33\xe3\x46\xb8\x2f\x81\xa4\xa0\x4c\x91\x08\x8d\xca\x64\xc8\xa4\x58\x81\x2e\xa5\xd0\x40\xbe\x47\x84\xa0\x7a\x4b\xd9\x56\x50\x34\x0a\x62\x92\xa2\x62\x62\x7b\x51\xe1\x13\xa5\xa4\x8a\xc9\x50\x4a\x0e\x54\x58\x90\x35\x41\xe2\x4e\xd8\xe8\x23\x8a\x8c\x60\x52\x74\x2c\xe4\x3f\x72\x23\x72\xc8\x64\x0e\xb9\x0f\xff\xa8\x3d\x19\x32\xca\xd9\x3b\xcc\xa8\x7a\x02\xac\xf1\x39\xbc\x2e\x54\x0e\xaa\x7e\x9f\x51\xcc\x76\x01\x32\x92\x42\x9b\x02\x26\x2f\x20\x50\x1f\x41\x2a\x32\xe0\x81\x5f\x0a\x88\x1c\x2e\x8d\xc8\xfb\xbc\x86\xfb\x11\x67\x20\x30\xc9\xa3\x8a\xb0\xde\x64\x2d\x59\xa5\x92\x5b\x45\x8b\x64\x9c\x88\x1c\xde\x62\x72\xc3\x04\xfe\xfb\x8f\x65\x86\x66\x99\x34\x02\x47\xf6\x27\x26\xa4\x6b\xd1\x16\xb5\xcf\xdd\xc1\x74\x6f\x8d\x39\x45\x3a\x05\xb1\xc5\x5d\x65\x6e\x8e\x59\x4b\x7d\xc4\x3e\x8d\x32\x50\xc9\xd2\xb1\x7c\x44\x81\xdc\x21\xa7\x83\x43\x16\xae\x8c\xc2\x41\xbe\xda\xba\xe4\x23\xc9\xc4\x5a\x3e\x81\x68\xe1\x4b\xc5\x32\xe8\x18\x32\xc9\xc4\x8c\x89\x20\x48\x69\x3d\x43\xf0\x6c\x4e\x2e\x97\x0d\xd5\x30\x95\x98\xb2\x77\xf0\xf9\x7c\x36\x12\xfb\xf0\x47\x80\x15\x45\x18\x96\xda\x47\x5f\xa8\xe1\x68\xdb\x18\xd4\x5c\x8a\xac\x1b\x69\x6c\x34\xae\x77\x0a\xf4\x4e\xf2\xbc\xb1\x7a\xf2\x90\xf8\x3c\x71\x4d\x2d\x07\xa4\x6a\x08\xb3\xe1\x2c\xbb\x82\x7d\xec\x2b\xc4\x74\x95\x8b\x15\xc9\x9f\x1e\x7d\xab\x18\xd2\x0d\x07\x6f\xa8\x02\x86\xea\x8e\xef\x55\xeb\x00\xda\x48\xb2\x04\xe1\x1c\x75\x00\x2b\x78\x36\xa0\xf1\xda\x80\x81\xc0\x50\xd2\xbd\xcd\xc6\x8f\xf0\x2a\x5a\x88\x55\xf4\x7f\xcb\x63\x78\x32\xeb\x62\xba\xe4\xae\x1f\x96\xd5\x34\xb4\x52\x10\x2d\x67\x55\x8c\x99\xce\xaa\xb9\x38\xe0\xfd\x45\xbb\x62\x35\xcb\xe1\xd8\xdb\x29\xcb\x61\xbd\x2f\xc1\xc6\xe1\xac\x60\xe8\x7a\x31\xf6\xc7\xa4\xa0\x6f\xd7\x86\x0a\x64\xb8\xf7\x45\x97\x36\xa0\x3d\x6a\x63\x2d\xea\x17\x57\x66\x35\xea\xe3\xd8\x1f\xb7\x60\x4c\x7b\x95\x68\x09\xd5\xac\xa2\x5f\x49\x75\x52\x13\xb0\x4b\xab\x0b\x6f\x58\x1e\x8a\x4a\xf5\x93\xee\x08\x75\x09\xb0\x82\x0c\xd8\x4b\xd5\x4d\x81\x60\xa7\x6c\x27\x0a\x70\x89\x3b\x76\xcf\xce\x45\xb7\xe2\x56\xc4\x60\x17\x07\xac\xf8\xdd\x7a\x57\xa7\x73\x7f\xd1\xcf\xd7\x09\x5a\xbe\xa4\xe8\xf0\xba\x38\x5f\x77\x7c\xa6\xaa\x76\xd8\xe6\x2a\xf9\xe3\x6e\x68\xcf\xe3\xa9\x0f\x78\x43\xe2\x4f\x87\xeb\xf7\x24\xf7\xb7\xb1\xcf\x7a\x0f\x9a\x72\x89\x3d\x17\xd8\xb1\xee\x6e\x4d\xad\xa4\xbc\x3b\xb5\xf6\xf8\xbd\x6d\xf5\x37\x2b\xc8\x7a\xde\x52\xce\xa1\xed\xda\x03\xea\xc3\x2e\xfe\xf4\x06\x7b\x04\xa5\x40\x2d\x47\xad\x58\x27\xcb\x77\x9a\xb4\xe7\xa5\x87\x9e\xd3\xaa\x36\xff\x40\xbe\x7a\xf3\x7f\xb6\xaf\x9a\x0c\xdc\x97\x9b\x2d\xf9\x99\x0e\xe9\xa6\xef\x3e\x04\xc2\x14\xc7\x2e\x75\x61\x07\xe9\x55\x44\xc8\x30\x19\x47\x84\xdc\xcc\xaf\xe6\x8b\xdb\xf9\xd1\xf1\xb8\xa5\x9d\xe7\x34\x99\x25\xeb\x88\x90\x64\x36\x9b\x8c\x93\xc1\x7a\xf2\xb0\x58\x3d\x8c\x06\xf3\xd1\x64\x1a\x11\xb2\x5c\xa4\xeb\x87\xc5\x7c\xfa\xcd\x8f\xf3\x33\x00\x00\xff\xff\xc7\x80\xb1\xed\xed\x0a\x00\x00")

func serum_instructionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_instructionGraphql,
		"serum_instruction.graphql",
	)
}

func serum_instructionGraphql() (*asset, error) {
	bytes, err := serum_instructionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_instruction.graphql", size: 2797, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x0, 0x97, 0x87, 0xf5, 0xf2, 0x68, 0x9e, 0x91, 0xaa, 0x70, 0xd6, 0x26, 0x51, 0x2a, 0x72, 0x35, 0xb3, 0x2e, 0xd7, 0x53, 0xc, 0xc, 0x18, 0xc0, 0x7f, 0xb, 0xfc, 0xd1, 0x82, 0xca, 0xbc, 0xbc}}
	return a, nil
}

var _serum_marketGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x4e\x2d\x2a\xcd\xf5\x4d\x2c\xca\x4e\x2d\x51\x50\xa8\xe6\x52\x50\x48\x4c\x49\x29\x4a\x2d\x2e\xb6\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\xc8\x4b\xcc\x4d\x45\xf0\xb9\x94\x15\x14\x14\xf4\xf5\x15\x0a\x32\x12\x8b\x53\x15\x8c\xc0\xdc\xe4\xfc\xcc\xbc\x90\xfc\xec\xd4\x3c\x05\x30\x09\x16\x2b\x48\x46\x16\xa9\x05\x04\x00\x00\xff\xff\x43\x18\x5c\x95\x6e\x00\x00\x00")

func serum_marketGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_serum_marketGraphql,
		"serum_market.graphql",
	)
}

func serum_marketGraphql() (*asset, error) {
	bytes, err := serum_marketGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "serum_market.graphql", size: 110, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcd, 0x9f, 0xd3, 0x7, 0x2c, 0x57, 0x3c, 0xd4, 0x62, 0x3b, 0x83, 0xb4, 0xfc, 0x4b, 0x33, 0xe0, 0x26, 0x15, 0x4a, 0x7a, 0x5c, 0x6b, 0x7e, 0x97, 0x2e, 0x5, 0x41, 0xd3, 0x3a, 0x29, 0x57, 0x44}}
	return a, nil
}

var _subscriptionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\x28\x4e\x2d\x2a\xcd\xf5\xcc\x2b\x2e\x29\x2a\x4d\x06\x09\x7b\x64\x16\x97\xe4\x17\x55\x6a\x24\x26\x27\xe7\x97\xe6\x95\x58\x29\x04\x97\x14\x65\xe6\xa5\x2b\x6a\x5a\x29\x04\xa3\x29\x0d\x4a\x2d\x2e\xc8\xcf\x2b\x4e\xe5\xaa\xe5\x02\x04\x00\x00\xff\xff\xba\xf2\xbd\x88\x5e\x00\x00\x00")

func subscriptionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_subscriptionGraphql,
		"subscription.graphql",
	)
}

func subscriptionGraphql() (*asset, error) {
	bytes, err := subscriptionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subscription.graphql", size: 94, mode: os.FileMode(0644), modTime: time.Unix(1608153851, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x54, 0x96, 0xf5, 0x98, 0x89, 0xb2, 0xd0, 0xf9, 0x8a, 0x9b, 0x7f, 0xcc, 0xb5, 0x9e, 0xee, 0xf6, 0x5f, 0xf7, 0x11, 0xe4, 0xcc, 0x82, 0x8b, 0x4a, 0x4f, 0x67, 0x9a, 0x59, 0xee, 0x8f, 0x2b}}
	return a, nil
}

var _subscription_alphaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x4d\xe9\x40\xe8\x15\x00\x00\x00")

func subscription_alphaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_subscription_alphaGraphql,
		"subscription_alpha.graphql",
	)
}

func subscription_alphaGraphql() (*asset, error) {
	bytes, err := subscription_alphaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subscription_alpha.graphql", size: 21, mode: os.FileMode(0644), modTime: time.Unix(1608153609, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x20, 0x13, 0xe7, 0x39, 0xf6, 0x4e, 0x44, 0x50, 0x61, 0x33, 0x62, 0x95, 0x35, 0x1d, 0x5c, 0x43, 0x43, 0xb0, 0xe4, 0xe2, 0xf9, 0xa2, 0x94, 0xdd, 0xcc, 0xad, 0x3b, 0x5d, 0x4d, 0xce, 0x16}}
	return a, nil
}

var _tokenGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3b\xae\x02\x31\x10\x04\x73\x9f\xa2\xdf\x1d\x9e\x08\x9c\x11\x92\xf2\x39\x80\x85\x1b\x18\xb1\x1e\x5b\x9e\xd9\xc0\x20\xee\x8e\x36\x5a\x02\x08\xbb\x4b\x25\x95\x8f\x46\x1c\xeb\x9d\x8a\x67\x00\x52\xce\x9d\x66\x11\x07\xef\xa2\xd7\xbf\x00\x68\x2a\x5c\xf7\x2b\x84\x55\xd9\xd3\x5a\x55\xe3\x2f\xb5\x88\xfa\x76\xf6\x5b\xed\xe2\xe3\x13\x5c\x3a\xf9\xe0\x57\x64\x73\x6b\xd3\x88\x38\x89\xfa\xe6\x7f\x79\x32\xcf\x52\xd2\x64\x11\x3b\xf5\xa5\xe0\x1d\x00\x00\xff\xff\x26\x2a\xf1\x44\xb4\x00\x00\x00")

func tokenGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_tokenGraphql,
		"token.graphql",
	)
}

func tokenGraphql() (*asset, error) {
	bytes, err := tokenGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "token.graphql", size: 180, mode: os.FileMode(0644), modTime: time.Unix(1608237307, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2b, 0x2c, 0x9d, 0x30, 0xbf, 0x9f, 0xfb, 0x37, 0xc8, 0x67, 0xe1, 0x9e, 0xc6, 0xb, 0xc4, 0xa3, 0xd9, 0xb7, 0x8f, 0x7b, 0x43, 0xa4, 0xec, 0x19, 0x34, 0x6, 0x65, 0x6c, 0x22, 0xbb, 0x2c, 0x9}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"query.graphql":              queryGraphql,
	"query_alpha.graphql":        query_alphaGraphql,
	"registered_token.graphql":   registered_tokenGraphql,
	"schema.graphql":             schemaGraphql,
	"serum_fill.graphql":         serum_fillGraphql,
	"serum_instruction.graphql":  serum_instructionGraphql,
	"serum_market.graphql":       serum_marketGraphql,
	"subscription.graphql":       subscriptionGraphql,
	"subscription_alpha.graphql": subscription_alphaGraphql,
	"token.graphql":              tokenGraphql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"query.graphql": {queryGraphql, map[string]*bintree{}},
	"query_alpha.graphql": {query_alphaGraphql, map[string]*bintree{}},
	"registered_token.graphql": {registered_tokenGraphql, map[string]*bintree{}},
	"schema.graphql": {schemaGraphql, map[string]*bintree{}},
	"serum_fill.graphql": {serum_fillGraphql, map[string]*bintree{}},
	"serum_instruction.graphql": {serum_instructionGraphql, map[string]*bintree{}},
	"serum_market.graphql": {serum_marketGraphql, map[string]*bintree{}},
	"subscription.graphql": {subscriptionGraphql, map[string]*bintree{}},
	"subscription_alpha.graphql": {subscription_alphaGraphql, map[string]*bintree{}},
	"token.graphql": {tokenGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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