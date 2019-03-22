package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

//r任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

var (
	//单例对象
	G_apiServer *ApiServer
)

//初始化服务
func InitApiServer(err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	//启动TCP舰艇
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	//创建http服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiiReadTimeout) * time.Millisecond,
		Handler:      mux,
	}
	//赋值单例
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	//启动服务端
	go httpServer.Serve(listener)
	return
}
