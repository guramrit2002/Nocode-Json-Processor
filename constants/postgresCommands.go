package constants

var WidgetSchemaCommand = `
	CREATE TABLE IF NOT EXISTS widgets (
		id TEXT PRIMARY KEY,          -- Unique identifier for the widget
		type TEXT NOT NULL,           -- Type of the widget
		style JSONB NOT NULL,         -- Style as key-value pairs (JSONB format)
		content TEXT,                 -- Optional content
		parent_id TEXT,               -- Parent widget ID (self-referential relationship)
		FOREIGN KEY (parent_id) REFERENCES widgets(id) ON DELETE CASCADE
	);
	`

var WidgetInsertCommand = `
		INSERT INTO widgets (id, type, style, content, parent_id)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO NOTHING;
	`
