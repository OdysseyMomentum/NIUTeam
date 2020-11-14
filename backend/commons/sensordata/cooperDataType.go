package sensordata

type CooperDataType struct {
	DeviceId         string  `json:"device_id"`
	Timestamp        uint64  `json:"timestamp"`
	Altitude         int     `json:"altitude"`
	Co2Concentration uint    `json:"co2_conc"`
	Humidity         string  `json:"humidity"`
	Illuminance      uint    `json:"illuminance"`
	MotionCount      uint64  `json:"motion_count"`
	Orientation      int     `json:"orientation,omitempty"`
	PressCount       int     `json:"press_count,omitempty"`
	Pressure         uint    `json:"pressure"`
	Rssi             int     `json:"rssi,omitempty"`
	Sequence         uint    `json:"sequence,omitempty"`
	SoundLevel       uint    `json:"sound_level"`
	Temperature      float32 `json:"temperature"`
	Type             string  `json:"type,omitempty"`
	UpTime           uint64  `json:"uptime,omitempty"`
	VocConentration  uint    `json:"voc_conc"`
	Voltage          float32 `json:"voltage,omitempty"`
}
