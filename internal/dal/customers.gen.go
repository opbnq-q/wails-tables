// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"app/internal/models"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCustomer(db *gorm.DB, opts ...gen.DOOption) customer {
	_customer := customer{}

	_customer.customerDo.UseDB(db, opts...)
	_customer.customerDo.UseModel(&models.Customer{})

	tableName := _customer.customerDo.TableName()
	_customer.ALL = field.NewAsterisk(tableName)
	_customer.Id = field.NewUint(tableName, "id")
	_customer.Title = field.NewString(tableName, "title")
	_customer.Contact = field.NewString(tableName, "contact")
	_customer.Orders = customerHasManyOrders{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Orders", "models.Order"),
		ProductType: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Orders.ProductType", "models.ProductType"),
		},
		Customer: struct {
			field.RelationField
			Orders struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Orders.Customer", "models.Customer"),
			Orders: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Orders.Customer.Orders", "models.Order"),
			},
		},
		Tasks: struct {
			field.RelationField
			ProductType struct {
				field.RelationField
			}
			Order struct {
				field.RelationField
			}
			PrepTasks struct {
				field.RelationField
				Task struct {
					field.RelationField
				}
				WorkArea struct {
					field.RelationField
					Workshop struct {
						field.RelationField
						WorkAreas struct {
							field.RelationField
						}
						Workers struct {
							field.RelationField
							Workshop struct {
								field.RelationField
							}
							TeamTasks struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}
						}
						Tasks struct {
							field.RelationField
						}
					}
					PrepTasks struct {
						field.RelationField
					}
					Shifts struct {
						field.RelationField
						ProductType struct {
							field.RelationField
						}
						WorkArea struct {
							field.RelationField
						}
					}
					TeamTasks struct {
						field.RelationField
					}
				}
			}
			Workshops struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Orders.Tasks", "models.Task"),
			ProductType: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Orders.Tasks.ProductType", "models.ProductType"),
			},
			Order: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Orders.Tasks.Order", "models.Order"),
			},
			PrepTasks: struct {
				field.RelationField
				Task struct {
					field.RelationField
				}
				WorkArea struct {
					field.RelationField
					Workshop struct {
						field.RelationField
						WorkAreas struct {
							field.RelationField
						}
						Workers struct {
							field.RelationField
							Workshop struct {
								field.RelationField
							}
							TeamTasks struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}
						}
						Tasks struct {
							field.RelationField
						}
					}
					PrepTasks struct {
						field.RelationField
					}
					Shifts struct {
						field.RelationField
						ProductType struct {
							field.RelationField
						}
						WorkArea struct {
							field.RelationField
						}
					}
					TeamTasks struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Orders.Tasks.PrepTasks", "models.PrepTask"),
				Task: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Orders.Tasks.PrepTasks.Task", "models.Task"),
				},
				WorkArea: struct {
					field.RelationField
					Workshop struct {
						field.RelationField
						WorkAreas struct {
							field.RelationField
						}
						Workers struct {
							field.RelationField
							Workshop struct {
								field.RelationField
							}
							TeamTasks struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}
						}
						Tasks struct {
							field.RelationField
						}
					}
					PrepTasks struct {
						field.RelationField
					}
					Shifts struct {
						field.RelationField
						ProductType struct {
							field.RelationField
						}
						WorkArea struct {
							field.RelationField
						}
					}
					TeamTasks struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea", "models.WorkArea"),
					Workshop: struct {
						field.RelationField
						WorkAreas struct {
							field.RelationField
						}
						Workers struct {
							field.RelationField
							Workshop struct {
								field.RelationField
							}
							TeamTasks struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}
						}
						Tasks struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop", "models.Workshop"),
						WorkAreas: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.WorkAreas", "models.WorkArea"),
						},
						Workers: struct {
							field.RelationField
							Workshop struct {
								field.RelationField
							}
							TeamTasks struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}
						}{
							RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers", "models.Worker"),
							Workshop: struct {
								field.RelationField
							}{
								RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.Workshop", "models.Workshop"),
							},
							TeamTasks: struct {
								field.RelationField
								TeamType struct {
									field.RelationField
								}
								TeamLeader struct {
									field.RelationField
								}
								WorkArea struct {
									field.RelationField
								}
								TeamMembers struct {
									field.RelationField
								}
							}{
								RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.TeamTasks", "models.TeamTask"),
								TeamType: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.TeamTasks.TeamType", "models.TeamType"),
								},
								TeamLeader: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.TeamTasks.TeamLeader", "models.Worker"),
								},
								WorkArea: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.TeamTasks.WorkArea", "models.WorkArea"),
								},
								TeamMembers: struct {
									field.RelationField
								}{
									RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Workers.TeamTasks.TeamMembers", "models.Worker"),
								},
							},
						},
						Tasks: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Workshop.Tasks", "models.Task"),
						},
					},
					PrepTasks: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.PrepTasks", "models.PrepTask"),
					},
					Shifts: struct {
						field.RelationField
						ProductType struct {
							field.RelationField
						}
						WorkArea struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Shifts", "models.Shift"),
						ProductType: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Shifts.ProductType", "models.ProductType"),
						},
						WorkArea: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.Shifts.WorkArea", "models.WorkArea"),
						},
					},
					TeamTasks: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Orders.Tasks.PrepTasks.WorkArea.TeamTasks", "models.TeamTask"),
					},
				},
			},
			Workshops: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Orders.Tasks.Workshops", "models.Workshop"),
			},
		},
	}

	_customer.fillFieldMap()

	return _customer
}

type customer struct {
	customerDo

	ALL     field.Asterisk
	Id      field.Uint
	Title   field.String
	Contact field.String
	Orders  customerHasManyOrders

	fieldMap map[string]field.Expr
}

func (c customer) Table(newTableName string) *customer {
	c.customerDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c customer) As(alias string) *customer {
	c.customerDo.DO = *(c.customerDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *customer) updateTableName(table string) *customer {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint(table, "id")
	c.Title = field.NewString(table, "title")
	c.Contact = field.NewString(table, "contact")

	c.fillFieldMap()

	return c
}

func (c *customer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *customer) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["id"] = c.Id
	c.fieldMap["title"] = c.Title
	c.fieldMap["contact"] = c.Contact

}

func (c customer) clone(db *gorm.DB) customer {
	c.customerDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c customer) replaceDB(db *gorm.DB) customer {
	c.customerDo.ReplaceDB(db)
	return c
}

type customerHasManyOrders struct {
	db *gorm.DB

	field.RelationField

	ProductType struct {
		field.RelationField
	}
	Customer struct {
		field.RelationField
		Orders struct {
			field.RelationField
		}
	}
	Tasks struct {
		field.RelationField
		ProductType struct {
			field.RelationField
		}
		Order struct {
			field.RelationField
		}
		PrepTasks struct {
			field.RelationField
			Task struct {
				field.RelationField
			}
			WorkArea struct {
				field.RelationField
				Workshop struct {
					field.RelationField
					WorkAreas struct {
						field.RelationField
					}
					Workers struct {
						field.RelationField
						Workshop struct {
							field.RelationField
						}
						TeamTasks struct {
							field.RelationField
							TeamType struct {
								field.RelationField
							}
							TeamLeader struct {
								field.RelationField
							}
							WorkArea struct {
								field.RelationField
							}
							TeamMembers struct {
								field.RelationField
							}
						}
					}
					Tasks struct {
						field.RelationField
					}
				}
				PrepTasks struct {
					field.RelationField
				}
				Shifts struct {
					field.RelationField
					ProductType struct {
						field.RelationField
					}
					WorkArea struct {
						field.RelationField
					}
				}
				TeamTasks struct {
					field.RelationField
				}
			}
		}
		Workshops struct {
			field.RelationField
		}
	}
}

func (a customerHasManyOrders) Where(conds ...field.Expr) *customerHasManyOrders {
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

func (a customerHasManyOrders) WithContext(ctx context.Context) *customerHasManyOrders {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a customerHasManyOrders) Session(session *gorm.Session) *customerHasManyOrders {
	a.db = a.db.Session(session)
	return &a
}

func (a customerHasManyOrders) Model(m *models.Customer) *customerHasManyOrdersTx {
	return &customerHasManyOrdersTx{a.db.Model(m).Association(a.Name())}
}

type customerHasManyOrdersTx struct{ tx *gorm.Association }

func (a customerHasManyOrdersTx) Find() (result []*models.Order, err error) {
	return result, a.tx.Find(&result)
}

func (a customerHasManyOrdersTx) Append(values ...*models.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a customerHasManyOrdersTx) Replace(values ...*models.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a customerHasManyOrdersTx) Delete(values ...*models.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a customerHasManyOrdersTx) Clear() error {
	return a.tx.Clear()
}

func (a customerHasManyOrdersTx) Count() int64 {
	return a.tx.Count()
}

type customerDo struct{ gen.DO }

type ICustomerDo interface {
	gen.SubQuery
	Debug() ICustomerDo
	WithContext(ctx context.Context) ICustomerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICustomerDo
	WriteDB() ICustomerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICustomerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICustomerDo
	Not(conds ...gen.Condition) ICustomerDo
	Or(conds ...gen.Condition) ICustomerDo
	Select(conds ...field.Expr) ICustomerDo
	Where(conds ...gen.Condition) ICustomerDo
	Order(conds ...field.Expr) ICustomerDo
	Distinct(cols ...field.Expr) ICustomerDo
	Omit(cols ...field.Expr) ICustomerDo
	Join(table schema.Tabler, on ...field.Expr) ICustomerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICustomerDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICustomerDo
	Group(cols ...field.Expr) ICustomerDo
	Having(conds ...gen.Condition) ICustomerDo
	Limit(limit int) ICustomerDo
	Offset(offset int) ICustomerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICustomerDo
	Unscoped() ICustomerDo
	Create(values ...*models.Customer) error
	CreateInBatches(values []*models.Customer, batchSize int) error
	Save(values ...*models.Customer) error
	First() (*models.Customer, error)
	Take() (*models.Customer, error)
	Last() (*models.Customer, error)
	Find() ([]*models.Customer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Customer, err error)
	FindInBatches(result *[]*models.Customer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Customer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICustomerDo
	Assign(attrs ...field.AssignExpr) ICustomerDo
	Joins(fields ...field.RelationField) ICustomerDo
	Preload(fields ...field.RelationField) ICustomerDo
	FirstOrInit() (*models.Customer, error)
	FirstOrCreate() (*models.Customer, error)
	FindByPage(offset int, limit int) (result []*models.Customer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICustomerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c customerDo) Debug() ICustomerDo {
	return c.withDO(c.DO.Debug())
}

func (c customerDo) WithContext(ctx context.Context) ICustomerDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customerDo) ReadDB() ICustomerDo {
	return c.Clauses(dbresolver.Read)
}

func (c customerDo) WriteDB() ICustomerDo {
	return c.Clauses(dbresolver.Write)
}

func (c customerDo) Session(config *gorm.Session) ICustomerDo {
	return c.withDO(c.DO.Session(config))
}

func (c customerDo) Clauses(conds ...clause.Expression) ICustomerDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customerDo) Returning(value interface{}, columns ...string) ICustomerDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c customerDo) Not(conds ...gen.Condition) ICustomerDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customerDo) Or(conds ...gen.Condition) ICustomerDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customerDo) Select(conds ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customerDo) Where(conds ...gen.Condition) ICustomerDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customerDo) Order(conds ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customerDo) Distinct(cols ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customerDo) Omit(cols ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customerDo) Join(table schema.Tabler, on ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customerDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customerDo) RightJoin(table schema.Tabler, on ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customerDo) Group(cols ...field.Expr) ICustomerDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customerDo) Having(conds ...gen.Condition) ICustomerDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customerDo) Limit(limit int) ICustomerDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customerDo) Offset(offset int) ICustomerDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICustomerDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customerDo) Unscoped() ICustomerDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customerDo) Create(values ...*models.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customerDo) CreateInBatches(values []*models.Customer, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customerDo) Save(values ...*models.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customerDo) First() (*models.Customer, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Customer), nil
	}
}

func (c customerDo) Take() (*models.Customer, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Customer), nil
	}
}

func (c customerDo) Last() (*models.Customer, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Customer), nil
	}
}

func (c customerDo) Find() ([]*models.Customer, error) {
	result, err := c.DO.Find()
	return result.([]*models.Customer), err
}

func (c customerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Customer, err error) {
	buf := make([]*models.Customer, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c customerDo) FindInBatches(result *[]*models.Customer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c customerDo) Attrs(attrs ...field.AssignExpr) ICustomerDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customerDo) Assign(attrs ...field.AssignExpr) ICustomerDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customerDo) Joins(fields ...field.RelationField) ICustomerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c customerDo) Preload(fields ...field.RelationField) ICustomerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c customerDo) FirstOrInit() (*models.Customer, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Customer), nil
	}
}

func (c customerDo) FirstOrCreate() (*models.Customer, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Customer), nil
	}
}

func (c customerDo) FindByPage(offset int, limit int) (result []*models.Customer, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c customerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c customerDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c customerDo) Delete(models ...*models.Customer) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *customerDo) withDO(do gen.Dao) *customerDo {
	c.DO = *do.(*gen.DO)
	return c
}
