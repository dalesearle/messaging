package tabledata

import "playground/messaging"

type TableData struct {
	data     map[string]interface{}
	database string
	pms      string
	table    string
	version  string
}

func (t *TableData) AsBuilder() *TableDataBuilder {
	return &TableDataBuilder{
		Database: t.database,
		PMS:      t.pms,
		Table:    t.table,
		Version:  t.version,
		Data:     t.data,
	}
}

func (t *TableData) ContentID() messaging.ContentID {
	return messaging.TableData
}

func (t *TableData) Data() map[string]interface{} {
	return t.data
}

func (t *TableData) Database() string {
	return t.database
}

func (t *TableData) PackageType() messaging.PackageType {
	return messaging.GobPackage
}

func (t *TableData) Pms() string {
	return t.pms
}

func (t *TableData) Table() string {
	return t.table
}

func (t *TableData) Version() string {
	return t.version
}
