package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	KuaiDi100QueryURI = "http://poll.kuaidi100.com/poll/query.do"

	kuaidiAPIKey   = "aHpmOsSO3530"
	kuaidiCustomer = "C705DACD09FC967E59070D3FA2B1F0C2"

	KuaidiPullURI     = "http://poll.kuaidi100.com/poll/query.do"
	KuaidiSubscibeURI = "http://poll.kuaidi100.com/poll"
)

type KuaidiQueryObject struct {
	Customer string             `json:"customer"`
	Sign     string             `json:"sign"`
	Param    DeliveryQueryParam `json:"param"`
}

type DeliveryQueryParam struct {
	Com  string `json:"com"`
	Num  string `json:"num"`
	From string `json:"from"`
	To   string `json:"to"`
}

/// froad.com Key => aHpmOsSO3530

func main() {
	fmt.Println("Kuaidi 100 Query URI =>", KuaiDi100QueryURI)

	//postValue := make(url.Values, 30)
	//postValue["schema"] = []string{"json"}
	//postValue["param"] = []string{url.Values{
	//	"company": []string{"shentong"},
	//	"number":  []string{"1234567890"},
	//	"from":    []string{""},
	//	"to":      []string{""},
	//	"key":     []string{"ppppppppppppppp"},
	//	"parameters": []string{url.Values{
	//		"callbackurl": []string{"http://aaaaaa.bbbbbbb.cccc/callback?platfrom=2"},
	//		"salt":        []string{""},
	//		"resultv2":    []string{"1"},
	//	}.Encode()},
	//}.Encode()}

	//signObject := KuaidiQueryObject{
	//	Customer:"froad.com",
	//
	//}
	corpCode, shippingNo := "shentong", "3319627503942"
	postValue := make(url.Values, 20)
	postValue["customer"] = []string{kuaidiCustomer}
	postValue["sign"] = []string{signEncode(corpCode, shippingNo, kuaidiAPIKey, kuaidiCustomer)}
	//postValue["sign"] = []string{"E85EAE3DC27BCE3A2B6AFF68A2FA030F"}
	param, _ := encodeExpressPullParam(corpCode, shippingNo)
	postValue["param"] = []string{param}

	fmt.Println("postValue encode =>", postValue.Encode())

	requset, err := http.PostForm(KuaidiPullURI, postValue)
	if err != nil {
		log.Println("SubscribeExpress HTTP POSTFORM failed, error:", err)
		return
	}
	defer requset.Body.Close()

	rb, err := ioutil.ReadAll(requset.Body)
	if err != nil {
		log.Println("SubscribeExpress ioutil.ReadAll failed, error:", err)
		return
	}
	log.Println("[Obtain Express Result]: ", string(rb))
	//startTime := time.Now()
	//request, err := http.PostForm(KuaiDi100QueryURI, postValue)
	//if err != nil {
	//	log.Fatal("http.POST failed, error:", err)
	//	return
	//}
	//endTime := time.Now()
	//defer request.Body.Close()
	//
	//log.Println("request status code => ", request.StatusCode)
	//
	//log.Println("http.POST success, request:", request.StatusCode)
	//if request.StatusCode == http.StatusOK {
	//	rb, err := ioutil.ReadAll(request.Body)
	//	if err != nil {
	//		log.Fatal("ioutil.ReadAll failed, error:", err)
	//		return
	//	}
	//	fmt.Println("request content:", string(rb))
	//	fmt.Println("used time:", endTime.Sub(startTime).Seconds(), "s")
	//}
}

func signEncode(com, num, key, customer string) string {
	log.Println("Got signEncode start, req: com:", com, " num:", num, " customer:", customer)
	paramString, err := encodeExpressPullParam(com, num)

	var sign bytes.Buffer
	if err != nil {
		log.Println("Build signEncode Failed, error:", err)
		return ""
	}
	sign.WriteString(paramString)
	sign.WriteString(key + customer)
	fmt.Println("sign =>", sign.String())
	md5Ctx := md5.New()
	_, err = md5Ctx.Write(sign.Bytes())
	if err != nil {
		log.Println("Write bytes to md5Ctx error:", err)
	}
	md5Rslt := strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
	defer log.Println("signEncode md5 sum success, result:", string(md5Rslt))
	return md5Rslt
}

type expressPullParam struct {
	Com  string `json:"com"`
	Num  string `json:"num"`
	From string `json:"from"`
	To   string `json:"to"`
}

func encodeExpressPullParam(com, num string) (string, error) {
	epp := expressPullParam{
		Com:  com,
		Num:  num,
		From: "",
		To:   "",
	}
	rb, err := json.Marshal(epp)
	if err != nil {
		log.Println("encodeExpressPullParam JSON Marshal failed, error:", err)
		return "", err
	}
	log.Println("Got encodeExpressPullParam success, result:", string(rb))
	return string(rb), nil
}

//package main
//
//import (
//"encoding/json"
//"fmt"
//"net/http"
//"log"
//"io/ioutil"
//"bytes"
//)
//
//const (
//	KuaiDi100QueryURI = "http://poll.kuaidi100.com/poll/query.do"
//)
//
//type KuaidiQueryObject struct {
//	Customer string             `json:"customer"`
//	Sign     string             `json:"sign"`
//	Param    DeliveryQueryParam `json:"param"`
//}
//
//type DeliveryQueryParam struct {
//	Com  string `json:"com"`
//	Num  string `json:"num"`
//	From string `json:"from"`
//	To   string `json:"to"`
//}
//
//func main() {
//	fmt.Println("Kuaidi 100 Query URI =>", KuaiDi100QueryURI)
//
//	postValueObject := KuaidiQueryObject{
//		Customer: "froad.com",
//		Sign:     "sdqugei1u s1234",
//		Param: DeliveryQueryParam{
//			Com:  "yuantong",
//			Num:  "123456789",
//			From: "广东深圳",
//			To:   "北京朝阳",
//		},
//	}
//
//	postValueJSON, err := json.Marshal(postValueObject)
//	if err != nil {
//		fmt.Println("json.Marshal failed, error:", err)
//	}
//	fmt.Println("json.Marshal success, result:", string(postValueJSON))
//
//	body := bytes.NewBufferString(postValueJSON)
//	request, err := http.Post(KuaiDi100QueryURI, "application/json", body)
//	if err != nil {
//		log.Fatal("http.POST failed, error:", err)
//		return
//	}
//	defer request.Body.Close()
//
//	log.Println("request status code => ", request.StatusCode)
//
//	log.Println("http.POST success, request:", request.StatusCode)
//	if request.StatusCode == http.StatusOK {
//		rb, err := ioutil.ReadAll(request.Body)
//		if err != nil {
//			log.Fatal("ioutil.ReadAll failed, error:", err)
//			return
//		}
//		fmt.Println("request content:", string(rb))
//	}
//}
