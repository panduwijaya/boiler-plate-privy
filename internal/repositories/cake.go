// Package repositories
// Automatic generated
package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"

	"golang.org/x/sync/errgroup"

	"cake-store/cake-store/internal/common"
	"cake-store/cake-store/internal/entity"
	"cake-store/cake-store/pkg/builderx"
	"cake-store/cake-store/pkg/databasex"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/tracer"
)

// Cakeer contract of Cake
type Cakeer interface {
	Storer
	Updater
	Deleter
	Counter
	FindOne(ctx context.Context, param interface{}) (*entity.Cake, error)
	Find(ctx context.Context, param interface{}) ([]entity.Cake, error)
	FindWithCount(ctx context.Context, param interface{}) ([]entity.Cake, int64, error)
	FileUpload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader)
}

type cake struct {
	db databasex.Adapter
}

// NewCake create new instance of Cake
func NewCake(db databasex.Adapter) Cakeer {
	return &cake{db: db}
}

func (r *cake) FileUpload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) {
	fmt.Println(file)
}

// FindOne cake
func (r *cake) FindOne(ctx context.Context, param interface{}) (*entity.Cake, error) {
	var (
		result entity.Cake
		err    error
	)

	ctx = tracer.SpanStart(ctx, "repo.cake_find_one")
	defer tracer.SpanFinish(ctx)
	logger.Info(param)
	wq, err := builderx.StructToMySqlQueryWhere(param, "db")
	if err != nil {
		tracer.SpanError(ctx, err)
		return nil, err
	}

	q := `SELECT 
			id,
			title,
			description,
			rating,
			image,
			created_at,
			updated_at
		 FROM cakes %s LIMIT 1`

	err = r.db.QueryRow(ctx, &result, fmt.Sprintf(q, wq.Query), wq.Values...)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

// Find cake
func (r *cake) Find(ctx context.Context, param interface{}) ([]entity.Cake, error) {
	var (
		result []entity.Cake
		err    error
	)

	ctx = tracer.SpanStart(ctx, "repo.cake_finds")
	defer tracer.SpanFinish(ctx)

	wq, err := builderx.StructToMySqlQueryWhere(param, "db")
	if err != nil {
		tracer.SpanError(ctx, err)
		return nil, err
	}

	q := `SELECT 
			id,
			title,
			description,
			rating,
			image,
			created_at,
			updated_at
		 FROM cakes %s LIMIT ? OFFSET ? `

	vals := wq.Values
	vals = append(vals, wq.Limit, common.PageToOffset(wq.Limit, wq.Page))
	err = r.db.Query(ctx, &result, fmt.Sprintf(q, wq.Query), vals...)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, err
}

// Store cake
func (r *cake) Store(ctx context.Context, param interface{}) (int64, error) {
	var (
		err      error
		affected int64
	)

	ctx = tracer.SpanStart(ctx, "repo.cake_store")
	defer tracer.SpanFinish(ctx)

	np := &param
	param = *np
	query, vals, err := builderx.StructToQueryInsert(param, "cakes", "db")
	if err != nil {
		tracer.SpanError(ctx, err)
		return 0, err
	}

	// See https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels.
	err = r.db.Transact(ctx, sql.LevelRepeatableRead, func(tx *databasex.DB) error {
		af, err := tx.Exec(ctx, query, vals...)
		affected = af
		return err
	})

	return affected, err

}

// Update cake data
func (r *cake) Update(ctx context.Context, input interface{}, where interface{}) (int64, error) {
	var (
		err      error
		affected int64
	)

	ctx = tracer.SpanStart(ctx, "repo.cake_update")
	defer tracer.SpanFinish(ctx)

	query, vals, err := builderx.StructToQueryUpdate(input, where, "cakes", "db")
	if err != nil {
		tracer.SpanError(ctx, err)
		return 0, err
	}

	// See https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels.
	err = r.db.Transact(ctx, sql.LevelRepeatableRead, func(tx *databasex.DB) error {
		af, err := tx.Exec(ctx, query, vals...)
		affected = af
		return err
	})

	return affected, err
}

// Delete cake from database
func (r *cake) Delete(ctx context.Context, param interface{}) (int64, error) {
	var (
		err      error
		affected int64
	)
	ctx = tracer.SpanStart(ctx, "repo.cake_delete")
	defer tracer.SpanFinish(ctx)

	query, vals, err := builderx.StructToQueryDelete(param, "cakes", "db", false)
	if err != nil {
		tracer.SpanError(ctx, err)
		return 0, err
	}

	// See https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels.
	err = r.db.Transact(ctx, sql.LevelRepeatableRead, func(tx *databasex.DB) error {
		af, err := tx.Exec(ctx, query, vals...)
		affected = af
		return err
	})

	return affected, err
}

// Count cake
func (r *cake) Count(ctx context.Context, p interface{}) (total int64, err error) {
	ctx = tracer.SpanStart(ctx, "repo.cake_count")
	defer tracer.SpanFinish(ctx)

	wq, err := builderx.StructToMySqlQueryWhere(p, "db")
	if err != nil {
		tracer.SpanError(ctx, err)
		return
	}

	q := fmt.Sprintf(`
		SELECT
        	COUNT(id) AS jumlah
		FROM cakes %s `, wq.Query)

	err = r.db.QueryRow(ctx, &total, q, wq.Values...)
	if err != nil {
		tracer.SpanError(ctx, err)
		err = err
		return
	}

	return
}

// FindWithCount find cake with count
func (r *cake) FindWithCount(ctx context.Context, param interface{}) ([]entity.Cake, int64, error) {

	var (
		cl    []entity.Cake
		count int64
	)

	ctx = tracer.SpanStart(ctx, "repo.cake_with_count")
	defer tracer.SpanFinish(ctx)

	group, newCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		l, err := r.Find(newCtx, param)
		cl = l
		return err
	})
	group.Go(func() error {
		c, err := r.Count(ctx, param)
		count = c
		return err
	})

	err := group.Wait()

	return cl, count, err
}
