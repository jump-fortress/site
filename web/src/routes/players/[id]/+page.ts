import { Client } from '$lib/internalApi';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
  const { data } = await Client.GET('/internal/players/profile/{id}', {
    fetch: fetch,
    params: {
      path: { id: params.id }
    }
  });

  return {
    playerProfile: data ?? null
  };
};
