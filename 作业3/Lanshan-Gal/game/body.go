package game

import (
	"Lanshan-gal/model"
	"Lanshan-gal/server"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func Body() {
	fmt.Println("----欢迎来到游戏:YourGal----")
	time.Sleep(1 * time.Second)
	fmt.Println("**做烂勿喷**")
	time.Sleep(1 * time.Second)
	fmt.Println("欢迎来到“重庆最高学府”重庆邮电大学大学！在这片被“3G”雕像祝福的土地上,“红岩”的传说与代码的魔法交织。你,一个误入此地的新生,必须在“蓝山工作室”的传奇（内卷）与吉祥物“邮宝”那看穿一切的眼神中活下来,并……开始你的约会大作战")
	fmt.Println("请选择游戏模式:1.生成小故事    2.旮旯给木模式（未完成）")
	scanner.Scan()
	choi := scanner.Text()
	switch choi {
	case "1":
		story()
	case "2":
		fmt.Println("此模式未完成")
	}
}

//生成爱情小故事
func story() {
	var playerName string
	for {
		fmt.Print("请输入你的名字: ")
		scanner.Scan()
		playerName = scanner.Text()                // 获取玩家输入的那一行文本
		playerName = strings.TrimSpace(playerName) // 去掉开头和结尾的空格
		if playerName != "" {
			break
		}
		fmt.Println("名字不能为空,请重新输入。")
	}
	var playerDescription string
	for {
		fmt.Printf("很好,%s。请用一句话描述你自己,不超过20字: ", playerName)
		scanner.Scan()
		playerDescription = scanner.Text()
		playerDescription = strings.TrimSpace(playerDescription)
		if playerDescription != "" {
			break
		}
		fmt.Println("描述不能为空,请重新输入。")
	}
	
	player := model.Player{}
	player.Add(playerName, playerDescription)
	object := &model.Upclassman{}
Loop:
	for {
		fmt.Println("请输入你想要攻略的角色(1.康桥	2.cy	3.郭瑞彤	4.王家宽	5.贺一鸣   6.沈峻宇)")
		scanner.Scan()
		choi := scanner.Text()
		switch choi {
		case "1":
			object = KqInit()
			fmt.Println("你选择了康桥学姐")
			break Loop
		
		case "2":
			object = CyInit()
			fmt.Println("你选择了cy学姐")
			break Loop
		case "3":
			object = GrtInit()
			fmt.Println("你选择了郭瑞彤学姐")
			break Loop
		case "4":
			object = WjkInit()
			fmt.Println("你选择了王家宽学姐")
			break Loop
		case "5":
			object = HymInit()
			fmt.Println("你选择了贺一鸣学姐")
			break Loop
		case "6":
			object = SjyInit()
			fmt.Println("你选择了沈峻宇学姐")
			break Loop
		default:
			fmt.Println("输入无效,没有这个选项,请重新输入。")
			fmt.Println()
		}
	}
	var shortDescription string
	fmt.Printf("%v学姐的基础背景是：%v\n", object.Info.Name, object.Info.Details)
	time.Sleep(1*time.Second)
	for {
		fmt.Printf("请输入对%v学姐的拓展背景:", object.Info.Name)
		scanner.Scan()
		shortDescription = scanner.Text()
		if playerDescription != "" {
			resp := server.CreateStory(object,player, shortDescription)
			fmt.Println(resp)
			break
		}
		fmt.Println("不能为空")
	}
}
