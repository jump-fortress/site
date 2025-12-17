import { Client } from '$lib/internalApi';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  // todo: error handling
  const data = Client.GET('/internal/consultant/players/requests/pending', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { playersWithRequests: data };
};
