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
  import SelectDivisions from '$lib/components/input/select/SelectDivisions.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { createMonthly } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap, SvelteSet } from 'svelte/reactivity';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Monthly } from '$lib/schema.js';

  let { data } = $props();
  // todo: prizes

  // [division: map]
  let competitionDivisions: SvelteMap<string, string> = $state(new SvelteMap());

  const nowDate = Temporal.Now.plainDateISO();
  const placeholderMonthly = {
    id: 0,
    competition: {
      id: 0,
      class: 'Soldier',
      starts_at: `${nowDate.add({ days: 7 }).toString()}T00:00:00Z`,
      ends_at: '',
      created_at: '',
      visible_at: `${nowDate.add({ days: 2 }).toString()}T00:00:00Z`,
      complete: false
    },
    divisions: []
  };

  function cdToArray(cd: SvelteMap<string, string>): Monthly['divisions'] {
    return Array.from(cd, ([division, map]) => ({
      division,
      map,
      id: 0,
      competition_id: 0
    }));
  }

  // state for a newly created competition
  let monthly: Monthly = $state(placeholderMonthly);

  let competitionDate = $state('');
  let competitionTime = $state('');
  let competitionDatetime = $derived(`${competitionDate}T${competitionTime}:00Z`);

  let visibleDate = $state('');
  let visibleTime = $state('');
  let visibleDatetime = $derived(`${visibleDate}T${visibleTime}:00Z`);

  let selectedDivisions: SvelteSet<string> = $state(new SvelteSet());

  let create = $state(false);
  let edit = $state(false);
</script>

<div class="relative -top-4">
  <MonthlyHeader header={true} {monthly} />
</div>

{#if create || edit}
  <div transition:slide>
    <Section label={create ? 'create' : 'edit'}>
      <div class="flex items-center gap-2">
        <SelectButton
          label=""
          options={[
            { src: soldier, value: 'Soldier' },
            { src: demo, value: 'Demo' }
          ]}
          selected={monthly.competition.class}
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

      <SelectDivisions
        label="set divisions"
        selected={selectedDivisions}
        onsubmit={(values) => {
          selectedDivisions = values;
          for (const division of divisions) {
            if (!values.has(division)) {
              competitionDivisions.delete(division);
            } else if (!competitionDivisions.has(division)) {
              competitionDivisions.set(division, '');
            }
          }
          monthly.divisions = cdToArray(competitionDivisions);
        }} />

      {#await data.maps then maps}
        {#if maps && maps.length !== 0}
          <div class="flex flex-col">
            {#each competitionDivisions as [div, map]}
              <div class="flex items-center gap-2">
                <div class="flex w-24 justify-center">
                  <DivisionTag {div} />
                </div>
                <Label label="">
                  <Select
                    type="text"
                    placeholder={map}
                    options={maps}
                    onsubmit={async (value) => {
                      monthly.divisions = cdToArray(competitionDivisions.set(div, value));
                      return null;
                    }} />
                </Label>
              </div>
            {/each}
          </div>
        {:else}
          <span>no maps found</span>
        {/if}
      {/await}

      <Button
        onsubmit={() => {
          const monthlyValue = monthly;
          monthlyValue.competition.ends_at = monthlyValue.competition.starts_at;
          monthlyValue.competition.created_at = monthlyValue.competition.starts_at;
          return createMonthly(monthlyValue);
        }}>{create ? 'create' : 'update'} monthly</Button>
    </Section>
  </div>
{/if}

<Button
  onsubmit={async () => {
    create = true;
    edit = false;
    monthly = placeholderMonthly;
    selectedDivisions.clear();
    competitionDivisions.clear();
    return null;
  }}>new</Button>

{#await data.monthlies then monthlies}
  {#if monthlies}
    <Table data={monthlies}>
      {#snippet header()}
        <th class="w-16">id</th>
        <th></th>
        <th class="w-32">visible date</th>
        <th class="w-32">start date</th>
        <th class="w-32">end date</th>
        <th class="w-24">complete</th>
      {/snippet}
      {#snippet row(m: Monthly)}
        <td>{m.competition.id}</td>
        <td
          onclick={() => {
            create = false;
            edit = true;
            monthly = m;
            competitionDivisions = new SvelteMap(m.divisions?.map((d) => [d.division, d.map]));
            selectedDivisions = new SvelteSet(m.divisions?.map((cd) => cd.division));
            monthly.divisions = cdToArray(competitionDivisions);
          }}><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
        <td><TableDate date={m.competition.visible_at} /></td>
        <td><TableDate date={m.competition.starts_at} /></td>
        <td><TableDate date={m.competition.ends_at} /></td>
        <td>{m.competition.complete}</td>
      {/snippet}
    </Table>
  {/if}
{/await}
