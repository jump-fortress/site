import { Client } from '$lib/internalApi';
import type { PlayerProfile, Session } from '$lib/schema';
import type { PageLoad } from './$types';

// parent is used here since +page.ts cannot access locals
export const load: PageLoad = async ({ parent, fetch }) => {
	// todo: error handling
	const session = (await parent()).session as Session;

	const { data } = await Client.GET('/internal/players/profile/{id}', {
		fetch: fetch,
		params: {
			path: { id: session.id }
		}
	});

	if (!data) {
		return { player: null };
	}

	return {
		player: data as PlayerProfile
	};
};
