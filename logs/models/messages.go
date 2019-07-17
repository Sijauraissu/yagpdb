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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Message is an object representing the database table.
type Message struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt      null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt      null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	MessageLogID   null.Int    `boil:"message_log_id" json:"message_log_id,omitempty" toml:"message_log_id" yaml:"message_log_id,omitempty"`
	MessageID      null.String `boil:"message_id" json:"message_id,omitempty" toml:"message_id" yaml:"message_id,omitempty"`
	AuthorUsername null.String `boil:"author_username" json:"author_username,omitempty" toml:"author_username" yaml:"author_username,omitempty"`
	AuthorDiscrim  null.String `boil:"author_discrim" json:"author_discrim,omitempty" toml:"author_discrim" yaml:"author_discrim,omitempty"`
	AuthorID       null.String `boil:"author_id" json:"author_id,omitempty" toml:"author_id" yaml:"author_id,omitempty"`
	Deleted        null.Bool   `boil:"deleted" json:"deleted,omitempty" toml:"deleted" yaml:"deleted,omitempty"`
	Content        null.String `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	Timestamp      null.String `boil:"timestamp" json:"timestamp,omitempty" toml:"timestamp" yaml:"timestamp,omitempty"`

	R *messageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageColumns = struct {
	ID             string
	CreatedAt      string
	UpdatedAt      string
	MessageLogID   string
	MessageID      string
	AuthorUsername string
	AuthorDiscrim  string
	AuthorID       string
	Deleted        string
	Content        string
	Timestamp      string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	MessageLogID:   "message_log_id",
	MessageID:      "message_id",
	AuthorUsername: "author_username",
	AuthorDiscrim:  "author_discrim",
	AuthorID:       "author_id",
	Deleted:        "deleted",
	Content:        "content",
	Timestamp:      "timestamp",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MessageWhere = struct {
	ID             whereHelperint
	CreatedAt      whereHelpernull_Time
	UpdatedAt      whereHelpernull_Time
	MessageLogID   whereHelpernull_Int
	MessageID      whereHelpernull_String
	AuthorUsername whereHelpernull_String
	AuthorDiscrim  whereHelpernull_String
	AuthorID       whereHelpernull_String
	Deleted        whereHelpernull_Bool
	Content        whereHelpernull_String
	Timestamp      whereHelpernull_String
}{
	ID:             whereHelperint{field: "\"messages\".\"id\""},
	CreatedAt:      whereHelpernull_Time{field: "\"messages\".\"created_at\""},
	UpdatedAt:      whereHelpernull_Time{field: "\"messages\".\"updated_at\""},
	MessageLogID:   whereHelpernull_Int{field: "\"messages\".\"message_log_id\""},
	MessageID:      whereHelpernull_String{field: "\"messages\".\"message_id\""},
	AuthorUsername: whereHelpernull_String{field: "\"messages\".\"author_username\""},
	AuthorDiscrim:  whereHelpernull_String{field: "\"messages\".\"author_discrim\""},
	AuthorID:       whereHelpernull_String{field: "\"messages\".\"author_id\""},
	Deleted:        whereHelpernull_Bool{field: "\"messages\".\"deleted\""},
	Content:        whereHelpernull_String{field: "\"messages\".\"content\""},
	Timestamp:      whereHelpernull_String{field: "\"messages\".\"timestamp\""},
}

// MessageRels is where relationship names are stored.
var MessageRels = struct {
	MessageLog string
}{
	MessageLog: "MessageLog",
}

// messageR is where relationships are stored.
type messageR struct {
	MessageLog *MessageLog
}

// NewStruct creates a new relationship struct
func (*messageR) NewStruct() *messageR {
	return &messageR{}
}

// messageL is where Load methods for each relationship are stored.
type messageL struct{}

var (
	messageAllColumns            = []string{"id", "created_at", "updated_at", "message_log_id", "message_id", "author_username", "author_discrim", "author_id", "deleted", "content", "timestamp"}
	messageColumnsWithoutDefault = []string{"created_at", "updated_at", "message_log_id", "message_id", "author_username", "author_discrim", "author_id", "deleted", "content", "timestamp"}
	messageColumnsWithDefault    = []string{"id"}
	messagePrimaryKeyColumns     = []string{"id"}
)

type (
	// MessageSlice is an alias for a slice of pointers to Message.
	// This should generally be used opposed to []Message.
	MessageSlice []*Message

	messageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageType                 = reflect.TypeOf(&Message{})
	messageMapping              = queries.MakeStructMapping(messageType)
	messagePrimaryKeyMapping, _ = queries.BindMapping(messageType, messageMapping, messagePrimaryKeyColumns)
	messageInsertCacheMut       sync.RWMutex
	messageInsertCache          = make(map[string]insertCache)
	messageUpdateCacheMut       sync.RWMutex
	messageUpdateCache          = make(map[string]updateCache)
	messageUpsertCacheMut       sync.RWMutex
	messageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single message record from the query using the global executor.
func (q messageQuery) OneG(ctx context.Context) (*Message, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single message record from the query.
func (q messageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Message, error) {
	o := &Message{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for messages")
	}

	return o, nil
}

// AllG returns all Message records from the query using the global executor.
func (q messageQuery) AllG(ctx context.Context) (MessageSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Message records from the query.
func (q messageQuery) All(ctx context.Context, exec boil.ContextExecutor) (MessageSlice, error) {
	var o []*Message

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Message slice")
	}

	return o, nil
}

// CountG returns the count of all Message records in the query, and panics on error.
func (q messageQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Message records in the query.
func (q messageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count messages rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q messageQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q messageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if messages exists")
	}

	return count > 0, nil
}

// MessageLog pointed to by the foreign key.
func (o *Message) MessageLog(mods ...qm.QueryMod) messageLogQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MessageLogID),
	}

	queryMods = append(queryMods, mods...)

	query := MessageLogs(queryMods...)
	queries.SetFrom(query.Query, "\"message_logs\"")

	return query
}

// LoadMessageLog allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageL) LoadMessageLog(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessage interface{}, mods queries.Applicator) error {
	var slice []*Message
	var object *Message

	if singular {
		object = maybeMessage.(*Message)
	} else {
		slice = *maybeMessage.(*[]*Message)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageR{}
		}
		if !queries.IsNil(object.MessageLogID) {
			args = append(args, object.MessageLogID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.MessageLogID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.MessageLogID) {
				args = append(args, obj.MessageLogID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`message_logs`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load MessageLog")
	}

	var resultSlice []*MessageLog
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice MessageLog")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for message_logs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for message_logs")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.MessageLog = foreign
		if foreign.R == nil {
			foreign.R = &messageLogR{}
		}
		foreign.R.Messages = append(foreign.R.Messages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.MessageLogID, foreign.ID) {
				local.R.MessageLog = foreign
				if foreign.R == nil {
					foreign.R = &messageLogR{}
				}
				foreign.R.Messages = append(foreign.R.Messages, local)
				break
			}
		}
	}

	return nil
}

// SetMessageLogG of the message to the related item.
// Sets o.R.MessageLog to related.
// Adds o to related.R.Messages.
// Uses the global database handle.
func (o *Message) SetMessageLogG(ctx context.Context, insert bool, related *MessageLog) error {
	return o.SetMessageLog(ctx, boil.GetContextDB(), insert, related)
}

// SetMessageLog of the message to the related item.
// Sets o.R.MessageLog to related.
// Adds o to related.R.Messages.
func (o *Message) SetMessageLog(ctx context.Context, exec boil.ContextExecutor, insert bool, related *MessageLog) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_log_id"}),
		strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.MessageLogID, related.ID)
	if o.R == nil {
		o.R = &messageR{
			MessageLog: related,
		}
	} else {
		o.R.MessageLog = related
	}

	if related.R == nil {
		related.R = &messageLogR{
			Messages: MessageSlice{o},
		}
	} else {
		related.R.Messages = append(related.R.Messages, o)
	}

	return nil
}

// RemoveMessageLogG relationship.
// Sets o.R.MessageLog to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *Message) RemoveMessageLogG(ctx context.Context, related *MessageLog) error {
	return o.RemoveMessageLog(ctx, boil.GetContextDB(), related)
}

// RemoveMessageLog relationship.
// Sets o.R.MessageLog to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Message) RemoveMessageLog(ctx context.Context, exec boil.ContextExecutor, related *MessageLog) error {
	var err error

	queries.SetScanner(&o.MessageLogID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("message_log_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.MessageLog = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Messages {
		if queries.Equal(o.MessageLogID, ri.MessageLogID) {
			continue
		}

		ln := len(related.R.Messages)
		if ln > 1 && i < ln-1 {
			related.R.Messages[i] = related.R.Messages[ln-1]
		}
		related.R.Messages = related.R.Messages[:ln-1]
		break
	}
	return nil
}

// Messages retrieves all the records using an executor.
func Messages(mods ...qm.QueryMod) messageQuery {
	mods = append(mods, qm.From("\"messages\""))
	return messageQuery{NewQuery(mods...)}
}

// FindMessageG retrieves a single record by ID.
func FindMessageG(ctx context.Context, iD int, selectCols ...string) (*Message, error) {
	return FindMessage(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Message, error) {
	messageObj := &Message{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"messages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, messageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from messages")
	}

	return messageObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Message) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Message) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageInsertCacheMut.RLock()
	cache, cached := messageInsertCache[key]
	messageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageType, messageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"messages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"messages\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into messages")
	}

	if !cached {
		messageInsertCacheMut.Lock()
		messageInsertCache[key] = cache
		messageInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Message record using the global executor.
// See Update for more documentation.
func (o *Message) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Message.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Message) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	messageUpdateCacheMut.RLock()
	cache, cached := messageUpdateCache[key]
	messageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, append(wl, messagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update messages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for messages")
	}

	if !cached {
		messageUpdateCacheMut.Lock()
		messageUpdateCache[key] = cache
		messageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q messageQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q messageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for messages")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o MessageSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all message")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Message) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Message) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

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

	messageUpsertCacheMut.RLock()
	cache, cached := messageUpsertCache[key]
	messageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messagePrimaryKeyColumns))
			copy(conflict, messagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"messages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messageType, messageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert messages")
	}

	if !cached {
		messageUpsertCacheMut.Lock()
		messageUpsertCache[key] = cache
		messageUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Message record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Message) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Message record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Message) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Message provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messagePrimaryKeyMapping)
	sql := "DELETE FROM \"messages\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o MessageSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Message) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Message provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Message) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty MessageSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"messages\".* FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MessageSlice")
	}

	*o = slice

	return nil
}

// MessageExistsG checks if the Message row exists.
func MessageExistsG(ctx context.Context, iD int) (bool, error) {
	return MessageExists(ctx, boil.GetContextDB(), iD)
}

// MessageExists checks if the Message row exists.
func MessageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"messages\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if messages exists")
	}

	return exists, nil
}
