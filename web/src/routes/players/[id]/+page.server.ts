import { Client } from '$lib/internalApi';
import type { PageServerLoad } from './$types';

// this call doesn't need credentials
export const load: PageServerLoad = async ({ fetch, params }) => {
  const data = Client.GET('/internal/players/profile/{id}', {
    fetch: fetch,
    params: {
      path: { id: params.id }
    }
  }).then((response) => {
    return response.data;
  });

  return { playerProfile: data };
};
