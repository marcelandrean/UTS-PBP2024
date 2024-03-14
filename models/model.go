package models

// Account...
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Game...
type Game struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"max_player"`
}

// Room...
type Room struct {
	ID        int    `json:"id"`
	Room_Name string `json:"room_name"`
}

// Participant...
type Participant struct {
	ID        int `json:"id"`
	IDRoom    int `json:"id_room"`
	IDAccount int `json:"id_account"`
}

// RoomsResponse...
type RoomsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomWrap `json:"data"`
}

// RoomWrap...
type RoomWrap struct {
	Rooms []Room `json:"rooms"`
}

// DetailRoomResponse...
type DetailRoomResponse struct {
	Room DetailRoom `json:"room"`
}

// DetailRoom...
type DetailRoom struct {
	ID           int                 `json:"id"`
	RoomName     string              `json:"room_name"`
	Participants []ParticipantDetail `json:"participants"`
}

// ParticipantDetail...
type ParticipantDetail struct {
	ID        int    `json:"id"`
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
}

// DetailRoomsResponse...
type DetailRoomsResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    DetailRoomResponse `json:"data"`
}

// SuccessResponse...
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ErrorResponse...
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
