INSERT INTO users (username, email, image_url) values ('sid', 'sid@sid.com', 'https://avatars.githubusercontent.com/u/73601258?v=4');
INSERT INTO login (hash, username, user_id) values ('$2a$14$EWIWa0JVXjjkdwg9Qw0Gv.ReelDDX6hsgc/7rQfCKyjkxnNAbi4PW' ,'sid', 1);
INSERT INTO users (username, email, image_url) values ('ram', 'ram@ram.com', 'https://avatars.githubusercontent.com/u/54383831?v=4');
INSERT INTO login (hash, username, user_id) values ('$2a$14$kPIHlK0fFuhCE1iWZvgL9ueKFTLm98wCbr6cde2lrm38VpWstEgLm' ,'ram', 2);

INSERT INTO stories (user_id, name, branch, clubs, motto, github_link, linkedin_link, image_url, journey) values (1,'Everyone has their own Story.','EIE','zairza','manners maketh man','https://github.com/sidmohanty11','https://linkedin.com/in/sidmohanty11','https://www.pixsy.com/wp-content/uploads/2021/04/ben-sweet-2LowviVHZ-E-unsplash-1.jpeg', '"If you want a happy ending, that depends, of course, on where you stop your story." – Orson Welles. Everyone has a story. But not everyone tells their story. And we do not always ask. Sharing the stories that shaped your life does not make you weak. It makes you human. When we share stories of the defining moments in our lives, we connect at a deeper level, beyond roles and goals.');

INSERT INTO questions (user_id, q_text) values (1, 'how can i solve this?');
INSERT INTO questions (user_id, q_text) values (2, 'how can i solve this? part2');

INSERT INTO answers (user_id, q_id, a_text) values (1, 1, 'i got itsid');
INSERT INTO answers (user_id, q_id, a_text) values (2, 1, 'i got itram');
INSERT INTO answers (user_id, q_id, a_text) values (1, 2, 'i got it');
INSERT INTO answers (user_id, q_id, a_text) values (2, 2, 'i got it');

INSERT INTO messages (from_user_id, to_user_id, content) values (1, 2, 'hi i am sid');
