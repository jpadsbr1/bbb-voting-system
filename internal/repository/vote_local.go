package repository

type VoteLocalRepository struct {
	total_votes int
}

func NewVoteLocalRepository() *VoteLocalRepository {
	return &VoteLocalRepository{total_votes: 0}
}

func (r *VoteLocalRepository) SaveVote(ParticipantID string) error {
	r.total_votes++
	return nil
}

func (r *VoteLocalRepository) GetTotalVotes() (int, error) {
	return r.total_votes, nil
}
