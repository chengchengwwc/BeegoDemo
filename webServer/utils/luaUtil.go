package utils



type PlayerMessage struct {
	//玩家基本信息
	ClassName string
	Name string
	GuildName string
}


type GameMessage struct{
	PlayerMessageList []PlayerMessage
	ZoneName string
	DateTime string
	Points   float64
	ItemName string
	Awardedby string
	Foritem string
}




