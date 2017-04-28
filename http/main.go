package main

import (
	"encoding/json"
	"log"
	"net/http/httptest"

	"bytes"
	"strings"

	"github.com/gin-gonic/gin"
)

type Man struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	m := Man{Name: "jack", Age: 12}
	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("can not marshal a man:", m)
	}

	var bf bytes.Buffer
	bf.Write(bs)

	// create a http response recorder
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", strings.NewReader(bf.String()))

	// create gin context and router
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	if err != nil {
		log.Fatalln("write body to context failed, error:", err)
	}

	var y Man
	err = ctx.BindJSON(&y)
	if err != nil {
		log.Fatalln("binding json in request failed, error:", err)
	}
	log.Println(y)
}
