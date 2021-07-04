INSERT INTO users (id, hash, username, email, image_url) values (1,'$2a$14$EWIWa0JVXjjkdwg9Qw0Gv.ReelDDX6hsgc/7rQfCKyjkxnNAbi4PW' ,'sid', 'sid@sid.com', 'https://avatars.githubusercontent.com/u/73601258?v=4');
INSERT INTO login (id, hash, username) values (1,'$2a$14$EWIWa0JVXjjkdwg9Qw0Gv.ReelDDX6hsgc/7rQfCKyjkxnNAbi4PW' ,'sid');
INSERT INTO users (id, hash, username, email, image_url) values (2, '$2a$14$kPIHlK0fFuhCE1iWZvgL9ueKFTLm98wCbr6cde2lrm38VpWstEgLm' ,'ram', 'ram@ram.com', 'https://avatars.githubusercontent.com/u/54383831?v=4');
INSERT INTO login (id, hash, username) values (2, '$2a$14$kPIHlK0fFuhCE1iWZvgL9ueKFTLm98wCbr6cde2lrm38VpWstEgLm' ,'ram');
INSERT INTO users (id, hash, username, email, image_url) values (3, '$2a$14$kPIHlK0fFuhCE1iWZvgL9ueKFTLm98wCbr6cde2lrm38VpWstEgLm', 'shyam', 'shyam@shyam.com', 'https://avatars.githubusercontent.com/u/54383831?v=4');
INSERT INTO login (id, hash, username) values (3, '$2a$14$kPIHlK0fFuhCE1iWZvgL9ueKFTLm98wCbr6cde2lrm38VpWstEgLm', 'shyam');

INSERT INTO stories (id, user_id, name, branch, clubs, motto, github_link, linkedin_link, image_url, journey) values (1,1,'sid','branch','clubs','my motto','github.com/sidmohanty11','linkedin.com/in/sidmohanty11','https://avatars.githubusercontent.com/u/73601258?v=4', 'journey');

INSERT INTO blogs (id, user_id, blog_title, blog_text) values (1,1,'blog title', 'lorem ipsum i cant type');

INSERT INTO questions (id, user_id, q_text) values (1, 1, 'how can i solve this?');
INSERT INTO questions (id, user_id, q_text) values (2, 2, 'how can i solve this? part2');

INSERT INTO answers (id, user_id, q_id, a_text) values (1, 1, 1, 'i got itsid');
INSERT INTO answers (id, user_id, q_id, a_text) values (2, 2, 1, 'i got itram');
INSERT INTO answers (id, user_id, q_id, a_text) values (3, 1, 2, 'i got it');
INSERT INTO answers (id, user_id, q_id, a_text) values (4, 2, 2, 'i got it');

INSERT INTO messages (id, from_user_id, to_user_id, content) values (1, 1, 2, 'hi i am sid');