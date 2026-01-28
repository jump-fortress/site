import { redirect } from '@sveltejs/kit';

import type { PageServerLoad } from './$types';
import { Client } from '$lib/api/api';
import { ApiPaths } from '$lib/schema';

export const load: PageServerLoad = async ({ parent, fetch }) => {
  const session = await (await parent()).session;
  if (!session) {
    redirect(302, '/');
  }

  const { data } = await Client.GET(ApiPaths.get_player, {
    fetch: fetch,
    params: { path: { player_id: session.id } }
  });
  return { player: data };
};
