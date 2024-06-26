// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Invite is an object representing the database table.
type Invite struct {
	ID        string            `boil:"id" json:"id" toml:"id" yaml:"id"`
	Sender    string            `boil:"sender" json:"sender" toml:"sender" yaml:"sender"`
	Receiver  types.StringArray `boil:"receiver" json:"receiver" toml:"receiver" yaml:"receiver"`
	TimesUsed int               `boil:"times_used" json:"times_used" toml:"times_used" yaml:"times_used"`
	GroupID   string            `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	CreatedAt null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *inviteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L inviteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InviteColumns = struct {
	ID        string
	Sender    string
	Receiver  string
	TimesUsed string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Sender:    "sender",
	Receiver:  "receiver",
	TimesUsed: "times_used",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var InviteTableColumns = struct {
	ID        string
	Sender    string
	Receiver  string
	TimesUsed string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "invites.id",
	Sender:    "invites.sender",
	Receiver:  "invites.receiver",
	TimesUsed: "invites.times_used",
	GroupID:   "invites.group_id",
	CreatedAt: "invites.created_at",
	UpdatedAt: "invites.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var InviteWhere = struct {
	ID        whereHelperstring
	Sender    whereHelperstring
	Receiver  whereHelpertypes_StringArray
	TimesUsed whereHelperint
	GroupID   whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"invites\".\"id\""},
	Sender:    whereHelperstring{field: "\"invites\".\"sender\""},
	Receiver:  whereHelpertypes_StringArray{field: "\"invites\".\"receiver\""},
	TimesUsed: whereHelperint{field: "\"invites\".\"times_used\""},
	GroupID:   whereHelperstring{field: "\"invites\".\"group_id\""},
	CreatedAt: whereHelpernull_Time{field: "\"invites\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"invites\".\"updated_at\""},
}

// InviteRels is where relationship names are stored.
var InviteRels = struct {
	Group string
}{
	Group: "Group",
}

// inviteR is where relationships are stored.
type inviteR struct {
	Group *Group `boil:"Group" json:"Group" toml:"Group" yaml:"Group"`
}

// NewStruct creates a new relationship struct
func (*inviteR) NewStruct() *inviteR {
	return &inviteR{}
}

func (r *inviteR) GetGroup() *Group {
	if r == nil {
		return nil
	}
	return r.Group
}

// inviteL is where Load methods for each relationship are stored.
type inviteL struct{}

var (
	inviteAllColumns            = []string{"id", "sender", "receiver", "times_used", "group_id", "created_at", "updated_at"}
	inviteColumnsWithoutDefault = []string{"id", "sender", "receiver", "group_id"}
	inviteColumnsWithDefault    = []string{"times_used", "created_at", "updated_at"}
	invitePrimaryKeyColumns     = []string{"id"}
	inviteGeneratedColumns      = []string{}
)

type (
	// InviteSlice is an alias for a slice of pointers to Invite.
	// This should almost always be used instead of []Invite.
	InviteSlice []*Invite
	// InviteHook is the signature for custom Invite hook methods
	InviteHook func(context.Context, boil.ContextExecutor, *Invite) error

	inviteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	inviteType                 = reflect.TypeOf(&Invite{})
	inviteMapping              = queries.MakeStructMapping(inviteType)
	invitePrimaryKeyMapping, _ = queries.BindMapping(inviteType, inviteMapping, invitePrimaryKeyColumns)
	inviteInsertCacheMut       sync.RWMutex
	inviteInsertCache          = make(map[string]insertCache)
	inviteUpdateCacheMut       sync.RWMutex
	inviteUpdateCache          = make(map[string]updateCache)
	inviteUpsertCacheMut       sync.RWMutex
	inviteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var inviteAfterSelectHooks []InviteHook

var inviteBeforeInsertHooks []InviteHook
var inviteAfterInsertHooks []InviteHook

var inviteBeforeUpdateHooks []InviteHook
var inviteAfterUpdateHooks []InviteHook

var inviteBeforeDeleteHooks []InviteHook
var inviteAfterDeleteHooks []InviteHook

var inviteBeforeUpsertHooks []InviteHook
var inviteAfterUpsertHooks []InviteHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Invite) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Invite) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Invite) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Invite) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Invite) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Invite) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Invite) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Invite) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Invite) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInviteHook registers your hook function for all future operations.
func AddInviteHook(hookPoint boil.HookPoint, inviteHook InviteHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		inviteAfterSelectHooks = append(inviteAfterSelectHooks, inviteHook)
	case boil.BeforeInsertHook:
		inviteBeforeInsertHooks = append(inviteBeforeInsertHooks, inviteHook)
	case boil.AfterInsertHook:
		inviteAfterInsertHooks = append(inviteAfterInsertHooks, inviteHook)
	case boil.BeforeUpdateHook:
		inviteBeforeUpdateHooks = append(inviteBeforeUpdateHooks, inviteHook)
	case boil.AfterUpdateHook:
		inviteAfterUpdateHooks = append(inviteAfterUpdateHooks, inviteHook)
	case boil.BeforeDeleteHook:
		inviteBeforeDeleteHooks = append(inviteBeforeDeleteHooks, inviteHook)
	case boil.AfterDeleteHook:
		inviteAfterDeleteHooks = append(inviteAfterDeleteHooks, inviteHook)
	case boil.BeforeUpsertHook:
		inviteBeforeUpsertHooks = append(inviteBeforeUpsertHooks, inviteHook)
	case boil.AfterUpsertHook:
		inviteAfterUpsertHooks = append(inviteAfterUpsertHooks, inviteHook)
	}
}

// One returns a single invite record from the query.
func (q inviteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Invite, error) {
	o := &Invite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for invites")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Invite records from the query.
func (q inviteQuery) All(ctx context.Context, exec boil.ContextExecutor) (InviteSlice, error) {
	var o []*Invite

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Invite slice")
	}

	if len(inviteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Invite records in the query.
func (q inviteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count invites rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q inviteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if invites exists")
	}

	return count > 0, nil
}

// Group pointed to by the foreign key.
func (o *Invite) Group(mods ...qm.QueryMod) groupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	return Groups(queryMods...)
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (inviteL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvite interface{}, mods queries.Applicator) error {
	var slice []*Invite
	var object *Invite

	if singular {
		var ok bool
		object, ok = maybeInvite.(*Invite)
		if !ok {
			object = new(Invite)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeInvite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeInvite))
			}
		}
	} else {
		s, ok := maybeInvite.(*[]*Invite)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeInvite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeInvite))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &inviteR{}
		}
		args = append(args, object.GroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &inviteR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`groups`),
		qm.WhereIn(`groups.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups")
	}

	if len(groupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &groupR{}
		}
		foreign.R.Invites = append(foreign.R.Invites, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.Invites = append(foreign.R.Invites, local)
				break
			}
		}
	}

	return nil
}

// SetGroup of the invite to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.Invites.
func (o *Invite) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Group) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"invites\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, invitePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID
	if o.R == nil {
		o.R = &inviteR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &groupR{
			Invites: InviteSlice{o},
		}
	} else {
		related.R.Invites = append(related.R.Invites, o)
	}

	return nil
}

// Invites retrieves all the records using an executor.
func Invites(mods ...qm.QueryMod) inviteQuery {
	mods = append(mods, qm.From("\"invites\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invites\".*"})
	}

	return inviteQuery{q}
}

// FindInvite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvite(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Invite, error) {
	inviteObj := &Invite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invites\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, inviteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from invites")
	}

	if err = inviteObj.doAfterSelectHooks(ctx, exec); err != nil {
		return inviteObj, err
	}

	return inviteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Invite) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invites provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(inviteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	inviteInsertCacheMut.RLock()
	cache, cached := inviteInsertCache[key]
	inviteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			inviteAllColumns,
			inviteColumnsWithDefault,
			inviteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(inviteType, inviteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(inviteType, inviteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invites\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invites\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into invites")
	}

	if !cached {
		inviteInsertCacheMut.Lock()
		inviteInsertCache[key] = cache
		inviteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Invite.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Invite) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	inviteUpdateCacheMut.RLock()
	cache, cached := inviteUpdateCache[key]
	inviteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			inviteAllColumns,
			invitePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update invites, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invites\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, invitePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(inviteType, inviteMapping, append(wl, invitePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update invites row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for invites")
	}

	if !cached {
		inviteUpdateCacheMut.Lock()
		inviteUpdateCache[key] = cache
		inviteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q inviteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for invites")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InviteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invites\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, invitePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invite")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Invite) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invites provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(inviteColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	inviteUpsertCacheMut.RLock()
	cache, cached := inviteUpsertCache[key]
	inviteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			inviteAllColumns,
			inviteColumnsWithDefault,
			inviteColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			inviteAllColumns,
			invitePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert invites, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(invitePrimaryKeyColumns))
			copy(conflict, invitePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invites\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(inviteType, inviteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(inviteType, inviteMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert invites")
	}

	if !cached {
		inviteUpsertCacheMut.Lock()
		inviteUpsertCache[key] = cache
		inviteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Invite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Invite) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Invite provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invitePrimaryKeyMapping)
	sql := "DELETE FROM \"invites\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for invites")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q inviteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no inviteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invites")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InviteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(inviteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"invites\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invitePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invites")
	}

	if len(inviteAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Invite) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvite(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InviteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InviteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invites\".* FROM \"invites\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invitePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InviteSlice")
	}

	*o = slice

	return nil
}

// InviteExists checks if the Invite row exists.
func InviteExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invites\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if invites exists")
	}

	return exists, nil
}

// Exists checks if the Invite row exists.
func (o *Invite) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return InviteExists(ctx, exec, o.ID)
}
