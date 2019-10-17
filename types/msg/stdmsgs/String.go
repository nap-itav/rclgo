package stdmsgs

// Automated generation with types/generator/app/generator.go
// This Generator is still under development,
// some imported comments may not make any sence,
// therefore it is recommended to look at the respective .msg/.srv original file.

// String Autogenerated Struct
type String struct {
	// Data Autogenerated Comment
	Data string
}

// ParseBytes TODO
func (msg *String) ParseBytes(data []byte) error {

	//stringData := string(data)

	msg.Data = string(data)
	//return String{Data: stringData}, nil
	return nil
}

// ToString TOOD
func (msg *String) ToString() string {
	return msg.Data
}