// Code generated by entc, DO NOT EDIT.

package route

const (
	// Label holds the string label denoting the route type in the database.
	Label = "route"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddr holds the string denoting the addr field in the database.
	FieldAddr = "addr"
	// FieldBits holds the string denoting the bits field in the database.
	FieldBits = "bits"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// Table holds the table name of the route in the database.
	Table = "routes"
	// HostTable is the table that holds the host relation/edge.
	HostTable = "routes"
	// HostInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	HostInverseTable = "nodes"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "node_routes"
)

// Columns holds all SQL columns for route fields.
var Columns = []string{
	FieldID,
	FieldAddr,
	FieldBits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "routes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"node_routes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// AddrValidator is a validator for the "addr" field. It is called by the builders before save.
	AddrValidator func(string) error
)
