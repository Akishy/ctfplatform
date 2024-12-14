CREATE TABLE IF NOT EXISTS public.teams_members (
    id serial PRIMARY KEY NOT NULL,
    team_id INTEGER NOT NULL REFERENCES public.teams(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    team_role_id INTEGER NOT NULL REFERENCES public.team_roles(id) ON DELETE CASCADE
);
