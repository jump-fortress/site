import { Client } from '$lib/src/api';

import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
  const monthlyData = Client.GET('/internal/admin/competitions/all/monthly', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return {
    monthlies: monthlyData
  };
};
