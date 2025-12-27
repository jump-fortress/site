import config from '$lib/config';
import { Client } from '$lib/internalApi';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
  //check for session before making a request
  if (!event.locals.session) {
    const { data, error } = await Client.GET('/internal/session', {
      baseUrl: config.apiBaseUrl,
      headers: event.request.headers,
      credentials: 'include'
    });

    event.locals.session = data ?? null;

    if (error && error.status !== 401) {
      console.error('There was an error retrieving the session', error);
    }

    // making internal api request
    if (event.url.pathname.startsWith('/internal')) {
      const result = await fetch(new URL(event.url.pathname, config.apiBaseUrl), event.request);
      return result;
    }
  }

  return await resolve(event, {
    filterSerializedResponseHeaders(name) {
      return name === 'content-length';
    }
  });
};
