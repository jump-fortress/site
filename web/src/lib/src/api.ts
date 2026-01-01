import createClient from 'openapi-fetch';

import type {
  operations,
  paths,
  PlayerProfile,
  PlayerRequestPreview,
  PlayerWithRequest
} from '../schema';

export const Client = createClient<paths>({
  baseUrl: 'http://localhost:5173/'
});

// returns response.ok, since an error wouldn't be used
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

// returns response.ok, since an error wouldn't be used
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

export async function updatePreferredMap(selectedMap: string): Promise<InputResponse> {
  const { error } = await Client.PUT('/internal/players/preferredmap/{map}', {
    fetch: fetch,
    params: {
      path: {
        map: selectedMap
      }
    }
  });

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updateTempusID(tempusId: number): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updateSteamTradeToken(url: string): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updateSteamAvatar(): Promise<InputResponse> {
  const { error } = await Client.PUT('/internal/players/steamavatarurl', {
    fetch: fetch
  });

  if (error) {
    return {
      error: true,
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updatePlayerDisplayName(id: string, name: string): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
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
): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function updatePlayerDemoDivision(
  id: string,
  division: string
): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'updated'
  };
}

export async function createPlayerRequest(type: string, request: string): Promise<InputResponse> {
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
      message: error.detail ?? 'unknown error'
    };
  }

  return {
    error: false,
    message: 'request submitted'
  };
}

export async function getSelfPlayerRequests(): Promise<PlayerRequestPreview[] | null> {
  const { data } = await Client.GET('/internal/players/requests', {
    fetch: fetch
  });

  return data ?? null;
}

export async function getAllPlayerRequests(): Promise<PlayerWithRequest[] | null> {
  const { data } = await Client.GET('/internal/consultant/players/requests/pending', {
    fetch: fetch
  });

  return data ?? null;
}

export async function getAllMapNames(): Promise<string[] | null> {
  const { data } = await Client.GET('/internal/maps/names', {
    fetch: fetch
  });

  return data ?? null;
}

// returns blank message on no error
export async function resolvePlayerRequest(id: number): Promise<InputResponse> {
  const { error } = await Client.PUT('/internal/moderator/players/requests/resolve/{id}', {
    fetch: fetch,
    params: {
      path: {
        id: id
      }
    }
  });
  return error
    ? { error: true, message: error.detail ?? 'unknown error' }
    : { error: false, message: '' };
}
