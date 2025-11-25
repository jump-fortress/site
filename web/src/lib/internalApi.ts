// todo: review is this is acceptable:
// use endpoint functions here only on non-page-load

import type {
	paths,
	PathsInternalPlayersPreferredclassClassPutParametersPathClass
} from '$lib/schema';
import createClient from 'openapi-fetch';

export const Client = createClient<paths>({
	baseUrl: 'http://localhost:5173/'
});

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
