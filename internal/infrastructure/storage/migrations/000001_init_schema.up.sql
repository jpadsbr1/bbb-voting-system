CREATE TABLE participants (
    participant_id CHAR(36) PRIMARY KEY,
    name TEXT NOT NULL,
    is_eliminated BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE bigwall (
    bigwall_id CHAR(36) PRIMARY KEY,
    start_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    end_time TIMESTAMPTZ,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    total_votes INT NOT NULL DEFAULT 0
);

CREATE TABLE participants_bigwall (
    bigwall_id CHAR(36) REFERENCES bigwall(bigwall_id) ON DELETE CASCADE,
    participant_id CHAR(36) REFERENCES participants(participant_id) ON DELETE CASCADE,
    votes INT NOT NULL DEFAULT 0,
    PRIMARY KEY (bigwall_id, participant_id)
);

CREATE TABLE votes_hourly (
    vote_id BIGSERIAL PRIMARY KEY,
    bigwall_id CHAR(36) REFERENCES bigwall(bigwall_id) ON DELETE CASCADE,
    participant_id CHAR(36) REFERENCES participants(participant_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);