package gen

const templateDaoDaoIndexContent = `
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"{TplImportPrefix}/dao/internal"
)

// {TplTableNameCamelLowerCase}Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal.{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.{TplTableNameCamelCase},
	}
)

// Fill with you ideas below.

`

const templateDaoDaoInternalContent = `
// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"time"
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"

	"{TplImportPrefix}/model"
)

// {TplTableNameCamelCase}Dao is the manager for logic model data accessing
// and custom defined data operations functions management.
type {TplTableNameCamelCase}Dao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns {TplTableNameCamelLowerCase}Columns
	ctx 	context.Context
}

// {TplTableNameCamelCase}Columns defines and stores column names for table {TplTableName}.
type {TplTableNameCamelLowerCase}Columns struct {
	{TplColumnDefine}
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelCase}Dao{
		M:     g.DB("{TplGroupName}").Model("{TplTableName}").Safe(),
		DB:    g.DB("{TplGroupName}"),
		Table: "{TplTableName}",
		Columns: {TplTableNameCamelLowerCase}Columns{
			{TplColumnNames}
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *{TplTableNameCamelCase}Dao) Ctx(ctx context.Context) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Ctx(ctx), ctx: ctx}
}

// As sets an alias name for current table.
func (d *{TplTableNameCamelCase}Dao) As(as string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.As(as), ctx: d.ctx}
}

// TX sets the transaction for current operation.
func (d *{TplTableNameCamelCase}Dao) TX(tx *gdb.TX) *{TplTableNameCamelCase}Dao {
	if tx != nil {
		return &{TplTableNameCamelCase}Dao{M: d.M.TX(tx), ctx: d.ctx}
	}
	return &{TplTableNameCamelCase}Dao{M: d.M, ctx: d.ctx}
}

// Master marks the following operation on master node.
func (d *{TplTableNameCamelCase}Dao) Master() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Master(), ctx: d.ctx}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *{TplTableNameCamelCase}Dao) Slave() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Slave(), ctx: d.ctx}
}

// Args sets custom arguments for model operation.
func (d *{TplTableNameCamelCase}Dao) Args(args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Args(args...), ctx: d.ctx}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) LeftJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LeftJoin(table...), ctx: d.ctx}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) RightJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.RightJoin(table...), ctx: d.ctx}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) InnerJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.InnerJoin(table...), ctx: d.ctx}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *{TplTableNameCamelCase}Dao) Fields(fieldNamesOrMapStruct ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Fields(fieldNamesOrMapStruct...), ctx: d.ctx}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *{TplTableNameCamelCase}Dao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.FieldsEx(fieldNamesOrMapStruct...), ctx: d.ctx}
}

// Option sets the extra operation option for the model.
func (d *{TplTableNameCamelCase}Dao) Option(option int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Option(option), ctx: d.ctx}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *{TplTableNameCamelCase}Dao) OmitEmpty() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.OmitEmpty(), ctx: d.ctx}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *{TplTableNameCamelCase}Dao) Filter() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Filter(), ctx: d.ctx}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *{TplTableNameCamelCase}Dao) Where(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Where(where, args...), ctx: d.ctx}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *{TplTableNameCamelCase}Dao) WherePri(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.WherePri(where, args...), ctx: d.ctx}
}

// And adds "AND" condition to the where statement.
func (d *{TplTableNameCamelCase}Dao) And(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.And(where, args...), ctx: d.ctx}
}

// Or adds "OR" condition to the where statement.
func (d *{TplTableNameCamelCase}Dao) Or(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Or(where, args...), ctx: d.ctx}
}

// Group sets the "GROUP BY" statement for the model.
func (d *{TplTableNameCamelCase}Dao) Group(groupBy string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Group(groupBy), ctx: d.ctx}
}

// Order sets the "ORDER BY" statement for the model.
func (d *{TplTableNameCamelCase}Dao) Order(orderBy ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Order(orderBy...), ctx: d.ctx}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *{TplTableNameCamelCase}Dao) Limit(limit ...int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Limit(limit...), ctx: d.ctx}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *{TplTableNameCamelCase}Dao) Offset(offset int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Offset(offset), ctx: d.ctx}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *{TplTableNameCamelCase}Dao) Page(page, limit int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Page(page, limit), ctx: d.ctx}
}

// Batch sets the batch operation number for the model.
func (d *{TplTableNameCamelCase}Dao) Batch(batch int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Batch(batch), ctx: d.ctx}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *{TplTableNameCamelCase}Dao) Cache(duration time.Duration, name ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Cache(duration, name...), ctx: d.ctx}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *{TplTableNameCamelCase}Dao) Data(data ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Data(data...), ctx: d.ctx}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.{TplTableNameCamelCase}.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *{TplTableNameCamelCase}Dao) All(where ...interface{}) ([]*model.{TplTableNameCamelCase}, error) {
	var all gdb.Result
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	dbName := gconv.String(d.ctx.Value("dbname"))
	if dbName == "" {
		all, err = d.M.All(where...)
	} else {
		all, err = d.M.Schema(dbName).All(where...)
	}
	if err != nil {
		return nil, err
	}
	var entities []*model.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.{TplTableNameCamelCase}.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *{TplTableNameCamelCase}Dao) One(where ...interface{}) (*model.{TplTableNameCamelCase}, error) {
	var one gdb.Record
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	dbName := gconv.String(d.ctx.Value("dbname"))
	if dbName == "" {
		one, err = d.M.One(where...)
	} else {
		one, err = d.M.Schema(dbName).One(where...)
	}
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *{TplTableNameCamelCase}Dao) FindOne(where ...interface{}) (*model.{TplTableNameCamelCase}, error) {
	var one gdb.Record
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	dbName := gconv.String(d.ctx.Value("dbname"))
	if dbName == "" {
		one, err = d.M.FindOne(where...)
	} else {
		one, err = d.M.Schema(dbName).FindOne(where...)
	}
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *{TplTableNameCamelCase}Dao) FindAll(where ...interface{}) ([]*model.{TplTableNameCamelCase}, error) {
	var all gdb.Result
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	dbName := gconv.String(d.ctx.Value("dbname"))
	if dbName == "" {
		all, err = d.M.FindAll(where...)
	} else {
		all, err = d.M.Schema(dbName).FindAll(where...)
	}
	if err != nil {
		return nil, err
	}
	var entities []*model.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *{TplTableNameCamelCase}Dao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *{TplTableNameCamelCase}Dao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *{TplTableNameCamelCase}Dao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *{TplTableNameCamelCase}Dao) Chunk(limit int, callback func(entities []*model.{TplTableNameCamelCase}, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.{TplTableNameCamelCase}
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *{TplTableNameCamelCase}Dao) LockUpdate() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LockUpdate(), ctx: d.ctx}
}

// LockShared sets the lock in share mode for current operation.
func (d *{TplTableNameCamelCase}Dao) LockShared() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LockShared(), ctx: d.ctx}
}

// Unscoped enables/disables the soft deleting feature.
func (d *{TplTableNameCamelCase}Dao) Unscoped() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Unscoped(), ctx: d.ctx}
}

// Platform check dbname to platform
func (d *{TplTableNameCamelCase}Dao) Platform() *{TplTableNameCamelCase}Dao {
	ctx := context.WithValue(d.ctx, "dbname", "platform")
	return &{TplTableNameCamelCase}Dao{M: d.M, ctx: ctx}
}

// FindOneByIDWithCache retrieves and returns a single Record with cache by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *{TplTableNameCamelCase}Dao) FindOneByIDWithCache(id int64) (*model.{TplTableNameCamelCase}, error) {
	var one gdb.Record
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	result, err := g.Redis().Do("GET", d.GetRowKey(id))
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if gconv.String(result) == "" {
		dbName := gconv.String(d.ctx.Value("dbname"))
		if dbName == "" {
			one, err = d.M.Cache(time.Hour*24, d.GetRowKey(id)).FindOne(id)
		} else {
			one, err = d.M.Schema(dbName).Cache(time.Hour*24, d.GetRowKey(id)).FindOne(id)
		}
		if err != nil {
			return nil, err
		}
		if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		encodeByte, err := gjson.Encode(entity)
		if err != nil {
			return nil, nil
		}
		_, err = g.Redis().Do("SET", d.GetRowKey(id), gconv.String(encodeByte))
		if err != nil {
			return nil, err
		}
		_, err = g.Redis().Do("EXPIRE", d.GetRowKey(id), 3600*24)
		if err != nil {
			return nil, err
		}
	} else {
		err = gconv.Struct(gconv.String(result), &entity)
		if err != nil {
			return nil, err
		}
	}
	return entity, nil
}

func (d *{TplTableNameCamelCase}Dao) FindOneByID(id int64) (*model.{TplTableNameCamelCase}, error) {
	var one gdb.Record
	var err error
	if d.ctx == nil {
		return nil, errors.New("必须传ctx")
	}
	dbName := gconv.String(d.ctx.Value("dbname"))
	if dbName == "" {
		one, err = d.M.FindOne(id)
	} else {
		one, err = d.M.Schema(dbName).FindOne(id)
	}
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

func (d *{TplTableNameCamelCase}Dao) GetRowKey(id int64) string {
	return fmt.Sprintf("system:table:%s:row:pri:%d", {TplTableNameCamelCase}.Table, id)
}
`
