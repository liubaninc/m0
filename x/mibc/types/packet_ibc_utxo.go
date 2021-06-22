package types

// ValidateBasic is used for validating the packet
func (p IbcUTXOPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcUTXOPacketData) GetBytes() ([]byte, error) {
	var modulePacket MibcPacketData

	modulePacket.Packet = &MibcPacketData_IbcUTXOPacket{&p}

	return modulePacket.Marshal()
}
