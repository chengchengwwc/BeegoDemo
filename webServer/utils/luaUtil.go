package utils



import (
	"log"
	"reflect"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gluamapper"
)


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

type LuaMessage struct{
	GameMessageList []GameMessage
}





func ParseMap(targetMap map[interface{}]interface{},gp *GameMessage){
	// 递归解析
	defer func(){
			err := recover()
			if err != nil{
				log.Println(err)
			}
		}()

	PlayerMessage := PlayerMessage{} 
	for key,value := range targetMap{
		res,ok := value.(map[interface{}]interface{})
		if ok{
			ParseMap(res,gp)
		}else{
			if(key == "Class"){
				PlayerMessage.ClassName = value.(string)
			}else if(key == "Name"){
				PlayerMessage.Name = value.(string)
			}else if(key == "Guild"){
				PlayerMessage.GuildName = value.(string)
			}else if(key == "Zone"){
				gp.ZoneName = value.(string) 
			}else if(key == "Date"){
				gp.DateTime = value.(string)
			}else if(key == "Points"){
				gp.Points = value.(float64)
			}else if(key == "Reason"){
				gp.ItemName = value.(string)
			}else if(key == "Awardedby"){
				gp.Awardedby = value.(string)
			}else if(key == "Foritem"){
				gp.Foritem = value.(string)
			}
		}
	}
	gp.PlayerMessageList = append(gp.PlayerMessageList,PlayerMessage)
}



func ReadLuaFile(filePath,logType string,lm *LuaMessage) error{
	// 解析lua文件
	defer func(){
		err := recover()
		if err != nil{
			log.Println(err)
		}
	}()
	luaMessage := make(map[string]interface{})
	L := lua.NewState()
	defer L.Close()
	err := L.DoFile(filePath)
	if err != nil{
		log.Println(err)
		return err
	}
	err = gluamapper.Map(L.GetGlobal(logType).(*lua.LTable), &luaMessage)
	if err != nil{
		log.Println(err)
		return err
	}
	for _,value := range luaMessage{
		OneMessage := GameMessage{}
		typeValue := reflect.TypeOf(value).String()
		if typeValue == "float64"{
			continue
		}else{
			ParseMap(value.(map[interface{}]interface{}),&OneMessage)
			lm.GameMessageList = append(lm.GameMessageList,OneMessage)
		}
	}
	return nil
}





