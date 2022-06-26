// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/miscord-dev/toxfu/persistent/ent/address"
	"github.com/miscord-dev/toxfu/persistent/ent/node"
	"github.com/miscord-dev/toxfu/persistent/ent/route"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPublicKey sets the "public_key" field.
func (nc *NodeCreate) SetPublicKey(s string) *NodeCreate {
	nc.mutation.SetPublicKey(s)
	return nc
}

// SetPublicDiscoKey sets the "public_disco_key" field.
func (nc *NodeCreate) SetPublicDiscoKey(s string) *NodeCreate {
	nc.mutation.SetPublicDiscoKey(s)
	return nc
}

// SetHostName sets the "host_name" field.
func (nc *NodeCreate) SetHostName(s string) *NodeCreate {
	nc.mutation.SetHostName(s)
	return nc
}

// SetOs sets the "os" field.
func (nc *NodeCreate) SetOs(s string) *NodeCreate {
	nc.mutation.SetOs(s)
	return nc
}

// SetGoos sets the "goos" field.
func (nc *NodeCreate) SetGoos(s string) *NodeCreate {
	nc.mutation.SetGoos(s)
	return nc
}

// SetGoarch sets the "goarch" field.
func (nc *NodeCreate) SetGoarch(s string) *NodeCreate {
	nc.mutation.SetGoarch(s)
	return nc
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (nc *NodeCreate) SetLastUpdatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetLastUpdatedAt(t)
	return nc
}

// SetEndpoints sets the "endpoints" field.
func (nc *NodeCreate) SetEndpoints(s []string) *NodeCreate {
	nc.mutation.SetEndpoints(s)
	return nc
}

// SetState sets the "state" field.
func (nc *NodeCreate) SetState(n node.State) *NodeCreate {
	nc.mutation.SetState(n)
	return nc
}

// SetID sets the "id" field.
func (nc *NodeCreate) SetID(i int64) *NodeCreate {
	nc.mutation.SetID(i)
	return nc
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (nc *NodeCreate) AddRouteIDs(ids ...int64) *NodeCreate {
	nc.mutation.AddRouteIDs(ids...)
	return nc
}

// AddRoutes adds the "routes" edges to the Route entity.
func (nc *NodeCreate) AddRoutes(r ...*Route) *NodeCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return nc.AddRouteIDs(ids...)
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (nc *NodeCreate) AddAddressIDs(ids ...int64) *NodeCreate {
	nc.mutation.AddAddressIDs(ids...)
	return nc
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (nc *NodeCreate) AddAddresses(a ...*Address) *NodeCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nc.AddAddressIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "Node.public_key"`)}
	}
	if _, ok := nc.mutation.PublicDiscoKey(); !ok {
		return &ValidationError{Name: "public_disco_key", err: errors.New(`ent: missing required field "Node.public_disco_key"`)}
	}
	if _, ok := nc.mutation.HostName(); !ok {
		return &ValidationError{Name: "host_name", err: errors.New(`ent: missing required field "Node.host_name"`)}
	}
	if _, ok := nc.mutation.Os(); !ok {
		return &ValidationError{Name: "os", err: errors.New(`ent: missing required field "Node.os"`)}
	}
	if _, ok := nc.mutation.Goos(); !ok {
		return &ValidationError{Name: "goos", err: errors.New(`ent: missing required field "Node.goos"`)}
	}
	if _, ok := nc.mutation.Goarch(); !ok {
		return &ValidationError{Name: "goarch", err: errors.New(`ent: missing required field "Node.goarch"`)}
	}
	if _, ok := nc.mutation.LastUpdatedAt(); !ok {
		return &ValidationError{Name: "last_updated_at", err: errors.New(`ent: missing required field "Node.last_updated_at"`)}
	}
	if _, ok := nc.mutation.Endpoints(); !ok {
		return &ValidationError{Name: "endpoints", err: errors.New(`ent: missing required field "Node.endpoints"`)}
	}
	if _, ok := nc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Node.state"`)}
	}
	if v, ok := nc.mutation.State(); ok {
		if err := node.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Node.state": %w`, err)}
		}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: node.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: node.FieldID,
			},
		}
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.PublicKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicKey,
		})
		_node.PublicKey = value
	}
	if value, ok := nc.mutation.PublicDiscoKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldPublicDiscoKey,
		})
		_node.PublicDiscoKey = value
	}
	if value, ok := nc.mutation.HostName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldHostName,
		})
		_node.HostName = value
	}
	if value, ok := nc.mutation.Os(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldOs,
		})
		_node.Os = value
	}
	if value, ok := nc.mutation.Goos(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldGoos,
		})
		_node.Goos = value
	}
	if value, ok := nc.mutation.Goarch(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: node.FieldGoarch,
		})
		_node.Goarch = value
	}
	if value, ok := nc.mutation.LastUpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: node.FieldLastUpdatedAt,
		})
		_node.LastUpdatedAt = value
	}
	if value, ok := nc.mutation.Endpoints(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: node.FieldEndpoints,
		})
		_node.Endpoints = value
	}
	if value, ok := nc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: node.FieldState,
		})
		_node.State = value
	}
	if nodes := nc.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.RoutesTable,
			Columns: []string{node.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.AddressesTable,
			Columns: []string{node.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Node.Create().
//		SetPublicKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeUpsert) {
//			SetPublicKey(v+v).
//		}).
//		Exec(ctx)
//
func (nc *NodeCreate) OnConflict(opts ...sql.ConflictOption) *NodeUpsertOne {
	nc.conflict = opts
	return &NodeUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nc *NodeCreate) OnConflictColumns(columns ...string) *NodeUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NodeUpsertOne{
		create: nc,
	}
}

type (
	// NodeUpsertOne is the builder for "upsert"-ing
	//  one Node node.
	NodeUpsertOne struct {
		create *NodeCreate
	}

	// NodeUpsert is the "OnConflict" setter.
	NodeUpsert struct {
		*sql.UpdateSet
	}
)

// SetPublicKey sets the "public_key" field.
func (u *NodeUpsert) SetPublicKey(v string) *NodeUpsert {
	u.Set(node.FieldPublicKey, v)
	return u
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *NodeUpsert) UpdatePublicKey() *NodeUpsert {
	u.SetExcluded(node.FieldPublicKey)
	return u
}

// SetPublicDiscoKey sets the "public_disco_key" field.
func (u *NodeUpsert) SetPublicDiscoKey(v string) *NodeUpsert {
	u.Set(node.FieldPublicDiscoKey, v)
	return u
}

// UpdatePublicDiscoKey sets the "public_disco_key" field to the value that was provided on create.
func (u *NodeUpsert) UpdatePublicDiscoKey() *NodeUpsert {
	u.SetExcluded(node.FieldPublicDiscoKey)
	return u
}

// SetHostName sets the "host_name" field.
func (u *NodeUpsert) SetHostName(v string) *NodeUpsert {
	u.Set(node.FieldHostName, v)
	return u
}

// UpdateHostName sets the "host_name" field to the value that was provided on create.
func (u *NodeUpsert) UpdateHostName() *NodeUpsert {
	u.SetExcluded(node.FieldHostName)
	return u
}

// SetOs sets the "os" field.
func (u *NodeUpsert) SetOs(v string) *NodeUpsert {
	u.Set(node.FieldOs, v)
	return u
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *NodeUpsert) UpdateOs() *NodeUpsert {
	u.SetExcluded(node.FieldOs)
	return u
}

// SetGoos sets the "goos" field.
func (u *NodeUpsert) SetGoos(v string) *NodeUpsert {
	u.Set(node.FieldGoos, v)
	return u
}

// UpdateGoos sets the "goos" field to the value that was provided on create.
func (u *NodeUpsert) UpdateGoos() *NodeUpsert {
	u.SetExcluded(node.FieldGoos)
	return u
}

// SetGoarch sets the "goarch" field.
func (u *NodeUpsert) SetGoarch(v string) *NodeUpsert {
	u.Set(node.FieldGoarch, v)
	return u
}

// UpdateGoarch sets the "goarch" field to the value that was provided on create.
func (u *NodeUpsert) UpdateGoarch() *NodeUpsert {
	u.SetExcluded(node.FieldGoarch)
	return u
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (u *NodeUpsert) SetLastUpdatedAt(v time.Time) *NodeUpsert {
	u.Set(node.FieldLastUpdatedAt, v)
	return u
}

// UpdateLastUpdatedAt sets the "last_updated_at" field to the value that was provided on create.
func (u *NodeUpsert) UpdateLastUpdatedAt() *NodeUpsert {
	u.SetExcluded(node.FieldLastUpdatedAt)
	return u
}

// SetEndpoints sets the "endpoints" field.
func (u *NodeUpsert) SetEndpoints(v []string) *NodeUpsert {
	u.Set(node.FieldEndpoints, v)
	return u
}

// UpdateEndpoints sets the "endpoints" field to the value that was provided on create.
func (u *NodeUpsert) UpdateEndpoints() *NodeUpsert {
	u.SetExcluded(node.FieldEndpoints)
	return u
}

// SetState sets the "state" field.
func (u *NodeUpsert) SetState(v node.State) *NodeUpsert {
	u.Set(node.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *NodeUpsert) UpdateState() *NodeUpsert {
	u.SetExcluded(node.FieldState)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(node.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NodeUpsertOne) UpdateNewValues() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(node.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Node.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NodeUpsertOne) Ignore() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeUpsertOne) DoNothing() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeCreate.OnConflict
// documentation for more info.
func (u *NodeUpsertOne) Update(set func(*NodeUpsert)) *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetPublicKey sets the "public_key" field.
func (u *NodeUpsertOne) SetPublicKey(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdatePublicKey() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePublicKey()
	})
}

// SetPublicDiscoKey sets the "public_disco_key" field.
func (u *NodeUpsertOne) SetPublicDiscoKey(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetPublicDiscoKey(v)
	})
}

// UpdatePublicDiscoKey sets the "public_disco_key" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdatePublicDiscoKey() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePublicDiscoKey()
	})
}

// SetHostName sets the "host_name" field.
func (u *NodeUpsertOne) SetHostName(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetHostName(v)
	})
}

// UpdateHostName sets the "host_name" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateHostName() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateHostName()
	})
}

// SetOs sets the "os" field.
func (u *NodeUpsertOne) SetOs(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateOs() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateOs()
	})
}

// SetGoos sets the "goos" field.
func (u *NodeUpsertOne) SetGoos(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetGoos(v)
	})
}

// UpdateGoos sets the "goos" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateGoos() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateGoos()
	})
}

// SetGoarch sets the "goarch" field.
func (u *NodeUpsertOne) SetGoarch(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetGoarch(v)
	})
}

// UpdateGoarch sets the "goarch" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateGoarch() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateGoarch()
	})
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (u *NodeUpsertOne) SetLastUpdatedAt(v time.Time) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetLastUpdatedAt(v)
	})
}

// UpdateLastUpdatedAt sets the "last_updated_at" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateLastUpdatedAt() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateLastUpdatedAt()
	})
}

// SetEndpoints sets the "endpoints" field.
func (u *NodeUpsertOne) SetEndpoints(v []string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetEndpoints(v)
	})
}

// UpdateEndpoints sets the "endpoints" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateEndpoints() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateEndpoints()
	})
}

// SetState sets the "state" field.
func (u *NodeUpsertOne) SetState(v node.State) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateState() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateState()
	})
}

// Exec executes the query.
func (u *NodeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NodeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NodeUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NodeUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
	conflict []sql.ConflictOption
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Node.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeUpsert) {
//			SetPublicKey(v+v).
//		}).
//		Exec(ctx)
//
func (ncb *NodeCreateBulk) OnConflict(opts ...sql.ConflictOption) *NodeUpsertBulk {
	ncb.conflict = opts
	return &NodeUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ncb *NodeCreateBulk) OnConflictColumns(columns ...string) *NodeUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NodeUpsertBulk{
		create: ncb,
	}
}

// NodeUpsertBulk is the builder for "upsert"-ing
// a bulk of Node nodes.
type NodeUpsertBulk struct {
	create *NodeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(node.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NodeUpsertBulk) UpdateNewValues() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(node.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NodeUpsertBulk) Ignore() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeUpsertBulk) DoNothing() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeCreateBulk.OnConflict
// documentation for more info.
func (u *NodeUpsertBulk) Update(set func(*NodeUpsert)) *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetPublicKey sets the "public_key" field.
func (u *NodeUpsertBulk) SetPublicKey(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdatePublicKey() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePublicKey()
	})
}

// SetPublicDiscoKey sets the "public_disco_key" field.
func (u *NodeUpsertBulk) SetPublicDiscoKey(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetPublicDiscoKey(v)
	})
}

// UpdatePublicDiscoKey sets the "public_disco_key" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdatePublicDiscoKey() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePublicDiscoKey()
	})
}

// SetHostName sets the "host_name" field.
func (u *NodeUpsertBulk) SetHostName(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetHostName(v)
	})
}

// UpdateHostName sets the "host_name" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateHostName() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateHostName()
	})
}

// SetOs sets the "os" field.
func (u *NodeUpsertBulk) SetOs(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateOs() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateOs()
	})
}

// SetGoos sets the "goos" field.
func (u *NodeUpsertBulk) SetGoos(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetGoos(v)
	})
}

// UpdateGoos sets the "goos" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateGoos() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateGoos()
	})
}

// SetGoarch sets the "goarch" field.
func (u *NodeUpsertBulk) SetGoarch(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetGoarch(v)
	})
}

// UpdateGoarch sets the "goarch" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateGoarch() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateGoarch()
	})
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (u *NodeUpsertBulk) SetLastUpdatedAt(v time.Time) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetLastUpdatedAt(v)
	})
}

// UpdateLastUpdatedAt sets the "last_updated_at" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateLastUpdatedAt() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateLastUpdatedAt()
	})
}

// SetEndpoints sets the "endpoints" field.
func (u *NodeUpsertBulk) SetEndpoints(v []string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetEndpoints(v)
	})
}

// UpdateEndpoints sets the "endpoints" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateEndpoints() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateEndpoints()
	})
}

// SetState sets the "state" field.
func (u *NodeUpsertBulk) SetState(v node.State) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateState() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateState()
	})
}

// Exec executes the query.
func (u *NodeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NodeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NodeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}