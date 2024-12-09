-- chairs
CREATE INDEX idx_owner_id ON chairs (owner_id);
CREATE INDEX idx_model ON chairs (model(255));
CREATE INDEX idx_is_active ON chairs (is_active);

-- chair_locations
CREATE INDEX idx_chair_id ON chair_locations (chair_id);
CREATE INDEX idx_created_at ON chair_locations (created_at);

-- coupons
CREATE INDEX idx_code ON coupons (code);
CREATE INDEX idx_coupons ON coupons (user_id);

-- owners
CREATE INDEX idx_owner_register_token ON owners (chair_register_token);

-- payment_tokens
CREATE INDEX idx_user_id ON payment_tokens (user_id);

-- rides
CREATE INDEX idx_chair_id ON rides (chair_id);
CREATE INDEX idx_user_id ON rides (user_id);

-- ride_statuses
CREATE INDEX idx_ride_id ON ride_statuses (ride_id);
CREATE INDEX idx_status ON ride_statuses (status);

-- settings
CREATE INDEX idx_name ON settings (name);

-- users
CREATE INDEX idx_user_invitation_code ON users (invitation_code);
