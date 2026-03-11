import config from '$lib/api/config';
import { Client } from '$lib/api/api';

import type { Handle } from '@sveltejs/kit';
import { ApiPaths } from '$lib/schema';

import { PUBLIC_JUMP_SESSION_COOKIE_SECURE } from '$env/static/public';

export const handle: Handle = async ({ event, resolve }) => {
  // internal api request
  if (event.url.pathname.startsWith('/internal')) {
    let url = new URL(event.url.pathname, config.apiBaseUrl);
    url.search = event.url.search;

    console.log(event.url.pathname);
    // todo: better way of handling auth paths..?
    // todo: 405 method not allowed on POSTs
    if (event.url.pathname.startsWith('/internal')) {
      console.log(event.request.method);
      const result = await fetch(url, {
        ...event.request,
        headers: {
          Cookie: `sessionid=${event.cookies.get('sessionid')}`
        }
      });
      return result;
    } else {
      const result = await fetch(url, {
        ...event.request
      });
      return result;
    }
  }

  // check for session before making a request
  // set session to Promise<Session> if not
  if (!event.locals.session) {
    try {
      const { data } = await Client.GET(ApiPaths.get_session, {
        fetch: fetch,
        baseUrl: config.apiBaseUrl,
        headers: event.request.headers,
        credentials: 'include'
      });
      event.locals.session = data;
    } catch (error) {
      console.log('erorr', error);
    }
  }

  return await resolve(event, {
    filterSerializedResponseHeaders(name) {
      return name === 'content-length';
    }
  });
};
