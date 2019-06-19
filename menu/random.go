package menu

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"cookbook/public/code"
	"math/rand"
	"cookbook/public/datastruct"
)

func HandleRandomMenu(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)
	spicyStr := c.Query("spicy")
	spicy, _ := strconv.Atoi(spicyStr)
	cntStr := c.Query("cnt")
	cnt, _ := strconv.Atoi(cntStr)
	if userId == 0 || spicy > 4 || cnt == 0 {
		c.JSON(http.StatusOK, code.RetInvalidParams)
		return
	}
	menuList, err := getMenuListByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	if cnt > len(menuList) {
		errResponse := code.RetOperationErr
		errResponse.Err = "你的菜品数量不足"
		c.JSON(http.StatusOK, errResponse)
		return
	}
	var firstMenus []menuInfo
	var secondsMenus []menuInfo
	for _, v := range menuList {
		if v.Spicy == spicy {
			firstMenus = append(firstMenus, v)
		}
		//spicy必须是清淡以上的口味,不然没有第二类菜
		if spicy-1 >= 0 && v.Spicy == spicy-1 {
			secondsMenus = append(secondsMenus, v)
		}
	}

	choseM := datastruct.NewSetWithParseKeyFunc(func(e interface{}) (key interface{}) {
		return e.(menuInfo).Name
	})
	if cnt == 1 && len(firstMenus) >= 1 {
		choseM.Add(firstMenus[rand.Intn(len(firstMenus))])
	} else if spicy == 0 { //全部选清淡口味
		num := cnt
		if len(firstMenus) < num {
			num = len(firstMenus)
		}
		for i := 0; i < num; i++ {
			choseM.Add(firstMenus[rand.Intn(len(firstMenus))])
		}
	} else {
		//保证总数量除以第一品类菜接近50%
		num := cnt / 2
		if len(firstMenus) < num {
			num = len(firstMenus)
		}
		for i := 0; i < num; i++ {
			choseM.Add(firstMenus[rand.Intn(len(firstMenus))])
		}
		if len(secondsMenus) < cnt-num {
			num = len(secondsMenus)
		}
		for i := 0; i < cnt-num; i++ {
			choseM.Add(secondsMenus[rand.Intn(len(secondsMenus))])
		}
		for choseM.Len() < cnt { //前面的菜不够,从个人菜谱中随机取来填充
			for i := 0; i < cnt-choseM.Len(); i++ {
				choseM.Add(menuList[rand.Intn(len(menuList))])
			}
		}
	}
	ret := code.RetNormal
	ret.RetData = gin.H{
		"list": choseM.ToArr(),
	}
	c.JSON(http.StatusOK, ret)
}
