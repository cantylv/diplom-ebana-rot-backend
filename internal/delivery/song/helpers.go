package song

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cantylv/online-music-lib/internal/entity/dto"
	e "github.com/cantylv/online-music-lib/internal/helpers/my/error"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

func validateFilterLibraryOpts(r *http.Request) (*dto.FilterLibraryOptions, []error) {
	var (
		hints []error
		err   error
	)

	queryParams := r.URL.Query()

	// нужно проверить список идентификаторов и имен песен
	var ids []string
	if queryParams.Get("ids") != "" {
		ids = strings.Split(queryParams.Get("ids"), "@")
		for _, songID := range ids {
			_, err = uuid.FromString(songID)
			if err != nil {
				hints = append(hints, e.ErrInvalidSongID)
			}
		}
	}

	var names []string
	if queryParams.Get("names") != "" {
		names = strings.Split(queryParams.Get("names"), "@")
	}
	// нужно проверить на соответствие формату времени
	layout := "02-01-2006"
	var fromReleaseDate dto.Date
	fromReleaseDateParam := queryParams.Get("from_release_date")
	if fromReleaseDateParam != "" {
		time, err := time.Parse(layout, fromReleaseDateParam)
		if err != nil {
			hints = append(hints, e.ErrInvalidFromReleaseDate)
		} else {
			fromReleaseDate.Valid = true
			fromReleaseDate.Time = time
		}
	} else {
		fromReleaseDate.Valid = false
	}

	var toReleaseDate dto.Date
	toReleaseDateParam := queryParams.Get("to_release_date")
	if toReleaseDateParam != "" {
		time, err := time.Parse(layout, toReleaseDateParam)
		if err != nil {
			hints = append(hints, e.ErrInvalidToReleaseDate)
		} else {
			toReleaseDate.Valid = true
			toReleaseDate.Time = time
		}
	} else {
		toReleaseDate.Valid = false
		toReleaseDate.Time = time.Now()
	}

	// нужно провалидировать числовые параметры
	var songLimitNum int
	songLimit := queryParams.Get("limit")
	if songLimit != "" {
		songLimitNum, err = strconv.Atoi(songLimit)
		if err != nil || songLimitNum < 0 {
			hints = append(hints, e.ErrInvalidSongLimit)
		}
	} else {
		songLimitNum = 2 // default value
	}

	var songOffsetNum int
	songOffset := queryParams.Get("offset")
	if songOffset != "" {
		songOffsetNum, err = strconv.Atoi(songOffset)
		if err != nil || songOffsetNum < 0 {
			hints = append(hints, e.ErrInvalidSongOffset)
		}
	} else {
		songOffsetNum = 0 // default value
	}

	if len(hints) != 0 {
		return nil, hints
	}

	return &dto.FilterLibraryOptions{
		SongIDs:         ids,   // format: id1,id2,id3
		SongNames:       names, // format: Shape of You, Blinding Lights
		FromReleaseDate: fromReleaseDate,
		ToReleaseDate:   toReleaseDate,
		TextPhrases:     queryParams.Get("text"),
		SongLimit:       songLimitNum,
		SongOffset:      songOffsetNum,
	}, nil
}

func validateFilterSongOpts(r *http.Request) (*dto.FilterSongOptions, []error) {
	var (
		hints []error
		err   error
	)

	queryParams := r.URL.Query()

	id, err := uuid.FromString(mux.Vars(r)["song_id"])
	if err != nil {
		hints = append(hints, e.ErrBadUUID)
	}

	// нужно провалидировать числовые параметры
	var coupletsLimitNum int
	songLimit := queryParams.Get("limit")
	if songLimit != "" {
		coupletsLimitNum, err = strconv.Atoi(songLimit)
		if err != nil || coupletsLimitNum < 0 {
			hints = append(hints, e.ErrInvalidSongLimit)
		}
	} else {
		coupletsLimitNum = 2 // default value
	}

	var coupletsOffsetNum int
	songOffset := queryParams.Get("offset")
	if songOffset != "" {
		coupletsOffsetNum, err = strconv.Atoi(songOffset)
		if err != nil || coupletsOffsetNum < 0 {
			hints = append(hints, e.ErrInvalidSongOffset)
		}
		coupletsOffsetNum++
	} else {
		coupletsOffsetNum = 1 // default value
	}

	if len(hints) != 0 {
		return nil, hints
	}

	return &dto.FilterSongOptions{
		SongID:        id.String(),
		CoupletLimit:  coupletsLimitNum,
		CoupletOffset: coupletsOffsetNum,
	}, nil
}
