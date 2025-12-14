// todo: review is this is acceptable:
// use endpoint functions here only on non-page-load
import createClient from 'openapi-fetch';
import type { operations, paths, PlayerProfile, SelfPlayerRequest } from './schema';

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
  const { error } = await Client.PUT('/internal/players/tempusinfo/{tempus_id}', {
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
    message: 'updated'
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
    message: 'updated'
  };
}

export async function updateSteamAvatar(): Promise<InputError> {
  const { error } = await Client.PUT('/internal/players/steamavatarurl', {
    fetch: fetch
  });

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'Unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updatePlayerDisplayName(id: string, name: string): Promise<InputError> {
  const { error } = await Client.PUT('/internal/moderator/players/displayname/{id}/{name}', {
    fetch: fetch,
    params: {
      path: {
        id: id,
        name: name
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
    message: 'updated'
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

export async function updatePlayerSoldierDivision(
  id: string,
  division: string
): Promise<InputError> {
  const { error } = await Client.PUT(
    '/internal/moderator/players/soldierdivision/{id}/{division}',
    {
      fetch: fetch,
      params: {
        path: {
          id: id,
          division: division
        }
      }
    }
  );

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'Unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updatePlayerDemoDivision(id: string, division: string): Promise<InputError> {
  const { error } = await Client.PUT('/internal/moderator/players/demodivision/{id}/{division}', {
    fetch: fetch,
    params: {
      path: {
        id: id,
        division: division
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
    message: 'updated'
  };
}

export async function createPlayerRequest(type: string, request: string) {
  const { error } = await Client.PUT('/internal/players/requests/{request_type}/{request_string}', {
    fetch: fetch,
    params: {
      path: {
        request_type:
          type as operations['insert-player-request']['parameters']['path']['request_type'],
        request_string: request
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
    message: 'request submitted'
  };
}

export async function getPlayerRequests() {
  const { data } = await Client.GET('/internal/players/requests', {
    fetch: fetch
  });

  return data as SelfPlayerRequest[];
}
