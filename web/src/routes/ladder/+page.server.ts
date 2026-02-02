import type { PageServerLoad } from './$types';
import { Client } from '$lib/api/api';
import { ApiPaths } from '$lib/schema';

export const load: PageServerLoad = async ({ fetch }) => {
  const { data } = await Client.GET(ApiPaths.get_players, {
    fetch: fetch
  });
  return { players: data };
};
