package song

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cantylv/online-music-lib/internal/entity"
	"github.com/cantylv/online-music-lib/internal/entity/dto"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type DBContract interface {
	GetAll(ctx context.Context, opts *dto.FilterLibraryOptions) ([]*entity.Song, error)
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

func (t *dbconnector) GetAll(ctx context.Context, opts *dto.FilterLibraryOptions) ([]*entity.Song, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select("*").From("song")
	if len(opts.SongIDs) != 0 {
		// преобразуем слайс строк в слайс interface{}
		values := make([]interface{}, len(opts.SongIDs))
		for i, v := range opts.SongIDs {
			values[i] = v
		}
		sb.Where(sb.In("id", values...))
	}
	if len(opts.SongNames) != 0 {
		values := make([]interface{}, len(opts.SongNames))
		for i, v := range opts.SongNames {
			values[i] = v
		}
		sb.Where(sb.In("name", values...))
	}
	if opts.FromReleaseDate.Valid {
		sb.Where(sb.Between("release_date", opts.FromReleaseDate.Time, opts.ToReleaseDate.Time))
	}
	if opts.TextPhrases != "" {
		opts.TextPhrases = escapeSQLString(opts.TextPhrases)
		sb.Where(fmt.Sprintf(
			"EXISTS (SELECT 1 FROM jsonb_array_elements(song.text->'couplets') AS couplet WHERE couplet::text ILIKE '%%%s%%')",
			sqlbuilder.Escape(opts.TextPhrases),
		))

	}
	// здесь именно пагинация песен
	sb.Limit(opts.SongLimit).Offset(opts.SongOffset)

	sqlQuery, args := sb.Build()
	var rows pgx.Rows
	var err error
	if len(args) != 0 {
		rows, err = t.pgConn.Query(ctx, sqlQuery, args...)
	} else {
		rows, err = t.pgConn.Query(ctx, sqlQuery)
	}
	defer rows.Close() // nolint: errcheck
	if err != nil {
		return nil, err
	}

	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song
		err := rows.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
		if err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}

	return songs, nil
}

func escapeSQLString(input string) string {
	return strings.ReplaceAll(input, "'", "''")
}

func (t *dbconnector) GetByID(ctx context.Context, ID string) (*entity.Song, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select("*").From("song")
	sb.Where(sb.EQ("id", ID))

	sqlQuery, args := sb.Build()
	row := t.pgConn.QueryRow(ctx, sqlQuery, args...)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (t *dbconnector) DeleteByID(ctx context.Context, ID string) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder().DeleteFrom("song")
	sb.Where(sb.EQ("id", ID))

	sqlQuery, args := sb.Build()
	tag, err := t.pgConn.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (t *dbconnector) UpdateByID(ctx context.Context, data *dto.UpdateSong) (*entity.Song, error) {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update("song")

	updates := []string{}
	if data.Name != "" {
		updates = append(updates, sb.Assign("name", data.Name))
	}
	if data.ReleaseDate != "" {
		updates = append(updates, sb.Assign("release_date", data.ReleaseDate))
	}
	if data.NewText.Couplets != nil {
		updates = append(updates, sb.Assign("text", data.NewText))
	}
	if data.Link != "" {
		updates = append(updates, sb.Assign("link", data.Link))
	}

	if len(updates) > 0 {
		sb.Set(updates...)
	}

	sb.Where(sb.Equal("id", data.ID))

	sqlQuery, args := sb.Build()
	sqlQuery += " RETURNING *"

	// Выполняем запрос
	row := t.pgConn.QueryRow(ctx, sqlQuery, args...)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (t *dbconnector) Create(ctx context.Context, data *dto.CreateData) (*entity.Song, error) {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto("song").
		Cols("name", "release_date", "text", "link").
		Values(data.Name, data.ReleaseDate, data.Text, data.Link)

	sqlQuery, args := sb.Build()
	sqlQuery += " RETURNING *"

	row := t.pgConn.QueryRow(ctx, sqlQuery, args...)
	var song entity.Song
	err := row.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
