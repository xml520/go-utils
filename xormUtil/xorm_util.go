package xormUtil

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm/names"
)
import "xorm.io/xorm"

const (
	Eq = iota
	In
	Like
	Between
	Gt
	Lt
)

type Engine struct {
	*xorm.Engine
}
type Opt struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
}
type BuilderQuery map[string]int
type Builder map[string]string

func (f BuilderQuery) Builder(v Builder) builder.Cond {
	var where = builder.NewCond()
	for key, val := range v {
		if typ, ok := f[key]; ok {
			switch typ {
			case Eq:
				fmt.Println(key, val)
				where = where.And(builder.Eq{key: val})
			case Like:
				where = where.And(builder.Like{key, fmt.Sprintf("%%%s%%", val)})
			case In:
				where = where.And(builder.In(key, strings.Split(val, ",")))
			case Between:
				where = where.And(builder.Between{Col: key, LessVal: val, MoreVal: v[key+"_end"]})
			case Gt:
				where = where.And(builder.Gt{key: val})
			case Lt:
				where = where.And(builder.Lt{key: val})
			default:
				panic("unhandled default case")
			}
		}
	}
	return where
}
func (f Builder) BuilderWhere(q BuilderQuery) (builder.Cond, error) {
	var where = builder.NewCond()
	for key, val := range f {
		if typ, ok := q[key]; ok {
			switch typ {
			case Eq:
				where = where.And(builder.Eq{key: val})
			case Like:
				where = where.And(builder.Like{key, fmt.Sprintf("%%%s%%", val)})
			case In:
				//intArr, err := f.SplitInt(key, ",")
				//if err != nil {
				//	return nil, err
				//}
				where = where.And(builder.Eq{key: strings.Split(f[key], ",")})
			case Between:
				where = where.And(builder.Between{Col: key, LessVal: val, MoreVal: f[key+"_end"]})
			case Gt:
				where = where.And(builder.Gt{key: val})
			case Lt:
				where = where.And(builder.Lt{key: val})
			default:
				panic("unhandled default case")
			}
		}
	}
	if ok := where.IsValid(); !ok {
		return nil, fmt.Errorf("查询参数验证失败")
	}
	return where, nil
}
func (f Builder) SplitInt(key, sep string) ([]int, error) {
	var err error
	strArr := strings.Split(f[key], sep)
	intArr := make([]int, len(strArr))
	for i, s := range strArr {
		if intArr[i], err = strconv.Atoi(s); err != nil {
			return nil, err
		}
	}
	return intArr, nil
}
func (f Builder) Limit() int {
	if limit, _ := strconv.Atoi(f["limit"]); limit == 0 {
		if limit > 0 {
			return limit
		}
	}
	return 20
}
func (f Builder) Page() int {
	if page, _ := strconv.Atoi(f["page"]); page == 0 {
		if page > 0 {
			return page
		}
	}
	return 1
}
func (f Builder) Select() string {
	return f["select"]
}
func (f Builder) OrderBy() string {
	var (
		orderID string
		order   string
	)
	if orderID = f["order_id"]; orderID == "" {
		orderID = "id"
	}
	if order = f["order"]; order != "asc" {
		order = "desc"
	}
	return fmt.Sprintf("%s %s", order, orderID)

}
func (f Builder) Set(k, v string) Builder {
	f[k] = v
	return f
}
func NewMysqlEngine(dsn string, opt *Opt) (engine *Engine, err error) {
	engine = new(Engine)
	engine.Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}
	engine.SetMapper(names.GonicMapper{})

	db := engine.DB()
	_, err = db.Exec(`SET default_storage_engine='INNODB'`)
	if err != nil {
		return nil, err
	}
	if opt == nil {
		opt = &Opt{
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxIdleTime: 59 * time.Second,
		}
	}
	db.SetConnMaxIdleTime(opt.ConnMaxIdleTime)
	db.SetMaxIdleConns(opt.MaxIdleConns)
	db.SetMaxOpenConns(opt.MaxOpenConns)
	return engine, err
}
