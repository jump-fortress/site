import { Client } from '$lib/internalApi';
import type { PlayerRequest } from '$lib/schema';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  // todo: error handling
  const { data } = await Client.GET('/internal/consultant/players/requests/pending', {
    fetch: fetch
  });

  return {
    request: data as PlayerRequest[]
  };
};
