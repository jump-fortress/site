import { Client } from '$lib/src/api';

import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
  const data = Client.GET('/internal/maps/names', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { maps: data };
};
