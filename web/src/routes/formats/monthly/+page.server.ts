import { Client } from '$lib/src/api';

import type { PageServerLoad } from './$types';

// this call doesn't need credentials
export const load: PageServerLoad = async ({ fetch }) => {
  const data = Client.GET('/internal/competitions/all/monthly', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { monthlies: data };
};
