// retrieve monthly from param id
import { Client } from '$lib/src/api';

import type { PageServerLoad } from './$types';

// this call doesn't need credentials
export const load: PageServerLoad = async ({ fetch, params }) => {
  const monthlyData = Client.GET('/internal/competitions/monthly/{id}', {
    fetch: fetch,
    params: {
      path: { id: parseInt(params.id) }
    }
  }).then((response) => {
    return response.data;
  });

  const timesData = Client.GET('/internal/players/times/{competition_type}/{id}', {
    fetch: fetch,
    params: {
      path: {
        competition_type: 'Monthly',
        id: parseInt(params.id)
      }
    }
  }).then((response) => {
    return response.data;
  });

  return { monthly: monthlyData, times: timesData };
};
