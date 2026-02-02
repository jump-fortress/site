import type { PageServerLoad } from './$types';
import { Client } from '$lib/api/api';
import { ApiPaths } from '$lib/schema';

export const load: PageServerLoad = async ({ params, fetch }) => {
  const { data } = await Client.GET(ApiPaths.get_event, {
    fetch: fetch,
    params: { path: { event_kind: params.kind, kind_id: parseInt(params.kind_id) } }
  });
  return { ewl: data };
};
