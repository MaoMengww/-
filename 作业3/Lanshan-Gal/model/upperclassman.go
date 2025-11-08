package model

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)
//可攻略学姐详情
type Upclassman struct {
	Info         Info //可攻略学姐详情
	Group        string   //可攻略学姐组织
	Favorability int      //好感度
	Events       []Event  //会遇到的事件
}

type Info struct {
	Name    string
	Age     int
	Details string //对学姐的描述
}

type Event struct {
	Name        string   //事件名称
	Description string   //事件名称
	Options     []Option //选项
}

type Option struct {
	Id                 int    //选项id(1,2,3)
	Content            string //选项内容
	FavorabilityChange int    //选项得分(好感度更改)
	Result             string //选项结果
}

type people interface{
	Add(name string, discription string)
}

type User struct {
	Name string
	ShortDescription string
}

func (m User) Add(name string, description string){
	m.Name = name
	m.ShortDescription = description
}

//修改好感度
func (u *Upclassman) UpdateFavorability(amount int) {
	u.Favorability += amount
}
//修改描述
/*func (u *Upclassman) UpdateDetails(player User, shortDescription string) error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	userPrompt := fmt.Sprintf("你是一个galgame编剧，请你根据可攻略对象的基础背景和玩家给出的拓展背景，和玩家本身的描述写一段玩家和可攻略对象的galgame剧情，无选项，基础背景是%v,拓展背景是%v,玩家姓名时%v,玩家描述是%v", u.Info.Details, shortDescription, player.Name, player.ShortDescription)

	 resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(userPrompt),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }



	if err != nil {
		return fmt.Errorf("创建GenAI客户端失败:%v", err)
	}
	if err != nil {
		return fmt.Errorf("创建GenAI客户端失败:%v", err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-2.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(fmt.Sprintf(
					"请你根据用户提供的角色信息和拓展背景，为角色重新生成一段生动，详细，富有魅力的背景"+
					"内容应该包含他的名字，组织，性格，传闻。",
			)),
		},
	}

	model.GenerationConfig.Temperature = genai.Ptr[float32](1.5)
	model.MaxOutputTokens = genai.Ptr[int32](300)

	userPrompt := fmt.Sprintf("基础背景是%v,拓展背景是%v", u.Info.Details, shortDescription)
	
	 resp, err := model.GenerateContent(ctx, genai.Text(userPrompt))

	if err != nil {
		return fmt.Errorf("Gemini API 请求失败: %v", err)
	}
	if len(resp.Candidates) == 0{
		return fmt.Errorf("无内容")
	}
	// 【修复】增加对 Parts 列表的检查
    if len(resp.Candidates[0].Content.Parts) == 0 {
        return fmt.Errorf("AI 返回了候选，但内容部分(Parts)为空")
    }
	if len(resp.Candidates) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if txt, ok := part.(genai.Text); ok {
			u.Info.Details = string(txt)
			return nil
		}
	}

	if resp.PromptFeedback != nil && resp.PromptFeedback.BlockReason != 0 {
		return fmt.Errorf("AI 未能返回有效内容，可能被安全策略阻止: %s", resp.PromptFeedback.BlockReason.String())
	}

	return fmt.Errorf("AI 未能返回有效内容")
}*/


//随机增添事件


func (u *Upclassman) Tempgame(player User, shortDescription string) string {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, nil)
	userPrompt := fmt.Sprintf("你是一个galgame编剧，请你根据可攻略对象的基础背景和玩家给出的拓展背景和对玩家本身的描述写一段玩家和可攻略对象的galgame完整剧情，大于1000字，无选项，基础背景是%v,拓展背景是%v,玩家姓名时%v,玩家本身的描述是%v", u.Info.Details, shortDescription, player.Name, player.ShortDescription)

	 resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(userPrompt),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
	return resp.Text()
}