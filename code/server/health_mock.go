package server

func (dsm *defaultServerMock) healthMock() {
	dsm.PongMock()
}

func (dsm *defaultServerMock) PongMock() {

}
