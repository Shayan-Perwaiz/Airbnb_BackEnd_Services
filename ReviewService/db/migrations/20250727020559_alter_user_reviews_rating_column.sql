-- +goose Up
ALTER TABLE user_reviews
MODIFY COLUMN rating DECIMAL(2,1);

-- +goose Down
ALTER TABLE user_reviews
MODIFY COLUMN rating VARCHAR(10);
