import { Client } from '$lib/src/api';

import type { PageServerLoad } from './$types';

// this call doesn't need credentials
export const load: PageServerLoad = async ({ fetch }) => {
  const data = Client.GET('/internal/players/all', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { playerPreviews: data };
};
