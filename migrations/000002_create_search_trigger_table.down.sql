DROP Trigger if EXISTS photo_request_changes
on search_photo ;
DROP Trigger if EXISTS video_request_changes
on search_video;
DROP Function if EXISTS log_increment_success_video_count;
DROP Function if EXISTS log_increment_success_photo_count;
DROP TRIGGER if EXISTS create_search_logs
On users;
DROP FUNCTION if EXISTS insert_search_logs;
DROP TABLE if EXISTS search_logs;