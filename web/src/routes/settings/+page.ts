import { Client } from '$lib/internalApi';

export const load = async ({ fetch }) => {
  const playerData = Client.GET('/internal/players', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  const playerRequestsData = Client.GET('/internal/players/requests', {
    fetch: fetch
  }).then((response) => {
    return response.data;
  });

  return { player: playerData, requestPreviews: playerRequestsData };
};
