<script lang="ts">
  import { Client } from '$lib/api/api';
  import Div from '$lib/components/display/Div.svelte';
  import EventHeader from '$lib/components/display/EventHeader.svelte';
  import TableEvents from '$lib/components/display/table/presets/TableEvents.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Errors from '$lib/components/input/Errors.svelte';
  import Select from '$lib/components/input/Select.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { divs } from '$lib/helpers/sort';
  import { datetimeToMs, validDateTime } from '$lib/helpers/temporal';
  import { ApiPaths, type Event, type EventWithLeaderboards, type Leaderboard } from '$lib/schema';
  import { Temporal } from 'temporal-polyfill';

  let event_id: number = $state(0);
  let player_class: 'Soldier' | 'Demo' = $state('Soldier');
  let event_kind: string = $state('event');
  let event_kind_id: number = $state(0);
  let visible_date: string = $state('');
  let visible_time: string = $state('');
  let start_date: string = $state('');
  let start_time: string = $state('');

  const visible_at: string = $derived(validDateTime(`${visible_date}T${visible_time}:00Z`));
  const starts_at: string = $derived(validDateTime(`${start_date}T${start_date}:00Z`));

  const divless: Leaderboard[] = $derived.by(() => {
    return [
      {
        id: 0,
        event_id: event_id,
        map: ''
      }
    ];
  });

  let leaderboards: Leaderboard[] = $state([]);

  // key to fetch updated events
  let reloadEvents: boolean = $state(true);
  let mode: 'create' | 'update' = $state('create');
  let oerror: OpenAPIError = $state(undefined);

  // update event with the selected event
  function loadEvent({ event: e, leaderboards: l }: EventWithLeaderboards): void {
    event_id = e.id;
    player_class = e.player_class;
    event_kind = e.kind;
    event_kind_id = e.kind_id;
    visible_date = e.visible_at.substring(0, e.visible_at.indexOf('T'));
    visible_time = e.visible_at.substring(
      e.visible_at.indexOf('T') + 1,
      e.visible_at.indexOf('Z') - 3
    );
    start_date = e.starts_at.substring(0, e.starts_at.indexOf('T'));
    start_time = e.starts_at.substring(e.starts_at.indexOf('T') + 1, e.starts_at.indexOf('Z') - 3);
    event.starts_at = e.starts_at;
    event.ends_at = e.ends_at;
    leaderboards = l ?? divless;
    leaderboards[0]!.event_id = e.id;
  }

  const event = $derived.by(() => {
    const event: Event = {
      id: event_id,
      kind: event_kind as Event['kind'],
      kind_id: event_kind_id,
      player_class: player_class,
      visible_at: visible_at,
      starts_at: starts_at,
      ends_at: starts_at,
      created_at: starts_at
    };
    return event;
  });
</script>

<Section label="update">
  <EventHeader event={{ event: event, leaderboards: leaderboards }} />

  <Errors {oerror} />

  <span class="text-content/75"
    >leaderboards can be edited, but not removed. if you have too many divisions, please delete the
    event.</span>
  {#await Client.GET(ApiPaths.get_maps)}
    <span></span>
  {:then { data: maps }}
    {#if maps && mode === 'update'}
      <div class="flex flex-col gap-1">
        {#each leaderboards as l, i}
          <div class="flex">
            <Select
              label="div"
              type="text"
              options={divs}
              placeholder={leaderboards[i]?.div}
              onsubmit={async (value) => {
                leaderboards[i]!.div = value;
                return true;
              }} />
            <Select
              label="map"
              type="text"
              placeholder={l.map}
              options={maps.map((m) => m.name)}
              onsubmit={async (value) => {
                leaderboards[i]!.map = value;
                return true;
              }} />
          </div>
        {/each}

        <!-- buttons -->
        <div class="flex gap-1">
          {#if leaderboards.length > 0 && leaderboards.length < divs.length && leaderboards[0]?.div}
            <Button
              onsubmit={async () => {
                leaderboards = leaderboards.concat(divless);
                return true;
              }}>add div</Button>
          {/if}
          <Button
            onsubmit={async () => {
              const resp = await Client.POST(ApiPaths.update_leaderboards, {
                body: leaderboards
              });
              oerror = resp.error;
              // reset, since leaderboard IDs can't be updated
              leaderboards = divless;
              if (resp.response.ok) {
                reloadEvents = !reloadEvents;
              }
              return resp.response.ok;
            }}>
            <span>update</span>
          </Button>
        </div>
      </div>
    {/if}
  {/await}
</Section>
<Section label={'updatable events'}>
  {#key reloadEvents}
    {#await Client.GET(ApiPaths.get_full_events)}
      <span></span>
    {:then { data: ewls }}
      {@const now = Temporal.Now.instant().epochMilliseconds}
      {@const editable = ewls?.filter(({ event }) => datetimeToMs(event.starts_at) < now) ?? []}
      <TableEvents
        data={editable}
        onclick={(ewl) => {
          mode = 'update';
          loadEvent(ewl);
        }}>
      </TableEvents>
    {/await}
  {/key}
</Section>
