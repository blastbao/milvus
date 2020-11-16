package msgstream

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/golang/protobuf/proto"

	commonPb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalPb "github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
)

type InsertTask struct {
	Tag string
	InsertMsg
}

func (tt *InsertTask) Marshal(input *TsMsg) ([]byte, error) {
	testMsg := (*input).(*InsertTask)
	insertRequest := &testMsg.InsertRequest
	mb, err := proto.Marshal(insertRequest)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (tt *InsertTask) Unmarshal(input []byte) (*TsMsg, error) {
	insertRequest := internalPb.InsertRequest{}
	err := proto.Unmarshal(input, &insertRequest)
	testMsg := &InsertTask{InsertMsg: InsertMsg{InsertRequest: insertRequest}}
	testMsg.Tag = testMsg.PartitionTag

	if err != nil {
		return nil, err
	}
	var tsMsg TsMsg = testMsg
	return &tsMsg, nil
}

func newRepackFunc(tsMsgs []*TsMsg, hashKeys [][]int32) (map[int32]*MsgPack, error) {
	result := make(map[int32]*MsgPack)
	for i, request := range tsMsgs {
		if (*request).Type() != internalPb.MsgType_kInsert {
			return nil, errors.New(string("msg's must be Insert"))
		}
		insertRequest := (*request).(*InsertTask).InsertRequest
		keys := hashKeys[i]

		timestampLen := len(insertRequest.Timestamps)
		rowIDLen := len(insertRequest.RowIDs)
		rowDataLen := len(insertRequest.RowData)
		keysLen := len(keys)

		if keysLen != timestampLen || keysLen != rowIDLen || keysLen != rowDataLen {
			return nil, errors.New(string("the length of hashValue, timestamps, rowIDs, RowData are not equal"))
		}
		for index, key := range keys {
			_, ok := result[key]
			if !ok {
				msgPack := MsgPack{}
				result[key] = &msgPack
			}

			sliceRequest := internalPb.InsertRequest{
				MsgType:        internalPb.MsgType_kInsert,
				ReqID:          insertRequest.ReqID,
				CollectionName: insertRequest.CollectionName,
				PartitionTag:   insertRequest.PartitionTag,
				SegmentID:      insertRequest.SegmentID,
				ChannelID:      insertRequest.ChannelID,
				ProxyID:        insertRequest.ProxyID,
				Timestamps:     []uint64{insertRequest.Timestamps[index]},
				RowIDs:         []int64{insertRequest.RowIDs[index]},
				RowData:        []*commonPb.Blob{insertRequest.RowData[index]},
			}

			var msg TsMsg = &InsertTask{
				InsertMsg: InsertMsg{InsertRequest: sliceRequest},
			}

			result[key].Msgs = append(result[key].Msgs, &msg)
		}
	}
	return result, nil
}

func getMsg(reqID UniqueID, hashValue int32) *TsMsg {
	var tsMsg TsMsg
	baseMsg := BaseMsg{
		BeginTimestamp: 0,
		EndTimestamp:   0,
		HashValues:     []int32{hashValue},
	}
	insertRequest := internalPb.InsertRequest{
		MsgType:        internalPb.MsgType_kInsert,
		ReqID:          reqID,
		CollectionName: "Collection",
		PartitionTag:   "Partition",
		SegmentID:      1,
		ChannelID:      1,
		ProxyID:        1,
		Timestamps:     []Timestamp{1},
		RowIDs:         []int64{1},
		RowData:        []*commonPb.Blob{{}},
	}
	insertMsg := InsertMsg{
		BaseMsg:       baseMsg,
		InsertRequest: insertRequest,
	}

	testTask := InsertTask{
		InsertMsg: insertMsg,
	}
	tsMsg = &testTask
	return &tsMsg
}

func TestStream_task_Insert(t *testing.T) {
	pulsarAddress := "pulsar://localhost:6650"
	producerChannels := []string{"insert1", "insert2"}
	consumerChannels := []string{"insert1", "insert2"}
	consumerSubName := "subInsert"

	msgPack := MsgPack{}
	msgPack.Msgs = append(msgPack.Msgs, getMsg(1, 1))
	msgPack.Msgs = append(msgPack.Msgs, getMsg(3, 3))

	inputStream := NewPulsarMsgStream(context.Background(), 100)
	inputStream.SetPulsarCient(pulsarAddress)
	inputStream.CreatePulsarProducers(producerChannels)
	inputStream.SetRepackFunc(newRepackFunc)
	inputStream.Start()

	outputStream := NewPulsarMsgStream(context.Background(), 100)
	outputStream.SetPulsarCient(pulsarAddress)
	unmarshalDispatcher := NewUnmarshalDispatcher()
	testTask := InsertTask{}
	unmarshalDispatcher.AddMsgTemplate(internalPb.MsgType_kInsert, testTask.Unmarshal)
	outputStream.CreatePulsarConsumers(consumerChannels, consumerSubName, unmarshalDispatcher, 100)
	outputStream.Start()

	err := inputStream.Produce(&msgPack)
	if err != nil {
		log.Fatalf("produce error = %v", err)
	}
	receiveCount := 0
	for {
		result := (*outputStream).Consume()
		if len(result.Msgs) > 0 {
			msgs := result.Msgs
			for _, v := range msgs {
				receiveCount++
				fmt.Println("msg type: ", (*v).Type(), ", msg value: ", *v, "msg tag: ", (*v).(*InsertTask).Tag)
			}
		}
		if receiveCount >= len(msgPack.Msgs) {
			break
		}
	}
	inputStream.Close()
	outputStream.Close()
}