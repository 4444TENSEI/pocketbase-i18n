package core

// Model defines an interface with common methods that all db models should have.
//
// Note: for simplicity composite pk are not supported.
type Model interface {
	TableName() string
	PK() any
	LastSavedPK() any
	IsNew() bool
	MarkAsNew()
	MarkAsNotNew()
}

// BaseModel defines a base struct that is intended to be embedded into other custom models.
type BaseModel struct {
	lastSavedPK string

	// Id is the primary key of the model.
	// It is usually autogenerated by the parent model implementation.
	Id string `db:"id" json:"id" form:"id" xml:"id"`
}

// LastSavedPK returns the last saved primary key of the model.
//
// Its value is updated to the latest PK value after MarkAsNotNew() or PostScan() calls.
func (m *BaseModel) LastSavedPK() any {
	return m.lastSavedPK
}

func (m *BaseModel) PK() any {
	return m.Id
}

// IsNew indicates what type of db query (insert or update)
// should be used with the model instance.
func (m *BaseModel) IsNew() bool {
	return m.lastSavedPK == ""
}

// MarkAsNew clears the pk field and marks the current model as "new"
// (aka. forces m.IsNew() to be true).
func (m *BaseModel) MarkAsNew() {
	m.lastSavedPK = ""
}

// MarkAsNew set the pk field to the Id value and marks the current model
// as NOT "new" (aka. forces m.IsNew() to be false).
func (m *BaseModel) MarkAsNotNew() {
	m.lastSavedPK = m.Id
}

// PostScan implements the [dbx.PostScanner] interface.
//
// It is usually executed right after the model is populated with the db row values.
func (m *BaseModel) PostScan() error {
	m.MarkAsNotNew()
	return nil
}