import { Client } from '$lib/internalApi';

export const load = async ({ fetch }) => {
  const mapData = Client.GET('/internal/admin/maps', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { maps: mapData };
};
