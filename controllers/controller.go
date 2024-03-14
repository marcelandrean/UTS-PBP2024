package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strconv"

	m "uts/models"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db, errDb := connectGorm()
	if errDb != nil {
		sendErrorResponse(w, "Failed to connect database")
		return
	}

	var rooms []m.Room
	query := db.Model(&m.Room{})

	// Read from Query Param
	gameID := r.URL.Query().Get("id_game")

	if gameID != "" {
		query = query.Where("id_game = ?", gameID)
	}

	err := query.Find(&rooms).Error
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		sendErrorResponse(w, "Failed to fetch rooms")
		return
	}

	var response m.RoomsResponse
	if len(rooms) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data.Rooms = rooms
	} else {
		response.Status = 400
		response.Message = "No rooms found"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.id, r.room_name, p.id, p.id_account, a.username
		FROM rooms r JOIN participants p ON r.id = p.id_room
		JOIN accounts a ON p.id_account = a.id
		WHERE r.id = ?`

	roomID := mux.Vars(r)["id"]

	detailRoomRow, err := db.Query(query, roomID)
	if err != nil {
		fmt.Println(roomID)
		print(err.Error())
		sendErrorResponse(w, "Invalid query")
		return
	}

	var detailRoom m.DetailRoom
	var participants []m.ParticipantDetail

	for detailRoomRow.Next() {
		var participant m.ParticipantDetail
		if err := detailRoomRow.Scan(
			&detailRoom.ID, &detailRoom.RoomName, &participant.ID, &participant.ID_Account, &participant.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Failed to fetch detail room")
			return
		}
		participants = append(participants, participant)
	}

	detailRoom.Participants = participants

	var response m.DetailRoomsResponse
	if len(participants) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data.Room = detailRoom
	} else {
		response.Status = 400
		response.Message = "No participants found"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
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

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.SuccessResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
