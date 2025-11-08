package model

//可攻略学姐详情
type Upclassman struct {
	Info         Info //可攻略学姐详情
	Group        string   //可攻略学姐组织
	Favorability int      //好感度
	Events       []Event  //会遇到的事件
}

//学姐事件
type Info struct {
	Name    string
	Age     int
	Details string //对学姐的描述
}


//随机事件
type Event struct {
	Name        string   //事件名称
	Description string   //事件名称
	Options     []Option //选项
}


//随机事件选项
type Option struct {
	Id                 int    //选项id(1,2,3)
	Content            string //选项内容
	FavorabilityChange int    //选项得分(好感度更改)
	Result             string //选项结果
}


//修改好感度
//func (u *Upclassman) UpdateFavorability(amount int) {
//	u.Favorability += amount