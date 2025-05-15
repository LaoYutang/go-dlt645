package dlt645_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/LaoYutang/go-dlt645"
	"github.com/tarm/serial"
)

func newTestClient() *dlt645.Dlt645Client {
	handler := dlt645.NewSerial2007Handler("COM13")
	handler.Baud = 2400
	handler.StopBits = 1
	handler.Parity = serial.ParityEven
	handler.Size = 8
	handler.ReadTimeout = 1 * time.Second
	handler.Logger = log.New(os.Stdout, "rs485: ", log.LstdFlags)
	client := dlt645.NewClient(handler)
	return client
}

func TestReadData(t *testing.T) {
	client := newTestClient()
	r, err := client.ReadData("000000000010", "02020100")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("result: %#v", r)
}

func TestSetParam(t *testing.T) {
	client := newTestClient()
	r, err := client.SetParam("000000000010", "04000104", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("result: %#v", r)
}

func TestSend(t *testing.T) {
	client := newTestClient()
	pdu := dlt645.NewCommonProtocolDataUnitByBytes(
		[]byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 0x13,
		[]byte{},
	)

	r, err := client.Send(pdu)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("result: %#v", r)
}
