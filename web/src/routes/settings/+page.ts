import { getPlayerProfile } from '$lib/internalApi';
import type { PlayerProfile, Session } from '$lib/schema';
import type { PageLoad } from './$types';

// parent is used here since +page.ts cannot access locals
export const load: PageLoad = async ({ parent }) => {
	// todo: error handling
	const session = (await parent()).session as Session;

	const player = await getPlayerProfile(session.id);
	return {
		player: player as PlayerProfile
	};
};
