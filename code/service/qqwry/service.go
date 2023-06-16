package qqwry

// ip数据库
type QQWRYService interface {
	Find(ip string) (WryFind, error)
}

var DefaultQQwry *QQwry

type stdQQWRYService struct {
}

func NewQQWRYService() QQWRYService {
	std := new(stdQQWRYService)
	return std
}

func (stdQQWRY *stdQQWRYService) Find(ip string) (WryFind, error) {
	return DefaultQQwry.Find(ip)
}
