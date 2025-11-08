package server

//这个包以后会分


import (
	"Lanshan-gal/model"
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)


func  CreateStory(u *model.Upclassman,player model.Player, shortDescription string) string {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, nil)
	PlayerPrompt := fmt.Sprintf(
`你是一个专业的Galgame（视觉小说）编剧。
你的任务是创作一个“无选项”的、详细的、具有画面感的开场剧情章节（大于1000字）。
风格应为日式Galgame,包含细腻的心理活动、生动的场景描述和自然的角色对话。

---
## 剧情数据与要求：

**1. 攻略对象（必须融合以下两点）：**
  - 基础设定: %v
  - 关键拓展: %v

**2. 玩家（必须符合以下两点）：**
  - 姓名: %v
  - 性格/描述: %v

---
## 编剧指令：

请你基于以上所有数据,编写一个玩家和攻略对象的“首次深入互动”场景。
剧情必须：
1. 将攻略对象的“基础设定”和“关键拓展”融合成一个统一的、立体的形象,并在此基础上展开故事。
2. 玩家的心理活动、对话和行动必须符合他/她的“性格/描述”。
3. 包含清晰的地点和氛围描写（例如：蓝山工作室、放学后的教室、图书馆等）。
4. 这是一个无选项的章节,请用旁白(narrator)和对话推动故事。
---
现在,请开始创作。`,   
    u.Info.Details,     
    shortDescription,   
    player.Name,        
    player.ShortDescription, 
)

	 resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(PlayerPrompt),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
	return resp.Text()
}

//修改好感度
//func (u *Upclassman) UpdateFavorability(amount int) {
//	u.Favorability += amount
//}


//修改描述
/*func (u *Upclassman) UpdateDetails(player Player, shortDescription string) error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	PlayerPrompt := fmt.Sprintf("你是一个galgame编剧,请你根据可攻略对象的基础背景和玩家给出的拓展背景,和玩家本身的描述写一段玩家和可攻略对象的galgame剧情,无选项,基础背景是%v,拓展背景是%v,玩家姓名时%v,玩家描述是%v", u.Info.Details, shortDescription, player.Name, player.ShortDescription)

	 resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(PlayerPrompt),
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
					"请你根据用户提供的角色信息和拓展背景,为角色重新生成一段生动,详细,富有魅力的背景"+
					"内容应该包含他的名字,组织,性格,传闻。",
			)),
		},
	}

	model.GenerationConfig.Temperature = genai.Ptr[float32](1.5)
	model.MaxOutputTokens = genai.Ptr[int32](300)

	PlayerPrompt := fmt.Sprintf("基础背景是%v,拓展背景是%v", u.Info.Details, shortDescription)

	 resp, err := model.GenerateContent(ctx, genai.Text(PlayerPrompt))

	if err != nil {
		return fmt.Errorf("Gemini API 请求失败: %v", err)
	}
	if len(resp.Candidates) == 0{
		return fmt.Errorf("无内容")
	}
	// 【修复】增加对 Parts 列表的检查
    if len(resp.Candidates[0].Content.Parts) == 0 {
        return fmt.Errorf("AI 返回了候选,但内容部分(Parts)为空")
    }
	if len(resp.Candidates) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if txt, ok := part.(genai.Text); ok {
			u.Info.Details = string(txt)
			return nil
		}
	}

	if resp.PromptFeedback != nil && resp.PromptFeedback.BlockReason != 0 {
		return fmt.Errorf("AI 未能返回有效内容,可能被安全策略阻止: %s", resp.PromptFeedback.BlockReason.String())
	}

	return fmt.Errorf("AI 未能返回有效内容")
}*/

//随机生成事件及选项


