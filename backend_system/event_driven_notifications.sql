CREATE DATABASE IF NOT EXISTS realtime_notification_system;

USE realtime_notification_system;

CREATE TABLE IF NOT EXISTS events  (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event TEXT NOT NULL,
    time_stamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP

    

);