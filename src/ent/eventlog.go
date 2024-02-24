// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/ent/eventlog"
)

// EventLog is the model entity for the EventLog schema.
type EventLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Block holds the value of the "block" field.
	Block uint64 `json:"block,omitempty"`
	// SignersCount holds the value of the "signersCount" field.
	SignersCount uint64 `json:"signersCount,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature []byte `json:"signature,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Chain holds the value of the "chain" field.
	Chain string `json:"chain,omitempty"`
	// Index holds the value of the "index" field.
	Index uint64 `json:"index,omitempty"`
	// Event holds the value of the "event" field.
	Event string `json:"event,omitempty"`
	// Transaction holds the value of the "transaction" field.
	Transaction []byte `json:"transaction,omitempty"`
	// Args holds the value of the "args" field.
	Args []datasets.EventLogArg `json:"args,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventLogQuery when eager-loading is set.
	Edges        EventLogEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EventLogEdges holds the relations/edges for other nodes in the graph.
type EventLogEdges struct {
	// Signers holds the value of the signers edge.
	Signers []*Signer `json:"signers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SignersOrErr returns the Signers value or an error if the edge
// was not loaded in eager-loading.
func (e EventLogEdges) SignersOrErr() ([]*Signer, error) {
	if e.loadedTypes[0] {
		return e.Signers, nil
	}
	return nil, &NotLoadedError{edge: "signers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventlog.FieldSignature, eventlog.FieldTransaction, eventlog.FieldArgs:
			values[i] = new([]byte)
		case eventlog.FieldID, eventlog.FieldBlock, eventlog.FieldSignersCount, eventlog.FieldIndex:
			values[i] = new(sql.NullInt64)
		case eventlog.FieldAddress, eventlog.FieldChain, eventlog.FieldEvent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventLog fields.
func (el *EventLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			el.ID = int(value.Int64)
		case eventlog.FieldBlock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block", values[i])
			} else if value.Valid {
				el.Block = uint64(value.Int64)
			}
		case eventlog.FieldSignersCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field signersCount", values[i])
			} else if value.Valid {
				el.SignersCount = uint64(value.Int64)
			}
		case eventlog.FieldSignature:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value != nil {
				el.Signature = *value
			}
		case eventlog.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				el.Address = value.String
			}
		case eventlog.FieldChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain", values[i])
			} else if value.Valid {
				el.Chain = value.String
			}
		case eventlog.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				el.Index = uint64(value.Int64)
			}
		case eventlog.FieldEvent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event", values[i])
			} else if value.Valid {
				el.Event = value.String
			}
		case eventlog.FieldTransaction:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field transaction", values[i])
			} else if value != nil {
				el.Transaction = *value
			}
		case eventlog.FieldArgs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field args", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &el.Args); err != nil {
					return fmt.Errorf("unmarshal field args: %w", err)
				}
			}
		default:
			el.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EventLog.
// This includes values selected through modifiers, order, etc.
func (el *EventLog) Value(name string) (ent.Value, error) {
	return el.selectValues.Get(name)
}

// QuerySigners queries the "signers" edge of the EventLog entity.
func (el *EventLog) QuerySigners() *SignerQuery {
	return NewEventLogClient(el.config).QuerySigners(el)
}

// Update returns a builder for updating this EventLog.
// Note that you need to call EventLog.Unwrap() before calling this method if this EventLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *EventLog) Update() *EventLogUpdateOne {
	return NewEventLogClient(el.config).UpdateOne(el)
}

// Unwrap unwraps the EventLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (el *EventLog) Unwrap() *EventLog {
	_tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: EventLog is not a transactional entity")
	}
	el.config.driver = _tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *EventLog) String() string {
	var builder strings.Builder
	builder.WriteString("EventLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", el.ID))
	builder.WriteString("block=")
	builder.WriteString(fmt.Sprintf("%v", el.Block))
	builder.WriteString(", ")
	builder.WriteString("signersCount=")
	builder.WriteString(fmt.Sprintf("%v", el.SignersCount))
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(fmt.Sprintf("%v", el.Signature))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(el.Address)
	builder.WriteString(", ")
	builder.WriteString("chain=")
	builder.WriteString(el.Chain)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", el.Index))
	builder.WriteString(", ")
	builder.WriteString("event=")
	builder.WriteString(el.Event)
	builder.WriteString(", ")
	builder.WriteString("transaction=")
	builder.WriteString(fmt.Sprintf("%v", el.Transaction))
	builder.WriteString(", ")
	builder.WriteString("args=")
	builder.WriteString(fmt.Sprintf("%v", el.Args))
	builder.WriteByte(')')
	return builder.String()
}

// EventLogs is a parsable slice of EventLog.
type EventLogs []*EventLog
