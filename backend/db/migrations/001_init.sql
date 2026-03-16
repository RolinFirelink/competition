CREATE TABLE IF NOT EXISTS users (
  street_id      VARCHAR(64) PRIMARY KEY,
  nick           VARCHAR(64) NOT NULL,
  tier           VARCHAR(32),
  role           VARCHAR(32),
  level          VARCHAR(32),
  plan_slots     INT DEFAULT 1,
  plan_used      INT DEFAULT 0,
  penalized      TINYINT(1) DEFAULT 0,
  penalty_reason VARCHAR(255),
  password_hash  VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS growth_routes (
  id           BIGINT PRIMARY KEY AUTO_INCREMENT,
  student_id   VARCHAR(64) NOT NULL,
  character_   VARCHAR(64),
  question     VARCHAR(255),
  video_id     VARCHAR(128),
  tier         VARCHAR(32),
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (student_id) REFERENCES users(street_id)
);

CREATE TABLE IF NOT EXISTS replies (
  id         BIGINT PRIMARY KEY AUTO_INCREMENT,
  route_id   BIGINT NOT NULL,
  author     VARCHAR(32),
  content    TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (route_id) REFERENCES growth_routes(id)
);

CREATE TABLE IF NOT EXISTS events (
  id         BIGINT PRIMARY KEY AUTO_INCREMENT,
  title      VARCHAR(128),
  announce   TEXT,
  groups     JSON,
  awards     JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS participants (
  id         BIGINT PRIMARY KEY AUTO_INCREMENT,
  event_id   BIGINT NOT NULL,
  street_id  VARCHAR(64),
  nick       VARCHAR(64),
  tier       VARCHAR(32),
  FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE IF NOT EXISTS banners (
  id     BIGINT PRIMARY KEY AUTO_INCREMENT,
  title  VARCHAR(128),
  image  VARCHAR(255),
  link   VARCHAR(255)
);

