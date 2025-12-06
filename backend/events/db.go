package events

import "database/sql"

type EventDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) EventDB {
	return EventDB{db}
}

func (edb *EventDB) Init() {
	scheme := `
	CREATE TABLE events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL, 
	text TEXT NOT NULL,
	location TEXT NOT NULL
	);

	CREATE TABLE creations(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER NOT NULL,
	organisator_id INTEGER NOT NULL,
	FOREIGN KEY event_id REFERENCES events(id),
	FOREIGN KEY organisator_id REFERENCES creds(id)
	);
	`
	edb.db.Exec(scheme)
}

func (edb *EventDB) CreateEvent(title, text, location string) error {
	_, err := edb.db.Exec("INSERT INTO events (title, text, location) VALUES (?,?,?)", title, text, location)
	return err
}

func (edb *EventDB) GetEvents(scheme string) ([]Event, error) {
	var events []Event
	res, err := edb.db.Query(scheme)
	if err != nil {
		return events, err
	}
	for res.Next() {
		var event Event
		err := res.Scan(&event.ID, &event.Title, &event.Text, &event.Location)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (edb *EventDB) UpdateEvent(newEvent Event) error {
	_, err := edb.db.Exec("UPDATE events SET title = ?, text = ?, location = ? WHERE id = ?", newEvent.Title, newEvent.Text, newEvent.Location, newEvent.ID)
	if err != nil {
		return err
	}
	return nil
}
func (edb *EventDB) RemoveEvent(id int64) error {
	_, err := edb.db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
