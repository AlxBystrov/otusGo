CREATE table IF NOT EXISTS events (
    id    		  uuid primary key,
    title 		  text,
	date  		  timestamp,
	duration      interval,
	description   text,
	user_id       bigint,
	notify_before interval
    );