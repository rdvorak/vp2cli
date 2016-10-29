package loop2packet

import (
	"github.com/derekpitt/weather_crc"
)

type Packet []byte

type BarTrend int8

const (
	FallingRapidly = -60
	FallingSlowly  = -20
	Steady         = 0
	RisingSlowly   = 20
	RisingRapidly  = 60
	NoInformation  = 80
)

func (bt BarTrend) String() string {
	switch bt {
	case FallingSlowly:
		return "FallingSlowly"
	case FallingRapidly:
		return "FallingRapidly"
	case Steady:
		return "Steady"
	case RisingRapidly:
		return "RisingRapidly"
	case RisingSlowly:
		return "RisingSlowly"
	}

	return "NoInformation"
}

type Tempeture float32
type RainClick float32
type Mph float32

type Loop2Packet struct {
	BarometerTrend            BarTrend
	Barometer                 float32
	InsideTemperature         Tempeture
	InsideHumidity            int
	OutsideTemerature         Tempeture
	WindSpeed                 Mph
	WindDirection             int
	AverageWindSpeed10Minute  Mph
	AverageWindSpeed2Minute   Mph
	WindGust10Minute          Mph
	WindDirectionGust10Minute int
	DewPoint                  Tempeture
	OutsideHumidity           int
	HeatIndex                 Tempeture
	WindChill                 Tempeture
	THSWIndex                 Tempeture
	RainRate                  RainClick
	UVIndex                   int
	SolarRadiation            int
	StormRain                 RainClick
	DailyRain                 RainClick
	Last15MinuteRain          RainClick
	LastHourRain              RainClick
	DailyET                   float32
	Last24Rain                RainClick
}

// Some helper methods
func (packet Packet) get2ByteFloat(position int) float32 {
	return float32(int(packet[position+1])<<8 + int(packet[position]))
}

func (packet Packet) get2ByteInt(position int) int {
	return int(packet[position+1])<<8 + int(packet[position])
}

func (packet Packet) getTempeture(position int) Tempeture {
	return Tempeture(packet.get2ByteFloat(position) / 10)
}

func (packet Packet) getRainClick(position int) RainClick {
	return RainClick(packet.get2ByteFloat(position) / 100)
}

func (packet Packet) getMph(position int) Mph {
	return Mph(packet[position])
}

func (packet Packet) getMph2Byte(position int) Mph {
	return Mph(packet.get2ByteFloat(position) / 10)
}

func (packet Packet) stripACK() (p Packet) {
	p = packet
	if len(packet) > 0 && packet[0] == 0x6 {
		p = packet[1:]
	}

	return
}

func (p Packet) convertPacket() (sample Loop2Packet) {
	sample.BarometerTrend = BarTrend(int8(p[3]))
	sample.Barometer = p.get2ByteFloat(7) / 1000

	sample.InsideTemperature = p.getTempeture(9)
	sample.InsideHumidity = int(p[11])

	sample.OutsideTemerature = p.getTempeture(12)
	sample.OutsideHumidity = int(p[33])

	sample.WindSpeed = p.getMph(14)
	sample.WindDirection = p.get2ByteInt(16)
	sample.AverageWindSpeed10Minute = p.getMph2Byte(18)
	sample.AverageWindSpeed2Minute = p.getMph2Byte(20)
	sample.WindGust10Minute = p.getMph2Byte(22)
	sample.WindDirectionGust10Minute = p.get2ByteInt(24)

	sample.DewPoint = p.getTempeture(30) * 10
	sample.HeatIndex = p.getTempeture(35) * 10
	sample.WindChill = p.getTempeture(37) * 10
	sample.THSWIndex = p.getTempeture(39) * 10

	sample.UVIndex = int(p[43])
	sample.SolarRadiation = p.get2ByteInt(44)

	sample.RainRate = p.getRainClick(41)
	sample.StormRain = p.getRainClick(46)
	sample.DailyRain = p.getRainClick(50)
	sample.Last15MinuteRain = p.getRainClick(52)
	sample.LastHourRain = p.getRainClick(54)

	sample.DailyET = p.get2ByteFloat(56) / 1000
	sample.Last24Rain = p.getRainClick(58)

	return
}

func (packet Packet) IsValid() bool {
	if len(packet) != 99 {
		return false
	}

	// this is a werid thing...  the docs say that the byte at offset 96 should be 0xd,
	// but most of my packets are 0xa.. weird..   if i change it to a 0xd then the crc check
	// passes
	packet[96] = 0xd

	crc := weather_crc.New()
	crc.Write(packet)
	return 0 == crc.Sum16()
}

type loop2PacketError struct {
	s string
}

func (e loop2PacketError) Error() string {
	return e.s
}

func Decode(b []byte) (loopPacket Loop2Packet, err error) {
	packet := Packet(b).stripACK()
	if !packet.IsValid() {
		err = &loop2PacketError{"Packet not valid"}
		return
	}

	loopPacket = packet.convertPacket()
	return
}
