package main

import (
	"fmt"
	"libraryMG/model"
)

var (
	curPos int = TopMenu
	curSel int
	exit   bool
	mgr    = model.NewBookMgr()
	user   *model.User
)

const (
	TopMenu = 1
)

func mainMenu() {
	curSel = 0
	fmt.Println("\n\t\t图书管理系统")
	fmt.Println("\t1. 管理员登陆")
	fmt.Println("\t2. 用户登陆")
	fmt.Println("\t3. 退出\n")
	fmt.Printf("序号:")
	fmt.Scan(&curSel)
	switch curSel {
	case 1:
		u, err := adminLogin()
		if err != nil {
			fmt.Printf("\nadmin login failed, err:%v\n", err)
			return
		}
		user = u
		showAdminMenu()
	case 2:
		u, err := userLogin()
		if err != nil {
			fmt.Printf("\nlogin failed, err:%v\n", err)
			return
		}
		user = u
		showUserMenu()
	case 3:
		exit = true
		return
	default:

	}
}

func main() {
	for !exit {
		if curPos == TopMenu {
			mainMenu()
		}
	}
}
