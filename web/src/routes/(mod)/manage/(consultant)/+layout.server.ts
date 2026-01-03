import { redirect } from '@sveltejs/kit';

import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  if (!locals.session) {
    redirect(302, '/');
  } else if (
    locals.session.role === 'Consultant' ||
    locals.session.role === 'Treasurer' ||
    locals.session.role === 'Mod' ||
    locals.session.role === 'Admin'
  ) {
    return;
  } else {
    redirect(302, '/');
  }
};
