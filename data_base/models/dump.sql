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



