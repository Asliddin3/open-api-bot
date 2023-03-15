CREATE TABLE if NOT EXISTS search_logs(
  "user_id" BIGINT,
  "success_image_req" BIGINT NOT NULL DEFAULT 0,
  "fail_image_req" BIGINT not NULL DEFAULT 0,
  "success_video_req" BIGINT not NULL DEFAULT 0,
  "fail_video_req" BIGINT not NULL DEFAULT 0,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  Foreign Key (user_id) REFERENCES users(tg_id) on delete CASCADE
);


CREATE OR REPLACE FUNCTION log_insert_search_logs()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
  INSERT INTO search_logs ("user_id") values(NEW.tg_id);
  RETURN NEW;
END;
$$;

CREATE TRIGGER create_search_logs
  AFTER INSERT on "users"
  FOR EACH ROW
  EXECUTE PROCEDURE log_insert_search_logs();



CREATE OR REPLACE FUNCTION log_increment_success_photo_count()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
  UPDATE search_photo as sp set sp.success_image_req=sp.success_image_req+1 where sp.user_id=new.user_id;
	RETURN NEW;

END;
$$;

CREATE TRIGGER  photo_request_changes
  AFTER INSERT on "search_photo"
  FOR EACH ROW
  EXECUTE PROCEDURE log_increment_success_photo_count();



CREATE OR REPLACE FUNCTION log_increment_success_video_count()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
  UPDATE search_video as sv set sv.success_video_req=sv.success_video_req+1 where sv.user_id=new.user_id;
	RETURN NEW;
END;
$$;

CREATE TRIGGER video_request_changes
  AFTER INSERT on "search_video"
  FOR EACH ROW
  EXECUTE PROCEDURE log_increment_success_video_count();