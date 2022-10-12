package db

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

func (r *EventsQuery) FilterByAddresFrom(addressFrom string) *EventsQuery {
	r.sql = r.sql.Where(sq.Eq{"address_from": addressFrom})
	return r
}
func (r *EventsQuery) FilterByAddresTo(addressTo string) *EventsQuery {
	r.sql = r.sql.Where(sq.Eq{"address_to": addressTo})
	return r
}
func (r *EventsQuery) FilterByAmountMore(amountMore string) *EventsQuery {
	r.sql = r.sql.Where("amount>?", amountMore)
	return r
}
func (r *EventsQuery) FilterByAmountLess(amountLess string) *EventsQuery {
	r.sql = r.sql.Where("amount<?", amountLess)
	return r
}
func (r *EventsQuery) FilterByEvents(logName string) *EventsQuery {
	r.sql = r.sql.Where(sq.Eq{"log_name": logName})
	return r
}
func (r *EventsQuery) Select() ([]Events, error) {
	var result []Events

	sql, args, err := r.sql.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to sql")
	}

	err = r.db.Select(&result, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec select stmt")
	}

	return result, nil
}
func (r *EventsQuery) Create(logName string, from string, to string, amount string) error {
	sql, args, err := sq.
		Insert("events").Columns("log_name", "address_from", "address_to", "amount").
		PlaceholderFormat(sq.Dollar).
		Values(logName, from, to, amount).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Insert error")
	}
	fmt.Println(sql, args)
	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "Exec errors")
	}
	return nil
}
func (r *EventsQuery) Page(limit string, cursor string, order string) *EventsQuery {
	l, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		log.Println(errors.Wrap(err, "Parse uint error"))
	}
	c, err := strconv.ParseUint(cursor, 10, 64)
	if err != nil {
		log.Println(errors.Wrap(err, "Parse uint error"))
	}
	r.sql = r.sql.
		Limit(l).
		Offset(c).
		OrderBy(fmt.Sprintf("%s %s", "id", order))
	return r
}
