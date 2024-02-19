// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KenshiTech/unchained/ent/assetprice"
	"github.com/KenshiTech/unchained/ent/predicate"
	"github.com/KenshiTech/unchained/ent/signer"
)

// AssetPriceUpdate is the builder for updating AssetPrice entities.
type AssetPriceUpdate struct {
	config
	hooks    []Hook
	mutation *AssetPriceMutation
}

// Where appends a list predicates to the AssetPriceUpdate builder.
func (apu *AssetPriceUpdate) Where(ps ...predicate.AssetPrice) *AssetPriceUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetBlock sets the "block" field.
func (apu *AssetPriceUpdate) SetBlock(u uint64) *AssetPriceUpdate {
	apu.mutation.ResetBlock()
	apu.mutation.SetBlock(u)
	return apu
}

// SetNillableBlock sets the "block" field if the given value is not nil.
func (apu *AssetPriceUpdate) SetNillableBlock(u *uint64) *AssetPriceUpdate {
	if u != nil {
		apu.SetBlock(*u)
	}
	return apu
}

// AddBlock adds u to the "block" field.
func (apu *AssetPriceUpdate) AddBlock(u int64) *AssetPriceUpdate {
	apu.mutation.AddBlock(u)
	return apu
}

// SetSignersCount sets the "signersCount" field.
func (apu *AssetPriceUpdate) SetSignersCount(u uint64) *AssetPriceUpdate {
	apu.mutation.ResetSignersCount()
	apu.mutation.SetSignersCount(u)
	return apu
}

// SetNillableSignersCount sets the "signersCount" field if the given value is not nil.
func (apu *AssetPriceUpdate) SetNillableSignersCount(u *uint64) *AssetPriceUpdate {
	if u != nil {
		apu.SetSignersCount(*u)
	}
	return apu
}

// AddSignersCount adds u to the "signersCount" field.
func (apu *AssetPriceUpdate) AddSignersCount(u int64) *AssetPriceUpdate {
	apu.mutation.AddSignersCount(u)
	return apu
}

// ClearSignersCount clears the value of the "signersCount" field.
func (apu *AssetPriceUpdate) ClearSignersCount() *AssetPriceUpdate {
	apu.mutation.ClearSignersCount()
	return apu
}

// SetPrice sets the "price" field.
func (apu *AssetPriceUpdate) SetPrice(b *big.Int) *AssetPriceUpdate {
	apu.mutation.SetPrice(b)
	return apu
}

// SetSignature sets the "signature" field.
func (apu *AssetPriceUpdate) SetSignature(b []byte) *AssetPriceUpdate {
	apu.mutation.SetSignature(b)
	return apu
}

// SetAsset sets the "asset" field.
func (apu *AssetPriceUpdate) SetAsset(s string) *AssetPriceUpdate {
	apu.mutation.SetAsset(s)
	return apu
}

// SetNillableAsset sets the "asset" field if the given value is not nil.
func (apu *AssetPriceUpdate) SetNillableAsset(s *string) *AssetPriceUpdate {
	if s != nil {
		apu.SetAsset(*s)
	}
	return apu
}

// ClearAsset clears the value of the "asset" field.
func (apu *AssetPriceUpdate) ClearAsset() *AssetPriceUpdate {
	apu.mutation.ClearAsset()
	return apu
}

// SetChain sets the "chain" field.
func (apu *AssetPriceUpdate) SetChain(s string) *AssetPriceUpdate {
	apu.mutation.SetChain(s)
	return apu
}

// SetNillableChain sets the "chain" field if the given value is not nil.
func (apu *AssetPriceUpdate) SetNillableChain(s *string) *AssetPriceUpdate {
	if s != nil {
		apu.SetChain(*s)
	}
	return apu
}

// ClearChain clears the value of the "chain" field.
func (apu *AssetPriceUpdate) ClearChain() *AssetPriceUpdate {
	apu.mutation.ClearChain()
	return apu
}

// SetPair sets the "pair" field.
func (apu *AssetPriceUpdate) SetPair(s string) *AssetPriceUpdate {
	apu.mutation.SetPair(s)
	return apu
}

// SetNillablePair sets the "pair" field if the given value is not nil.
func (apu *AssetPriceUpdate) SetNillablePair(s *string) *AssetPriceUpdate {
	if s != nil {
		apu.SetPair(*s)
	}
	return apu
}

// ClearPair clears the value of the "pair" field.
func (apu *AssetPriceUpdate) ClearPair() *AssetPriceUpdate {
	apu.mutation.ClearPair()
	return apu
}

// AddSignerIDs adds the "signers" edge to the Signer entity by IDs.
func (apu *AssetPriceUpdate) AddSignerIDs(ids ...int) *AssetPriceUpdate {
	apu.mutation.AddSignerIDs(ids...)
	return apu
}

// AddSigners adds the "signers" edges to the Signer entity.
func (apu *AssetPriceUpdate) AddSigners(s ...*Signer) *AssetPriceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apu.AddSignerIDs(ids...)
}

// Mutation returns the AssetPriceMutation object of the builder.
func (apu *AssetPriceUpdate) Mutation() *AssetPriceMutation {
	return apu.mutation
}

// ClearSigners clears all "signers" edges to the Signer entity.
func (apu *AssetPriceUpdate) ClearSigners() *AssetPriceUpdate {
	apu.mutation.ClearSigners()
	return apu
}

// RemoveSignerIDs removes the "signers" edge to Signer entities by IDs.
func (apu *AssetPriceUpdate) RemoveSignerIDs(ids ...int) *AssetPriceUpdate {
	apu.mutation.RemoveSignerIDs(ids...)
	return apu
}

// RemoveSigners removes "signers" edges to Signer entities.
func (apu *AssetPriceUpdate) RemoveSigners(s ...*Signer) *AssetPriceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apu.RemoveSignerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AssetPriceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AssetPriceUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AssetPriceUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AssetPriceUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AssetPriceUpdate) check() error {
	if v, ok := apu.mutation.Signature(); ok {
		if err := assetprice.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "AssetPrice.signature": %w`, err)}
		}
	}
	return nil
}

func (apu *AssetPriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(assetprice.Table, assetprice.Columns, sqlgraph.NewFieldSpec(assetprice.FieldID, field.TypeInt))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.Block(); ok {
		_spec.SetField(assetprice.FieldBlock, field.TypeUint64, value)
	}
	if value, ok := apu.mutation.AddedBlock(); ok {
		_spec.AddField(assetprice.FieldBlock, field.TypeUint64, value)
	}
	if value, ok := apu.mutation.SignersCount(); ok {
		_spec.SetField(assetprice.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := apu.mutation.AddedSignersCount(); ok {
		_spec.AddField(assetprice.FieldSignersCount, field.TypeUint64, value)
	}
	if apu.mutation.SignersCountCleared() {
		_spec.ClearField(assetprice.FieldSignersCount, field.TypeUint64)
	}
	if value, ok := apu.mutation.Price(); ok {
		vv, err := assetprice.ValueScanner.Price.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(assetprice.FieldPrice, field.TypeString, vv)
	}
	if value, ok := apu.mutation.Signature(); ok {
		_spec.SetField(assetprice.FieldSignature, field.TypeBytes, value)
	}
	if value, ok := apu.mutation.Asset(); ok {
		_spec.SetField(assetprice.FieldAsset, field.TypeString, value)
	}
	if apu.mutation.AssetCleared() {
		_spec.ClearField(assetprice.FieldAsset, field.TypeString)
	}
	if value, ok := apu.mutation.Chain(); ok {
		_spec.SetField(assetprice.FieldChain, field.TypeString, value)
	}
	if apu.mutation.ChainCleared() {
		_spec.ClearField(assetprice.FieldChain, field.TypeString)
	}
	if value, ok := apu.mutation.Pair(); ok {
		_spec.SetField(assetprice.FieldPair, field.TypeString, value)
	}
	if apu.mutation.PairCleared() {
		_spec.ClearField(assetprice.FieldPair, field.TypeString)
	}
	if apu.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedSignersIDs(); len(nodes) > 0 && !apu.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.SignersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AssetPriceUpdateOne is the builder for updating a single AssetPrice entity.
type AssetPriceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetPriceMutation
}

// SetBlock sets the "block" field.
func (apuo *AssetPriceUpdateOne) SetBlock(u uint64) *AssetPriceUpdateOne {
	apuo.mutation.ResetBlock()
	apuo.mutation.SetBlock(u)
	return apuo
}

// SetNillableBlock sets the "block" field if the given value is not nil.
func (apuo *AssetPriceUpdateOne) SetNillableBlock(u *uint64) *AssetPriceUpdateOne {
	if u != nil {
		apuo.SetBlock(*u)
	}
	return apuo
}

// AddBlock adds u to the "block" field.
func (apuo *AssetPriceUpdateOne) AddBlock(u int64) *AssetPriceUpdateOne {
	apuo.mutation.AddBlock(u)
	return apuo
}

// SetSignersCount sets the "signersCount" field.
func (apuo *AssetPriceUpdateOne) SetSignersCount(u uint64) *AssetPriceUpdateOne {
	apuo.mutation.ResetSignersCount()
	apuo.mutation.SetSignersCount(u)
	return apuo
}

// SetNillableSignersCount sets the "signersCount" field if the given value is not nil.
func (apuo *AssetPriceUpdateOne) SetNillableSignersCount(u *uint64) *AssetPriceUpdateOne {
	if u != nil {
		apuo.SetSignersCount(*u)
	}
	return apuo
}

// AddSignersCount adds u to the "signersCount" field.
func (apuo *AssetPriceUpdateOne) AddSignersCount(u int64) *AssetPriceUpdateOne {
	apuo.mutation.AddSignersCount(u)
	return apuo
}

// ClearSignersCount clears the value of the "signersCount" field.
func (apuo *AssetPriceUpdateOne) ClearSignersCount() *AssetPriceUpdateOne {
	apuo.mutation.ClearSignersCount()
	return apuo
}

// SetPrice sets the "price" field.
func (apuo *AssetPriceUpdateOne) SetPrice(b *big.Int) *AssetPriceUpdateOne {
	apuo.mutation.SetPrice(b)
	return apuo
}

// SetSignature sets the "signature" field.
func (apuo *AssetPriceUpdateOne) SetSignature(b []byte) *AssetPriceUpdateOne {
	apuo.mutation.SetSignature(b)
	return apuo
}

// SetAsset sets the "asset" field.
func (apuo *AssetPriceUpdateOne) SetAsset(s string) *AssetPriceUpdateOne {
	apuo.mutation.SetAsset(s)
	return apuo
}

// SetNillableAsset sets the "asset" field if the given value is not nil.
func (apuo *AssetPriceUpdateOne) SetNillableAsset(s *string) *AssetPriceUpdateOne {
	if s != nil {
		apuo.SetAsset(*s)
	}
	return apuo
}

// ClearAsset clears the value of the "asset" field.
func (apuo *AssetPriceUpdateOne) ClearAsset() *AssetPriceUpdateOne {
	apuo.mutation.ClearAsset()
	return apuo
}

// SetChain sets the "chain" field.
func (apuo *AssetPriceUpdateOne) SetChain(s string) *AssetPriceUpdateOne {
	apuo.mutation.SetChain(s)
	return apuo
}

// SetNillableChain sets the "chain" field if the given value is not nil.
func (apuo *AssetPriceUpdateOne) SetNillableChain(s *string) *AssetPriceUpdateOne {
	if s != nil {
		apuo.SetChain(*s)
	}
	return apuo
}

// ClearChain clears the value of the "chain" field.
func (apuo *AssetPriceUpdateOne) ClearChain() *AssetPriceUpdateOne {
	apuo.mutation.ClearChain()
	return apuo
}

// SetPair sets the "pair" field.
func (apuo *AssetPriceUpdateOne) SetPair(s string) *AssetPriceUpdateOne {
	apuo.mutation.SetPair(s)
	return apuo
}

// SetNillablePair sets the "pair" field if the given value is not nil.
func (apuo *AssetPriceUpdateOne) SetNillablePair(s *string) *AssetPriceUpdateOne {
	if s != nil {
		apuo.SetPair(*s)
	}
	return apuo
}

// ClearPair clears the value of the "pair" field.
func (apuo *AssetPriceUpdateOne) ClearPair() *AssetPriceUpdateOne {
	apuo.mutation.ClearPair()
	return apuo
}

// AddSignerIDs adds the "signers" edge to the Signer entity by IDs.
func (apuo *AssetPriceUpdateOne) AddSignerIDs(ids ...int) *AssetPriceUpdateOne {
	apuo.mutation.AddSignerIDs(ids...)
	return apuo
}

// AddSigners adds the "signers" edges to the Signer entity.
func (apuo *AssetPriceUpdateOne) AddSigners(s ...*Signer) *AssetPriceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apuo.AddSignerIDs(ids...)
}

// Mutation returns the AssetPriceMutation object of the builder.
func (apuo *AssetPriceUpdateOne) Mutation() *AssetPriceMutation {
	return apuo.mutation
}

// ClearSigners clears all "signers" edges to the Signer entity.
func (apuo *AssetPriceUpdateOne) ClearSigners() *AssetPriceUpdateOne {
	apuo.mutation.ClearSigners()
	return apuo
}

// RemoveSignerIDs removes the "signers" edge to Signer entities by IDs.
func (apuo *AssetPriceUpdateOne) RemoveSignerIDs(ids ...int) *AssetPriceUpdateOne {
	apuo.mutation.RemoveSignerIDs(ids...)
	return apuo
}

// RemoveSigners removes "signers" edges to Signer entities.
func (apuo *AssetPriceUpdateOne) RemoveSigners(s ...*Signer) *AssetPriceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apuo.RemoveSignerIDs(ids...)
}

// Where appends a list predicates to the AssetPriceUpdate builder.
func (apuo *AssetPriceUpdateOne) Where(ps ...predicate.AssetPrice) *AssetPriceUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AssetPriceUpdateOne) Select(field string, fields ...string) *AssetPriceUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AssetPrice entity.
func (apuo *AssetPriceUpdateOne) Save(ctx context.Context) (*AssetPrice, error) {
	return withHooks(ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AssetPriceUpdateOne) SaveX(ctx context.Context) *AssetPrice {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AssetPriceUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AssetPriceUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AssetPriceUpdateOne) check() error {
	if v, ok := apuo.mutation.Signature(); ok {
		if err := assetprice.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "AssetPrice.signature": %w`, err)}
		}
	}
	return nil
}

func (apuo *AssetPriceUpdateOne) sqlSave(ctx context.Context) (_node *AssetPrice, err error) {
	if err := apuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(assetprice.Table, assetprice.Columns, sqlgraph.NewFieldSpec(assetprice.FieldID, field.TypeInt))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssetPrice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assetprice.FieldID)
		for _, f := range fields {
			if !assetprice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != assetprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.Block(); ok {
		_spec.SetField(assetprice.FieldBlock, field.TypeUint64, value)
	}
	if value, ok := apuo.mutation.AddedBlock(); ok {
		_spec.AddField(assetprice.FieldBlock, field.TypeUint64, value)
	}
	if value, ok := apuo.mutation.SignersCount(); ok {
		_spec.SetField(assetprice.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := apuo.mutation.AddedSignersCount(); ok {
		_spec.AddField(assetprice.FieldSignersCount, field.TypeUint64, value)
	}
	if apuo.mutation.SignersCountCleared() {
		_spec.ClearField(assetprice.FieldSignersCount, field.TypeUint64)
	}
	if value, ok := apuo.mutation.Price(); ok {
		vv, err := assetprice.ValueScanner.Price.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(assetprice.FieldPrice, field.TypeString, vv)
	}
	if value, ok := apuo.mutation.Signature(); ok {
		_spec.SetField(assetprice.FieldSignature, field.TypeBytes, value)
	}
	if value, ok := apuo.mutation.Asset(); ok {
		_spec.SetField(assetprice.FieldAsset, field.TypeString, value)
	}
	if apuo.mutation.AssetCleared() {
		_spec.ClearField(assetprice.FieldAsset, field.TypeString)
	}
	if value, ok := apuo.mutation.Chain(); ok {
		_spec.SetField(assetprice.FieldChain, field.TypeString, value)
	}
	if apuo.mutation.ChainCleared() {
		_spec.ClearField(assetprice.FieldChain, field.TypeString)
	}
	if value, ok := apuo.mutation.Pair(); ok {
		_spec.SetField(assetprice.FieldPair, field.TypeString, value)
	}
	if apuo.mutation.PairCleared() {
		_spec.ClearField(assetprice.FieldPair, field.TypeString)
	}
	if apuo.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedSignersIDs(); len(nodes) > 0 && !apuo.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.SignersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   assetprice.SignersTable,
			Columns: assetprice.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AssetPrice{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}
