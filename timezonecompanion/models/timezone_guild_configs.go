// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// TimezoneGuildConfig is an object representing the database table.
type TimezoneGuildConfig struct {
	GuildID             int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	DisabledInChannels  types.Int64Array `boil:"disabled_in_channels" json:"disabled_in_channels,omitempty" toml:"disabled_in_channels" yaml:"disabled_in_channels,omitempty"`
	EnabledInChannels   types.Int64Array `boil:"enabled_in_channels" json:"enabled_in_channels,omitempty" toml:"enabled_in_channels" yaml:"enabled_in_channels,omitempty"`
	NewChannelsDisabled bool             `boil:"new_channels_disabled" json:"new_channels_disabled" toml:"new_channels_disabled" yaml:"new_channels_disabled"`

	R *timezoneGuildConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L timezoneGuildConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TimezoneGuildConfigColumns = struct {
	GuildID             string
	DisabledInChannels  string
	EnabledInChannels   string
	NewChannelsDisabled string
}{
	GuildID:             "guild_id",
	DisabledInChannels:  "disabled_in_channels",
	EnabledInChannels:   "enabled_in_channels",
	NewChannelsDisabled: "new_channels_disabled",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var TimezoneGuildConfigWhere = struct {
	GuildID             whereHelperint64
	DisabledInChannels  whereHelpertypes_Int64Array
	EnabledInChannels   whereHelpertypes_Int64Array
	NewChannelsDisabled whereHelperbool
}{
	GuildID:             whereHelperint64{field: "\"timezone_guild_configs\".\"guild_id\""},
	DisabledInChannels:  whereHelpertypes_Int64Array{field: "\"timezone_guild_configs\".\"disabled_in_channels\""},
	EnabledInChannels:   whereHelpertypes_Int64Array{field: "\"timezone_guild_configs\".\"enabled_in_channels\""},
	NewChannelsDisabled: whereHelperbool{field: "\"timezone_guild_configs\".\"new_channels_disabled\""},
}

// TimezoneGuildConfigRels is where relationship names are stored.
var TimezoneGuildConfigRels = struct {
}{}

// timezoneGuildConfigR is where relationships are stored.
type timezoneGuildConfigR struct {
}

// NewStruct creates a new relationship struct
func (*timezoneGuildConfigR) NewStruct() *timezoneGuildConfigR {
	return &timezoneGuildConfigR{}
}

// timezoneGuildConfigL is where Load methods for each relationship are stored.
type timezoneGuildConfigL struct{}

var (
	timezoneGuildConfigAllColumns            = []string{"guild_id", "disabled_in_channels", "enabled_in_channels", "new_channels_disabled"}
	timezoneGuildConfigColumnsWithoutDefault = []string{"guild_id", "disabled_in_channels", "enabled_in_channels", "new_channels_disabled"}
	timezoneGuildConfigColumnsWithDefault    = []string{}
	timezoneGuildConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// TimezoneGuildConfigSlice is an alias for a slice of pointers to TimezoneGuildConfig.
	// This should generally be used opposed to []TimezoneGuildConfig.
	TimezoneGuildConfigSlice []*TimezoneGuildConfig

	timezoneGuildConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	timezoneGuildConfigType                 = reflect.TypeOf(&TimezoneGuildConfig{})
	timezoneGuildConfigMapping              = queries.MakeStructMapping(timezoneGuildConfigType)
	timezoneGuildConfigPrimaryKeyMapping, _ = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, timezoneGuildConfigPrimaryKeyColumns)
	timezoneGuildConfigInsertCacheMut       sync.RWMutex
	timezoneGuildConfigInsertCache          = make(map[string]insertCache)
	timezoneGuildConfigUpdateCacheMut       sync.RWMutex
	timezoneGuildConfigUpdateCache          = make(map[string]updateCache)
	timezoneGuildConfigUpsertCacheMut       sync.RWMutex
	timezoneGuildConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single timezoneGuildConfig record from the query using the global executor.
func (q timezoneGuildConfigQuery) OneG(ctx context.Context) (*TimezoneGuildConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single timezoneGuildConfig record from the query.
func (q timezoneGuildConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TimezoneGuildConfig, error) {
	o := &TimezoneGuildConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for timezone_guild_configs")
	}

	return o, nil
}

// AllG returns all TimezoneGuildConfig records from the query using the global executor.
func (q timezoneGuildConfigQuery) AllG(ctx context.Context) (TimezoneGuildConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TimezoneGuildConfig records from the query.
func (q timezoneGuildConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (TimezoneGuildConfigSlice, error) {
	var o []*TimezoneGuildConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TimezoneGuildConfig slice")
	}

	return o, nil
}

// CountG returns the count of all TimezoneGuildConfig records in the query, and panics on error.
func (q timezoneGuildConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TimezoneGuildConfig records in the query.
func (q timezoneGuildConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count timezone_guild_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q timezoneGuildConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q timezoneGuildConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if timezone_guild_configs exists")
	}

	return count > 0, nil
}

// TimezoneGuildConfigs retrieves all the records using an executor.
func TimezoneGuildConfigs(mods ...qm.QueryMod) timezoneGuildConfigQuery {
	mods = append(mods, qm.From("\"timezone_guild_configs\""))
	return timezoneGuildConfigQuery{NewQuery(mods...)}
}

// FindTimezoneGuildConfigG retrieves a single record by ID.
func FindTimezoneGuildConfigG(ctx context.Context, guildID int64, selectCols ...string) (*TimezoneGuildConfig, error) {
	return FindTimezoneGuildConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindTimezoneGuildConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTimezoneGuildConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*TimezoneGuildConfig, error) {
	timezoneGuildConfigObj := &TimezoneGuildConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"timezone_guild_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, timezoneGuildConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from timezone_guild_configs")
	}

	return timezoneGuildConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TimezoneGuildConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TimezoneGuildConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no timezone_guild_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(timezoneGuildConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	timezoneGuildConfigInsertCacheMut.RLock()
	cache, cached := timezoneGuildConfigInsertCache[key]
	timezoneGuildConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			timezoneGuildConfigAllColumns,
			timezoneGuildConfigColumnsWithDefault,
			timezoneGuildConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"timezone_guild_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"timezone_guild_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into timezone_guild_configs")
	}

	if !cached {
		timezoneGuildConfigInsertCacheMut.Lock()
		timezoneGuildConfigInsertCache[key] = cache
		timezoneGuildConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single TimezoneGuildConfig record using the global executor.
// See Update for more documentation.
func (o *TimezoneGuildConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TimezoneGuildConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TimezoneGuildConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	timezoneGuildConfigUpdateCacheMut.RLock()
	cache, cached := timezoneGuildConfigUpdateCache[key]
	timezoneGuildConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			timezoneGuildConfigAllColumns,
			timezoneGuildConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update timezone_guild_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"timezone_guild_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, timezoneGuildConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, append(wl, timezoneGuildConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update timezone_guild_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for timezone_guild_configs")
	}

	if !cached {
		timezoneGuildConfigUpdateCacheMut.Lock()
		timezoneGuildConfigUpdateCache[key] = cache
		timezoneGuildConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q timezoneGuildConfigQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q timezoneGuildConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for timezone_guild_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for timezone_guild_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TimezoneGuildConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TimezoneGuildConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timezoneGuildConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"timezone_guild_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, timezoneGuildConfigPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in timezoneGuildConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all timezoneGuildConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TimezoneGuildConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TimezoneGuildConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no timezone_guild_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(timezoneGuildConfigColumnsWithDefault, o)

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

	timezoneGuildConfigUpsertCacheMut.RLock()
	cache, cached := timezoneGuildConfigUpsertCache[key]
	timezoneGuildConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			timezoneGuildConfigAllColumns,
			timezoneGuildConfigColumnsWithDefault,
			timezoneGuildConfigColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			timezoneGuildConfigAllColumns,
			timezoneGuildConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert timezone_guild_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(timezoneGuildConfigPrimaryKeyColumns))
			copy(conflict, timezoneGuildConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"timezone_guild_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(timezoneGuildConfigType, timezoneGuildConfigMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert timezone_guild_configs")
	}

	if !cached {
		timezoneGuildConfigUpsertCacheMut.Lock()
		timezoneGuildConfigUpsertCache[key] = cache
		timezoneGuildConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single TimezoneGuildConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TimezoneGuildConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TimezoneGuildConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TimezoneGuildConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TimezoneGuildConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), timezoneGuildConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"timezone_guild_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from timezone_guild_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for timezone_guild_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q timezoneGuildConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no timezoneGuildConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from timezone_guild_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for timezone_guild_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TimezoneGuildConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TimezoneGuildConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timezoneGuildConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"timezone_guild_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, timezoneGuildConfigPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from timezoneGuildConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for timezone_guild_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TimezoneGuildConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TimezoneGuildConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TimezoneGuildConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTimezoneGuildConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TimezoneGuildConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TimezoneGuildConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TimezoneGuildConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TimezoneGuildConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timezoneGuildConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"timezone_guild_configs\".* FROM \"timezone_guild_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, timezoneGuildConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TimezoneGuildConfigSlice")
	}

	*o = slice

	return nil
}

// TimezoneGuildConfigExistsG checks if the TimezoneGuildConfig row exists.
func TimezoneGuildConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return TimezoneGuildConfigExists(ctx, boil.GetContextDB(), guildID)
}

// TimezoneGuildConfigExists checks if the TimezoneGuildConfig row exists.
func TimezoneGuildConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"timezone_guild_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if timezone_guild_configs exists")
	}

	return exists, nil
}
