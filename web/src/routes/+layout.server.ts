import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  if (!locals.session) {
    return {
      session: null
    };
  }
  return {
    session: locals.session
  };
};
