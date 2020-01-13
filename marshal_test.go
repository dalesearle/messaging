package messaging

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

func TestPackageMarshall(t *testing.T) {
	var contentID ContentID
	var packageType PackageType
	var postmark int64
	var returnAddress uint32
	var vertical Vertical
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, TableData); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(buf, binary.BigEndian, GobPackage); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(buf, binary.BigEndian, time.Now().Unix()); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(buf, binary.BigEndian, uint32(65530)); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(buf, binary.BigEndian, ClassicVertical); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(buf, binary.BigEndian, []byte("now is the time for all good men to come to the aid of their country")); err != nil {
		t.Fatal(err)
	}

	rbuf := bytes.NewBuffer(buf.Bytes())
	fmt.Println(rbuf.Len())
	if err := binary.Read(rbuf, binary.BigEndian, uint16(contentID)); err != nil {
		t.Fatal(err)
	}
	fmt.Println(rbuf.Len())
	if err := binary.Read(rbuf, binary.BigEndian, byte(packageType)); err != nil {
		t.Fatal(err)
	}
	fmt.Println(rbuf.Len())
	if err := binary.Read(rbuf, binary.BigEndian, postmark); err != nil {
		t.Fatal(err)
	}
	fmt.Println(rbuf.Len())
	if err := binary.Read(rbuf, binary.BigEndian, returnAddress); err != nil {
		t.Fatal(err)
	}
	fmt.Println(rbuf.Len())
	if err := binary.Read(rbuf, binary.BigEndian, byte(vertical)); err != nil {
		t.Fatal(err)
	}
	//msg := make([]byte, rbuf.Len())
	fmt.Println(string(rbuf.Bytes()))
}
