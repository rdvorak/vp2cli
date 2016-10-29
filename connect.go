package vp2cli

import (
	"github.com/derekpitt/weather_station/loop2packet"
	"github.com/tarm/goserial"
	"io"
	"sync"
)

const baudRate = 19200
const loop2Command = "LPS 2 1\n"
const loop2ResultLength = 100

var mutexi = make(map[string]*sync.Mutex)

type Session struct {
	dev string
	rw  io.ReadWriteCloser
	err error
}

func Connect() (ws Session, err error) {
	config := &serial.Config{Name: dev, Baud: baudRate}
	ws.dev = dev
	ws.rw, err = serial.OpenPort(config)
	mutexi[dev] = &sync.Mutex{}
	ws.err = err

	return ws, err
}

func (ws *weatherStation) GetSample() (packet loop2packet.Loop2Packet, err error) {
	if ws.err != nil {
		return packet, ws.err
	}

	mutexi[ws.dev].Lock()
	defer mutexi[ws.dev].Unlock()

	_, err = ws.rw.Write([]byte(loop2Command))
	if err != nil {
		return
	}

	buffer := make([]byte, loop2ResultLength)
	_, err = io.ReadFull(ws.rw, buffer)
	if err != nil {
		return
	}

	packet, err = loop2packet.Decode(buffer)
	if err != nil {
		return
	}

	return
}
