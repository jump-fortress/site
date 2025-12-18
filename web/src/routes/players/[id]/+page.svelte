<script lang="ts">
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import DataSection from '$lib/components/DataSection.svelte';
  import DivisionTag from '$lib/components/DivisionTag.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import TableMap from '$lib/components/TableMap.svelte';
  import type { Session } from '$lib/schema';
  import type { PageData } from './$types';
  import PlayerHeader from '$lib/components/PlayerHeader.svelte';
  import escape from '$lib/assets/static/maps/jump_escape_rc4.jpg';

  let { data }: { data: PageData } = $props();
</script>

{#await data.playerProfile}
  <span>waiting...</span>
{:then playerProfile}
  {#if playerProfile}
    <div
      class="absolute top-0 left-0 h-54 w-full mask-b-from-75% bg-cover bg-center opacity-50"
      style:background-image={`url(${escape})`}
    ></div>
    <PlayerHeader player={playerProfile.player} points={playerProfile.points} />
  {:else}
    <span>no player</span>
  {/if}
{/await}

<DataSection title="Bounties Claimed">
  <div class="grid w-full grid-cols-3 gap-2 text-base leading-5">
    {#each { length: 1 }}
      {@render testBounty()}
    {/each}
  </div>
</DataSection>

<DataSection title="Trophies">
  <div class="grid w-full grid-cols-4 gap-2 text-base leading-5">
    {#each { length: 1 }}
      {@render testTrophy()}
    {/each}
  </div>
</DataSection>

<DataSection title="Quests Claimed">
  <div class="grid w-full grid-cols-4 gap-2 text-base leading-5">
    {#each { length: 1 }}
      {@render testQuest()}
    {/each}
  </div>
</DataSection>

<DataSection title="Competition History">
  <Table data={[{}]}>
    {#snippet header()}
      <th class="w-12 text-ctp-lavender-50/75"></th>
      <th class="w-1/10 text-ctp-lavender-50/75"></th>
      <th class="w-1/8 text-ctp-lavender-50/75">format</th>
      <th class="text-ctp-lavender-50/75">map</th>
      <th class="w-1/10 text-ctp-lavender-50/75">time</th>
      <th class="w-1/16 text-ctp-lavender-50/75">place</th>
      <th class="w-1/12 text-ctp-lavender-50/75">points</th>
    {/snippet}
    {#snippet row()}
      <td><img src={rocket} class="m-auto size-8" alt="" /></td>
      <td><div class="flex justify-center"><DivisionTag div="platinum" /></div></td>
      <td class="cursor-pointer decoration-1 hover:underline">monthly #1</td>
      <td class="relative">
        <TableMap map={'jump_escape_rc4'} />
      </td>
      <td>3:53.00</td>
      <td>2nd</td>
      <td>9000</td>
    {/snippet}
  </Table>
</DataSection>

{#snippet testTrophy()}
  <div class="relative flex h-20 overflow-hidden bg-jfgray-900">
    <span class="absolute -left-10 z-0 icon-[ant-design--trophy-outlined] size-20 text-ctp-teal"
    ></span>
    <div class="absolute z-10 size-full p-1">
      <div class="ml-10 flex h-full flex-col justify-between">
        <div class="flex flex-col leading-4">
          <div class="flex items-center gap-2">
            <span class="text-lg text-ctp-yellow text-shadow-2xs text-shadow-ctp-yellow">1st</span>
            <span class="span-ellipsis">jump_escape_rc4</span>
          </div>
          <span class="font-semibold">3:53.00</span>
        </div>
        <span class="text-base opacity-75">monthly #1</span>
      </div>
    </div>
    <div
      style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg")`}
      class="size-full mask-l-from-50% mask-l-to-75% bg-cover bg-right bg-no-repeat opacity-50"
    ></div>
  </div>
{/snippet}

{#snippet testBounty()}
  <div class="relative flex h-28 items-center overflow-hidden bg-jfgray-900">
    <span
      class="absolute -left-12 z-0 icon-[ri--star-line] size-28 bg-linear-to-b from-ctp-lavender-950 to-ctp-lavender-50"
    ></span>
    <div class="absolute z-10 size-full p-2">
      <div class=" ml-8 flex h-full flex-col">
        <div class="flex items-center gap-1.5">
          <span class="text-lg text-shadow-ctp-lavender text-shadow-xs">100 keys</span>
          <span class="opacity-90">for sub 4:00</span>
        </div>
        <span class="span-ellipsis ml-8 text-base">jump_escape_rc4</span>
        <span class="mt-2 ml-8 text-xl font-semibold text-ctp-lavender">3:53.00</span>
      </div>
    </div>
    <div
      style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg")`}
      class="size-full mask-l-from-50% mask-l-to-75% bg-cover bg-right bg-no-repeat opacity-50"
    ></div>
  </div>
{/snippet}

{#snippet testQuest()}
  <div class="relative flex h-14 overflow-hidden bg-jfgray-900">
    <span class="absolute -left-6 z-0 icon-[ri--treasure-map-line] size-14 text-ctp-teal"></span>
    <div class="absolute z-10 size-full p-1">
      <div class="flex h-full flex-col">
        <div class="ml-10 flex flex-col">
          <span class="span-ellipsis">jump_escape_rc4</span>
          <span class="opacity-75">completion</span>
        </div>
      </div>
    </div>
    <div
      class="size-full bg-[url(https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg)] mask-l-from-50% mask-l-to-75% bg-cover bg-right bg-no-repeat opacity-50"
    ></div>
  </div>
{/snippet}
