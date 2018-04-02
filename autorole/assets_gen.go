// Code generated by "esc -o assets_gen.go -pkg autorole -ignore .go assets/"; DO NOT EDIT.

package autorole

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/assets/settings.html": {
		local:   "assets/settings.html",
		size:    3221,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/8xXzW7jNhC++ymmvHQXWFlIjwtbQNAstlsgTbEN0GNBSSOJDcVhyZFdQ9C7F6RkxXbs
uD/YYnOIHQ5nOPN985e+L7FSBkEU9jfZMTnSKIZhseh7xtZqyaOsQVkKWA7DYlWqDRRaer8WjrYiWwAA
HJ4WpBNdJzffTbIob272YitrTII9dCK7nd5cpc3NwW2bfVQbBOoYgtRD8K2VrAqp9Q4qctBim6Pzq9RO
HqSl2kxfv0kSSJezH5Ak2WKSn8QlNTr2U2SjmqPtqFCRa6FFbqhcC0ueBciCFZm1SFtpZI1p3y9vC1Yb
/NgpXS4/3Q1DOsP4EplnvP4OZqd3rDSoIf5OSqxkp/nk9lmNJKdyd+biBPT3jTQ1emjlDlg+IXQWmEBC
q0zHCDlW5BC4wR00coMgzQ6wqrDgGfvXfAgoJrWjzl7wISpomaMOvK7FHr9kBPH2iHnpvarNnnzgRnkY
8ydaeOUFjxoLBlWevnDkaEGGHWkBRra4Fp+febz00/fByoMNieHhKB0+x9Rd/qDqBj2Hv2C5T/goBPET
GYQ3pfIy11i+DaV3MYJ0DOEC5M/p/2XYKDsnQ4wiu4+J4YGqPQ+NsuDwj045LGNxBg14s8qzO+ULLVWL
7v0qzTP4VWkNhhi25J5AVbCjbkwrxeCRQ+qR0TuoQ/2PtU8Gfidl3l7nWBnbMfDO4lqYLrh2gd2jLJgj
25M+RXI3n1stC2xIl+jWQsBG6g7XIhT/zOaJzjCI/4unbHobQipRFSrV76FjghzBOvRoOADJDU6k7dFc
5S69XjbHyMTEnqFtO81quncMc5RYjdMdG6otQD/lSqw//0/q6z6YOVdkRml4yUaUfUUllX2qTWimFslq
hK3iJjJSkda0VaYeafs31IyW/ysxKlr5YrwcOPl1sHI4gRssnnL680rYqys96Gwvmm1PZD0YvXswP5Iy
AvpeVQcQPYuGIaph2fdoymHIIMj2EzB22G2DZpzMoUG+g5Jib429UzHksngKn4q/9eCwpU1oz47aoNOC
rBjdVrrSL18P+WrffYWSUXRx+eicQ8NzWPsSCI1rVVCJWd8vf3ZUoPfK1MOwSuPpfvAs4cPj7ftzVz88
3s6322lgvbm/uuNATgyepWN/4pJ/u7y879jskcAz2XfjEGtwDiOOeHFe9wxw547yjjn07phLvstbxXON
52wgZ5NYp1rpdvG7ruNHrql4EpD9Ije4SkcjGQAcGx+X3nG3DGvvBUfO7tTn9+5pgT58ZbFKQ0G+WL8r
Ig5Dejn+yxHzfPFXAAAA///a/D7flQwAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/assets": {
		isDir: true,
		local: "assets",
	},
}