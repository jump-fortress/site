import { Client } from '$lib/internalApi';
import type { PlayerProfile } from '$lib/schema';
import type { PageLoad } from './$types';

// parent is used here since +page.ts cannot access locals
export const load: PageLoad = async ({ fetch, params, parent }) => {
  // todo: error handling
  const session = (await parent()).session;

  const { data } = await Client.GET('/internal/players/profile/{id}', {
    fetch: fetch,
    params: {
      path: { id: params.id }
    }
  });

  if (!data) {
    return {
      player: null
    };
  }
  return {
    player: data as PlayerProfile,
    ownProfile: session ? data.id === session.id : false
  };
};
