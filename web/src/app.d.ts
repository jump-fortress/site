// See https://svelte.dev/docs/kit/types#app.d.ts

import type { Session } from '$lib/schema';

// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      session: Promise<Session | undefined>;
    }
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
  }
}

export {};
