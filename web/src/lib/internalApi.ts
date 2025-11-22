import type {
	paths,
	PathsInternalPlayersPreferredclassClassPutParametersPathClass,
	PlayerProfile
} from '$lib/schema';
import createClient from 'openapi-fetch';

export const Client = createClient<paths>({
	baseUrl: 'http://localhost:5173/'
});

export async function getPlayerProfile(id: number): Promise<PlayerProfile | null> {
	const { data } = await Client.GET('/internal/players/profile/{id}', {
		fetch: fetch,
		params: {
			path: { id: id }
		}
	});

	if (!data) {
		return null;
	}

	return data as PlayerProfile;
}

// todo: handle errors, use response
export async function updatePreferredClass(
	selectedClass: PathsInternalPlayersPreferredclassClassPutParametersPathClass
) {
	const { response } = await Client.PUT('/internal/players/preferredclass/{class}', {
		fetch: fetch,
		params: {
			path: { class: selectedClass }
		}
	});
}
