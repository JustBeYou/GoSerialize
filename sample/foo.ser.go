package sample

// THIS FILE WAS GENERATED BY serialize
// PLEASE DO NOT EDIT

func (self Foo) Serialize() ([]byte, error) {
	var output, bytesTemp []byte

	bytesTemp = IntAsBytes(self.Bar)
	output = append(output, bytesTemp...)

	bytesTemp = UintAsBytes(self.Fizz)
	output = append(output, bytesTemp...)
	return output, nil
}