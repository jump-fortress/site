import type { PageServerLoad } from './$types';
import { Client } from '$lib/api/api';
import { ApiPaths } from '$lib/schema';

export const load: PageServerLoad = async ({ params, fetch }) => {
  const { data } = await Client.GET(ApiPaths.get_player, {
    fetch: fetch,
    params: { path: { player_id: params.id } }
  });
  return { player: data };
};
