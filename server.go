package main

import (
    "github.com/ant0ine/go-json-rest"
    "log"
    "net/http"
)

// 入力の定義
type postHelloInput struct {
    Name string
}

// 出力の定義
type postHelloOutput struct {
    Result string
}

func postHello(w *rest.ResponseWriter, req *rest.Request) {
    input := postHelloInput{}
    err := req.DecodeJsonPayload(&input)

    // そもそも入力の形式と違うとここでエラーになる
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 適当なバリデーション
    if input.Name == "" {
        rest.Error(w, "Name is required", 400)
        return
    }

    log.Printf("%#v", input)

    // 結果を返す部分
    w.WriteJson(&postHelloOutput{
        "Hello, " + input.Name,
    })
}

func main() {
    handler := rest.ResourceHandler{}
    handler.SetRoutes(
        rest.Route{"POST", "/api/hello", postHello},
    )
    log.Printf("Server started")
    http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
