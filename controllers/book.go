package controllers

import (
	"api/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

//  BookController operations for Book
type BookController struct {
	beego.Controller
}


type BookPage struct {
	Sum int 	`json:"sum"`
	Offset int 	`json:"offset"`
	Info []map[string]interface{}	`json:"info"`

}



// URLMapping ...
func (c *BookController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create Book
// @Param	body		body 	models.Book	true		"body for Book content"
// @Success 201 {int} models.Book
// @Failure 403 body is empty
// @router / [post]
func (c *BookController) Post() {
	//var v models.Book
	//json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	//
	//r :=c.Ctx.Request
	//fmt.Println(c.Ctx.Request.Header.Get("Content-Type"))
	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//a :=string(body)
	//fmt.Printf("%T,%v",a,a)
	////fmt.Println(a)
	//json.Unmarshal(body, &v)
	//fmt.Println(c.GetString("action"))
	//fmt.Print()
	//c.ParseForm(&v)
	var action string
	action = c.GetString("action")
	switch action {
	case "add":
		fmt.Println("add")
		add(c)
		break
	case "show":
		fmt.Println("show")
		show(c)
		break
	case "update":
		fmt.Println("update")
		update(c)
		break
	}
	return
}

func add(c *BookController)  {
	var v models.Book
	v.Name = c.GetString("name")
	v.Year, _ = c.GetInt64("year")

	if _, err := models.AddBook(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
/**
 * @Param ac offset num
 *
 */
func show(c *BookController)  {
	var data BookPage
	cfg:= models.GetPageFactory()

	offset,err:=c.GetInt("offset")
	if err==nil{
		cfg.Offset= offset
	}
	num,err:=c.GetInt("num")
	if err ==nil {
		cfg.Num= num
	}
	//if num >100 {
	//	num = 100
	//}
	if num<1 {
		c.Ctx.WriteString("参数错误，num需要为正数")
		return
	}
	index := cfg.Offset * cfg.Num
	//查询 总行数
	nums := models.GetBooksNums()
	//查询 分页结果集
	l,err:= models.GetPageFunc(index,cfg.Num)
	if err !=nil {
		fmt.Errorf("models.Test :%v",err)
	}
	data.Sum = nums
	data.Info = l
	c.Data["json"] = data
	c.ServeJSON()
}

func checkParam(){
	
}

type Rule struct {

} 

func update(c *BookController){
	//offset,_:=c.GetInt("offset")
	//num,_:=c.GetInt("num")
	//val := make(models.Book)
	var v models.Book
	v.Id,_ = c.GetInt64("id")
	//if v.Id != 0 {
	//	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	//	//err :=c.ParseForm(&v)
	//	//fmt.Println(err)
	//}
	v.Name = c.GetString("name")
	v.Year, _ = c.GetInt64("year")
	fmt.Println(v)
	//var rule interface{}
	//rule := {
	//	value: 1,
	//	key: "id",
	//	//reg: 
	//}
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//id :=models.UpdateFunc(v)
	//fmt.Println(v,id)
	c.ServeJSON()

}

// GetOne ...
// @Title Get One
// @Description get Book by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Book
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BookController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetBookById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}


// GetAll ...
// @Title Get All
// @Description get Book
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Book
// @Failure 403
// @router / [get]

func (c *BookController) GetAll() {


	c.ServeJSON()
}
