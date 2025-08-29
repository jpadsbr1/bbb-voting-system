package domain

type BigWallParticipant struct {
	Participant Participant `json:"participant" binding:"required"`
	Votes       int         `json:"votes" binding:"required"`
}

type BigWall struct {
	Participants []BigWallParticipant `json:"participants" binding:"required"`
}

type BigWallRepository interface {
	GetBigWallInfo() (BigWall, error)
	CreateBigWall(Participants []Participant) (BigWall, error)
}
