//Package csvauto GENERATED BY CSV AUTO; DO NOT EDIT
package csvauto

//Robot auto
type Robot struct {
	ID                    int     //#
	Type                  int     //机器人类型（1.普通；2.俱乐部）
	GroupNum              int     //机器人组号
	GroupWeight           int     //机器人组内权重
	TargetType            int     //目标类型（1.关卡;2.收集物）
	InitialTime           []int   //首次开始行动条件（秒）
	InitialNum            []int   //首次生成机器人时的初始积分
	ActiviteRange         []int   //行动次数区间（左开右闭，最左侧从0开始，最右侧为无限大）
	TimeGap               [][]int //与行动次数区间对应的行动间隔（每个次数区间对应两个值，表示下限和上限；不同区间之间用;分隔）0到1 若有开始行动等待时间则按照开始等待时间来算
	LevelRange            [][]int //与行动次数区间对应的关卡步长（每个次数区间对应两个值，表示下限和上限；不同区间之间用;分隔）
	TargetRange           []int   //单关产生的收集物范围
	RobotSleepRule1       []int   //收集目标大于等于排名最高玩家该值时休眠
	RobotSleepRule2Time   []int   //活动时间占比（机器人初始化到活动结束时时间占比）
	RobotSleepRule2Target string  //时间分段对应的目标获取上限
}

//IRobot auto
type IRobot interface {
	GetID() int
	GetType() int
	GetGroupNum() int
	GetGroupWeight() int
	GetTargetType() int
	GetInitialTimeLen() int
	GetInitialTimeByIndex(index int) int
	GetInitialNumLen() int
	GetInitialNumByIndex(index int) int
	GetActiviteRangeLen() int
	GetActiviteRangeByIndex(index int) int
	GetTimeGapLen() int
	GetTimeGapByIndex(index int) []int
	GetLevelRangeLen() int
	GetLevelRangeByIndex(index int) []int
	GetTargetRangeLen() int
	GetTargetRangeByIndex(index int) int
	GetRobotSleepRule1Len() int
	GetRobotSleepRule1ByIndex(index int) int
	GetRobotSleepRule2TimeLen() int
	GetRobotSleepRule2TimeByIndex(index int) int
	GetRobotSleepRule2Target() string
}

//GetID auto
func (r *Robot) GetID() int {
	return r.ID
}

//GetType auto
func (r *Robot) GetType() int {
	return r.Type
}

//GetGroupNum auto
func (r *Robot) GetGroupNum() int {
	return r.GroupNum
}

//GetGroupWeight auto
func (r *Robot) GetGroupWeight() int {
	return r.GroupWeight
}

//GetTargetType auto
func (r *Robot) GetTargetType() int {
	return r.TargetType
}

//GetInitialTimeLen auto
func (r *Robot) GetInitialTimeLen() int {
	return len(r.InitialTime)
}

//GetInitialTimeByIndex auto
func (r *Robot) GetInitialTimeByIndex(index int) int {
	return r.InitialTime[index]
}

//GetInitialNumLen auto
func (r *Robot) GetInitialNumLen() int {
	return len(r.InitialNum)
}

//GetInitialNumByIndex auto
func (r *Robot) GetInitialNumByIndex(index int) int {
	return r.InitialNum[index]
}

//GetActiviteRangeLen auto
func (r *Robot) GetActiviteRangeLen() int {
	return len(r.ActiviteRange)
}

//GetActiviteRangeByIndex auto
func (r *Robot) GetActiviteRangeByIndex(index int) int {
	return r.ActiviteRange[index]
}

//GetTimeGapLen auto
func (r *Robot) GetTimeGapLen() int {
	return len(r.TimeGap)
}

//GetTimeGapByIndex auto
func (r *Robot) GetTimeGapByIndex(index int) []int {
	return r.TimeGap[index]
}

//GetLevelRangeLen auto
func (r *Robot) GetLevelRangeLen() int {
	return len(r.LevelRange)
}

//GetLevelRangeByIndex auto
func (r *Robot) GetLevelRangeByIndex(index int) []int {
	return r.LevelRange[index]
}

//GetTargetRangeLen auto
func (r *Robot) GetTargetRangeLen() int {
	return len(r.TargetRange)
}

//GetTargetRangeByIndex auto
func (r *Robot) GetTargetRangeByIndex(index int) int {
	return r.TargetRange[index]
}

//GetRobotSleepRule1Len auto
func (r *Robot) GetRobotSleepRule1Len() int {
	return len(r.RobotSleepRule1)
}

//GetRobotSleepRule1ByIndex auto
func (r *Robot) GetRobotSleepRule1ByIndex(index int) int {
	return r.RobotSleepRule1[index]
}

//GetRobotSleepRule2TimeLen auto
func (r *Robot) GetRobotSleepRule2TimeLen() int {
	return len(r.RobotSleepRule2Time)
}

//GetRobotSleepRule2TimeByIndex auto
func (r *Robot) GetRobotSleepRule2TimeByIndex(index int) int {
	return r.RobotSleepRule2Time[index]
}

//GetRobotSleepRule2Target auto
func (r *Robot) GetRobotSleepRule2Target() string {
	return r.RobotSleepRule2Target
}