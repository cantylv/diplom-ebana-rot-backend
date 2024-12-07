package song

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cantylv/online-music-lib/internal/entity"
	"github.com/cantylv/online-music-lib/internal/entity/dto"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type DBContract interface {
	GetAll(ctx context.Context, opts *dto.FilterLibraryOptions) ([]entity.Song, error)
	GetByID(ctx context.Context, ID string) (*entity.Song, error)
	DeleteByID(ctx context.Context, ID string) error
	UpdateByID(ctx context.Context, data *dto.UpdateSong) (*entity.Song, error)
	Create(ctx context.Context, data *dto.CreateData) (*entity.Song, error)
}

var _ DBContract = (*dbconnector)(nil)

type dbconnector struct {
	pgConn *pgx.Conn
}

func NewDatabaseConnector(pgConn *pgx.Conn) *dbconnector {
	return &dbconnector{
		pgConn: pgConn,
	}
}

func (t *dbconnector) GetAll(ctx context.Context, opts *dto.FilterLibraryOptions) ([]entity.Song, error) {
	sb := sqlbuilder.Select("*").From("song")
	if opts.SongIDs != "" {
		sb.Where(sb.In("id", opts.SongIDs))
	}
	if opts.SongNames != "" {
		sb.Where(sb.In("name", opts.SongIDs))
	}
	if opts.FromReleaseDate != "" {
		sb.Where(sb.Between("release_date", opts.FromReleaseDate, opts.ToReleaseDate))
	}
	if opts.TextPhrases != "" {
		sb.Where(sb.Like("text", opts.TextPhrases))
	}
	// здесь именно пагинация песен
	sb.Limit(opts.SongLimit).Offset(opts.SongOffset)

	sqlQuery, args := sb.Build()
	rows, err := t.pgConn.Query(ctx, sqlQuery, args)
	defer rows.Close()// nolint: errcheck
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	var songs []entity.Song
	for rows.Next() {
		var song entity.Song
		err := rows.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (t *dbconnector) GetByID(ctx context.Context, ID string) (*entity.Song, error) {
	sb := sqlbuilder.Select("*").From("song")
	sb.Where(sb.EQ("id", ID))

	sqlQuery, args := sb.Build()
	row := t.pgConn.QueryRow(ctx, sqlQuery, args)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &song, nil
}

func (t *dbconnector) DeleteByID(ctx context.Context, ID string) error {
	sb := sqlbuilder.DeleteFrom("song")
	sb.Where(sb.EQ("id", ID))

	sqlQuery, args := sb.Build()
	tag, err := t.pgConn.Exec(ctx, sqlQuery, args)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (t *dbconnector) UpdateByID(ctx context.Context, data *dto.UpdateSong) (*entity.Song, error) {
	sb := sqlbuilder.Update("song")
	if data.Name != "" {
		sb.Set(sb.EQ("name", data.Name))
	}
	if data.ReleaseDate.GoString() != "" {
		sb.Set(sb.EQ("release_date", data.ReleaseDate))
	}
	if data.NewText.Couplets != nil {
		sb.Set(sb.EQ("text", data.NewText))
	}
	if data.Link != "" {
		sb.Set(sb.EQ("link", data.Link))
	}
	sb.Where(sb.EQ("id", data.ID))

	sqlQuery, args := sb.Build()
	sqlQuery += " RETURNING *"

	row := t.pgConn.QueryRow(ctx, sqlQuery, args)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (t *dbconnector) Create(ctx context.Context, data *dto.CreateData) (*entity.Song, error) {
	sb := sqlbuilder.InsertInto("song").
		Cols("name", "release_date", "text", "link").
		Values(data.Name, data.ReleaseDate, data.Text, data.Link)

	sqlQuery, args := sb.Build()
	sqlQuery += " RETURNING *"

	row := t.pgConn.QueryRow(ctx, sqlQuery, args)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
