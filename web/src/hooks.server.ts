import config from '$lib/config';
import { Client } from '$lib/src/api';

import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
  //check for session before making a request
  if (!event.locals.session) {
    const data = Client.GET('/internal/session', {
      fetch: fetch,
      baseUrl: config.apiBaseUrl,
      headers: event.request.headers,
      credentials: 'include'
    }).then((response) => {
      return response.data;
    });

    event.locals.session = data;

    // making internal api request
    if (event.url.pathname.startsWith('/internal')) {
      const result = fetch(new URL(event.url.pathname, config.apiBaseUrl), event.request);
      return result;
    }
  }

  return await resolve(event, {
    filterSerializedResponseHeaders(name) {
      return name === 'content-length';
    }
  });
};
