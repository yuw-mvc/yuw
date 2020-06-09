package models

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
	C "github.com/yuw-mvc/yuw/configs"
	E "github.com/yuw-mvc/yuw/exceptions"
)

const (
	ONE   string = "ONE"
	ALL   string = "ALL"
	ByAsc string = "ASC"
	ByEsc string = "DESC"
)

type Models struct {
	Engine *xorm.Engine
}

func New(engine *xorm.Engine) *Models {
	return &Models {
		Engine:engine,
	}
}

func (m *Models) Insert(data interface{}) (i int64, err error) {
	xm := m.Engine.NewSession()
	defer xm.Close()

	i, err = xm.Insert(data)
	return
}

func (m *Models) Update(mPoT *C.MPoT, data interface{}) (i int64, err error) {
	if mPoT.Query == nil || mPoT.QueryArgs == nil {
		err = E.Err("m^aa", E.ErrPosition())
		return
	}

	xm := m.Engine.NewSession()
	defer xm.Close()

	i, err = xm.Where(mPoT.Query, mPoT.QueryArgs ...).Update(data)
	return
}

func (m *Models) Delete(mPoT *C.MPoT, data interface{}) (i int64, err error) {
	if mPoT.Query == nil || mPoT.QueryArgs == nil {
		err = E.Err("m^aa", E.ErrPosition())
		return
	}

	xm := m.Engine.NewSession()
	defer xm.Close()

	i, err = xm.Where(mPoT.Query, mPoT.QueryArgs ...).Delete(data)
	return
}

func (m *Models) Total(data interface{}) (nums int64, err error) {
	xm := m.Engine.NewSession()
	defer xm.Close()

	nums, err = xm.Count(data)
	return
}

func (m *Models) Select(mPoT *C.MPoT, data interface{}) (err error) {
	xm := m.Engine.NewSession()

	/**
	 * Todo: Throw Error
	 */
	defer func() {
		xm.Close()
	}()

	if mPoT.Table != "" && mPoT.Field != "" {
		if ok, _ := xm.IsTableExist(mPoT.Table); ok == false {
			err = E.Err("", E.ErrPosition())
			return
		}

		xm = xm.Table(mPoT.Table).Select(mPoT.Field)
	} else {
		if mPoT.Columns != nil {
			xm = xm.Cols(mPoT.Columns ...)
		} else {
			xm = xm.Cols()
		}
	}

	/**
	 * Todo: Add Join Table (INNER|LEFT|RIGHT)
	**/
	if mPoT.Joins != nil {
		for _, join := range mPoT.Joins {
			if len(join) > 2 {
				xm = xm.Join(
					cast.ToString(join[0]),
					join[1],
					cast.ToString(join[2]),
				)
			}
		}
	}

	/**
	 * Todo: Add Condition
	 */
	if mPoT.Query != nil && mPoT.QueryArgs != nil {
		xm = xm.Where(mPoT.Query, mPoT.QueryArgs ...)
	}

	switch mPoT.Types {
	case ONE:
		_, err = xm.Get(data)
		return

	case ALL:
		/**
		 * Todo: Add OrderBy (ASC|DESC)
		**/
		if len(mPoT.OrderArgs) > 0 {
			if mPoT.OrderType == ByAsc {
				xm = xm.Asc(mPoT.OrderArgs ...)
			}

			if mPoT.OrderType == ByEsc {
				xm = xm.Desc(mPoT.OrderArgs ...)
			}
		}

		/**
		 * Todo: Limit & Start
		**/
		if len(mPoT.Start) > 0 && mPoT.Limit != 0 {
			xm = xm.Limit(mPoT.Limit, mPoT.Start ...)
		}

		err = xm.Find(data)
		return

	default:
		err = E.Err("yuw^m_db_e", E.ErrPosition())
		return
	}

	return
}

func (m *Models) ToSerialized() {

}
