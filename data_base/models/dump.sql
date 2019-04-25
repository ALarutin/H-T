CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;

-- table person
CREATE TABLE person
(
  id       serial not null,
  nickname citext not null,
  email    citext not null,
  fullname text   not null,
  about    text   not null
);

CREATE UNIQUE INDEX person_email_ui
  ON public.person (email);

ALTER TABLE public.person
  ADD CONSTRAINT person_pk PRIMARY KEY (nickname);


-- table forum
CREATE TABLE forum
(
  id      serial not null,
  slug    citext not null,
  author  citext not null,
  title   text   not null,
  posts   int    not null,
  threads int    not null
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
  on forum
  FOR EACH ROW
EXECUTE PROCEDURE update_forum_users_on_forum();


-- table thread
CREATE TABLE thread
(
  id      serial                                             not null,
  slug    citext                                             not null,
  author  citext                                             not null,
  forum   citext                                             not null,
  title   text                                               not null,
  message text                                               not null,
  votes   int                                                not null,
  created timestamp with time zone DEFAULT CURRENT_TIMESTAMP not null
);

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
  on thread
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
  on thread
  FOR EACH ROW
EXECUTE PROCEDURE update_forum_users_on_thread();


-- table post
CREATE TABLE post
(
  id        serial                                             not null,
  author    citext                                             not null,
  thread    citext                                             not null,
  forum     citext                                             not null,
  message   text                                               not null,
  is_edited boolean                  default false             not null,
  parent    int,
  created   timestamp with time zone DEFAULT CURRENT_TIMESTAMP not null,
  post_path integer[]                DEFAULT '{}'::integer[]
);

ALTER TABLE public.post
  ADD CONSTRAINT post_pk PRIMARY KEY (id);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_author_fk" FOREIGN KEY (author) REFERENCES public.person (nickname);

ALTER TABLE ONLY public.post
  ADD CONSTRAINT "post_thread_fk" FOREIGN KEY (thread) REFERENCES public.thread (slug);

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
  on post
  FOR EACH ROW
EXECUTE PROCEDURE update_posts_quantity();


-- table forum_users
CREATE TABLE forum_users
(
  forum_slug    citext not null,
  user_nickname citext not null
);

ALTER TABLE ONLY public.forum_users
  ADD CONSTRAINT "forum_users_forum_slug_fk" FOREIGN KEY (forum_slug) REFERENCES public.forum (slug);

ALTER TABLE ONLY public.forum_users
  ADD CONSTRAINT "forum_users_user_nickname_fk" FOREIGN KEY (user_nickname) REFERENCES public.person (nickname);


-- table vote
CREATE TABLE vote
(
  thread_slug   citext not null,
  user_nickname citext not null,
  voice         int    not null
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
  AFTER INSERT OR UPDATE
  on vote
  FOR EACH ROW
EXECUTE PROCEDURE update_votes();
