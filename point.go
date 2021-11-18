package conf

type Point struct {
	DI   []*DI     `toml:"di"`   // di
	DO   []*DO     `toml:"do"`   // do
	IPMI []*IPMI   `toml:"ipmi"` // ipmi
	RTU  []*Modbus `toml:"rtu"`  // modbus rtu
	S7   []*S7     `toml:"s7"`   // s7
	Snmp []*Snmp   `toml:"snmp"` // snmp
	TCP  []*Modbus `toml:"tcp"`  // modbus tcp
}

// NewPoint new point
func NewPoint() *Point {
	return &conf.Point
}
