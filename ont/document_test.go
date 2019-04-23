package ont

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestNew(t *testing.T) {
	doc, id, err := New()
	if err != nil {
		fmt.Println(err)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	d, err := json.Marshal(doc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("doc:%s\n", d)
	fmt.Println("PublicKey:[")
	for _, v := range doc.PublicKey {
		fmt.Printf("%v,\n", v)
	}
	fmt.Println("]")
	fmt.Printf("id:%v\n", id)

}
