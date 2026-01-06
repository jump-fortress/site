import { redirect } from '@sveltejs/kit';

import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  const session = await locals.session;
  if (!session) {
    redirect(302, '/');
  }

  if (session.role === 'Admin') {
    return;
  } else {
    redirect(302, '/');
  }
};
