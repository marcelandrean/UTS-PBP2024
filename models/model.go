package models

// Account adalah ...
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Game adalah ...
type Game struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"max_player"`
}

// Room adalah ...
type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
}

// Participant adalah ...
type Participant struct {
	ID        int `json:"id"`
	IDRoom    int `json:"id_room"`
	IDAccount int `json:"id_account"`
}

// RoomsResponse adalah ...
type RoomsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomWrap `json:"data"`
}

// RoomWrap adalah ...
type RoomWrap struct {
	Rooms []Room `json:"rooms"`
}

// DetailRoomResponse adalah ...
type DetailRoomResponse struct {
	Room DetailRoom `json:"room"`
}

// DetailRoom adalah ...
type DetailRoom struct {
	ID           int                 `json:"id"`
	RoomName     string              `json:"room_name"`
	Participants []ParticipantDetail `json:"participants"`
}

// ParticipantDetail adalah ...
type ParticipantDetail struct {
	ID        int    `json:"id"`
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
}

// DetailRoomsResponse adalah ...
type DetailRoomsResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    DetailRoomResponse `json:"data"`
}

// SuccessResponse adalah ...
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ErrorResponse adalah ...
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
