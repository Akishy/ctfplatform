CREATE TABLE IF NOT EXISTS public.teams_members (
    id serial PRIMARY KEY NOT NULL,
    team_id INTEGER NOT NULL REFERENCES public.teams(team_id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES public.users(user_id) ON DELETE CASCADE,
    is_captain BOOLEAN NOT NULL
);
