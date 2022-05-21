// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/squadra-ricordo/ent/predicate"
	"github.com/squadra-ricordo/ent/skill"
	"github.com/squadra-ricordo/ent/userskill"
)

// SkillUpdate is the builder for updating Skill entities.
type SkillUpdate struct {
	config
	hooks    []Hook
	mutation *SkillMutation
}

// Where appends a list predicates to the SkillUpdate builder.
func (su *SkillUpdate) Where(ps ...predicate.Skill) *SkillUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SkillUpdate) SetName(s string) *SkillUpdate {
	su.mutation.SetName(s)
	return su
}

// AddUserSkillIDs adds the "user_skills" edge to the UserSkill entity by IDs.
func (su *SkillUpdate) AddUserSkillIDs(ids ...int) *SkillUpdate {
	su.mutation.AddUserSkillIDs(ids...)
	return su
}

// AddUserSkills adds the "user_skills" edges to the UserSkill entity.
func (su *SkillUpdate) AddUserSkills(u ...*UserSkill) *SkillUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserSkillIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (su *SkillUpdate) Mutation() *SkillMutation {
	return su.mutation
}

// ClearUserSkills clears all "user_skills" edges to the UserSkill entity.
func (su *SkillUpdate) ClearUserSkills() *SkillUpdate {
	su.mutation.ClearUserSkills()
	return su
}

// RemoveUserSkillIDs removes the "user_skills" edge to UserSkill entities by IDs.
func (su *SkillUpdate) RemoveUserSkillIDs(ids ...int) *SkillUpdate {
	su.mutation.RemoveUserSkillIDs(ids...)
	return su
}

// RemoveUserSkills removes "user_skills" edges to UserSkill entities.
func (su *SkillUpdate) RemoveUserSkills(u ...*UserSkill) *SkillUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserSkillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SkillUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SkillUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SkillUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SkillUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skill.Table,
			Columns: skill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skill.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skill.FieldName,
		})
	}
	if su.mutation.UserSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUserSkillsIDs(); len(nodes) > 0 && !su.mutation.UserSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SkillUpdateOne is the builder for updating a single Skill entity.
type SkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillMutation
}

// SetName sets the "name" field.
func (suo *SkillUpdateOne) SetName(s string) *SkillUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AddUserSkillIDs adds the "user_skills" edge to the UserSkill entity by IDs.
func (suo *SkillUpdateOne) AddUserSkillIDs(ids ...int) *SkillUpdateOne {
	suo.mutation.AddUserSkillIDs(ids...)
	return suo
}

// AddUserSkills adds the "user_skills" edges to the UserSkill entity.
func (suo *SkillUpdateOne) AddUserSkills(u ...*UserSkill) *SkillUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserSkillIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (suo *SkillUpdateOne) Mutation() *SkillMutation {
	return suo.mutation
}

// ClearUserSkills clears all "user_skills" edges to the UserSkill entity.
func (suo *SkillUpdateOne) ClearUserSkills() *SkillUpdateOne {
	suo.mutation.ClearUserSkills()
	return suo
}

// RemoveUserSkillIDs removes the "user_skills" edge to UserSkill entities by IDs.
func (suo *SkillUpdateOne) RemoveUserSkillIDs(ids ...int) *SkillUpdateOne {
	suo.mutation.RemoveUserSkillIDs(ids...)
	return suo
}

// RemoveUserSkills removes "user_skills" edges to UserSkill entities.
func (suo *SkillUpdateOne) RemoveUserSkills(u ...*UserSkill) *SkillUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserSkillIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SkillUpdateOne) Select(field string, fields ...string) *SkillUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Skill entity.
func (suo *SkillUpdateOne) Save(ctx context.Context) (*Skill, error) {
	var (
		err  error
		node *Skill
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SkillUpdateOne) SaveX(ctx context.Context) *Skill {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SkillUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SkillUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SkillUpdateOne) sqlSave(ctx context.Context) (_node *Skill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skill.Table,
			Columns: skill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skill.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Skill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skill.FieldID)
		for _, f := range fields {
			if !skill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skill.FieldName,
		})
	}
	if suo.mutation.UserSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUserSkillsIDs(); len(nodes) > 0 && !suo.mutation.UserSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.UserSkillsTable,
			Columns: []string{skill.UserSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Skill{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
