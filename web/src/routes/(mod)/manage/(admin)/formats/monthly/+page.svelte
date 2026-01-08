<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';
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
  import { createMonthly, updateMonthly } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap, SvelteSet } from 'svelte/reactivity';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Monthly } from '$lib/schema.js';

  let { data } = $props();
  // todo: prizes

  // [division: map]
  // used for easier updating of division <-> map name pairs instead of an array
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

  // convert map to a competition.divisions array
  function cdToArray(cd: SvelteMap<string, string>): Monthly['divisions'] {
    return Array.from(cd, ([division, map]) => ({
      division,
      map,
      id: monthlyId,
      competition_id: competitionId
    }));
  }

  // needed for cdToArray(), when updating a competition's divisions
  let competitionId = $state(0);
  let monthlyId = $state(0);

  let competitionClass = $state('Soldier');
  let competitionDate = $state(nowDate.add({ days: 7 }).toString());
  let competitionTime = $state('00:00');
  let visibleDate = $state(nowDate.add({ days: 2 }).toString());
  let visibleTime = $state('00:00');

  // used for the competition preview & creation / updating
  let monthly: Monthly = $derived({
    id: monthlyId,
    competition: {
      id: competitionId,
      class: competitionClass,
      starts_at: `${competitionDate}T${competitionTime}:00Z`,
      ends_at: '',
      created_at: '',
      visible_at: `${visibleDate}T${visibleTime}:00Z`,
      complete: false
    },
    divisions: cdToArray(competitionDivisions)
  });

  // initialize to defaults
  // note: assignments are made to reactive variables for monthly instead of monthly directly
  //       monthly doesn't seem to update without this, (due to $derived / deep state (being an object?)
  monthly = placeholderMonthly;

  // used only for division select component
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
            competitionClass = value;
            return null;
          }} />

        <div class="flex flex-col">
          <Label label="start date" width={'max-w-60'}>
            <Input
              type="date"
              onsubmit={async (value) => {
                competitionDate = value;
                return null;
              }} />
          </Label>
          <Label label="start time" width={'max-w-60'}>
            <Input
              type="time"
              onsubmit={async (value) => {
                competitionTime = value;
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
                return null;
              }} />
          </Label>
          <Label label="visible time" width={'max-w-60'}>
            <Input
              type="time"
              onsubmit={async (value) => {
                visibleTime = value;
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
                      competitionDivisions.set(div, value);
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

      {#if create}
        <Button
          onsubmit={() => {
            const monthlyBody = monthly;
            monthlyBody.competition.ends_at = monthlyBody.competition.starts_at;
            monthlyBody.competition.created_at = monthlyBody.competition.starts_at;
            return createMonthly(monthlyBody);
          }}>create monthly</Button>
      {:else if edit}
        <Button
          onsubmit={() => {
            const monthlyBody = monthly;
            monthlyBody.competition.ends_at = monthlyBody.competition.starts_at;
            monthlyBody.competition.created_at = monthlyBody.competition.starts_at;
            console.log(monthlyBody);
            return updateMonthly(monthlyBody);
          }}>update monthly</Button>
      {/if}
    </Section>
  </div>
{/if}

<Button
  onsubmit={async () => {
    create = true;
    edit = false;
    monthly = placeholderMonthly; // reset
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

            competitionId = m.competition.id;
            monthlyId = m.id;
            competitionClass = m.competition.class;

            const starts_at = Temporal.Instant.from(m.competition.starts_at).toZonedDateTimeISO(
              'UTC'
            );
            competitionDate = starts_at.toPlainDate().toString();
            competitionTime = starts_at.toPlainTime().toString().replace(':00', '');
          }}><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
        <td><TableDate date={m.competition.visible_at} /></td>
        <td><TableDate date={m.competition.starts_at} /></td>
        <td><TableDate date={m.competition.ends_at} /></td>
        <td>{m.competition.complete}</td>
      {/snippet}
    </Table>
  {/if}
{/await}
