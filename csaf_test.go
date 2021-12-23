package protocsaf

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"testing"
)

func TestLoad(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/test.json")
	assert.NoError(t, err)
	msg := &CsafMessage{}

	err = protojson.Unmarshal(data, msg)
	assert.NoError(t, err)

	_, err = protojson.Marshal(msg)
	assert.NoError(t, err)
}

func TestLoad2(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/test2.json")
	assert.NoError(t, err)
	msg := &CsafMessage{}

	err = protojson.Unmarshal(data, msg)
	assert.NoError(t, err)

	_, err = protojson.Marshal(msg)
	assert.NoError(t, err)
}

//func TestLoadLarge(t *testing.T) {
//	_, err := ioutil.ReadDir("./data/years")
//	if err != nil {
//		t.Skipf("please download the NVD dataset to run a full test")
//	}
//	data, err := ioutil.ReadFile("./data/years/2021.json")
//	assert.NoError(t, err)
//
//	msg := &NvdMessage{}
//	err = protojson.Unmarshal(data, msg)
//	assert.NoError(t, err)
//
//	assert.Len(t, msg.CveItems, int(msg.CveDataNumberOfCves))
//}
//
//func TestFullLoad(t *testing.T) {
//	dir, err := ioutil.ReadDir("./data/years")
//	if err != nil {
//		t.Skipf("please download the NVD dataset to run a full test")
//	}
//	assert.NoError(t, err)
//	count := 0
//	for _, file := range dir {
//		path := fmt.Sprintf("./data/years/%s", file.Name())
//		data, err := ioutil.ReadFile(path)
//		assert.NoError(t, err)
//
//		msg := &NvdMessage{}
//		err = protojson.Unmarshal(data, msg)
//		assert.NoError(t, err)
//
//		assert.Len(t, msg.CveItems, int(msg.CveDataNumberOfCves))
//		count = count + int(msg.CveDataNumberOfCves)
//	}
//}
//
//func BenchmarkLoadSingleJson(b *testing.B) {
//	count := 0
//	for i := 0; i < b.N; i++ {
//		msg := &NvdMessage{}
//
//		err := protojson.Unmarshal([]byte(str), msg)
//		assert.NoError(b, err)
//		count = count + int(msg.CveDataNumberOfCves)
//	}
//	b.StopTimer()
//}
//
//func BenchmarkSerializeSingleJson(b *testing.B) {
//	count := 0
//	msg := &NvdMessage{}
//	err := protojson.Unmarshal([]byte(str), msg)
//	assert.NoError(b, err)
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_, err = protojson.Marshal(msg)
//		assert.NoError(b, err)
//		count = count + int(msg.CveDataNumberOfCves)
//	}
//	b.StopTimer()
//}
//
//func BenchmarkLoadSingleProto(b *testing.B) {
//	count := 0
//	msg := &NvdMessage{}
//
//	err := protojson.Unmarshal([]byte(str), msg)
//	assert.NoError(b, err)
//
//	data, err := proto.Marshal(msg)
//	assert.NoError(b, err)
//
//	for i := 0; i < b.N; i++ {
//		err = proto.Unmarshal(data, msg)
//		assert.NoError(b, err)
//		count = count + int(msg.CveDataNumberOfCves)
//	}
//	b.StopTimer()
//}
//
//func BenchmarkSerializeSingleProto(b *testing.B) {
//	count := 0
//	msg := &NvdMessage{}
//	err := protojson.Unmarshal([]byte(str), msg)
//	assert.NoError(b, err)
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_, err = proto.Marshal(msg)
//		assert.NoError(b, err)
//		count = count + int(msg.CveDataNumberOfCves)
//	}
//	b.StopTimer()
//}
