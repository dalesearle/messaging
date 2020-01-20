package container

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"playground/messaging"
	"playground/messaging/content/tabledata"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var buf = new(bytes.Buffer)
	var cont = createContent()
	var err error
	var pkg = NewPackage().SetContent(cont)
	var read int
	var enc = NewEncoder(pkg)
	if err := enc.EncodeContainer(buf); err != nil {
		t.Fatal(err)
	}

	dec := NewDecoder(newMockConn(buf.Bytes()))
	pkg = NewPackage()
	if read, err = dec.DecodeContainer(pkg); err != nil {
		t.Fatal(err)
	}
	if read != buf.Len() {
		t.Fatalf("container length mismatch, expected %d, got %d", buf.Len(), read)
	}
	if cont, err = pkg.Content(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(cont)
}

func TestContentID(t *testing.T) {
	var id messaging.ContentID
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, uint16(1)); err != nil {
		t.Fatal(err)
	}
	wr := bytes.NewReader(buf.Bytes())
	if err := binary.Read(wr, binary.LittleEndian, &id); err != nil {
		t.Fatal(err)
	}
	if id != 1 {
		t.Fatalf("expeced 1, got %d", id)
	}
}

func createContent() messaging.Content {
	return tabledata.NewBuilder().
		SetDatabase("MSSQL").
		SetData("SOMEID", 13).
		SetData("SOMENAME", "Name").
		SetData("SOMETIME", time.Now()).
		SetData("SOMEFLOAT", 3.14).
		SetPMS("DENTRIXENTERPRISE").
		SetVersion("8.0.23.456")
}
