CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    notif_title TEXT NOT NULL,
    notif_heading TEXT NOT NULL,
    notif_body TEXT NOT NULL,
    days_to_Last INTEGER NOT NULL,
    notif_time TEXT NOT NULL,
    notif_id TEXT NOT NULL
);