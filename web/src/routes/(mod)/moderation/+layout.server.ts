import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { Client } from '$lib/internalApi';
import type { FullPlayer } from '$lib/schema';

// todo: is it okay to run this fetch on the server, since we need to verify
//       a player's role first?
export const load: LayoutServerLoad = async ({ fetch, locals }) => {
  if (!locals.session) {
    redirect(302, '/');
  }

  if (locals.session.role === 'Admin' || locals.session.role === 'Mod') {
    // todo: error handling
    const { data } = await Client.GET('/internal/moderator/players/all', {
      fetch: fetch
    });

    return {
      fullPlayers: data as FullPlayer[]
    };
  } else {
    redirect(302, '/');
  }
};
