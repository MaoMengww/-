package game

import "Lanshan-gal/model"

func KqInit() *model.Upclassman {
	kq := model.Upclassman{
		Info: model.Info{
			Name:    "康桥",
			Age:     20,
			Details: "康桥学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的组长,作为技术代表,平日里保持着无可挑剔的完美,没有人知道她的真实面目",
			Group:   "蓝山工作室",
		},
		Favorability: 30,
		Events:       make([]model.Event, 10),
	}
	return &kq
}

func CyInit() *model.Upclassman {
	cy := model.Upclassman{
		Info: model.Info{
			Name:    "cy",
			Age:     20,
			Details: "cy学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的成员,平日里非常神秘,似乎永远也看不见她,传闻所有叫出她名字的人都消失了（可能是谣言）",
			Group:        "蓝山工作室",
		},
		Favorability: 0,
		Events:       make([]model.Event, 10),
	}
	return &cy
}

func WjkInit() *model.Upclassman {
	wjk := model.Upclassman{
		Info: model.Info{
			Name:    "王家宽",
			Age:     20,
			Details: "王家宽学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的成员,同时还是人工智能的大神,python组的名誉成员,两界都听说过她的传说",
			Group:        "蓝山工作室",
		},
		Favorability: 30,
		Events:       make([]model.Event, 10),
	}
	return &wjk
}

func HymInit() *model.Upclassman {
	hym := model.Upclassman{
		Info: model.Info{
			Name:    "贺一鸣",
			Age:     20,
			Details: "贺一鸣学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的成员,同时还是游戏大神，各种游戏都特别厉害。传说他的乌鲁鲁堵桥震慑整个学校",
			Group:        "蓝山工作室",
		},
		Favorability: 30,
		Events:       make([]model.Event, 10),
	}
	return &hym
}

func GrtInit() *model.Upclassman {
	grt := model.Upclassman{
		Info: model.Info{
			Name:    "郭瑞彤",
			Age:     19,
			Details: "贺一鸣学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的成员,qq顶着幸运星柊镜的头像,然而没有人会因此小看他",
			Group:        "蓝山工作室",
		},
		Favorability: 30,
		Events:       make([]model.Event, 10),
	}
	return &grt
}

func SjyInit() *model.Upclassman {
	grt := model.Upclassman{
		Info: model.Info{
			Name:    "沈峻宇",
			Age:     19,
			Details: "沈峻宇学姐是重庆最高学府重庆邮电大学蓝山工作室的前端组的组长,是一位长发美女,写的前端界面美轮美奂,跟本人一样魅",
			Group:        "蓝山工作室",
		},
		Favorability: 30,
		Events:       make([]model.Event, 10),
	}
	return &grt
}
