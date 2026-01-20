// See https://svelte.dev/docs/kit/types#app.d.ts
import type {
  Session,
  Player,
  PlayerPreview,
  PlayerRequestPreview,
  PlayerRequest,
  PlayerWithRequest,
  Map,
  CompetitionDivisionTimes,
  PlayerWithPoints
} from '$lib/schema';

// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      session?: Promise<Session | undefined>;
    }
    interface PageData {
      session?: Promise<Session | undefined>;
      player?: Promise<Player>;
      players?: Promise<Player[]>;
      playerPreview?: Promise<PlayerPreview>;
      playersWithPoints?: Promise<PlayerWithPoints[]>;
      playerWithPoints?: Promise<PlayerWithPoints>;
      requestPreviews?: Promise<PlayerRequestPreview[]>;
      requests?: Promise<PlayerRequest[]>;
      playersWithRequests?: Promise<PlayerWithRequest[]>;
      maps?: Promise<Map[]>;
      monthly?: Promise<Monthly>;
      times?: Promise<CompetitionDivisionTimes[]>;
    }
    // interface PageState {}
    // interface Platform {}
  }

  // used for input components
  type InputResponse = {
    error: boolean;
    message: string;
  } | null;
}

export {};
