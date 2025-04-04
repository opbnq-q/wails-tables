// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWoodSpec(db *gorm.DB, opts ...gen.DOOption) woodSpec {
	_woodSpec := woodSpec{}

	_woodSpec.woodSpecDo.UseDB(db, opts...)
	_woodSpec.woodSpecDo.UseModel(&models.WoodSpec{})

	tableName := _woodSpec.woodSpecDo.TableName()
	_woodSpec.ALL = field.NewAsterisk(tableName)
	_woodSpec.Id = field.NewUint(tableName, "id")
	_woodSpec.WoodSpecTypeId = field.NewUint(tableName, "wood_spec_type_id")
	_woodSpec.Poroda = field.NewString(tableName, "poroda")
	_woodSpec.Sort = field.NewString(tableName, "sort")
	_woodSpec.DlinaBrevna = field.NewString(tableName, "dlina_brevna")
	_woodSpec.DiameterBrevna = field.NewString(tableName, "diameter_brevna")
	_woodSpec.SechenieDoski = field.NewString(tableName, "sechenie_doski")
	_woodSpec.ProcentVlajnosti = field.NewInt(tableName, "procent_vlajnosti")
	_woodSpec.DiameterGranyl = field.NewString(tableName, "diameter_granyl")
	_woodSpec.WoodSpecType = woodSpecBelongsToWoodSpecType{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("WoodSpecType", "models.WoodSpecType"),
		Receivers: struct {
			field.RelationField
			Exporter struct {
				field.RelationField
				Receivers struct {
					field.RelationField
				}
			}
			Spec struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("WoodSpecType.Receivers", "models.Receiver"),
			Exporter: struct {
				field.RelationField
				Receivers struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("WoodSpecType.Receivers.Exporter", "models.Exporter"),
				Receivers: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("WoodSpecType.Receivers.Exporter.Receivers", "models.Receiver"),
				},
			},
			Spec: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("WoodSpecType.Receivers.Spec", "models.WoodSpecType"),
			},
		},
	}

	_woodSpec.fillFieldMap()

	return _woodSpec
}

type woodSpec struct {
	woodSpecDo

	ALL              field.Asterisk
	Id               field.Uint
	WoodSpecTypeId   field.Uint
	Poroda           field.String
	Sort             field.String
	DlinaBrevna      field.String
	DiameterBrevna   field.String
	SechenieDoski    field.String
	ProcentVlajnosti field.Int
	DiameterGranyl   field.String
	WoodSpecType     woodSpecBelongsToWoodSpecType

	fieldMap map[string]field.Expr
}

func (w woodSpec) Table(newTableName string) *woodSpec {
	w.woodSpecDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w woodSpec) As(alias string) *woodSpec {
	w.woodSpecDo.DO = *(w.woodSpecDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *woodSpec) updateTableName(table string) *woodSpec {
	w.ALL = field.NewAsterisk(table)
	w.Id = field.NewUint(table, "id")
	w.WoodSpecTypeId = field.NewUint(table, "wood_spec_type_id")
	w.Poroda = field.NewString(table, "poroda")
	w.Sort = field.NewString(table, "sort")
	w.DlinaBrevna = field.NewString(table, "dlina_brevna")
	w.DiameterBrevna = field.NewString(table, "diameter_brevna")
	w.SechenieDoski = field.NewString(table, "sechenie_doski")
	w.ProcentVlajnosti = field.NewInt(table, "procent_vlajnosti")
	w.DiameterGranyl = field.NewString(table, "diameter_granyl")

	w.fillFieldMap()

	return w
}

func (w *woodSpec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *woodSpec) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.Id
	w.fieldMap["wood_spec_type_id"] = w.WoodSpecTypeId
	w.fieldMap["poroda"] = w.Poroda
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["dlina_brevna"] = w.DlinaBrevna
	w.fieldMap["diameter_brevna"] = w.DiameterBrevna
	w.fieldMap["sechenie_doski"] = w.SechenieDoski
	w.fieldMap["procent_vlajnosti"] = w.ProcentVlajnosti
	w.fieldMap["diameter_granyl"] = w.DiameterGranyl

}

func (w woodSpec) clone(db *gorm.DB) woodSpec {
	w.woodSpecDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w woodSpec) replaceDB(db *gorm.DB) woodSpec {
	w.woodSpecDo.ReplaceDB(db)
	return w
}

type woodSpecBelongsToWoodSpecType struct {
	db *gorm.DB

	field.RelationField

	Receivers struct {
		field.RelationField
		Exporter struct {
			field.RelationField
			Receivers struct {
				field.RelationField
			}
		}
		Spec struct {
			field.RelationField
		}
	}
}

func (a woodSpecBelongsToWoodSpecType) Where(conds ...field.Expr) *woodSpecBelongsToWoodSpecType {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a woodSpecBelongsToWoodSpecType) WithContext(ctx context.Context) *woodSpecBelongsToWoodSpecType {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a woodSpecBelongsToWoodSpecType) Session(session *gorm.Session) *woodSpecBelongsToWoodSpecType {
	a.db = a.db.Session(session)
	return &a
}

func (a woodSpecBelongsToWoodSpecType) Model(m *models.WoodSpec) *woodSpecBelongsToWoodSpecTypeTx {
	return &woodSpecBelongsToWoodSpecTypeTx{a.db.Model(m).Association(a.Name())}
}

type woodSpecBelongsToWoodSpecTypeTx struct{ tx *gorm.Association }

func (a woodSpecBelongsToWoodSpecTypeTx) Find() (result *models.WoodSpecType, err error) {
	return result, a.tx.Find(&result)
}

func (a woodSpecBelongsToWoodSpecTypeTx) Append(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a woodSpecBelongsToWoodSpecTypeTx) Replace(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a woodSpecBelongsToWoodSpecTypeTx) Delete(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a woodSpecBelongsToWoodSpecTypeTx) Clear() error {
	return a.tx.Clear()
}

func (a woodSpecBelongsToWoodSpecTypeTx) Count() int64 {
	return a.tx.Count()
}

type woodSpecDo struct{ gen.DO }

type IWoodSpecDo interface {
	gen.SubQuery
	Debug() IWoodSpecDo
	WithContext(ctx context.Context) IWoodSpecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWoodSpecDo
	WriteDB() IWoodSpecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWoodSpecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWoodSpecDo
	Not(conds ...gen.Condition) IWoodSpecDo
	Or(conds ...gen.Condition) IWoodSpecDo
	Select(conds ...field.Expr) IWoodSpecDo
	Where(conds ...gen.Condition) IWoodSpecDo
	Order(conds ...field.Expr) IWoodSpecDo
	Distinct(cols ...field.Expr) IWoodSpecDo
	Omit(cols ...field.Expr) IWoodSpecDo
	Join(table schema.Tabler, on ...field.Expr) IWoodSpecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWoodSpecDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWoodSpecDo
	Group(cols ...field.Expr) IWoodSpecDo
	Having(conds ...gen.Condition) IWoodSpecDo
	Limit(limit int) IWoodSpecDo
	Offset(offset int) IWoodSpecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWoodSpecDo
	Unscoped() IWoodSpecDo
	Create(values ...*models.WoodSpec) error
	CreateInBatches(values []*models.WoodSpec, batchSize int) error
	Save(values ...*models.WoodSpec) error
	First() (*models.WoodSpec, error)
	Take() (*models.WoodSpec, error)
	Last() (*models.WoodSpec, error)
	Find() ([]*models.WoodSpec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WoodSpec, err error)
	FindInBatches(result *[]*models.WoodSpec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.WoodSpec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWoodSpecDo
	Assign(attrs ...field.AssignExpr) IWoodSpecDo
	Joins(fields ...field.RelationField) IWoodSpecDo
	Preload(fields ...field.RelationField) IWoodSpecDo
	FirstOrInit() (*models.WoodSpec, error)
	FirstOrCreate() (*models.WoodSpec, error)
	FindByPage(offset int, limit int) (result []*models.WoodSpec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWoodSpecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w woodSpecDo) Debug() IWoodSpecDo {
	return w.withDO(w.DO.Debug())
}

func (w woodSpecDo) WithContext(ctx context.Context) IWoodSpecDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w woodSpecDo) ReadDB() IWoodSpecDo {
	return w.Clauses(dbresolver.Read)
}

func (w woodSpecDo) WriteDB() IWoodSpecDo {
	return w.Clauses(dbresolver.Write)
}

func (w woodSpecDo) Session(config *gorm.Session) IWoodSpecDo {
	return w.withDO(w.DO.Session(config))
}

func (w woodSpecDo) Clauses(conds ...clause.Expression) IWoodSpecDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w woodSpecDo) Returning(value interface{}, columns ...string) IWoodSpecDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w woodSpecDo) Not(conds ...gen.Condition) IWoodSpecDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w woodSpecDo) Or(conds ...gen.Condition) IWoodSpecDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w woodSpecDo) Select(conds ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w woodSpecDo) Where(conds ...gen.Condition) IWoodSpecDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w woodSpecDo) Order(conds ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w woodSpecDo) Distinct(cols ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w woodSpecDo) Omit(cols ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w woodSpecDo) Join(table schema.Tabler, on ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w woodSpecDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w woodSpecDo) RightJoin(table schema.Tabler, on ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w woodSpecDo) Group(cols ...field.Expr) IWoodSpecDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w woodSpecDo) Having(conds ...gen.Condition) IWoodSpecDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w woodSpecDo) Limit(limit int) IWoodSpecDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w woodSpecDo) Offset(offset int) IWoodSpecDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w woodSpecDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWoodSpecDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w woodSpecDo) Unscoped() IWoodSpecDo {
	return w.withDO(w.DO.Unscoped())
}

func (w woodSpecDo) Create(values ...*models.WoodSpec) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w woodSpecDo) CreateInBatches(values []*models.WoodSpec, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w woodSpecDo) Save(values ...*models.WoodSpec) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w woodSpecDo) First() (*models.WoodSpec, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.WoodSpec), nil
	}
}

func (w woodSpecDo) Take() (*models.WoodSpec, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.WoodSpec), nil
	}
}

func (w woodSpecDo) Last() (*models.WoodSpec, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.WoodSpec), nil
	}
}

func (w woodSpecDo) Find() ([]*models.WoodSpec, error) {
	result, err := w.DO.Find()
	return result.([]*models.WoodSpec), err
}

func (w woodSpecDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WoodSpec, err error) {
	buf := make([]*models.WoodSpec, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w woodSpecDo) FindInBatches(result *[]*models.WoodSpec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w woodSpecDo) Attrs(attrs ...field.AssignExpr) IWoodSpecDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w woodSpecDo) Assign(attrs ...field.AssignExpr) IWoodSpecDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w woodSpecDo) Joins(fields ...field.RelationField) IWoodSpecDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w woodSpecDo) Preload(fields ...field.RelationField) IWoodSpecDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w woodSpecDo) FirstOrInit() (*models.WoodSpec, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.WoodSpec), nil
	}
}

func (w woodSpecDo) FirstOrCreate() (*models.WoodSpec, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.WoodSpec), nil
	}
}

func (w woodSpecDo) FindByPage(offset int, limit int) (result []*models.WoodSpec, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w woodSpecDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w woodSpecDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w woodSpecDo) Delete(models ...*models.WoodSpec) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *woodSpecDo) withDO(do gen.Dao) *woodSpecDo {
	w.DO = *do.(*gen.DO)
	return w
}
