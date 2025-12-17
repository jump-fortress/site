// See https://svelte.dev/docs/kit/types#app.d.ts

import type {
  Player,
  PlayerPreview,
  PlayerProfile,
  PlayerRequest,
  PlayerRequestPreview,
  PlayerWithRequest,
  Session
} from '$lib/schema';

// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      session: Session | null;
    }
    interface PageData {
      session?: Session | null;
      player?: Promise<Player>;
      players?: Promise<Players[]>;
      playerPreview?: Promise<PlayerPreview>;
      playerPreviews?: Promise<PlayerPreview[]>;
      playerProfile?: Promise<PlayerProfile>;
      requestPreviews?: Promise<PlayerRequestPreview[]>;
      requests?: Promise<PlayerRequest[]>;
      playersWithRequests?: Promise<PlayerWithRequest[]>;
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
