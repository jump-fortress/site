// See https://svelte.dev/docs/kit/types#app.d.ts

import type { FullPlayer, PlayerProfile, Session } from '$lib/schema';

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			session: Session;
		}
		interface PageData {
			session?: Session | null;
			player?: PlayerProfile;
			players?: PlayerProfile[];
			fullPlayers?: FullPlayer[];
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
