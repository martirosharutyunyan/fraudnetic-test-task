CREATE KEYSPACE fraudnetic WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor' : 1};

CREATE TABLE events (
     id uuid,
     browser_fingerprint bigint,
     canvas_fingerprint bigint,
     created_at timestamp,
     device_language text,
     device_timezone varint,
     event_name text,
     font_fingerprint bigint,
     incognito boolean,
     ip inet,
     periodic_wave bigint,
     processed boolean,
     screen_resolution text,
     session text,
     status boolean,
     storage bigint,
     updated_at timestamp,
     user_agent text,
     user_id varint,
     utm text,
     webgl_fingerprint bigint,
     PRIMARY KEY (id, created_at)
 )

