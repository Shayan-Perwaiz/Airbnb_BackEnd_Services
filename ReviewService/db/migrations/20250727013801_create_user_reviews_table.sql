-- +goose Up
CREATE TABLE user_reviews (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    hotel_id BIGINT NOT NULL,
    booking_id BIGINT NOT NULL,
    comment TEXT,
    rating VARCHAR(10),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE user_reviews;
