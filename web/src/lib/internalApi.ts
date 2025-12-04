// todo: review is this is acceptable:
// use endpoint functions here only on non-page-load
import createClient from 'openapi-fetch';
import type { operations, paths, PlayerProfile } from './schema';

export const Client = createClient<paths>({
  baseUrl: 'http://localhost:5173/'
});

export async function updatePreferredClass(selectedClass: string): Promise<boolean> {
  const { response } = await Client.PUT('/internal/players/preferredclass/{class}', {
    fetch: fetch,
    params: {
      path: {
        class:
          selectedClass as operations['set-player-preferredclass']['parameters']['path']['class']
      }
    }
  });

  return response.ok;
}

export async function updatePreferredLauncher(selectedLauncher: string): Promise<boolean> {
  const { response } = await Client.PUT('/internal/players/preferredlauncher/{launcher}', {
    fetch: fetch,
    params: {
      path: {
        launcher:
          selectedLauncher as operations['set-player-preferredlauncher']['parameters']['path']['launcher']
      }
    }
  });

  return response.ok;
}

export async function updateTempusID(tempusId: number): Promise<InputError> {
  const { error } = await Client.PUT('/internal/players/tempusid/{tempus_id}', {
    fetch: fetch,
    params: {
      path: {
        tempus_id: tempusId
      }
    }
  });

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'Unknown error'
    };
  }

  return {
    error: false,
    message: 'success!'
  };
}

export async function updateSteamTradeToken(url: string): Promise<InputError> {
  const { error } = await Client.PUT('/internal/players/steamtradetoken/{url}', {
    fetch: fetch,
    params: {
      path: {
        url: url
      }
    }
  });

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'Unknown error'
    };
  }

  return {
    error: false,
    message: 'success!'
  };
}

export async function getPlayerProfile(id: string): Promise<PlayerProfile | null> {
  const { data } = await Client.GET('/internal/players/profile/{id}', {
    fetch: fetch,
    params: {
      path: {
        id: id
      }
    }
  });

  if (data) {
    return data as PlayerProfile;
  }

  return null;
}
