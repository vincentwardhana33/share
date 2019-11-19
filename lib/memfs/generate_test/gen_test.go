// Code generated by github.com/shuLhan/share/lib/memfs DO NOT EDIT.

package test

import (
	"github.com/shuLhan/share/lib/memfs"
)

func generate_() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata",
		Path:            "/",
		Name:            "/",
		ContentType:     "",
		ContentEncoding: "",
		Mode:            2147484141,
		Size:            4096,
	}
	return node
}

func generate_direct() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/direct",
		Path:            "/direct",
		Name:            "direct",
		ContentType:     "",
		ContentEncoding: "",
		Mode:            2147484141,
		Size:            0,
	}
	return node
}

func generate_direct_add() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/direct/add",
		Path:            "/direct/add",
		Name:            "add",
		ContentType:     "",
		ContentEncoding: "",
		Mode:            2147484141,
		Size:            0,
	}
	return node
}

func generate_direct_add_file() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/direct/add/file",
		Path:            "/direct/add/file",
		Name:            "file",
		ContentType:     "text/plain; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            22,
		V: []byte{
			84, 101, 115, 116, 32, 100, 105, 114, 101, 99, 116, 32, 97, 100, 100, 32,
			102, 105, 108, 101, 46, 10,
		},
	}
	return node
}

func generate_direct_add_file2() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/direct/add/file2",
		Path:            "/direct/add/file2",
		Name:            "file2",
		ContentType:     "text/plain; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            24,
		V: []byte{
			84, 101, 115, 116, 32, 100, 105, 114, 101, 99, 116, 32, 97, 100, 100, 32,
			102, 105, 108, 101, 32, 50, 46, 10,
		},
	}
	return node
}

func generate_exclude() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "/testdata/exclude",
		Path:            "/exclude",
		Name:            "exclude",
		ContentType:     "",
		ContentEncoding: "",
		Mode:            2147484141,
		Size:            0,
	}
	return node
}

func generate_exclude_index_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/exclude/index.css",
		Path:            "/exclude/index.css",
		Name:            "index.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            9,
		V: []byte{
			98, 111, 100, 121, 32, 123, 10, 125, 10,
		},
	}
	return node
}

func generate_exclude_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/exclude/index.html",
		Path:            "/exclude/index.html",
		Name:            "index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            14,
		V: []byte{
			60, 104, 116, 109, 108, 62, 60, 47, 104, 116, 109, 108, 62, 10,
		},
	}
	return node
}

func generate_exclude_index_js() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/exclude/index.js",
		Path:            "/exclude/index.js",
		Name:            "index.js",
		ContentType:     "application/javascript",
		ContentEncoding: "",
		Mode:            420,
		Size:            16,
		V: []byte{
			102, 117, 110, 99, 116, 105, 111, 110, 32, 88, 40, 41, 32, 123, 125, 10,
		},
	}
	return node
}

func generate_include() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/include",
		Path:            "/include",
		Name:            "include",
		ContentType:     "",
		ContentEncoding: "",
		Mode:            2147484141,
		Size:            0,
	}
	return node
}

func generate_include_index_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/include/index.css",
		Path:            "/include/index.css",
		Name:            "index.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            9,
		V: []byte{
			98, 111, 100, 121, 32, 123, 10, 125, 10,
		},
	}
	return node
}

func generate_include_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/include/index.html",
		Path:            "/include/index.html",
		Name:            "index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            14,
		V: []byte{
			60, 104, 116, 109, 108, 62, 60, 47, 104, 116, 109, 108, 62, 10,
		},
	}
	return node
}

func generate_include_index_js() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/include/index.js",
		Path:            "/include/index.js",
		Name:            "index.js",
		ContentType:     "application/javascript",
		ContentEncoding: "",
		Mode:            420,
		Size:            16,
		V: []byte{
			102, 117, 110, 99, 116, 105, 111, 110, 32, 88, 40, 41, 32, 123, 125, 10,
		},
	}
	return node
}

func generate_index_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/index.css",
		Path:            "/index.css",
		Name:            "index.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            9,
		V: []byte{
			98, 111, 100, 121, 32, 123, 10, 125, 10,
		},
	}
	return node
}

func generate_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/index.html",
		Path:            "/index.html",
		Name:            "index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            14,
		V: []byte{
			60, 104, 116, 109, 108, 62, 60, 47, 104, 116, 109, 108, 62, 10,
		},
	}
	return node
}

func generate_index_js() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/index.js",
		Path:            "/index.js",
		Name:            "index.js",
		ContentType:     "application/javascript",
		ContentEncoding: "",
		Mode:            420,
		Size:            16,
		V: []byte{
			102, 117, 110, 99, 116, 105, 111, 110, 32, 88, 40, 41, 32, 123, 125, 10,
		},
	}
	return node
}

func generate_plain() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "xxx/testdata/plain",
		Path:            "/plain",
		Name:            "plain",
		ContentType:     "text/plain; charset=utf-8",
		ContentEncoding: "",
		Mode:            420,
		Size:            22,
		V: []byte{
			84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 112, 108, 97, 105, 110, 32,
			116, 101, 120, 116, 46, 10,
		},
	}
	return node
}

func init() {
	memfs.GeneratedPathNode = memfs.NewPathNode()
	memfs.GeneratedPathNode.Set("/", generate_())
	memfs.GeneratedPathNode.Set("/direct", generate_direct())
	memfs.GeneratedPathNode.Set("/direct/add", generate_direct_add())
	memfs.GeneratedPathNode.Set("/direct/add/file", generate_direct_add_file())
	memfs.GeneratedPathNode.Set("/direct/add/file2", generate_direct_add_file2())
	memfs.GeneratedPathNode.Set("/exclude", generate_exclude())
	memfs.GeneratedPathNode.Set("/exclude/index.css", generate_exclude_index_css())
	memfs.GeneratedPathNode.Set("/exclude/index.html", generate_exclude_index_html())
	memfs.GeneratedPathNode.Set("/exclude/index.js", generate_exclude_index_js())
	memfs.GeneratedPathNode.Set("/include", generate_include())
	memfs.GeneratedPathNode.Set("/include/index.css", generate_include_index_css())
	memfs.GeneratedPathNode.Set("/include/index.html", generate_include_index_html())
	memfs.GeneratedPathNode.Set("/include/index.js", generate_include_index_js())
	memfs.GeneratedPathNode.Set("/index.css", generate_index_css())
	memfs.GeneratedPathNode.Set("/index.html", generate_index_html())
	memfs.GeneratedPathNode.Set("/index.js", generate_index_js())
	memfs.GeneratedPathNode.Set("/plain", generate_plain())
}
