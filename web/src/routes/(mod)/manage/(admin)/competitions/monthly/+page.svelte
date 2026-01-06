<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import ClassImage from '$lib/components/display/ClassImage.svelte';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableCompetition from '$lib/components/display/table/TableCompetition.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import SelectButton from '$lib/components/input/select/SelectButton.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { createMonthly } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap } from 'svelte/reactivity';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Monthly } from '$lib/schema.js';

  let { data } = $props();

  // [division: map]
  let competitionDivisions: SvelteMap<string, string> = new SvelteMap();

  // todo: prizes
  const nowDate = Temporal.Now.plainDateISO();

  // placeholder values for creation
  const defaultMonthly = {
    id: 0,
    competition: {
      id: 0,
      class: 'Soldier',
      starts_at: `${nowDate.add({ days: 7 }).toString()}T00:00:00Z`,
      ends_at: '',
      created_at: '',
      visible_at: `${nowDate.add({ days: 1 }).toString()}T00:00:00Z`,
      complete: false
    },
    divisions: Array.from(competitionDivisions, ([division, map]) => ({
      division,
      map,
      id: 0,
      competition_id: 0
    }))
  };

  let competitionDate = $state('');
  let competitionTime = $state('00:00');
  let competitionDatetime = $derived(`${competitionDate}T${competitionTime}:00Z`);

  let visibleDate = $state('');
  let visibleTime = $state('00:00');
  let visibleDatetime = $derived(`${visibleDate}T${visibleTime}:00Z`);

  let monthly: Monthly = $state(defaultMonthly);

  let create = $state(false);
  let edit = $state(false);
</script>

{#await data.maps then maps}
  {#if maps}
    <div class="flex flex-col gap-1">
      {#each competitionDivisions as [div, _]}
        <div class="flex items-center gap-2">
          <div class="flex w-24 justify-center">
            <DivisionTag {div} />
          </div>
          <Label label="">
            <Select
              type="text"
              options={maps}
              onsubmit={async (value) => {
                competitionDivisions.set(div, value);
                return null;
              }} />
          </Label>
        </div>
      {/each}
    </div>
  {:else}
    <span>couldn't load any maps</span>
  {/if}
{/await}

<div class="relative -top-8">
  <MonthlyHeader header={true} {monthly} />
</div>

{#if create}
  <div in:slide>
    <Section label="create">
      <div class="flex items-center gap-2">
        <SelectButton
          label=""
          options={[
            { src: soldier, value: 'Soldier' },
            { src: demo, value: 'Demo' }
          ]}
          selected={'Soldier'}
          onsubmit={async (value) => {
            monthly.competition.class = value;
            return null;
          }} />

        <div class="flex flex-col">
          <Label label="start date" width={'max-w-60'}>
            <Input
              type="date"
              onsubmit={async (value) => {
                competitionDate = value;
                monthly.competition.starts_at = competitionDatetime;
                return null;
              }} />
          </Label>
          <Label label="start time" width={'max-w-60'}>
            <Input
              type="time"
              onsubmit={async (value) => {
                competitionTime = value;
                monthly.competition.starts_at = competitionDatetime;
                return null;
              }} />
          </Label>
        </div>

        <div class="flex flex-col">
          <Label label="visible date" width={'max-w-60'}>
            <Input
              type="date"
              onsubmit={async (value) => {
                visibleDate = value;
                monthly.competition.visible_at = visibleDatetime;
                return null;
              }} />
          </Label>
          <Label label="visible time" width={'max-w-60'}>
            <Input
              type="time"
              onsubmit={async (value) => {
                visibleTime = value;
                monthly.competition.visible_at = visibleDatetime;
                return null;
              }} />
          </Label>
        </div>

        <div class="ml-auto flex flex-col">
          <span>
            <span class="text-primary">input timezone:</span>
            UTC
          </span>
          <span>
            <span class="text-primary">your timezone:</span>
            {Temporal.Now.timeZoneId()}
          </span>
          <span>
            <span class="text-primary">your offset:</span>
            {Temporal.Now.zonedDateTimeISO().offset} UTC
          </span>
        </div>
      </div>

      <Label label="add division">
        <Select
          type="text"
          options={divisions}
          onsubmit={async (value) => {
            competitionDivisions.set(value, '');
            return null;
          }} />
      </Label>

      <Button
        onsubmit={() => {
          const monthlyValue = monthly;
          monthlyValue.competition.ends_at = monthlyValue.competition.starts_at;
          monthlyValue.competition.created_at = monthlyValue.competition.starts_at;
          return createMonthly(monthlyValue);
        }}>submit monthly</Button>
    </Section>
  </div>
{/if}

{#if edit}
  <div in:slide>
    <span>edit mode!</span>
  </div>
{/if}

<Button
  onsubmit={async () => {
    create = true;
    edit = false;
    return null;
  }}>create monthly</Button>

{#await data.monthlies then monthlies}
  {#if monthlies}
    <Table data={monthlies}>
      {#snippet header()}
        <th class="w-16">id</th>
        <th></th>
        <th class="w-24">start date</th>
        <th class="w-32">visible date</th>
        <th class="w-24">complete</th>
      {/snippet}
      {#snippet row(m: Monthly)}
        <td>{m.competition.id}</td>
        <td
          onclick={() => {
            create = false;
            edit = true;
            monthly = m;
          }}><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
        <td><TableDate date={m.competition.starts_at} /></td>
        <td><TableDate date={m.competition.visible_at} /></td>
        <td>{m.competition.complete}</td>
      {/snippet}
    </Table>
  {/if}
{/await}
