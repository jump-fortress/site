// See https://svelte.dev/docs/kit/types#app.d.ts

import type { FullPlayer, PlayerProfile, Session } from '$lib/schema';

// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      session: Session | null;
    }
    interface PageData {
      session?: Session | null;
      playerProfile?: PlayerProfile;
      players?: PlayerProfile[];
      fullPlayer?: FullPlayer;
      fullPlayers?: FullPlayer[];
    }
    // interface PageState {}
    // interface Platform {}
  }

  interface InputError {
    error: boolean;
    message: string;
  }
}

export {};
