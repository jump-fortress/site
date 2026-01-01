import { redirect } from '@sveltejs/kit';

import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  if (!locals.session) {
    redirect(302, '/');
  }
  if (
    locals.session.role === 'Admin' ||
    locals.session.role === 'Mod' ||
    locals.session.role === 'Consultant' ||
    locals.session.role === 'Treasurer'
  ) {
    return;
  } else {
    redirect(302, '/');
  }
};
