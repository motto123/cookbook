package menu

import (
	"net/http"
	"cookbook/db"
	"cookbook/public/logger"
	"cookbook/public/code"
	"strconv"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)

type menuInfo struct {
	Name   string `json:"name"`
	Spicy  int    `json:"spicy"`
	Weight int    `json:"weight"`
	IsMeat bool   `json:"is_meat"`
	Pic    string `json:"pic"`
}

func HandleGetAllMenu(c *gin.Context) {
	rows, err := db.DbApp.Query(`SELECT name, spicy FROM menu`)
	if err != nil {
		logger.Sugar.Error(err)
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	menuList := make([]menuInfo, 0)
	for rows.Next() {
		var m menuInfo
		rows.Scan(&m.Name, &m.Spicy)
		menuList = append(menuList, m)
	}
	ret := code.RetNormal
	ret.RetData = menuList
	c.JSON(http.StatusOK, ret)
}

func HandleLookMenu(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)
	if userId == 0 {
		c.JSON(http.StatusOK, code.RetInvalidParams)
		return
	}
	menuList, err := getMenuListByUserId(userId)
	if err != nil {
		logger.Sugar.Error(err)
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	ret := code.RetNormal
	ret.RetData = menuList
	c.JSON(http.StatusOK, ret)
}

func getMenuListByUserId(userId int) (menuList []menuInfo, err error) {
	rows, err := db.DbApp.Query(`SELECT m.name, m.spicy, um.weight, m.is_meat, m.spicy
		FROM menu m JOIN user_menu um ON um.menu_id = m.id WHERE um.user_id = ?;`, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	menuList = make([]menuInfo, 0)
	for rows.Next() {
		var isMeat int
		var m menuInfo
		rows.Scan(&m.Name, &m.Spicy, &m.Weight, &isMeat, &m.Pic)
		if isMeat == 1 {
			m.IsMeat = true
		}
		menuList = append(menuList, m)
	}
	return
}

type insertMenu struct {
	menuInfo
	UserId int `json:"user_id"`
}
type reqInsertMenu struct {
	List []insertMenu `json:"list"`
}

func HandleInsertMenu(c *gin.Context) {
	var r reqInsertMenu
	err := c.BindJSON(&r)
	fmt.Printf("%+v\n", r.List)
	if err != nil {
		c.JSON(http.StatusOK, code.GeneralErrRet(err))
		return
	}
	if len(r.List) == 0 {
		c.JSON(http.StatusOK, code.RetInvalidParams)
		return
	}
	tx, err := db.DbApp.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	defer tx.Rollback()
	for _, menu := range r.List {
		if menu.UserId == 0 {
			c.JSON(http.StatusOK, code.RetInvalidParams)
			return
		}
		var menuId int64
		row := db.DbApp.QueryRow(`SELECT id FROM menu WHERE name = ?;`, menu.Name)
		row.Scan(&menuId)
		if menuId == 0 {
			result, err := tx.Exec(`INSERT menu(name, spicy, is_meat) VALUE(?, ?, ?);`, menu.Name, menu.Spicy, menu.IsMeat)
			if err != nil {
				c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
				return
			}
			menuId, _ = result.LastInsertId()
			if err != nil {
				c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
				return
			}
		}
		_, err = tx.Exec(`INSERT user_menu(user_id, menu_id, weight) VALUE(?, ?, ?);`,
			menu.UserId, menuId, 1)
		if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, code.RetNormal)
}

type reqDeleteMenu struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

func HandleDeleteMenu(c *gin.Context) {
	var r reqDeleteMenu
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	if r.UserId == 0 || r.Name == "" {
		c.JSON(http.StatusOK, code.RetInvalidParams)
		return
	}
	id := 0
	row := db.DbApp.QueryRow(`SELECT id FROM menu WHERE name = ?;`, r.Name)
	row.Scan(&id)
	_, err = db.DbApp.Exec(`DELETE FROM user_menu WHERE user_id = ? AND  id = ?;`, r.UserId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	c.JSON(http.StatusOK, code.RetNormal)
}

type reqAlterMenu struct {
	OldMenuName string `json:"old_menu_name"`
	NewMenuName string `json:"new_menu_name"`
	NewSpicy    int    `json:"new_spicy"`
}

func HandleAlterMenu(c *gin.Context) {
	var r reqAlterMenu
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	if r.OldMenuName == "" || (r.NewMenuName == "" && r.NewSpicy == 0) {
		c.JSON(http.StatusOK, code.RetInvalidParams)
		return
	}
	var mi menuInfo
	row := db.DbApp.QueryRow(`SELECT name, spicy FROM menu WHERE name = ?;`, r.OldMenuName)
	row.Scan(&mi.Name, &mi.Spicy)
	if r.NewSpicy == 0 {
		r.NewSpicy = mi.Spicy
	}
	if r.NewMenuName == "" {
		r.NewMenuName = mi.Name
	}
	_, err = db.DbApp.Exec(`UPDATE menu SET name = ?, spicy = ? WHERE name = ?;`, r.NewMenuName, r.NewSpicy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.GeneralErrRet(err))
		return
	}
	c.JSON(http.StatusOK, code.RetNormal)
}
