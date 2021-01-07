package main

import (
	"fmt"
	"libraryMG/model"
	"time"
)

func showAdminMenu() {

	for {
		fmt.Println()
		fmt.Println("图书管理系统>管理员登录:")
		fmt.Println("\t1. 添加书籍")
		fmt.Println("\t2. 添加用户")
		fmt.Println("\t3. 书籍列表")
		fmt.Println("\t4. 用户列表")
		fmt.Println("\t5. 返回")
		fmt.Println()
		fmt.Printf("序号:")
		fmt.Scan(&curSel)

		switch curSel {
		case 1:
			addBook()
		case 2:
			addUser()
		case 3:
			listBook()
		case 4:
			listUser()
		case 5:
			return
		default:
		}
	}

}

func showUserMenu() {

	for {
		fmt.Println()
		fmt.Println("图书管理系统>用户登录:")
		fmt.Println("\t1. 借书")
		fmt.Println("\t2. 还书")
		fmt.Println("\t3. 已借列表")
		fmt.Println("\t4. 返回")
		fmt.Println()
		fmt.Printf("序号:")
		fmt.Scan(&curSel)

		switch curSel {
		case 1:
			borrowBook()
		case 2:
			backBook()
		case 3:
			listBorrowedBook()
		case 4:
			return
		default:
		}
	}

}

func addBook() {
	var (
		Sn      string
		Name    string
		Publish string
		Date    string
		Author  string
	)
	fmt.Println("请输入准备添加的书籍信息：")

	fmt.Printf("Sn:")
	fmt.Scan(&Sn)
	fmt.Printf("Name:")
	fmt.Scan(&Name)
	fmt.Printf("Publish:")
	fmt.Scan(&Publish)
	fmt.Printf("Auther:")
	fmt.Scan(&Author)

	newbook := model.NewBook(Sn, Name, Publish, Date, Author)
	mgr := model.NewBookMgr()
	err := mgr.AddBook(newbook)
	if err != nil {
		fmt.Println("书籍添加失败：", err)
		return
	}
	fmt.Println("书籍添加成功！")
}

func addUser() {
	var (
		username string
		passwd   string
		grade    string
		sex      string
		age      int
		id       string
	)
	fmt.Println("请输入准备添加的用户信息：")

	fmt.Printf("username:")
	fmt.Scan(&username)
	fmt.Printf("passwd:")
	fmt.Scan(&passwd)
	fmt.Printf("grade:")
	fmt.Scan(&grade)
	fmt.Printf("sex:")
	fmt.Scan(&sex)
	fmt.Printf("age:")
	fmt.Scan(&age)
	fmt.Printf("id:")
	fmt.Scan(&id)

	newuser := model.NewUser(username, passwd, grade, sex, age, id)
	mgr := model.NewBookMgr()
	err := mgr.AddUser(newuser)
	if err != nil {
		fmt.Println("用户添加失败：", err)
		return
	}
	fmt.Println("用户添加成功！")
}

func listBook() {
	var (
		offset, size int
	)
	fmt.Printf("请输入起始查询编号：")
	fmt.Scan(&offset)
	fmt.Printf("请输入查询的数量：")
	fmt.Scan(&size)

	mgr := model.NewBookMgr()
	books, err := mgr.ListAllBook(offset, size)
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	if size > len(books)-offset {
		size = len(books)
	}
	for i := offset; i < len(books); i++ {
		fmt.Printf("%d :%v\n", i, *books[i])
	}
}
func listUser() {
	var (
		offset, size int
	)
	fmt.Printf("请输入起始查询编号：")
	fmt.Scan(&offset)
	fmt.Printf("请输入查询的数量：")
	fmt.Scan(&size)

	mgr := model.NewBookMgr()
	users, err := mgr.ListUser(offset, size)
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	if size > len(users)-offset {
		size = len(users)
	}
	for i := offset; i < len(users); i++ {
		fmt.Printf("%d :%v\n", i, *users[i])
	}
}
func borrowBook() {
	var (
		uname    string
		sn       string
		interval time.Duration
	)
	fmt.Printf("请输入用户名称：")
	fmt.Scan(&uname)
	fmt.Printf("请输入书籍SN：")
	fmt.Scan(&sn)
	fmt.Printf("请输入借用时间间隔：")
	fmt.Scan(&interval)

	mgr := model.NewBookMgr()
	book, err := mgr.Borrow(uname, sn, interval)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v 已成功借出！\n", book.Name)
}

func backBook() {
	var (
		uname string
		sn    string
	)
	fmt.Printf("请输入用户名称：")
	fmt.Scan(&uname)
	fmt.Printf("请输入书籍SN：")
	fmt.Scan(&sn)

	mgr := model.NewBookMgr()
	err := mgr.BackBook(uname, sn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("已成功还书！\n")
}

func listBorrowedBook() {
	user, err := userLogin()
	if err != nil {
		fmt.Println(err)
		return
	}

	bookRecord := user.GetBookRecords()
	if len(bookRecord) == 0 {
		fmt.Println("暂无借书！")
		return
	}
	for bookSN, _ := range bookRecord {
		fmt.Printf("SN: %v\n", bookSN)
	}
}
