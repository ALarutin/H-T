CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


-- table person
CREATE TABLE person
(
  id       SERIAL NOT NULL,
  nickname citext NOT NULL,
  email    citext NOT NULL,
  fullname text   NOT NULL,
  about    text   NOT NULL
);

CREATE UNIQUE INDEX person_email_ui
  ON public.person (email);

ALTER TABLE public.person
  ADD CONSTRAINT person_pk PRIMARY KEY (nickname);

CREATE TYPE public.type_person AS (
  is_new BOOLEAN,
  id BIGINT,
  nickname citext,
  email citext,
  about text,
  fullname text
  );

INSERT INTO public."person" (email, about, fullname, nickname)
VALUES ('admin@admin.com', 'something', 'admin', 'admin');

-- table forum
CREATE TABLE forum
(
  id      SERIAL          NOT NULL,
  slug    citext          NOT NULL,
  author  citext          NOT NULL,
  title   text DEFAULT '' NOT NULL,
  posts   INT  DEFAULT 0  NOT NULL,
  threads INT  DEFAULT 0  NOT NULL
);

ALTER TABLE public.forum
  ADD CONSTRAINT forum_pk PRIMARY KEY (slug);

ALTER TABLE ONLY public.forum
  ADD CONSTRAINT "forum_user_fk" FOREIGN KEY (author) REFERENCES public.person (nickname);

CREATE OR REPLACE FUNCTION update_forum_users_on_forum()
  RETURNS trigger AS
$BODY$
BEGIN
  INSERT INTO public."forum_users"(forum_slug, user_nickname)
  VALUES (NEW."slug", NEW."author");
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_forum_users_on_forum
  AFTER INSERT
  ON forum
  FOR EACH ROW
EXECUTE PROCEDURE update_forum_users_on_forum();

-- table forum_users
CREATE TABLE forum_users
(
  forum_slug    citext NOT NULL,
  user_nickname citext NOT NULL
);

ALTER TABLE ONLY public.forum_users
  ADD CONSTRAINT "forum_users_forum_slug_fk" FOREIGN KEY (forum_slug) REFERENCES public.forum (slug);

ALTER TABLE ONLY public.forum_users
  ADD CONSTRAINT "forum_users_user_nickname_fk" FOREIGN KEY (user_nickname) REFERENCES public.person (nickname);

INSERT INTO public."forum" (author, slug)
VALUES ('admin', 'admin');


-- table thread
CREATE TABLE thread
(
  id      SERIAL                                             NOT NULL,
  slug    citext                                             NOT NULL,
  author  citext                                             NOT NULL,
  forum   citext                                             NOT NULL,
  title   text                     DEFAULT ''                NOT NULL,
  message text                     DEFAULT ''                NOT NULL,
  votes   INT                      DEFAULT 0                 NOT NULL,
  created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX thread_id_ui
  ON public.thread (id);

ALTER TABLE public.thread
  ADD CONSTRAINT thread_pk PRIMARY KEY (slug);

ALTER TABLE ONLY public.thread
  ADD CONSTRAINT "thread_author_fk" FOREIGN KEY (author) REFERENCES public.person (nickname);

ALTER TABLE ONLY public.thread
  ADD CONSTRAINT "thread_forum_fk" FOREIGN KEY (forum) REFERENCES public.forum (slug);

CREATE OR REPLACE FUNCTION update_threads_quantity()
  RETURNS trigger AS
$BODY$
BEGIN
  UPDATE public."forum"
  SET threads = threads + 1
  WHERE "slug" = NEW."forum";
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_forum_threads
  AFTER INSERT
  ON thread
  FOR EACH ROW
EXECUTE PROCEDURE update_threads_quantity();

CREATE OR REPLACE FUNCTION update_forum_users_on_thread()
  RETURNS trigger AS
$BODY$
BEGIN
  INSERT INTO public."forum_users"(forum_slug, user_nickname)
  VALUES (NEW."forum", NEW."author");
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_forum_users_on_thread
  AFTER INSERT
  ON thread
  FOR EACH ROW
EXECUTE PROCEDURE update_forum_users_on_thread();

INSERT INTO public."thread" (author, forum, slug)
VALUES ('admin', 'admin', 'admin');


-- table post
CREATE TABLE post
(
  id        SERIAL                                             NOT NULL,
  author    citext                                             NOT NULL,
  thread    INT                                                NOT NULL,
  forum     citext                                             NOT NULL,
  message   text                     DEFAULT ''                NOT NULL,
  is_edited boolean                  DEFAULT FALSE             NOT NULL,
  parent    int                                                NOT NULL,
  created   timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  post_path integer[]                DEFAULT '{}'::integer[]
);

ALTER TABLE public.post
  ADD CONSTRAINT post_pk PRIMARY KEY (id);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_author_fk" FOREIGN KEY (author) REFERENCES public.person (nickname);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_thread_fk" FOREIGN KEY (thread) REFERENCES public.thread (id);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_forum_fk" FOREIGN KEY (forum) REFERENCES public.forum (slug);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_parent_fk" FOREIGN KEY (parent) REFERENCES public.post (id);

CREATE OR REPLACE FUNCTION update_posts_quantity()
  RETURNS trigger AS
$BODY$
BEGIN
  UPDATE public."forum"
  SET posts = posts + 1
  WHERE "slug" = NEW."forum";
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_forum_posts
  AFTER INSERT
  ON post
  FOR EACH ROW
EXECUTE PROCEDURE update_posts_quantity();

CREATE OR REPLACE FUNCTION update_post_quantity()
  RETURNS trigger AS
$BODY$
BEGIN
  UPDATE public."post"
  SET is_edited = TRUE
  WHERE "id" = NEW."id";
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_post
  AFTER UPDATE OF message
  ON post
  FOR EACH ROW
EXECUTE PROCEDURE update_post_quantity();

INSERT INTO public."post" (id, author, thread, forum, parent)
VALUES (0, 'admin', '1', 'admin', 0);


-- table vote
CREATE TABLE vote
(
  thread_slug   citext NOT NULL,
  user_nickname citext NOT NULL,
  voice         INT    NOT NULL
);

ALTER TABLE public.vote
  ADD CONSTRAINT vote_pk PRIMARY KEY (thread_slug, user_nickname);

ALTER TABLE ONLY public.vote
  ADD CONSTRAINT "vote_thread_slug_fk" FOREIGN KEY (thread_slug) REFERENCES public.thread (slug);

ALTER TABLE ONLY public.vote
  ADD CONSTRAINT "vote_user_nickname_fk" FOREIGN KEY (user_nickname) REFERENCES public.person (nickname);

CREATE OR REPLACE FUNCTION update_votes()
  RETURNS trigger AS
$BODY$
BEGIN
  UPDATE public."thread"
  SET votes = votes + New."voice"
  WHERE "slug" = NEW."thread_slug";
  RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE TRIGGER update_thread_votes
  AFTER INSERT
  on vote
  FOR EACH ROW
EXECUTE PROCEDURE update_votes();

CREATE OR REPLACE FUNCTION add_admin()
  RETURNS VOID AS
$BODY$
BEGIN
  INSERT INTO public."person" (email, about, fullname, nickname)
  VALUES ('admin@admin.com', 'something', 'admin', 'admin');
  INSERT INTO public."forum" (author, slug)
  VALUES ('admin', 'admin');
  INSERT INTO public."thread" (author, forum, slug)
  VALUES ('admin', 'admin', 'admin');
  INSERT INTO public."post" (id, author, thread, forum, parent)
  VALUES (0, 'admin', '1', 'admin', 0);
END;
$BODY$
  LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION clear_database()
  RETURNS VOID AS
$BODY$
BEGIN
  TRUNCATE TABLE public.forum, public.forum_users, public.person, public.post, public.thread, public.vote
    RESTART IDENTITY;
  PERFORM add_admin();
END;
$BODY$
  LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION func_create_user(arg_nickname citext, arg_email citext, arg_fullname text, arg_about text)
  RETURNS SETOF public.type_person
AS
$BODY$
DECLARE
  result public.type_person;
  rec    RECORD;
BEGIN
  INSERT INTO person (nickname, email, fullname, about)
  VALUES (arg_nickname, arg_email, arg_fullname, arg_about) RETURNING *
    INTO result.id, result.nickname, result.fullname, result.about, result.email;
  result.is_new := true;
  RETURN next result;
EXCEPTION
  WHEN unique_violation THEN
    FOR rec IN SELECT *
               FROM public.person
               WHERE nickname = arg_nickname
                  OR email = arg_email
      LOOP
        result.id := rec.id;
        result.nickname := rec.nickname;
        result.fullname := rec.fullname;
        result.about := rec.about;
        result.email := rec.email;
        result.is_new := false;
        RETURN NEXT result;
      END LOOP;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION func_get_user(arg_nickname citext)
  RETURNS public.type_person
AS
$BODY$
DECLARE
  result public.type_person;
BEGIN
  SELECT *
    INTO result.id, result.nickname, result.fullname, result.about, result.email
  FROM public.person
    WHERE nickname = arg_nickname;
  result.is_new := FALSE;
  IF NOT FOUND THEN
    RAISE no_data_found;
  END IF;
  RETURN result;
END;
$BODY$
  LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION func_update_user(arg_nickname citext, arg_email citext, arg_fullname text, arg_about text)
  RETURNS public.type_person
AS
$BODY$
DECLARE
  result public.type_person;
BEGIN
  UPDATE public.person
  SET email = arg_email, fullname = arg_fullname, about = arg_about
  WHERE nickname = arg_nickname RETURNING *
    INTO result.id, result.nickname, result.email, result.fullname, result.about;
  result.is_new := FALSE;
  IF NOT FOUND THEN
    RAISE no_data_found;
  END IF;
  RETURN result;
EXCEPTION
  WHEN unique_violation THEN
    RAISE  unique_violation;
END;
$BODY$
  LANGUAGE plpgsql;