package song

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/cantylv/online-music-lib/internal/entity/dto"
	f "github.com/cantylv/online-music-lib/internal/helpers/function"
	mc "github.com/cantylv/online-music-lib/internal/helpers/my/constant"
	e "github.com/cantylv/online-music-lib/internal/helpers/my/error"
	"github.com/cantylv/online-music-lib/internal/usecase/song"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"go.uber.org/zap"
)

type SongHandlerManager struct {
	ucSong song.Contract
	logger *zap.Logger
}

func NewSongHandlerManager(ucSong song.Contract, logger *zap.Logger) *SongHandlerManager {
	return &SongHandlerManager{
		ucSong: ucSong,
		logger: logger,
	}
}

//	@tags			song
//
// GetLibrarySongs godoc
//
//	@Summary		Retrieve a list of library songs
//	@Description	Get songs from the library based on filtering options
//	@ID				get-library-songs
//	@Produce		json
//	@Param			ids					query		string				false	"Song identifiers in format: 'uuid_id1@uuid_id2'"
//	@Param			names				query		string				false	"Song names in format: 'song_name1@song_name2'"
//	@Param			from_release_date	query		string				false	"Left boundary of the interval in format: 'DD-MM-YYYY'"
//	@Param			to_release_date		query		string				false	"Right boundary of the interval in format: 'DD-MM-YYYY'"	default(current time)
//	@Param			text				query		string				false	"Search text in couplets of song, e.g.: 'love'"
//	@Param			limit				query		int					false	"Max number of visible songs, e.g.: 3"	default(2)
//	@Param			offset				query		int					false	"Number of skipped songs, e.g.: 3"		default(0)
//	@Success		200					{array}		entity.Song			"List of songs when songs are found"
//	@Failure		400					{object}	dto.ResponseError	"Bad request - validation errors"
//	@Failure		404					{object}	dto.ResponseError	"Message when no songs are found"
//	@Failure		500					{object}	dto.ResponseError	"Internal server error"
//	@Router			/songs [get]
func (t *SongHandlerManager) GetLibrarySongs(w http.ResponseWriter, r *http.Request) {
	requestID, err := f.GetCtxRequestID(r)
	if err != nil {
		t.logger.Warn(err.Error(), zap.String(mc.RequestID, requestID))
	}
	// получаем параметры фильтрации и валидируем их
	opts, errs := validateFilterLibraryOpts(r)
	if len(errs) != 0 {
		var respErr dto.ResponseError
		for _, err := range errs {
			respErr.Errors = append(respErr.Errors, err.Error())
		}
		f.Response(w, respErr, http.StatusBadRequest)
		return
	}
	songs, err := t.ucSong.GetLibrarySongs(r.Context(), opts)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrInternal.Error()}}, http.StatusInternalServerError)
		return
	}
	if len(songs) == 0 {
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrNotExist.Error()}}, http.StatusNotFound)
		return
	}

	f.Response(w, songs, http.StatusOK)
}

//	@tags			song
// AddNewSongToLibrary godoc
//	@Summary		Create new song
//	@Description	Add new song to the library. Saves it in database.
//	@ID				add-new-song-to-library
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		dto.CreateData		true	"Request body"
//	@Success		200				{object}	entity.Song			"Created song"
//	@Failure		400				{object}	dto.ResponseError	"Bad request - invalid request body"
//	@Failure		500				{object}	dto.ResponseError	"Internal server error"
//	@Router			/songs [post]
func (t *SongHandlerManager) AddNewSongToLibrary(w http.ResponseWriter, r *http.Request) {
	requestID, err := f.GetCtxRequestID(r)
	if err != nil {
		t.logger.Warn(err.Error(), zap.String(mc.RequestID, requestID))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrBadPayload.Error()}}, http.StatusBadRequest)
		return
	}

	var createData dto.CreateData
	err = json.Unmarshal(body, &createData)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrBadPayload.Error()}}, http.StatusBadRequest)
		return
	}

	song, err := t.ucSong.AddNewSongToLibrary(r.Context(), &createData)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrInternal.Error()}}, http.StatusInternalServerError)
		return
	}
	f.Response(w, song, http.StatusOK)
}

//	@tags			song
// GetLibrarySong godoc
//	@Summary		Get library song.
//	@Description	Receiving library song by id.
//	@ID				get-library-song
//	@Produce		json
//	@Param			song_id	path		string				true	"Song identifier in format: 'uuid_id'"
//	@Param			limit	query		int					false	"Max number of visible songs, e.g.: 3"	default(2)
//	@Param			offset	query		int					false	"Number of skipped songs, e.g.: 3"		default(0)
//	@Success		200		{object}	entity.Song			"Song from library"
//	@Failure		400		{object}	dto.ResponseError	"Bad request - invalid request body"
//	@Failure		404		{object}	dto.ResponseError	"Message when no song is found"
//	@Failure		500		{object}	dto.ResponseError	"Internal server error"
//	@Router			/songs/{song_id} [get]
func (t *SongHandlerManager) GetLibrarySong(w http.ResponseWriter, r *http.Request) {
	requestID, err := f.GetCtxRequestID(r)
	if err != nil {
		t.logger.Warn(err.Error(), zap.String(mc.RequestID, requestID))
	}

	opts, errs := validateFilterSongOpts(r)
	if len(errs) != 0 {
		var respErr dto.ResponseError
		for _, err := range errs {
			respErr.Errors = append(respErr.Errors, err.Error())
		}
		f.Response(w, respErr, http.StatusBadRequest)
		return
	}

	song, err := t.ucSong.GetLibrarySong(r.Context(), opts)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrInternal.Error()}}, http.StatusInternalServerError)
		return
	}
	if song == nil {
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrNotExist.Error()}}, http.StatusNotFound)
		return
	}

	f.Response(w, song, http.StatusOK)
}


//	@tags			song
// UpdateLibrarySong godoc
//	@Summary		Update library song.
//	@Description	Update library song by id.
//	@ID				udpate-library-song
//	@Accept			json
//	@Produce		json
//	@Param			song_id	path		string				true	"Song identifier in format: 'uuid_id'"
//	@Param			payload	body		dto.UpdateSong		true	"Data for update library song"
//	@Success		200		{object}	entity.Song			"Updated song from library"
//	@Failure		400		{object}	dto.ResponseError	"Bad request - invalid request body"
//	@Failure		500		{object}	dto.ResponseError	"Internal server error"
//	@Router			/songs/{song_id} [put]
func (t *SongHandlerManager) UpdateLibrarySong(w http.ResponseWriter, r *http.Request) {
	requestID, err := f.GetCtxRequestID(r)
	if err != nil {
		t.logger.Warn(err.Error(), zap.String(mc.RequestID, requestID))
	}

	songID := mux.Vars(r)["song_id"]
	id, err := uuid.FromString(songID)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrBadUUID.Error()}}, http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrBadPayload.Error()}}, http.StatusBadRequest)
		return
	}

	var updateData dto.UpdateSong
	err = json.Unmarshal(body, &updateData)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, e.ErrBadPayload, http.StatusBadRequest)
		return
	}

	updateData.ID = id.String()

	song, err := t.ucSong.UpdateLibrarySong(r.Context(), &updateData)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrInternal.Error()}}, http.StatusInternalServerError)
		return
	}

	f.Response(w, song, http.StatusOK)
}


//	@tags			song
// DeleteLibrarySong godoc
//	@Summary		Delete library song.
//	@Description	Delete library song by id.
//	@ID				delete-library-song
//	@Produce		json
//	@Param			song_id	path		string				true	"Song identifier in format: 'uuid_id'"
//	@Success		200		{object}	dto.ResponseDetail	"Song from library"
//	@Failure		400		{object}	dto.ResponseError	"Bad request - invalid request body"
//	@Failure		404		{object}	dto.ResponseError	"Message when no song is found"
//	@Failure		500		{object}	dto.ResponseError	"Internal server error"
//	@Router			/songs/{song_id} [delete]
func (t *SongHandlerManager) DeleteLibrarySong(w http.ResponseWriter, r *http.Request) {
	requestID, err := f.GetCtxRequestID(r)
	if err != nil {
		t.logger.Warn(err.Error(), zap.String(mc.RequestID, requestID))
	}

	songID := mux.Vars(r)["song_id"]
	id, err := uuid.FromString(songID)
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrBadUUID.Error()}}, http.StatusBadRequest)
		return
	}

	err = t.ucSong.DeleteLibrarySong(r.Context(), id.String())
	if err != nil {
		t.logger.Error(err.Error(), zap.String(mc.RequestID, requestID))
		if errors.Is(err, e.ErrNotExist) {
			f.Response(w, dto.ResponseError{Errors: []string{e.ErrNotExist.Error()}}, http.StatusNotFound)
			return
		}
		f.Response(w, dto.ResponseError{Errors: []string{e.ErrInternal.Error()}}, http.StatusInternalServerError)
		return
	}

	f.Response(w, dto.ResponseDetail{Detail: "song was deleted succesful"}, http.StatusOK)
}
