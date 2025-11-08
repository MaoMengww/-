package initialize

import "Lanshan-gal/model"

func KqInit() *model.Upclassman{
	kq := model.Upclassman{
		Info: model.Info{
			Name: "康桥",
			Age: 18,
			Details: "康桥学姐是重庆最高学府重庆邮电大学蓝山工作室的go组的组长，作为技术代表，平日里保持着无可挑剔的完美,没有人知道她的真实面目",
		},
		Group: "蓝山工作室",
		Favorability: 30,
		Events: make([]model.Event, 10),
	}
	return &kq
}