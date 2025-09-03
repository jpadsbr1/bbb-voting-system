-- Participants
CREATE TABLE participants (
    participant_id CHAR(36) PRIMARY KEY,
    name TEXT NOT NULL
);

-- Big Wall
CREATE TABLE bigwall (
    bigwall_id CHAR(36) PRIMARY KEY,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Relation Participant_BigWall
CREATE TABLE participants_bigwall (
    bigwall_id CHAR(36) REFERENCES bigwall(bigwall_id) ON DELETE CASCADE,
    participant_id CHAR(36) REFERENCES participants(participant_id) ON DELETE CASCADE,
    PRIMARY KEY (bigwall_id, participant_id)
);

-- Votes
CREATE TABLE votes (
    vote_id BIGSERIAL PRIMARY KEY,
    bigwall_id CHAR(36) REFERENCES bigwall(bigwall_id) ON DELETE CASCADE,
    participant_id CHAR(36) REFERENCES participants(participant_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);