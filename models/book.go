package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id   int64
	Name string
	Year int64	`form:"year"`
}
type BookPage struct {
	Sum int 	`json:"sum"`
	Offset int 	`json:"offset"`
	Info []map[string]interface{}	`json:"info"`

}
type BookRows struct {

}

type  PageFactory struct {
	Offset int
	Num int
}

func init() {
	orm.RegisterModel(new(Book))
}

// AddBook insert a new Book into database and returns
// last inserted Id on success.
func AddBook(m *Book) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBookById retrieves Book by Id. Returns error if
// Id doesn't exist
func GetBookById(id int64) (v *Book, err error) {
	o := orm.NewOrm()
	v = &Book{Id: id}
	if err = o.QueryTable(new(Book)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//type GetPageFactory()


func GetBooksNums()(nums int)  {
	o := orm.NewOrm()
	o.Raw("select count(*) as n from book;").QueryRow(&nums)
	//fmt.Println("res",res)
	return  nums
}

func GetPageFactory()PageFactory  {
	return PageFactory{
		Offset: 0,
		Num: 5,
	}
}

func GetPageFunc(offset int,n int) (rows []map[string]interface{}, err error) {
	var maps []orm.Params
	o := orm.NewOrm()
	f:=fmt.Sprintf("select * from book limit %v,%v;",offset,n)
	num, err := o.Raw(f).Values(&maps)
	if err != nil && num == 0 {
		return rows,err
	}
	fmt.Print("%T",maps) // slene
	var s []map[string]interface{}
	for _,v := range maps{
		//fmt.Println("map's val:", v)
		s=append(s, v)
	}
	//fmt.Printf("a %v",s)
	rows = s
	return rows,err
}

func UpdateFunc(data Book)(id int64){
	//o := orm.NewOrm()
	//f:=fmt.Sprintf("update book set name='jsj12' where 1=1 and id =%v",data.Id)
	//res,err :=o.Raw(f).Exec()
	//num, _ := res.RowsAffected()
	//fmt.Println(res,err)
	//fmt.Println(num)
	fmt.Println(data)
	o := orm.NewOrm()
	user := Book{Id: data.Id}
	if o.Read(&user) == nil {
		//user.Name = "MyName"
		if num, err := o.Update(&data); err == nil {
			fmt.Println(num)
		}
	}
	id =data.Id
	return  id
}



