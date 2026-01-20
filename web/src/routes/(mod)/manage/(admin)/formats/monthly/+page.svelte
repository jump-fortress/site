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
  import { cancelCompetition, createMonthly, updateMapList, updateMonthly } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap, SvelteSet } from 'svelte/reactivity';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Monthly } from '$lib/schema.js';

  let { data } = $props();

  // [division: map]
  // used for easier updating of division <-> map name pairs instead of an array
  const competitionDivisions: SvelteMap<string, string> = new SvelteMap();

  // used only for division select component
  const selectedDivisions: SvelteSet<string> = new SvelteSet();

  const now = Temporal.Now.plainDateISO();
  const defaultMonthly = {
    id: 0,
    competition: {
      competition_type: 'Monthly',
      id: 0,
      class: 'Soldier',
      prizepool: 0,
      starts_at: `${now.add({ days: 7 }).toString()}T00:00:00Z`,
      ends_at: '',
      created_at: '',
      visible_at: `${now.add({ days: 2 }).toString()}T00:00:00Z`,
      complete: false
    },
    divisions: []
  };

  // needed for cdToArray(), when updating a competition's divisions
  let competitionId = $state(0);
  let monthlyId = $state(0);

  let competitionClass = $state('Soldier');
  let competitionDate = $state(now.add({ days: 7 }).toString());
  let competitionTime = $state('00:00');
  let visibleDate = $state(now.add({ days: 2 }).toString());
  let visibleTime = $state('00:00');

  // used for the competition preview & creation / updating
  let monthly: Monthly = $derived.by(() => {
    return {
      id: monthlyId,
      competition: {
        id: competitionId,
        competition_type: 'Monthly',
        class: competitionClass,
        prizepool: 0,
        starts_at: `${competitionDate}T${competitionTime}:00Z`,
        ends_at: '',
        created_at: '',
        visible_at: `${visibleDate}T${visibleTime}:00Z`,
        complete: false
      },
      divisions: cdToArray(competitionDivisions)
    };
  });

  monthly = defaultMonthly;

  // convert map to a competition.divisions array
  function cdToArray(cd: SvelteMap<string, string>): Monthly['divisions'] {
    return Array.from(cd, ([division, map]) => ({
      division,
      map,
      id: monthlyId,
      competition_id: competitionId
    }));
  }

  function loadMonthly(m: Monthly) {
    create = false;
    edit = true;
    monthly = m;

    competitionDivisions.clear();
    selectedDivisions.clear();
    for (let d of m.divisions ?? []) {
      competitionDivisions.set(d.division, d.map);
      selectedDivisions.add(d.division);
    }

    competitionId = m.competition.id;
    monthlyId = m.id;
    competitionClass = m.competition.class;

    const starts_at = Temporal.Instant.from(m.competition.starts_at).toZonedDateTimeISO('UTC');
    competitionDate = starts_at.toPlainDate().toString();
    competitionTime = starts_at.toPlainTime().toString().replace(':00', '');
  }

  let create = $state(false);
  let edit = $state(false);
</script>

<div class="relative -top-4">
  <MonthlyHeader header={false} {monthly} ends_at={false} />
</div>

{#if create || edit}
  <div transition:slide>
    <Section label={create ? 'create' : 'edit'}>
      <!-- times and class -->
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

      <!-- divisions -->
      <SelectDivisions
        label="set divisions"
        selected={selectedDivisions}
        onsubmit={(values) => {
          for (let v of values) {
            selectedDivisions.add(v);
          }
          for (const division of divisions) {
            if (!values.has(division)) {
              competitionDivisions.delete(division);
            } else if (!competitionDivisions.has(division)) {
              competitionDivisions.set(division, '');
            }
          }
        }} />

      <!-- maps, as divisions are selected -->
      {#await data.maps then maps}
        {#if maps && maps.length !== 0}
          <div class="flex flex-col">
            {#each competitionDivisions as [div, map]}
              <div class="flex items-center gap-1">
                <div class="flex w-24">
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

      <!-- create / edit buttons -->
      {#if create}
        <Button
          onsubmit={() => {
            const monthlyBody = monthly;

            // need to be valid datetimes for request body, but not used for creation / editing
            monthlyBody.competition.ends_at = monthlyBody.competition.starts_at;
            monthlyBody.competition.created_at = monthlyBody.competition.starts_at;

            return createMonthly(monthlyBody);
          }}>
          <span class="icon-[mdi--plus]"></span>
          <span>create monthly</span>
        </Button>
      {:else if edit}
        <Button
          onsubmit={() => {
            const monthlyBody = monthly;
            monthlyBody.competition.ends_at = monthlyBody.competition.starts_at;
            monthlyBody.competition.created_at = monthlyBody.competition.starts_at;

            return updateMonthly(monthlyBody);
          }}>
          <span class="icon-[mdi--edit]"></span>
          <span>update monthly</span>
        </Button>

        <Button
          onsubmit={() => {
            return cancelCompetition(monthly.competition.id);
          }}>
          <span class="icon-[mdi--close]"></span>
          <span>cancel monthly</span>
        </Button>
      {/if}
    </Section>
  </div>
{/if}

{#if !create && !edit}
  <div class="flex gap-2">
    <Button
      onsubmit={async () => {
        // reset state for monthly creation
        create = true;
        edit = false;
        monthly = defaultMonthly;
        selectedDivisions.clear();
        competitionDivisions.clear();
        return null;
      }}>
      <span class="icon-[mdi--plus]"></span>
      <span>new monthly</span>
    </Button>

    <Button onsubmit={() => updateMapList()}>
      <span class="icon-[mdi--arrow-up]"></span>
      <span>update map list</span>
    </Button>
  </div>
{/if}

<Section label="editable monthlies">
  {#await data.monthlies then monthlies}
    {#if monthlies}
      {@const nowInstant = Temporal.Now.instant()}
      <Table
        data={monthlies.filter(
          ({ competition }) =>
            nowInstant.until(Temporal.Instant.from(competition.starts_at)).total('minutes') > 0
        )}>
        {#snippet header()}
          <th class="w-12">id</th>
          <th></th>
          <th class="w-20">prizepool</th>
          <th class="w-24">visible</th>
          <th class="w-24">start</th>
          <th class="w-24">end</th>
          <th class="w-20">complete</th>
        {/snippet}
        {#snippet row(m: Monthly)}
          <td>{m.competition.id}</td>
          <td
            onclick={() => {
              loadMonthly(m);
            }}
            ><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
          <td>{m.competition.prizepool ?? 0} keys</td>
          <td><TableDate date={m.competition.visible_at} /></td>
          <td><TableDate date={m.competition.starts_at} /></td>
          <td><TableDate date={m.competition.ends_at} /></td>
          <td>{m.competition.complete}</td>
        {/snippet}
      </Table>
    {/if}
  {/await}
</Section>
