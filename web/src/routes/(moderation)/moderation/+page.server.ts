import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({ locals }) => {
	if (!locals.session) {
		redirect(302, '/');
	}
	if (locals.session.role === 'admin' || locals.session.role === 'mod') {
		return {
			session: locals.session
		};
	} else {
		redirect(302, '/');
	}
};
