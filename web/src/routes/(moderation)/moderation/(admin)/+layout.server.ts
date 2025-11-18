import type { LayoutServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: LayoutServerLoad = ({ locals }) => {
	if (!locals.session) {
		redirect(302, '/');
	}
	if (locals.session.role === 'admin') {
		return {
			session: locals.session
		};
	} else {
		redirect(302, '/');
	}
};
