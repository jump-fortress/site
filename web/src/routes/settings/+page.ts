import { Client } from '$lib/internalApi';
import type { FullPlayer } from '$lib/schema';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  // todo: error handling
  const { data } = await Client.GET('/internal/players', {
    fetch: fetch
  });

  if (!data) {
    return { fullPlayer: null };
  }

  return {
    fullPlayer: data as FullPlayer
  };
};
