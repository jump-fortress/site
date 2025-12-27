<script lang="ts">
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import DataSection from '$lib/components/DataSection.svelte';
  import Table from '$lib/components/table/Table.svelte';

  const activeBountiesDummyData = [
    {
      amount: '100 keys',
      map: 'jump_beef',
      class: 'soldier',
      goal: '00:55.00',
      placedby: 'hiei, nvg',
      date: '10/22/25'
    }
  ];

  const claimedBountiesDummyData = [
    {
      amount: '100 keys',
      claimedby: 'vice',
      map: 'jump_beef',
      class: 'soldier',
      run: '00:54.23',
      goal: '00:55.00',
      placedby: 'hiei, nvg',
      date: '10/22/25'
    }
  ];
</script>

<DataSection title={'Active Bounties'}>
  <Table data={activeBountiesDummyData}>
    {#snippet header()}
      <th class="w-12"></th>
      <th class="w-1/10">amount</th>
      <th>map</th>
      <th class="w-1/10">target time</th>
      <th class="w-1/6">placed by</th>
      <th class="w-1/10">date active</th>
    {/snippet}
    {#snippet row(d)}
      <td><img src={rocket} class="m-auto size-8" alt="" /></td>
      <td class="text-shadow-ctp-lavender text-shadow-xs">{d.amount}</td>
      <td class="relative">{@render TableMap(d.map)}</td>
      <td>{d.goal}</td>
      <td>{d.placedby}</td>
      <td>{d.date}</td>
    {/snippet}
  </Table>
</DataSection>

<DataSection title="Claimed Bounties">
  <Table data={claimedBountiesDummyData}>
    {#snippet header()}
      <th class="w-12"></th>
      <th class="w-1/10">amount</th>
      <th class="w-1/6">claimed by</th>
      <th>map</th>
      <th class="w-1/8">run time</th>
      <th class="w-1/10">target time</th>
      <th class="w-1/6">placed by</th>
      <th class="w-1/10">date active</th>
    {/snippet}
    {#snippet row(d)}
      <td><img src={rocket} class="m-auto size-8" alt="" /></td>
      <td class="text-shadow-ctp-lavender text-shadow-xs">{d.amount}</td>
      <td class="text-ctp-sky">{d.claimedby}</td>
      <td class="relative">{@render TableMap(d.map)}</td>
      <td>
        <div class="flex items-center justify-center gap-1">
          <span>{d.run}</span>
          {@render TestLink()}
        </div>
      </td>
      <td>{d.goal}</td>
      <td>{d.placedby}</td>
      <td>{d.date}</td>
    {/snippet}
  </Table>
</DataSection>

{#snippet TableMap(map: string)}
  <span class="span-ellipsis relative z-10">{map}</span>
  <div
    style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}
    class="absolute top-0 size-full mask-x-from-80% bg-cover bg-right bg-no-repeat opacity-50">
  </div>
{/snippet}

{#snippet TestLink()}
  <a href="/ababa" title="link" class="opacity-50 transition-opacity hover:opacity-100">
    <span class="icon-[ri--link-m] size-6"></span></a>
{/snippet}
