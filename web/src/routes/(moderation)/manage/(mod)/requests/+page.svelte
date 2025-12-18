<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData } from './$types';
  import Table from '$lib/components/table/Table.svelte';
  import type { Player, PlayerWithRequest, Session } from '$lib/schema';
  import PlayerPreview from '$lib/components/PlayerPreview.svelte';
  import { slide } from 'svelte/transition';
  import { resolvePlayerRequest } from '$lib/internalApi';
  import CheckButton from '$lib/components/input/CheckButton.svelte';
  import ManagePlayer from '$lib/components/ManagePlayer.svelte';
  import { onMount } from 'svelte';

  let { data }: { data: PageData } = $props();

  let session = $derived(data.session as Session);

  // requests are initialized with load() data, then removed with successful fulfillments
  let requests: PlayerWithRequest[] | null = $state(null);
  let selectedPlayer: Player | null = $state(null);
  onMount(async () => {
    requests = (await data.playersWithRequests) as PlayerWithRequest[];
  });
</script>

{#if selectedPlayer}
  <div in:slide>
    <DataSection title={'Manage Player'}>
      <ManagePlayer player={selectedPlayer} role={session.role} />
    </DataSection>
  </div>
{/if}

<DataSection title={'Pending Requests'}>
  {#if requests}
    <Table data={requests}>
      {#snippet header()}
        <th class="w-48">request</th>
        <th class="w-32">content</th>
        <th></th>
        <th class="w-division">soldier</th>
        <th class="w-division">demo</th>
        <th class="w-date">date</th>
        <th class="w-0"></th>
      {/snippet}
      {#snippet row({ request, player }: PlayerWithRequest)}
        <td>{request.request_type}</td>
        <td>{request.request_string}</td>
        <td
          onclick={() => {
            selectedPlayer = player;
          }}
          class="td-hover-preview"
        >
          <PlayerPreview {player} />
        </td>
        <td class="text-division-{player.soldier_division?.toLowerCase()}"
          >{player.soldier_division}</td
        >
        <td class="text-division-{player.demo_division?.toLowerCase()}">{player.demo_division}</td>
        <td>{Temporal.Instant.from(request.created_at).toLocaleString()}</td>
        <td
          ><CheckButton
            onSelect={async () => {
              const response = await resolvePlayerRequest(request.id);
              if (!response.error) {
                requests = requests!.filter(({ request: r }) => r.id !== request.id);
              }
              return response;
            }}
          />
        </td>
      {/snippet}
    </Table>
  {:else}
    <span>no requests!</span>
  {/if}
</DataSection>
