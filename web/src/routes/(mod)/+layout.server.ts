import { redirect } from '@sveltejs/kit';

import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ parent }) => {
  const session = await (await parent()).session;
  if (!session) {
    redirect(302, '/');
  }
  if (
    session.role === 'Admin' ||
    session.role === 'Mod' ||
    session.role === 'Consultant' ||
    session.role === 'Treasurer'
  ) {
    return;
  } else {
    redirect(302, '/');
  }
};
