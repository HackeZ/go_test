package main

// === Usage 1 ===
// type P struct {
// 	X, Y, Z int
// 	Name    string
// }

// type Q struct {
// 	X, Y *int32
// 	Name string
// }

// func main() {
// 	var network bytes.Buffer
// 	enc := gob.NewEncoder(&network)
// 	dec := gob.NewDecoder(&network)
// 	// Encode (Send) the value.
// 	err := enc.Encode(P{3, 4, 5, "HackerZ"})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}

// 	// Decode (Receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}
// 	fmt.Println(q)
// 	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
// }
// func main() {
// 	var network bytes.Buffer
// 	enc := gob.NewEncoder(&network)
// 	dec := gob.NewDecoder(&network)
// 	// Encode (Send) the value.
// 	err := enc.Encode(&P{3, 4, 5, "HackerZ"})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}

// 	// Decode (Receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}
// 	fmt.Println(q)
// 	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
// }

// === Usage 1 ===

/*******************
 *******************
 *******分隔符*******
 *******************
 *******************/

// === Usage 2 ===
// type P struct {
// 	X, Y, Z int
// 	Name    interface{}
// }

// type Q struct {
// 	X, Y *int32
// 	Name interface{}
// }

// // 当编解码中有一个字段是 interface{} 的时候，
// // 需要对 interface{} 的可能产生的类型进行注册
// type Inner struct {
// 	Test int
// }

// func main() {
// 	var network bytes.Buffer
// 	enc := gob.NewEncoder(&network)
// 	dec := gob.NewDecoder(&network)

// 	gob.Register(Inner{})

// 	// Encode (Send) the value.
// 	inner := Inner{1}
// 	err := enc.Encode(P{1, 2, 3, inner})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}
// 	// Decode (Receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}

// 	// 这里使用了 gob.Register(Inner{}) 告诉系统：
// 	// 所有的 interface{} 是有可能为 Inner 结构的
// 	fmt.Println(q)
// 	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
// }

// === Usage 2 ===

/*******************
 *******************
 *******分隔符*******
 *******************
 *******************/

// === Usage 3 ===
// 这里我的P实现了GobEncoder接口
// 因此在enc.Encode的时候会调用
// func (this *P)GobEncode() ([]byte, error)

// 当然我这个函数直接返回的是空byte
// 因此在解码的时候会报错：
// decode error:gob: type mismatch in decoder: want struct type main.Q; got non-struct

// 这两个接口暴露出来就代表你为自己定义的结构进行编解码规则制定
//当然，如果使用自己的编解码规则，在编码和解码的过程就需要是一对的。
// type P struct {
// 	X, Y, Z int
// 	Name    string
// }

// func (this *P) GobEncode() ([]byte, error) {
// 	return []byte{}, nil
// }

// type Q struct {
// 	X, Y *int32
// 	Name string
// }

// func main() {
// 	var network bytes.Buffer
// 	enc := gob.NewEncoder(&network)
// 	dec := gob.NewDecoder(&network)
// 	// Encode (send) the value.
// 	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}
// 	// Decode (receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}
// 	fmt.Println(q)
// 	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

// }

// === Usage 3 ===
