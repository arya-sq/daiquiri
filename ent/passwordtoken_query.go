// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/arya-sq/daiquiri/ent/passwordtoken"
	"github.com/arya-sq/daiquiri/ent/predicate"
	"github.com/arya-sq/daiquiri/ent/user"
)

// PasswordTokenQuery is the builder for querying PasswordToken entities.
type PasswordTokenQuery struct {
	config
	ctx        *QueryContext
	order      []passwordtoken.OrderOption
	inters     []Interceptor
	predicates []predicate.PasswordToken
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PasswordTokenQuery builder.
func (ptq *PasswordTokenQuery) Where(ps ...predicate.PasswordToken) *PasswordTokenQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PasswordTokenQuery) Limit(limit int) *PasswordTokenQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PasswordTokenQuery) Offset(offset int) *PasswordTokenQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PasswordTokenQuery) Unique(unique bool) *PasswordTokenQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PasswordTokenQuery) Order(o ...passwordtoken.OrderOption) *PasswordTokenQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryUser chains the current query on the "user" edge.
func (ptq *PasswordTokenQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(passwordtoken.Table, passwordtoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, passwordtoken.UserTable, passwordtoken.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PasswordToken entity from the query.
// Returns a *NotFoundError when no PasswordToken was found.
func (ptq *PasswordTokenQuery) First(ctx context.Context) (*PasswordToken, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{passwordtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PasswordTokenQuery) FirstX(ctx context.Context) *PasswordToken {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PasswordToken ID from the query.
// Returns a *NotFoundError when no PasswordToken ID was found.
func (ptq *PasswordTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{passwordtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PasswordTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PasswordToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PasswordToken entity is found.
// Returns a *NotFoundError when no PasswordToken entities are found.
func (ptq *PasswordTokenQuery) Only(ctx context.Context) (*PasswordToken, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{passwordtoken.Label}
	default:
		return nil, &NotSingularError{passwordtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PasswordTokenQuery) OnlyX(ctx context.Context) *PasswordToken {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PasswordToken ID in the query.
// Returns a *NotSingularError when more than one PasswordToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PasswordTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{passwordtoken.Label}
	default:
		err = &NotSingularError{passwordtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PasswordTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PasswordTokens.
func (ptq *PasswordTokenQuery) All(ctx context.Context) ([]*PasswordToken, error) {
	ctx = setContextOp(ctx, ptq.ctx, "All")
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PasswordToken, *PasswordTokenQuery]()
	return withInterceptors[[]*PasswordToken](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PasswordTokenQuery) AllX(ctx context.Context) []*PasswordToken {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PasswordToken IDs.
func (ptq *PasswordTokenQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, "IDs")
	if err = ptq.Select(passwordtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PasswordTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PasswordTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Count")
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PasswordTokenQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PasswordTokenQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PasswordTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Exist")
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PasswordTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PasswordTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PasswordTokenQuery) Clone() *PasswordTokenQuery {
	if ptq == nil {
		return nil
	}
	return &PasswordTokenQuery{
		config:     ptq.config,
		ctx:        ptq.ctx.Clone(),
		order:      append([]passwordtoken.OrderOption{}, ptq.order...),
		inters:     append([]Interceptor{}, ptq.inters...),
		predicates: append([]predicate.PasswordToken{}, ptq.predicates...),
		withUser:   ptq.withUser.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PasswordTokenQuery) WithUser(opts ...func(*UserQuery)) *PasswordTokenQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withUser = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PasswordToken.Query().
//		GroupBy(passwordtoken.FieldHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PasswordTokenQuery) GroupBy(field string, fields ...string) *PasswordTokenGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PasswordTokenGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = passwordtoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash,omitempty"`
//	}
//
//	client.PasswordToken.Query().
//		Select(passwordtoken.FieldHash).
//		Scan(ctx, &v)
func (ptq *PasswordTokenQuery) Select(fields ...string) *PasswordTokenSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PasswordTokenSelect{PasswordTokenQuery: ptq}
	sbuild.label = passwordtoken.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PasswordTokenSelect configured with the given aggregations.
func (ptq *PasswordTokenQuery) Aggregate(fns ...AggregateFunc) *PasswordTokenSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PasswordTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !passwordtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PasswordTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PasswordToken, error) {
	var (
		nodes       = []*PasswordToken{}
		withFKs     = ptq.withFKs
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withUser != nil,
		}
	)
	if ptq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PasswordToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PasswordToken{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withUser; query != nil {
		if err := ptq.loadUser(ctx, query, nodes, nil,
			func(n *PasswordToken, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PasswordTokenQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*PasswordToken, init func(*PasswordToken), assign func(*PasswordToken, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PasswordToken)
	for i := range nodes {
		if nodes[i].password_token_user == nil {
			continue
		}
		fk := *nodes[i].password_token_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "password_token_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptq *PasswordTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PasswordTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(passwordtoken.Table, passwordtoken.Columns, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.FieldID)
		for i := range fields {
			if fields[i] != passwordtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PasswordTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(passwordtoken.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = passwordtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PasswordTokenGroupBy is the group-by builder for PasswordToken entities.
type PasswordTokenGroupBy struct {
	selector
	build *PasswordTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PasswordTokenGroupBy) Aggregate(fns ...AggregateFunc) *PasswordTokenGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PasswordTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, "GroupBy")
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordTokenQuery, *PasswordTokenGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PasswordTokenGroupBy) sqlScan(ctx context.Context, root *PasswordTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PasswordTokenSelect is the builder for selecting fields of PasswordToken entities.
type PasswordTokenSelect struct {
	*PasswordTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PasswordTokenSelect) Aggregate(fns ...AggregateFunc) *PasswordTokenSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PasswordTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, "Select")
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordTokenQuery, *PasswordTokenSelect](ctx, pts.PasswordTokenQuery, pts, pts.inters, v)
}

func (pts *PasswordTokenSelect) sqlScan(ctx context.Context, root *PasswordTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
