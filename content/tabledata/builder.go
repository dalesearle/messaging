package tabledata

import "playground/messaging"

type TableDataBuilder struct {
	Data     map[string]interface{}
	Database string
	PMS      string
	Table    string
	Version  string
}

func NewBuilder() *TableDataBuilder {
	return &TableDataBuilder{
		Data: make(map[string]interface{}),
	}
}

func (t *TableDataBuilder) AsImmutable() *TableData {
	return &TableData{
		data:     t.Data,
		database: t.Database,
		pms:      t.PMS,
		table:    t.Table,
		version:  t.Version,
	}
}

func (t *TableDataBuilder) SetData(col string, value interface{}) *TableDataBuilder {
	t.Data[col] = value
	return t
}

func (t *TableDataBuilder) SetDatabase(database string) *TableDataBuilder {
	t.Database = database
	return t
}

func (t *TableDataBuilder) SetPMS(pms string) *TableDataBuilder {
	t.PMS = pms
	return t
}

func (t *TableDataBuilder) SetTable(table string) *TableDataBuilder {
	t.Table = table
	return t
}

func (t *TableDataBuilder) SetVersion(version string) *TableDataBuilder {
	t.Version = version
	return t
}

func (t *TableDataBuilder) ContentID() messaging.ContentID {
	return messaging.TableData
}

func (t *TableDataBuilder) PackageType() messaging.PackageType {
	return messaging.GobPackage
}
