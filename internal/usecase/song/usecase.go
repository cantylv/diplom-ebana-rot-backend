package song

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cantylv/online-music-lib/internal/entity"
	"github.com/cantylv/online-music-lib/internal/entity/dto"
	e "github.com/cantylv/online-music-lib/internal/helpers/my/error"
	"github.com/cantylv/online-music-lib/internal/repo/song"
)

type Contract interface {
	GetLibrarySongs(ctx context.Context, opts *dto.FilterLibraryOptions) ([]*entity.Song, error)
	GetLibrarySong(ctx context.Context, opts *dto.FilterSongOptions) (*entity.Song, error)
	DeleteLibrarySong(ctx context.Context, songID string) error
	UpdateLibrarySong(ctx context.Context, data *dto.UpdateSong) (*entity.Song, error)
	AddNewSongToLibrary(ctx context.Context, data *dto.CreateData) (*entity.Song, error)
}

var _ Contract = (*proccessor)(nil)

type proccessor struct {
	repoSong song.DBContract
}

func Newproccessor(repoSong song.DBContract) *proccessor {
	return &proccessor{
		repoSong: repoSong,
	}
}

func (t *proccessor) GetLibrarySongs(ctx context.Context, opts *dto.FilterLibraryOptions) ([]*entity.Song, error) {
	songs, err := t.repoSong.GetAll(ctx, opts)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return songs, nil
}

func (t *proccessor) GetLibrarySong(ctx context.Context, opts *dto.FilterSongOptions) (*entity.Song, error) {
	song, err := t.repoSong.GetByID(ctx, opts.SongID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	countCouplet := len(song.Text.Couplets)
	song.Text.Couplets = song.Text.Couplets[min(opts.CoupletOffset-1, countCouplet-1):min(opts.CoupletOffset+opts.CoupletLimit-1, countCouplet)]
	return song, nil
}

func (t *proccessor) DeleteLibrarySong(ctx context.Context, songID string) error {
	err := t.repoSong.DeleteByID(ctx, songID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.ErrNotExist
		}
		return err
	}
	return nil
}

func (t *proccessor) UpdateLibrarySong(ctx context.Context, data *dto.UpdateSong) (*entity.Song, error) {
	updatedSong, err := t.repoSong.UpdateByID(ctx, data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.ErrNotExist
		}
		return nil, err
	}
	return updatedSong, nil
}

func (t *proccessor) AddNewSongToLibrary(ctx context.Context, data *dto.CreateData) (*entity.Song, error) {
	return t.repoSong.Create(ctx, data)
}
