//Package csvauto GENERATED BY CSV AUTO; DO NOT EDIT
package csvauto

//MEDALRUSHTIMECONFIG auto
type MEDALRUSHTIMECONFIG struct {
	ID          int   //#
	STARTDAY    int   //活动开启时间，每周的星期几；1代表星期1，以此类推
	STARTTIME   []int //24小时制，小时|分钟|秒
	JOINENDTIME int   //允许加入持续时间，单位为分钟；活动开启时间+允许加入持续时间=活动加入结束时间，在活动加入结束时间后，玩家不能再加入活动，持续到当期活动结束
	DURATION    int   //活动持续时间，单位为分钟；活动开启时间+活动持续时间即为活动运行结束时间
}

//IMEDALRUSHTIMECONFIG auto
type IMEDALRUSHTIMECONFIG interface {
	GetID() int
	GetSTARTDAY() int
	GetSTARTTIMELen() int
	GetSTARTTIMEByIndex(index int) int
	GetJOINENDTIME() int
	GetDURATION() int
}

//GetID auto
func (m *MEDALRUSHTIMECONFIG) GetID() int {
	return m.ID
}

//GetSTARTDAY auto
func (m *MEDALRUSHTIMECONFIG) GetSTARTDAY() int {
	return m.STARTDAY
}

//GetSTARTTIMELen auto
func (m *MEDALRUSHTIMECONFIG) GetSTARTTIMELen() int {
	return len(m.STARTTIME)
}

//GetSTARTTIMEByIndex auto
func (m *MEDALRUSHTIMECONFIG) GetSTARTTIMEByIndex(index int) int {
	return m.STARTTIME[index]
}

//GetJOINENDTIME auto
func (m *MEDALRUSHTIMECONFIG) GetJOINENDTIME() int {
	return m.JOINENDTIME
}

//GetDURATION auto
func (m *MEDALRUSHTIMECONFIG) GetDURATION() int {
	return m.DURATION
}
