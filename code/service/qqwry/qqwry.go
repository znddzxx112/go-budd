package qqwry

import (
	"encoding/binary"
	"fmt"
	"github.com/yinheli/mahonia"
	"net"
	"os"
	"sync"
)

const (
	INDEX_LEN       = 7
	REDIRECT_MODE_1 = 0x01
	REDIRECT_MODE_2 = 0x02
)

type QQwry struct {
	sync.RWMutex
	closed   bool
	filepath string
	file     *os.File
}

func NewQQwry(file string) (qqwry *QQwry, err error) {
	qqwry = new(QQwry)
	qqwry.filepath = file
	qqwry.closed = false
	qqwry.file, err = os.OpenFile(qqwry.filepath, os.O_RDONLY, 0400)
	if err != nil {
		return
	}
	return
}

func (wry *QQwry) Close() error {
	wry.Lock()
	defer wry.Unlock()
	err := wry.file.Close()
	if err != nil {
		return fmt.Errorf("wry.file.close():%s", err.Error())
	}
	return nil
}

type WryFind struct {
	Ip      string
	Country string
	City    string
}

func (wry *QQwry) Find(ip string) (WryFind, error) {
	wry.RLock()
	if wry.closed == true {
		wry.RUnlock()
		return WryFind{}, fmt.Errorf("%s", "ip地址库已关闭")
	}
	wrf := wry.find(ip)
	wry.RUnlock()
	return wrf, nil
}

func (wry *QQwry) find(ip string) WryFind {
	wf := WryFind{
		Ip: ip,
	}
	offset := wry.searchIndex(binary.BigEndian.Uint32(net.ParseIP(ip).To4()))
	// log.Println("loc offset:", offset)
	if offset <= 0 {
		return wf
	}

	var country []byte
	var area []byte

	mode := wry.readMode(offset + 4)
	// log.Println("mode", mode)
	if mode == REDIRECT_MODE_1 {
		countryOffset := wry.readUInt24()
		mode = wry.readMode(countryOffset)
		// log.Println("1 - mode", mode)
		if mode == REDIRECT_MODE_2 {
			c := wry.readUInt24()
			country = wry.readString(c)
			countryOffset += 4
		} else {
			country = wry.readString(countryOffset)
			countryOffset += uint32(len(country) + 1)
		}
		area = wry.readArea(countryOffset)
	} else if mode == REDIRECT_MODE_2 {
		countryOffset := wry.readUInt24()
		country = wry.readString(countryOffset)
		area = wry.readArea(offset + 8)
	} else {
		country = wry.readString(offset + 4)
		area = wry.readArea(offset + uint32(5+len(country)))
	}

	enc := mahonia.NewDecoder("gbk")
	wf.Country = enc.ConvertString(string(country))
	wf.City = enc.ConvertString(string(area))
	return wf
}

func (wry *QQwry) readMode(offset uint32) byte {
	wry.file.Seek(int64(offset), 0)
	mode := make([]byte, 1)
	wry.file.Read(mode)
	return mode[0]
}

func (wry *QQwry) readArea(offset uint32) []byte {
	mode := wry.readMode(offset)
	if mode == REDIRECT_MODE_1 || mode == REDIRECT_MODE_2 {
		areaOffset := wry.readUInt24()
		if areaOffset == 0 {
			return []byte("")
		} else {
			return wry.readString(areaOffset)
		}
	} else {
		return wry.readString(offset)
	}
}

func (wry *QQwry) readString(offset uint32) []byte {
	wry.file.Seek(int64(offset), 0)
	data := make([]byte, 0, 30)
	buf := make([]byte, 1)
	for {
		wry.file.Read(buf)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

func (wry *QQwry) searchIndex(ip uint32) uint32 {
	header := make([]byte, 8)
	wry.file.Seek(0, 0)
	wry.file.Read(header)

	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])

	// log.Printf("len info %v, %v ---- %v, %v", start, end, hex.EncodeToString(header[:4]), hex.EncodeToString(header[4:]))

	for {
		mid := wry.getMiddleOffset(start, end)
		wry.file.Seek(int64(mid), 0)
		buf := make([]byte, INDEX_LEN)
		wry.file.Read(buf)
		_ip := binary.LittleEndian.Uint32(buf[:4])

		// log.Printf(">> %v, %v, %v -- %v", start, mid, end, hex.EncodeToString(buf[:4]))

		if end-start == INDEX_LEN {
			offset := byte3ToUInt32(buf[4:])
			wry.file.Read(buf)
			if ip < binary.LittleEndian.Uint32(buf[:4]) {
				return offset
			} else {
				return 0
			}
		}

		// 找到的比较大，向前移
		if _ip > ip {
			end = mid
		} else if _ip < ip { // 找到的比较小，向后移
			start = mid
		} else if _ip == ip {
			return byte3ToUInt32(buf[4:])
		}

	}
	return 0
}

func (wry *QQwry) readUInt24() uint32 {
	buf := make([]byte, 3)
	wry.file.Read(buf)
	return byte3ToUInt32(buf)
}

func (wry *QQwry) getMiddleOffset(start uint32, end uint32) uint32 {
	records := ((end - start) / INDEX_LEN) >> 1
	return start + records*INDEX_LEN
}

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
