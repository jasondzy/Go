package service

import (
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
	"net/http"
	"fmt"
	"net"
	"time"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/callistaenterprise/goblog/accountservice/message"
	"github.com/callistaenterprise/goblog/accountservice/model"
)

var DBClient dbclient.IBoltClient
var MessagingClient message.IMessagingClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]   //这里通过mux库的功能来获取path传递的参数

	account, err := DBClient.QueryAccount(accountId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	account.ServedBy = getIP()

	notifyVIP(account)

	data, _ := json.Marshal(account) //将查询到的数据进行json序列化
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)

	w.Write(data) //将序列化后的数据传输到client端
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	panic("unable to determine local ip address (non loopback ).Exting")

}

func notifyVIP(account model.Account) {
	if account.Id == "10000" {
		go func(account model.Account) {

			vipNotification := model.VipNotification{AccountId: account.Id, ReadAt: time.Now().UTC().String()}
			
			data,_ := json.Marshal(vipNotification)
			err := MessagingClient.PublishOnQueue(data, "vipQueue")
			if err != nil {
				 fmt.Println(err.Error())
			}
		}(account)
	}
}