package controllers

import (
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// EnterRoom...
func EnterRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse form")
		return
	}

	roomID := r.Form.Get("id_room")
	if roomID == "" {
		sendErrorResponse(w, "Room ID is required")
		return
	}

	// Check apakah ada room dengan ID tersebut
	var countRoom int
	var gameID int
	err = db.QueryRow("SELECT COUNT(*), id_game FROM rooms WHERE id = ?", roomID).Scan(&countRoom, &gameID)

	if err != nil {
		sendErrorResponse(w, "Failed to fetch rooms data")
		return
	}
	if countRoom == 0 {
		sendErrorResponse(w, "Room does not exist")
		return
	}

	// Get max player limit dari game
	var maxPlayer int
	err = db.QueryRow("SELECT max_player FROM games WHERE id = ?", gameID).Scan(&maxPlayer)
	if err != nil {
		sendErrorResponse(w, "Failed to fetch games data")
		return
	}

	// Check apakah room udah melebihi limit max player
	var currentPlayers int
	err = db.QueryRow("SELECT COUNT(*) FROM participants WHERE id_room = ?", roomID).Scan(&currentPlayers)
	if err != nil {
		sendErrorResponse(w, "Failed to fetch participants data")
		return
	}

	if currentPlayers >= maxPlayer {
		sendErrorResponse(w, "Room has reached maximum players limit")
		return
	}

	accountID := r.Form.Get("id_account")
	if accountID == "" {
		sendErrorResponse(w, "Account ID is required")
		return
	}

	// Check apakah account sudah masuk di salah satu room
	var countParticipant int
	err = db.QueryRow("SELECT COUNT(*) FROM participants WHERE id_account = ?", accountID).Scan(&countParticipant)
	if err != nil {
		sendErrorResponse(w, "Failed to check existence of participant")
		return
	}

	if countParticipant > 0 {
		sendErrorResponse(w, "Account / participant already exists in any room.")
		return
	}

	_, errQuery := db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if errQuery == nil {
		sendSuccessResponse(w, "Insert participant to room success")
	} else {
		sendErrorResponse(w, "Failed to insert")
	}
}

// LeaveRoom...
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	acountID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorResponse(w, "Invalid participant ID")
		return
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM participants WHERE id_account=?", acountID).Scan(&count)
	if err != nil {
		sendErrorResponse(w, "Failed to check participant existence")
		return
	}

	if count == 0 {
		sendErrorResponse(w, "Participant does not exist")
		return
	}

	_, errQuery := db.Exec("DELETE FROM participants WHERE id_account=?", acountID)

	if errQuery == nil {
		sendSuccessResponse(w, "Leave room success")
	} else {
		sendErrorResponse(w, "Failed to leave room")
	}
}
