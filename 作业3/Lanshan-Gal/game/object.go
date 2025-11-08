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
