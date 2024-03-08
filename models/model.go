package models

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_Player int    `json:"max_player"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_Name string `json:"room_name"`
}

type Participant struct {
	ID         int `json:"id"`
	ID_Room    int `json:"id_room"`
	ID_Account int `json:"id_account"`
}

type RoomsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomWrap `json:"data"`
}

type RoomWrap struct {
	Rooms []Room `json:"rooms"`
}

type DetailRoomResponse struct {
	Room DetailRoom `json:"room"`
}

type DetailRoom struct {
	ID           int                `json:"id"`
	RoomName     string             `json:"room_name"`
	Participants []ParticipantDetail `json:"participants"`
}

type ParticipantDetail struct {
	ID         int    `json:"id"`
	ID_Account int    `json:"id_account"`
	Username   string `json:"username"`
}

type DetailRoomsResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    DetailRoomResponse `json:"data"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
