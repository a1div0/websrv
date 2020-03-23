package websrv

import (
    "testing"
    "fmt"
    "io/ioutil"
    "net/http/httptest"
)

func TestResponse500(t *testing.T) {
    var p WebPages
    msg := "Проверка 500-й"
    w := httptest.NewRecorder()
    p.Response500(w, fmt.Errorf(msg))
    resp := w.Result()
    if (resp.StatusCode != 500) {
        t.Error("Код ответа = ", resp.StatusCode, ", а должен быть = 500\n")
    }

	body, _ := ioutil.ReadAll(resp.Body)
    if (string(body) != msg) {
        t.Error("Тело ответа = ", body, ", а должно быть:\n", msg)
    }
}

func TestPage(t *testing.T) {
    var p WebPages
    err:= p.Init("test_pages")
    if (err != nil) {
        t.Error(err)
    }

    w := httptest.NewRecorder()
    p.Page(w, "/1.html")

    resp := w.Result()
    if (resp.StatusCode != 200) {
        t.Error("Код ответа = ", resp.StatusCode, ", а должен быть = 200\n")
    }

	body, _ := ioutil.ReadAll(resp.Body)
    if (string(body) != "check") {
        t.Error("Тело ответа = ", body, ", а должно быть:\ncheck")
    }
}
