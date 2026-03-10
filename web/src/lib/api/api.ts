import type { paths } from '$lib/schema';
import { PUBLIC_JUMP_WEB_URL } from '$env/static/public';
import createClient from 'openapi-fetch';

export const Client = createClient<paths>({
  baseUrl: PUBLIC_JUMP_WEB_URL
});
