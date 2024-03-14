package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
