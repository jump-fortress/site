import { Client } from '$lib/src/api';

import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  const data = Client.GET('/internal/consultant/players/all', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return {
    players: data
  };
};
