package server

import (
	"../../interface/network"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type server struct{}

type Server interface {
	Start(addr string)
	Get(endPoint string, apiFunc func(hc *HttpContext))
	Post(endPoint string, apiFunc func(hc *HttpContext))
}

func New() Server {
	return &server{}
}

func (s *server) Start(addr string) {
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func (s *server) Get(endPoint string, apiFunc func(hc *HttpContext)) {
	http.HandleFunc(endPoint, httpMethod(apiFunc, http.MethodGet))
}

func (s *server) Post(endPoint string, apiFunc func(hc *HttpContext)) {
	http.HandleFunc(endPoint, httpMethod(apiFunc, http.MethodPost))
}

func httpMethod(apiFunc func(hc *HttpContext), method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		// https://developer.mozilla.org/ja/docs/Web/HTTP/CORS
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// CORS プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}

		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		writer.Header().Add("Content-Type", "application/json")

		httpContext := HttpContext{
			ResponseWriter: writer,
			Request:        HttpRequest{Request: *request},
		}

		apiFunc(&httpContext)
	}
}

// interface以下で使用する通信プロトコル
type HttpContext struct {
	ResponseWriter http.ResponseWriter
	Request        HttpRequest
}

func (hc *HttpContext) GetResponseWriter() network.ResponseWriter {
	return hc.ResponseWriter
}

func (hc *HttpContext) GetRequest() network.Request {
	return &hc.Request
}

func (hc *HttpContext) SetRequestContext(ctx context.Context) {
	hc.Request.WithContext(ctx)
}

func (hc *HttpContext) GetRequestContext() context.Context {
	return hc.Request.Context()
}

func (hc *HttpContext) Success(jsonData interface{}) {
	data, err := json.Marshal(jsonData)
	if err != nil {
		log.Println(err)
		httpError(hc.ResponseWriter, http.StatusInternalServerError, "Failed Json Marshal")
		return
	}
	hc.ResponseWriter.Write(data)
}

func (hc *HttpContext) BadRequest(message string) {
	httpError(hc.ResponseWriter, http.StatusBadRequest, message)
}

func (hc *HttpContext) InternalServerError(message string) {
	httpError(hc.ResponseWriter, http.StatusInternalServerError, message)
}

func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		writer.Write(data)
	}
}

type HttpRequest struct {
	Request http.Request
}

func (hr *HttpRequest) GetBody() io.Reader {
	return hr.Request.Body
}

func (hr *HttpRequest) GetHeaderValue(Key string) string {
	return hr.Request.Header.Get(Key)
}

func (hr *HttpRequest) Context() context.Context {
	return hr.Request.Context()
}

func (hr *HttpRequest) WithContext(ctx context.Context) {
	hr.Request = *hr.Request.WithContext(ctx)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
