package websrv

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

type BytesBuf []byte
type PagesMap map[string]BytesBuf

type WebPages struct {
    pages_map PagesMap // эта переменная хранит в памяти все файлы сайта
    server_path string // папка с файлами на сервере
}

func (p *WebPages) Response500(w http.ResponseWriter, err error) {
    w.WriteHeader(500)
    fmt.Fprintf(w, err.Error())
}

func (p *WebPages) Init(server_path string) (error) {
    p.pages_map = make(PagesMap)
    p.server_path = server_path
    return p.LoadDirToMemory("/")
}

// Функция загружает файлы в память
func (p *WebPages) LoadDirToMemory(path string) (error) {
    files, err := ioutil.ReadDir(p.server_path + path)
    if err != nil {
        return err
    }

    for _, f := range files {
        if (f.IsDir()) {
            err = p.LoadDirToMemory(path + f.Name() + "/")
            if err != nil {
                return err
            }
        }else{
            var full_name = path + f.Name()

            content, err := ioutil.ReadFile(p.server_path + full_name)
        	if err != nil {
        		return err
        	}

            p.pages_map[full_name] = content
        }
    }

    return nil
}

func (p *WebPages) Page(w http.ResponseWriter, file_name string) {

    var content BytesBuf

    content = p.pages_map[file_name]
    if (content == nil) {
        w.WriteHeader(404)
    }

    if (strings.HasSuffix(file_name, ".html")) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        if (content == nil) {
            content = p.pages_map["404.html"]
        }
    } else if (strings.HasSuffix(file_name, ".css")) {
        w.Header().Set("Content-Type", "text/css; charset=utf-8")
    } else if (strings.HasSuffix(file_name, ".js")) {
        w.Header().Set("Content-Type", "application/javascript")
    } else {

    }

    var content_len = len(content)
    var content_len_str = fmt.Sprintf("%d", content_len)

    w.Header().Set("Content-Length", content_len_str)
    w.Write(content)
}
