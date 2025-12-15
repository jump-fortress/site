import { Client } from '$lib/internalApi';
import type { FullPlayer } from '$lib/schema';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  // todo: error handling
  const { data } = await Client.GET('/internal/consultant/players/all', {
    fetch: fetch
  });
  if (!data) {
    return {
      fullPlayers: null
    };
  }

  return {
    fullPlayers: data as FullPlayer[]
  };
};
