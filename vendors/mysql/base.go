package mysql

import "github.com/astaxie/beego/orm"

//选择数据库
func GetOrmer(genre ...string) orm.Ormer {
	var g string
	if len(genre) > 0 && genre[0] != "" {
		g = genre[0]
	} else {
		g = "default"
	}
	if ormers[g] != nil {
		return ormers[g]
	}
	o := orm.NewOrm()
	_ = o.Using(g)
	ormers[g] = o
	return o
}

//选择数据库中的指定表
func GetQuerySetter(ptrStructOrTableName interface{}, genre ...string) orm.QuerySeter {
	var o orm.Ormer
	if len(genre) > 0 && genre[0] != "" {
		o = GetOrmer(genre[0])
	} else {
		o = GetOrmer()
	}
	return o.QueryTable(ptrStructOrTableName)
}
