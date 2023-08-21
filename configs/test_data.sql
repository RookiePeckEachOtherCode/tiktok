INSERT INTO `tiktok-test`.user_infos (id, name, follow_count, follower_count, is_follow, avatar)
VALUES (5,'TestJudgeUserPassword0',0,0,0,'defualt.jpg');
INSERT INTO `tiktok-test`.user_logins (id, user_info_id, username, password)
VALUES (5,5,'TestJudgeUserPassword','123456');

INSERT INTO `tiktok-test`.user_infos (id, name, follow_count, follower_count, is_follow, avatar)
VALUES (6,'TestGetUserInfoById',0,0,0,'defualt.jpg');
INSERT INTO  `tiktok-test`.user_logins (id, user_info_id, username, password)
VALUES (6,6,'TestGetUserInfoById','123456');


