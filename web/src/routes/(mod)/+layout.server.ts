import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ parent }) => {
  const session = (await parent()).session;
  if (!session) {
    redirect(302, '/');
  }
  if (session.role === 'admin' || session.role === 'mod' || session.role === 'dev') {
    return;
  } else {
    redirect(302, '/');
  }
};
